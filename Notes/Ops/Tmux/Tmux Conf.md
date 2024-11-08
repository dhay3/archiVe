# Tmux Conf

## ~/.tmux.conf

```
# Plugins
# Use prefix + I to install plugins
set-option -g @plugin 'tmux-plugins/tpm'
set-option -g @plugin 'tmux-plugins/tmux-sensible'
# Use prefix + Ctrl + s to save tmux status
# Use prefix + Ctrl + r to restore tmux status
set-option -g @plugin 'tmux-plugins/tmux-resurrect'
set-option -g @plugin 'fabioluciano/tmux-tokyo-night'

run '~/.tmux/plugins/tpm/tpm'

# Global Opotions
set-option -gq prefix C-x
set-option -gq prefix2 None
set-option -gq default-terminal "screen-256color"
set-option -gq mode-keys vi
set-option -gq monitor-bell off
set-option -gq renumber-windows on
set-option -gq set-clipboard on
set-option -gq set-titles on
set-option -gq mouse on
set-option -gq status on
set-option -gq status-position bottom
set-option -gq base-index 1
set-option -gq pane-base-index 1
set-option -gq pane-border-lines simple
set-option -gp pane-border-indicators both
set-option -gq pane-border-style 'bg=default fg=#AAFF00'
set-option -gq pane-active-border-style 'bg=default fg=#FF1493 bold'
#xclip for xorg, wl-copy for wayland
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {set-option -gq copy-command 'wl-copy -n'}
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {set-option -gq copy-command 'xclip -i'}


# Bindings
#unbind-key -a

bind-key -T prefix C-b send-prefix
bind-key -T prefix Space next-layout
bind-key -T prefix h list-keys -N
bind-key -T prefix H list-keys
bind-key -T prefix s source-file ~/.tmux.conf
bind-key -T prefix ? display-message
bind-key -T prefix : command-prompt
bind-key -T prefix q display-panes
bind-key -T prefix [ previous-window
bind-key -T prefix ] next-window
bind-key -T prefix \{ swap-window -d -t -1
bind-key -T prefix \} swap-window -d -t +1
#bind-key -T prefix \{ swap-pane -U
#bind-key -T prefix \} swap-pane -D
bind-key -T prefix ( switch-client -p
bind-key -T prefix ) switch-client -n
bind-key -T prefix < display-menu -T "#[align=centre]#{window_index}:#{window_name}" -t = -x W -y W "#{?#{>:#{session_windows},1},,-}Swap Left" l { swap-window -t :-1 } "#{?#{>:#{session_windows},1},,-}Swap Right" r { swap-window -t :+1 } "#{?pane_marked_set,,-}Swap Marked" s { swap-window } '' Kill k { kill-window } Respawn x { respawn-window -k } "#{?pane_marked,Unmark,Mark}" m { select-pane -m } Rename r { command-prompt -F -I "#W" { rename-window -t "#{window_id}" "%%" } } '' "New After" C { new-window -a } "New At End" c { new-window }
bind-key -T prefix > display-menu -T "#[align=centre]#{pane_index} (#{pane_id})" -x P -y P "#{?#{m/r:(copy|view)-mode,#{pane_mode}},Go To Top,}" < { send-keys -X history-top } "#{?#{m/r:(copy|view)-mode,#{pane_mode}},Go To Bottom,}" > { send-keys -X history-bottom } '' "#{?mouse_word,Search For #[underscore]#{=/9/...:mouse_word},}" C-r { if-shell -F "#{?#{m/r:(copy|view)-mode,#{pane_mode}},0,1}" "copy-mode -t=" ; send-keys -X -t = search-backward "#{q:mouse_word}" } "#{?mouse_word,Type #[underscore]#{=/9/...:mouse_word},}" C-y { copy-mode -q ; send-keys -l "#{q:mouse_word}" } "#{?mouse_word,Copy #[underscore]#{=/9/...:mouse_word},}" c { copy-mode -q ; set-buffer "#{q:mouse_word}" } "#{?mouse_line,Copy Line,}" l { copy-mode -q ; set-buffer "#{q:mouse_line}" } '' "Horizontal Split" h { split-window -h } "Vertical Split" v { split-window -v } '' "#{?#{>:#{window_panes},1},,-}Swap Up" u { swap-pane -U } "#{?#{>:#{window_panes},1},,-}Swap Down" d { swap-pane -D } "#{?pane_marked_set,,-}Swap Marked" s { swap-pane } '' Kill X { kill-pane } Respawn R { respawn-pane -k } "#{?pane_marked,Unmark,Mark}" m { select-pane -m } "#{?#{>:#{window_panes},1},,-}#{?window_zoomed_flag,Unzoom,Zoom}" z { resize-pane -Z }
bind-key -T prefix y copy-mode
bind-key -T prefix p paste-buffer -p
bind-key -T prefix x respawn-window
bind-key -T prefix X respawn-pane
bind-key -T prefix l choose-tree -Zw
bind-key -T prefix L choose-tree -Zs
bind-key -T prefix c new-window
bind-key -T prefix C command-prompt { new-window -n "%%" }
bind-key -T prefix C-a command-prompt { new-window -an "%%" }
bind-key -T prefix C-b command-prompt { new-window -bn "%%" }
bind-key -T prefix r command-prompt -I "#W" { rename-window "%%" }
bind-key -T prefix R command-prompt -I "#S" { rename-session "%%" }
bind-key -T prefix k confirm-before -p "kill-window #W? (y/n)" kill-window
bind-key -T prefix K confirm-before -p "kill-session #S? (y/n)" kill-session
bind-key -T prefix C-k confirm-before -p "kill-server? (y/n)" kill-server
bind-key -T prefix w split-window -h -c "#{pane_current_path}"
bind-key -T prefix W split-window -v -c "#{pane_current_path}"
bind-key -T prefix m command-prompt -T target { move-window -t "%%" }
bind-key -T prefix M command-prompt -T target { move-pane -t "%%" }
bind-key -T prefix z resize-pane -Z
bind-key -T prefix 0 select-window -t :=0
bind-key -T prefix 1 select-window -t :=1
bind-key -T prefix 2 select-window -t :=2
bind-key -T prefix 3 select-window -t :=3
bind-key -T prefix 4 select-window -t :=4
bind-key -T prefix 5 select-window -t :=5
bind-key -T prefix 6 select-window -t :=6
bind-key -T prefix 7 select-window -t :=7
bind-key -T prefix 8 select-window -t :=8
bind-key -T prefix 9 select-window -t :=9

#MouseDown1 == rightclick
#MouseDown3 == Leftclick
bind-key -T root WheelUpPane if-shell -F "#{||:#{pane_in_mode},#{mouse_any_flag}}" { send-keys -M } { copy-mode -e }
bind-key -T root MouseDown1Status select-window -t =
bind-key -T root MouseDown3Status display-menu -T "#[align=centre]#{window_index}:#{window_name}" -t = -x W -y W "#{?#{>:#{session_windows},1},,-}Swap Left" l { swap-window -t :-1 } "#{?#{>:#{session_windows},1},,-}Swap Right" r { swap-window -t :+1 } "#{?pane_marked_set,,-}Swap Marked" s { swap-window } '' Kill k { kill-window } Respawn x { respawn-window -k } "#{?pane_marked,Unmark,Mark}" m { select-pane -m } Rename r { command-prompt -F -I "#W" { rename-window -t "#{window_id}" "%%" } } '' "New After" C { new-window -a } "New At End" c { new-window }
bind-key -T root MouseDown1Pane select-pane -t = \; send-keys -M
bind-key -T root MouseDrag1Pane select-pane -t = \; if-shell -F "#{||:#{pane_in_mode},#{mouse_any_flag}}" { send-keys -M } { copy-mode }
bind-key -T root MouseDrag1Border resize-pane -M
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {bind-key -T root DoubleClick1Pane select-pane -t = \; copy-mode -M \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"}
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {bind-key -T root TripleClick1Pane select-pane -t = \; copy-mode -M \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"}
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {bind-key -T root DoubleClick1Pane select-pane -t = \; copy-mode -M \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"}
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {bind-key -T root TripleClick1Pane select-pane -t = \; copy-mode -M \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"}
bind-key -r -T prefix Up select-pane -U
bind-key -r -T prefix Down select-pane -D
bind-key -r -T prefix Left select-pane -L
bind-key -r -T prefix Right select-pane -R
bind-key -r -T prefix M-Up resize-pane -U 5 
bind-key -r -T prefix M-Down resize-pane -D 5
bind-key -r -T prefix M-Left resize-pane -L 5
bind-key -r -T prefix M-Right resize-pane -R 5

bind-key -T copy-mode C-c send-keys -X cancel
bind-key -T copy-mode / command-prompt -T search -p "(search down)" { send-keys -X search-forward "%%" }
bind-key -T copy-mode n send-keys -X search-again
bind-key -T copy-mode N send-keys -X search-reverse
bind-key -T copy-mode v send-keys -X begin-selection
bind-key -T copy-mode q send-keys -X cancel
bind-key -T copy-mode Escape send-keys -X clear-selection
bind-key -T copy-mode Space send-keys -X begin-selection
bind-key -T copy-mode MouseUp1Pane send-keys -X clear-selection
bind-key -T copy-mode MouseDrag1Pane select-pane \; send-keys -X begin-selection
bind-key -T copy-mode WheelUpPane select-pane \; send-keys -X -N 1 scroll-up
bind-key -T copy-mode WheelDownPane select-pane \; send-keys -X -N 1 scroll-down
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {bind-key -T copy-mode Enter send-keys -X copy-pipe-and-cancel "wl-copy && wl-paste -n | wl-copy -p"}
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {bind-key -T copy-mode MouseDragEnd1Pane select-pane \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"} 
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {bind-key -T copy-mode DoubleClick1Pane select-pane \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"}
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {bind-key -T copy-mode TripleClick1Pane select-pane \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"}
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {bind-key -T copy-mode Enter send-keys -X copy-pipe-and-cancel "xclip -selection clipboard -in"}
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {bind-key -T copy-mode MouseDragEnd1Pane select-pane \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"}
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {bind-key -T copy-mode DoubleClick1Pane select-pane \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"}
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {bind-key -T copy-mode TripleClick1Pane select-pane \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"}

bind-key -T copy-mode-vi C-c send-keys -X cancel
bind-key -T copy-mode-vi / command-prompt -T search -p "(search down)" { send-keys -X search-forward "%%" }
bind-key -T copy-mode-vi n send-keys -X search-again
bind-key -T copy-mode-vi N send-keys -X search-reverse
bind-key -T copy-mode-vi v send-keys -X begin-selection
bind-key -T copy-mode-vi q send-keys -X cancel
bind-key -T copy-mode-vi Escape send-keys -X clear-selection
bind-key -T copy-mode-vi Space send-keys -X begin-selection
bind-key -T copy-mode-vi MouseUp1Pane send-keys -X clear-selection
bind-key -T copy-mode-vi MouseDrag1Pane select-pane \; send-keys -X begin-selection
bind-key -T copy-mode-vi WheelUpPane select-pane \; send-keys -X -N 1 scroll-up
bind-key -T copy-mode-vi WheelDownPane select-pane \; send-keys -X -N 1 scroll-down
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {bind-key -T copy-mode-vi Enter send-keys -X copy-pipe-and-cancel "wl-copy && wl-paste -n | wl-copy -p"}
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {bind-key -T copy-mode-vi MouseDragEnd1Pane select-pane \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"} 
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {bind-key -T copy-mode-vi DoubleClick1Pane select-pane \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"}
if-shell  "[[ ${XDG_SESSION_TYPE} == x11 ]]" {bind-key -T copy-mode-vi TripleClick1Pane select-pane \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"}
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {bind-key -T copy-mode-vi Enter send-keys -X copy-pipe-and-cancel "xclip -selection clipboard -in"}
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {bind-key -T copy-mode-vi MouseDragEnd1Pane select-pane \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"}
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {bind-key -T copy-mode-vi DoubleClick1Pane select-pane \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"}
if-shell  "[[ ${XDG_SESSION_TYPE} == wayland ]]" {bind-key -T copy-mode-vi TripleClick1Pane select-pane \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"}
```

**references**

[^1]:https://gist.github.com/mzmonsour/8791835
