package model

// swagger:parameters welcome
type WelcomeInput struct {
	// Kullanıcının adı
	// required: true
	FirstName string `json:"firstName"`

	// Kullanıcının soyadı
	LastName string `json:"lastName"`

	// HTTP header içinde gönderilen bir bilgi
	// in:header
	SomeHeader string `json:"some-header"`

	// HTTP body içinde gönderilen bir bilgi
	// in:body
	SomeBody string `json:"some-body"`
}

// Hoşgeldin servisinden dönen cevap.
// swagger:response
type WelcomeResponse struct {
	// Hoşgeldin mesajı
	// in: body
	Message string `json:"message"`
}
