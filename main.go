package main

import (
	"time"
)
import . "demo/pkg"

func main() {
	// go Doit_const()
	// go Doit_variable()
	// go Doit_strconv()
	// go Doit_regexp()
	// go Doit_runtime()

	go Channel()
	// go ChannelCheckClose()
	// go ChannelApplicationCrond()
	// go ChannelApplicationDecouple()
	// go ChannelApplicationControlConCurrency()

	//保证goroutine走完
	time.Sleep(10 * time.Second)
}


