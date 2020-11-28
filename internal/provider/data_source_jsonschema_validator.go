package provider

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xeipuuv/gojsonschema"
)

func dataSourceJsonschemaValidator() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJsonschemaValidatorRead,

		Schema: map[string]*schema.Schema{
			"document": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "body of a yaml or json file",
			},

			"schema": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "json schema to validate content by",
			},

			"validated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJsonschemaValidatorRead(d *schema.ResourceData, m interface{}) error {
	var (
		err    error = nil
		result       = new(gojsonschema.Result)
	)

	document := d.Get("document").(string)

	schemaLoader := gojsonschema.NewStringLoader(
		d.Get("schema").(string))
	documentLoader := gojsonschema.NewStringLoader(document)

	result, err = gojsonschema.Validate(schemaLoader, documentLoader)
	if err == nil {
		if result.Valid() {
			err = d.Set("validated", document)
		} else {
			message := "The document is not valid. see errors :\n"
			for _, desc := range result.Errors() {
				message += fmt.Sprintf("[%s]\n", desc)
			}
			err = errors.New(message)
		}
	}

	if err != nil {
		return err
	}

	d.SetId(hash(document))
	return nil
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}
