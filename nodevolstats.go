// Copyright 2021 Hewlett Packard Enterprise Development LP

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/hpe-storage/csi-driver/pkg/driver"
)

func TestNodeGetVolumeStats(req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	driver, err := driver.NewDriver(
		"csi.hpe.com",
		"2.0",
		"unix:///var/lib/kubelet/csi.hpe.com/csi.sock",
		"",
		true,
		"",
		"",
		false,
		30,
	)

	if err != nil {
		return nil, err
	}

	if req.VolumeId == "" || req.VolumePath == "" {
		return nil, fmt.Errorf("empty volumeID or mountPath")
	}

	response, err := driver.NodeGetVolumeStats(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func main() {
	volumeID := flag.String("volumeID", "", "volume id")
	mountPath := flag.String("mountPath", "", "mount path")
	flag.Parse()

	req := &csi.NodeGetVolumeStatsRequest{
		VolumeId:   *volumeID,
		VolumePath: *mountPath,
	}
	resp, err := TestNodeGetVolumeStats(req)
	if err != nil {
		fmt.Printf("error=%s", err.Error())
	}

	fmt.Printf("NodeGetVolumeStats for volume with id %s, %+v", resp)

}
