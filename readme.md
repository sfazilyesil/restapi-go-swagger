# restapi-go-swagger

**Golang** ile yazılmış web servislerin **swagger 2.0** formatına uygun
**API dökümanı**nı otomatik olarak nasıl üretilebileceğimizi göstermek amacıyla 
hazırlanmış bir **örnek uygulama**.

Örnek uygulamanın genel anlatımı için: [bkz](https://github.com/sfazilyesil/restapi-go-swagger/blob/master/tutorial.md)
  
  
#### Hazır üretilmiş API dökümanını görüntülemek için:
API swagger 2.0 JSON dökümanı: [swagger.json](https://github.com/sfazilyesil/restapi-go-swagger/blob/master/swagger.json)
  
API html dosyasını görüntülek için projeyi indirip, proje ana dizinindeki **apidoc.html**
dosyasını tarayıcıda açmamız gerekmektedir.
  
  
#### API dökümanını koddan yeniden üretmek için:
##### Gereksinimler
- Golang
- bash

##### İndirme/Yükleme
```go
go get github.com/sfazilyesil/restapi-go-swagger
```

##### API dökümanının üretilmesi
Proje dizininde
```go
./apidoc.sh
```
ile örnek uygulamanın api dökümanını yeniden üretebiliyoruz.



### Problem Tanımı
1- Go ile yazılmış bir REST API uygulamasının Swagger 2.0 formatına uygun API dökümanını 
Go kodundan nasıl otomatik olarak üretebiliriz?

2- Üretilen dökümanı nasıl doğrulayabiliriz?

3- Bu işleri günlük çalışma akışına (ilgili build scriptlerine) nasıl dahil edebiliriz?

4- Üretilen dökümanı nasıl görüntüleyebiliriz?

### Çözüm
Swagger 2.0 için ekosistemdeki en iyi golang kütüphanesi go-swagger. Koddan döküman, 
dökümandan istemci ve sunucu, döküman doğrulaması vb. şeyleri yapma yeteneğine sahip.

Bizim odaklandığımız problemler ise koddan otomatik olarak döküman üretimi ve 
üretilen dökümanın doğrulanması ile sınırlı. Go-swagger kütüphanesi bu ihtiyaçları
tamamıyla karşılayabiliyor. Bazı noktalarda başka kütüphanelerin sunduğu bazı kısa 
yolları sunmasa da şu an için en bütünlüklü çözüm sunan kütüphane. Bazı şeyleri 
yapış biçimini beğenmesek de bir yol öneriyor.

### İlgili Bağlantılar
- http://swagger.io/
- https://goswagger.io/
- https://github.com/Rebilly/ReDoc

