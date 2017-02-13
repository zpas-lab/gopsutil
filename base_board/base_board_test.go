package system

import (
	"os/exec"
	"strings"
	"encoding/json"
)

func GetBaseBoardStat() (BaseBoardInfo, error) {
	path, err := exec.LookPath("dmidecode")
	if err != nil {
		return BaseBoardInfo{}, err
	}

	// 2 is mapped into "Base Board"
	out, err := invoke.Command(path, "-t", "2")
	if err != nil {
		return BaseBoardInfo{}, err
	}

	stat := &BaseBoardInfo{}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		field := strings.Split(line, ":")
		if len(field) < 2 {
			continue
		}
		name := strings.TrimSpace(field[0])
		value := strings.TrimSpace(field[1])
		switch name {
		case "Manufacturer":
			stat.Manufacturer = value
		case "Product Name":
			stat.ProductName = value
		case "Version":
			stat.Version = value
		case "Serial Number":
			stat.SerialNumber = value
		}
	}

	return *stat, nil
}

func (c BaseBoardInfo) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}