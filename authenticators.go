package auth

// StandardAuthenticator interface uses email and password authentication
type StandardAuthenticator interface {
	Login(email, password string) (userID uint, err error)
	SignUp(email, password string) (userID uint, err error)
	Logout()
}

// Authenticator interface combines all authenticator interfaces
type Authenticator interface {
	StandardAuthenticator
}
