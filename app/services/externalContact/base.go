package externalContact

import (
	"errors"
	"workwx/app/models/qyWechatAccountModel"
	"workwx/pkg/logger"
)

type syncFunc func(account *qyWechatAccountModel.ScrmQyWechatAccount)

func accountSync(syf syncFunc) {
	accounts, err := qyWechatAccountModel.All()
	if err != nil {
		logger.LogError(err)
		panic(err)
	}
	if len(accounts) == 0 {
		err := errors.New("no account can use")
		logger.LogError(err)
		panic(err)
	}

	for _, account := range accounts {
		syf(account)
	}

}
