package types

type APIError struct {
	Status int
	Msg    string
}

func (e APIError) Error() string {
	return e.Msg
}
