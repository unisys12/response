package schema

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type hideBannerDef struct{}

func (d *hideBannerDef) Name() string {
	return "hide_banner"
}

func (d *hideBannerDef) Spec() hcldec.Spec {
	return &hcldec.ValidateSpec{
		Wrapped: &hcldec.AttrSpec{
			Name:     "hide_banner",
			Type:     cty.Bool,
			Required: false,
		},
		Func: func(value cty.Value) hcl.Diagnostics {
			diag := hcl.Diagnostics{}

			if value.IsNull() {
				return hcl.Diagnostics{}
			}

			return diag
		},
	}
}

type hideServingDef struct{}

func (d *hideServingDef) Name() string {
	return "hide_serving"
}

func (d *hideServingDef) Spec() hcldec.Spec {
	return &hcldec.ValidateSpec{
		Wrapped: &hcldec.AttrSpec{
			Name:     "hide_serving",
			Type:     cty.Bool,
			Required: false,
		},
		Func: func(value cty.Value) hcl.Diagnostics {
			diag := hcl.Diagnostics{}

			if value.IsNull() {
				return hcl.Diagnostics{}
			}

			return diag
		},
	}
}
