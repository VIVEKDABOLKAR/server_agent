package collector

import (
    "time"

    "github.com/shirou/gopsutil/v3/cpu"
)

type CPUCollector struct{}

func NewCPUCollector() *CPUCollector {
    return &CPUCollector{}
}

func (c *CPUCollector) Collect() (float64, error) {
    percent, err := cpu.Percent(1*time.Second, false)
    if err != nil {
        return 0, err
    }
    
    if len(percent) > 0 {
        return percent[0], nil
    }
    return 0, nil
}