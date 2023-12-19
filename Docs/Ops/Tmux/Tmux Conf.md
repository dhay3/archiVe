# Tmux Conf

## ~/.tmux.conf

```
(base) 0x00 in ~ λ cat ~/.tmux.conf
set -g mouse off
set -g default-terminal "screen-256color"
set -g mode-keys vi
set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'fabioluciano/tmux-tokyo-night'

run '~/.tmux/plugins/tpm/tpm'
```
