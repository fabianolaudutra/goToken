package main
/*
import "fmt"

func main() {
	fmt.Printf("AAAAQQQUUUU\n")
}*/

import ( "fmt"
 
		"goToken/config"
		"reflect"
		"github.com/fabianolaudutra/goToken/app"
)

 func main() {
	fmt.Printf("SAdasdasIU\n")
	  
	config := config.GetConfig()
	fmt.Println(reflect.TypeOf(config).String())
	
 	app1 := app.App{}
 	app1.Initialize(config)
	app1.Run(":3000")
	
 }
