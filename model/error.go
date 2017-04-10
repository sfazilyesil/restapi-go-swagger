package model

// Servisten dönen hata modeli.
//
// swagger:model
type ErrorResponse struct {
	// Hata kodu
	Code string `json:"code"`

	// Hata mesajı
	Message string `json:"message"`
}

// Hatalı istek. Servise gönderilen parametreler hatalı.
// swagger:response
type ErrorResponse400 struct {
	//in:body
	Body ErrorResponse
}

// Kayıt bulunamadı. Verilen kriterlere uygun kayıt bulunamadı.
// swagger:response
type ErrorResponse404 struct {
	//in:body
	Body ErrorResponse
}

// Alt servis çağrısından hata alındı.
// swagger:response
type ErrorResponse502 struct {
	//in:body
	Body ErrorResponse
}

// Diğer hatalar.
// swagger:response
type GenericErrorResponse struct {
	//in:body
	Body ErrorResponse
}
