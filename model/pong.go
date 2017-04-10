package model

// Ping isteğine dönülen pong cevap modeli.
//
// swagger:model
type Pong struct {
	// Pong mesajı
	Message string `json:"message"`
}

// Ping isteğine gönderilen pong cevabı
// swagger:response
type PongResponse struct {
	// in: body
	Pong Pong `json:"pong"`
}
