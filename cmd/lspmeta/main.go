// Copyright 2021 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/glepnir/lspmeta.nvim/pkg"
	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	plugin.Main(pkg.Register)
}
