package prod

import "example/infrastructure/database"

type sqlGateway struct {
	database.GormWithTrxImpl
	database.GormWithoutTrxImpl
}
