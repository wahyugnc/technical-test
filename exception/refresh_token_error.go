package exception

type RefreshTokenError struct {
	Message string
}

type RefreshTokenIOSError struct {
	Message string
}

func (refreshTokenError RefreshTokenError) Error() string {
	return refreshTokenError.Message
}

func (refreshTokenIOSError RefreshTokenIOSError) Error() string {
	return refreshTokenIOSError.Message
}
