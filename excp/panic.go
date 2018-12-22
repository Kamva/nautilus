package excp

// PanicIfErr panic exception if the given error is not nil
func PanicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// PanicIfErrCustomMsg panic with custom message if given error is not nil
func PanicIfErrCustomMsg(err error, message string) {
	if err != nil {
		panic(message)
	}
}
