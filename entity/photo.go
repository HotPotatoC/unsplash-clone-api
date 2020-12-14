package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	// PhotoCollectionName represents the photo entity collection name
	PhotoCollectionName = "photos"
)

// Photo entity
type Photo struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Label string             `json:"label,omitempty" bson:"label,omitempty"`
	URL   string             `json:"url,omitempty" bson:"url,omitempty"`
}
