// Package primitive is a support library for generic Go primitives.
//
// GooseLang provides models for all of these operations.
package primitive

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// UInt64Get converts the first 8 bytes of p (in little-endian order) to a
// uint64.
//
// Requires p be at least 8 bytes long.
func UInt64Get(p []byte) uint64 {
	return binary.LittleEndian.Uint64(p)
}

// UInt32Get converts the first 4 bytes of p (in little endian order) to a
// uint32.
//
// Requires p be at least 4 bytes long.
func UInt32Get(p []byte) uint32 {
	return binary.LittleEndian.Uint32(p)
}

// UInt64Put stores n to the first 8 bytes of p in little-endian order.
//
// Requires p to be at least 8 bytes long.
func UInt64Put(p []byte, n uint64) {
	binary.LittleEndian.PutUint64(p, n)
}

// UInt32Put stores n to the first 4 bytes of p in little-endian order.
//
// Requires p to be at least 4 bytes long.
func UInt32Put(p []byte, n uint32) {
	binary.LittleEndian.PutUint32(p, n)
}

// RandomUint64 returns a random uint64 using the global seed.
func RandomUint64() uint64 {
	return rand.Uint64()
}

// UInt64ToString formats a number as a string.
//
// Assumed to be pure and injective in the Coq model.
func UInt64ToString(x uint64) string {
	return fmt.Sprintf("%d", x)
}

// Linearize does nothing.
//
// Translates to an atomic step that supports opening invariants conveniently for
// the sake of executing a simulation fancy update at the linearization point of
// a procedure.
func Linearize() {}

// Assume lets the proof assume that `c` is true.
//
// In Go, if the assumption is violated this function will panic, whereas in the
// GooseLang model it will loop infinitely.
func Assume(c bool) {
	if !c {
		panic("Assume condition violated")
	}
}

// Assert induces a proof obligation that `c` is true.
//
// The Go implementation will panic (quit the process in a controlled manner) if
// `c` is not true. In GooseLang, it will make the machine stuck, i.e., cause UB.
//
// Using `panic()` directly is preferred (which is also modeled as the machine
// getting stuck), unless the extra control flow is unsupported.
func Assert(c bool) {
	if !c {
		panic("Assert condition violated")
	}
}

// Exit terminates the program with the given exit code.
//
// Modeled as an infinite loop since no more steps will be taken.
func Exit(n uint64) {
	os.Exit(int(n))
}

// WaitTimeout is like cond.Wait(), but waits for a maximum time of timeoutMs
// milliseconds.
//
// Not provided by sync.Cond, so we have to (inefficiently) implement this
// ourselves.
func WaitTimeout(cond *sync.Cond, timeoutMs uint64) {
	done := make(chan struct{})
	go func() {
		cond.Wait()
		cond.L.Unlock()
		close(done)
	}()
	select {
	case <-time.After(time.Duration(timeoutMs) * time.Millisecond):
		// timed out
		cond.L.Lock()
		return
	case <-done:
		// Wait returned
		cond.L.Lock()
		return
	}
}

// TimeNow returns the current time in nanoseconds.
func TimeNow() uint64 {
	return uint64(time.Now().UnixNano())
}

// Sleep waits for ns nanoseconds.
//
// Modeled as a no-op.
func Sleep(ns uint64) {
	time.Sleep(time.Duration(ns) * time.Nanosecond)
}
