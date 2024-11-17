package utilkit

// Refer: https://gist.github.com/yakuter/c0df0f4253ea639529f3589e99dc940b

import (
	"unsafe"
)

// b2s converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes converts a string to a byte slice without memory allocation.
//
// Warning: This function is unsafe and should be used with caution. Modifying
// the resulting byte slice may lead to undefined behavior if the original
// string is immutable. Ensure compatibility with future Go versions.
func StringToBytes(s string) []byte {
	// Get the pointer to the string data
	data := unsafe.StringData(s)

	// Convert to a byte slice using unsafe.Slice
	return unsafe.Slice(data, len(s))
}
