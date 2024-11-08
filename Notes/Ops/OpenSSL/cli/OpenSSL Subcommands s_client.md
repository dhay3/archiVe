# OpenSSL Subcommands s_client

ref

https://www.feistyduck.com/library/openssl-cookbook/online/testing-with-openssl/connecting-to-tls-services.html

## Digest

syntax

```
openssl s_client [-help] [-ssl_config section] [-connect host:port] [-host hostname] [-port port] [-bind host:port] [-proxy host:port]
[-proxy_user userid] [-proxy_pass arg] [-unix path] [-4] [-6] [-servername name] [-noservername] [-verify depth] [-verify_return_error]
[-verify_quiet] [-verifyCAfile filename] [-verifyCApath dir] [-verifyCAstore uri] [-cert filename] [-certform DER|PEM|P12] [-cert_chain filename]
[-build_chain] [-CRL filename] [-CRLform DER|PEM] [-crl_download] [-key filename|uri] [-keyform DER|PEM|P12|ENGINE] [-pass arg] [-chainCAfile
filename] [-chainCApath directory] [-chainCAstore uri] [-requestCAfile filename] [-dane_tlsa_domain domain] [-dane_tlsa_rrdata rrdata]
[-dane_ee_no_namechecks] [-reconnect] [-showcerts] [-prexit] [-debug] [-trace] [-nocommands] [-security_debug] [-security_debug_verbose] [-msg]
[-timeout] [-mtu size] [-no_etm] [-keymatexport label] [-keymatexportlen len] [-msgfile filename] [-nbio_test] [-state] [-nbio] [-crlf] [-ign_eof]
[-no_ign_eof] [-psk_identity identity] [-psk key] [-psk_session file] [-quiet] [-sctp] [-sctp_label_bug] [-fallback_scsv] [-async] [-maxfraglen
len] [-max_send_frag] [-split_send_frag] [-max_pipelines] [-read_buf] [-ignore_unexpected_eof] [-bugs] [-comp] [-no_comp] [-brief]
[-legacy_server_connect] [-no_legacy_server_connect] [-allow_no_dhe_kex] [-sigalgs sigalglist] [-curves curvelist] [-cipher cipherlist]
[-ciphersuites val] [-serverpref] [-starttls protocol] [-name hostname] [-xmpphost hostname] [-name hostname] [-tlsextdebug] [-no_ticket]
[-sess_out filename] [-serverinfo types] [-sess_in filename] [-serverinfo types] [-status] [-alpn protocols] [-nextprotoneg protocols] [-ct]
[-noct] [-ctlogfile] [-keylogfile file] [-early_data file] [-enable_pha] [-use_srtp value] [-srpuser value] [-srppass value] [-srp_lateuser]
[-srp_moregroups] [-srp_strength number] [-nameopt option] [-no_ssl3] [-no_tls1] [-no_tls1_1] [-no_tls1_2] [-no_tls1_3] [-ssl3] [-tls1] [-tls1_1]
[-tls1_2] [-tls1_3] [-dtls] [-dtls1] [-dtls1_2] [-xkey infile] [-xcert file] [-xchain file] [-xchain_build file] [-xcertform DER|PEM]> [-xkeyform
DER|PEM]> [-CAfile file] [-no-CAfile] [-CApath dir] [-no-CApath] [-CAstore uri] [-no-CAstore] [-bugs] [-no_comp] [-comp] [-no_ticket]
[-serverpref] [-client_renegotiation] [-legacy_renegotiation] [-no_renegotiation] [-no_resumption_on_reneg] [-legacy_server_connect]
[-no_legacy_server_connect] [-no_etm] [-allow_no_dhe_kex] [-prioritize_chacha] [-strict] [-sigalgs algs] [-client_sigalgs algs] [-groups groups]
[-curves curves] [-named_curve curve] [-cipher ciphers] [-ciphersuites 1.3ciphers] [-min_protocol minprot] [-max_protocol maxprot]
[-record_padding padding] [-debug_broken_protocol] [-no_middlebox] [-rand files] [-writerand file] [-provider name] [-provider-path path]
[-propquery propq] [-engine id] [-ssl_client_engine id] [-allow_proxy_certs] [-attime timestamp] [-no_check_time] [-check_ss_sig] [-crl_check]
[-crl_check_all] [-explicit_policy] [-extended_crl] [-ignore_critical] [-inhibit_any] [-inhibit_map] [-partial_chain] [-policy arg]
[-policy_check] [-policy_print] [-purpose purpose] [-suiteB_128] [-suiteB_128_only] [-suiteB_192] [-trusted_first] [-no_alt_chains] [-use_deltas]
[-auth_level num] [-verify_depth num] [-verify_email email] [-verify_hostname hostname] [-verify_ip ip] [-verify_name name] [-x509_strict]
[-issuer_checks] [host:port]
```

`s_client` 是一个用于校验 TLS 证书的诊断工具

## Optional args

- `-connect host:port`

  指定需要校验的站点和端口

- `-bind host:port`

  指定使用源站点和端口

- `-servername name`

  设置 clientHello 中的 serverName Indication, 如果没有指定默认使用 `-connect` 中指定的 DNS 域名

- `-verify_quiet`

  只输出错误信息

- `-quiet`

  不输出证书和 session 相关的信息，只输出证书 DN 信息

- `-brief`

  只输出简短的信息

  ```
  
  openssl s_client -brief \
  -connect www.feistyduck.com:443 \
  -servername www.feistyduck.com
  CONNECTION ESTABLISHED
  Protocol version: TLSv1.2
  Ciphersuite: ECDHE-RSA-AES128-GCM-SHA256
  Peer certificate: CN = blog.ivanristic.com
  Hash used: SHA512
  Signature type: RSA
  Verification: OK
  Supported Elliptic Curve Point Formats: uncompressed:ansiX962_compressed_prime:ansiX962_compressed_char2
  Server Temp Key: ECDH, prime256v1, 256 bits
  ```

- `-verify_hostname hostname`

  校验指定 hostname，如果域名和证书不匹配会返回 62 status code (功能比较鸡肋)

  ```
  openssl s_client -connect www.feistyduck.com:443 -verify_hostname www.feistyduck1.com -brief
  depth=0 CN = blog.ivanristic.com
  verify error:num=62:hostname mismatch
  CONNECTION ESTABLISHED
  Protocol version: TLSv1.2
  Ciphersuite: ECDHE-RSA-AES128-GCM-SHA256
  Peer certificate: CN = blog.ivanristic.com
  Hash used: SHA512
  Signature type: RSA
  Verification error: hostname mismatch
  Supported Elliptic Curve Point Formats: uncompressed:ansiX962_compressed_prime:ansiX962_compressed_char2
  Server Temp Key: ECDH, prime256v1, 256 bits
  ```

  