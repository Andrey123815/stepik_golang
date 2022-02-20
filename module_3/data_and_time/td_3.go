package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	s := strings.Split(scanner.Text(), ",")
	dt1, _ := time.Parse("02.01.2006 15:04:05", s[0])
	dt2, _ := time.Parse("02.01.2006 15:04:05", s[1])
	t1 := time.Since(dt1)
	t2 := time.Since(dt2)
	if t1 > t2 {
		fmt.Println(t1.Round(time.Second) - t2.Round(time.Second))
		return
	}
	fmt.Println(t2.Round(time.Second) - t1.Round(time.Second))
}
