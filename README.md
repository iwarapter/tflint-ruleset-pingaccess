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
|pingaccess_site_skip_hostname_verification_check|Rule for checking `pingaccess_site` `skip_hostname_verification` is not set to `true`|WARNING|✔||
|pingaccess_site_secure_check|Rule for checking `pingaccess_site` `secure` is not set to `false`|WARNING|✔||
|pingaccess_application_requite_https_enabled_check|Rule for checking `pingaccess_application` `require_https` is not set to `false`|WARNING|✔||
|pingaccess_application_resource_audit_level_on_check|Rule for checking `pingaccess_application_resource` `audit_level` is not set to `OFF`|WARNING|✔||
|pingaccess_third_party_service_secure_enabled_check|Rule for checking `pingaccess_third_party_service` `secure` is not set to `false`|WARNING|✔||
|pingaccess_third_party_service_skip_hostname_verification_check|Rule for checking `pingaccess_third_party_service` `skip_hostname_verification` is not set to `true`|WARNING|✔||
|pingaccess_trusted_certificate_group_ignore_all_certificate_errors_check|Rule for checking `pingaccess_trusted_certificate` `ignore_all_certificate_errors` is not set to `true`|WARNING|✔||
|pingaccess_trusted_certificate_group_skip_certificate_date_check|Rule for checking `pingaccess_trusted_certificate` `skip_certificate_date_check` is not set to `true`|WARNING|✔||
|pingaccess_engine_listener_secure_check|Rule for checking `pingaccess_engine_listener` `secure` is not set to `false`|WARNING|✔||
|pingaccess_websession_secure_cookie_check|Rule for checking `pingaccess_websession` `secure_cookie` is not set to `false`|WARNING|✔||
|pingaccess_websession_pkce_challenge_type_check|Rule for checking `pingaccess_websession` `pkce_challenge_type` is not set to `OFF`|WARNING|✔||

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```
