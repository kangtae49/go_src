# go_src

# .bash_profile
export GOROOT=/usr/local/go

export PATH=$PATH:$GOROOT/bin

# .emacs
(set-terminal-coding-system 'utf-8)

(set-keyboard-coding-system 'utf-8)

(prefer-coding-system 'utf-8)

(setenv "PATH" "/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin:/usr/local/go/bin")

(setenv "GOPATH" "/usr/local/go/")

(setq exec-path (cons "/usr/local/go/bin" exec-path))

(add-to-list 'load-path "~/Misc/emacs/go-mode.el")

(require 'go-mode-autoloads)

(add-hook 'before-save-hook 'gofmt-before-save)




# 
git clone https://github.com/dominikh/go-mode.el.git
