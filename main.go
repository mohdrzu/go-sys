package main

import (
	"log"
	"net/http"

	"peerac/go-sys/model"
	"peerac/go-sys/utils"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	r := gin.Default()

	// Static file
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("views/*.html")

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

	// Get memory Information
	memInfo, _ := mem.VirtualMemory()
	memData := model.RAM{
		Total:       utils.HumanFileSize(float64(memInfo.Total)),
		Available:   utils.HumanFileSize(float64(memInfo.Available)),
		Used:        utils.HumanFileSize(float64(memInfo.Used)),
		UsedPercent: memInfo.UsedPercent,
		Free:        utils.HumanFileSize(float64(memInfo.Free)),
	}

	// Populate response
	finalResponse := model.InfoResponse{
		Machine: hostData,
		Memory:  memData,
	}

	// Return response
	c.HTML(http.StatusOK, "index", gin.H{
		"data": finalResponse,
	})
}
