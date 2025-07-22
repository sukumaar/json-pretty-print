package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return
		}

		var obj interface{}
		err := json.Unmarshal([]byte(line), &obj)
		if err != nil {
			fmt.Println("❌ Invalid JSON:", err)
			return
		}

		prettyJSON, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			fmt.Println("❌ Error formatting JSON:", err)
			return
		}

		fmt.Println(string(prettyJSON))

	}

}
