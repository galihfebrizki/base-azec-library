package baseModel

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type CredentialsResponse struct {
	Token string `json:"token"`
}
