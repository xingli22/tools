// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/tools/internal/lsp/command/generate"
)

func main() {
	content, err := generate.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	ioutil.WriteFile("command_gen.go", content, 0644)
}
