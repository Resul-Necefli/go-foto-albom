package model

import (
	"fmt"
)

type Photo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

var Photos = []Photo{
	{ID: 1, Title: "Beautiful Sun", URL: "https://picsum.photos/200/300"},
	{ID: 2, Title: "Seaside", URL: "https://picsum.photos/300/200"},
	{ID: 3, Title: "Mountain View", URL: "https://picsum.photos/400/300"},
}

func FindPhotoById(ID int) (*Photo, error) {

	for _, v := range Photos {

		if v.ID == ID {

			return &v, nil

		}
	}

	return nil, fmt.Errorf("photo not found")

}
