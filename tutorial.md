# Golang Otomatik API Dökümantasyonu Üretimi

Hikayesi başka zamana. Şimdilik lafı uzatmadan kestirmeden gidelim.

## Problem Tanımı
1- Go ile yazılmış bir REST API uygulamasının Swagger 2.0 formatına uygun API dökümanını 
Go kodundan nasıl otomatik olarak üretebiliriz?

2- Üretilen dökümanı nasıl doğrulayabiliriz?

3- Bu işleri günlük çalışma akışına (ilgili build scriptlerine) nasıl dahil edebiliriz?

4- Üretilen dökümanı nasıl görüntüleyebiliriz?

## Çözüm
Görüp inceleyebildiğim kadarıyla go-swagger, Swagger 2.0 için ekosistemdeki en iyi/bütünlüklü 
çözüm sunan golang kütüphanesi . Koddan döküman üretimi, dökümandan istemci ve sunucu oluşturulması, 
döküman doğrulaması vb. şeyleri yapma yeteneğine sahip.

Bizim odaklandığımız problemler ise koddan otomatik olarak döküman üretimi ve 
üretilen dökümanın doğrulanması ile sınırlı. Go-swagger kütüphanesi bu ihtiyaçları
tamamıyla karşılayabiliyor. Bazı noktalarda başka kütüphanelerin sunduğu bazı kısa 
yolları sunmasa da şu an için en bütünlüklü çözüm sunan kütüphane. Bazı şeyleri 
yapış biçimini beğenmesek de bir yol öneriyor.
 

## Giriş
Go dilinde Java'daki anotasyonlara benzer bir yapı yok. Go-swagger döküman üretimini
yorum satırlarına konulan özel `swagger:...` benzeri direktifleri algılayıp çözümleyerek,
go kod ve yorum satırlarından API döküman üretiyor.


## Dökümanın kodda ifade edilmesi

### Temel go-swagger direktifleri ve açıklamaları:

- `swagger:meta`: Dökümandaki versiyon, şema, içerik, lisans, genel açıklama 
benzeri bilgilerin ifade etmeyi sağlayan direktif.

- `swagger:route`: Api metotlarını (endpoint) ifade eden direktif. Girdi, çıktı, 
içerik tipi vb. bilgiler belirtmeyi sağlıyor.

- `swagger:params`: Api metodunun beklediği parmetreleri ifade etmeyi sağlıyor. 
Parametrelerin nerede olduğu header/body/path/query olarak belirtilebiliyor. 

- `swagger:operation`: `swagger:route` alternatifi. Direk yalın Swagger 2.0 yaml 
formatında yazma olanağı sağlıyor.

- `swagger:response`: İsteklere verilen cevabı ifade etmeyi sağlıyor.

- `swagger:model`: JSON veri modellerini ifade etmeyi sağlıyor. Go struct'ları
bu direktif ile işaretlendiğinde, JSON veri modeli olarak API dökümanındaki
"definitions" bölümüne yansıtılıyor. Tabi bir yerlerde referans verilmek şartıyla.
 
- `swagger:allOf`: İçiçe veri modellerini ifade etmede kullanılıyor. İçtekinde 
kullanıldığında ayrı bir model olduğu belli oluyor, ancak kullanılmazsa içtekinin 
alanları dıştakinin içine serpiştiriliyor, ayrı bir modelden söz edilmiyor bu 
durumda.

- `swagger:strfmt`: Veri formatlarılarını ifade etmede kullanılıyor. date, date-time,
email, uuid4 vs.

Örnek uygulama: github.com/sfazilyesil/restapi-go-swagger
Aşağıda alıntılanan kod parçaları örnek uygulamaya ait.


### API Dökümanı Giriş Noktası
*restapi-go-swagger/main.go* giriş noktasını oluşturuyor.
```go
//go:generate swagger generate spec
func main() {
	s := server.NewServer()
	s.Run()
}
```
 
### Meta Bilgilerin İfade Edilmesi
*restapi-go-swagger/doc.go* içeriği:
```go
// RESTAPI-GO-SWAGGER uygulaması API dökümanı.
//
// Bu uygulamanın amacı kendi kodunu kullanarak API dökümanını da
// üreten örnek bir go uygulaması ortaya koymaktır.
//
// Uygulama go kodu ve yorum satırlarında kullanılan özel direktifleri
// kullanarak swagger 2.0 uyumlu bir API spesikasyon dökümanı üretmektedir.
//
// Schemes: http, https
// Host: localhost
// BasePath: /v2
// Version: 0.0.1
// License: MIT http://opensource.org/licenses/MIT
// Contact: sfy<sfazilyesil@github.com>
//
// Consumes:
//   - application/json
//   - application/xml
//
// Produces:
//   - application/json
//   - application/xml
//
// swagger:meta
package main
```

### API Metotlarının İfade Edilmesi

