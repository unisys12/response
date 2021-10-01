package schema

import (
	"github.com/contaim/spec/parser"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// verify we meet the interface definition
var _ parser.BlockDefinition = (*developerDef)(nil)
var _ parser.NamedBlockDefinition = (*developerDef)(nil)

type developerDef struct{}

func (d *developerDef) Name() string {
	return "developer"
}

func (d *developerDef) Spec() hcldec.Spec {
	return &hcldec.ObjectSpec{
		"profiling": &hcldec.AttrSpec{
			Name: "profiling",
			Type: cty.String,
		},
		"ui_proxy_port": &hcldec.AttrSpec{
			Name: "ui_proxy_port",
			Type: cty.Number,
		},
	}
}
