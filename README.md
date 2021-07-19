# PingAccess Ruleset

[![Build Status](https://github.com/iwarapter/tflint-ruleset-pingaccess/workflows/build/badge.svg?branch=master)](https://github.com/iwarapter/tflint-ruleset-pingaccess/actions)

This ruleset provides various rules and best practices around the PingAccess terraform provider.

## Requirements

- TFLint v0.30+

## Installation

The plugin can be configured with the following config block to download and install the required version.
```hcl
plugin "pingaccess" {
    enabled = true
    version = "0.0.1"
    source  = "github.com/iwarapter/tflint-ruleset-pingaccess"

    signing_key = <<-KEY
    -----BEGIN PGP PUBLIC KEY BLOCK-----

    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
    xcB8MONKVAjD
    =MCW6
    -----END PGP PUBLIC KEY BLOCK-----
    KEY
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|pingaccess_pingfederate_runtime_duplicate_check|Rule for checking for duplicate `singleton` resource type|ERROR|✔||
|pingaccess_site_skip_hostname_verification_check|Rule for checking `pingaccess_site` `skip_hostname_verification` is not set to `true`|WARNING|✔||
|pingaccess_site_secure_check|Rule for checking `pingaccess_site` `secure` is not set to `false`|WARNING|✔||
|pingaccess_application_requite_https_enabled_check|Rule for checking `pingaccess_application` `require_https` is not set to `false`|WARNING|✔||
|pingaccess_application_resource_audit_level_on_check|Rule for checking `pingaccess_application_resource` `audit_level` is not set to `OFF`|WARNING|✔||
|pingaccess_third_party_service_secure_enabled_check|Rule for checking `pingaccess_third_party_service` `secure` is not set to `false`|WARNING|✔||
|pingaccess_third_party_service_skip_hostname_verification_check|Rule for checking `pingaccess_third_party_service` `skip_hostname_verification` is not set to `true`|WARNING|✔||
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
