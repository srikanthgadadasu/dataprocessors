package main

import(
	"time"
	"./utilities"
)

// Schema of the input csv file.  This needs to be a parameter and comes from data pipeline

type iptype struct{
	course_id int
	course_title string
	url	string
	is_paid	bool
	price float64
	num_subscribers int	
	num_reviews	int
	num_lectures int
	level string
	content_duration float64	
	published_timestamp time.Time
	subject string
}

// Read the input csv file and converts to dataframe using the above schema

func main(){
	ipfile := "/Users/sripri/Downloads/udemy_courses.csv"

	csvtodf.CsvToDataframe(ipfile, iptype)
}


