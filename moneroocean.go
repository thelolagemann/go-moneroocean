package go_moneroocean

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

var (
	client Client = &http.Client{}
)

// StatsResponse is the struct returned by the Stats function.
type StatsResponse struct {
	Hash             float32 `json:"hash"`
	Hash2            float32 `json:"hash2"`
	ID               string  `json:"identifier"`
	LastHash         float32 `json:"lastHash"`
	TotalHashes      float32 `json:"totalHashes"`
	ValidShares      uint    `json:"validShares"`
	InvalidShares    uint    `json:"invalidShares"`
	AmountPaid       uint    `json:"amtPaid"`
	AmountDue        uint    `json:"amtDue"`
	TransactionCount uint    `json:"txnCount"`
}

// PoolStatsResponse is the struct returned by the PoolStats function.
type PoolStatsResponse struct {
	PoolList       []string `json:"pool_list"`
	PoolStatistics struct {
		HashRate            float64 `json:"hashRate"`
		Miners              int     `json:"miners"`
		TotalHashes         int64   `json:"totalHashes"`
		LastBlockFoundTime  int     `json:"lastBlockFoundTime"`
		LastBlockFound      int     `json:"lastBlockFound"`
		TotalBlocksFound    int     `json:"totalBlocksFound"`
		TotalMinersPaid     int     `json:"totalMinersPaid"`
		TotalPayments       int     `json:"totalPayments"`
		RoundHashes         float64 `json:"roundHashes"`
		TotalAltBlocksFound int     `json:"totalAltBlocksFound"`
		AltBlocksFound      struct {
			Num2086  int `json:"2086"`
			Num8545  int `json:"8545"`
			Num8766  int `json:"8766"`
			Num9053  int `json:"9053"`
			Num9231  int `json:"9231"`
			Num9998  int `json:"9998"`
			Num11181 int `json:"11181"`
			Num11812 int `json:"11812"`
			Num11898 int `json:"11898"`
			Num12211 int `json:"12211"`
			Num13007 int `json:"13007"`
			Num13102 int `json:"13102"`
			Num16000 int `json:"16000"`
			Num17750 int `json:"17750"`
			Num18181 int `json:"18181"`
			Num18981 int `json:"18981"`
			Num19091 int `json:"19091"`
			Num19281 int `json:"19281"`
			Num19734 int `json:"19734"`
			Num19950 int `json:"19950"`
			Num19994 int `json:"19994"`
			Num20189 int `json:"20189"`
			Num20206 int `json:"20206"`
			Num22023 int `json:"22023"`
			Num24182 int `json:"24182"`
			Num25182 int `json:"25182"`
			Num26968 int `json:"26968"`
			Num31014 int `json:"31014"`
			Num33124 int `json:"33124"`
			Num34568 int `json:"34568"`
			Num38081 int `json:"38081"`
			Num48782 int `json:"48782"`
		} `json:"altBlocksFound"`
		ActivePort       int           `json:"activePort"`
		ActivePorts      []interface{} `json:"activePorts"`
		ActivePortProfit float64       `json:"activePortProfit"`
		CoinProfit       struct {
			Num2086  float64     `json:"2086"`
			Num8545  float64     `json:"8545"`
			Num8766  float64     `json:"8766"`
			Num9053  float64     `json:"9053"`
			Num9231  float64     `json:"9231"`
			Num9998  float64     `json:"9998"`
			Num11812 float64     `json:"11812"`
			Num11898 float64     `json:"11898"`
			Num12211 float64     `json:"12211"`
			Num13007 float64     `json:"13007"`
			Num16000 float64     `json:"16000"`
			Num17750 float64     `json:"17750"`
			Num18081 float64     `json:"18081"`
			Num18981 interface{} `json:"18981"`
			Num19281 float64     `json:"19281"`
			Num19734 float64     `json:"19734"`
			Num19950 float64     `json:"19950"`
			Num19994 float64     `json:"19994"`
			Num20206 float64     `json:"20206"`
			Num25182 float64     `json:"25182"`
			Num33124 float64     `json:"33124"`
			Num38081 float64     `json:"38081"`
			Num48782 float64     `json:"48782"`
		} `json:"coinProfit"`
		CoinComment struct {
			Num2086  string `json:"2086"`
			Num8545  string `json:"8545"`
			Num8766  string `json:"8766"`
			Num9053  string `json:"9053"`
			Num9231  string `json:"9231"`
			Num9998  string `json:"9998"`
			Num11812 string `json:"11812"`
			Num11898 string `json:"11898"`
			Num12211 string `json:"12211"`
			Num13007 string `json:"13007"`
			Num16000 string `json:"16000"`
			Num17750 string `json:"17750"`
			Num18081 string `json:"18081"`
			Num18981 string `json:"18981"`
			Num19281 string `json:"19281"`
			Num19734 string `json:"19734"`
			Num19950 string `json:"19950"`
			Num19994 string `json:"19994"`
			Num20206 string `json:"20206"`
			Num25182 string `json:"25182"`
			Num33124 string `json:"33124"`
			Num38081 string `json:"38081"`
			Num48782 string `json:"48782"`
		} `json:"coinComment"`
		MinBlockRewards struct {
			Num2086  float64 `json:"2086"`
			Num8545  float64 `json:"8545"`
			Num8766  float64 `json:"8766"`
			Num9053  float64 `json:"9053"`
			Num9231  float64 `json:"9231"`
			Num9998  float64 `json:"9998"`
			Num11812 float64 `json:"11812"`
			Num11898 float64 `json:"11898"`
			Num12211 float64 `json:"12211"`
			Num13007 float64 `json:"13007"`
			Num16000 float64 `json:"16000"`
			Num17750 float64 `json:"17750"`
			Num18081 float64 `json:"18081"`
			Num19281 float64 `json:"19281"`
			Num19734 float64 `json:"19734"`
			Num19950 float64 `json:"19950"`
			Num19994 float64 `json:"19994"`
			Num20206 float64 `json:"20206"`
			Num25182 float64 `json:"25182"`
			Num33124 float64 `json:"33124"`
			Num38081 float64 `json:"38081"`
			Num48782 float64 `json:"48782"`
		} `json:"minBlockRewards"`
		Pending float64 `json:"pending"`
		Price   struct {
			Btc float64 `json:"btc"`
			Usd float64 `json:"usd"`
			Eur float64 `json:"eur"`
		} `json:"price"`
		CurrentEfforts struct {
			Num2086  float64 `json:"2086"`
			Num8545  int64   `json:"8545"`
			Num8766  float64 `json:"8766"`
			Num9053  int64   `json:"9053"`
			Num9231  float64 `json:"9231"`
			Num9998  float64 `json:"9998"`
			Num11812 float64 `json:"11812"`
			Num11898 float64 `json:"11898"`
			Num12211 float64 `json:"12211"`
			Num13007 float64 `json:"13007"`
			Num16000 float64 `json:"16000"`
			Num17750 float64 `json:"17750"`
			Num18081 float64 `json:"18081"`
			Num18981 float64 `json:"18981"`
			Num19281 float64 `json:"19281"`
			Num19734 float64 `json:"19734"`
			Num19950 float64 `json:"19950"`
			Num19994 float64 `json:"19994"`
			Num20206 float64 `json:"20206"`
			Num25182 float64 `json:"25182"`
			Num33124 float64 `json:"33124"`
			Num38081 float64 `json:"38081"`
			Num48782 float64 `json:"48782"`
		} `json:"currentEfforts"`
		PplnsPortShares struct {
			Num8545  float64 `json:"8545"`
			Num8766  float64 `json:"8766"`
			Num9053  float64 `json:"9053"`
			Num9231  float64 `json:"9231"`
			Num9998  float64 `json:"9998"`
			Num11812 float64 `json:"11812"`
			Num11898 float64 `json:"11898"`
			Num12211 float64 `json:"12211"`
			Num13007 float64 `json:"13007"`
			Num16000 float64 `json:"16000"`
			Num17750 float64 `json:"17750"`
			Num18081 float64 `json:"18081"`
			Num19734 float64 `json:"19734"`
			Num19950 float64 `json:"19950"`
			Num19994 float64 `json:"19994"`
			Num20206 float64 `json:"20206"`
			Num25182 float64 `json:"25182"`
			Num33124 float64 `json:"33124"`
			Num38081 float64 `json:"38081"`
			Num48782 float64 `json:"48782"`
		} `json:"pplnsPortShares"`
		PplnsWindowTime float64 `json:"pplnsWindowTime"`
		PortHash        struct {
			Num8545  float64 `json:"8545"`
			Num8766  float64 `json:"8766"`
			Num9231  float64 `json:"9231"`
			Num9998  float64 `json:"9998"`
			Num11812 float64 `json:"11812"`
			Num11898 float64 `json:"11898"`
			Num12211 float64 `json:"12211"`
			Num13007 float64 `json:"13007"`
			Num16000 float64 `json:"16000"`
			Num17750 float64 `json:"17750"`
			Num18081 float64 `json:"18081"`
			Num19734 float64 `json:"19734"`
			Num19950 float64 `json:"19950"`
			Num19994 float64 `json:"19994"`
			Num20206 float64 `json:"20206"`
			Num25182 float64 `json:"25182"`
			Num33124 float64 `json:"33124"`
			Num38081 float64 `json:"38081"`
			Num48782 float64 `json:"48782"`
		} `json:"portHash"`
		PortMinerCount struct {
			Num8545  int `json:"8545"`
			Num8766  int `json:"8766"`
			Num9231  int `json:"9231"`
			Num9998  int `json:"9998"`
			Num11812 int `json:"11812"`
			Num11898 int `json:"11898"`
			Num12211 int `json:"12211"`
			Num13007 int `json:"13007"`
			Num16000 int `json:"16000"`
			Num17750 int `json:"17750"`
			Num18081 int `json:"18081"`
			Num19734 int `json:"19734"`
			Num19950 int `json:"19950"`
			Num19994 int `json:"19994"`
			Num20206 int `json:"20206"`
			Num25182 int `json:"25182"`
			Num38081 int `json:"38081"`
		} `json:"portMinerCount"`
		PortCoinAlgo struct {
			Num2086  string `json:"2086"`
			Num8545  string `json:"8545"`
			Num8766  string `json:"8766"`
			Num9053  string `json:"9053"`
			Num9231  string `json:"9231"`
			Num9998  string `json:"9998"`
			Num11812 string `json:"11812"`
			Num11898 string `json:"11898"`
			Num12211 string `json:"12211"`
			Num13007 string `json:"13007"`
			Num16000 string `json:"16000"`
			Num17750 string `json:"17750"`
			Num18081 string `json:"18081"`
			Num18981 string `json:"18981"`
			Num19281 string `json:"19281"`
			Num19734 string `json:"19734"`
			Num19950 string `json:"19950"`
			Num19994 string `json:"19994"`
			Num20206 string `json:"20206"`
			Num25182 string `json:"25182"`
			Num33124 string `json:"33124"`
			Num38081 string `json:"38081"`
			Num48782 string `json:"48782"`
		} `json:"portCoinAlgo"`
	} `json:"pool_statistics"`
	LastPayment int `json:"last_payment"`
}

