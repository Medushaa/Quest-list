package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct { //generic (can handle other types including todos)
	FileName string
}

//initiate storage instance with a filename
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName} //addr of the new storage struct
}

//save any data of type T into the file (json)
func (s *Storage[T]) Save(data T) error {
	//serialise go data into pretty json string ([]byte type)
	fileData, err := json.MarshalIndent(data, "", "")	
	//1st "" is prefix of every json line. 2nd "" is the identation
	if err != nil {
		return err
	}	
	return os.WriteFile(s.FileName, fileData, 0644) //create/edit in folder
}

func (s *Storage[T]) Load(data *T) error { //retrive
	fileData, err := os.ReadFile(s.FileName) //read the file
	if err != nil { //if file no exist
		return err
	}
	//json to go object data and put it in data (a pointer)
	return json.Unmarshal(fileData, data) 
}