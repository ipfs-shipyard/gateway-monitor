package tasks

import (
	"github.com/prometheus/client_golang/prometheus"

	logging "github.com/ipfs/go-log"

	"github.com/ipfs-shipyard/gateway-monitor/pkg/task"
)

// This file contains the list of tasks to be run (see All)
// as well as common metrics that might be useful for more than one task.

func init() {
	prometheus.Register(common_fetch_speed)
	prometheus.Register(common_fetch_latency)
}

const (
	kiB = 1024
	miB = 1024 * kiB
	giB = 1024 * miB
)

var (
	log = logging.Logger("tasks")

	All = []task.Task{
		NewRandomLocalBench("10,30,50 * * * *", 16*miB),
		NewRandomLocalBench("20 * * * *", 256*miB),
		NewIpnsBench("10,30,50 * * * *", 16*miB),
		NewIpnsBench("40 * * * *", 256*miB),
		NewKnownGoodCheck("* * * * *", map[string][]byte{
			"/ipfs/Qmc5gCcjYypU7y28oCALwfSvxCBskLuPKWpK4qpterKC7z": []byte("Hello World!\r\n"),
		}),
		NewNonExistCheck("0 * * * *"),
	}

	common_fetch_speed = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "gatewaymonitor_task",
			Subsystem: "common",
			Name:      "fetch_speed",
		})
	common_fetch_latency = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "gatewaymonitor_task",
			Subsystem: "common",
			Name:      "fetch_latency",
		})
)
