package conn

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGrpcClient)
