package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "university-library/proto/github.com/university/proto"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Create clients
	bookClient := pb.NewBookServiceClient(conn)
	studentClient := pb.NewStudentServiceClient(conn)
	loanClient := pb.NewLoanServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	log.Println("=== University Library gRPC Client Demo ===")

	// Test Book Service
	log.Println("\n--- Testing Book Service ---")
	testBookService(ctx, bookClient)

	// Test Student Service
	log.Println("\n--- Testing Student Service ---")
	testStudentService(ctx, studentClient)

	// Test Loan Service
	log.Println("\n--- Testing Loan Service ---")
	testLoanService(ctx, loanClient)
}

func testBookService(ctx context.Context, client pb.BookServiceClient) {
	// List books
	log.Println("1. Listing all books:")
	listResp, err := client.ListBooks(ctx, &pb.ListBooksRequest{})
	if err != nil {
		log.Printf("Error listing books: %v", err)
		return
	}
	for _, book := range listResp.Books {
		log.Printf("   - %s: %s by %s (Stock: %d)", book.Id, book.Title, book.Author, book.Stock)
	}

	// Get a specific book
	log.Println("\n2. Getting book-1:")
	getResp, err := client.GetBook(ctx, &pb.GetBookRequest{Id: "book-1"})
	if err != nil {
		log.Printf("Error getting book: %v", err)
	} else {
		book := getResp.Book
		log.Printf("   Found: %s - %s (ISBN: %s)", book.Title, book.Author, book.Isbn)
	}

	// Create a new book
	log.Println("\n3. Creating a new book:")
	newBook := &pb.Book{
		Title:     "Design Patterns",
		Author:    "Gang of Four",
		Isbn:      "978-0201633610",
		Publisher: "Addison-Wesley",
		PageCount: 395,
		Stock:     2,
	}
	createResp, err := client.CreateBook(ctx, &pb.CreateBookRequest{Book: newBook})
	if err != nil {
		log.Printf("Error creating book: %v", err)
	} else {
		log.Printf("   Created book with ID: %s", createResp.Book.Id)
	}

	// Update the book
	log.Println("\n4. Updating book stock:")
	if createResp != nil {
		updateBook := createResp.Book
		updateBook.Stock = 5
		updateResp, err := client.UpdateBook(ctx, &pb.UpdateBookRequest{Book: updateBook})
		if err != nil {
			log.Printf("Error updating book: %v", err)
		} else {
			log.Printf("   Updated book stock to: %d", updateResp.Book.Stock)
		}
	}
}

func testStudentService(ctx context.Context, client pb.StudentServiceClient) {
	// List students
	log.Println("1. Listing all students:")
	listResp, err := client.ListStudents(ctx, &pb.ListStudentsRequest{})
	if err != nil {
		log.Printf("Error listing students: %v", err)
		return
	}
	for _, student := range listResp.Students {
		log.Printf("   - %s: %s (%s) - Active: %t", student.Id, student.Name, student.StudentNumber, student.IsActive)
	}

	// Get a specific student
	log.Println("\n2. Getting student-1:")
	getResp, err := client.GetStudent(ctx, &pb.GetStudentRequest{Id: "student-1"})
	if err != nil {
		log.Printf("Error getting student: %v", err)
	} else {
		student := getResp.Student
		log.Printf("   Found: %s - %s (%s)", student.Name, student.Email, student.StudentNumber)
	}

	// Create a new student
	log.Println("\n3. Creating a new student:")
	newStudent := &pb.Student{
		Name:          "Mehmet Demir",
		StudentNumber: "20210003",
		Email:         "mehmet.demir@university.edu",
		IsActive:      true,
	}
	createResp, err := client.CreateStudent(ctx, &pb.CreateStudentRequest{Student: newStudent})
	if err != nil {
		log.Printf("Error creating student: %v", err)
	} else {
		log.Printf("   Created student with ID: %s", createResp.Student.Id)
	}
}

func testLoanService(ctx context.Context, client pb.LoanServiceClient) {
	// List loans
	log.Println("1. Listing all loans:")
	listResp, err := client.ListLoans(ctx, &pb.ListLoansRequest{})
	if err != nil {
		log.Printf("Error listing loans: %v", err)
		return
	}
	if len(listResp.Loans) == 0 {
		log.Println("   No loans found")
	} else {
		for _, loan := range listResp.Loans {
			log.Printf("   - %s: Student %s borrowed Book %s on %s (Status: %s)",
				loan.Id, loan.StudentId, loan.BookId, loan.LoanDate, loan.Status.String())
		}
	}

	// Borrow a book
	log.Println("\n2. Student borrowing a book:")
	borrowResp, err := client.BorrowBook(ctx, &pb.BorrowBookRequest{
		StudentId: "student-1",
		BookId:    "book-1",
	})
	if err != nil {
		log.Printf("Error borrowing book: %v", err)
	} else {
		loan := borrowResp.Loan
		log.Printf("   Loan created: %s (Date: %s)", loan.Id, loan.LoanDate)
	}

	// List loans again to see the new loan
	log.Println("\n3. Listing loans after borrowing:")
	listResp2, err := client.ListLoans(ctx, &pb.ListLoansRequest{})
	if err != nil {
		log.Printf("Error listing loans: %v", err)
	} else {
		for _, loan := range listResp2.Loans {
			log.Printf("   - %s: Student %s, Book %s, Status: %s",
				loan.Id, loan.StudentId, loan.BookId, loan.Status.String())
		}
	}

	// Return a book (if we have a loan)
	if borrowResp != nil {
		log.Println("\n4. Returning the borrowed book:")
		returnResp, err := client.ReturnBook(ctx, &pb.ReturnBookRequest{
			LoanId: borrowResp.Loan.Id,
		})
		if err != nil {
			log.Printf("Error returning book: %v", err)
		} else {
			loan := returnResp.Loan
			log.Printf("   Book returned: Loan %s (Return Date: %s, Status: %s)",
				loan.Id, loan.ReturnDate, loan.Status.String())
		}
	}
} 