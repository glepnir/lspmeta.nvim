" Copyright 2016 The nvim-go Authors. All rights reserved.
" Use of this source code is governed by a BSD-style
" license that can be found in the LICENSE file.

let s:save_cpo = &cpo
set cpo&vim

if exists('g:load_lspmeta')
  finish
endif

let g:load_lspmeta = 1

let s:plugin_name   = 'lspmeta'
let s:plugin_cmd    = ['lspmeta']

function! s:JobStart(host) abort
  try
    return jobstart(s:plugin_cmd, {'rpc': v:true, 'detach': v:false})
  catch
    echomsg v:throwpoint
    echomsg v:exception
  endtry
endfunction

" plugin manifest
call remote#host#Register(s:plugin_name, 'x', function('s:JobStart'))

let &cpo = s:save_cpo
unlet s:save_cpo
