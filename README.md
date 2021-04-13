# PingAccess Ruleset
[![Build Status](https://github.com/iwarapter/tflint-ruleset-pingaccess/workflows/build/badge.svg?branch=master)](https://github.com/iwarapter/tflint-ruleset-pingaccess/actions)

This ruleset provides various rules and best practices around the PingAccess terraform provider.

## Requirements

- TFLint v0.24+
- Go v1.16

## Installation

Download the plugin and place it in `~/.tflint.d/plugins/tflint-ruleset-pingaccess` (or `./.tflint.d/plugins/tflint-ruleset-pingaccess`). When using the plugin, configure as follows in `.tflint.hcl`:

```hcl
plugin "pingaccess" {
    enabled = true
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|pingaccess_pingfederate_runtime_duplicate|Rule for checking for duplicate `singleton` resource type|ERROR|✔||
|pingaccess_site_skip_hostname_verification_check|Rule for checking `pingaccess_site` `skip_hostname_verification` is not enabled|WARNING|✔||

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```
