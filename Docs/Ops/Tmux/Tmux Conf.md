# Tmux Conf

## ~/.tmux.conf

```
# Plugins
set-option -g @plugin 'tmux-plugins/tpm'
set-option -g @plugin 'tmux-plugins/tmux-sensible'
set-option -g @plugin 'fabioluciano/tmux-tokyo-night'

run '~/.tmux/plugins/tpm/tpm'

# Global Opotions
set-option -gq default-terminal "screen-256color"
set-option -gq prefix C-b
set-option -gq prefix2 None
set-option -gq mouse on
set-option -gq base-index 1
set-option -gq pane-base-index 1
set-option -gq renumber-windows on
set-option -gq mode-keys vi
set-option -gq set-clipboard on
set-option -gq copy-command 'xclip -i'

# Bindings
unbind-key '"'
unbind-key %

bind-key -T prefix C-b send-prefix
bind-key -T prefix ? list-keys -N
bind-key -T prefix h list-keys
bind-key -T prefix : command-prompt
bind-key -T prefix q display-panes
bind-key -T prefix [ previous-window
bind-key -T prefix ] next-window
bind-key -T prefix ( switch-client -p
bind-key -T prefix ) switch-client -n
bind-key -T prefix y copy-mode
bind-key -T prefix p paste-buffer -p
bind-key -T prefix x respawn-window
bind-key -T prefix X respawn-pane
bind-key -T prefix l choose-tree -Zw
bind-key -T prefix L choose-tree -Zs
bind-key -T prefix c new-window
bind-key -T prefix r command-prompt -I #W "rename-window '%%'"
bind-key -T prefix R command-prompt -I #S "rename-session '%%'"
bind-key -T prefix k confirm-before -p "kill-window #W? (y/n)" kill-window
bind-key -T prefix K confirm-before -p "kill-pane #P? (y/n)" kill-pane
bind-key -T prefix w split-window -h
bind-key -T prefix W split-window -v
bind-key -T prefix z resize-pane -Z

bind-key -T copy-mode-vi v send-keys -X begin-selection
bind-key -T copy-mode-vi Enter send-keys -X copy-pipe-and-cancel "xclip -selection clipboard -in"
bind-key -T copy-mode-vi MouseUp1Pane send-keys -X clear-selection
bind-key -T copy-mode-vi q send-keys -X cancel
bind-key -T copy-mode-vi MouseDrag1Pane select-pane \; send-keys -X begin-selection
bind-key -T copy-mode-vi MouseDragEnd1Pane select-pane \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in" 
bind-key -T copy-mode-vi DoubleClick1Pane select-pane \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in" 
bind-key -T copy-mode-vi TripleClick1Pane select-pane \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"
bind-key -T copy-mode-vi MouseDragEnd1Pane send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"
bind-key -T copy-mode-vi WheelUpPane select-pane \; send-keys -X -N 1 scroll-up
bind-key -T copy-mode-vi WheelDownPane select-pane \; send-keys -X -N 1 scroll-down
bind-key -T root WheelUpPane if-shell -F "#{||:#{pane_in_mode},#{mouse_any_flag}}" { send-keys -M } { copy-mode -e }
bind-key -T root MouseDrag1Pane select-pane -t = \; if-shell -F "#{||:#{pane_in_mode},#{mouse_any_flag}}" { send-keys -M } { copy-mode -H }
bind-key -T root DoubleClick1Pane select-pane -t = \; if-shell -F "#{||:#{pane_in_mode},#{mouse_any_flag}}" { send-keys -M } { copy-mode -H ; send-keys -X select-word ; run-shell -d 0.3 ; send-keys -X copy-pipe-and-no-clear "xclip -selection clipboard -in" }
bind-key -T root TripleClick1Pane select-pane -t = \; if-shell -F "#{||:#{pane_in_mode},#{mouse_any_flag}}" { send-keys -M } { copy-mode -H ; send-keys -X select-line ; run-shell -d 0.3 ; send-keys -X copy-pipe-and-no-clear "xclip -selection clipboard -in" }
```

**references**

[^1]:https://gist.github.com/mzmonsour/8791835
