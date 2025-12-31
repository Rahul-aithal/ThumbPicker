package routers

import (
	"github.com/Rahul-aithal/ThumbPicker/db"
	"github.com/Rahul-aithal/ThumbPicker/handlers"
	"github.com/go-chi/chi/v5"
)

// TODO: Take all the routers form different files put it all here an retturn one
// TODO: First list all the endpoints
// 1. Home ('/') return Home page HTML
// 2. Upload Video ('/video') Retrun list of images
// 3. Select Image and Download ('/image/{slug}') Return Video in inline Header

func Routers(q *db.Queries) *chi.Mux {
	r := chi.NewRouter()
	h := handlers.NewHandler(q)
	r.Get("/", h.Home)
	r.Get("/video/{id}", h.VideoPage)
	r.Post("/video", h.Video)
	// r.Post("/downloadVideo/")
	return r
}
