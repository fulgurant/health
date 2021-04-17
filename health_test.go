package health

import "testing"

func TestInterface(t *testing.T) {
	var _ IHealth = New(nil)
}
