package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gogetsms-api/model"
)

func GetRentList() (*model.RentListResponse, error) {
	apiKey := os.Getenv("GOGETSMS_API_KEY")
	url := fmt.Sprintf("https://www.gogetsms.com/handler_api.php?api_key=%s&action=getRentList", apiKey)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	
	var result model.RentListResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
