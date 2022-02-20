package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// вам это понадобится
const now = 1589570165

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var dur_str string
	dur_str = scanner.Text()
	dur_str = strings.Replace(dur_str, " мин. ", "m", 1)
	dur_str = strings.Replace(dur_str, " сек.", "s", 1)

	//fmt.Print(dur_str)

	dur, err := time.ParseDuration(dur_str)
	if err != nil {
		panic(err)
	}

	unixTime := time.Unix(now, 0).UTC()
	//fmt.Print(unixTime.Format(time.UnixDate))

	timein := unixTime.Add(dur)
	fmt.Print(timein.UTC().Format(time.UnixDate))
	// - replace-ом заменил " мин. " и " сек." на формат удобный для ParseDuration
	// - к Unix времени добавил (time.Add) полученный duration, ну и Format(time.UnixDate) для вывода
}
