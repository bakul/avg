package avg

import (
	"avg"
	//"fmt"
	"testing"
)

var (
	pic = []string{
		"0123456789",
		"ABCDEFGHJK",
		"LMNPQRSTUV",
		"WXYZabcdef",
		"ghjklmnpqr",
		"stuvwxyz",
	}

	pic1 = []string{
		"0..0....1....3..",
		"......1...1.....",
		"0..0....1.......",
		"....7....4...5..",
		".........7..9...",
	}
)

func Test1(t *testing.T) {
	fig := avg.ToFig(pic)
	if len(fig) != 1 {
		t.Errorf("expected %d, got %d:\n%v", len(pic), len(fig), fig)
	}
	fig = avg.ToFig(pic1)
	t.Logf("%v", fig)
	if len(fig) != 5 {
		t.Errorf("expected %d, got %d:\n%v", len(pic), len(fig), fig)
	}
}
