set rtp+=~/workstation/vim/lspmeta.nvim
set rtp+=~/workstation/vim/nvim-lspconfig

lua <<EOF
require('lspconfig').gopls.setup{}
require('lspconfig').clangd.setup{}
EOF
