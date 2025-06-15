# University Library gRPC System

Bu proje, bir üniversite kütüphanesi için geliştirilmiş gRPC tabanlı bir servistir. Protocol Buffers kullanılarak API tanımları yapılmış ve Go programlama dili ile sunucu ve istemci uygulamaları geliştirilmiştir.

## 🏗️ Sistem Mimarisi

Sistem üç ana servisten oluşmaktadır:

### 📚 BookService
- **ListBooks**: Tüm kitapları listeler
- **GetBook**: Belirli bir kitabı getirir
- **CreateBook**: Yeni kitap ekler
- **UpdateBook**: Mevcut kitabı günceller
- **DeleteBook**: Kitabı siler

### 👤 StudentService
- **ListStudents**: Tüm öğrencileri listeler
- **GetStudent**: Belirli bir öğrenciyi getirir
- **CreateStudent**: Yeni öğrenci ekler
- **UpdateStudent**: Mevcut öğrenciyi günceller
- **DeleteStudent**: Öğrenciyi siler

### 🔄 LoanService
- **ListLoans**: Tüm ödünç verme kayıtlarını listeler
- **GetLoan**: Belirli bir ödünç kaydını getirir
- **BorrowBook**: Kitap ödünç verme işlemi
- **ReturnBook**: Kitap iade işlemi

## 📋 Veri Modelleri

### Book
```proto
message Book {
  string id = 1;
  string title = 2;
  string author = 3;
  string isbn = 4;
  string publisher = 5;
  int32 page_count = 6;
  int32 stock = 7;
}
```

### Student
```proto
message Student {
  string id = 1;
  string name = 2;
  string student_number = 3;
  string email = 4;
  bool is_active = 5;
}
```

### Loan
```proto
message Loan {
  string id = 1;
  string student_id = 2;
  string book_id = 3;
  string loan_date = 4;
  string return_date = 5;
  LoanStatus status = 6;
}
```

### LoanStatus Enum
```proto
enum LoanStatus {
  LOAN_STATUS_UNSPECIFIED = 0;
  ONGOING = 1;
  RETURNED = 2;
  LATE = 3;
}
```

## 🚀 Kurulum ve Çalıştırma

### Gereksinimler
- Go 1.21 veya üzeri
- Protocol Buffers compiler (protoc)
- grpcurl (test için)

### 1. Bağımlılıkları İndirin
```bash
go mod tidy
```

### 2. Protocol Buffers Kodlarını Oluşturun
```bash
# Proto dizini oluştur
mkdir -p proto

# Protocol buffers kodlarını oluştur
protoc --go_out=proto --go-grpc_out=proto university.proto
```

### 3. Sunucuyu Çalıştırın
```bash
go run server/main.go
```

Sunucu `localhost:50051` portunda çalışmaya başlayacaktır.

### 4. İstemciyi Çalıştırın (Ayrı Terminal)
```bash
go run client/main.go
```

## 🧪 grpcurl ile Test

### Servisleri Keşfetme
```bash
# Mevcut servisleri listele
grpcurl -plaintext localhost:50051 list

# BookService metotlarını listele
grpcurl -plaintext localhost:50051 list university.BookService
```

### Book Service Test
```bash
# Kitapları listele
grpcurl -plaintext localhost:50051 university.BookService/ListBooks

# Belirli kitabı getir
grpcurl -plaintext -d '{"id": "book-1"}' localhost:50051 university.BookService/GetBook

# Yeni kitap ekle
grpcurl -plaintext -d '{
  "book": {
    "title": "Effective Go",
    "author": "Go Team",
    "isbn": "978-1234567890",
    "publisher": "Google",
    "page_count": 200,
    "stock": 5
  }
}' localhost:50051 university.BookService/CreateBook
```

### Student Service Test
```bash
# Öğrencileri listele
grpcurl -plaintext localhost:50051 university.StudentService/ListStudents

# Belirli öğrenciyi getir
grpcurl -plaintext -d '{"id": "student-1"}' localhost:50051 university.StudentService/GetStudent

# Yeni öğrenci ekle
grpcurl -plaintext -d '{
  "student": {
    "name": "Ali Veli",
    "student_number": "20210004",
    "email": "ali.veli@university.edu",
    "is_active": true
  }
}' localhost:50051 university.StudentService/CreateStudent
```

### Loan Service Test
```bash
# Ödünç kayıtlarını listele
grpcurl -plaintext localhost:50051 university.LoanService/ListLoans

# Kitap ödünç ver
grpcurl -plaintext -d '{
  "student_id": "student-1",
  "book_id": "book-1"
}' localhost:50051 university.LoanService/BorrowBook

# Kitap iade et
grpcurl -plaintext -d '{"loan_id": "loan-1"}' localhost:50051 university.LoanService/ReturnBook
```

## 📁 Proje Yapısı

```
university-library/
├── university.proto          # Protocol Buffers tanımları
├── go.mod                   # Go modül dosyası
├── server/
│   └── main.go             # gRPC sunucu uygulaması
├── client/
│   └── main.go             # gRPC istemci uygulaması
├── proto/                  # Oluşturulan Go kodları (gitignore'da)
├── README.md               # Bu dosya
├── grpcurl-tests.md        # Test dokümantasyonu
└── gRPC_Delivery.md        # Ödev teslim formu
```

## 🎯 Özellikler

- ✅ Protocol Buffers v3 kullanımı
- ✅ Üç ana servis (Book, Student, Loan)
- ✅ CRUD operasyonları
- ✅ Enum kullanımı (LoanStatus)
- ✅ İlişkisel veri yönetimi
- ✅ Hata yönetimi
- ✅ Mock veri ile test
- ✅ grpcurl uyumluluğu

## 🔧 Teknik Detaylar

- **Programlama Dili**: Go 1.21
- **Protocol Buffers**: proto3
- **gRPC Framework**: google.golang.org/grpc
- **Port**: 50051
- **Veri Depolama**: In-memory (Mock data)

## 📝 Notlar

- Stub dosyalar repository'e dahil edilmemiştir
- Veri kalıcılığı için in-memory storage kullanılmıştır
- Production ortamında gerçek bir veritabanı kullanılmalıdır
- grpcurl testleri için detaylı dokümantasyon `grpcurl-tests.md` dosyasında bulunmaktadır

## 👨‍💻 Geliştirici

Bu proje Açık Kaynak Kodlu Yazılımlar dersi kapsamında geliştirilmiştir. 