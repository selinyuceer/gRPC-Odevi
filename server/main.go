package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "university-library/proto/github.com/university/proto"
)

// Mock data storage
var (
	books    = make(map[string]*pb.Book)
	students = make(map[string]*pb.Student)
	loans    = make(map[string]*pb.Loan)
)

// BookService implementation
type bookService struct {
	pb.UnimplementedBookServiceServer
}

func (s *bookService) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	var bookList []*pb.Book
	for _, book := range books {
		bookList = append(bookList, book)
	}
	return &pb.ListBooksResponse{Books: bookList}, nil
}

func (s *bookService) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	book, exists := books[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Book with id %s not found", req.Id)
	}
	return &pb.GetBookResponse{Book: book}, nil
}

func (s *bookService) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := req.Book
	if book.Id == "" {
		book.Id = fmt.Sprintf("book-%d", len(books)+1)
	}
	books[book.Id] = book
	return &pb.CreateBookResponse{Book: book}, nil
}

func (s *bookService) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	book := req.Book
	if _, exists := books[book.Id]; !exists {
		return nil, status.Errorf(codes.NotFound, "Book with id %s not found", book.Id)
	}
	books[book.Id] = book
	return &pb.UpdateBookResponse{Book: book}, nil
}

func (s *bookService) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	if _, exists := books[req.Id]; !exists {
		return nil, status.Errorf(codes.NotFound, "Book with id %s not found", req.Id)
	}
	delete(books, req.Id)
	return &pb.DeleteBookResponse{Success: true}, nil
}

// StudentService implementation
type studentService struct {
	pb.UnimplementedStudentServiceServer
}

func (s *studentService) ListStudents(ctx context.Context, req *pb.ListStudentsRequest) (*pb.ListStudentsResponse, error) {
	var studentList []*pb.Student
	for _, student := range students {
		studentList = append(studentList, student)
	}
	return &pb.ListStudentsResponse{Students: studentList}, nil
}

func (s *studentService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {
	student, exists := students[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Student with id %s not found", req.Id)
	}
	return &pb.GetStudentResponse{Student: student}, nil
}

func (s *studentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {
	student := req.Student
	if student.Id == "" {
		student.Id = fmt.Sprintf("student-%d", len(students)+1)
	}
	students[student.Id] = student
	return &pb.CreateStudentResponse{Student: student}, nil
}

func (s *studentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentResponse, error) {
	student := req.Student
	if _, exists := students[student.Id]; !exists {
		return nil, status.Errorf(codes.NotFound, "Student with id %s not found", student.Id)
	}
	students[student.Id] = student
	return &pb.UpdateStudentResponse{Student: student}, nil
}

func (s *studentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	if _, exists := students[req.Id]; !exists {
		return nil, status.Errorf(codes.NotFound, "Student with id %s not found", req.Id)
	}
	delete(students, req.Id)
	return &pb.DeleteStudentResponse{Success: true}, nil
}

// LoanService implementation
type loanService struct {
	pb.UnimplementedLoanServiceServer
}

func (s *loanService) ListLoans(ctx context.Context, req *pb.ListLoansRequest) (*pb.ListLoansResponse, error) {
	var loanList []*pb.Loan
	for _, loan := range loans {
		loanList = append(loanList, loan)
	}
	return &pb.ListLoansResponse{Loans: loanList}, nil
}

func (s *loanService) GetLoan(ctx context.Context, req *pb.GetLoanRequest) (*pb.GetLoanResponse, error) {
	loan, exists := loans[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Loan with id %s not found", req.Id)
	}
	return &pb.GetLoanResponse{Loan: loan}, nil
}

func (s *loanService) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	// Check if student exists
	if _, exists := students[req.StudentId]; !exists {
		return nil, status.Errorf(codes.NotFound, "Student with id %s not found", req.StudentId)
	}

	// Check if book exists and has stock
	book, exists := books[req.BookId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Book with id %s not found", req.BookId)
	}
	if book.Stock <= 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "Book is out of stock")
	}

	// Create loan
	loanId := fmt.Sprintf("loan-%d", len(loans)+1)
	loan := &pb.Loan{
		Id:        loanId,
		StudentId: req.StudentId,
		BookId:    req.BookId,
		LoanDate:  time.Now().Format("2006-01-02"),
		Status:    pb.LoanStatus_ONGOING,
	}

	// Update book stock
	book.Stock--
	books[req.BookId] = book

	// Store loan
	loans[loanId] = loan

	return &pb.BorrowBookResponse{Loan: loan}, nil
}

func (s *loanService) ReturnBook(ctx context.Context, req *pb.ReturnBookRequest) (*pb.ReturnBookResponse, error) {
	loan, exists := loans[req.LoanId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Loan with id %s not found", req.LoanId)
	}

	if loan.Status != pb.LoanStatus_ONGOING {
		return nil, status.Errorf(codes.FailedPrecondition, "Loan is not ongoing")
	}

	// Update loan
	loan.Status = pb.LoanStatus_RETURNED
	loan.ReturnDate = time.Now().Format("2006-01-02")
	loans[req.LoanId] = loan

	// Update book stock
	if book, exists := books[loan.BookId]; exists {
		book.Stock++
		books[loan.BookId] = book
	}

	return &pb.ReturnBookResponse{Loan: loan}, nil
}

func initMockData() {
	// Sample books
	books["book-1"] = &pb.Book{
		Id:        "book-1",
		Title:     "The Go Programming Language",
		Author:    "Alan Donovan, Brian Kernighan",
		Isbn:      "978-0134190440",
		Publisher: "Addison-Wesley",
		PageCount: 380,
		Stock:     5,
	}

	books["book-2"] = &pb.Book{
		Id:        "book-2",
		Title:     "Clean Code",
		Author:    "Robert C. Martin",
		Isbn:      "978-0132350884",
		Publisher: "Prentice Hall",
		PageCount: 464,
		Stock:     3,
	}

	// Sample students
	students["student-1"] = &pb.Student{
		Id:            "student-1",
		Name:          "Ahmet Yılmaz",
		StudentNumber: "20210001",
		Email:         "ahmet.yilmaz@university.edu",
		IsActive:      true,
	}

	students["student-2"] = &pb.Student{
		Id:            "student-2",
		Name:          "Ayşe Kaya",
		StudentNumber: "20210002",
		Email:         "ayse.kaya@university.edu",
		IsActive:      true,
	}
}

func main() {
	// Initialize mock data
	initMockData()

	// Create listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server
	s := grpc.NewServer()

	// Register services
	pb.RegisterBookServiceServer(s, &bookService{})
	pb.RegisterStudentServiceServer(s, &studentService{})
	pb.RegisterLoanServiceServer(s, &loanService{})

	// Register reflection service for grpcurl
	reflection.Register(s)

	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
} 