package main

import (
	_ "github.com/mattn/go-sqlite3"

	"github.com/smxlong/kit/service"

	"github.com/smxlong/mud"
)

func main() {
	service.Main("mud", "run the mud", mud.New())
}
