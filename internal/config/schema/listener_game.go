package schema

import (
	"fmt"
	"github.com/contaim/spec/parser"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"strings"
)

// verify we meet the interface definition
var _ parser.BlockDefinition = (*gameDef)(nil)
var _ parser.NamedBlockDefinition = (*gameDef)(nil)

type gameDef struct{}

func (d *gameDef) Name() string {
	return "game_listener"
}

func (d *gameDef) Spec() hcldec.Spec {
	return &hcldec.BlockSpec{
		TypeName: "game_listener",
		Nested: &hcldec.ValidateSpec{
			Wrapped: &hcldec.ObjectSpec{
				"bind_address": &hcldec.AttrSpec{
					Name: "bind_address",
					Type: cty.String,
				},
				"port": &hcldec.AttrSpec{
					Name: "port",
					Type: cty.Number,
				},
				"tls": &hcldec.BlockSpec{
					TypeName: "tls",
					Nested: &hcldec.ValidateSpec{
						Wrapped: &hcldec.ObjectSpec{
							"port": &hcldec.AttrSpec{
								Name: "port",
								Type: cty.Number,
							},
							"cert_path": &hcldec.AttrSpec{
								Name: "cert_path",
								Type: cty.String,
							},
							"key_path": &hcldec.AttrSpec{
								Name: "key_path",
								Type: cty.String,
							},
							"auto": &hcldec.BlockSpec{
								TypeName: "auto",
								Nested: &hcldec.ObjectSpec{
									"production": &hcldec.AttrSpec{
										Name: "production",
										Type: cty.Bool,
									},
									"email": &hcldec.AttrSpec{
										Name:     "email",
										Type:     cty.String,
										Required: true,
									},
									"domains": &hcldec.AttrSpec{
										Name:     "domains",
										Type:     cty.Set(cty.String),
										Required: true,
									},
									"dns": &hcldec.BlockSpec{
										TypeName: "dns",
										Nested: &hcldec.ObjectSpec{
											"solver": &hcldec.ValidateSpec{
												Wrapped: &hcldec.BlockLabelSpec{
													Index: 0,
													Name:  "solver",
												},
												Func: func(value cty.Value) hcl.Diagnostics {
													var allowedSolvers = []string{"cloudflare", "vultr", "digitalocean"}

													if value.IsNull() {
														return hcl.Diagnostics{}
													}

													valStr := value.AsString()
													for _, solver := range allowedSolvers {
														if solver == valStr {
															return hcl.Diagnostics{}
														}
													}

													return hcl.Diagnostics{
														{
															Severity: hcl.DiagError,
															Summary:  fmt.Sprintf("The provided solver (%q) is not a valid solver type", valStr),
															Detail:   fmt.Sprintf("The solver type must be one of: %s", strings.Join(allowedSolvers, ", ")),
														},
													}
												},
											},
											"token": &hcldec.AttrSpec{
												Name:     "token",
												Type:     cty.String,
												Required: true,
											},
										},
									},
								},
							},
						},
						Func: func(value cty.Value) hcl.Diagnostics {
							if value.IsNull() {
								return hcl.Diagnostics{}
							}

							v := value.AsValueMap()
							if !v["cert_path"].IsNull() && (v["key_path"].IsNull() || v["key_path"].AsString() == "") {
								return hcl.Diagnostics{
									{
										Severity: hcl.DiagError,
										Summary:  "The `key_path` attribute must be set when `cert_path` is set",
										Detail: "You provided a `cert_path` but did not provide the corresponding `key_path` attribute. Set the" +
											"`key_path` attribute equal to the relative or absolute path to your certificate's key file.",
									},
								}
							}

							if !v["key_path"].IsNull() && (v["cert_path"].IsNull() || v["cert_path"].AsString() == "") {
								return hcl.Diagnostics{
									{
										Severity: hcl.DiagError,
										Summary:  "The `cert_path` attribute must be set when `key_path` is set",
										Detail: "You provided a `key_path` but did not provide the corresponding `cert_path` attribute. Set the" +
											"`cert_path` attribute equal to the relative or absolute path to your certificate file.",
									},
								}
							}

							if v["cert_path"].IsNull() && v["key_path"].IsNull() && v["auto"].IsNull() {
								return hcl.Diagnostics{
									{
										Severity: hcl.DiagError,
										Summary:  "TLS is enabled but no certificate is configured",
										Detail: "You provided the `tls` block but did not provide a static certificate or configure automatic " +
											"TLS. You must configure a static certificate or configure automatic TLS. To disable TLS remove the `tls` block altogether.",
									},
								}
							}

							return hcl.Diagnostics{}
						},
					},
				},
			},
			Func: func(value cty.Value) hcl.Diagnostics {
				m := value.AsValueMap()

				if !m["tls"].IsNull() {
					tls := m["tls"].AsValueMap()

					if m["port"].IsNull() || tls["port"].IsNull() {
						return hcl.Diagnostics{}
					}

					return hcl.Diagnostics{
						{
							Severity: hcl.DiagError,
							Summary:  "Only one port may be specified for `game_listener`",
							Detail: "You provided a `port` and a `tls.port` but the `game_listener` may only " +
								"listen on a single port at a time. If you intend to use TLS, remove the root `port` from the `game_listener` block.",
						},
					}
				}

				return hcl.Diagnostics{}
			},
		},
	}
}
