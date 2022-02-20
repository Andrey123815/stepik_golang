package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	t, _ := time.Parse(time.RFC3339, scanner.Text())
	fmt.Println(string(t.Format(time.UnixDate)))
}
