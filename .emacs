(setenv "PATH" "/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin:/home/kangtae49/go/bin")
(setenv "GOROOT" "/home/kangtae49/go")
(setenv "GOPATH" "/home/kangtae49/src/go_src")

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
