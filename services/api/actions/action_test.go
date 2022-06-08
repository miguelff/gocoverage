package actions

import "testing"

func TestTransform(t *testing.T) {
    res := transform("foo")
    if res != "transformed foo" {
        t.Error("fail")
    }
}
