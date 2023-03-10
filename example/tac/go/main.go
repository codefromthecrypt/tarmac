// Tac is a small, simple Go program that is an example WASM module for Tarmac. This program will accept a Tarmac
// server request, log it, and echo back the payload in reverse.
package main

import (
	"fmt"
	wapc "github.com/wapc/wapc-guest-tinygo"
)

func main() {
	// Tarmac uses waPC to facilitate WASM module execution. Modules must register their custom handlers 
	wapc.RegisterFunctions(wapc.Functions{
		// Register request handler
		"handler": Handler,
	})
}

// Handler is the custom Tarmac Handler function that will receive a payload and
// must return a payload along with a nil error.
func Handler(payload []byte) ([]byte, error) {
	// Perform a host callback to log the incoming request
	_, err := wapc.HostCall("tarmac", "logger", "trace", []byte(fmt.Sprintf("Reversing Payload: %s", payload)))
	if err != nil {
		return []byte(""), fmt.Errorf("Unable to call callback - %s", err)
	}

	// Flip it and reverse
	if len(payload) > 0 {
		for i, n := 0, len(payload)-1; i < n; i, n = i+1, n-1 {
			payload[i], payload[n] = payload[n], payload[i]
		}
	}

	// Return the payload via a ServerResponse JSON
	return payload, nil
}
