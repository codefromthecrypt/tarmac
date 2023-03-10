// This program is a test program used to facilitate unit testing with Tarmac.
package main

import (
	"fmt"
	wapc "github.com/wapc/wapc-guest-tinygo"
)

func main() {
	// Tarmac uses waPC to facilitate WASM module execution. Modules must register their custom handlers 
	wapc.RegisterFunctions(wapc.Functions{
    "handler": Handler,
	})
}

func Handler(payload []byte) ([]byte, error) {
  fmt.Printf("Output from WASM Modules goes here\n")

	// Log the payload
	_, err := wapc.HostCall("tarmac", "logger", "info", []byte(`Testdata Function Starting Execution`))
	if err != nil {
		return []byte(""), fmt.Errorf("Unable to call host callback - %s", err)
	}

	// KVStore
	_, err = wapc.HostCall("tarmac", "kvstore", "set", []byte(`{"key":"test-data","data":"aSBhbSBhIGxpdHRsZSB0ZWFwb3Q="}`))
	if err != nil {
		return []byte(""), fmt.Errorf(`Failed to call host callback - %s`, err)
	}

	_, err = wapc.HostCall("tarmac", "kvstore", "get", []byte(`{"key":"test-data"}`))
	if err != nil {
		return []byte(""), fmt.Errorf(`Failed to call host callback - %s`, err)
	}

	// SQL Query
	_, err = wapc.HostCall("tarmac", "sql", "query", []byte(`{"query":"Q1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgd2FzbWd1ZXN0ICggaWQgaW50IE5PVCBOVUxMLCBuYW1lIHZhcmNoYXIoMjU1KSwgUFJJTUFSWSBLRVkgKGlkKSApOw=="}`))
	if err != nil {
		return []byte(""), fmt.Errorf(`Failed to call host callback - %s`, err)
	}

	// Return a happy message
	return []byte("Howdie"), nil
}
