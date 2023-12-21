package users

type FindByIDRequest struct {
	ID int64 `json:"id" param:"id" query:"id"`
}
