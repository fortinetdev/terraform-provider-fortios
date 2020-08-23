---
layout: "fortios"
page_title: "FortiOS: fortios_system_license_vm"
sidebar_current: "docs-fortios-resource-system-license-vm"
subcategory: "FortiGate System"
description: |-
  Provides a resource to update VM license using uploaded file for FortiOS. Reboots immediately if successful.
---

# fortios_system_license_vm
Provides a resource to update VM license using uploaded file for FortiOS. Reboots immediately if successful.

## Example Usage
```hcl
resource "fortios_system_license_vm" "test2" {
  file_content = "LS0tLS1CRUdJTiBGR1QgVk0gTElDRU5TRS0tLS0tDQpRQUFBQUxXaTdCVnVkV2x3QXJZcC92S2J2Yk5zME5YNWluUW9sVldmcFoxWldJQi9pL2g4c01oR0psWWc5Vkl1DQorSlBJRis1aFphMWwyNm9yNHdiEQE3RnJDeVZnQUFBQWhxWjliWHFLK1hGN2o3dnB3WTB6QXRTaTdOMVM1ZWNxDQpWYmRRREZyYklUdnRvUWNyRU1jV0ltQzFqWWs5dmVoeGlYTG1OV0MwN25BSitYTTJFNmh2b29DMjE1YUwxK2wrDQovUHl5M0VLVnNTNjJDT2hMZHc3UndXajB3V3RqMmZiWg0KLS0tLS1FTkQgRkdUIFZNIExJQ0VOU0UtLS0tLQ0K"
}
```

## Argument Reference
The following arguments are supported:

* `file_content` - (Required) The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.

