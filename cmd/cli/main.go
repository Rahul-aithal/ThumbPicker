package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Rahul-aithal/ThumbPicker/internals/video"
)



func main(){
	if len(os.Args)<2{
	fmt.Println("Video path must be given") 
		return
	}
	file:= os.Args[1]
	fmt.Printf("The Video will be proccessed %s\n",file)
	filepath,err:= filepath.Abs(file)
	if err!=nil{
		fmt.Println("No file found",err)
	}
	video.Service(filepath,8)
}
