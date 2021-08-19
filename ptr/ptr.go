// LICENSE: UNLICENSE - See LICENSE file

// Package ptr creates pointer from value inline.
//
// For all non-structual builtin types, including type alias:
//	var str *string = ptr.NewString("hello")
//	var n *int = ptr.NewInt(10)
//	var r *rune = ptr.NewRune('a')
// And for any type:
//	var str *interface{} = ptr.New("hello")
package ptr

func New(v interface{}) *interface{} {
	return &v
}
func NewBool(v bool) *bool {
	return &v
}
func NewUint8(v uint8) *uint8 {
	return &v
}
func NewUint16(v uint16) *uint16 {
	return &v
}
func NewUint32(v uint32) *uint32 {
	return &v
}
func NewUint64(v uint64) *uint64 {
	return &v
}
func NewInt8(v int8) *int8 {
	return &v
}
func NewInt16(v int16) *int16 {
	return &v
}
func NewInt32(v int32) *int32 {
	return &v
}
func NewInt64(v int64) *int64 {
	return &v
}
func NewFloat32(v float32) *float32 {
	return &v
}
func NewFloat64(v float64) *float64 {
	return &v
}
func NewComplex64(v complex64) *complex64 {
	return &v
}
func NewComplex128(v complex128) *complex128 {
	return &v
}
func NewString(v string) *string {
	return &v
}
func NewInt(v int) *int {
	return &v
}
func NewUint(v uint) *uint {
	return &v
}
func NewUintptr(v uintptr) *uintptr {
	return &v
}
func NewByte(v byte) *byte {
	return &v
}
func NewRune(v rune) *rune {
	return &v
}
