package main

import (
	"fmt"

	controller "github.com/ddld93/goauth/controllers"
	"github.com/ddld93/goauth/model"
	routes "github.com/ddld93/goauth/routes"
)
func main()  {
	userCtrl := controller.NewUserCtrl("localhost", 4321)
	route := routes.UserRoute{UserCtrl: userCtrl}
 	
	
	
}