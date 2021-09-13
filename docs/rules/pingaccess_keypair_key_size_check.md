# pingaccess_keypair_key_size_check

Checks keypairs have an appropriate `key_size`.

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |
| rsa_key_sizes | [2048, 4096] | Accepted key sizes for RSA keys |
| ec_key_sizes | [256, 384, 521] | Accepted key sizes for EC keys |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_keypair_key_size_check" {
  enabled = true
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_keypair" "test_generate" {
  alias             = "acctest_test2"
  city              = "Test"
  common_name       = "Test"
  country           = "GB"
  key_algorithm     = "RSA"
  key_size          = 1024
  organization      = "Test"
  organization_unit = "Test"
  state             = "Test"
  valid_days        = 365
}
```

```console
$ tflint
1 issue(s) found:

Warning: key_size is 1024 (pingaccess_keypair_key_size_check)

  on main.tf line 16:
  16:   key_size          = 1024
```

## Why

Avoid creating keys that are considered weak due to its small key size.

## How To Fix

Set `key_size` to an accepted value.