package main

import (
	"bytes"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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
	watch     = "wchs"
)

type ZKCollector struct {
	OK                  *prometheus.Desc
	AvgLatency          *prometheus.Desc
	MinLatency          *prometheus.Desc
	MaxLatency          *prometheus.Desc
	PackageReceived     *prometheus.Desc
	PackageSent         *prometheus.Desc
	NumAliveConnections *prometheus.Desc
	OutstandingRequests *prometheus.Desc
	ZnodeCount          *prometheus.Desc
	WatchCount          *prometheus.Desc
	ServerState         *prometheus.Desc
	Version             *prometheus.Desc
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
		ServerState: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "serverstate"),
			"Current state of the zk instance: 1 = follower, 2 = leader, 3 = standalone, -1 if unknown.", []string{"zk_instance"}, nil),
		Version: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "version"),
			"Zookeeper version", []string{"zk_instance", "zk_version"}, nil),
	}
}

func (c *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.OK
	ch <- c.Version
	ch <- c.AvgLatency
	ch <- c.MaxLatency
	ch <- c.MinLatency
	ch <- c.NumAliveConnections
	ch <- c.OutstandingRequests
	ch <- c.PackageReceived
	ch <- c.PackageSent
	ch <- c.ServerState
	ch <- c.ZnodeCount
	ch <- c.WatchCount
}

func (c *Exporter) Collect(ch chan<- prometheus.Metric) {
	c.CollectOK(ch)
	c.CollectState(ch)
	c.CollectWatch(ch)
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
			if ruok == "imok" {
				okMetric = 1.
			}
			ch <- prometheus.MustNewConstMetric(c.OK, prometheus.GaugeValue, okMetric, zk.Instance)
		}(address, ch)
		wg.Wait()
	}
}

func (c *Exporter) CollectState(ch chan<- prometheus.Metric) {
	var wg sync.WaitGroup
	for _, address := range c.instance {
		wg.Add(1)
		go func(address string, ch chan<- prometheus.Metric) {
			defer wg.Done()
			zk := newZk(address)
			state, err := zk.sendMsg(state)
			if err != nil {
				log.Error(err)
			}
			// fmt.Println(state)
			viper.SetConfigType("yaml")
			viper.ReadConfig(bytes.NewBuffer([]byte(state)))

			// zk version
			zkVersion := viper.GetString("Zookeeper version")
			ch <- prometheus.MustNewConstMetric(c.Version, prometheus.GaugeValue, 1, zk.Instance, zkVersion)

			// zk latency
			Latency := viper.GetString("Latency min/avg/max")
			minLatency, _ := strconv.ParseFloat(strings.Split(Latency, "/")[0], 64)
			avgLatency, _ := strconv.ParseFloat(strings.Split(Latency, "/")[1], 64)
			maxLatency, _ := strconv.ParseFloat(strings.Split(Latency, "/")[2], 64)
			ch <- prometheus.MustNewConstMetric(c.MinLatency, prometheus.GaugeValue, minLatency, zk.Instance)
			ch <- prometheus.MustNewConstMetric(c.AvgLatency, prometheus.GaugeValue, avgLatency, zk.Instance)
			ch <- prometheus.MustNewConstMetric(c.MaxLatency, prometheus.GaugeValue, maxLatency, zk.Instance)

			// zk AliveConnections Nums
			AliveCon := viper.GetFloat64("Connections")
			ch <- prometheus.MustNewConstMetric(c.NumAliveConnections, prometheus.GaugeValue, AliveCon, zk.Instance)

			// zk PackageReceived
			PackageRecv := viper.GetFloat64("Received")
			ch <- prometheus.MustNewConstMetric(c.PackageReceived, prometheus.GaugeValue, PackageRecv, zk.Instance)

			// zk PackageSent
			PackageSent := viper.GetFloat64("Sent")
			ch <- prometheus.MustNewConstMetric(c.PackageSent, prometheus.GaugeValue, PackageSent, zk.Instance)

			// zk Outstanding
			Outstanding := viper.GetFloat64("Outstanding")
			ch <- prometheus.MustNewConstMetric(c.OutstandingRequests, prometheus.GaugeValue, Outstanding, zk.Instance)

			// zk znodecount
			ZnodeCount := viper.GetFloat64("Node count")
			ch <- prometheus.MustNewConstMetric(c.ZnodeCount, prometheus.GaugeValue, ZnodeCount, zk.Instance)

			// zk state
			ZKState := viper.GetString("Mode")
			var ZKStateID float64
			switch ZKState {
			case "follower":
				ZKStateID = 1
			case "leader":
				ZKStateID = 2
			case "standalone":
				ZKStateID = 3
			default:
				ZKStateID = -1
			}
			ch <- prometheus.MustNewConstMetric(c.ServerState, prometheus.GaugeValue, ZKStateID, zk.Instance)
		}(address, ch)
		wg.Wait()
	}
}

func (c *Exporter) CollectWatch(ch chan<- prometheus.Metric) {
	var wg sync.WaitGroup
	for _, address := range c.instance {
		wg.Add(1)
		go func(address string, ch chan<- prometheus.Metric) {
			defer wg.Done()
			zk := newZk(address)
			wa, err := zk.sendMsg(watch)
			if err != nil {
				log.Error(err)
			}
			reg, _ := regexp.Compile(".*:.*")
			watchCount, _ := strconv.ParseFloat(strings.Split(string(reg.Find([]byte(wa))), ":")[1], 64)

			// zk watch
			ch <- prometheus.MustNewConstMetric(c.WatchCount, prometheus.GaugeValue, watchCount, zk.Instance)
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
