package models

import "encoding/json"

// Confession a model confession
type Confession struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ToJson transform struct to json
func (c *Confession) ToJson() ([]byte, error) {
	return json.Marshal(c)
}
