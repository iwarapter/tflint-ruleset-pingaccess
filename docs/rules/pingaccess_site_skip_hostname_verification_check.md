# pingaccess_site_skip_hostname_verification_check

Checks sites don't skip hostname verification.

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_site_skip_hostname_verification_check" {
  enabled = true
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_site" "example" {
  name                       = "example"
  targets                    = ["www.google.com"]
  max_connections            = -1
  max_web_socket_connections = -1
  availability_profile_id    = 1
  skip_hostname_verification = true
  use_target_host_header     = false
  secure                     = true
}
```

```console
$ tflint 
1 issue(s) found:

Warning: skip_hostname_verification is true (pingaccess_site_skip_hostname_verification_check)

  on main.tf line 16:
  16:   skip_hostname_verification = true
```

## Why

Unverified hostname certificate checks allow for potential man in the middle attacks. 

## How To Fix

Change the `skip_hostname_verification` attribute to `false` or remove as the default is `false`. 