package anywhat

import "github.com/golang/protobuf/ptypes/timestamp"

type createAnythingRequest struct {
	ID          string               `json:"id"`
	Name        string               `json:"name" validate:"required,min=2,max=50"`
	Description string               `json:"description"`
	CreatedAt   *timestamp.Timestamp `json:"createdAt"`
	UpdatedAt   *timestamp.Timestamp `json:"updatedAt"`
}

type updateAnythingRequest struct {
	ID          string               `json:"id"`
	Name        string               `json:"name" validate:"required,min=2,max=50"`
	Description string               `json:"description"`
	CreatedAt   *timestamp.Timestamp `json:"createdAt"`
	UpdatedAt   *timestamp.Timestamp `json:"updatedAt"`
}
