package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	EmployeeID string `json:"employee_id,omitempty" bson:"employee_id"`
	Name       string `json:"name,omitempty" bson:"name"`
	Department string `json:"department,omitempty" bson:"department"`
}

type Blog struct {
	// ID uint `json:"id" gorm:"primaryKey"`
	ID    string `json:"blog_id" bson:"blog_id, omitempty"`
	Title string `json:"title" bson:"title, omitempty"`
	Post  string `json:"post" bson:"post, omitempty"`
	Image string `json:"image" bson:"image, omitempty"`
}

type Song struct {
	ID string `json:"song_id" bson:"song_id,omitempty"` // MongoDB 自動生成的唯一 ID

	// Get from scraper
	Artist      string             `json:"artist" bson:"artist"` // Singer/Artist
	Album       string             `json:"album" bson:"album"`   // Album Name
	ReleaseDate string             `json:"release_date" bson:"release_date"`
	CoverImage  primitive.ObjectID `json:"cover_image" bson:"cover_image,omitempty"`
	SongURL     string             `json:"song_url" bson:"song_url, omitempty"`
	Lyricst     string             `json:"lyricist" bson:"lyricist"`
	Composer    string             `json:"composer"`
	CoverURL    string             `json:"cover_url" bson:"cover_url, omitempty"`
	Title       string             `json:"title" bson:"title"`

	// User Input
	AddedDate string `json:"added_date" bson:"added_date, omitempty"`
}
