package client

import (
	"encoding/json"
	"fmt"
	"github.com/kendoow/SportApp/backend/internal/controllers/dto/auth"
	"io"
	"net/http"
)

const yaIdUrl = "https://login.yandex.ru/info"

type YaIdClient struct {
}

func (client *YaIdClient) Authorized(token string) *auth.YandexOAuthResponse {
	url := fmt.Sprintf("%s?format=json", yaIdUrl)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		//TODO error handle
	}

	oauthHeader := fmt.Sprintf("OAuth %s", token)
	req.Header.Add("Authorization", oauthHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//TODO err handle
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var userInfo auth.YandexOAuthResponse
	json.Unmarshal(body, &userInfo) //TODO err handle

	return &userInfo
}
