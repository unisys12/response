package response

import "github.com/hashicorp/hcl/v2"

func NewWithConfigGlob(glob string) (*Core, hcl.Diagnostics) {
	return &Core{}, nil
}
