package main

import (
	"ginarticle/mysql"
	"ginarticle/router"
	"time"
)

func main() {
	mysql.Dbinit()
	go router.Router()

	for {
		time.Sleep(0 * time.Microsecond)
	}
}
