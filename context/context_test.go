package context

import (
	"testing"
)

func TestGetLocals_OK(t *testing.T) {
	res, err := GetLocals()

	if err != nil {
		t.Errorf("fails to get locals %s", err)
		t.FailNow()
	}

	t.Logf("Result: %v", res)

}
