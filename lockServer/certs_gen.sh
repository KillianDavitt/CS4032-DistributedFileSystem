openssl req  -x509  -nodes -days 365 -newkey rsa:1024 -keyout lock.key.pem -out lock.crt.pem &&
        openssl x509 -pubkey -noout -in lock.crt.pem  > lock.pub.pem
