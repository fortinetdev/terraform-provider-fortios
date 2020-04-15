# Terraform Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)
- FortiOS 6.0 or 6.2. This provider uses the FortiOS API. All the resources are validated with FortiOS 6.0 and 6.2.

## Building the Provider

1. Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-fortios`.

    ```sh
    $ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
    $ git clone git@github.com:terraform-providers/terraform-provider-fortios
    ```

2. Enter the provider directory and build the provider.

    ```sh
    $ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-fortios
    $ make build
    ```

### Modifying vendor source code

Since this repository is using a vendor folder, in order to build changes in that folder we'll need add a flag upon building.

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-fortios
$ go build -mod=vendor
```

## Using the Provider

If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory,  run `terraform init` to initialize it.

```sh
$ terraform init
```

## Developing the Provider

If you wish to work on the provider, you'll first need Go installed on your machine (version 1.11+ is required). You'll also need to correctly setup a GOPATH, as well as adding $GOPATH/bin to your $PATH.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the $GOPATH/bin directory.

```sh
$ make build
...
```

## Testing

Running the acceptance test suite requires a FortiGate VM/device to test against.

To run the test:
1. Set `FORTIOS_ACCESS_HOSTNAME` to point to a FortiGate VM/device.
2. Set `FORTIOS_ACCESS_TOKEN` to the access token of a Rest API user on that device.
3. Run `make testacc`.

    ```sh
    make testacc
    ```

For more information about acceptance testing in Terraform, see [Running an Acceptance Test](https://github.com/hashicorp/terraform/blob/master/.github/CONTRIBUTING.md#running-an-acceptance-test).
