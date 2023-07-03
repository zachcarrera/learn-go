package isbn

import (
	"fmt"
	"strings"
)

func IsValidISBN(isbn string) bool {
	if !(len(isbn) == 10 || len(isbn) == 13) {
		return false
	}
	fmt.Printf("isbn: %v\n", isbn)
	groups := strings.Split(isbn, "-")
	fmt.Printf("groups: %v\n", groups)
	isbn = strings.Join(groups, "")
	fmt.Printf("isbn: %v\n", isbn)
	return true
}
