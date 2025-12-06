package video

import "fmt"

func Service(path string, options int) {
	metap := &Meta{}
	metap.FilePath = path
	metap.NumberOfFrames = options
	err := MetaData(metap)
	if err != nil {
		panic("Error while getting meta data")
	}
	err1 := Generator(metap)
	if err1 != nil {
		panic("Error while Generator")
	}
	err2 := Extractor(metap)
	if err2 != nil {
		fmt.Println(err2)
		panic("Error while Extractor")
	}
	metap.SeletedThumbIndex = 1
	err3 := Combainer(metap)
	if err3 != nil {
		fmt.Println(err3)
		panic("Error while Combaining")
	}

	fmt.Println("The location of thumbnail mentioned: ", metap.ThumbLocation)
}
