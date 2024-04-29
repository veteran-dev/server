package initialize

import (
	"github.com/smartwalle/alipay/v3"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/veteran-dev/server/global"
	"github.com/veteran-dev/server/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)

	global.GVA_AliPay, err = alipay.New(global.GVA_CONFIG.Alipay.AppID, global.GVA_CONFIG.Alipay.PrivateKey, false)

	if err != nil {
		panic(err)
	}

	// 加载证书
	if err = global.GVA_AliPay.LoadAppCertPublicKeyFromFile(global.GVA_CONFIG.Alipay.AppPublicCert); err != nil {
		panic(err)
	}
	if err = global.GVA_AliPay.LoadAliPayRootCertFromFile(global.GVA_CONFIG.Alipay.AlipayRootCert); err != nil {
		panic(err)

	}
	if err = global.GVA_AliPay.LoadAlipayCertPublicKeyFromFile(global.GVA_CONFIG.Alipay.AlipayPublicCert); err != nil {
		panic(err)
	}

}
