package sayajim

import (
	"testing"
)

func TestYouIsSayajim(t *testing.T) {
	waitTest := true
	receiveTest := YouIsSayajim(701, "terra")

	if waitTest != receiveTest {
		t.Error("Fail Test")
	}
}
