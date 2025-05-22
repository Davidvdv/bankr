package main

import "time"
import "fmt"

func main() {
	start := time.Now()
    defer func() {
        elapsed := time.Since(start)
        ms := elapsed.Milliseconds()
        fmt.Printf("Execution time: %dms\n", ms)    
	}()
	
	app := ApplicationFactory()
	app.Go()
}