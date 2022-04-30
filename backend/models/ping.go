package models

type Ping struct {
	Command string `json:"command" binding:"required"`
}

// Ping logic here if there were any lol
// do it as methods
