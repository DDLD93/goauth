package utilities

import (
	"errors"

	"github.com/ddld93/goauth/model"
)

func UserModelValidate(user *model.User)  (*model.User, error){
	// checking required fields 
	if user.FirstName == "" {
		return user, errors.New("first name field cannot be empty")
	}
	if user.LastName == "" {
		return user, errors.New("last name field cannot be empty")
	}
	if user.Email == "" {
		return user, errors.New("email field cannot be empty")
	}
	if user.Password == "" {
		return user, errors.New("password field cannot be empty")
	}
	if user.DOB == "" {
		return user, errors.New("date of birth field cannot be empty")
	}
	if user.Phone == "" {
		return user, errors.New("phone field cannot be empty")
	}
	if user.Address == "" {
		return user, errors.New("address field cannot be empty")
	}
	if user.UserRole == "" {
		return user, errors.New("user role must be provided")
	}
	// if user.Intitution.Name == "" {
	// 	return user, errors.New("institution name field cannot be empty")
	// }
	// if user.Intitution.Type == "" {
	// 	return user, errors.New("institution type field cannot be empty")
	// }
	// assigning default values
	user.Status = "pending"
	user.EmailVerified = false
	user.Avatar="/profile/nopic.jpg"
	return user, nil
}