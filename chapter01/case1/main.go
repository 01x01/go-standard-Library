// retrieving the golang version

package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Printf(`
		the application %s	starting,
		the go version is %s
	`,"example",runtime.Version())
}