package authentication

type UserFetcher interface {
	getUserByUsername(string) (string, error) // TODO return a user type
}

type AuthStore struct {
	userFetcher UserFetcher
}
