// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet)

// Description: Manage port assignment for FortiOS system virtual switches.

package framework

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	forticlient "github.com/terraform-providers/terraform-provider-fortios/sdk/sdkcore"
)

var _ resource.Resource = &resourceSystemVirtualSwitchPortAssignment{}
var _ resource.ResourceWithValidateConfig = &resourceSystemVirtualSwitchPortAssignment{}

func NewResourceSystemVirtualSwitchPortAssignment() resource.Resource {
	return &resourceSystemVirtualSwitchPortAssignment{}
}

type resourceSystemVirtualSwitchPortAssignment struct {
	client *forticlient.FortiSDKClient
}

func (r *resourceSystemVirtualSwitchPortAssignment) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_virtualswitch_port_assignment"
}

// Schema uses a DynamicAttribute for port_map so the map-keyed HCL can carry
// an optional per-switch "vdom" alongside the "ports" list. Protocol v5 (this
// provider is muxed at v5) does not support nested attributes, and a MapAttribute
// with an object element forces every object field to be required, so a dynamic
// value parsed with tftypes is used instead.
func (r *resourceSystemVirtualSwitchPortAssignment) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage port assignments for FortiOS system virtual switches. Does not create or delete virtual switches, only assigns ports.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"port_map": schema.DynamicAttribute{
				Required: true,
				MarkdownDescription: "Map of virtual switch name to an object with an optional `vdom` and a list of `ports` to assign. " +
					"Each port object has `name` and `alias`. If `vdom` is not specified, the provider-level vdom is used. " +
					"Example: `port_map = { \"switch1\" = { vdom = \"root\", ports = [{ name = \"port1\", alias = \"a1\" }] } }`.",
			},
		},
	}
}

func (r *resourceSystemVirtualSwitchPortAssignment) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, err := getFortiSDKClient(req.ProviderData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *FortiClient with Client field, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

// ValidateConfig checks the plan for port conflicts across the configured virtual
// switches before apply. A physical port can only be assigned to one virtual switch,
// so detecting duplicate port names at plan time avoids a failed apply.
func (r *resourceSystemVirtualSwitchPortAssignment) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resourceSystemVirtualSwitchPortAssignmentModel
	diags := &resp.Diagnostics

	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	portMap, d := extractPortMap(ctx, data.PortMap)
	diags.Append(d...)
	if diags.HasError() {
		return
	}

	if conflicts := validatePortAssignmentConflicts(portMap); conflicts != "" {
		diags.AddError("Port Conflict", conflicts)
	}
}

// resourceSystemVirtualSwitchPortAssignmentModel describes the resource data model.
type resourceSystemVirtualSwitchPortAssignmentModel struct {
	ID      types.String  `tfsdk:"id"`
	PortMap types.Dynamic `tfsdk:"port_map"`
}

// portItem holds a single port assignment: name and optional alias.
type portItem struct {
	Name  string
	Alias string
}

// switchEntry holds the parsed contents of a single port_map entry: the
// optional per-switch vdom and the list of port items.
type switchEntry struct {
	Vdom  string
	Ports []portItem
}

