package google

type LoginError struct {
	message string
}

func (e *LoginError) Error() string {
	return "auth/google: " + e.message
}
