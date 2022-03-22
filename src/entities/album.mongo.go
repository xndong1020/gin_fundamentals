package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type AlbumMongoDB struct {
	ID     primitive.ObjectID    `bson:"_id"`
	Name  string                 `bson:"name,omitempty"`
	Content string               `bson:"content,omitempty"`
	AlbumId uint  				 `bson:"albumId,omitempty"`
}