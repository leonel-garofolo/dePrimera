package weekend

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	fmt.Println("Liga con 8 equipos:")
	weekends := getWeekendFromDate(time.Now())
	for i := 0; i < len(weekends); i++ {
		fmt.Println(weekends[i])
	}
}
