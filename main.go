package main

import (
	"fmt"

	controller "github.com/ddld93/goauth/controllers"
	"github.com/ddld93/goauth/model"

	// "github.com/ddld93/goauth/model"
	routes "github.com/ddld93/goauth/routes"
)
func main()  {
	userCtrl := controller.NewUserCtrl("localhost", 4333)
	route := routes.UserRoute{UserCtrl: userCtrl}
    newUser:= model.User{
		FirstName: "umar",
		LastName: "jere",
		Email: "umar.jere@ymail.com",
		Password: "secretecode",
		Phone:"07055793353",
		Address: "secrete location",
		DOB: "11/05/1993",
		Intitution: model.Institution{
			Name: "ABU Zaria",
			Type: "University",
		},
		
	}




	fmt.Println(route.Register(&newUser, "admin"))
	// fmt.Println(route.Login("umar.jere@yahoo.com","secretecode" ))
	// fmt.Println(route.Activate("v2.local.TEAWnzCavuwhosROHyd6sYhaVKH2UpLuKEQtXFmcfxnKAN6v589MRMQQGUX_JEwmqsUa1VMh6nCEZ9K35cjTnyo1i3VBhtZmFwgp5xzNzRslQenmElJQZr70h5_srHjL_vvDi3tX_TfT-i3WLGgHhJhZBMLlZu8fTdey9OvPq2EyzClD4hAMmmNbXTU9y-l43UETDTthWTzXPcpW6zz2-OZE7agreeeROUNxnpL9SMR5O6mIT82FsrzG.bnVsbA","umar.jere@yahoo.com","active"))
 	
	
	
}