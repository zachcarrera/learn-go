package grep

type lineComparison func(string, string) bool

func Search(pattern string, flags, files []string) []string {
	panic("Please implement the Search function")
}

func hasFlag(flags []string, flag string) bool {
	for _, v := range flags {
		if v == flag {
			return true
		}
	}
	return false
}
