package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
)

type Exporter struct {
	*ZKCollector
	instance []string
}

type ZKOpt struct {
	IP []string
}

func NewExporter(opt ZKOpt, collector *ZKCollector) (*Exporter, error) {
	return &Exporter{collector, opt.IP}, nil
}

const (
	namespace = "zk"
	state     = "state"
	ok        = "ruok"
)

type ZKCollector struct {
	OK                      *prometheus.Desc
	AvgLatency              *prometheus.Desc
	MinLatency              *prometheus.Desc
	MaxLatency              *prometheus.Desc
	PackageReceived         *prometheus.Desc
	PackageSent             *prometheus.Desc
	NumAliveConnections     *prometheus.Desc
	OutstandingRequests     *prometheus.Desc
	ZnodeCount              *prometheus.Desc
	WatchCount              *prometheus.Desc
	EphemeralsCount         *prometheus.Desc
	ApproximateDataSize     *prometheus.Desc
	OpenFileDescriptorCount *prometheus.Desc
	MaxFileDescriptorCount  *prometheus.Desc
	Followers               *prometheus.Desc
	SyncedFollowers         *prometheus.Desc
	PendingSyncs            *prometheus.Desc
	ServerState             *prometheus.Desc
	FsyncThresholdExceeded  *prometheus.Desc
	Version                 *prometheus.Desc
}

func newZKCollector() *ZKCollector {
	return &ZKCollector{
		OK: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "ok"),
			"Was the last query of Zookeeper successful.", []string{"zk_instance"}, nil),
		AvgLatency: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "avglatency"),
			"Average Latency for ZooKeeper network requests.", []string{"zk_instance"}, nil),
		MinLatency: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "minlatency"),
			"Minimum latency for Zookeeper network requests.", []string{"zk_instance"}, nil),
		MaxLatency: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "maxlatency"),
			"Maximum latency for ZooKeeper network requests.", []string{"zk_instance"}, nil),
		PackageReceived: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "packagereceived"),
			"Number of network packets received by the ZooKeeper instance.", []string{"zk_instance"}, nil),
		PackageSent: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "packagesent"),
			"Number of network packets sent by the ZooKeeper instance.", []string{"zk_instance"}, nil),
		NumAliveConnections: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "numaliveconnections"),
			"Number of currently alive connections to the ZooKeeper instance.", []string{"zk_instance"}, nil),
		OutstandingRequests: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "outstandingrequests"),
			"Number of requests currently waiting in the queue.", []string{"zk_instance"}, nil),
		ZnodeCount: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "znodecount"),
			"Znode count.", []string{"zk_instance"}, nil),
		WatchCount: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "watchcount"),
			"Watch count.", []string{"zk_instance"}, nil),
		EphemeralsCount: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "ephemeralscount"),
			"Ephemerals Count.", []string{"zk_instance"}, nil),
		ApproximateDataSize: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "approximatedatasize"),
			"Approximate data size.", []string{"zk_instance"}, nil),
		OpenFileDescriptorCount: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "openfiledescriptiorcount"),
			"Number of currently open file descriptors.", []string{"zk_instance"}, nil),
		MaxFileDescriptorCount: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "maxfiledescriptorcount"),
			"Maximum number of open file descriptors", []string{"zk_instance"}, nil),
		Followers: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "followers"),
			"Leader only: number of followers.", []string{"zk_instance"}, nil),
		SyncedFollowers: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "syncedfollowers"),
			"Leader only: number of followers currently in sync.", []string{"zk_instance"}, nil),
		PendingSyncs: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "pendingsync"),
			"Current number of pending syncs", []string{"zk_instance"}, nil),
		ServerState: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "serverstate"),
			"Current state of the zk instance: 1 = follower, 2 = leader, 3 = standalone, -1 if unknown.", []string{"zk_instance"}, nil),
		FsyncThresholdExceeded: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "fsyncthresholdexceeded"),
			"Number of times File sync exceeded fsyncWarningThresholdMS", []string{"zk_instance"}, nil),
		Version: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "version"),
			"Zookeeper version", []string{"zk_instance", "zk_version"}, nil),
	}
}

func (c *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.OK
	ch <- c.Version
	ch <- c.ApproximateDataSize
	ch <- c.AvgLatency
	ch <- c.MaxLatency
	ch <- c.MinLatency
	ch <- c.EphemeralsCount
	ch <- c.Followers
	ch <- c.FsyncThresholdExceeded
	ch <- c.MaxFileDescriptorCount
	ch <- c.NumAliveConnections
	ch <- c.OpenFileDescriptorCount
	ch <- c.OutstandingRequests
	ch <- c.PackageReceived
	ch <- c.PackageSent
	ch <- c.ServerState
	ch <- c.ZnodeCount
	ch <- c.WatchCount
	ch <- c.SyncedFollowers
	ch <- c.PendingSyncs
}

func (c *Exporter) Collect(ch chan<- prometheus.Metric) {
	c.CollectOK(ch)
}

func (c *Exporter) CollectOK(ch chan<- prometheus.Metric) {
	var wg sync.WaitGroup
	for _, address := range c.instance {
		wg.Add(1)
		go func(address string, ch chan<- prometheus.Metric) {
			defer wg.Done()
			zk := newZk(address)
			ruok, err := zk.sendMsg(ok)
			if err != nil {
				log.Error(err)
			}
			var okMetric float64 = 0.
			if string(ruok) == "imok" {
				okMetric = 1.
			}
			ch <- prometheus.MustNewConstMetric(c.OK, prometheus.GaugeValue, okMetric, zk.Instance)
		}(address, ch)
		wg.Wait()
	}
}

type zkInstance struct {
	Instance string
	Conn     net.Conn
}

func newZk(address string) *zkInstance {
	if conn, err := net.Dial("tcp", address); err != nil {
		return &zkInstance{}
	} else {
		return &zkInstance{
			Instance: address,
			Conn:     conn,
		}
	}
}

func (z zkInstance) sendMsg(cmd string) (string, error) {
	z.Conn.Write([]byte(cmd))
	defer z.Conn.Close()
	if ret, err := ioutil.ReadAll(z.Conn); err != nil {
		return string(ret), err
	} else {
		return string(ret), nil
	}
}

func main() {
	exporter, err := NewExporter(ZKOpt{
		IP: []string{"192.168.55.161:2181", "192.168.55.162:2181", "192.168.55.171:2181"}},
		newZKCollector())
	if err != nil {
		log.Fatalf("exporter create failed! ", err)
	}
	prometheus.MustRegister(exporter)
	log.Info("beging to server on Port: 18082")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":18082", nil))
}
