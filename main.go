package main

// #include "bridge.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type Datagram struct {
	len C.int
	results unsafe.Pointer
	Name *C.char
	MessageType C.int
}

const MODULE_NAME = "helloworld"

func CreateDatagram(results []byte) *Datagram {
	return &Datagram{
		len: C.int(len(results)),
		results: C.CBytes(results),
		Name: C.CString(MODULE_NAME),
		MessageType: C.int(0),
	}
}


func doStuffWithData(data []byte) {
	fmt.Printf("Library says: Received data of %d length, and it was: %s\n", len(data), data)
}

//export helloworldCallback
func helloworldCallback(cData unsafe.Pointer) {
	var dg *Datagram
	dg = (*Datagram)(cData)
	bData := C.GoBytes(dg.results, dg.len)
	// Could potentially route this data based on dg.MessageType
	doStuffWithData(bData)
}

//export helloworld
func helloworld(cData unsafe.Pointer, appCallbackPtr unsafe.Pointer, resultsPtr unsafe.Pointer) {
	// cData is a datagram pointer, appCallbackPtr is the application core callback function pointer,
	// and resultsPtr is a pointer to a resultant datagram to populate.
	var dg *Datagram
	dg = (*Datagram)(cData)
	data := C.GoBytes(dg.results, dg.len)
	fmt.Printf("Library received data of %d length: %s\n", len(data), data)
	// Prepare data to send to application core callback
	sModData := fmt.Sprintf("Data from %s library!", MODULE_NAME)
	bModData := []byte(sModData)
	cbData := CreateDatagram(bModData)
	// Cast the appCallbackPtr to the callback function defined in bridge.h
	f := C.callbackFunc(appCallbackPtr)
	// Invoke the callback pointer using our bridge function
	C.bridge_callback(f, unsafe.Pointer(cbData))
	// could potentially block and wait for data from application core
	// before continuing to populate results structure
	sRet := fmt.Sprintf("Return value from %s", MODULE_NAME)
	bRet := []byte(sRet)
	dataRet := CreateDatagram(bRet)
	tmp := (*Datagram)(resultsPtr)
	tmp.len = dataRet.len
	tmp.results = dataRet.results
}

func main() {}
