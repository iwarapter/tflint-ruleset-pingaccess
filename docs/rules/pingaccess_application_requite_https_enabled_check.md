# pingaccess_application_requite_https_enabled_check

Checks applications require a HTTPS connection.

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_application_requite_https_enabled_check" {
  enabled = true
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_application" "demo" {
  application_type  = "API"
  name              = "api-demo"
  context_root      = "/"
  default_auth_type = "API"
  destination       = "Site"
  site_id           = pingaccess_site.example1.id
  virtual_host_ids  = [pingaccess_virtualhost.demo.id]
  require_https     = false
}
```

```console
$ tflint
1 issue(s) found:

Warning: require_https is false (pingaccess_application_requite_https_enabled_check)

  on main.tf line 18:
  18:   require_https     = false
```

## Why

Sets the application to use a secure HTTPS connection.

## How To Fix

Set the attribute `require_https` to `true`. 