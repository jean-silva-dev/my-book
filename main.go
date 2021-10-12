package main

import (
	"book/cli"
	"book/database/migration"
)

func main() {
	migration.Init()
	cli.Init()
}
