package bearertoken

// BearerToken is used for single-application developer accounts
type Token struct {
	AccessToken string `json:"access_token"`
}
