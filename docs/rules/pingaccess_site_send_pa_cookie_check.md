# pingaccess_site_send_pa_cookie_check

Checks sites do not have `send_pa_cookie` enabled

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_site_send_pa_cookie_check" {
  enabled = true
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_site" "example1" {
  name                       = "example1"
  targets                    = ["www.google.com"]
  max_connections            = -1
  max_web_socket_connections = -1
  availability_profile_id    = 1
  use_target_host_header     = false
  send_pa_cookie             = true
}
```

```console
$ tflint
1 issue(s) found:

Warning: send_pa_cookie is true (pingaccess_site_send_pa_cookie_check)

  on main.tf line 17:
  17:   send_pa_cookie             = true
```

## Why

Ensures PingAccess Token or OAuth Access Token stop at PingAccess and are not passed down. 

## How To Fix

Set `send_pa_cookie` to false. Ensure the downstream application will handle this change.