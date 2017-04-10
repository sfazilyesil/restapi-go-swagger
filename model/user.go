package model

import "time"

// swagger:model
type User struct {
	// Kullanıcının  referansı
	ID int64 `json:"id"`

	// Kullanıcının ismi
	// min length: 3
	Name string `json:"name,required"`

	// Kullanıcının doğum günü
	//
	// swagger:strfmt date-time
	Birthday time.Time `json:"birthday"`

	// Kullanıcının cinsiyeti
	// enum: E,K
	// default: K
	Gender string `json:"gender"` // enum

	// swagger:allOf
	Address *Address `json:"address"`

	// Kullanıcının eposta adresi
	// unique: true
	// swagger:strfmt email
	Email string `json:"email"`
}

// Adres bilgisi
type Address struct {
	// Ülke
	Country string `json:"country"`

	// Şehir
	City string `json:"city"`

	// Sokak
	Street string `json:"street"`
}

// swagger:parameters create-user
type CreateUserInput struct {
	// Kullanıcıya ait bilgiler
	// in:body
	User User
}

// swagger:parameters update-user
type UpdateUserInput struct {
	// Kullanıcı referansı
	// in:path
	Id int64 `json:"id"`

	// Kullanıcıya ait bilgiler
	// in:body
	User User
}
