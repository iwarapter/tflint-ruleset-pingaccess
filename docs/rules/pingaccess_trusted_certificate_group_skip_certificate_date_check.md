# pingaccess_trusted_certificate_group_skip_certificate_date_check

Checks trusted certificate groups are configured to check dates (expired/valid from etc).

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_trusted_certificate_group_skip_certificate_date_check" {
  enabled = true
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_trusted_certificate_group" "example" {
  name = "example"
  use_java_trust_store = true
  skip_certificate_date_check = true
}
```

```console
$  tflint       
1 issue(s) found:

Warning: skip_certificate_date_check is true (pingaccess_trusted_certificate_group_skip_certificate_date_check)
  
  on main.tf line 13:
  13:   skip_certificate_date_check = true
```

## Why

Trusting expired certificates is a security issue. 

## How To Fix

If the running on https set the `secure_cookie` attribute to `true`. 