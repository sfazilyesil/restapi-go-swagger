# restapi-go-swagger

**Golang** ile yazılmış web servislerin **[swagger 2.0](http://swagger.io/)** formatına uygun
**API dökümanı**nı otomatik olarak nasıl üretilebileceğimizi göstermek amacıyla 
hazırlanmış bir **örnek uygulama**.

Örnek uygulamanın genel anlatımı için: [bkz](https://github.com/sfazilyesil/restapi-go-swagger/blob/master/tutorial.md)

<br/>

### Problem Tanımı
1- Go ile yazılmış bir REST API uygulamasının Swagger 2.0 formatına uygun API dökümanını 
Go kodundan nasıl otomatik olarak üretebiliriz?

2- Üretilen dökümanı nasıl doğrulayabiliriz?

3- Bu işleri günlük çalışma akışına (ilgili build scriptlerine) nasıl dahil edebiliriz?

4- Üretilen dökümanı nasıl görüntüleyebiliriz?


### Çözüm
[go-swagger](https://goswagger.io/) kütüphanesi ile go kodu ve yorum satırındaki dökümantasyon 
direktifleriyle otomatik olarak döküman üretmek ve doğrulamak mümkün. Görüntülemek içinse
[ReDoc](https://github.com/Rebilly/ReDoc) güzel seçeneklerden biri. Build akışına entegre etmek 
içinse apidoc.sh gibi bir script yazıp onu ilgili akıştan çağırmak yeterli.

<br/>

*** 
  
#### Hazır üretilmiş API dökümanını görüntülemek için:
API swagger 2.0 JSON dökümanı: [swagger.json](https://github.com/sfazilyesil/restapi-go-swagger/blob/master/swagger.json)
  
API html dosyasını görüntülek için projeyi indirip, proje ana dizinindeki **apidoc.html**
dosyasını tarayıcıda açmamız gerekmektedir.

***

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


***

<br/>

### İlgili Bağlantılar
- http://swagger.io/
- https://goswagger.io/
- https://github.com/Rebilly/ReDoc

