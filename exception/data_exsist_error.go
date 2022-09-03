package exception

type DataExsistError struct {
	Message string
}

func (dataExsistError DataExsistError) Error() string {
	return dataExsistError.Message
}
