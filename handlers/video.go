package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Rahul-aithal/ThumbPicker/db"
	"github.com/Rahul-aithal/ThumbPicker/internals/video"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *handler) Video(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "File too large or invalid form data", http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile("file")

	fileName := filepath.Base(header.Filename)
	if fileName == "" {
		http.Error(w, "Invalid filename", http.StatusBadRequest)
		return
	}

	uploads := "./pub/uploads"
	os.MkdirAll(uploads, os.ModePerm)
	filePath := filepath.Join(uploads, fileName)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Upload successful: %s (Size: %d bytes)", fileName, header.Size)

	fmt.Printf("Starting to Processs the video: %s", fileName)
	result := video.Service(filePath, 8)
	video,err:=h.q.CreateVideo(r.Context(),db.CreateVideoParams{
		Src:  result.FilePath,
		Dur: strconv.FormatFloat(result.Duration,'f',-1,64) ,
		ThumbnailsCount:  int32(result.NumberOfFrames),
	})
	if err!=nil{
		log.Fatal(err,"While Inserting  video")
	}
	for i := range result.NumberOfFrames {
		h.q.InsertThumbs(r.Context(),db.InsertThumbsParams{
			Src:  result.ThumbLocation[i],
			Video: video.ID,
			Timestamp: strconv.FormatFloat(result.TimeStamps[i],'f',-1,64) ,
			Idx: pgtype.Int4{Int32: int32(i),Valid: true},
	 })

	}
	http.Redirect(w, r, "/video/"+video.ID.String(), http.StatusFound)
	// components.VideoImage(result.FilePath, thumbDisplay).Render(r.Context(), w)
}
// thumbDisplay := make([]types.ThumbData, 0)
// thumbDisplay = append(thumbDisplay, types.ThumbData{
// 			Path:      result.ThumbLocation[i],
// 			TimeStamp: result.TimeStamps[i],
// 			Index:     int(i),
// 		})
