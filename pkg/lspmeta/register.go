// Copyright 2021 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package pkg

import (
	"log"
	"os"

	"github.com/neovim/go-client/nvim"
	"github.com/neovim/go-client/nvim/plugin"
)

func Register(p *plugin.Plugin) error {
	stdin := os.Stdin
	stdout := os.Stdout
	v, err := nvim.New(stdin, stdout, stdout, log.Printf)
	if err != nil {
		return err
	}
	c := NewCommand(v)
	p.HandleCommand(&plugin.CommandOptions{
		Name:     "LspInstall",
		NArgs:    "",
		Range:    "",
		Count:    "",
		Addr:     "",
		Eval:     "",
		Complete: "",
		Bang:     true,
		Register: false,
		Bar:      false,
	}, func(args []string, bang bool) {
		c.LspInstall(args, bang)
	})
	return nil
}
