---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_certificate_management_local"
description: |-
  Manage a fortios system local certificate.
---

# fortios_system_certificate_management_local
Use this resource to manage a fortios system local certificate

## Example Usage

```hcl
locals {
  cert = <<EOF
-----BEGIN CERTIFICATE-----
MIIDKjCCAhKgAwIBAgIUU8dDvvKDVqXom29wrEPzIvMjLoIwDQYJKoZIhvcNAQEL
BQAwHDEaMBgGA1UEAwwRdGVzdDIuZXhhbXBsZS5jb20wHhcNMjEwNjE1MjExNzQ4
WhcNMjEwNjE2MjExNzQ4WjAcMRowGAYDVQQDDBF0ZXN0Mi5leGFtcGxlLmNvbTCC
ASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALk7oiDavyvqtQFarTPNm08f
Z3viEU7w8+7Kdlg/iPkxZ6HXQZpfR2KeWCjpq+CqWQx1/M+piJVkoyfW3zMSYzw+
u/+5PAkuc9YYhSXPqBJhw4VI59mQhchSckxq/7UmebMMy8tPGEJbQLF76lfga9mx
xMml1e8AQpU3VRlVzBKWwYE+wY5p+vIIyp8VBmN02ijbze0PAPEvQ62vi7qCHszS
ZkrxEPer5Oh6FSH409LoQQPhuBBbNau8TcyKv4vJO0HshkzM7xm69iDZVW3Pk5Se
1CsxOBn5W4SM+Bhb/w7oHlL4W09520cGtOn9PM+bokgsdsxRk2Es114AXw7MchEC
AwEAAaNkMGIwHQYDVR0OBBYEFIhW1JaPO2KKyngS0KxNcHJtdbYZMB8GA1UdIwQY
MBaAFIhW1JaPO2KKyngS0KxNcHJtdbYZMA8GA1UdEwEB/wQFMAMBAf8wDwYDVR0T
AQH/BAUwAwIBATANBgkqhkiG9w0BAQsFAAOCAQEAROoSayvNuTjRw00syveHvudl
2wKorrTTSvkoYBCwfEuV4sCDmQZqPk4RNi9XvDMJ73p5/g1I51yF+8MbV68iOFnl
hLzT5Thor3g1GZGihDBJiipTPsvMCx4FwdDiFqoVWHSYRrkzNsAeFJ6Z+ctnLzeL
VIkEmji4/kd0k5gEpy9kiS+cfi+8Rk4Hd1+jVp27YOty4v2UrwgWnDO0fIQtxI5n
/hSi8c7U0rKj0AKGrr79IFMcFgkSWN7zjfer5uSylz5ctply++tMwqUPE2wh5rTk
aN+30Pte6pA85vzPVE/oRVEENSixkxsM2b/GoCIix+4Y/KLqH+R6pvn1yB3WNg==
-----END CERTIFICATE-----
EOF
  key  = <<EOF
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC5O6Ig2r8r6rUB
Wq0zzZtPH2d74hFO8PPuynZYP4j5MWeh10GaX0dinlgo6avgqlkMdfzPqYiVZKMn
1t8zEmM8Prv/uTwJLnPWGIUlz6gSYcOFSOfZkIXIUnJMav+1JnmzDMvLTxhCW0Cx
e+pX4GvZscTJpdXvAEKVN1UZVcwSlsGBPsGOafryCMqfFQZjdNoo283tDwDxL0Ot
r4u6gh7M0mZK8RD3q+ToehUh+NPS6EED4bgQWzWrvE3Mir+LyTtB7IZMzO8ZuvYg
2VVtz5OUntQrMTgZ+VuEjPgYW/8O6B5S+FtPedtHBrTp/TzPm6JILHbMUZNhLNde
AF8OzHIRAgMBAAECggEALVLuJOPhizlu+NnbL6XLrtycUa/LVGmZBoD73DQPrAnu
tacaIk/WA8eDAt/Kcrq791SXe6icBxIM6h6llrWVGpSvI6+LhSOcrHJrggkBsx3A
3cgtEwtN8OpblV9JGmZDuRAUfbbo1LPHbKZJfR8oxKe+4yh05HpH0IMti3l26cUX
BndSQHI13t7lsiOGvupITTLWeG8+16q6ZiHf1sPLT3ZWvNeAUW1O04aqxIRRc4Ja
CBrGI/ECTif7TJN6eOTlRar3lqTeTtcHZkCH+cfHYmlIhPADXCxaHJJGU2itf1UZ
0AEMYfHBdd/VJ2JR6QTQmsvLPdPZYx7QQpc3D2ZdUQKBgQDiWMDq76DHS0Hbghsb
Imd1lkHvtnONvEPtbGeyVJ5Vhr08pbRbAgjQlGoOPnADZ6uZSCmnc8DCUh34AAkM
q0bH4rf7/xqxLtXIx0lVVjpEygiax/11QlxtO3HfwBfABN9+VG4uSldgjaCP/eJ8
04q7u/I32+aaoGbsCsKiusigMwKBgQDRgABJDfPSICfb4Dq4EPO8MRsurxOjoI5U
QwAbak0+ZeuvT1CNNFFnM1gl30VJXZMn6uVd9lXE6G23niqdKO7Hv8COnrj+1Qj2
/NgfWEYLIe59nU2N5wHJja5t4qi925GUGXECorKtLTa26Qw74F9q/fMDdofY57jn
Bj0idODQqwKBgQDIOQSmjkTmJp9iQswhi6SDcuBu4TGEvnZELvHn4UySkXcSj+5j
1v/fKnpKkVba3DkChcA7HXz1KFjUSYu3xkb9iIOCCd8dvzVjv04SjA1NTn6gFKsT
sBk8kyofaLhZprXg9WTl0+NJSN8woMBZ9XMysIYKfZ6XR67jvH0CIEJa2wKBgDDd
QDPJ81LReTqJxGhmW9NLSOHMqDIEIu54ai/6zWV+dEiBoXIt/8aobSj4OLrx/n7T
BQizijHRXLX1SE872uXwTcN2NgQKQHsLYoV9G2lBUtUtuYcdmIcgaszqjx0pd30p
qfhJZo+J+jcTiGz22oqkDOiD3w+yjKgmBmBhDobdAoGAIeHx1HgzmpRdOTIsOikv
hWZKZ+XRULaVDWfCne0kYtWVj1w6bk2dB2w1ityAmolkSMC9asdmrzzHApTmyewP
lfe4NTzxUOcu8lqMASsaWBb70PuAAeVkQ7VVMHUq2HzjN5UlJtnLfExnLdeWWAlK
Bjc76rLN8hpVicES4SOChok=
-----END PRIVATE KEY-----
EOF
}

resource "fortios_certificate_management_local" "example" {
  name        = "Global_Example"
  comments    = "Managed by Terraform"
  certificate = local.cert
  private_key = local.key
  scope       = "global"
  type        = "regular"

  lifecycle {
    create_before_destroy = true
  }
}

output "debug" {
  value = fortios_certificate_management_local.example.certificate_details
}
```

