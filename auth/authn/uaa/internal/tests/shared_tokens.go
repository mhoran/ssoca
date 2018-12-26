package tests

const SharedPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuzUaj3MHrNsm8X4JQs4b
qNpTnQmL57DWIaDe9wtRMVSqRZm7ovFkFBIooew/dQc3zfyFZo54t2IPu2cEeOdF
rcs2gpaXQAV9o5uYkZ6migckmQKj2fH8DZzG4aI0NCCFXwkS46kzBcavF1cOrAo4
4pKJQ4yoTOmLFq7Ac8urMnwM2mFhOY1FlizlwNY+W4SX572NNT3PKEDTmyd+9vLE
a8sMBAz034hUeBoXxBuJv18cwIaCs+sz2/v8iyPoB4D6C3hdJCkHWXDsCjs0zQyp
HRudXc9xuJ868LPjqpFtz5g2YZNHboMDP2yLNPTBdTta+u0hvRSWCPG711K8NoYU
QQIDAQAB
-----END PUBLIC KEY-----`

const SharedPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAuzUaj3MHrNsm8X4JQs4bqNpTnQmL57DWIaDe9wtRMVSqRZm7
ovFkFBIooew/dQc3zfyFZo54t2IPu2cEeOdFrcs2gpaXQAV9o5uYkZ6migckmQKj
2fH8DZzG4aI0NCCFXwkS46kzBcavF1cOrAo44pKJQ4yoTOmLFq7Ac8urMnwM2mFh
OY1FlizlwNY+W4SX572NNT3PKEDTmyd+9vLEa8sMBAz034hUeBoXxBuJv18cwIaC
s+sz2/v8iyPoB4D6C3hdJCkHWXDsCjs0zQypHRudXc9xuJ868LPjqpFtz5g2YZNH
boMDP2yLNPTBdTta+u0hvRSWCPG711K8NoYUQQIDAQABAoIBAQCUgeLpCxJ6oHs7
NYV5NnGguFl+CPWwPYsQsEJP7e1h2c0dW4ALGb0PYvxSv0vztV8ijADzQ3Al4aej
PWzB0I1o+fzrCTVA91ocsLEOS7HHazUwIHUEzr7puYXXtkn9VBMZOeR/L1knat65
ADnGZnGL4zz6lhloshfBTa2j7lyHtZz7bVLMgOlMzpltym7S/yi0MuVtkLxoHZ55
OC0ImwJg7EF0u6Pu7EXjiJiLt2UtkgZUgrnNE+5GzgH7MV8X6Z5TdyCVMCVBFPnh
eb0l1FdqIA5pXEmnroFFCb+HPBfCO0oM5Pj0rllnLRqiwAUh3ybkLcUuSpK452pD
r/AGCIBxAoGBAMLo/BUdMQHQ73Ah+4XFKKjpepNEm/dGSR7fYrPyjTfIUxb0cbCH
IpbWjfpfcSBqC1ruO47yhLQDGqfXO1OcZ/R5zoDvO4S73FTz/I2gXAU2ZT7zEbJb
uD35x3BOESyRlfi+sxuwa6x3+9eDLBuUA17YkENZULbEYgAX/w3N+ay1AoGBAPXi
FIaf+8hS/wAytOaZ4oamPXQLxewlawVhqLlso9q0syr9bl2WpuwXT2ZpVD/vuM90
JfXtj6sQA5+cxfT1j3mlE7vBWGQVu5MZ3upNYjy6kfyd8tbbJCGT3rbn0imk7XHg
E9IYguGd7q53PNpSeNYyeye6rpSG9mfFXOQMBYzdAoGAVXKEiNEaWfx568PKG7P6
hkqNJSDP61DR34n5IIB2Ln1RB+A0MzqbDFuVNK1A0AANaFjGnX1udehieDBgXuBT
ppb71ASNKDyfdFLeZkxl6xcUvcsb+ABDvk/RgkeHOe29PKaFqN7n5V1Yd/uWi6SQ
8bSpUg1G2Ekx2W6rryacjX0CgYBohaF6FPoADx/tIqD5wgKkgHhxIywQlVtoHzqN
EBybNONQ7AiFaiwtAj+zZiT1RN9H+LTGVSUUb6tKp2cTGWvTJxQ7ZRZA/WgOQRoc
eaa/8ANs4mu/X8p9J5GuKN+S6lQxazKhxd+lLvCuY2uP1y0cyLrHibj61GBcfm5d
VqrpVQKBgQDAuA7In/4hRaiggKcThXGBKektnmBmvG8Yow+zb41VLNB6Rvus0fJq
Bhq6wRt33OEKJADCHLQL8Au/4QdJzwMgyqlaYV8mcGLheFtiGC+JynJ7VSyxHf+X
ODuydwjQ/6ger8SJVemH2wO56ntTy9NKN6vj7F1Dbpegn+7jNA5VwQ==
-----END RSA PRIVATE KEY-----`

const SharedToken = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4MjkwMDE2MDAsImp0aSI6ImZha2UtdXVpZCIsImlhdCI6MTU0NTAwNDgwMCwiaXNzIjoiZmFrZS1vcmlnaW4iLCJuYmYiOjE1NDUwMDQ4MDAsImF1ZCI6WyJmYWtlLWF1ZGllbmNlMSJdLCJ1c2VyX25hbWUiOiJmYWtlLXVzZXIxIiwic2NvcGUiOlsic2NvcGUxIiwic2NvcGUyIl19.dNIrUNTa8APl9UUYNEODMT1GxWX1CQakIs4pDrwN2guRw8hjXqCrBijs02R5xM1vVqt7beHUjs5tQcviuD28ZVM5dntwKV-y0oHo1xmtkrvzlz35ECZkpjnK2HVx7GnMDMQECfnoTelB8_BGC-50pIBQ2ueZ2HmNuOzO5KTa464SxzNVCvYmj7FZcEUBZcguwUS2VvTpBSy_1324ASMJ0_8Jq0xb9VuyGPqUmooVStr3CGMbLJB0SwkDuw3DtH1Oqcywi1zzwECb6hq4VF2DHa_kivaht6_vYoXUG0M5YqAyRfmt1ahSh1QTBureHLolQXLwBwFexnAtJUCPjCKZnQ`