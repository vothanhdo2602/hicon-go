package service

import (
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/internal/rd"
)

func isDisableCache(localCache bool) bool {
	var (
		env = config.GetENV()
	)

	if !env.DB.DBConfiguration.DisableCache || rd.GetRedis() == nil {
		return true
	}

	return localCache
}
