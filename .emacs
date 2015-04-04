(set-terminal-coding-system 'utf-8)
(set-keyboard-coding-system 'utf-8)
(prefer-coding-system 'utf-8)

(setenv "PATH" "/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin:/usr/local/go/bin")
(setenv "GOPATH" "/usr/local/go/")

(setq exec-path (cons "/usr/local/go/bin" exec-path))
(add-to-list 'load-path "~/Misc/emacs/go-mode.el")
(require 'go-mode-autoloads)
(add-hook 'before-save-hook 'gofmt-before-save)

(defun my-go-mode-hook ()
  ; Call Gofmt before saving
  (add-hook 'before-save-hook 'gofmt-before-save)
  ; Customize compile command to run go build
  (if (not (string-match "go" compile-command))
      (set (make-local-variable 'compile-command)
           "go build -v && go test -v && go vet"))
  ; Godef jump key binding
  (local-set-key (kbd "M-.") 'godef-jump))
(add-hook 'go-mode-hook 'my-go-mode-hook)
