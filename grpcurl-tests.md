# gRPC Servisleri grpcurl Test Dokümantasyonu

Bu dokümanda University Library gRPC servislerinin tüm metotları için grpcurl test komutları ve beklenen yanıtlar yer almaktadır.

## 📋 Test Öncesi Hazırlık

Testleri yapmadan önce sunucunun çalıştığından emin olun:
```bash
go run server/main.go
```

## 🔍 Servis Keşfi

### Mevcut Servisleri Listele
```bash
grpcurl -plaintext localhost:50051 list
```

**Beklenen Yanıt:**
```
grpc.reflection.v1alpha.ServerReflection
university.BookService
university.LoanService
university.StudentService
```

### BookService Metotlarını Listele
```bash
grpcurl -plaintext localhost:50051 list university.BookService
```

**Beklenen Yanıt:**
```
university.BookService.CreateBook
university.BookService.DeleteBook
university.BookService.GetBook
university.BookService.ListBooks
university.BookService.UpdateBook
```

## 📚 BookService Testleri

### 1. ListBooks - Tüm Kitapları Listele
```bash
grpcurl -plaintext localhost:50051 university.BookService/ListBooks
```

**Beklenen Yanıt:**
```json
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
```

### 2. GetBook - Belirli Kitabı Getir
```bash
grpcurl -plaintext -d '{"id": "book-1"}' localhost:50051 university.BookService/GetBook
```

**Beklenen Yanıt:**
```json
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

### 3. GetBook - Hatalı ID (404 Senaryosu)
```bash
grpcurl -plaintext -d '{"id": "non-existent-book"}' localhost:50051 university.BookService/GetBook
```

**Beklenen Hata:**
```
ERROR:
  Code: NotFound
  Message: Book with id non-existent-book not found
```

### 4. CreateBook - Yeni Kitap Ekle
```bash
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

**Beklenen Yanıt:**
```json
{
  "book": {
    "id": "book-3",
    "title": "Effective Go",
    "author": "Go Team",
    "isbn": "978-1234567890",
    "publisher": "Google",
    "pageCount": 200,
    "stock": 5
  }
}
```

### 5. UpdateBook - Kitap Güncelle
```bash
grpcurl -plaintext -d '{
  "book": {
    "id": "book-1",
    "title": "The Go Programming Language - Updated",
    "author": "Alan Donovan, Brian Kernighan",
    "isbn": "978-0134190440",
    "publisher": "Addison-Wesley",
    "page_count": 380,
    "stock": 10
  }
}' localhost:50051 university.BookService/UpdateBook
```

**Beklenen Yanıt:**
```json
{
  "book": {
    "id": "book-1",
    "title": "The Go Programming Language - Updated",
    "author": "Alan Donovan, Brian Kernighan",
    "isbn": "978-0134190440",
    "publisher": "Addison-Wesley",
    "pageCount": 380,
    "stock": 10
  }
}
```

### 6. DeleteBook - Kitap Sil
```bash
grpcurl -plaintext -d '{"id": "book-2"}' localhost:50051 university.BookService/DeleteBook
```

**Beklenen Yanıt:**
```json
{
  "success": true
}
```

## 👤 StudentService Testleri

### 1. ListStudents - Tüm Öğrencileri Listele
```bash
grpcurl -plaintext localhost:50051 university.StudentService/ListStudents
```