## Argument Reference

* `name` - (Required) Name of certificate. Changing the name forces a new resource to be created.
* `type` - (Required) Type of certificate (local|pkcs12|regular). Changing the type forces a new resource to be created.
* `scope` - (Required) Scope of local certificate (vdom|global). Global scope is only accessible for global administrators. Changing the scope forces a new resource to be created.
* `password` - Optional password for pkcs12 and regular certificate types. Changing the password forces a new resource to be created.
* `comments` - Comment.
* `private_key` - PEM format key. Changing the private_key forces a new resource to be created.
* `certificate` - PEM format certificate. Changing the certificate forces a new resource to be created.
* `ike_localid` - Local ID the FortiGate uses for authentication as a VPN client.
* `ike_localid_type` - IKE local ID type. (asn1dn:ASN.1 distinguished name. fqdn:Fully qualified domain name).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:

* `certificate_details` - Certificate details. The structure of `certificate_details` block is documented below.

The `certificate_details` block contains:

* `signature_algorithm` - The algorithm used to sign the certificate.
* `public_key_algorithm` - Tag category.
* `serial_number` - Number that uniquely identifies the certificate with the CA's system. The `format`
    function can be used to convert this base 10 number into other bases, such as hex.
* `is_ca` - `true` if this certificate is a ca certificate.
* `is_valid` - `true` if the certificate is within valid time period.
* `version` - The version the certificate is in.
* `issuer` - Who verified and signed the certificate, roughly following RFC2253.
* `subject` - The entity the certificate belongs to, roughly following RFC2253.
* `not_before` - The time after which the certificate is valid, as an RFC3339 timestamp.
* `not_after` - The time until which the certificate is invalid, as an RFC3339 timestamp.
* `sha1_fingerprint` - The SHA1 fingerprint of the public key of the certificate.