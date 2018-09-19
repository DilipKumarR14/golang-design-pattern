package main

import (
	"log"
	"encoding/json"
	"fmt"
)
type Message struct{
	Title string
	BOdy string
}
func main(){
	message:=Message{
		Title:"hello world",
		BOdy:"new mEssaGe",
	}
	data,err:=json.Marshal(message)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(string(data))
}