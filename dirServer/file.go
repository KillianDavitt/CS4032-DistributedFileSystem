package main

import (
	"net"
	"log"
	"encoding/json"
)
type file struct{
	Filename string
	Ip net.IP
	Hash []byte
}

func (f *file) MarshalFile() ([]byte){
	data, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func UnmarshalFile(data []byte) (file){
	var f file
	err := json.Unmarshal(data, &f)
	if err != nil{
		log.Fatal(err)
	}
	return f
}

func (f *file) UpdateRedisFile(){
	data := f.MarshalFile()
	fileClient := getFileRedis()
	err := fileClient.Set(string(f.Filename), string(data), 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func NewRedisFile(filename string, ip net.IP, hash []byte){
	newFile := &file{}
	newFile.Ip = ip
	newFile.Filename = filename
	newFile.Hash = hash
	newFile.UpdateRedisFile()
}
