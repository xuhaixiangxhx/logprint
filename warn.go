package logprint

import (
	"fmt"
	"time"
)

func Warn(msg interface{}) {
	fmt.Printf("[warn] [%s] %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), msg)
}
