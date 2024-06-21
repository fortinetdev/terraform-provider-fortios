---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sshconfig"
description: |-
  Configure SSH config.
---

# fortios_system_sshconfig
Configure SSH config. Applies to FortiOS Version `>= 7.4.4`.

## Argument Reference

The following arguments are supported:

* `ssh_kex_algo` - Select one or more SSH kex algorithms. Valid values: `diffie-hellman-group1-sha1`, `diffie-hellman-group14-sha1`, `diffie-hellman-group14-sha256`, `diffie-hellman-group16-sha512`, `diffie-hellman-group18-sha512`, `diffie-hellman-group-exchange-sha1`, `diffie-hellman-group-exchange-sha256`, `curve25519-sha256@libssh.org`, `ecdh-sha2-nistp256`, `ecdh-sha2-nistp384`, `ecdh-sha2-nistp521`.
* `ssh_enc_algo` - Select one or more SSH ciphers. Valid values: `chacha20-poly1305@openssh.com`, `aes128-ctr`, `aes192-ctr`, `aes256-ctr`, `arcfour256`, `arcfour128`, `aes128-cbc`, `3des-cbc`, `blowfish-cbc`, `cast128-cbc`, `aes192-cbc`, `aes256-cbc`, `arcfour`, `rijndael-cbc@lysator.liu.se`, `aes128-gcm@openssh.com`, `aes256-gcm@openssh.com`.
* `ssh_mac_algo` - Select one or more SSH MAC algorithms. Valid values: `hmac-md5`, `hmac-md5-etm@openssh.com`, `hmac-md5-96`, `hmac-md5-96-etm@openssh.com`, `hmac-sha1`, `hmac-sha1-etm@openssh.com`, `hmac-sha2-256`, `hmac-sha2-256-etm@openssh.com`, `hmac-sha2-512`, `hmac-sha2-512-etm@openssh.com`, `hmac-ripemd160`, `hmac-ripemd160@openssh.com`, `hmac-ripemd160-etm@openssh.com`, `umac-64@openssh.com`, `umac-128@openssh.com`, `umac-64-etm@openssh.com`, `umac-128-etm@openssh.com`.
* `ssh_hsk_algo` - Select one or more SSH hostkey algorithms. Valid values: `ssh-rsa`, `ecdsa-sha2-nistp521`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp256`, `rsa-sha2-256`, `rsa-sha2-512`, `ssh-ed25519`.
* `ssh_hsk_override` - Enable/disable SSH host key override in SSH daemon. Valid values: `disable`, `enable`.
* `ssh_hsk_password` - Password for ssh-hostkey.
* `ssh_hsk` - Config SSH host key.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SshConfig can be imported using any of these accepted formats:
```
$ terraform import fortios_system_sshconfig.labelname SystemSshConfig

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_sshconfig.labelname SystemSshConfig
$ unset "FORTIOS_IMPORT_TABLE"
```
