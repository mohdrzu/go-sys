package model

type CPU struct {
	CpuNumber  int32
	VendorID   string
	Cores      int32
	ModelName  string
	ClockSpeed float64
}

type Host struct {
	OS              string
	Platform        string
	PlatformVersion string
	KernelVersion   string
	Architecture    string
}

type RAM struct {
	Total       string
	Available   string
	Used        string
	UsedPercent float64
	Free        string
}

type InfoResponse struct {
	Processors []CPU `json:"Processors"`
	Machine    Host  `json:"Host"`
	Memory     RAM   `json:"Memory"`
}
