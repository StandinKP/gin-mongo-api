package responses

import "gin-mongo-api/models"

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
	User    models.User            `json:"user"`
}
