package video

import "fmt"

func Generator(meta* Meta) error{
	n:= meta.NumberOfFrames
		
	if n==0{
	meta.NumberOfFrames = 8
	meta.TimeStamps = divideDuration(8,meta.Duration)	
		return nil
	}
	meta.TimeStamps = divideDuration(n,meta.Duration)	
	fmt.Println("The timestamp are  decided as ",meta.TimeStamps)
	return  nil
}


func divideDuration(n int,dur float64) []float64{
	timestamps := make([]float64,0,n)
	for i:=range n{
		timestamp:=(dur/float64(n+1))*float64(i+1)
		timestamps = append(timestamps, timestamp)
	}	
	return  timestamps
}
