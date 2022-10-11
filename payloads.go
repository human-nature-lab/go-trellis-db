package go_trellis_db

type LoginPayload struct {
	Username string
	Pass     string
}

type LoginResult struct {
	User  User
	Token Token
}
