// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"text/template"
	"time"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

var (
	authInternallyKnownOAuthProviders = []string{
		"discord",
		"twitch",
	}
)

const (
	DiagSessionSecretEncoding       = "The session secret provided is not base64 encoded"
	DiagSessionSecretEncodingDetail = "Session secrets must be base64 encoded to prevent the need for escaping when passing the key in via" +
		"environment variables and/or config files."

	DiagSessionSecretBelowLength  = "The session secret provided is less than 64 bytes in length"
	DiagSessionSecretAboveLength  = "The session secret provided is more than 64 bytes in length"
	DiagSessionSecretLengthDetail = "Session secrets must be between 32 and 64 bytes in length to enable session signing. You can use" +
		"the 'response operator keygen' command to generate a suitable session secret."

	SessionSecretByteRequirement = 64
)

type authDef struct{}

func (s *authDef) Name() string {
	return "auth"
}

func (s *authDef) Spec() hcldec.Spec {
	return &hcldec.BlockSpec{
		TypeName: s.Name(),
		Required: false,
		Nested: &hcldec.ValidateSpec{
			Wrapped: &hcldec.ObjectSpec{
				"session_root_url": &hcldec.ValidateSpec{
					Wrapped: &hcldec.AttrSpec{
						Name: "session_root_url",
						Type: cty.String,
					},
					Func: validateURL,
				},
				"session_domain": &hcldec.AttrSpec{
					Name: "session_domain",
					Type: cty.String,
				},
				"session_secret": &hcldec.ValidateSpec{
					Wrapped: &hcldec.AttrSpec{
						Name: "session_secret",
						Type: cty.String,
					},
					Func: func(value cty.Value) hcl.Diagnostics {
						diag := hcl.Diagnostics{}

						if value.IsNull() {
							return hcl.Diagnostics{}
						}

						key, err := base64.StdEncoding.DecodeString(value.AsString())
						if err != nil {
							diag = append(diag, &hcl.Diagnostic{
								Severity: hcl.DiagError,
								Summary:  DiagSessionSecretEncoding,
								Detail:   DiagSessionSecretEncodingDetail,
							})

							// not valid anyway, fail early
							return diag
						}

						if len(key) < SessionSecretByteRequirement {
							diag = append(diag, &hcl.Diagnostic{
								Severity: hcl.DiagError,
								Summary:  DiagSessionSecretBelowLength,
								Detail:   DiagSessionSecretLengthDetail,
							})
						}

						if len(key) > SessionSecretByteRequirement {
							diag = append(diag, &hcl.Diagnostic{
								Severity: hcl.DiagError,
								Summary:  DiagSessionSecretAboveLength,
								Detail:   DiagSessionSecretLengthDetail,
							})
						}

						return diag
					},
				},
				"session_duration": &hcldec.ValidateSpec{
					Wrapped: &hcldec.AttrSpec{
						Name: "session_duration",
						Type: cty.String,
					},
					Func: func(value cty.Value) hcl.Diagnostics {
						if value.IsNull() {
							return hcl.Diagnostics{}
						}

						_, err := time.ParseDuration(value.AsString())
						if err != nil {
							return hcl.Diagnostics{
								{
									Severity: hcl.DiagError,
									Summary:  "Value is not a valid duration string",
									Detail:   fmt.Sprintf("Response was unable to parse the duration string. Duration strings should be valid Go Duration strings.\n %s", err.Error()),
								},
							}
						}

						return hcl.Diagnostics{}
					},
				},
				"password": &hcldec.BlockSpec{
					TypeName: "password",
					Nested: &hcldec.ObjectSpec{
						"disabled": &hcldec.AttrSpec{
							Name: "disabled",
							Type: cty.Bool,
						},
						"bcrypt_cost": &hcldec.ValidateSpec{
							Wrapped: &hcldec.AttrSpec{
								Name: "bcrypt_cost",
								Type: cty.Number,
							},
							Func: func(value cty.Value) hcl.Diagnostics {
								if value.IsNull() {
									return hcl.Diagnostics{}
								}

								if value.LessThan(cty.NumberIntVal(4)) == cty.True {
									return hcl.Diagnostics{
										{
											Severity: hcl.DiagError,
											Summary:  "Hashing cost below minimum of 4",
											Detail:   "You've provided a hashing cost that is below the minimum allowable cost of 4.",
										},
									}
								}

								if value.GreaterThan(cty.NumberIntVal(31)) == cty.True {
									return hcl.Diagnostics{
										{
											Severity: hcl.DiagError,
											Summary:  "Hashing cost above maximum allowable value of 31",
											Detail:   "You've provided a hashing cost that is above the maximum allowable cost of 31.",
										},
									}
								}

								return hcl.Diagnostics{}
							},
						},
					},
				},
				"oauth2": &hcldec.ValidateSpec{
					Wrapped: &hcldec.BlockMapSpec{
						TypeName:   "oauth2",
						LabelNames: []string{"provider"},
						Nested: &hcldec.ObjectSpec{
							//"provider": &hcldec.BlockLabelSpec{
							//	Name:  "provider",
							//	Index: 0,
							//},
							"name": &hcldec.AttrSpec{
								Name: "name",
								Type: cty.String,
							},
							"icon": &hcldec.AttrSpec{
								Name: "icon",
								Type: cty.String,
							},
							"update_profile_on_login": &hcldec.AttrSpec{
								Name: "update_profile_on_login",
								Type: cty.Bool,
							},
							"client_id": &hcldec.AttrSpec{
								Name:     "client_id",
								Type:     cty.String,
								Required: true,
							},
							"client_secret": &hcldec.AttrSpec{
								Name:     "client_secret",
								Type:     cty.String,
								Required: true,
							},
							"scopes": &hcldec.AttrSpec{
								Name: "scopes",
								Type: cty.Set(cty.String),
							},
							"authorization_url": &hcldec.ValidateSpec{
								Wrapped: &hcldec.AttrSpec{
									Name: "authorization_url",
									Type: cty.String,
								},
								Func: validateURL,
							},
							"token_url": &hcldec.ValidateSpec{
								Wrapped: &hcldec.AttrSpec{
									Name: "token_url",
									Type: cty.String,
								},
								Func: validateURL,
							},
							"userinfo_url": &hcldec.ValidateSpec{
								Wrapped: &hcldec.AttrSpec{
									Name: "userinfo_url",
									Type: cty.String,
								},
								Func: validateURL,
							},
							"userinfo_map": &hcldec.BlockSpec{
								TypeName: "userinfo_map",
								Nested: &hcldec.ObjectSpec{
									"id": &hcldec.ValidateSpec{
										Wrapped: &hcldec.AttrSpec{
											Name: "id",
											Type: cty.String,
										},
										Func: validateMapTemplate,
									},
									"name": &hcldec.ValidateSpec{
										Wrapped: &hcldec.AttrSpec{
											Name: "name",
											Type: cty.String,
										},
										Func: validateMapTemplate,
									},
									"email": &hcldec.ValidateSpec{
										Wrapped: &hcldec.AttrSpec{
											Name: "email",
											Type: cty.String,
										},
										Func: validateMapTemplate,
									},
									"metadata": &hcldec.AttrSpec{
										Name: "metadata",
										Type: cty.Map(cty.String),
									},
								},
							},
						},
					},
					Func: func(value cty.Value) hcl.Diagnostics {
						if value.IsNull() {
							return hcl.Diagnostics{}
						}

						oauth := value.AsValueMap()

						for provider, _ := range oauth {
							for _, wellKnown := range authInternallyKnownOAuthProviders {
								if wellKnown == provider {
									return hcl.Diagnostics{}
								}
							}
						}

						diags := hcl.Diagnostics{}
						if oauth["authorization_url"].IsNull() {
							diags = diags.Append(&hcl.Diagnostic{
								Severity: hcl.DiagError,
								Summary:  "Authorization URL is required for unknown OAuth 2.0 providers",
								Detail:   "Response supports any OAuth 2.0 provider, however, you'll need to specify an `authorization_url` for providers that Response does not provide internal integrations with.",
							})
						}

						if oauth["token_url"].IsNull() {
							diags = diags.Append(&hcl.Diagnostic{
								Severity: hcl.DiagError,
								Summary:  "Token URL is required for unknown OAuth 2.0 providers",
								Detail:   "Response supports any OAuth 2.0 provider, however, you'll need to specify a `token_url` for providers that Response does not provide internal integrations with.",
							})
						}

						if oauth["userinfo_url"].IsNull() {
							diags = diags.Append(&hcl.Diagnostic{
								Severity: hcl.DiagError,
								Summary:  "Userinfo URL is required for unknown OAuth 2.0 providers",
								Detail:   "Response supports any OAuth 2.0 provider, however, you'll need to specify a `userinfo_url` for providers that Response does not provide internal integrations with.",
							})
						}

						if oauth["userinfo_map"].IsNull() {
							diags = diags.Append(&hcl.Diagnostic{
								Severity: hcl.DiagError,
								Summary:  "Userinfo Map is required for unknown OAuth 2.0 providers",
								Detail:   "Response supports any OAuth 2.0 provider, however, you'll need to manually provide a mapping between the provider's User object and an internal Response User. This allows Response to map/update user data when a user logs in with this provider.",
							})
						}

						return diags
					},
				},
			},
			Func: func(value cty.Value) hcl.Diagnostics {
				if value.IsNull() {
					return hcl.Diagnostics{}
				}

				values := value.AsValueMap()
				if !values["oauth2"].IsNull() && (values["session_root_url"].IsNull() || values["session_root_url"].AsString() == "") {
					return hcl.Diagnostics{
						{
							Severity: hcl.DiagError,
							Summary:  "Missing `session_root_url` in `auth` block.",
							Detail:   "You must provide a `session_root_url` when configuring `oauth2` in order to ensure that consistent OAuth 2.0 redirect URLs are generated. This attribute should be equal to the root of your Response URL, such as https://response.mydomain.com. Consult the Response auth documentation for help, if needed.",
						},
					}
				}

				return nil
			},
		},
	}
}

