---
subcategory: "FortiGate Certificate"
layout: "fortios"
page_title: "FortiOS: fortios_certificate_local"
description: |-
  Local keys and certificates.
---

# fortios_certificate_local
Local keys and certificates.

Due to the limitations of the current system, the feature is temporarily unavailable. Please use the following resource configuration as an alternative.

## Example
### Import Certificate:

**Step1: Prepare certificate**

The following key is a randomly generated example key for testing. In actual use, please replace it with your own key.

```
# cat xxx.key
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAy7Vj/9qM4LiWJjd+GMUijSrIMJLmoHDTLKPiT4CAQvgn+MDo
1J/ygkVJJSwifWn8ouGSycasjmgFJJIiu2stQzm4EQ3u85fvqazoo806pBOQm1gW
YEY5Ug3yMpddd7xUQO8s60VD/ZQ8pNzsKZlRWelGnNw9TzZXABi/FBxLFQ4TOnw5
zfibZ1dnvMV3mdg0MY2qoe96wpiSZKt4gkSmzPwRR7QnYzBtF8vVLlJ2I5fqJo4/
nt9j3Kd2X2wUaZ194aIrruCFU6CFxAb2ZTJLw4/97EXR8wqoDfhU55qk7pUR22Uc
uftobWYCMrO7e51WdBYEWIldAtcZkUsk7bUUzQIDAQABAoIBABlV8x0EOpdMfeg8
6KL+CcES/BkGfEaiIbGgpGoM6mbp5FbM72haiFfpdCJ6bcO5ZeGAOrh7zERd7Z3R
yx4SQ2vkBt+gIwMK95Tb24db5Bo6ELcxan8I3OI2t9PQ/aABvVziIm0UjVNBl5VN
oNW/qt2K5Oxne/yZHpL1gPZoWnJAuBNcDZDNI5qQfT1JTWhmbu9pkjiNg3h0l5O4
boETBdb3W2jlvCIegIQJ+xPkChS3I4cMoZ4LBRWMLpzK+Q57+zml75drsQ7PA0XH
lx2nzUFCByu27pM6kiajXleUSGVH2VCUSNysQAYBSa77g5t7O+m+o3iNUslUDor3
LY5eiKUCgYEA6dqJJxF28Wt6UbMgywQuv5koo8v8nyUhR4hZhvy5qge5v5Sh82UE
RyVfSvMDR9oWnXs8tBZaf1sgsUEaFl2I/5kmFWTe7aGj1eXprOxOuMNk51AN2w9T
jHWici/rj0JEjvAteDvLjY6CTAi8Zg9OxuJWNyV2gZ1LpZ2cIlULzLsCgYEA3wAH
JQ0jVeJ70v2I01m16d/klTNcqv9sTIPwowz0RFkOG8v482SSQ7Q43xkSYJvxKg12
BO6qq+RzYwDPgA7/7kLrq1Ye2VobhrsRey6dEWGdrgA/TfoCgSjK0LEBh4Gn5h7l
DycvfNRbum1D+uheyTPC94fJChwutihUsrXuEBcCgYAaoUczCrsXvNx+Bz75v20v
ZlqJZIZM/SZwBefkBk2CPkT5uwxCMkOtcmUKnOfHu98NaeY8v7roe9EaPkahO1+J
c8AxeX4lY13L0tWsWnCQe7e225foVTN3cEHibPCPLMWv3UvgQDbq1Mqjq+8AVEft
QAL/XqXDFs1xe6Q3CKZCVwKBgQCaCBbnTM/ffvUwo9dixVCWHwRw2m1j39Iad/g7
Z7NBkpHgOV/YHtu40D+IOnUrLgvClFG0znYtDTt2YxTwy2uUU70dN/tO/qKMyaIl
h+kOHHMhwSH45nvcYyTUSa9YvgIPPb/SW6q9eqFxgA+4u9DdAVfmSnBe/2B0ih8W
4ftyOQKBgF0puFMyA7ySE2tQ5quiPBO95Vls4SMl59ofhEgMghmUErEFGvTHPxff
2UF7AWahrhNq02cF8iTU/lS77D0W01TpEFl8xC5LyqewKzLatgFTFhFPPGt/wK0s
uIAliRuV1Iyv2vDYmYaugpeZakogVBpkVPqvEzfEPgn9VEZKLQ7y
-----END RSA PRIVATE KEY-----

#cat xxxx.crt
-----BEGIN CERTIFICATE-----
MIIDuTCCAqGgAwIBAgIJAKCr2aCM9Je5MA0GCSqGSIb3DQEBCwUAMHMxCzAJBgNV
BAYTAkdCMRMwEQYDVQQIDApTb21lLVN0YXRlMQ8wDQYDVQQHDAZMb25kb24xDDAK
BgNVBAoMA0ZETDESMBAGA1UEAwwJbG9jYWxob3N0MRwwGgYJKoZIhvcNAQkBFg1z
c0Bzc3Nzc3MuY29tMB4XDTE5MDUyOTE1MDIzMFoXDTIwMDUyODE1MDIzMFowczEL
MAkGA1UEBhMCR0IxEzARBgNVBAgMClNvbWUtU3RhdGUxDzANBgNVBAcMBkxvbmRv
bjEMMAoGA1UECgwDRkRMMRIwEAYDVQQDDAlsb2NhbGhvc3QxHDAaBgkqhkiG9w0B
CQEWDXNzQHNzc3Nzcy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQDLtWP/2ozguJYmN34YxSKNKsgwkuagcNMso+JPgIBC+Cf4wOjUn/KCRUklLCJ9
afyi4ZLJxqyOaAUkkiK7ay1DObgRDe7zl++prOijzTqkE5CbWBZgRjlSDfIyl113
vFRA7yzrRUP9lDyk3OwpmVFZ6Uac3D1PNlcAGL8UHEsVDhM6fDnN+JtnV2e8xXeZ
2DQxjaqh73rCmJJkq3iCRKbM/BFHtCdjMG0Xy9UuUnYjl+omjj+e32Pcp3ZfbBRp
nX3hoiuu4IVToIXEBvZlMkvDj/3sRdHzCqgN+FTnmqTulRHbZRy5+2htZgIys7t7
nVZ0FgRYiV0C1xmRSyTttRTNAgMBAAGjUDBOMB0GA1UdDgQWBBTinUGXSHLwDOVm
OMdVbk0NdJNcRzAfBgNVHSMEGDAWgBTinUGXSHLwDOVmOMdVbk0NdJNcRzAMBgNV
HRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBX7ZHWH3N8EN+XcmUak7RG9qzy
mnWPvyBiM8YI11rs87SkD+L8Q/ylxdmoI57cfPHpnqtkGAseRdN1EtzqILpyOo4Q
Q2aF3ZHKUOEPBbblWqv+AbyXPHhODrm+Nlyu42axcqfLwLGAIRhVkVerX24lV/u2
s3W/G5cse7RfNtxWPVUah7jAmsIv1Yc7Yi4TEIlQDImQI5TAoGTQOpPjYZXCtHXS
xUIGOKDTds9X2wWb3lM7ANecrINh6CNB/tg3GNdGV8SCJvJnYtEgfipjS7cQoc5C
pBmnz+IlqzrwBZBxmB+1xFrYATm/hZ3HMFrLKQVoTJgTP74/PIpCaO/mjis4
-----END CERTIFICATE-----

```


