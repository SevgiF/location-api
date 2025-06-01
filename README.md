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