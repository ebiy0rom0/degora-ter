package infra

import (
	"fmt"
	"math"
	"strconv"
)

// 行当たり最大文字数
const maxLength int = 50

// スペーシング用文字列
const spacer string = " "

// 装飾用 文字
const (
	corner  string = "+"
	sideBar string = "|"
	topBar  string = "-"
)

func DecorateDescription(title string, selectors []string) {
	displayHeadline()
	displayTitle(title)
	displaySelectors(selectors)
	displayBottomline()
}

func displayHeadline() {
	var s string
	for i := 0; i < maxLength; i++ {
		s += topBar
	}
	fmt.Printf(corner+"%s"+corner+"\n", s)
}

func displayBottomline() {
	displayHeadline()
}

func displayTitle(str string) {
	// タイトルが表示領域を越える場合に折り返す
	l := len(str)
	if l < maxLength {
		fmt.Printf(sideBar+"%s"+sideBar+"\n", makeJustifyFullSpacing(str))
	} else {
		wl := math.Ceil(float64(l) / float64(maxLength))
		for i := 0; i < int(wl); i++ {
			s := i * maxLength
			e := (i+1)*maxLength - 1
			if e > l {
				e = l - 1
			}
			fmt.Printf(sideBar+"%s"+sideBar+"\n", makeJustifyFullSpacing(str[s:e]))
		}
	}
}

func displaySelectors(lists []string) {
	l := len(lists)
	if l == 0 {
		return
	}

	fmt.Println(sideBar + makeJustifyFullSpacing("") + sideBar)
	if l < 4 {
		for i, line := range lists {
			no := strconv.Itoa(i + 1)
			fmt.Printf(sideBar+"%s"+sideBar+"\n", makeJustifyFullSpacing(no+". "+line))
		}
	} else {
		// 1行 = 2項目で構成するか
		b, o, e := is(lists)
		if b {
			d := maxLength - o - e
			o += d / 2
			e += d / 2
			if d%2 == 1 {
				o++
			}

			for i, line := range lists {
				no := strconv.Itoa(i + 1)
				if i%2 == 0 {
					fmt.Printf(sideBar+"%s", makeJustifySpacing(no+". "+line, o))
				} else {
					fmt.Printf("%s"+sideBar+"\n", makeJustifySpacing(no+". "+line, e))
				}
			}
		} else {
			for i, line := range lists {
				no := strconv.Itoa(i + 1)
				fmt.Printf(sideBar+"%s"+sideBar+"\n", makeJustifyFullSpacing(no+". "+line))
			}
		}
	}
}

func is(lists []string) (b bool, e int, o int) {
	for i, line := range lists {
		if i%2 == 0 {
			if e == 0 {
				e = len(line)
			} else {
				e = max(e, len(line))
			}
		} else {
			if o == 0 {
				o = len(line)
			} else {
				o = max(o, len(line))
			}
		}
		if o+e > maxLength {
			return false, 0, 0
		}
	}
	b = true
	return
}

func makeJustifyFullSpacing(str string) string {
	return makeJustifySpacing(str, maxLength)
}

func makeJustifySpacing(str string, max int) string {
	str = spacer + str

	l := len(str)
	for i := 0; i < max-l; i++ {
		str += spacer
	}
	return str
}

func max(n1 int, n2 int) int {
	if n1 >= n2 {
		return n1
	}
	return n2
}
