package x509Util

/*
Order of execution

1. Auth serv starts
2. Auth Serv reads ca file into DS
3. DS starts
4. DS reads in rsa keys
5. DS connects to AS
6. DS prints out the keyprint
7. DS accepts via cmd
8. AS accepts via cmd
9.

*/
func createTlsConfig() {

}
