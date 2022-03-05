package validate_bracket

const (
	leftB  = "("
	rightB = ")"
)

func isValidSimple(txt string) bool {
	var r int
	for i := 0; i < len(txt); i++ {
		c := txt[i : i+1]
		//for left
		if c == leftB {
			r += 1
			continue
		}
		//for right
		if c == rightB {
			if r == 0 {
				//if r is equal zero, it mean right brack come before the left.
				return false
			}
			r -= 1
		}
	}
	return !(r > 0)
}

