package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)



type User struct {
	Id    bson.ObjectId 	`json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string        `json:"firstName" bson:"firstName"`  			//require
	LastName string        	`json:"lastName" bson:"lastname"`				//require
	FullName string        	`json:"fullName" bson:"fullName"`				
	Email string        	`json:"email" bson:"email"`						//require && unique
	NarrEmail string        `json:"narrEmail" bson:"narrEmail"`
	EmailVerified bool		`json:"emailVerified" bson:"emailVerified"`
	Password string        	`json:"password" bson:"password"`				//require
	UserRole string			`json:"userRole" bson:"userRole"`
	DOB string				`json:"DOB" bson:"DOB"`							//require
	Phone string        	`json:"phone" bson:"phone"`						//require
	Address string        	`json:"address" bson:"address"`					//require
	Intitution Institution	`json:"institution" bson:"institution"`	
	Avatar string			`json:"avatar" bson:"avatar"`
	Bvn string				`json:"bvn" bson:"bvn"`
	Bank Bank				`json:"bank" bson:"bank"`
	Blockchain Blockchain	`json:"blockchain" bson:"blockchain"`
	TotalUploads string 	`json:"totalUploads" bson:"totalUploads"`
	Specialisation []string	`json:"specialization" bson:"ownership"`
	Interest []string		`json:"intrest" bson:"intrest"`
	LastLoggedIn time.Time	`json:"lastLogin" bson:"lastLogin"`
	CreatedAt time.Time		`json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time		`json:"updatedAt" bson:"updatedAt"`
	Status string			`json:"status" bson:"status"`

}
type Bank struct{
	BankName string          	`json:"bankName" bson:"bankName"`
	AcccountName string 		`json:"accountName" bson:"accountName"`
	AcccountNumber string 		`json:"accountNumber" bson:"accountNumber"`
	DefaultAccount bool 		`json:"defaultAccount" bson:"defaultAccount"`
	AccountStatus string 		`json:"accountStatus" bson:"accountStatus"`
	CreatedAt time.Time			`json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time			`json:"updatedAt" bson:"updatedAt"`
}
type Institution struct{
	Name string					`json:"Name" bson:"Name"`
	Type string					`json:"type" bson:"type"`
	Acronym string				`json:"acronym" bson:"acronym"`
	Ownership string			`json:"ownership" bson:"ownership"`
	Url string 					`json:"ulr" bson:"ulr"`
	Year string					`json:"year" bson:"year"`
	Logo string					`json:"logo" bson:"logo"`
}
type Blockchain struct{
	PublicKey string			`json:"publicKey" bson:"publicKey"`
	Token string				`json:"token" bson:"token"`
	GasBalance string			`json:"gasBalance" bson:"gasBalance"`
}