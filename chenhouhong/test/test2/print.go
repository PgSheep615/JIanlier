/* package print
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"
func Print(s string) {
    cs := C.CString(s)
    defer C.free(unsafe.Pointer(cs))
    C.fputs(cs, (*C.FILE)(C.stdout))
} */
//pkgm
package test2

func Hello() string{
	return "hello"
}

