package user

type Handler struct{}

func NewHandler() *Handler{
	return &Handler{} //return empty handler object
}