func (r *resourceSystemVirtualSwitchPortAssignment) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSystemVirtualSwitchPortAssignmentModel
	diags := &resp.Diagnostics

	diags.Append(req.Plan.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.client
	c.Retries = 1

	diags.Append(r.applyPortMap(ctx, &data)...)
	if diags.HasError() {
		return
	}

	data.ID = types.StringValue("system_virtualswitch_port_assignment")

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSystemVirtualSwitchPortAssignment) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resourceSystemVirtualSwitchPortAssignmentModel
	diags := &resp.Diagnostics

	diags.Append(req.State.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	diags.Append(r.refreshPortMap(ctx, &data)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSystemVirtualSwitchPortAssignment) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceSystemVirtualSwitchPortAssignmentModel
	diags := &resp.Diagnostics

	diags.Append(req.Plan.Get(ctx, &plan)...)
	if diags.HasError() {
		return
	}

	var state resourceSystemVirtualSwitchPortAssignmentModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	c := r.client
	c.Retries = 1

	newPortMap, d := extractPortMap(ctx, plan.PortMap)
	diags.Append(d...)
	if diags.HasError() {
		return
	}

	oldPortMap, d := extractPortMap(ctx, state.PortMap)
	diags.Append(d...)
	if diags.HasError() {
		return
	}

	// Phase 1: Clear ports from switches that have been removed from the plan
	// entirely (present in old state but absent from new). This must happen before
	// any adds so the freed ports can be reclaimed by other switches. The vdom is
	// taken from the old entry, since the new plan no longer describes this switch.
	for switchName, oldEntry := range oldPortMap {
		if _, ok := newPortMap[switchName]; ok {
			// Still present in the plan; handled in later phases.
			continue
		}

		if len(oldEntry.Ports) == 0 {
			// Nothing assigned to clear.
			continue
		}

		vdom := r.resolveVdom(oldEntry)
		current, err := c.ReadSystemVirtualSwitch(switchName, vdom)
		if err != nil || current == nil {
			// Virtual switch no longer exists, nothing to clear.
			continue
		}

		updateObj := make(map[string]interface{})
		updateObj["name"] = switchName
		if v, ok := current["vdom"]; ok && v != nil {
			updateObj["vdom"] = v
		}
		updateObj["port"] = make([]map[string]interface{}, 0)

		_, err = c.UpdateSystemVirtualSwitch(&updateObj, switchName, vdom)
		if err != nil {
			diags.AddError("Update Error", fmt.Sprintf("Error clearing ports from removed virtual switch %s: %v", switchName, err))
			return
		}
	}

	// Phase 2: For switches present in both old and new, do a removal-only update
	// first. This frees up ports that are being moved away from a switch so that
	// another switch can claim them in Phase 3 without a port conflict.
	for switchName, newEntry := range newPortMap {
		oldEntry, ok := oldPortMap[switchName]
		if !ok {
			// New switch has no old ports to remove.
			continue
		}

		oldPorts := getPortAssignmentPortSet(oldEntry.Ports)
		newPorts := getPortAssignmentPortSet(newEntry.Ports)
		removed, _ := portAssignmentPortDiff(oldPorts, newPorts)

		if len(removed) > 0 {
			vdom := r.resolveVdom(newEntry)
			removalObj := buildRemovalOnlyAPIObject(switchName, oldEntry.Ports, removed)
			_, err := c.UpdateSystemVirtualSwitch(&removalObj, switchName, vdom)
			if err != nil {
				diags.AddError("Update Error", fmt.Sprintf("Error removing ports from virtual switch %s: %v", switchName, err))
				return
			}
		}
	}

	// Phase 3: Full port-list update for every switch in the plan. Ports freed in
	// Phases 1 and 2 can now be added to their new switch. New switches (not in old
	// state) are handled here too.
	for switchName, entry := range newPortMap {
		vdom := r.resolveVdom(entry)

		current, err := c.ReadSystemVirtualSwitch(switchName, vdom)
		if err != nil {
			diags.AddError("Read Error", fmt.Sprintf("Error reading virtual switch %s: %v", switchName, err))
			return
		}

		updateObj := make(map[string]interface{})
		updateObj["name"] = switchName
		if v, ok := current["vdom"]; ok && v != nil {
			updateObj["vdom"] = v
		}
		updateObj["port"] = expandPortListToAPI(entry.Ports)

		_, err = c.UpdateSystemVirtualSwitch(&updateObj, switchName, vdom)
		if err != nil {
			diags.AddError("Update Error", fmt.Sprintf("Error updating ports for virtual switch %s: %v", switchName, err))
			return
		}
	}

	plan.ID = state.ID

	diags.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourceSystemVirtualSwitchPortAssignment) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Port assignment resource does not delete the virtual switches themselves.
	// On delete, remove all port assignments from the managed virtual switches.
	var data resourceSystemVirtualSwitchPortAssignmentModel
	diags := &resp.Diagnostics

	diags.Append(req.State.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	portMap, d := extractPortMap(ctx, data.PortMap)
	diags.Append(d...)
	if diags.HasError() {
		return
	}

	c := r.client

	for switchName, entry := range portMap {
		vdom := r.resolveVdom(entry)
		current, err := c.ReadSystemVirtualSwitch(switchName, vdom)
		if err != nil || current == nil {
			continue
		}

		updateObj := make(map[string]interface{})
		updateObj["name"] = switchName
		if v, ok := current["vdom"]; ok && v != nil {
			updateObj["vdom"] = v
		}
		updateObj["port"] = make([]map[string]interface{}, 0)

		_, err = c.UpdateSystemVirtualSwitch(&updateObj, switchName, vdom)
		if err != nil {
			diags.AddError("Update Error", fmt.Sprintf("Error clearing ports from virtual switch %s: %v", switchName, err))
			return
		}
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

// applyPortMap reads the plan and updates port assignments on all virtual switches.
func (r *resourceSystemVirtualSwitchPortAssignment) applyPortMap(ctx context.Context, data *resourceSystemVirtualSwitchPortAssignmentModel) diag.Diagnostics {
	var diags diag.Diagnostics

	portMap, d := extractPortMap(ctx, data.PortMap)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}

	c := r.client

	for switchName, entry := range portMap {
		vdom := r.resolveVdom(entry)

		current, err := c.ReadSystemVirtualSwitch(switchName, vdom)
		if err != nil {
			diags.AddError("Read Error", fmt.Sprintf("Error reading virtual switch %s: %v", switchName, err))
			return diags
		}

		updateObj := make(map[string]interface{})
		updateObj["name"] = switchName
		if v, ok := current["vdom"]; ok && v != nil {
			updateObj["vdom"] = v
		}
		updateObj["port"] = expandPortListToAPI(entry.Ports)

		_, err = c.UpdateSystemVirtualSwitch(&updateObj, switchName, vdom)
		if err != nil {
			diags.AddError("Update Error", fmt.Sprintf("Error updating ports for virtual switch %s: %v", switchName, err))
			return diags
		}
	}

	return diags
}

// refreshPortMap reads port assignments from the FortiOS API and updates the state.
// It preserves the exact tftypes type from the original value (including which
// attributes were present, e.g. an omitted "vdom") to avoid type mismatch and
// spurious diffs.
func (r *resourceSystemVirtualSwitchPortAssignment) refreshPortMap(ctx context.Context, data *resourceSystemVirtualSwitchPortAssignmentModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if data.PortMap.IsNull() || data.PortMap.IsUnknown() {
		return diags
	}

	// Get the original tftypes.Value to preserve its exact type.
	originalTfVal, err := data.PortMap.ToTerraformValue(ctx)
	if err != nil {
		diags.AddError("Conversion Error", fmt.Sprintf("cannot get terraform value: %v", err))
		return diags
	}

	originalType := originalTfVal.Type()

	// Extract switch names and their original values.
	var originalSwitches map[string]tftypes.Value
	switch originalType.(type) {
	case tftypes.Object:
		var objVals map[string]tftypes.Value
		if err := originalTfVal.As(&objVals); err != nil {
			diags.AddError("Conversion Error", fmt.Sprintf("cannot extract object values: %v", err))
			return diags
		}
		originalSwitches = objVals
	case tftypes.Map:
		var mapVals map[string]tftypes.Value
		if err := originalTfVal.As(&mapVals); err != nil {
			diags.AddError("Conversion Error", fmt.Sprintf("cannot extract map values: %v", err))
			return diags
		}
		originalSwitches = mapVals
	default:
		diags.AddError("Type Error", fmt.Sprintf("expected map or object for port_map, got %s", originalType.String()))
		return diags
	}

	// Build refreshed values preserving the original type.
	newSwitches := make(map[string]tftypes.Value)
	for name, originalSwitchVal := range originalSwitches {
		if originalSwitchVal.IsNull() {
			continue
		}

		// Extract vdom from original value to read the right switch.
		var vdom string
		var switchObj map[string]tftypes.Value
		if err := originalSwitchVal.As(&switchObj); err == nil {
			if vdomVal, ok := switchObj["vdom"]; ok && !vdomVal.IsNull() {
				var s string
				vdomVal.As(&s)
				vdom = s
			}
		}
		if vdom == "" && r.client != nil && r.client.Config.Auth.Vdom != "" {
			vdom = r.client.Config.Auth.Vdom
		}

		// Read from API.
		apiData, err := r.client.ReadSystemVirtualSwitch(name, vdom)
		if err != nil {
			// Switch no longer exists, skip it.
			continue
		}

		// Build refreshed tftypes.Value with the same type as original.
		refreshedVal := buildRefreshedSwitchValue(originalSwitchVal, apiData)
		newSwitches[name] = refreshedVal
	}

	// Reconstruct the top-level value with the original type.
	var newTfVal tftypes.Value
	switch originalType.(type) {
	case tftypes.Object:
		newTfVal = tftypes.NewValue(originalType.(tftypes.Object), newSwitches)
	case tftypes.Map:
		newTfVal = tftypes.NewValue(originalType.(tftypes.Map), newSwitches)
	}

	// Convert tftypes.Value to types.Dynamic.
	dynamicVal, err := types.DynamicType.ValueFromTerraform(ctx, newTfVal)
	if err != nil {
		diags.AddError("Conversion Error", fmt.Sprintf("cannot convert to dynamic: %v", err))
		return diags
	}

	data.PortMap = dynamicVal.(types.Dynamic)
	return diags
}

// buildRefreshedSwitchValue builds a new tftypes.Value for a single switch entry,
// preserving the exact attribute types from the original value. Only "ports" is
// refreshed from the API; "vdom" (and any other attribute) is preserved as-is
// from the original config.
func buildRefreshedSwitchValue(original tftypes.Value, apiData map[string]interface{}) tftypes.Value {
	if original.IsNull() {
		return original
	}

	originalType := original.Type().(tftypes.Object)

	var originalAttrs map[string]tftypes.Value
	_ = original.As(&originalAttrs)

	newAttrs := make(map[string]tftypes.Value)
	for attrName, attrType := range originalType.AttributeTypes {
		originalAttrVal, wasPresent := originalAttrs[attrName]
		if !wasPresent {
			// Attribute not in original config, preserve as null.
			newAttrs[attrName] = tftypes.NewValue(attrType, nil)
			continue
		}

		if originalAttrVal.IsNull() {
			// Was explicitly null in config, keep null.
			newAttrs[attrName] = tftypes.NewValue(attrType, nil)
			continue
		}

		switch attrName {
		case "vdom":
			// Preserve the configured vdom; do not overwrite from API.
			newAttrs[attrName] = originalAttrVal
		case "ports":
			newAttrs[attrName] = buildRefreshedPortValue(originalAttrVal, apiData)
		default:
			// Unknown attribute, preserve original value.
			newAttrs[attrName] = originalAttrVal
		}
	}

	return tftypes.NewValue(originalType, newAttrs)
}

// buildRefreshedPortValue builds a refreshed port list tftypes.Value from API data,
// preserving the original list/tuple element type.
func buildRefreshedPortValue(original tftypes.Value, apiData map[string]interface{}) tftypes.Value {
	if original.IsNull() {
		return original
	}

	originalType := original.Type()

	// Get the element type.
	var elemType tftypes.Type
	switch tt := originalType.(type) {
	case tftypes.List:
		elemType = tt.ElementType
	case tftypes.Tuple:
		if len(tt.ElementTypes) > 0 {
			elemType = tt.ElementTypes[0]
		}
	default:
		return original
	}

	portObjType, ok := elemType.(tftypes.Object)
	if !ok {
		return original
	}

	// Get API port data.
	var apiPorts []interface{}
	if v, ok := apiData["port"]; ok && v != nil {
		if ports, ok := v.([]interface{}); ok {
			apiPorts = ports
		}
	}

	// A Tuple has a fixed arity: it can only ever hold exactly len(ElementTypes)
	// elements. If the API returns a different number of ports (including zero),
	// the value cannot be represented in the original Tuple type, so preserve the
	// original value to avoid a panic / type mismatch. A List can hold any count.
	if tupleType, isTuple := originalType.(tftypes.Tuple); isTuple {
		if len(apiPorts) != len(tupleType.ElementTypes) {
			return original
		}
	}

	if len(apiPorts) == 0 {
		// Only reachable for List (Tuple is handled above). An empty list is valid.
		return tftypes.NewValue(originalType, []tftypes.Value{})
	}

	portElements := make([]tftypes.Value, 0, len(apiPorts))
	for _, item := range apiPorts {
		m, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		portAttrs := make(map[string]tftypes.Value)
		for attrName := range portObjType.AttributeTypes {
			switch attrName {
			case "name":
				if v, ok := m["name"]; ok && v != nil {
					portAttrs[attrName] = tftypes.NewValue(tftypes.String, v)
				} else {
					portAttrs[attrName] = tftypes.NewValue(tftypes.String, nil)
				}
			case "alias":
				if v, ok := m["alias"]; ok && v != nil {
					portAttrs[attrName] = tftypes.NewValue(tftypes.String, v)
				} else {
					portAttrs[attrName] = tftypes.NewValue(tftypes.String, nil)
				}
			default:
				portAttrs[attrName] = tftypes.NewValue(tftypes.String, nil)
			}
		}
		portElements = append(portElements, tftypes.NewValue(portObjType, portAttrs))
	}

	// For tuples, the element count must match the type definition.
	// If API returns a different count, we can't represent it in the original type.
	if tupleType, isTuple := originalType.(tftypes.Tuple); isTuple {
		if len(portElements) != len(tupleType.ElementTypes) {
			// Return original to avoid type mismatch crash.
			return original
		}
	}

	return tftypes.NewValue(originalType, portElements)
}

// extractPortMap parses a types.Dynamic containing a map/object of switch entries
// into a map[string]switchEntry. Each entry carries the optional per-switch vdom
// and the list of port items. Missing attributes (e.g. an omitted "vdom") are
// treated as absent, so vdom is optional.
func extractPortMap(ctx context.Context, dyn types.Dynamic) (map[string]switchEntry, diag.Diagnostics) {
	var diags diag.Diagnostics
	result := make(map[string]switchEntry)

	if dyn.IsNull() || dyn.IsUnknown() {
		return result, diags
	}

	tfVal, err := dyn.ToTerraformValue(ctx)
	if err != nil {
		diags.AddError("Conversion Error", fmt.Sprintf("cannot get terraform value: %v", err))
		return result, diags
	}

	tfType := tfVal.Type()

	// Collect the per-switch object values. Both a map and an object (whose
	// attribute names are the switch names) are accepted.
	var switchVals map[string]tftypes.Value
	switch tt := tfType.(type) {
	case tftypes.Map:
		_, isObject := tt.ElementType.(tftypes.Object)
		if !isObject {
			diags.AddError("Type Error", "port_map map elements must be objects")
			return result, diags
		}
		if err := tfVal.As(&switchVals); err != nil {
			diags.AddError("Conversion Error", fmt.Sprintf("cannot extract map values: %v", err))
			return result, diags
		}
	case tftypes.Object:
		if err := tfVal.As(&switchVals); err != nil {
			diags.AddError("Conversion Error", fmt.Sprintf("cannot extract object values: %v", err))
			return result, diags
		}
	default:
		diags.AddError("Type Error", fmt.Sprintf("expected map or object for port_map, got %s", tfType.String()))
		return result, diags
	}

	for name, entryVal := range switchVals {
		if entryVal.IsNull() {
			continue
		}

		entry, d := parseSwitchEntry(entryVal)
		diags.Append(d...)
		if !diags.HasError() {
			result[name] = entry
		}
	}

	return result, diags
}

// parseSwitchEntry converts a single switch entry tftypes.Value into a switchEntry.
func parseSwitchEntry(entryVal tftypes.Value) (switchEntry, diag.Diagnostics) {
	var diags diag.Diagnostics
	entry := switchEntry{}

	if _, ok := entryVal.Type().(tftypes.Object); !ok {
		diags.AddError("Type Error", fmt.Sprintf("expected object for port_map entry, got %s", entryVal.Type().String()))
		return entry, diags
	}

	var objMap map[string]tftypes.Value
	if err := entryVal.As(&objMap); err != nil {
		diags.AddError("Conversion Error", fmt.Sprintf("cannot extract port_map entry: %v", err))
		return entry, diags
	}

	// Optional per-switch vdom.
	if vdomVal, ok := objMap["vdom"]; ok && !vdomVal.IsNull() {
		var s string
		if err := vdomVal.As(&s); err != nil {
			diags.AddError("Conversion Error", fmt.Sprintf("cannot convert vdom to string: %v", err))
			return entry, diags
		}
		entry.Vdom = s
	}

	// ports list (optional; an empty/absent list means no ports).
	if portsVal, ok := objMap["ports"]; ok && !portsVal.IsNull() {
		ports, d := parsePortList(portsVal)
		diags.Append(d...)
		entry.Ports = ports
	}

	return entry, diags
}

// parsePortList converts a tftypes.Value (list/tuple of port objects) into []portItem.
func parsePortList(portsVal tftypes.Value) ([]portItem, diag.Diagnostics) {
	var diags diag.Diagnostics
	result := make([]portItem, 0)

	portsType := portsVal.Type()
	var listVals []tftypes.Value

	switch tt := portsType.(type) {
	case tftypes.List:
		if err := portsVal.As(&listVals); err != nil {
			diags.AddError("Conversion Error", fmt.Sprintf("cannot extract port list: %v", err))
			return result, diags
		}
	case tftypes.Tuple:
		if err := portsVal.As(&listVals); err != nil {
			diags.AddError("Conversion Error", fmt.Sprintf("cannot extract port tuple: %v", err))
			return result, diags
		}
	default:
		diags.AddError("Type Error", fmt.Sprintf("expected list or tuple for ports, got %s", tt.String()))
		return result, diags
	}

	for i, portVal := range listVals {
		if portVal.IsNull() {
			continue
		}

		if _, ok := portVal.Type().(tftypes.Object); !ok {
			diags.AddError("Type Error", fmt.Sprintf("port element %d is not an object", i))
			continue
		}

		var portMap map[string]tftypes.Value
		if err := portVal.As(&portMap); err != nil {
			diags.AddError("Conversion Error", fmt.Sprintf("cannot extract port %d: %v", i, err))
			continue
		}

		item := portItem{}
		if nameVal, ok := portMap["name"]; ok && !nameVal.IsNull() {
			var s string
			nameVal.As(&s)
			item.Name = s
		}
		if aliasVal, ok := portMap["alias"]; ok && !aliasVal.IsNull() {
			var s string
			aliasVal.As(&s)
			item.Alias = s
		}
		result = append(result, item)
	}

	return result, diags
}

// resolveVdom returns the vdom to use for a switch: the per-switch vdom when
// specified, otherwise the provider-level vdom.
func (r *resourceSystemVirtualSwitchPortAssignment) resolveVdom(entry switchEntry) string {
	if entry.Vdom != "" {
		return entry.Vdom
	}
	return r.getClientVdom()
}

// expandPortListToAPI converts a slice of port items into the API representation:
// name and alias are only included when they are not empty.
func expandPortListToAPI(ports []portItem) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(ports))
	for _, p := range ports {
		item := make(map[string]interface{})
		if p.Name != "" {
			item["name"] = p.Name
		}
		if p.Alias != "" {
			item["alias"] = p.Alias
		}
		result = append(result, item)
	}
	return result
}

