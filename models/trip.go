package models

import "gopkg.in/mgo.v2/bson"

type Trip struct {
	Id	    bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	UserId	bson.ObjectId 	`json:"userid,omitempty"`
	Name    string   	    `json:"name"`
    Code    string   	    `json:"code"`
	Places	[]Place  	    `json:"places"`
}