**Step2: Prepare TF file with fortios_json_generic_api resource**

```
#cat main.tf
provider "fortios" {
  hostname = "192.168.52.177"
  insecure = "true"
  token    = "GNH7r40H65GNb46kd4rG8rtrmn0fr1"
}

data "local_file" "key_file" {
    filename = "./test.key"
}

data "local_file" "crt_file" {
    filename = "./test.crt"
}

resource "fortios_json_generic_api" "genericapi1" {
  path   = "/api/v2/monitor/vpn-certificate/local/import"
  method = "POST"
  json   = <<EOF
{
    "type": "regular",
    "certname": "testcer",
    "password": "",
    "key_file_content": "${data.local_file.key_file.content_base64}",
    "file_content": "${data.local_file.crt_file.content_base64}"
}
EOF
}

```

**Step3: Apply**
```
# terraform apply
data.local_file.key_file: Refreshing state...
data.local_file.crt_file: Refreshing state...

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # fortios_json_generic_api.genericapi1 will be created
  + resource "fortios_json_generic_api" "genericapi1" {
      + id       = (known after apply)
      + json     = jsonencode(
            {
              + certname         = "testcer"
              + file_content     = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUR1VENDQXFHZ0F3SUJBZ0lKQUtDcjJhQ005SmU1TUEwR0NTcUdTSWIzRFFFQkN3VUFNSE14Q3pBSkJnTlYKQkFZVEFrZENNUk13RVFZRFZRUUlEQXBUYjIxbExWTjBZWFJsTVE4d0RRWURWUVFIREFaTWIyNWtiMjR4RERBSwpCZ05WQkFvTUEwWkVUREVTTUJBR0ExVUVBd3dKYkc5allXeG9iM04wTVJ3d0dnWUpLb1pJaHZjTkFRa0JGZzF6CmMwQnpjM056YzNNdVkyOXRNQjRYRFRFNU1EVXlPVEUxTURJek1Gb1hEVEl3TURVeU9ERTFNREl6TUZvd2N6RUwKTUFrR0ExVUVCaE1DUjBJeEV6QVJCZ05WQkFnTUNsTnZiV1V0VTNSaGRHVXhEekFOQmdOVkJBY01Ca3h2Ym1SdgpiakVNTUFvR0ExVUVDZ3dEUmtSTU1SSXdFQVlEVlFRRERBbHNiMk5oYkdodmMzUXhIREFhQmdrcWhraUc5dzBCCkNRRVdEWE56UUhOemMzTnpjeTVqYjIwd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUIKQVFETHRXUC8yb3pndUpZbU4zNFl4U0tOS3Nnd2t1YWdjTk1zbytKUGdJQkMrQ2Y0d09qVW4vS0NSVWtsTENKOQphZnlpNFpMSnhxeU9hQVVra2lLN2F5MURPYmdSRGU3emwrK3ByT2lqelRxa0U1Q2JXQlpnUmpsU0RmSXlsMTEzCnZGUkE3eXpyUlVQOWxEeWszT3dwbVZGWjZVYWMzRDFQTmxjQUdMOFVIRXNWRGhNNmZEbk4rSnRuVjJlOHhYZVoKMkRReGphcWg3M3JDbUpKa3EzaUNSS2JNL0JGSHRDZGpNRzBYeTlVdVVuWWpsK29tamorZTMyUGNwM1pmYkJScApuWDNob2l1dTRJVlRvSVhFQnZabE1rdkRqLzNzUmRIekNxZ04rRlRubXFUdWxSSGJaUnk1KzJodFpnSXlzN3Q3Cm5WWjBGZ1JZaVYwQzF4bVJTeVR0dFJUTkFnTUJBQUdqVURCT01CMEdBMVVkRGdRV0JCVGluVUdYU0hMd0RPVm0KT01kVmJrME5kSk5jUnpBZkJnTlZIU01FR0RBV2dCVGluVUdYU0hMd0RPVm1PTWRWYmswTmRKTmNSekFNQmdOVgpIUk1FQlRBREFRSC9NQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUJYN1pIV0gzTjhFTitYY21VYWs3Ukc5cXp5Cm1uV1B2eUJpTThZSTExcnM4N1NrRCtMOFEveWx4ZG1vSTU3Y2ZQSHBucXRrR0FzZVJkTjFFdHpxSUxweU9vNFEKUTJhRjNaSEtVT0VQQmJibFdxditBYnlYUEhoT0RybStObHl1NDJheGNxZkx3TEdBSVJoVmtWZXJYMjRsVi91MgpzM1cvRzVjc2U3UmZOdHhXUFZVYWg3akFtc0l2MVljN1lpNFRFSWxRREltUUk1VEFvR1RRT3BQallaWEN0SFhTCnhVSUdPS0RUZHM5WDJ3V2IzbE03QU5lY3JJTmg2Q05CL3RnM0dOZEdWOFNDSnZKbll0RWdmaXBqUzdjUW9jNUMKcEJtbnorSWxxenJ3QlpCeG1CKzF4RnJZQVRtL2haM0hNRnJMS1FWb1RKZ1RQNzQvUElwQ2FPL21qaXM0Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0="
              + key_file_content = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBeTdWai85cU00TGlXSmpkK0dNVWlqU3JJTUpMbW9IRFRMS1BpVDRDQVF2Z24rTURvCjFKL3lna1ZKSlN3aWZXbjhvdUdTeWNhc2ptZ0ZKSklpdTJzdFF6bTRFUTN1ODVmdnFhem9vODA2cEJPUW0xZ1cKWUVZNVVnM3lNcGRkZDd4VVFPOHM2MFZEL1pROHBOenNLWmxSV2VsR25OdzlUelpYQUJpL0ZCeExGUTRUT253NQp6ZmliWjFkbnZNVjNtZGcwTVkycW9lOTZ3cGlTWkt0NGdrU216UHdSUjdRbll6QnRGOHZWTGxKMkk1ZnFKbzQvCm50OWozS2QyWDJ3VWFaMTk0YUlycnVDRlU2Q0Z4QWIyWlRKTHc0Lzk3RVhSOHdxb0RmaFU1NXFrN3BVUjIyVWMKdWZ0b2JXWUNNck83ZTUxV2RCWUVXSWxkQXRjWmtVc2s3YlVVelFJREFRQUJBb0lCQUJsVjh4MEVPcGRNZmVnOAo2S0wrQ2NFUy9Ca0dmRWFpSWJHZ3BHb002bWJwNUZiTTcyaGFpRmZwZENKNmJjTzVaZUdBT3JoN3pFUmQ3WjNSCnl4NFNRMnZrQnQrZ0l3TUs5NVRiMjRkYjVCbzZFTGN4YW44STNPSTJ0OVBRL2FBQnZWemlJbTBValZOQmw1Vk4Kb05XL3F0Mks1T3huZS95WkhwTDFnUFpvV25KQXVCTmNEWkROSTVxUWZUMUpUV2htYnU5cGtqaU5nM2gwbDVPNApib0VUQmRiM1cyamx2Q0llZ0lRSit4UGtDaFMzSTRjTW9aNExCUldNTHB6SytRNTcrem1sNzVkcnNRN1BBMFhICmx4Mm56VUZDQnl1MjdwTTZraWFqWGxlVVNHVkgyVkNVU055c1FBWUJTYTc3ZzV0N08rbStvM2lOVXNsVURvcjMKTFk1ZWlLVUNnWUVBNmRxSkp4RjI4V3Q2VWJNZ3l3UXV2NWtvbzh2OG55VWhSNGhaaHZ5NXFnZTV2NVNoODJVRQpSeVZmU3ZNRFI5b1duWHM4dEJaYWYxc2dzVUVhRmwySS81a21GV1RlN2FHajFlWHByT3hPdU1OazUxQU4ydzlUCmpIV2ljaS9yajBKRWp2QXRlRHZMalk2Q1RBaThaZzlPeHVKV055VjJnWjFMcFoyY0lsVUx6THNDZ1lFQTN3QUgKSlEwalZlSjcwdjJJMDFtMTZkL2tsVE5jcXY5c1RJUHdvd3owUkZrT0c4djQ4MlNTUTdRNDN4a1NZSnZ4S2cxMgpCTzZxcStSell3RFBnQTcvN2tMcnExWWUyVm9iaHJzUmV5NmRFV0dkcmdBL1Rmb0NnU2pLMExFQmg0R241aDdsCkR5Y3ZmTlJidW0xRCt1aGV5VFBDOTRmSkNod3V0aWhVc3JYdUVCY0NnWUFhb1VjekNyc1h2TngrQno3NXYyMHYKWmxxSlpJWk0vU1p3QmVma0JrMkNQa1Q1dXd4Q01rT3RjbVVLbk9mSHU5OE5hZVk4djdyb2U5RWFQa2FoTzErSgpjOEF4ZVg0bFkxM0wwdFdzV25DUWU3ZTIyNWZvVlROM2NFSGliUENQTE1XdjNVdmdRRGJxMU1xanErOEFWRWZ0ClFBTC9YcVhERnMxeGU2UTNDS1pDVndLQmdRQ2FDQmJuVE0vZmZ2VXdvOWRpeFZDV0h3UncybTFqMzlJYWQvZzcKWjdOQmtwSGdPVi9ZSHR1NDBEK0lPblVyTGd2Q2xGRzB6bll0RFR0Mll4VHd5MnVVVTcwZE4vdE8vcUtNeWFJbApoK2tPSEhNaHdTSDQ1bnZjWXlUVVNhOVl2Z0lQUGIvU1c2cTllcUZ4Z0ErNHU5RGRBVmZtU25CZS8yQjBpaDhXCjRmdHlPUUtCZ0YwcHVGTXlBN3lTRTJ0UTVxdWlQQk85NVZsczRTTWw1OW9maEVnTWdobVVFckVGR3ZUSFB4ZmYKMlVGN0FXYWhyaE5xMDJjRjhpVFUvbFM3N0QwVzAxVHBFRmw4eEM1THlxZXdLekxhdGdGVEZoRlBQR3Qvd0swcwp1SUFsaVJ1VjFJeXYydkRZbVlhdWdwZVpha29nVkJwa1ZQcXZFemZFUGduOVZFWktMUTd5Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t"
              + password         = ""
              + type             = "regular"
            }
        )
      + method   = "POST"
      + path     = "/api/v2/monitor/vpn-certificate/local/import"
      + response = (known after apply)
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

fortios_json_generic_api.genericapi1: Creating...
fortios_json_generic_api.genericapi1: Creation complete after 0s [id=JsonGenericApie23d30b4-9845-43f5-ac28-f082fd5b95ef]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.

```
**Step4: Check the results**

![111](https://user-images.githubusercontent.com/49291382/99555680-86cc5400-29fb-11eb-8ae8-2c437f13595e.png)

### Delete Certificate:
```
resource "fortios_system_autoscript" "trname1" {
  interval    = 1
  name        = "delcerttest"
  output_size = 10
  repeat      = 1
  script      = <<EOF
config vpn certificate local
delete testcer
end
EOF
  start       = "auto"
}
```

