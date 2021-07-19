# pingaccess_pingfederate_runtime_duplicate_check

Checks for duplicate resources which are singletons.

## Configuration

| Name | Default | Description |
|---|---|---|
| enabled | `true` | Enable the rule |

## Example

#### Rule Configuration

```hcl
rule "pingaccess_pingfederate_runtime_duplicate_check" {
  enabled = true
}
```

#### Sample terraform source file
```hcl
resource "pingaccess_pingfederate_runtime" "foo" {
  issuer = "https://foo"
}

resource "pingaccess_pingfederate_runtime" "bar" {
  issuer = "https://bar"
}
```

```console
$ tflint
1 issue(s) found:

Error: duplicate instance of pingaccess_pingfederate_runtime (pingaccess_pingfederate_runtime_duplicate_check)

  on main.tf line 14:
  14: resource "pingaccess_pingfederate_runtime" "bar" {
```

## Why

Certain resources in the PingFederate provider are singletons and should only be defined once.

## How To Fix

Ensure only one correctly configured resource.