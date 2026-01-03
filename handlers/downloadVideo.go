package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Rahul-aithal/ThumbPicker/common"
	"github.com/Rahul-aithal/ThumbPicker/internals/video"
	"github.com/go-chi/chi/v5"
)

type RequestBody struct {
	Index string `json:"index"`
}

func (h *handler) HandleDownload(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	verifyFormValue(id, "No id found")
	thumb, err := h.q.GetThumb(r.Context(), common.GetUUID(id))

	if err != nil || len(thumb.Src) <= 0 {
		panic("No thumbnail found")
	}
	vi, err := h.q.GetVideo(r.Context(), thumb.Video)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	meta, err := video.Combainer(video.CombainerStruct{
		FilePath:      vi.Src,
		ThumbLocation: thumb.Src,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while Combaining", http.StatusInternalServerError)
	}

	file, err := os.Open(meta.FilePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set(
		"Content-Disposition",
		"attachment; filename=\""+meta.FileName,
	)
	w.Header().Set("Content-Type", "application/octet-stream")
	stat, err := file.Stat()
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
	_, err = io.CopyN(w, file, stat.Size())
	if err != nil {
		log.Printf("Error streaming file: %v", err)
	}
}

func verifyFormValue(input string, err string) {
	if len(input) <= 0 {
		panic(err)
	}
}
