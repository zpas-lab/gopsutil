package base_board

import (
	"github.com/zpas-lab/gopsutil/internal/common"
)

var invoke common.Invoker

func init() {
	invoke = common.Invoke{}
}

type BaseBoardInfo struct {
	SerialNumber string `json:"serialNumber"`
	Manufacturer string `json:"serialNumber"`
	ProductName string `json:"serialNumber"`
	Version string `json:"serialNumber"`
}

