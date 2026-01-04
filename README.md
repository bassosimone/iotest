# Golang Utilities for IO Testing

[![GoDoc](https://pkg.go.dev/badge/github.com/bassosimone/iotest)](https://pkg.go.dev/github.com/bassosimone/iotest) [![Build Status](https://github.com/bassosimone/iotest/actions/workflows/go.yml/badge.svg)](https://github.com/bassosimone/iotest/actions) [![codecov](https://codecov.io/gh/bassosimone/iotest/branch/main/graph/badge.svg)](https://codecov.io/gh/bassosimone/iotest)

The `iotest` Go package contains code to test `io` types (e.g., `io.Writer`, `io.Reader`).

For example:

```Go
import (
	"errors"
	"fmt"

	"github.com/bassosimone/iotest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Create a writer that always fails with a mocked error
w := &FuncWriter{
	WriteFunc: func(b []byte) (int, error) {
		return 0, errors.New("mocked error")
	},
}

// Use the writer with a function that will fail if it fails
count, err := fmt.Fprintf(w, "%s\n", "Hello, world!\n")

// Verify the test
require.Error(t, err)
assert.Equal(t, count, 0)
```

## Installation

To add this package as a dependency to your module:

```sh
go get github.com/bassosimone/iotest
```

## Development

To run the tests:
```sh
go test -v .
```

To measure test coverage:
```sh
go test -v -cover .
```

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```

## History

Adapted from [ooni/probe-cli](https://github.com/ooni/probe-cli/tree/v3.20.1/internal/mocks).
