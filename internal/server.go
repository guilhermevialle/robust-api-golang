package app

func StartServer() {
	app := NewApp()
	app.Run(":8080")
}
