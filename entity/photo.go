package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	// ImageCollectionName represents the Image entity collection name
	ImageCollectionName = "images"
)

// Image entity
type Image struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Label string             `json:"label,omitempty" bson:"label,omitempty"`
	URL   string             `json:"url,omitempty" bson:"url,omitempty"`
}
