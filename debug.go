package logimpl

import (
	"fmt"
	"time"
)

func Debug(msg interface{}) {
	t := time.Now()
	fmt.Printf("[Debug] %s:%s\n", t.Format("2006-01-02 15:04:05"), msg)
}