// NetworkStatsResponse is the struct returned by the NetworkStats function.
type NetworkStatsResponse struct {
	Num2086 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int    `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"2086"`
	Num8545 struct {
		Difficulty int64  `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Ts         int    `json:"ts"`
	} `json:"8545"`
	Num8766 struct {
		Difficulty int         `json:"difficulty"`
		Hash       string      `json:"hash"`
		Height     int         `json:"height"`
		Value      int64       `json:"value"`
		Ts         interface{} `json:"ts"`
	} `json:"8766"`
	Num9053 struct {
		Difficulty int64  `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Ts         int    `json:"ts"`
	} `json:"9053"`
	Num9231 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int    `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"9231"`
	Num9998 struct {
		Difficulty int64       `json:"difficulty"`
		Hash       string      `json:"hash"`
		Height     int         `json:"height"`
		Value      int64       `json:"value"`
		Ts         interface{} `json:"ts"`
	} `json:"9998"`
	Num11812 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int    `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"11812"`
	Num11898 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int    `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"11898"`
	Num12211 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"12211"`
	Num13007 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int    `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"13007"`
	Num16000 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int    `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"16000"`
	Num17750 struct {
		Difficulty int64  `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"17750"`
	Num18081 struct {
		Difficulty int64  `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"18081"`
	Num18981 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"18981"`
	Num19281 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"19281"`
	Num19734 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"19734"`
	Num19950 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"19950"`
	Num19994 struct {
		Difficulty int64  `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"19994"`
	Num20206 struct {
		Difficulty int64  `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"20206"`
	Num25182 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"25182"`
	Num33124 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"33124"`
	Num38081 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"38081"`
	Num48782 struct {
		Difficulty int    `json:"difficulty"`
		Hash       string `json:"hash"`
		Height     int    `json:"height"`
		Value      int64  `json:"value"`
		Ts         int    `json:"ts"`
	} `json:"48782"`
	Difficulty int64  `json:"difficulty"`
	Hash       string `json:"hash"`
	MainHeight int    `json:"main_height"`
	Height     int    `json:"height"`
	Value      int64  `json:"value"`
	Ts         int    `json:"ts"`
}

