package contx

import (
	"context"
	"testing"
)

func Test_GetSet(t *testing.T) {
	ctx := context.Background()
	set := 1
	got := GetUserID(SetUserID(ctx, set))

	if got != set {
		t.Fatalf("Returned ID is not the same as ID given to SetUserID")
	}
}
