package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Type     string             `json:"type,omitempty" validate:"required"`
	Label    string             `json:"label,omitempty" validate:"required"`
	SongUrl  string             `json:"song_url,omitempty" validate:"required"`
	Cover    string             `json:"cover,omitempty" validate:"required"`
	Movie    string             `json:"movie,omitempty" validate:"required"`
	Mood     string             `json:"mood,omitempty" validate:"required"`
	Language string             `json:"language,omitempty" validate:"required"`
}