// buildRemovalOnlyAPIObject builds an API object that keeps only the old ports
// minus the removed ones. No new ports are added. This frees up ports for other
// switches to claim.
func buildRemovalOnlyAPIObject(name string, oldPorts []portItem, removedPorts map[string]bool) map[string]interface{} {
	obj := make(map[string]interface{})
	obj["name"] = name

	keptPorts := make([]map[string]interface{}, 0)
	for _, p := range oldPorts {
		if removedPorts[p.Name] {
			continue
		}
		item := make(map[string]interface{})
		if p.Name != "" {
			item["name"] = p.Name
		}
		if p.Alias != "" {
			item["alias"] = p.Alias
		}
		keptPorts = append(keptPorts, item)
	}

	// Always send port (even empty) so FortiOS removes all ports when none are kept.
	obj["port"] = keptPorts

	return obj
}

// getPortAssignmentPortSet extracts port names from a slice of port items into a set.
func getPortAssignmentPortSet(ports []portItem) map[string]bool {
	result := make(map[string]bool)
	for _, p := range ports {
		if p.Name != "" {
			result[p.Name] = true
		}
	}
	return result
}

// portAssignmentPortDiff returns the removed and added ports between old and new sets.
func portAssignmentPortDiff(oldSet, newSet map[string]bool) (removed, added map[string]bool) {
	removed = make(map[string]bool)
	added = make(map[string]bool)

	for port := range oldSet {
		if !newSet[port] {
			removed[port] = true
		}
	}

	for port := range newSet {
		if !oldSet[port] {
			added[port] = true
		}
	}

	return removed, added
}

