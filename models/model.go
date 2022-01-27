package models

// model of an article
type Article struct {
	ID                 int    `json:"id"`
	Title              string    `json:"title"`
	SubTitle           string    `json:"subtitle"`
	Content            string    `json:"content"`
	Creation_Timestamp  string `json:"creation"`
}

