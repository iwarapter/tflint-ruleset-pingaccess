# pingaccess_site_secure_check

Checks sites are configured to use secure protocol (https).

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_site_secure_check" {
  enabled = true
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_site" "example1" {
	name                         = "example1"
	targets                      = ["www.google.com"]
	max_connections              = -1
	max_web_socket_connections   = -1
	availability_profile_id      = 1
	use_target_host_header       = false
	secure                       = false
}
```

```console
$ tflint
1 issue(s) found:

Warning: secure is false (pingaccess_site_secure_check)

  on main.tf line 26:
  26:   secure                     = false
```

## Why

Use of unencrypted backend sites is in insecure.

## How To Fix

If the backend site supports secure communication then set the `secure` attribute to `true`. 