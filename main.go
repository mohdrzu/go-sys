package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

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

var (
	suffixes [5]string
)

func HumanFileSize(size float64) string {
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024)
	getSize := utils.Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
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
	fmt.Println(memInfo.Total)
	fmt.Println(memInfo.Available)
	fmt.Println(memInfo.Used)

	memData := model.RAM{
		Total:       HumanFileSize(float64(memInfo.Total)),
		Available:   HumanFileSize(float64(memInfo.Available)),
		Used:        HumanFileSize(float64(memInfo.Used)),
		UsedPercent: memInfo.UsedPercent,
		Free:        HumanFileSize(float64(memInfo.Free)),
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
