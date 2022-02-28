package models

type Album struct {
    Id     int  `json:"id" binding:"required,numeric,min=1"`
    Title  string  `json:"title" binding:"required"`
    Artist string  `json:"artist" binding:"required"`
    Price  float64 `json:"price" binding:"required,numeric,min=0"`
}