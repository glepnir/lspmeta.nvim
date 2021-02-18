// Copyright 2021 glepnir. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package pkg

import (
	"log"

	"github.com/neovim/go-client/nvim"
)

type Command struct {
	v *nvim.Nvim
}

func NewCommand(v *nvim.Nvim) *Command {
	return &Command{}
}

type userLspServers []string

func (c *Command) LspInstall() {
	userConfig := userLspServers{}
	print("here")
	getUserConfig := "require('lspmeta').get_user_config_server()"
	err := c.v.ExecLua(getUserConfig, userConfig)
	print(userConfig)
	if err != nil {
		log.Fatal("[LspMeta] Get User config failed")
	}
}
