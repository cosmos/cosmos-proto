package zeropb

type Slice[T any] struct {
	offset int16
	length uint16
}
