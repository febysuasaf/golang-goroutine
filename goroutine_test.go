package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Display Number ", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 10; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
