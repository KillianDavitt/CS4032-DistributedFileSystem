package main

import (
	"net"
)

type server struct{
	ips[]net.IP
}

func (s *server) getIP() net.IP{
	ip := s.ips[0]
	s.ips = append(s.ips[1:], ip)
	return ip
}

type files struct{
	filesMap map[string]*server
}

func (f *files) addFile(filename string, ip net.IP){
	// Check if already exists
	_ = f.filesMap[filename]
	//if tst != nil {
	//	log.Fatal("File already exists!")
	//}
	
	newServer := &server{}
	newServer.ips = make([]net.IP, 5)
	f.filesMap[filename] = newServer

	// This might need to be atomic??

	///V/f := os.Open(filenam)
	//f.Write(ip)
	//f.Close()
	return
}

func (f *files) getFile(filename string) net.IP {
	return f.filesMap[filename].getIP()
}

 
func readFiles() (*files){

	newFiles := &files{}
	newFiles.filesMap = make(map[string]*server)
	newFiles.filesMap["test.txt"] = &server{}
	newFiles.filesMap["test.txt"].ips = make([]net.IP, 1, 20)
	newFiles.filesMap["test.txt"].ips[0] = net.ParseIP("10.1.0.10")
	return newFiles
}
