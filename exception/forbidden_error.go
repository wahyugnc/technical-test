package exception

type ForbiddenError struct {
	Message string
}

func (forbiddenError ForbiddenError) Error() string {
	return forbiddenError.Message
}
