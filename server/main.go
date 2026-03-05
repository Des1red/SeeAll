package main

import (
	"SeeAll/cmd"

	_ "SeeAll/internal/sources/sourcepages"
)

func main() {
	cmd.Execute()
}
