package config

import (
	"fmt"
	"github.com/contaim/spec"
	"github.com/hashicorp/hcl/v2"
	"github.com/responserms/response/internal/config/schema"
)

func New() (*Config, *spec.Diagnostics) {
	// New blank root
	cfg := newRoot()

	// Apply all the defaults and return a pointer to Config
	config, err := cfg.ToConfig()
	if err != nil {
		return nil, &spec.Diagnostics{
			Diags: hcl.Diagnostics{
				{
					Summary: "Unable to create Response Core configuration",
					Detail:  err.Error(),
				},
			},
		}
	}

	return config, &spec.Diagnostics{}
}

func NewFromGlobPath(glob string) (*Config, *spec.Diagnostics) {
	sch := spec.New(schema.Schema())
	ctx := schema.NewContext()

	if diags := sch.FileGlob(glob); diags.HasErrors() {
		return nil, diags
	}

	if diags := sch.Parse(ctx); diags.HasErrors() {
		return nil, diags
	}

	cfg := newRoot()
	if diags := sch.Decode(ctx, cfg); diags.HasErrors() {
		return nil, diags
	}

	config, err := cfg.ToConfig()
	if err != nil {
		fmt.Println(err.Error())

		return nil, &spec.Diagnostics{
			Diags: hcl.Diagnostics{
				{
					Summary: "Unable to create Response Core configuration",
					Detail:  err.Error(),
				},
			},
		}
	}

	return config, &spec.Diagnostics{}
}
