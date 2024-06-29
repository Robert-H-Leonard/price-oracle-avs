package common

import (
	_ "embed"
)

//go:embed abis/PriceAggregatorTaskManager.json
var PriceAggregatorTaskManagerAbi []byte
