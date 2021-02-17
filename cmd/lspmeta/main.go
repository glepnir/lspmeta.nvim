// Copyright 2021 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	plugin.Main(func(p *plugin.Plugin) error {
		return nil
	})

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-termChan
}
