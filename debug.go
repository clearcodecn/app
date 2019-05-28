package app

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

var (
	IsDebug bool
	one sync.Once
)

func Print(message string)  {
	output(message)
}

func Printf(format string,args ...interface{})  {
	output(fmt.Sprintf(format,args))
}

func output(message string)  {
	if !IsDebug {
		return
	}

	one.Do(func() {
		log.SetFlags(log.Ltime)
	})

	_, file,line,ok := runtime.Caller(2)
	if !ok {
		file = "?"
	}
	log.Println(fmt.Sprintf("INFO\t%s:%d\t%s",file,line,message))
}