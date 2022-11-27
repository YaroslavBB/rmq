package postgres

import (
	"database/sql"
	"rabbitmq/rest_service/internal/entity/global"
	"rabbitmq/rest_service/internal/entity/phone"
	"rabbitmq/rest_service/internal/repository"

	"github.com/jmoiron/sqlx"
)

type phoneRepository struct {
	*sqlx.DB
}

func NewNumberRepository(db *sqlx.DB) repository.PhoneRepository {
	return &phoneRepository{DB: db}
}

func (r *phoneRepository) FindPhoneNumber(numberID int) (phone.PhoneNumber, error) {
	var data phone.PhoneNumber
	err := r.DB.Get(&data, `
	select number_id, number 
	from phone_numbers 
	where number_id = $1`, numberID)
	switch err {
	case nil:
		return data, err
	case sql.ErrNoRows:
		return data, global.ErrNoData
	default:
		return data, err
	}
}
