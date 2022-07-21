package dego

func (d *degorater) makeJustifyFullSpacing(str string) string {
	return d.makeJustifySpacing(str, d.width)
}

func (d *degorater) makeJustifySpacing(str string, max int) string {
	str = d.sp + str

	l := len(str)
	for i := 0; i < max-l; i++ {
		str += d.sp
	}
	return str
}
