package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sfazilyesil/restapi-go-swagger/model"
)

// NewServer yeni bir sunucu çalıştırır ve tanımlı API
// metotlarını dinlemeye başlar.
func NewServer() *gin.Engine {
	r := gin.Default()

	//---------------------------------
	//---------- Basic API ------------
	//---------------------------------

	// swagger:route GET /ping Genel send-ping
	// Ping isteği gönder
	//
	// Uygulamanın ayakta olup olmadığını basitçe anlamaya
	// yarayan servis. Geriye pong cevabını gönderir.
	//
	// responses:
	// 200: PongResponse
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, model.Pong{Message: "pong..."})
	})

	// swagger:route GET /welcome Genel welcome
	// Hoşgeldin mesajı al
	//
	// Gönderilen parametleri kullanarak hoş geldin mesajı üretip
	// gönderir.
	//
	// responses:																																																																																 responses:
	//   200: WelcomeResponse
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Misafir")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Merhaba %s %s", firstname, lastname)
	})

	//---------------------------------
	//-------- Products API -----------
	//---------------------------------

	// swagger:route POST /products Ürünler create-product
	// Yeni bir ürün ekle
	//
	// Servis ürünlere yeni bir ürün ekler ve eklenen ürünü geriye
	// gönderir.
	//
	// responses:
	//   201: body:Product
	//   400: ErrorResponse400
	//   502: ErrorResponse502
	//   default: GenericErrorResponse
	r.POST("/products", func(ctx *gin.Context) {
		// TODO ürün ekle
	})

	// swagger:route GET /products/{id} Ürünler get-product-by-id
	// Referans ile ürün getir
	//
	// Referansı verilen ürünü iletir.
	//
	// responses:
	//   200: body:Product
	//   400: ErrorResponse400
	//   404: ErrorResponse404
	//   default: GenericErrorResponse
	r.GET("/products/:id", func(ctx *gin.Context) {
		// TODO ürünü getir
	})

	//---------------------------------
	//---------- Users API ------------
	//---------------------------------

	// swagger:route POST /users Kulanıcılar create-user
	// Yeni bir kullanıcı ekle
	//
	// Servis kullanıcılara yeni bir kullanıcı ekler ve eklenen kullanıcıyı
	// geriye gönderir.
	//
	// responses:
	//   201: body:User
	//   400: ErrorResponse400
	//   default: GenericErrorResponse
	r.POST("/users", func(ctx *gin.Context) {
		// TODO kullanıcı ekle
	})

	// swagger:operation GET /users/{id} Kulanıcılar get-user-by-id
	// Referans ile kullanıcı getir
	//
	// Referansı verilen kullanıcıyı iletir.
	//
	// ---
	// parameters:
	// - name: id
	//   in: path
	//   description: Kullanıcı referansı
	//   type: integer
	//   format: int64
	//   required: true
	//
	// responses:
	//   '200':
	//     description: OK
	//     schema:
	//       $ref: "#/definitions/User"
	//   '404':
	//     description: NotFound
	//     schema:
	//       $ref: "#/definitions/ErrorResponse"
	r.GET("/users/:id", func(ctx *gin.Context) {
		// TODO kullanıcıyı getir
	})

	// swagger:route PUT /users/{id} Kulanıcılar update-user
	// Kullanıcıyı güncelle
	//
	// Referansı verilen kullanıcıyı günceller.
	//
	// responses:
	//   200: body:User
	//   400: ErrorResponse400
	//   404: ErrorResponse404
	//   default: GenericErrorResponse
	r.PUT("/users/:id", func(ctx *gin.Context) {
		// TODO kullanıcıyı güncelle
	})

	return r
}
