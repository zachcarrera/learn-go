package isbn

func IsValidISBN(isbn string) bool {
	if len(isbn) == 10 || len(isbn) == 13 {
		return true
	}
	return false
}
