package routes

import (
	"errors"
	"fmt"

	controller "github.com/ddld93/goauth/controllers"
	"github.com/ddld93/goauth/model"
	"github.com/ddld93/goauth/utilities"
)

type UserRoute struct {
	UserCtrl *controller.DB_Connect
}
var tokenMaker,_ = utilities.NewPasetoMaker("yhjkiuytgrtyujikjtedewertyhujkiu")

func (ur *UserRoute) Login(email, password string) (string,error){
	user, err :=	ur.UserCtrl.GetUser(email)
	if err != nil {
		return "", err
	}
	//comparing password 
	resp:= utilities.CheckPasswordHash(password ,user.Password)
	if !resp {
		return "", errors.New("invalid password")
	}
	token,err:= tokenMaker.CreateToken(user.Email,user.UserRole,100)
	
	return token,nil
}

func (ur *UserRoute) Register(user *model.User) (string,error){
	// hashing password using Bcrypt
	hash, err := utilities.HashPassword(user.Password)
	if err != nil {
		return "", errors.New(" error harshing password")
	}
	user.Password = hash
	resp,err := ur.UserCtrl.CreateUser(user)
	if err != nil {
		return "", err
	}
	return resp,nil
}
func (ur *UserRoute) Activate(token,email,status string) (string,error){
	claims,err := tokenMaker.VerifyToken(token)
	if err != nil {
		return "", err
	}
	err3 := claims.Valid()
	if err != nil {
	 return "", err3
	}
	if claims.AccoutType != "admin" {
		return "", errors.New("invalid account previllages")
	}
	user, err := ur.UserCtrl.GetUser(email)
	if err != nil {
		return "", err
	}
	user.UserRole = "active"
	err2 := ur.UserCtrl.UpdateUser(user)

	if err2 != nil {
		return "", err2
	}
	return fmt.Sprintf("user activated succesifully"), nil

}
func (ur *UserRoute) ResetPassword(email string) (error){
	user, err := ur.UserCtrl.GetUser(email)
	if err != nil {
		return err
	}
 // send instructions to user.Email email 
 // check for errors
fmt.Println("instruction email sent to", user.Email)
 return nil

}
func (ur *UserRoute) ChangePassword(token string) error{
	claims,err := tokenMaker.VerifyToken(token)
	if err != nil {
		return  err
	}
	// validating tokens
   err2 := claims.Valid()
   if err != nil {
	return  err2
   }
	
	user, err := ur.UserCtrl.GetUser(claims.Username)
	if err != nil {
		return  err
	}
	hash, err := utilities.HashPassword(user.Password)
	if err != nil {
		return  err
	}
	user.Password = hash
	err3 := ur.UserCtrl.UpdateUser(user)
	if err2 != nil {
		return  err3
	}
	return  nil

}


