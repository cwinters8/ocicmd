# ocicmd

A command line utility that provides additional functionality beyond that provided by the `oci` CLI from Oracle.

Currently the functionality is limited to generating compartment information. There is a lot of potential for features, like being able to list or manipulate resources across multiple compartments. Pull requests are very much welcome if you would like to contribute.

Being able to get a list of compartments that match certain filters can be useful in combination with the [official `oci`](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/cliconcepts.htm) CLI from Oracle. You can pass the compartment IDs to search for instances, databases, etc.

For example, you can list all compute instances in compartments whose names contain the word 'application':

```sh
ocicmd compartments list -r --filter 'application' | jq '.[].id' | tr -d '"' | while read line ; do oci compute instance list --compartment-id $line | jq '[.data[] | {id, "display-name", "availability-domain"}]' ; done
```

This process of passing the compartments to `oci` is an improvement over manually checking each compartment, but its still very clunky. This is a big area for potential improvements IMO, as we can just make calls to the SDK to get the instance information instead of having to do this nonsense with `jq`, `while read line; do`, etc.

## Prerequisites

You will need a local configuration file in your environment. Follow [Oracle's instructions](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/apisigningkey.htm#apisigningkey_topic_How_to_Generate_an_API_Signing_Key_Console) on generating an API signing key for your user and creating the OCI config file.

`go` must be installed and available in your `$PATH`.

## Build from source

Clone this repo then you can either install in your path or build and use from a location of your choosing.

### Install

To install in $PATH:

```sh
go install
```

### Build without installing

This can be useful if you need to test changes quickly and don't need to have the tool installed in your $PATH

The target directory (`bin`) can be changed depending on your preferences.

```sh
go build -o bin/ocicmd
```

Then you can run from the target directory:

```sh
./bin/ocicmd
```

## Usage

Commands below assume you have installed `ocicmd` in your $PATH.

View command usage

```sh
ocicmd -h

# specifically compartments usage
ocicmd compartments -h
```

Listing all compartments recursively

```sh
ocicmd compartments -r
```

Filtering for specific compartments

```sh
ocicmd compartments -r --filter 'application'
```
