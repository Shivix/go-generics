package optional

import (
    "testing"
)

func TestRun(t *testing.T) {
    if runk[int32]() != 5 {
        t.Error("expected 5")
    }
    var test Optional[int32]
    if *test.value != 0 {
        t.Error("expected 3")
    }
}
