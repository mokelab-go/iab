#!bin/sh

# before executing this script, please make your private key
# $ openssl genrsa 1024 > private_key.pem

openssl dgst -sha1 -sign private_key.pem $1 | base64 > $1_sign64.sig  
