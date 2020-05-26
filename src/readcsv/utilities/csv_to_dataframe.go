package csvtodf

import(
	"encoding/csv"
	"fmt"
	"os"
)

func CsvToDataframe(ipfile string, iptype struct){
	csvfile, err := os.Open(ipfile)
	Checkerror(err error)
	defer csvfile.Close()

	data := csv.NewReader(csvfile)
	rawdata, err := data.ReadAll()
	Checkerror(err error)

	var line iptype
	var lines []iptype

	for _, record := range raw_data {
		line.course_id = record[0]
		line.course_title = record[1]
		line.url = record[2]
		line.is_paid = record[3]
		line.price = record[4]
		line.num_subscribers = record[5]
		line.num_reviews = record[6]
		line.num_lectures = record[7]
		line.level = record[8]
		line.content_duration = record[9]
		line.published_timestamp = record[10]
		line.subject = record[11]

		lines = append(lines, line)
	}
	fmt.Println(lines)
}