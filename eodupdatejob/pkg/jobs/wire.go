//+build wireinject

package jobs

import (
	"github.com/google/wire"
	"github.com/pwestlake/equity-fund/eodupdatejob/pkg/service"
	commons "github.com/pwestlake/equity-fund/commons/pkg/service"
	"github.com/pwestlake/equity-fund/commons/pkg/dao"
)

func InitializeBackFillJob() BackFillJob {
	wire.Build(NewBackFillJob, service.NewMarketStackService, 
		commons.NewEquityCatalogService, dao.NewEquityCatalogItemDAO,
		commons.NewEndOfDayService, dao.NewEndOfDayItemDAO)
	return BackFillJob{}
}