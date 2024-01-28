package initialize

import (
	"os"

	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"github.com/5asp/gin-vue-admin/server/model/city"
	"github.com/5asp/gin-vue-admin/server/model/order"

	"github.com/5asp/gin-vue-admin/server/model/carCombination"
	"github.com/5asp/gin-vue-admin/server/model/cityCarCombination"
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
		system.SysExportTemplate{}, city.CityData{}, order.Order{}, carCombination.CarCombination{}, cityCarCombination.CityCarCombination{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
