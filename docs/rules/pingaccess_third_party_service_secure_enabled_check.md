# pingaccess_third_party_service_secure_enabled_check

Checks third party services are using secure (HTTPS) connections.

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_third_party_service_secure_enabled_check" {
  enabled = true
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_third_party_service" "example" {
  name    = "example"
  targets = ["localhost:1234"]
  secure  = false
}
```

```console
$ tflint 
1 issue(s) found:

Warning: secure is false (pingaccess_third_party_service_secure_enabled_check)

  on main.tf line 13:
  13:   secure  = false
```

## Why

Use of unencrypted backend sites is in insecure.

## How To Fix

If the backend site supports secure communication then set the `secure` attribute to `true`. 