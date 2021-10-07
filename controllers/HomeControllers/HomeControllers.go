package Homecontrollers

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Description string `json:"description"`
}
type messages []Message

//get individual user
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	Welcome := messages{
		Message{Description: "GolikePost👍🏻"},
		Message{Description: "Welcome to coolest tech stack ever"},
	}
	json.NewEncoder(w).Encode(Welcome)

}
