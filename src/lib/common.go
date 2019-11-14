package lib

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// IsBlank - a function that returns a boolean indicating whether the string passed in is blank or not
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// FullTrim - trim spaces and remove new lines
func FullTrim(str string) string {
	return strings.Replace(strings.TrimSpace(str), "\n", "", -1)
}

// ReadLine - a function to read a line of text (\n delimited) and return the trimmed (no newline or space) input
func ReadLine(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return FullTrim(text)
}

// DownloadFile - a function used to download a file from a url and store it in the given file name
func DownloadFile(url string, filepath string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
