package users

type FindByIDRequest struct {
	ID int64 `json:"id" param:"id" query:"id"`
}

type FindByEmailRequest struct {
	Email string `json:"email" param:"email" query:"email"`
}
