package models

type CreateAlbumDto struct {
    Title  string  `json:"title" binding:"required"`
    Artist string  `json:"artist" binding:"required"`
    Price  float64 `json:"price" binding:"required,numeric,min=0"`
    Content string  `json:"content" binding:"required"`
}