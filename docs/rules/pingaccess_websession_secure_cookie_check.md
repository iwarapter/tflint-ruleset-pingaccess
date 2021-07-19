# pingaccess_websession_secure_cookie_check

Checks websessions are configured to use secure cookie.

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_websession_secure_cookie_check" {
  enabled = true
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_websession" "example" {
  name     = "example"
  audience = "aud"
  client_credentials {
    client_id = "websession"
    client_secret {
      value = "top_secret"
    }
  }
  scopes              = ["profile", "address", "email", "phone"]
  secure_cookie       = false
  pkce_challenge_type = "SHA256"
}
```

```console
$ tflint   
1 issue(s) found:

Warning: secure_cookie is false (pingaccess_websession_secure_cookie_check)

  on main.tf line 20:
  20:   secure_cookie       = false
```

## Why

Secure cookies are only sent over HTTPS preventing interception.

## How To Fix

If running on https set the `secure_cookie` attribute to `true`. 