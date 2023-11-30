package zeropb

type List struct {
	slice
}

type String struct {
	slice
}

type Bytes struct {
	slice
}

type slice struct {
	offset int16
	length uint16
}
