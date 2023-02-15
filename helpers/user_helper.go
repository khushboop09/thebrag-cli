package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"thebrag/requests"
	"thebrag/responses"
)

func LoginUser(name string, email string, password string) (responses.PostUserResponse, int) {
	user := requests.CreateUserRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}
	json_data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	api_url := fmt.Sprintf("%s/user", os.Getenv("API_HOST"))
	response, err := http.Post(api_url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var res responses.PostUserResponse
	json.Unmarshal(responseBody, &res)
	return res, response.StatusCode
}

func SaveUserIdToDataFile(user_id int, task string) {
	if task == "create" {
		file, err := os.Create(".data")
		if err != nil {
			fmt.Printf("failed login: %s", err)
			return
		}
		_, fileErr := file.WriteString(fmt.Sprintf("USER_ID=%d\n", user_id))
		if fileErr != nil {
			fmt.Printf("failed writing to file: %s", fileErr)
			return
		}
	} else {
		f, err := os.OpenFile(".data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if _, err := f.Write([]byte(fmt.Sprintf("USER_ID=%d\n", user_id))); err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := f.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}
}
