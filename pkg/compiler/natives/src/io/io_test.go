// +build gijit

package io_test

import (
	"testing"
)

func TestMultiWriter_WriteStringSingleAlloc(t *testing.T) {
	t.Skip()
}

func TestMultiReaderFlatten(t *testing.T) {
	t.Skip("test relies on runtime.Callers and runtime.CallersFrames, which GopherJS doesn't support")
}

func TestMultiWriterSingleChainFlatten(t *testing.T) {
	t.Skip("test relies on runtime.Callers and runtime.CallersFrames, which GopherJS doesn't support")
}

func TestMultiReaderFreesExhaustedReaders(t *testing.T) {
	t.Skip("test relies on runtime.SetFinalizer, which GopherJS does not implement")
}
