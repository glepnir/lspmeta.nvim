// Copyright 2021 glepnir. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package lspmeta

import (
	"github.com/neovim/go-client/nvim"
	"github.com/neovim/go-client/nvim/plugin"
)

type Command struct {
	nvim *nvim.Nvim
	// 	logger *logger.Logger
}

func NewCommand(v *nvim.Nvim) *Command {
	return &Command{
		nvim: v,
	}
}

func (c *Command) LspInstall(p *plugin.Plugin) ([]string, error) {
	userConfig := make([]string, 0)
	getUserConfig := "require('lspmeta').get_user_config_server()"
	if err := c.nvim.ExecLua(getUserConfig, userConfig); err != nil {
		return nil, err
	}
	return userConfig, nil
}
