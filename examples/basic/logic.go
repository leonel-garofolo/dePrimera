package basic

import (
	"fmt"
	"time"
)

func currentTime() {
	t := time.Now()
	fmt.Println("origin : ", t.String())
	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(formatted)
}
