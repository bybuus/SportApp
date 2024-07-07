package auth

// SignInRequest use in sign up user
// we only sign in, we don't sign-up
type SignInRequest struct {
	OAuthToken string //header
}

// YandexOAuthResponse ya-auth
// returns from yandex-api
type YandexOAuthResponse struct {
	Login    string `json:"login"`
	Phone    string `json:"default_phone.number"`
	ClientId string `json:"client_id"`
}

type AuthResponse struct {
	Id    string `json:"id"` //save that on frontend
	Login string `json:"login"`
}
