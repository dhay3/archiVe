# IPtables-extensions SYNPROXY Target

ref

https://wiki.nftables.org/wiki-nftables/index.php/Synproxy

## Digest

## Optional args

- `--mss maximum_segment_size`

  maximum segment size announced to clients. This must match the backend

- `--wscale window_scale`

  window scale announced to clients. This must match the backend

- `--sack-perm`

  pass client selective acknowledgement option to backend( will be disabled if not present )

- `--timestamps`

  pass client timestamp option to backend ( will be disabled if not present, also needed for selective acknowledgment and window scalling )

