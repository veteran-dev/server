package initialize

import (
	"os"

	"github.com/veteran-dev/server/global"
	"github.com/veteran-dev/server/model/system"

	"github.com/veteran-dev/server/model/city"
	"github.com/veteran-dev/server/model/order"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/veteran-dev/server/model/carCombination"
	"github.com/veteran-dev/server/model/cityCarCombination"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysExportTemplate{}, city.City{}, order.Order{}, carCombination.CarCombination{}, cityCarCombination.CityCarCombination{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
