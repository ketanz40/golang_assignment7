package main

import (
	"fmt"
	"assignment7/staticWebsite"
	"assignment7/fileServer"
)

func main(){
	var userChoice int 
	fmt.Println("Enter 1 for a Static Website\nEnter 2 for a File Server for a directory")
	fmt.Scanln(&userChoice)
	switch userChoice {
	case 1:
		staticWebsite.StaticFunc()
		break
	case 2:
		fileServer.FileServerFunc()
		break
	}
}