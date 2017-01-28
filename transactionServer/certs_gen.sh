openssl req  -x509  -nodes -days 365 -newkey rsa:1024 -keyout tran.key.pem -out tran.crt.pem &&
        openssl x509 -pubkey -noout -in tran..crt.pem  > tran..pub.pem
