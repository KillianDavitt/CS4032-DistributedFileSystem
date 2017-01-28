openssl req  -x509  -nodes -days 365 -newkey rsa:1024 -keyout dir.key.pem -out dir.crt.pem &&
openssl x509 -pubkey -noout -in dir.crt.pem  > dir.pub.pem
