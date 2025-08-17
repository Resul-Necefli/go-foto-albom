package storage

import (
	"Resul-Necefli/go-foto-albom/model"
	"errors"
	"fmt"
)

func FindPhotoById(ID int) (*model.Photo, error) {

	for _, v := range model.Photos {

		if v.ID == ID {

			return &v, nil

		}

	}

	return nil, fmt.Errorf("photo not found")

}

func NewPhotoCreate(photo *model.Photo) error {

	if photo == nil {

		return errors.New("zero value is  referencd")

	}

	for _, v := range model.Photos {

		if v.ID == photo.ID {

			return errors.New("the enter ID already exists")
		}

	}
	model.Photos = append(model.Photos, *photo)
	return nil

}
