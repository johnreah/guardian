package guardian

import (
	"fmt"
	"time"
	"testing"
)

func TestListTables(t *testing.T) {
	fmt.Println("Starting...")
	startTime := time.Now()

	ListTables()

	fmt.Printf("\nFinished in %dms\n", time.Now().Sub(startTime)/1000000)
}
