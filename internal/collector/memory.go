package collector

import (
    "github.com/shirou/gopsutil/v3/mem"
)

type MemoryCollector struct{}

func NewMemoryCollector() *MemoryCollector {
    return &MemoryCollector{}
}

func (m *MemoryCollector) Collect() (float64, error) {
    virtual, err := mem.VirtualMemory()
    if err != nil {
        return 0, err
    }
    
    return virtual.UsedPercent, nil
}