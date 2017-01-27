#openssl genrsa -out ca.key.pem 1024 &&  openssl req -subj '/CN=www.mydom.com/O=My Company Name LTD./C=US'\
#      -key ca.key.pem -new -x509 -days 365  -out ca.cert.pem


openssl req  -x509  -nodes -days 365 -newkey rsa:1024 -keyout fs.key.pem -out fs.crt.pem &&
        openssl x509 -pubkey -noout -in fs.crt.pem  > fs.pub.pem
