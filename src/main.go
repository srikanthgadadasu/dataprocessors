package main

import(
	"fmt"
	"github.com/srikanthgadadasu/nadi/src/readcsv"
)

func main(){
	Ipfile := "/Users/sriyan/Documents/nadi/udemy_courses.csv"
	rows := readcsv.Csvtodataframe(Ipfile)
	fmt.Println(rows[0]) //sample record
	fmt.Println(rows[0].Course_title)
}