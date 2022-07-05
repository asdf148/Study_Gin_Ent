// main
package main

func main() {
	r := initializeRoutes()
	r.Run(":8080")
}
