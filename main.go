/*package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("testing.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println(string(data))
}
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// readWordFromStandardInput()
	// readLineFromStandardInput()
	// readFromFile()
}

func readWordFromStandardInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	fmt.Println(scanner.Text())
}

func readLineFromStandardInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
}

func readFromFile() {
	file, err := os.Open("testing.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}
}
