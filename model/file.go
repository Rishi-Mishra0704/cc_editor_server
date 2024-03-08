package model

type File struct {
	Content       string `json:"content"`
	FileExtension string `json:"fileExtension"`
	Room          string `json:"room"`
}
