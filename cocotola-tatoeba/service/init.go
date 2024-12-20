package service

import (
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
)

type TransactionManager rslibgateway.TransactionManagerT[RepositoryFactory]
