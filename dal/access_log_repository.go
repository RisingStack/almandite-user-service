package dal

import (
	"log"

	"github.com/RisingStack/almandite-user-service/models"
	"github.com/go-pg/pg"
)

// AccessLogRepository interface
type AccessLogRepository interface {
	Fetch() (*[]models.AccessLog, error)
	GetByUserID(userID int) (*[]models.AccessLog, error)
	Create(accessLog *models.AccessLog) error
}

type accessLogRepository struct {
	DB *pg.DB
}

// NewUserRepository returns a repository that implements the UserRepository interface
func newAccessLogRepository(dbConn *pg.DB) AccessLogRepository {
	return &accessLogRepository{
		DB: dbConn,
	}
}

func (a *accessLogRepository) Fetch() (*[]models.AccessLog, error) {
	var accessLogs []models.AccessLog

	if err := a.DB.Model(&accessLogs).Select(); err != nil {
		return nil, err
	}

	log.Println("ACCESSLOGS", accessLogs)

	return &accessLogs, nil
}

func (a *accessLogRepository) GetByUserID(userID int) (*[]models.AccessLog, error) {
	var accessLogs []models.AccessLog

	err := a.DB.Model(&accessLogs).
		Where("user_id = ?", userID).
		Select()

	if err != nil {
		return nil, err
	}

	return &accessLogs, nil
}

func (a *accessLogRepository) Create(accessLog *models.AccessLog) error {
	log.Println("CREATE", accessLog)
	return a.DB.Insert(accessLog)
}
