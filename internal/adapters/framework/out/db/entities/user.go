package entities

type User struct {
	ID        int32  `json:"id" bson:"id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
	Balance   int32  `json:"balance" bson:"balance"`
}
