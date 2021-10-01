package reloader

import (
	"testing"

	"github.com/matryer/is"
)

func Test_Reload(t *testing.T) {
	is := is.New(t)
	r := New()

	var reloaded bool
	r.Register(func() error {
		reloaded = true

		return nil
	})

	err := r.Reload()
	is.NoErr(err)

	is.True(reloaded)
}
