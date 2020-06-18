package main

import "fmt"

// Creating a program that will help process data for comic book store
func main() {
	// defining all of our string data
	var publisher, writer, artist, title string
	// defining all of our numeric data
	var year, pageNumber int
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	// Creating a print statement for our values.
	fmt.Println(title, "written by", writer, "drawn by", artist, "publisher: ", publisher, "year:", year, "# of pages:", pageNumber, "grade:", grade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	year = 2013
	pageNumber = 160
	grade = 9.0

	// Creating a print statement for our values.
	fmt.Println(title, "written by", writer, "drawn by", artist, "publisher: ", publisher, "year:", year, "# of pages:", pageNumber, "grade:", grade)

}
