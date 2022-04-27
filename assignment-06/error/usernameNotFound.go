package error

type UsernameNotFound struct {
	Message string `json:"errorMessage"`
}


func (u *UsernameNotFound) Error() string {
	return u.Message
}

func GetUsernameNotFound() *UsernameNotFound {
	return &UsernameNotFound{Message: "username not found"}
}