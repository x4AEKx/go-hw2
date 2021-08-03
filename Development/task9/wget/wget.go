package wget

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Wget takes fileName, urlPath returns error
func Wget(fileName string, urlPath string) error {
	fmt.Printf("Connecting to %s...\n", urlPath)

	response, err := http.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Printf("HTTP request send, awaiting response... %s\n", response.Status)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	size, err := io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Length: %v\n", size)
	fmt.Printf("Saving to: %s\n", fileName)

	return nil
}
