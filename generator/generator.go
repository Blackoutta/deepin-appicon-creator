package generator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GenerateDesktopIcon() {
	params := []string{
		"version",
		"name",
		"comment",
		"exec",
		"icon",
		"terminal",
		"type",
		"catetories",
	}
	de := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	for _, v := range params {
		r, err := getInput(v, scanner)
		if err != nil {
			log.Fatalf("error while getting user input: %v\n", err)
		}
		de[v] = r
	}
	template := "[Desktop Entry]\nVersion=%s\nName=%s\nComment=%s\nExec=%s\nIcon=%s\nTerminal=%s\nType=%s\nCategories=%s\n"
	content := fmt.Sprintf(template,
		de["version"],
		de["name"],
		de["comment"],
		de["exec"],
		de["icon"],
		de["terminal"],
		de["type"],
		de["catetories"])
	filePath := fmt.Sprintf("/usr/share/applications/%s.desktop", de["name"])
	if err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm); err != nil {
		log.Fatalf("error while writing file: %v", err)
	}
}

func getInput(s string, scanner *bufio.Scanner) (string, error) {
	fmt.Printf(s + ": ")
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Printf("error while scanning: %v\n", scanner.Err().Error())
		return "", scanner.Err()
	}
	return scanner.Text(), nil
}
