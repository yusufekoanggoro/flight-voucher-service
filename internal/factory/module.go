package factory

import (
	"database/sql"

	"github.com/yusufekoanggoro/flight-voucher-service/internal/factory/interfaces"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher"
)

type ModuleFactory interface {
	RestHandler() interfaces.RestHandler
}

func InitAllModule(db *sql.DB) []ModuleFactory {

	modules := []ModuleFactory{
		voucher.NewModule(db),

		// Add initialization for other modules below
		// "modulename": modulePackage.NewModule(db),
	}

	return modules
}
