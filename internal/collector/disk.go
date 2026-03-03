package collector

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type DiskCollector struct {
	path string
}

func NewDiskCollector(path string) *DiskCollector {
	return &DiskCollector{
		path: path,
	}
}

func (d *DiskCollector) Collect() (float64, error) {
	usage, err := disk.Usage(d.path)
	if err != nil {
		return 0, err
	}

	return usage.UsedPercent, nil
}