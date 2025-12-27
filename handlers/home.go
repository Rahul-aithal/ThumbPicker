package handlers

import (
	"net/http"

	"github.com/Rahul-aithal/ThumbPicker/components"
)

func (h *handler) Home (w http.ResponseWriter, r *http.Request) {
	home := components.Home()
	home.Render(r.Context(), w)
}
