package types

type User struct {
	ID    int
	Name  string
	Email string
}

type ListUsersRequest struct {
	IDs []int `json:"ids"`
}

type ListUsersResponse struct {
	Users []User
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func toModel(r *CreateUserRequest) User {
	return User{
		Name:  r.Name,
		Email: r.Email,
	}
}

func toAPI(r *User) CreateUserRequest {
	return CreateUserRequest{
		Name:  r.Name,
		Email: r.Email,
	}
}

func Validate(r *CreateUserRequest) error {
	if r.Name == "" {
		return APIError{
			Status: 403,
			Msg:    "name is required",
		}
	}
	if r.Email == "" {
		return APIError{
			Status: 403,
			Msg:    "email is required",
		}
	}
	return nil
}
