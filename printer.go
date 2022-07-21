package dego

import (
	"fmt"
	"math"
	"strconv"
)

func (d *degorater) printTopline() {
	var s string
	for i := 0; i < d.width; i++ {
		s += d.top
	}
	fmt.Printf(d.corner+"%s"+d.corner+"\n", s)
}

func (d *degorater) printUnderline() {
	d.printTopline()
}

func (d *degorater) printMessage(msg string) {
	// wrap message if does not fit
	l := len(msg)
	if l < d.width {
		fmt.Printf(d.side+"%s"+d.side+"\n", d.makeJustifyFullSpacing(msg))
	} else {
		wl := math.Ceil(float64(l) / float64(d.width))
		for i := 0; i < int(wl); i++ {
			s := i * d.width
			e := (i+1)*d.width - 1
			if e > l {
				e = l - 1
			}
			fmt.Printf(d.side+"%s"+d.side+"\n", d.makeJustifyFullSpacing(msg[s:e]))
		}
	}
}

func (d *degorater) printSelectors(lists []string) {
	l := len(lists)
	if l == 0 {
		return
	}

	fmt.Println(d.side + d.makeJustifyFullSpacing("") + d.side)
	if l < 4 {
		for i, line := range lists {
			no := strconv.Itoa(i + 1)
			fmt.Printf(d.side+"%s"+d.side+"\n", d.makeJustifyFullSpacing(no+". "+line))
		}
	} else {
		b, o, e := d.canTwoItemsAlignment(lists)
		if b {
			dw := d.width - o - e
			o += dw / 2
			e += dw / 2
			if dw%2 == 1 {
				o++
			}

			for i, line := range lists {
				no := strconv.Itoa(i + 1)
				if i%2 == 0 {
					fmt.Printf(d.side+"%s", d.makeJustifySpacing(no+". "+line, o))
				} else {
					fmt.Printf("%s"+d.side+"\n", d.makeJustifySpacing(no+". "+line, e))
				}
			}
		} else {
			for i, line := range lists {
				no := strconv.Itoa(i + 1)
				fmt.Printf(d.side+"%s"+d.side+"\n", d.makeJustifyFullSpacing(no+". "+line))
			}
		}
	}
}

func (d *degorater) canTwoItemsAlignment(lists []string) (b bool, e int, o int) {
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
		if o+e > d.width {
			return false, 0, 0
		}
	}
	b = true
	return
}

func max(n1 int, n2 int) int {
	if n1 >= n2 {
		return n1
	}
	return n2
}
