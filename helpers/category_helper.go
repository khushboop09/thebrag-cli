package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"thebrag/requests"
	"thebrag/responses"
)

func GetAllCategories() (responses.GetAllCategoriesResponse, int) {
	apiUrl := fmt.Sprintf("%s/%s/categories", os.Getenv("API_HOST"), os.Getenv("USER_ID"))
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var categoriesResponse responses.GetAllCategoriesResponse
	json.Unmarshal(responseBody, &categoriesResponse)
	return categoriesResponse, response.StatusCode
}

func GetCategoryId(categoryName string) int {
	var categoryId int
	categoriesResp, statusCode := GetAllCategories()
	if statusCode != 200 {
		fmt.Println("Categories not found, please add to create a brag")
		return 0
	}
	categories := categoriesResp.Data
	for i := range categories {
		if strings.EqualFold(categories[i].Name, categoryName) {
			categoryId = categories[i].ID
		}
	}
	return categoryId
}

func AddACategory(name string) (responses.AddCategoryResponse, int) {
	category := requests.AddCategoryRequest{
		Name: name,
	}
	json_data, err := json.Marshal(category)
	if err != nil {
		fmt.Println(err)
	}
	api_url := fmt.Sprintf("%s/%s/category", os.Getenv("API_HOST"), os.Getenv("USER_ID"))
	response, err := http.Post(api_url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var res responses.AddCategoryResponse
	json.Unmarshal(responseBody, &res)
	return res, response.StatusCode
}
