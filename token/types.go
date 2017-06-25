package token

type TokeResponse struct {
	AccsessToken string `json:"access_token"`
	TokeType     string `json:"token_type"`
	ExpiresIn    int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
