package model

type Comment struct {
	Id string `json:"id"`
	Content string `json:"content"`
	ManifestationId string `json:"manifestationId"`
}