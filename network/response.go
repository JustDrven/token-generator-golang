package network

type Token struct {
	Token string `json:"token"`
}

type Error struct {
	Error bool `json:"error"`
	Message string `json:"message"`
}

