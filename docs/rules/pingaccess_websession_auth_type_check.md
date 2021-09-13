# pingaccess_websession_auth_type_check

Checks websessions are configured to use specified `credential_types`.

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `false` | Enable the rule |
| allowed_credential_types | [] | Specify the allowed authentication types |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_websession_auth_type_check" {
  enabled = true
  allowed_credential_types = ["PRIVATE_KEY_JWT"]
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_websession" "example" {
  name     = "example"
  audience = "aud"
  client_credentials {
    client_id        = "websession"
    credentials_type = "SECRET"
    client_secret {
      value = "top_secret"
    }
  }
  scopes              = ["profile", "address", "email", "phone"]
  secure_cookie       = true
  pkce_challenge_type = "SHA256"
}
```

```console
$ tflint
1 issue(s) found:

Warning: credentials_type is SECRET, allowed types are PRIVATE_KEY_JWT (pingaccess_websession_auth_type_check)

  on main.tf line 15:
  15:     credentials_type = "SECRET"
```

## Why

Allows users to specify the types of authentication mechanisms they wish to enforce for websessions.

## How To Fix

Use the specified authentication type. 