package main

import(
	"fmt"
	"github.com/srikanthgadadasu/nadi/src/readcsv"
)

func main(){
	Ipfile := "github.com/srikanthgadadasu/nadi/blob/master/udemy_courses.csv"
	rows := readcsv.Csvtodataframe(Ipfile)
	fmt.Println(rows[0]) //sample record
	fmt.Println(rows[0].Course_title)
}