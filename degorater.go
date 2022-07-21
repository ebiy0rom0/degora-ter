package dego

import (
	"syscall"

	"golang.org/x/term"
)

const defaultWidth int = 30

type Options struct {
	width  int
	sp     string
	corner string
	side   string
	top    string
}

type degorater struct {
	*Options
}

var dego *degorater

func init() {
	// default size is terminal width
	// if failed to got terminal size, use definition size
	_, sw, err := term.GetSize(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		sw = defaultWidth
	}

	dego = &degorater{
		&Options{
			width:  sw,
			sp:     " ",
			corner: "+",
			side:   "|",
			top:    "-",
		},
	}
}

func ChangeOptions(opt *Options) {
	if opt.width > 0 {
		ChangeWidth(opt.width)
	}
	// frame string can use single byte character only
	if len(opt.sp) == 1 {
		ChangeSp(opt.sp)
	}
	if len(opt.corner) == 1 {
		ChangeCorner(opt.corner)
	}
	if len(opt.side) == 1 {
		ChangeSide(opt.side)
	}
	if len(opt.top) == 1 {
		ChangeTop(opt.top)
	}
}

func ChangeWidth(width int) {
	dego.width = width
}

func ChangeSp(sp string) {
	dego.sp = sp
}

func ChangeCorner(corner string) {
	dego.corner = corner
}

func ChangeSide(side string) {
	dego.corner = side
}

func ChangeTop(top string) {
	dego.top = top
}

func Print(msg string, selectors []string) {
	dego.printTopline()
	dego.printMessage(msg)
	dego.printSelectors(selectors)
	dego.printUnderline()
}
