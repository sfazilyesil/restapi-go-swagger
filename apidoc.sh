#!/usr/bin/env bash

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
