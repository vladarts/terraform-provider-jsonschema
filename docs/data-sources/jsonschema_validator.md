# jsonschema_validator Data Source

The `jsonschema_validator` data source validates a json document using [json-schema](https://json-schema.org/).

## Example Usage

```hcl-terraform
data "jsonschema_validator" "values" {
  document = file("/path/to/document.json")
  schema = file("/path/to/schema.json")
}
```

## Argument Reference

List arguments this data source takes:

* `document` &mdash; (Required) Content of a json document.
* `schema` &mdash; (Required) Content of a [json-schema](https://json-schema.org/) file.

## Attributes Reference

List attributes that this data source exports:

* `validated` &mdash; equivalent to `document` argument.
