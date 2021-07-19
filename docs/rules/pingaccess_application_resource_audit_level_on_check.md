# pingaccess_application_resource_audit_level_on_check

Checks application resource have audit enabled.

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_application_resource_audit_level_on_check" {
  enabled = true
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_application_resource" "demo" {
  name    = "example"
  methods = ["*"]
  path_patterns {
    pattern = "/as/token.oauth2"
    type    = "WILDCARD"
  }
  path_patterns {
    pattern = "%s"
    type    = "WILDCARD"
  }
  path_prefixes = [
    "/as/token.oauth2",
    "%s"
  ]
  audit_level    = "OFF"
  application_id = pingaccess_application.demo.id
}
```

```console
$ tflint
1 issue(s) found:

Warning: audit_level is OFF (pingaccess_application_resource_audit_level_on_check)

  on main.tf line 25:
  25:   audit_level    = "OFF"
```

## Why

Ensures all application resources are logging.

## How To Fix

Set the attribute `audit_level` to `ON` or remove (defaults to `ON`). 