package travel

/*
#cgo CFLAGS: -I${SRCDIR}/callC
#cgo LDFLAGS: ${SRCDIR}/callC/libC.a
#include<stdbool.h>
#include "task2.h"
*/
import "C"

func calculateC(days int64, Country int, Season int,
	guide bool, luxury bool) float64 {

	result := (float64)(C.main(C.int(days), C.int(Country), C.int(Season),
		C.bool(guide), C.bool(luxury)))

	return result
}
