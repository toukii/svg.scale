package scale

import (
	"fmt"
	"strconv"

	"github.com/toukii/bytes"
	"github.com/toukii/goutils"
)

type Symbool string

const (
	DM     Symbool = "M"
	Dm     Symbool = "m"
	DL     Symbool = "L"
	Dl     Symbool = "l"
	DH     Symbool = "H"
	Dh     Symbool = "h"
	DV     Symbool = "V"
	Dv     Symbool = "v"
	DC     Symbool = "C"
	Dc     Symbool = "c"
	DS     Symbool = "S"
	Ds     Symbool = "s"
	DQ     Symbool = "Q"
	Dq     Symbool = "q"
	DT     Symbool = "T"
	Dt     Symbool = "t"
	DA     Symbool = "A"
	Da     Symbool = "a"
	DZ     Symbool = "Z"
	Dz     Symbool = "z"
	Dblank Symbool = " "
	Dcomma Symbool = ","
	Dminus Symbool = "-"
)

var (
	Symbools map[rune]Symbool = map[rune]Symbool{
		rune('M'): DM,
		rune('m'): Dm,
		rune('L'): DL,
		rune('l'): Dl,
		rune('H'): DH,
		rune('h'): Dh,
		rune('V'): DV,
		rune('v'): Dv,
		rune('C'): DC,
		rune('c'): Dc,
		rune('S'): DS,
		rune('s'): Ds,
		rune('Q'): DQ,
		rune('q'): Dq,
		rune('T'): DT,
		rune('t'): Dt,
		rune('A'): DA,
		rune('a'): Da,
		rune('Z'): DZ,
		rune('z'): Dz,
		rune(' '): Dblank,
		rune(','): Dcomma,
		rune('-'): Dminus,
	}
)

type Path struct {
	Symbool Symbool
	Value   float32
	islast  bool
}

func (p *Path) String() string {
	if p.islast && (p.Symbool == DZ || p.Symbool == Dz) {
		return string(p.Symbool)
	}
	return fmt.Sprintf("%s%+v", p.Symbool, p.Value)
}

func findSymbools(pathd []rune) []int {
	loc := make([]int, 0, 30)
	preloc := 0
	loc = append(loc, 0)
	for i, r := range pathd {
		if _, ex := Symbools[r]; ex && (preloc+1 < i) {
			if r == rune(' ') && i < len(pathd)-1 {
				// fmt.Println(string(r), string(pathd[i+1]))
				if _, ex := Symbools[pathd[i+1]]; ex {
					loc = append(loc, i+1)
					preloc = i + 1
					continue
				}
			}
			loc = append(loc, i)
			preloc = i
		}
	}
	return loc
}

func ConvPath(rs []rune, islast bool) []*Path {
	if len(rs) <= 0 {
		return nil
	}
	symb, ex := Symbools[rs[0]]
	pd := &Path{
		Symbool: symb,
	}
	str := string(rs[1:])
	if !ex {
		pd.Symbool = Dblank
	}
	if str == "" {
		pd.islast = true
		return []*Path{pd}
	}

	if rs[0] == rune('.') {
		str = fmt.Sprintf("0%s", string(rs))
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		// fmt.Printf("conv %s, err:%+v", str, err)
		loc := 0
		for i := loc; i < len(rs); i++ {
			if rs[i] == rune('.') {
				loc = i
			}
		}
		if loc > 0 {
			ps1 := ConvPath(rs[:loc], false)
			ps2 := ConvPath(rs[loc:], false)
			ps1 = append(ps1, ps2...)
			return ps1
		} else {
			panic(err)
		}
	}
	pd.Value = float32(f)
	return []*Path{pd}
}

func conv(pathd []rune, loc []int) []*Path {
	ps := make([]*Path, 0, len(loc))
	for i := 1; i < len(loc); i++ {
		ps = append(ps, ConvPath(pathd[loc[i-1]:loc[i]], false)...)
	}
	ps = append(ps, ConvPath(pathd[loc[len(loc)-1]:], true)...)
	return ps
}

func Scale(pathd string, scale float32) string {
	rs := []rune(pathd)
	loc := findSymbools(rs)
	ps := conv(rs, loc)
	for _, p := range ps {
		p.Value *= scale
	}
	return Path2D(ps)
}

func Path2D(ps []*Path) string {
	wr := bytes.NewWriter(make([]byte, 0, 2014))
	for _, p := range ps {
		wr.Write(goutils.ToByte(p.String()))
	}
	return goutils.ToString(wr.Bytes())
}
