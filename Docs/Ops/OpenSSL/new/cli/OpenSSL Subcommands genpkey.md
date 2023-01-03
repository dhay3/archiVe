ref
[https://www.feistyduck.com/library/openssl-cookbook/online/openssl-command-line/key-generation.html](https://www.feistyduck.com/library/openssl-cookbook/online/openssl-command-line/key-generation.html)
## Digest
syntax
```
openssl genpkey [-help] [-out filename] [-outform PEM|DER] [-pass arg]
[-cipher] [-engine id] [-paramfile file] [-algorithm alg] [-pkeyopt
opt:value] [-genparam] [-text]
```
genpkey 用于生成 private key
## Optional args

- `-out filename`

指定生成的 private key 文件名，如果没有指定直接输出到 stdout

- `-outform DER|PEM`

指定输出的文件格式，默认使用 PEM

- `-pass arg`

指定 private key 使用的 passphrase

- `-[cipher]`

指定 private key 使用的 cipher

- `-algorithm alg`

private key 使用的加密算法，例如 RSA, DSA, DH。必须在`-pkeyopt`之前指定

- `-pkeyopt opt:value`

设置 public key algorithm option opt to value，具体查看 Options

- `-text`

以格式化的形式输出 private key 使用的参数和 PEM 文件
## Options
> 这里只记录 RSA, 其他算法具体查看 KEY GENERATION OPTIONS / PARAMETER GENERATION OPTIONS

### RSA key generation options

- `rsa_keygen_bits:numbits`

私钥使用的加密比特数，默认 2048

- `rsa_keygen_primes:numprimes`

私钥使用的加密质数，默认 2
## Examples
```
#Generate an RSA private key using default parameters
openssl genpkey -algorithm RSA -out key.pem

#Generate an RSA private key using 4096 bits and aes-128-cbc cipher with 1234 passphrase
openssl genpkey -out fd.key -algorithm RSA -pkeyopt rsa_keygen_bits:4096 \
-aes-128-cbc -pass pass:1234


```
