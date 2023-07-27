package tournament

import "io"

type teamRecord struct {
	name   string
	played int
	wins   int
	losses int
	draws  int
}

func Tally(reader io.Reader, writer io.Writer) error {
	panic("Please implement the Tally function")
}
