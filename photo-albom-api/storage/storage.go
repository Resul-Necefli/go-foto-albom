package storage

import (
	"Resul-Necefli/go-foto-albom/model"
	"fmt"
)

func GetByIDPhoto(id int) (model.Photo, error) {

	if v, ok := model.Photos[id]; ok {

		return v, nil

	}

	return model.Photo{}, fmt.Errorf("ID %d already used", id)

}

func AddPhoto(p model.Photo) error {

	if _, ok := model.Photos[p.ID]; ok {

		return fmt.Errorf("ID %d already used", p.ID)
	}

	model.Photos[p.ID] = p
	return nil

}

func UpdatePhoto(p model.Photo) {

	model.Photos[p.ID] = p
}

func DeletePhoto(p model.Photo) {

	delete(model.Photos, p.ID)
}

func GetPhoto(id int) (*model.Photo, error) {

	photoObj, err := GetByIDPhoto(id)

	if err != nil {

		return &model.Photo{}, fmt.Errorf("GetPhoto to : %v ", err)

	}

	return &photoObj, nil

}
