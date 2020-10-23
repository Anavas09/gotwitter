package models

//Tweet : Struct for tweet that came from the r.Body
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
