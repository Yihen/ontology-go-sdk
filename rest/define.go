/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */
package rest

import (
	"encoding/json"
)

const (
	GET_GEN_BLK_TIME      = "/api/v1/node/generateblocktime"
	GET_CONN_COUNT        = "/api/v1/node/connectioncount"
	GET_BLK_TXS_BY_HEIGHT = "/api/v1/block/transactions/height/"
	GET_BLK_BY_HEIGHT     = "/api/v1/block/details/height/"
	GET_BLK_BY_HASH       = "/api/v1/block/details/hash/"
	GET_BLK_HEIGHT        = "/api/v1/block/height"
	GET_BLK_HASH          = "/api/v1/block/hash/"
	GET_TX                = "/api/v1/transaction/"
	GET_STORAGE           = "/api/v1/storage/"
	GET_BALANCE           = "/api/v1/balance/"
	GET_CONTRACT_STATE    = "/api/v1/contract/"
	GET_SMTCOCE_EVT_TXS   = "/api/v1/smartcode/event/transactions/"
	GET_SMTCOCE_EVTS      = "/api/v1/smartcode/event/txhash/"
	GET_BLK_HGT_BY_TXHASH = "/api/v1/block/height/txhash/"
	GET_MERKLE_PROOF      = "/api/v1/merkleproof/"
	GET_GAS_PRICE         = "/api/v1/gasprice"
	GET_ALLOWANCE         = "/api/v1/allowance/"
	GET_UNBOUNDONG        = "/api/v1/unboundong/"
	GET_MEMPOOL_TXCOUNT   = "/api/v1/mempool/txcount"
	GET_MEMPOOL_TXSTATE   = "/api/v1/mempool/txstate/"
	GET_VERSION           = "/api/v1/version"
	GET_NETWORK_ID        = "/api/v1/networkid"
	POST_RAW_TX           = "/api/v1/transaction"
)

const (
	ACTION_SEND_RAW_TRANSACTION = "sendrawtransaction"
)

const REST_VERSION = "1.0.0"

type RestfulReq struct {
	Action  string
	Version string
	Type    int
	Data    string
}

type RestfulResp struct {
	Action  string          `json:"action"`
	Result  json.RawMessage `json:"result"`
	Error   int64           `json:"error"`
	Desc    string          `json:"desc"`
	Version string          `json:"version"`
}
