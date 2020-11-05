package anywhat

type createAnythingRequest struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}

type updateAnythingRequest struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}
