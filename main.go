package main

import (
	"fmt"
	"log"
	"net/http"

	"peerac/go-sys/model"
	"peerac/go-sys/utils"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	r := gin.Default()

	r.GET("/", SysInfoHandler)

	err := r.Run(":9000")
	if err != nil {
		log.Fatal("Error menjalankan server pada port 9000")
	}

}

func SysInfoHandler(c *gin.Context) {
	// Get Host Information
	hostInfo, _ := host.Info()
	hostData := model.Host{
		OS:              hostInfo.OS,
		Platform:        hostInfo.Platform,
		PlatformVersion: hostInfo.PlatformVersion,
		KernelVersion:   hostInfo.KernelVersion,
		Architecture:    hostInfo.KernelArch,
	}

	// Get CPU information
	cpuInfo, _ := cpu.Info()
	var cpus []model.CPU
	for _, v := range cpuInfo {
		cpuData := model.CPU{
			CpuNumber:  v.CPU,
			VendorID:   v.VendorID,
			Cores:      v.Cores,
			ModelName:  v.ModelName,
			ClockSpeed: v.Mhz,
		}

		cpus = append(cpus, cpuData)
	}

	// Get memory Information
	memInfo, _ := mem.VirtualMemory()
	memData := model.RAM{
		Total:       utils.HumanFileSize(float64(memInfo.Total)),
		Available:   utils.HumanFileSize(float64(memInfo.Available)),
		Used:        utils.HumanFileSize(float64(memInfo.Used)),
		UsedPercent: fmt.Sprintf("%v %%", memInfo.UsedPercent),
		Free:        utils.HumanFileSize(float64(memInfo.Free)),
	}

	// Populate response
	finalResponse := model.InfoResponse{
		Processors: cpus,
		Machine:    hostData,
		Memory:     memData,
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"data": finalResponse,
	})
}
