package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource
	for resource, err = opener(); err != nil; resource, err = opener() {
		switch err.(type) {
		case TransientError:
			continue
		default:
			return err
		}
	}
	defer resource.Close()

	defer func() {
		if e := recover(); e != nil {
			if f, ok := e.(FrobError); ok {
				resource.Defrob(f.defrobTag)
			}
			err = e.(error)
		}
	}()

	resource.Frob(input)
	return err
}
