package authentication

import "github.com/RisingStack/almandite-user-service/internal/dal"

// AuthStore encapsulates repository needed for authentication middlewares
type AuthStore struct {
	UserRepository dal.UserRepository
}
