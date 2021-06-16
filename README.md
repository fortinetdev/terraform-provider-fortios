# Terraform Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x +
- [Go](https://golang.org/doc/install) 1.13.x (to build the provider plugin)
- The provider can cover FortiOS 6.0, 6.2, 6.4, 7.0 versions, the configuration of all parameters should be based on the relevant FortiOS version manual. The provider can cover both FortiManager 6.0 and 6.2 versions. When using FortiManager, make sure the versions of FortiManager and the FortiGates controlled by it are the same.

## Building the Provider

1. Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-fortios`.

    ```sh
    $ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
    $ git clone git@github.com:fortinetdev/terraform-provider-fortios
    ```

2. Enter the provider directory and build the provider.

    ```sh
    $ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-fortios
    $ make build
    ```

## Using the Provider

If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory,  run `terraform init` to initialize it.

```sh
$ terraform init
```

## Developing the Provider

If you wish to work on the provider, you'll first need Go installed on your machine (version 1.13+ is required). You'll also need to correctly setup a GOPATH, as well as adding $GOPATH/bin to your $PATH.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the $GOPATH/bin directory.

```sh
$ make build
...
```
