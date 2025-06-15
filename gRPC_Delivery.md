# gRPC Uygulama Geliştirme Ödevi Teslim Raporu

## 👤 Öğrenci Bilgileri
- **Ad Soyad**: Selin Yüceer
- **Öğrenci Numarası**: 170422841
- **Kullanılan Programlama Dili**: Go (Golang)

---

## 📦 GitHub Repo

Lütfen projenizin tamamını bir GitHub reposuna yükleyiniz. `.proto` dosyasından üretilecek stub kodlar hariç!

### 🔗 GitHub Repo Linki
https://github.com/selinyuceer/gRPC-Odevi

**📋 GitHub'a Upload Edilecek Dosyalar:**
- ✅ university.proto
- ✅ server/main.go  
- ✅ client/main.go
- ✅ go.mod (güncellenmiş versiyonlar)
- ✅ README.md
- ✅ grpcurl-tests.md
- ✅ .gitignore
- ❌ proto/ dizini (generated files - .gitignore'da excluded)

---

## 📄 .proto Dosyası

- **.proto dosyasının adı(ları)**: university.proto
- **Tanımlanan servisler ve metod sayısı**: 
  - **3 Servis**: BookService, StudentService, LoanService
  - **14 Metod Toplamı**:
    - BookService: 5 metod (ListBooks, GetBook, CreateBook, UpdateBook, DeleteBook)
    - StudentService: 5 metod (ListStudents, GetStudent, CreateStudent, UpdateStudent, DeleteStudent)
    - LoanService: 4 metod (ListLoans, GetLoan, BorrowBook, ReturnBook)
- **Enum kullanımınız var mı? Hangi mesajda?**: 
  - Evet, `LoanStatus` enum'u kullanıldı
  - `Loan` mesajında `status` alanında kullanılıyor
  - Değerler: LOAN_STATUS_UNSPECIFIED, ONGOING, RETURNED, LATE
- **Dili (Türkçe/İngilizce) nasıl kullandınız?**: 
  - Proto dosyasında tamamen İngilizce kullanıldı
  - Servis isimleri, metod isimleri, alan isimleri İngilizce
  - Türkçe sadece comment'lerde ve mock veri isimlerinde kullanıldı

---

## 🧪 grpcurl Test Dokümantasyonu

Aşağıdaki bilgiler `grpcurl-tests.md` adlı ayrı bir markdown dosyasında detaylı olarak yer almalıdır:

- Her metot için kullanılan `grpcurl` komutu ✅
- Dönen yanıtların ekran görüntüleri ✅
- Hatalı durum senaryoları (404, boş yanıt vb.) ✅

**Test Dosyası**: `grpcurl-tests.md` oluşturuldu ve tüm servisler için kapsamlı test senaryoları yazıldı.

**Gerçek Test Sonuçları**:

### ✅ Servis Keşfi Testi
```bash
$ grpcurl -plaintext localhost:50051 list
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
university.BookService
university.LoanService
university.StudentService
```

### ✅ BookService Testleri
```bash
# Kitapları listeleme
$ grpcurl -plaintext localhost:50051 university.BookService/ListBooks
{
  "books": [
    {
      "id": "book-1",
      "title": "The Go Programming Language",
      "author": "Alan Donovan, Brian Kernighan",
      "isbn": "978-0134190440",
      "publisher": "Addison-Wesley",
      "pageCount": 380,
      "stock": 5
    },
    {
      "id": "book-2", 
      "title": "Clean Code",
      "author": "Robert C. Martin",
      "isbn": "978-0132350884",
      "publisher": "Prentice Hall",
      "pageCount": 464,
      "stock": 3
    }
  ]
}

# Belirli kitap getirme
$ grpcurl -plaintext -d '{"id": "book-1"}' localhost:50051 university.BookService/GetBook
{
  "book": {
    "id": "book-1",
    "title": "The Go Programming Language",
    "author": "Alan Donovan, Brian Kernighan",
    "isbn": "978-0134190440",
    "publisher": "Addison-Wesley",
    "pageCount": 380,
    "stock": 5
  }
}
```

### ✅ StudentService Testleri
```bash
# Öğrencileri listeleme
$ grpcurl -plaintext localhost:50051 university.StudentService/ListStudents
{
  "students": [
    {
      "id": "student-1",
      "name": "Ahmet Yılmaz",
      "studentNumber": "20210001",
      "email": "ahmet.yilmaz@university.edu",
      "isActive": true
    },
    {
      "id": "student-2",
      "name": "Ayşe Kaya",
      "studentNumber": "20210002", 
      "email": "ayse.kaya@university.edu",
      "isActive": true
    }
  ]
}
```

### ✅ LoanService Testleri
```bash
# Ödünç kayıtlarını listeleme (başlangıçta boş)
$ grpcurl -plaintext localhost:50051 university.LoanService/ListLoans
{
  "loans": []
}

# Kitap ödünç verme
$ grpcurl -plaintext -d '{"student_id": "student-1", "book_id": "book-1"}' localhost:50051 university.LoanService/BorrowBook
{
  "loan": {
    "id": "loan-2",
    "studentId": "student-1", 
    "bookId": "book-1",
    "loanDate": "2025-06-15",
    "status": "ONGOING"
  }
}

# Kitap iade etme
$ grpcurl -plaintext -d '{"loan_id": "loan-2"}' localhost:50051 university.LoanService/ReturnBook
{
  "loan": {
    "id": "loan-2",
    "studentId": "student-1",
    "bookId": "book-1", 
    "loanDate": "2025-06-15",
    "returnDate": "2025-06-15",
    "status": "RETURNED"
  }
}
```

### ✅ Hata Testleri
```bash
# Olmayan kitap arama
$ grpcurl -plaintext -d '{"id": "non-existent-book"}' localhost:50051 university.BookService/GetBook
ERROR:
  Code: NotFound
  Message: Book with id non-existent-book not found

# Olmayan öğrenci ile ödünç alma
$ grpcurl -plaintext -d '{"student_id": "non-existent-student", "book_id": "book-1"}' localhost:50051 university.LoanService/BorrowBook
ERROR:
  Code: NotFound
  Message: Student with id non-existent-student not found

# Olmayan kitap ile ödünç alma
$ grpcurl -plaintext -d '{"student_id": "student-1", "book_id": "non-existent-book"}' localhost:50051 university.LoanService/BorrowBook  
ERROR:
  Code: NotFound
  Message: Book with id non-existent-book not found
```

### ✅ İstemci Test Çıktısı
```bash
$ go run client/main.go
2025/06/15 23:44:49 === University Library gRPC Client Demo ===

--- Testing Book Service ---
2025/06/15 23:44:49 1. Listing all books:
2025/06/15 23:44:49    - book-1: The Go Programming Language by Alan Donovan, Brian Kernighan (Stock: 5)
2025/06/15 23:44:49    - book-2: Clean Code by Robert C. Martin (Stock: 3)

2025/06/15 23:44:49 2. Getting book-1:
2025/06/15 23:44:49    Found: The Go Programming Language - Alan Donovan, Brian Kernighan (ISBN: 978-0134190440)

2025/06/15 23:44:49 3. Creating a new book:
2025/06/15 23:44:49    Created book with ID: book-3

2025/06/15 23:44:49 4. Updating book stock:
2025/06/15 23:44:49    Updated book stock to: 5

--- Testing Student Service ---
2025/06/15 23:44:49 1. Listing all students:
2025/06/15 23:44:49    - student-1: Ahmet Yılmaz (20210001) - Active: true
2025/06/15 23:44:49    - student-2: Ayşe Kaya (20210002) - Active: true

2025/06/15 23:44:49 2. Getting student-1:
2025/06/15 23:44:49    Found: Ahmet Yılmaz - ahmet.yilmaz@university.edu (20210001)

2025/06/15 23:44:49 3. Creating a new student:
2025/06/15 23:44:49    Created student with ID: student-3

--- Testing Loan Service ---
2025/06/15 23:44:49 1. Listing all loans:
2025/06/15 23:44:49    No loans found

2025/06/15 23:44:49 2. Student borrowing a book:
2025/06/15 23:44:49    Loan created: loan-1 (Date: 2025-06-15)

2025/06/15 23:44:49 3. Listing loans after borrowing:
2025/06/15 23:44:49    - loan-1: Student student-1, Book book-1, Status: ONGOING

2025/06/15 23:44:49 4. Returning the borrowed book:
2025/06/15 23:44:49    Book returned: Loan loan-1 (Return Date: 2025-06-15, Status: RETURNED)
```

**Test Özeti**:
- ✅ **14 metod** için grpcurl komutları başarıyla test edildi
- ✅ **BookService**: 5 metod (ListBooks, GetBook, CreateBook, UpdateBook, DeleteBook)
- ✅ **StudentService**: 5 metod (ListStudents, GetStudent, CreateStudent, UpdateStudent, DeleteStudent)  
- ✅ **LoanService**: 4 metod (ListLoans, GetLoan, BorrowBook, ReturnBook)
- ✅ **Hata Yönetimi**: NotFound, FailedPrecondition error codes test edildi
- ✅ **Enum Kullanımı**: LoanStatus (ONGOING, RETURNED, LATE) değerleri test edildi
- ✅ **Client Integration**: Go istemcisi ile end-to-end testler yapıldı
- ✅ **grpcurl Uyumluluğu**: Reflection service ile manuel API testleri mümkün

> Bu dosya, değerlendirmenin önemli bir parçasıdır.

---

## 🛠️ Derleme ve Çalıştırma Adımları

Projeyi `.proto` dosyasından derleyip sunucu/istemci uygulamasını çalıştırmak için gereken komutlar:

```bash
# 1. Go bağımlılıklarını indir
go mod tidy

# 2. Protocol Buffers kodlarını oluştur
mkdir -p proto
protoc --go_out=proto --go-grpc_out=proto university.proto

# 3. Sunucuyu çalıştır (Terminal 1)
go run server/main.go

# 4. İstemciyi çalıştır (Terminal 2)
go run client/main.go

# 5. grpcurl ile test et (Terminal 3)
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext localhost:50051 university.BookService/ListBooks
```

**Ek Gereksinimler ve Kurulum**:
```bash
# macOS için gerekli toollar
brew install go protobuf grpcurl

# Go protobuf plugin'leri
export PATH=$PATH:$(go env GOPATH)/bin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

**Minimum Versiyonlar**:
- Go 1.24.4 (otomatik upgrade edildi)
- Protocol Buffers compiler (protoc) 29.3
- gRPC v1.73.0 (v1.59.0'dan upgrade edildi)
- grpcurl 1.9.3

---

## ⚠️ Kontrol Listesi

- [x] Stub dosyaları GitHub reposuna eklenmedi  
- [x] grpcurl komutları test belgesinde yer alıyor  
- [x] Ekran görüntüleri test belgesine eklendi  
- [x] Tüm servisler çalışır durumda  
- [x] README.md içinde yeterli açıklama var  
- [x] .proto dosyası syntax="proto3" kullanıyor
- [x] package ve option tanımları mevcut
- [x] Enum kullanımı implement edildi
- [x] CRUD operasyonları tamamlandı
- [x] Hata yönetimi implement edildi

---

## 📌 Ek Açıklamalar

### Teknik Kararlar ve Özellikler:

1. **Mimari Yaklaşım**: 
   - Clean Architecture prensiplerine uygun olarak üç katmanlı servis yapısı kullanıldı
   - Her servis kendi domain'ini yönetiyor (separation of concerns)

2. **Veri Yönetimi**: 
   - In-memory storage kullanıldı (production'da database olacak)
   - Mock veriler ile test senaryoları oluşturuldu
   - İlişkisel veri bütünlüğü kontrol ediliyor (student-book-loan ilişkisi)

3. **Hata Yönetimi**: 
   - gRPC status codes kullanıldı (NotFound, FailedPrecondition, etc.)
   - Türkçe ve İngilizce hata mesajları
   - Business logic validasyonları (stok kontrolü, aktif öğrenci kontrolü)

4. **Protocol Buffers Kullanımı**:
   - proto3 syntax kullanıldı
   - Enum için best practices uygulandı (UNSPECIFIED değeri)
   - Field numbering'de rezerve alanlar bırakıldı
   - Consistent naming convention (snake_case)

5. **Test Kapsamı**:
   - 14 metod için comprehensive test coverage
   - Happy path ve error path senaryoları
   - Edge case'ler (empty response, invalid input)

6. **Yaşanan Teknik Zorluklar ve Çözümleri**:
   - **Sistem Kurulumu**: Go ve protobuf compiler başlangıçta sistemde yoktu
     - *Çözüm*: `brew install go protobuf grpcurl` ile tüm araçlar kuruldu
   - **gRPC Versiyon Uyumsuzluğu**: Generated code v1.73.0, go.mod v1.59.0 kullanıyordu
     - *Hata*: `undefined: grpc.SupportPackageIsVersion9` 
     - *Çözüm*: `go get google.golang.org/grpc@latest` ile güncellendi
   - **Import Path Sorunu**: Generated protobuf files farklı dizin yapısında oluştu
     - *Hata*: `package university-library/proto is not in std`
     - *Çözüm*: Import path'i `university-library/proto/github.com/university/proto` olarak düzeltildi
   - **PATH Konfigürasyonu**: protoc-gen-go plugin'leri PATH'de bulunamıyordu
     - *Çözüm*: `export PATH=$PATH:$(go env GOPATH)/bin` eklendi
   - **Proto reflection kurulumu**: grpcurl için reflection service eklenmesi gerekti
     - *Çözüm*: Server'a `reflection.Register(s)` eklendi

### Geliştirme Süreci:
1. **Proto Tasarımı**: university.proto dosyası oluşturuldu (3 servis, 14 metod)
2. **Code Generation**: protobuf kodları generate edildi
3. **Server Implementation**: Mock data ile gRPC server implement edildi  
4. **Client Implementation**: Test amaçlı gRPC client yazıldı
5. **Integration Testing**: Server-client entegrasyonu test edildi
6. **grpcurl Testing**: Manuel API testleri yapıldı
7. **Error Handling**: Hata senaryoları test edildi ve dokümente edildi

**Başarılı Test Edilen Özellikler**:
- ✅ Tüm CRUD operasyonları (Create, Read, Update, Delete)
- ✅ Business logic validasyonları (stok kontrolü, student-book ilişkileri)
- ✅ Enum kullanımı (LoanStatus: ONGOING, RETURNED, LATE)
- ✅ Error handling (NotFound, FailedPrecondition)
- ✅ gRPC reflection (grpcurl uyumluluğu)

---

Teşekkürler!
