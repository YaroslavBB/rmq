package repository

import "rabbitmq/rest_service/internal/entity/phone"

type PhoneRepository interface {
	FindPhoneNumber(numberID int) (phone.PhoneNumber, error)
}
