package main

import "fmt"
import "unsafe"

// #cgo CFLAGS: -I${SRCDIR}/greetings
// #cgo LDFLAGS: ${SRCDIR}/greetings/build/liblanguage.a
// #include <stdlib.h>
// #include <greetings.h>
import "C"

func main() {
	fmt.Printf("Go says: static C library greetings coming..\n")
	john := C.CString("John")
	defer C.free(unsafe.Pointer(john))
	johannes := C.CString("Johannes")
	defer C.free(unsafe.Pointer(johannes))
	johnOz := C.CString("JohnOz")
	defer C.free(unsafe.Pointer(johnOz))

	C.greet_in_english(john)
	C.greet_in_german(johannes)

        ozReturn := C.greet_in_oz(johnOz)
        fmt.Printf("Go says: %s\n", C.GoString(ozReturn))
        C.free(unsafe.Pointer(johnOz))         // free memory right now ...
        defer C.free(unsafe.Pointer(ozReturn)) // ... or later

	fmt.Printf("Go says bye!\n")
}
