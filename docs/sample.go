package docs

func do() {
	// ignore指定されているので、この部分はlintの指摘
	returnError()
}

func returnError() error {
	return nil
}
