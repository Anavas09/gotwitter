package models

//Relation : Model of the relation between users (follow)
type Relation struct {
	UserID         string `bson:"userid" json:"userId"`
	UserRelationID string `bson:"userrelationid" json:"userRelationId"`
}
