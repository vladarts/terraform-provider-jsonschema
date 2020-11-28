# jsonschema_validator

The `jsonschema_validator` data source validates a json document using [json-schema](https://json-schema.org/).

## Example Usage

```hcl-terraform
data "jsonschema_validator" "values" {
  document = file("/path/to/document.json")
  schema = file("/path/to/schema.json")
}
```

## Argument Reference

The following arguments are supported:

* `document` &mdash; (Required) Content of a json document.
* `schema` &mdash; (Required) Content of a [json-schema](https://json-schema.org/) file.

## Attributes Reference

The following attribute are exported:

* `validated` &mdash; equivalent to `document` argument.
