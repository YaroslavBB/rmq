package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rabbitmq/trigger_listener/internal/entity"
	"rabbitmq/trigger_listener/internal/entity/global"
)

type RestRepository struct{}

func NewRestRepository() RestRepository {
	return RestRepository{}
}

func (r *RestRepository) FindPhoneNumber(numberID int) (entity.PhoneNumber, error) {
	var data entity.PhoneNumber
	postBody, _ := json.Marshal(map[string]interface{}{
		"request_id": numberID,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://127.0.0.1:8080/get_number", "application/json", responseBody)
	if err != nil {
		fmt.Println(1)
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return data, err
		}

		if err := json.Unmarshal(body, &data); err != nil {
			return data, err
		}
		return data, nil

	case http.StatusNotFound:
		return data, global.ErrNoData

	default:
		return data, global.ErrInternalError
	}
}