// Worker a historical time series of personal stats
type Worker []struct {
	Ts  int     `json:"ts"`
	Hs  float64 `json:"hs"`
	Hs2 float64 `json:"hs2"`
}

// NetworkStats returns the global network stats
func NetworkStats() (*NetworkStatsResponse, error) {
	var networkStats *NetworkStatsResponse
	_, err := newRequestDecode("https://api.moneroocean.stream/network/stats", &networkStats)

	return networkStats, err
}

// PoolStats returns the global pool stats
func PoolStats() (*PoolStatsResponse, error) {
	var poolStats *PoolStatsResponse
	_, err := newRequestDecode("https://api.moneroocean.stream/pool/stats", &poolStats)

	return poolStats, err
}

// Stats returns personal stats
func Stats(address string) (*StatsResponse, error) {
	var info *StatsResponse
	_, err := newRequestDecode(fmt.Sprintf("https://api.moneroocean.stream/miner/%v/stats", address), &info)

	return info, err
}

// Workers returns personal worker stats
func Workers(address string) (map[string]Worker, error) {
	var workers map[string]Worker
	_, err := newRequestDecode(fmt.Sprintf("https://api.moneroocean.stream/miner/%v/allWorkers", address), &workers)

	return workers, err
}

// newRequestDecode a small helper function to create a new request and execute it
// with the configured Client, decoding the resulting value into v.
func newRequestDecode(url string, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(v)
	return res, err
}
