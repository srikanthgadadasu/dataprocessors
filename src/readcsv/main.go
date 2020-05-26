package main

import(
	"time"
	"github.com/srikanthgadadasu/nadi/src/readcsv/functions/csvtodf"
	"github.com/srikanthgadadasu/nadi/src/readcsv/schemas/ipschemas"
)

// Read the input csv file and converts to dataframe using the above schema

func main(){
	Ipfile := "/udemy_courses.csv"
	csvtodf.csvtodataframe(Ipfile, ipschemas.Iptype)
}


