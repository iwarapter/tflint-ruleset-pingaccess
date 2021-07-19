# Rules

This documentation describes a list of rules available by enabling this ruleset.

### Possible Errors

| Name                                                                                                    | Description                                               | Severity | Enabled |
| ------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- | -------- | ------- |
| [pingaccess_pingfederate_runtime_duplicate_check](./pingaccess_pingfederate_runtime_duplicate_check.md) | Rule for checking for duplicate `singleton` resource type | ERROR    | ✔       |

### Best Practices

| Name                                                                                                                                      | Description                                                                                           | Severity | Enabled |
| ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | -------- | ------- |
| [pingaccess_site_skip_hostname_verification_check](./pingaccess_site_skip_hostname_verification_check.md)                                 | Rule for checking `pingaccess_site` `skip_hostname_verification` is not set to `true`                 | WARNING  | ✔       |
| [pingaccess_site_secure_check](./pingaccess_site_secure_check.md)                                                                         | Rule for checking `pingaccess_site` `secure` is not set to `false`                                    | WARNING  | ✔       |
| [pingaccess_application_requite_https_enabled_check](./pingaccess_application_requite_https_enabled_check.md)                             | Rule for checking `pingaccess_application` `require_https` is not set to `false`                      | WARNING  | ✔       |
| [pingaccess_application_resource_audit_level_on_check](./pingaccess_application_resource_audit_level_on_check.md)                         | Rule for checking `pingaccess_application_resource` `audit_level` is not set to `OFF`                 | WARNING  | ✔       |
| [pingaccess_third_party_service_secure_enabled_check](./pingaccess_third_party_service_secure_enabled_check.md)                           | Rule for checking `pingaccess_third_party_service` `secure` is not set to `false`                     | WARNING  | ✔       |
| [pingaccess_third_party_service_skip_hostname_verification_check](./pingaccess_third_party_service_skip_hostname_verification_check.md)   | Rule for checking `pingaccess_third_party_service` `skip_hostname_verification` is not set to `true`  | WARNING  | ✔       |
| [pingaccess_trusted_certificate_group_skip_certificate_date_check](./pingaccess_trusted_certificate_group_skip_certificate_date_check.md) | Rule for checking `pingaccess_trusted_certificate` `skip_certificate_date_check` is not set to `true` | WARNING  | ✔       |
| [pingaccess_engine_listener_secure_check](./pingaccess_engine_listener_secure_check.md)                                                   | Rule for checking `pingaccess_engine_listener` `secure` is not set to `false`                         | WARNING  | ✔       |
| [pingaccess_websession_secure_cookie_check](./pingaccess_websession_secure_cookie_check.md)                                               | Rule for checking `pingaccess_websession` `secure_cookie` is not set to `false`                       | WARNING  | ✔       |
| [pingaccess_websession_pkce_challenge_type_check](./pingaccess_websession_pkce_challenge_type_check.md)                                   | Rule for checking `pingaccess_websession` `pkce_challenge_type` is not set to `OFF`                   | WARNING  | ✔       |
