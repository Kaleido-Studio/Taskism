package main

import (
	"taskism/controllers"
)

func main() {
	r := controllers.GinEngine()
	r.Run("localhost:3001")
}
