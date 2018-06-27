package handlers

import "zonstfe_api/common/my_context"


type Handler struct {
	*my_context.Context
}

func NewHandler(content *my_context.Context) *Handler {
	return &Handler{content}
}
