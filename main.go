package main 

import (
	"fmt"
	"log"
	"github.com/katyahohlova-dev/datafile"
)

func main(){
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil{
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers{
		sum += number
	}
	sampleCount := float64(len(numbers)) 
	fmt.Println("Average: %0.2f\n", sum/sampleCount)
}