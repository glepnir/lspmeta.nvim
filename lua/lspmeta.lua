local lspmeta = {}
local npcall = vim.F.npcall

-- get user lsp config from neovim/nvim-lspconfig
-- if not found lspconfig,try to load it
function lspmeta.get_user_config_server()
  local lspconfig = npcall(require,'lspconfig')
  if not lspconfig then
    -- check the packer exist
    if next(packer_plugins) ~= nil then
      vim.cmd [[packadd nvim-lspconfig]]
      lspconfig = require('lspconfig')
    else
      return
    end
  end

  local configs = require('lspconfig.configs')
  return vim.tbl_keys(configs)
end

return lspmeta