// validatePortAssignmentConflicts checks whether multiple virtual switches in the
// same port_map share the same port names. Returns a non-empty string describing
// conflicts if found.
func validatePortAssignmentConflicts(portMap map[string]switchEntry) string {
	portToVS := make(map[string][]string)

	for vsName, entry := range portMap {
		for _, p := range entry.Ports {
			if p.Name != "" {
				portToVS[p.Name] = append(portToVS[p.Name], vsName)
			}
		}
	}

	var conflicts []string
	for portName, vsList := range portToVS {
		if len(vsList) > 1 {
			conflicts = append(conflicts, fmt.Sprintf("port %q is assigned to multiple virtual switches: %s", portName, strings.Join(vsList, ", ")))
		}
	}

	if len(conflicts) > 0 {
		return strings.Join(conflicts, "\n")
	}
	return ""
}

// getClientVdom returns the default vdom from the client config.
func (r *resourceSystemVirtualSwitchPortAssignment) getClientVdom() string {
	if r.client != nil && r.client.Config.Auth.Vdom != "" {
		return r.client.Config.Auth.Vdom
	}
	return ""
}

// getFortiSDKClient extracts the FortiSDKClient from the provider data using reflection.
func getFortiSDKClient(providerData interface{}) (*forticlient.FortiSDKClient, error) {
	if providerData == nil {
		return nil, fmt.Errorf("provider data is nil")
	}

	val := reflect.ValueOf(providerData)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return nil, fmt.Errorf("expected pointer, got %T", providerData)
	}

	clientField := val.Elem().FieldByName("Client")
	if !clientField.IsValid() {
		return nil, fmt.Errorf("Client field not found")
	}

	client, ok := clientField.Interface().(*forticlient.FortiSDKClient)
	if !ok {
		return nil, fmt.Errorf("Client field is not *FortiSDKClient")
	}

	return client, nil
}
