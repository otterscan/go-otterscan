package sigs

import (
	"os"
	"path"
	"testing"
)

type testCase struct {
	hex  string
	want string
}

var tc_noname = []testCase{}
var tc_withname = []testCase{}

func init() {
	root := "../_hack/gen/4bytes/signatures"
	dirs, err := os.ReadDir(path.Clean(root))
	if err != nil {
		panic(err)
	}
	for _, v := range dirs {
		data, err := os.ReadFile(path.Join(root, v.Name()))
		if err != nil {
			panic(err)
		}
		tc_noname = append(tc_noname, testCase{
			hex:  v.Name(),
			want: string(data),
		})
	}
}

func init() {
	root := "../_hack/gen/4bytes/with_parameter_names"
	dirs, err := os.ReadDir(path.Clean(root))
	if err != nil {
		panic(err)
	}
	for _, v := range dirs {
		data, err := os.ReadFile(path.Join(root, v.Name()))
		if err != nil {
			panic(err)
		}
		tc_withname = append(tc_withname, testCase{
			hex:  v.Name(),
			want: string(data),
		})
	}
}

func TestNoName(t *testing.T) {
	for _, v := range tc_noname {
		w := v.want
		h := HexAbi(v.hex)
		if w != h {
			t.Errorf("wanted %s got %s", w, h)
		}
	}
}
func TestWithName(t *testing.T) {
	for _, v := range tc_withname {
		w := v.want
		h := HexNamed(v.hex)
		if w != h {
			t.Errorf("wanted %s got %s", w, h)
		}
	}
}

func BenchmarkRead10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx, v := range tc_noname {
			if idx > 10 {
				break
			}
			Hex(v.hex)
		}
	}
}

func BenchmarkRead500000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx, v := range tc_noname {
			if idx > 5000000 {
				break
			}
			Hex(v.hex)
		}
	}
}
