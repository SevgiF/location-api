# 📍LOCATION API
Bu proje, Golang kullanılarak geliştirilen bir konum yönetimi uygulamasıdır. Uygulama, kullanıcıların konum eklemesini, listelemesini, düzenlemesini ve harita üzerinde rota oluşturmasını sağlar.

## 📌ENDPOINTS

### 1. Konum Ekleme
- Gönderilen JSON verileri içerisinde:
    - Enlem (latitude)
    - Boylam (longitude)
    - Konum adı (name)
    - Marker rengi (color) – hexadecimal formatında
- Bu bilgiler veritabanına kaydedilir.

### 2. Konumları Listeleme
- Kaydedilmiş tüm konumları listeler.
- Her bir kayıt şu bilgileri içerir:
    - Konum adı
    - Kullanıcının seçtiği marker rengi
    - Enlem ve boylam bilgileri

### 3. Konum Detayı
- Tek bir konuma ait bilgileri döner.
- Parametre olarak konum ID’si beklenir.

### 4. Konum Düzenleme Sayfası
- Var olan konumların bilgileri güncellenebilir.
- Güncellenebilir alanlar:
    - Konum adı
    - Enlem, boylam
    - Marker rengi

### 5. Rotalama
- Gönderilen konuma en yakın konumdan başlayarak rota listesi verir.
- Rota Haversine formülüne göre hesaplanır.

## 🚀 Otomatik Deploy Süreci

Bu proje, GitHub Actions ile otomatik build ve deploy sürecine sahiptir. Ana özellikler:

🔧 Build
•	Her push işlemi sonrası (main ve dev branch’leri), proje Docker imajı olarak build edilir ve Docker Hub’a gönderilir.
•	Kullanılan image: sevgifidan/location-app:latest

💻 Local Deploy (self-hosted runner)
•	Eğer işlem bir self-hosted runner (örn. kendi bilgisayarınız) üzerinde çalışıyorsa:
•	docker-compose.yml dosyası lokal bir klasöre kopyalanır
•	İmaj çekilir ve docker-compose up -d ile çalıştırılır

☁️ EC2 Deploy (sadece göstermek için dev branch'e eklendi)
•	dev branch’e yapılan push’lar ayrıca bir AWS EC2 sunucusuna deploy edilir:
•	docker-compose.yml dosyası EC2’ye SCP ile kopyalanır
•	Uzakta docker-compose pull && up -d komutları çalıştırılır


## Test

Benchmark testi eklendi. Integration testleri de eklenebilir.

