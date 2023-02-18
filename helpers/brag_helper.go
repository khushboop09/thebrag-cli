package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"thebrag/requests"
	"thebrag/responses"
)

func AddABrag(title string, details string, categoryId int) (responses.PostBragResponse, int) {
	brag := requests.AddBragRequest{
		Title:      title,
		Details:    details,
		CategoryID: categoryId,
	}
	json_data, err := json.Marshal(brag)
	if err != nil {
		fmt.Println(err)
	}
	api_url := fmt.Sprintf("%s/%s/brag", os.Getenv("API_HOST"), os.Getenv("USER_ID"))
	response, err := http.Post(api_url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var res responses.PostBragResponse
	json.Unmarshal(responseBody, &res)
	return res, response.StatusCode
}

func GetABrag(id int) (responses.GetABragResponse, int) {
	apiUrl := fmt.Sprintf("%s/%s/brag/%d", os.Getenv("API_HOST"), os.Getenv("USER_ID"), id)
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
	apiUrl := fmt.Sprintf("%s/%s/brags?skip=%d&limit=%d", os.Getenv("API_HOST"), os.Getenv("USER_ID"), skip, limit)
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

func UpdateABrag(id int, title string, details string, categoryName string, categoryId int) (responses.PostBragResponse, int) {
	if categoryId == 0 {
		categoryId = GetCategoryId(categoryName)
		if categoryId == 0 {
			log.Fatal("Category not found, please check the spelling or add the category first")
		}
	}

	brag := requests.UpdateBragRequest{
		ID:         id,
		Title:      title,
		Details:    details,
		CategoryID: categoryId,
	}
	json_data, err := json.Marshal(brag)
	if err != nil {
		log.Fatal(err)
	}
	api_url := fmt.Sprintf("%s/%s/brag", os.Getenv("API_HOST"), os.Getenv("USER_ID"))
	response, err := http.Post(api_url, "application/json", bytes.NewBuffer(json_data))
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
	apiUrl := fmt.Sprintf("%s/%s/brag/%d", os.Getenv("API_HOST"), os.Getenv("USER_ID"), id)
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

func ExportBrags(from string, to string, category string) (responses.PostBragResponse, int) {
	categoryId := 0
	if category != "" {
		categoryId = GetCategoryId(category)
	}

	brags := requests.ExportBragsRequest{
		From:       from,
		To:         to,
		CategoryId: categoryId,
	}
	json_data, err := json.Marshal(brags)
	if err != nil {
		log.Fatal(err)
	}
	api_url := fmt.Sprintf("%s/%s/brags/export", os.Getenv("API_HOST"), os.Getenv("USER_ID"))
	response, err := http.Post(api_url, "application/json", bytes.NewBuffer(json_data))
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
