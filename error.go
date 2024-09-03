package logimpl

import (
	"fmt"
	"time"
)

func Error(msg interface{}) {
	t := time.Now()
	fmt.Printf("[Error] %s:%s\n", t.Format("2006-01-02 15:04:05"), msg)
}
