//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/ooni/probe-cli/blob/v3.20.1/internal/mocks/writer.go
//

package iotest

import (
	"io"

	"github.com/bassosimone/runtimex"
)

// FuncWriter allows to mock any [io.FuncWriter].
type FuncWriter struct {
	WriteFunc func(b []byte) (int, error)
}

var _ io.Writer = &FuncWriter{}

// Write implements [io.Writer].
func (w *FuncWriter) Write(b []byte) (int, error) {
	runtimex.Assert(w.WriteFunc != nil)
	return w.WriteFunc(b)
}
