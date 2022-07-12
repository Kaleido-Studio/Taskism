package main

import (
	"taskism/controllers"
)

func main() {
	controllers.InitDB()
	r := controllers.GinEngine()
	r.Run("localhost:3001")
}
