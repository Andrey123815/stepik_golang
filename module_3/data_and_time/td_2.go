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

	//t, _ := time.Parse(time.RFC3339, scanner.Text())

	t, err := time.Parse("2006-01-02 15:04:05", scanner.Text())

	if err != nil {
		panic(err)
	}

	if t.Hour() > 13 {
		new_dt := t.AddDate(0, 0, 1)
		fmt.Print(new_dt.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Print(scanner.Text())
	}
}
