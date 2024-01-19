package logprint

import (
	"fmt"
	"time"
)

func Error(msg interface{}) {
	fmt.Printf("[error] [%s] %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), msg)
}