func validateURL(v cty.Value) hcl.Diagnostics {
	if v.IsNull() {
		return hcl.Diagnostics{}
	}

	u, err := url.Parse(v.AsString())
	if err != nil {
		return hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Attribute must be a valid URL",
				Detail:   fmt.Sprintf("This attribute must be a valid URL. Check and fix the attribute's value.\n %s", err.Error()),
			},
		}
	}

	var diags hcl.Diagnostics
	if u.Scheme == "" {
		diags = diags.Append(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Attribute is missing URL scheme",
			Detail:   "The provided URL should contain a scheme. Specify either http or https in the URL.",
		})
	}

	if u.Host == "" {
		diags = diags.Append(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Attribute is missing Host",
			Detail:   "The provided URL should contain a host. Specify the full host of the URL, including the port if necessary.",
		})
	}

	return diags
}

func validateMapTemplate(value cty.Value) hcl.Diagnostics {
	if value.IsNull() {
		return hcl.Diagnostics{}
	}

	if value.AsString() == "" {
		return hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "An empty string is not valid here",
				Detail:   "All mapping templates should be produce a valid value to associate with the Response User. See the docs for more info.",
			},
		}
	}

	_, err := template.New("").
		Parse(value.AsString())

	if err != nil {
		return hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Invalid mapping template applied",
				Detail:   fmt.Sprintf("Response is unable to use the provided template. Check your template syntax.\n %s", err.Error()),
			},
		}
	}

	return hcl.Diagnostics{}
}
