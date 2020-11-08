package models

import "go.mongodb.org/mongo-driver/bson/primitive"



type Article struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id   string             `json:"id,omitempty" bson:"id,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	SubTitle  string             `json:"Stitle" bson:"Stitle,omitempty"`
	Content  string             `json:"content" bson:"content,omitempty"`

}



var Articles []Article
