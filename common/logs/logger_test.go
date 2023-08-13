package logs

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	Infos("hi")
	fmt.Println(Logger.Flags())
}
