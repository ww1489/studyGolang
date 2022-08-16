package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	v, _ := mem.VirtualMemory()
	res, err := cpu.Times(true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res[0].Total(), res[0].User, res[0].System, res[0].Idle)
	memTotal := float64(v.Total) / 1024 / 1024 / 1024
	memUsed := float64(v.Used) / 1024 / 1024 / 1024

	fmt.Printf("总内存: %.2f GB, 已使用:%.2f GB, 已使用百分比:%.2f%%\n", memTotal,
		memUsed, (memUsed/memTotal)*100)
}
