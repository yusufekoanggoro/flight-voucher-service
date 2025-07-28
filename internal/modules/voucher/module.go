package voucher

import (
	"database/sql"

	"github.com/yusufekoanggoro/flight-voucher-service/internal/factory/base"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/factory/interfaces"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/delivery/resthandler"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/repository"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/usecase"
)

type Module struct {
	restHandler interfaces.RestHandler
}

func NewModule(db *sql.DB) *Module {
	var module Module

	repo := repository.NewVoucherRepository(db)
	uc := usecase.NewVoucherUsecase(repo)
	restHandler := resthandler.NewRestHandler(uc)

	module.restHandler = restHandler
	return &module
}

func (m *Module) Name() base.ModuleType {
	return base.ModuleVoucher
}

func (m *Module) RestHandler() (d interfaces.RestHandler) {
	return m.restHandler
}
