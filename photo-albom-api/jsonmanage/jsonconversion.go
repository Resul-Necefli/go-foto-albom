package jsonmanage

import (
	"Resul-Necefli/go-foto-albom/model"
	"encoding/json"
	"fmt"
)

func Jsonconversion() ([]byte, error) {

	data, err := json.MarshalIndent(model.Photos, "", "  ")

	if err != nil {
		return nil, fmt.Errorf("Conversion to json failed %v", err)
	}

	return data, nil

}
