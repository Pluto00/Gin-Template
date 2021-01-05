package e

import "testing"

func TestGetMsg(t *testing.T) {
	t.Log(GetMsg(SUCCESS))
	t.Log(GetMsg(ERROR))
}
