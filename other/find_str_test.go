package other

import "testing"

func findStr(original, target string) int {
	var r int
	for i := 0; i < len(original); i++ {
		next := i + len(target)
		if next > len(original) {
			break
		}
		hook := original[i:next]
		if hook == target {
			r += 1
		}
	}
	return r
}
func TestFindStr(t *testing.T) {
	t.Parallel()
	want := 2
	got := findStr("ABCDECDECD", "CDE")
	if want != got {
		t.Errorf("want: %v, but got: %v", want, got)
	}
}
