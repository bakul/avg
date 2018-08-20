package avg

import (
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

	pic2 =  []string{
                "0..0...1....",
                ".....1...1..",
                "0..0...1....",
                "2..3....A..A",
                "......4....9",
                "6..5........",
        }
)

func Test1(t *testing.T) {
	fig := ToFig(pic)
	if len(fig) != 1 {
		t.Errorf("expected %d, got %d:\n%v", len(pic), len(fig), fig)
	}
	fig = ToFig(pic1)
	t.Logf("%v", fig)
	if len(fig) != 5 {
		t.Errorf("expected %d, got %d:\n%v", len(pic), len(fig), fig)
	}
	fig = ToFig(pic2)
	t.Logf("%v", fig)
	if len(fig) != 5 {
		t.Errorf("expected %d, got %d:\n%v", len(pic), len(fig), fig)
	}
}
