package models

type Patch struct {
	Id     string                 `json:"id"`
	Fields map[string]interface{} `json:"fields"`
}
