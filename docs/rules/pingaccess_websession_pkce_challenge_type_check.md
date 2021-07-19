# pingaccess_websession_pkce_challenge_type_check

Checks websessions are configured to use secure cookie.

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_websession_pkce_challenge_type_check" {
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
  secure_cookie       = true
  pkce_challenge_type = "OFF"
}
```

```console
$ tflint
1 issue(s) found:

Warning: pkce_challenge_type is OFF (pingaccess_websession_pkce_challenge_type_check)

  on main.tf line 21:
  21:   pkce_challenge_type = "OFF"
```

## Why

Proof Key Code Exchange ensures an additional layer of protection during code exchange.

## How To Fix

If the PKCE is supported set to `SHA256`. 