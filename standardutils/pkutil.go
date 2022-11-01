// standardutils module is intended to have common utilities, for me and anyone else.
package standardutils

import (
	"bufio"
	"go/scanner"
	"os"
)

// ErrorMessage prints a custom error message passed by you and an error message of type error or string via argument by parameter "anyError", if you want you can also pass a custom error code as a third argument.
func ErrorMessage(message string, anyError any, errorCode...int) {

	println(message)
	if str, ok := anyError.(string); ok {
		println("Error message:", str)
	} else if err, ok := anyError.(error); ok {
		println("Error message:", err.Error())
	}

	if len(errorCode) == 0 {
		os.Exit(1)
	} else {
		os.Exit(errorCode[0])
	}

}

// StringReader creates a NewReader and reads a string from STDIN and returns.
func StringReader() string {

	read := bufio.NewReader(os.Stdin)
	input, err := read.ReadString('\n')

	if err != nil {
		ErrorMessage("Error on input data!", err)
	}

	return input

}

// FileLineReader reads a file line by line and returns a slice of strings with all of them.
func FileLineReader(filePath string) ([]string, string) {

	var lines []string

	file, err := os.Open(filePath)
	if err != nil {
		ErrorMessage("Error on file opening!", err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}
	defer file.Close()

	return lines, scanner.Error{}.Msg

}
