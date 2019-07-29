package main


import ( 
 		"github.com/fabianolaudutra/goToken/config"
		"github.com/fabianolaudutra/goToken/app"
)

 func main() {
		  
	config := config.GetConfig()		
 	app1 := app.App{}
 	app1.Initialize(config)
	app1.Run(":3000")
	
 }
