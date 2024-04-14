package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var Token string

func RequestApi(apiURL, apiToken string, data interface{}) {
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+apiToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la requête:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}
}

func GetAccessToken() string {
	if Token == "" {
		ReloadApi()
	}
	return Token
}

func ReloadApi() {
	urltoken := "https://accounts.spotify.com/api/token"
	const clientId = "4188d3429ea04119a7db0a90ce7354cf"
	const clientSecret = "c92d845ebe0049fe9189a224fa62d837"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}
	bodyRequette := bytes.NewBufferString("grant_type=client_credentials&client_id=" + clientId + "&client_secret=" + clientSecret)
	req, err := http.NewRequest("POST", urltoken, bodyRequette)
	if err != nil {
		fmt.Println("error creating request: ", err.Error())
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, errRes := httpClient.Do(req)
	if resp.Body != nil {
		defer resp.Body.Close()
	} else {
		fmt.Println("Error creating response :", errRes.Error())
		os.Exit(2)
	}
	var reponseMap map[string]interface{}

	decoder := json.NewDecoder(resp.Body)
	errJSON := decoder.Decode(&reponseMap)
	if errJSON != nil {
		fmt.Println("Error reading JSON :", errJSON.Error())
		os.Exit(4)
	}

	Token = reponseMap["access_token"].(string)
}


