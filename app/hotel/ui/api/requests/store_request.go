package requests

import (
	"gin2018/app/container/ui/api/requests"
)

type StoreRequest struct {
	*requests.FormRequest
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
