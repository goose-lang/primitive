// Package primitive is a support library for generic Go primitives.
//
// GooseLang provides models for all of these operations.
package primitive

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

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

// Mutex is a wrapper around sync.Mutex.
//
// This exists primarily to allow the channel model to be used to bootstrap
// goose: to avoid a circular dependency, the model's dependencies cannot use
// channels, so it's easiest if primitive provides locks rather than using sync.
type Mutex struct {
	m sync.Mutex
}

// Lock locks m
func (m *Mutex) Lock() {
	m.m.Lock()
}

// Unlock unlocks m (which should be held)
func (m *Mutex) Unlock() {
	m.m.Unlock()
}
