//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import "fmt"
import "time"

type Title string // Book's title
type Name string  // Member's name

type BookEntry struct {
	total  int // Total of books owned by the library
	lended int // Total of books lended out
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Library struct {
	books   map[Title]BookEntry
	members map[Name]Member
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnedTime string

		if audit.checkIn.IsZero() {
			returnedTime = "[Not returned yet]"
		} else {
			returnedTime = audit.checkIn.String()
		}

		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnedTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println("---", title, "---")
		fmt.Println("Total:", book.total)
		fmt.Println("Lended:", book.lended)
		fmt.Println()
	}
}

func checkOutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("The book", title, "is not part of the library")
		return false
	}

	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}

	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}

	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("The book", title, "is not part of the library")
		return false
	}

	audit, found := member.books[title]

	if !found {
		fmt.Println("The member", member.name, "have not yet lend the book", title)
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit

	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	// Adding books
	library.books["Astrophysics I"] = BookEntry{
		total: 4,
	}
	library.books["Effective Go"] = BookEntry{
		total: 3,
	}
	library.books["Native Go"] = BookEntry{
		total: 8,
	}
	library.books["Atomic Habits"] = BookEntry{
		total: 1,
	}

	// Adding members
	library.members["Jayson"] = Member{"Jayson", make(map[Title]LendAudit)}
	library.members["Alisson"] = Member{"Alisson", make(map[Title]LendAudit)}
	library.members["Patrice"] = Member{"Patrice", make(map[Title]LendAudit)}

	fmt.Println("\nInitial library status:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	// Lending a book
	member := library.members["Patrice"]
	checkingOutBook := checkOutBook(&library, "Atomic Habits", &member)
	fmt.Println("check out a book...")
	if checkingOutBook {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

  // Returning a book
  returningBook := returnBook(&library, "Atomic Habits", &member)
  fmt.Println("Returning a book...")
  if returningBook {
		printLibraryBooks(&library)
		printMemberAudits(&library)
  }

}
