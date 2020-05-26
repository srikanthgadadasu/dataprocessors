package main

import(
	"time"
	"readcsv/functions/csvtodf"
	"readcsv/vars/vars"
)

// Read the input csv file and converts to dataframe using the above schema

func main(){
	csvtodf.Csvtodataframe(Ipfile, Iptype)
}


