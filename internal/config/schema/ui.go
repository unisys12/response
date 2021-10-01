package schema

import (
	"os"
	"path"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type uiDirDef struct {
}

func (d *uiDirDef) Name() string {
	return "ui_dir"
}

func (d *uiDirDef) Spec() hcldec.Spec {
	return &hcldec.BlockSpec{
		Nested: &hcldec.ObjectSpec{
			"console_dist": &hcldec.ValidateSpec{
				Wrapped: &hcldec.AttrSpec{
					Name: "console_dist",
					Type: cty.String,
				},
				Func: func(value cty.Value) hcl.Diagnostics {
					diag := hcl.Diagnostics{}

					if value.IsNull() {
						return hcl.Diagnostics{}
					}

					_, err := os.Stat(path.Join(value.AsString(), "index.html"))
					if os.IsNotExist(err) {
						return diag.Append(&hcl.Diagnostic{
							Summary: "The UI Directory provided was not found or does not contain an index.html file, which is required to serve the Response UI.",
						})
					}

					if os.IsPermission(err) {
						return diag.Append(&hcl.Diagnostic{
							Summary: "The user running Response does not have the permissions necessary to access the UI Directory.",
						})
					}

					return diag
				},
			},
			"auth_pages": &hcldec.ValidateSpec{
				Wrapped: &hcldec.AttrSpec{
					Name: "auth_pages",
					Type: cty.String,
				},
				Func: func(value cty.Value) hcl.Diagnostics {
					diag := hcl.Diagnostics{}

					if value.IsNull() {
						return hcl.Diagnostics{}
					}

					_, err := os.Stat(path.Join(value.AsString(), "base.html.tmpl"))
					if os.IsNotExist(err) {
						return diag.Append(&hcl.Diagnostic{
							Summary: "The Auth UI Directory provided was not found or does not contain an base.html.tmpl file, which is required to serve the Response Auth UI.",
						})
					}

					if os.IsPermission(err) {
						return diag.Append(&hcl.Diagnostic{
							Summary: "The user running Response does not have the permissions necessary to access the Auth UI Directory.",
						})
					}

					return diag
				},
			},
		},
	}
}
