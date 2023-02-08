package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"thebrag/requests"
	"thebrag/responses"
)

func GetABrag(id int) (responses.GetABragResponse, int) {
	apiUrl := fmt.Sprintf("http://localhost:8080/brag/%d", id)
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var bragResponse responses.GetABragResponse
	json.Unmarshal(responseBody, &bragResponse)
	return bragResponse, response.StatusCode
}

func GetAllBrags(skip int, limit int) (responses.GetAllBragResponse, int) {
	apiUrl := fmt.Sprintf("http://localhost:8080/brags?skip=%d&limit=%d", skip, limit)
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var bragResponse responses.GetAllBragResponse
	json.Unmarshal(responseBody, &bragResponse)
	return bragResponse, response.StatusCode
}

func UpdateABrag(id int, title string, details string) (responses.PostBragResponse, int) {
	brag := requests.UpdateBragRequest{
		ID:      id,
		Title:   title,
		Details: details,
	}
	json_data, err := json.Marshal(brag)
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.Post("http://localhost:8080/brag", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var res responses.PostBragResponse
	json.Unmarshal(responseBody, &res)
	return res, response.StatusCode
}

func DeleteABrag(id int) responses.DeleteBragResponse {
	apiUrl := fmt.Sprintf("http://localhost:8080/brag/%d", id)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", apiUrl, nil)
	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var bragResponse responses.DeleteBragResponse
	json.Unmarshal(responseBody, &bragResponse)
	return bragResponse
}