Farklı durumlar ve aynı şeyi farklı biçimlerde yapmaya birkaç örnek 
(*restapi-go-swagger/server/server.go* dosyasından):

```go
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
```

#### Modeller, İstekler ve Cevaplar
İlgili tanımlar *restapi-go-swagger/models* modülünde bulunuyor.

Hata modeli:
```go
// Servisten dönen hata modeli.
//
// swagger:model
type ErrorResponse struct {
	// Hata kodu
	Code string `json:"code"`

	// Hata mesajı
	Message string `json:"message"`
}
```

Sık dönülen hatalar için ortak cevaplar (*restapi-go-swagger/models* modülünden):
```go
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
```

*GET /welcome* metodunun beklediği parametreler. Bir kısmı istek başlığında, bir kısmı sorgu parametresi olarak,
bir kısmı istek body içinde gelmesi bekleniyor.
```go
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
```

*GET /welcome* metodu geriye düz string cevap dönüyor:
```go
// Hoşgeldin servisinden dönen cevap.
// swagger:response
type WelcomeResponse struct {
	// Hoşgeldin mesajı
	// in: body
	Message string `json:"message"`
}
```

Kullanıcı modelinin tanımı:
```go
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
```

*POST /users* metodunun parametre tanımı:
```go
// swagger:parameters create-user
type CreateUserInput struct {
	// Kullanıcıya ait bilgiler
	// in:body
	User User
}
```

*GET /produts/{id}* metodunun parametre/girdi tanımı:
```go
// swagger:parameters get-product-by-id
type ProductByIdInput struct {
	// Ürün referansı
	// in:path
	Id int64 `json:"id"`
}
```

## Dökümanın Otomatik Olarak Üretimi ve Doğrulanması

Döküman üretiliyor, doğrulanıyor ve hata oluştuğunda akış kesiliyor.
Dosyayı günlük kullandığımız *build* akışına entegre ettiğimizde başarılı ve doğru
biçimde güncel dökümantasyonun üretildiğinden emin olabileceğiz. Tabi insan 
faktörü her zaman mevcut. Yazılan kodlarda ilgili API dökümantasyon direktiflerinin
eklenmiş olduğunu varsayıyoruz şimdilik. Bir sonraki adım olarak yazılan kodların
üretilen API dökümanına uyup uymadığını otomatik olarak test eden bir kodu 
*build script*lerine entegre etmek olabilir.

*restapi-go-swagger/apidoc.sh* dosyası. Yorumlar kaynak kodda mevcut.
```bash
# go-swagger swagger komutunu derle.
if ! [ -x swagger ]; then
    echo "go-swagger swagger komutu kaynak koddan derleniyor..."
    go build ./vendor/github.com/go-swagger/go-swagger/cmd/swagger/swagger.go
fi

# API dökümanı oluştur. Hata varsa akışı kes.
echo "API dökümanı oluşturuluyor..."
OUTPUT=$(./swagger generate spec -o swagger.json 2>&1)
if [ "$OUTPUT" ]; then
   echo $OUTPUT
   echo "API dökümanı oluşturulamadı, işleme devam edilemiyor :("
   exit 1
fi


# API dökümanını doğrula. Hata varsa akışı kes.
echo "Oluşturulan API dökümanı swagger 2.0 versiyonuna göre doğrulanıyor..."
OUTPUT=$(./swagger validate swagger.json 2>&1)
if echo $OUTPUT |grep -q 'errors :'; then
   echo $OUTPUT
   echo "API dökümanı hatalı, işleme devam edilemiyor :("
   exit 1
fi

# Hata yok hayata devam et.
echo "API dökümanı geçerli :)"
```


## Redoc entegrasyonu
swagger-ui haricinde üretilen dökümanı göze hoş gelecek biçimde görüntülemenin, 
sunmanın yollarından biri de **ReDoc**. 

Tıpkı *Unix pipes* gibi... Eğer makine tarafından okunabilir bir şey varsa elimizde,
onu farklı işlemlere girdi olarak verebiliriz. 

**Kod -> API dökümanı -> Html dökümanı**  

ReDoc swagger 2.0 formatında yazılmış bir API dökümanını güzel bir Html dökümanına çeviriyor.

![alt text](https://github.com/sfazilyesil/restapi-go-swagger/blob/master/redoc.png)


## İleriye Dönük Fanteziler

Hayal kurmak bedava. :)

API dökümanından:
- otomatik servis spesifikasyon testinin oluşturulması ve *build* akışına entegrasyonu.
- özelleştirilmiş şablon oluşturup kullanarak istemci modülünün otomatik olarak oluşturulması.
- çalışma anında isteklerin girdi/parametre denetimlerinin döküman üzerinden yapılması.
- ...


## Kaynaklar
- http://swagger.io/
- https://goswagger.io/
- https://github.com/Rebilly/ReDoc
  
  