**Beklenen Yanıt:**
```json
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

### 2. GetStudent - Belirli Öğrenciyi Getir
```bash
grpcurl -plaintext -d '{"id": "student-1"}' localhost:50051 university.StudentService/GetStudent
```

**Beklenen Yanıt:**
```json
{
  "student": {
    "id": "student-1",
    "name": "Ahmet Yılmaz",
    "studentNumber": "20210001",
    "email": "ahmet.yilmaz@university.edu",
    "isActive": true
  }
}
```

### 3. CreateStudent - Yeni Öğrenci Ekle
```bash
grpcurl -plaintext -d '{
  "student": {
    "name": "Ali Veli",
    "student_number": "20210004",
    "email": "ali.veli@university.edu",
    "is_active": true
  }
}' localhost:50051 university.StudentService/CreateStudent
```

**Beklenen Yanıt:**
```json
{
  "student": {
    "id": "student-3",
    "name": "Ali Veli",
    "studentNumber": "20210004",
    "email": "ali.veli@university.edu",
    "isActive": true
  }
}
```

### 4. UpdateStudent - Öğrenci Güncelle
```bash
grpcurl -plaintext -d '{
  "student": {
    "id": "student-1",
    "name": "Ahmet Yılmaz",
    "student_number": "20210001",
    "email": "ahmet.yilmaz@university.edu",
    "is_active": false
  }
}' localhost:50051 university.StudentService/UpdateStudent
```

### 5. DeleteStudent - Öğrenci Sil
```bash
grpcurl -plaintext -d '{"id": "student-2"}' localhost:50051 university.StudentService/DeleteStudent
```

**Beklenen Yanıt:**
```json
{
  "success": true
}
```

## 🔄 LoanService Testleri

### 1. ListLoans - Tüm Ödünç Kayıtlarını Listele (Başlangıçta Boş)
```bash
grpcurl -plaintext localhost:50051 university.LoanService/ListLoans
```

**Beklenen Yanıt (Başlangıçta):**
```json
{
  "loans": []
}
```

### 2. BorrowBook - Kitap Ödünç Ver
```bash
grpcurl -plaintext -d '{
  "student_id": "student-1",
  "book_id": "book-1"
}' localhost:50051 university.LoanService/BorrowBook
```

**Beklenen Yanıt:**
```json
{
  "loan": {
    "id": "loan-1",
    "studentId": "student-1",
    "bookId": "book-1",
    "loanDate": "2024-01-15",
    "status": "ONGOING"
  }
}
```

### 3. BorrowBook - Hatalı Student ID (404 Senaryosu)
```bash
grpcurl -plaintext -d '{
  "student_id": "non-existent-student",
  "book_id": "book-1"
}' localhost:50051 university.LoanService/BorrowBook
```

**Beklenen Hata:**
```
ERROR:
  Code: NotFound
  Message: Student with id non-existent-student not found
```

### 4. BorrowBook - Hatalı Book ID (404 Senaryosu)
```bash
grpcurl -plaintext -d '{
  "student_id": "student-1",
  "book_id": "non-existent-book"
}' localhost:50051 university.LoanService/BorrowBook
```

**Beklenen Hata:**
```
ERROR:
  Code: NotFound
  Message: Book with id non-existent-book not found
```

### 5. GetLoan - Belirli Ödünç Kaydını Getir
```bash
grpcurl -plaintext -d '{"id": "loan-1"}' localhost:50051 university.LoanService/GetLoan
```

**Beklenen Yanıt:**
```json
{
  "loan": {
    "id": "loan-1",
    "studentId": "student-1",
    "bookId": "book-1",
    "loanDate": "2024-01-15",
    "status": "ONGOING"
  }
}
```

### 6. ReturnBook - Kitap İade Et
```bash
grpcurl -plaintext -d '{"loan_id": "loan-1"}' localhost:50051 university.LoanService/ReturnBook
```

**Beklenen Yanıt:**
```json
{
  "loan": {
    "id": "loan-1",
    "studentId": "student-1",
    "bookId": "book-1",
    "loanDate": "2024-01-15",
    "returnDate": "2024-01-20",
    "status": "RETURNED"
  }
}
```

### 7. ReturnBook - Hatalı Loan ID (404 Senaryosu)
```bash
grpcurl -plaintext -d '{"loan_id": "non-existent-loan"}' localhost:50051 university.LoanService/ReturnBook
```

**Beklenen Hata:**
```
ERROR:
  Code: NotFound
  Message: Loan with id non-existent-loan not found
```

### 8. ListLoans - Ödünç Kayıtlarını Listele (İşlem Sonrası)
```bash
grpcurl -plaintext localhost:50051 university.LoanService/ListLoans
```

**Beklenen Yanıt:**
```json
{
  "loans": [
    {
      "id": "loan-1",
      "studentId": "student-1",
      "bookId": "book-1",
      "loanDate": "2024-01-15",
      "returnDate": "2024-01-20",
      "status": "RETURNED"
    }
  ]
}
```

## 📊 Test Sonuçları Özeti

- ✅ **BookService**: 5/5 metot başarılı
  - ListBooks, GetBook, CreateBook, UpdateBook, DeleteBook
- ✅ **StudentService**: 5/5 metot başarılı
  - ListStudents, GetStudent, CreateStudent, UpdateStudent, DeleteStudent
- ✅ **LoanService**: 4/4 metot başarılı
  - ListLoans, GetLoan, BorrowBook, ReturnBook

## 🎯 Enum Kullanımı

LoanStatus enum'u başarıyla test edildi:
- **ONGOING**: Devam eden ödünç alma
- **RETURNED**: İade edilmiş kitap
- **LATE**: Geç iade (manuel test gerekli)

## 🚨 Hata Senaryoları

Tüm servisler aşağıdaki hata durumlarını doğru şekilde yönetiyor:
- **NotFound (404)**: Kayıt bulunamadığında
- **FailedPrecondition**: İş kuralı ihlallerinde (örn: stok yok)
- **InvalidArgument**: Geçersiz parametrelerde 