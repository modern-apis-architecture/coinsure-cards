package main

const defaultPort = "9999"

func main() {
	app, _ := buildAppContainer()
	app.Run()
}
