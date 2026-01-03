package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rahul-aithal/ThumbPicker/common/types"
	"github.com/Rahul-aithal/ThumbPicker/components"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *handler) VideoPage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) != 36 {
		fmt.Println(id, len(id))
		panic("Invalid Id")
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		log.Fatal(err, "While finding video")
		return
	}
	video, err := h.q.GetVideo(r.Context(), uid)

	if err != nil {
		log.Fatal(err, "While finding video")
		return
	}

	thumbs, err := h.q.GetAllThubmsOfVideo(r.Context(), uid)
	thumbDisplay := make([]types.ThumbData, 0)
	for i, thumb := range thumbs {
		thumbDisplay = append(thumbDisplay, types.ThumbData{
			Path:      thumb.Src[1:],
			TimeStamp: thumb.Timestamp,
			Index:     int(i),
			Id:        thumb.ID,
		})
	}

	components.VideoImage(video.Src,  thumbDisplay).Render(r.Context(), w)
}
