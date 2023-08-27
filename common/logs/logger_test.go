package logs

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	Info("hi")
	fmt.Println(Logger.Flags())
}
