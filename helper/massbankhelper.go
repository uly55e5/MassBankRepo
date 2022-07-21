package main

// #include "../../pwizmzlib/mz_pwiz_c_interface.h"
// #cgo LDFLAGS: -L../../pwizmzlib/dist/debug -lpwizmzlib -lzstd -lbz2 -llzma
// #cgo CFLAGS: -I../../pwizmzlib/libs/pwiz
import "C"

func main() {
	var msdata *C.struct_MSDataFile = C.openMSData(C.CString("../../test-data/mzML_Files/005_TOM_CYS.mzML"))
	idata := C.getInstrumentInfo(msdata)
	runinfo := C.getRunInfo(msdata)
	println(msdata, runinfo, idata)
}
