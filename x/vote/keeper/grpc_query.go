package keeper

import (
	"github.com/cosmonaut/vote/x/vote/types"
)

var _ types.QueryServer = Keeper{}
