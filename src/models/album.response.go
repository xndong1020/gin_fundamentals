package models

type AlbumResponse struct {
    Id     int  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
	Content string `json:"content"`
}