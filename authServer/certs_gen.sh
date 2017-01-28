openssl req  -x509  -nodes -days 365 -newkey rsa:1024 -keyout auth.key.pem -out auth.crt.pem &&
        openssl x509 -pubkey -noout -in auth.crt.pem  > auth.pub.pem
