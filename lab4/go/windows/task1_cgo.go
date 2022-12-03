package windows

/*
#cgo CFLAGS: -I${SRCDIR}/callC
#cgo LDFLAGS: ${SRCDIR}/callC/libC.a
#include<stdbool.h>
#include "task1.h"
*/
import "C"

func calculateC(selectedMaterial int, selectedSklop int,
	sillChecked bool, height float64, width float64) float64 {

	result := (float64)(C.main(C.int(selectedMaterial), C.int(selectedSklop),
		C.bool(sillChecked), C.double(height), C.double(width)))

	return result
}
