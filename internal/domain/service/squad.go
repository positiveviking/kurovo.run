package service

import "damihailov85/kurovo.run/internal/domain/entity"

type Squad interface {
	Boost(from, to entity.SquadAccount) error
}
