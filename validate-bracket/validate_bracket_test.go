package validate_bracket
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	tt := []struct {
		name string
		txt  string
		want bool
	}{
		{
			name: "valid 1",
			txt:  "(avc)",
			want: true,
		},
		{
			name: "valid 2",
			txt:  "(aa(b)c)",
			want: true,
		},
		{
			name: "valid 3",
			txt:  "(aa(b)c()()())",
			want: true,
		},
		{
			name: "valid 4",
			txt:  "(aazxc(safe)(derwr)(asdew)(sdfsf(fsda)))",
			want: true,
		},
		{
			name: "invalid 1",
			txt:  ")sdafas(",
			want: false,
		},
		{
			name: "invalid 2",
			txt:  ")sdafas(()",
			want: false,
		},
		{
			name: "invalid 3",
			txt:  "steast(",
			want: false,
		},
		{
			name: "invalid 4",
			txt:  "ssdaf())))))",
			want: false,
		},
		{
			name: "invalid 5",
			txt:  "sdsaf(asdfasf))df(",
			want: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidSimple(tc.txt)
			assert.Equal(t, tc.want, got)
		})
	}
}
