package plerrors

// MockBadInputError This method should be called only in test cases
func MockBadInputError(code string) (pErr *AppError) {
	return &AppError{
		StatusCode: BadInput,
		Code:       code,
		Message:    "mock internal service error for testing flow",
	}
}
