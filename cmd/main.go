package main

import "encryption-algorithm/internal/app"

func main() {
	r := app.Init()
	r.Run(":8080")
}
