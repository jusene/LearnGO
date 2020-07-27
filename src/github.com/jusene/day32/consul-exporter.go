package main

import (
	"fmt"
	consul_api "github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
)

type Exporter struct {
	*consulCollector
	client *consul_api.Client
}

type ConsulOpt struct {
	ip string
	port int
}

func NewExporter(opt ConsulOpt, collector *consulCollector) (*Exporter, error) {
	conf := consul_api.DefaultConfig()
	conf.Address = fmt.Sprintf("%s:%d", opt.ip, opt.port)
	client, err := consul_api.NewClient(conf)
	if err != nil {
		return &Exporter{}, err
	}
	return &Exporter{
		collector, client,
	}, nil

}

const namespace = "consul"

type consulCollector struct {
	up                 *prometheus.Desc
	clusterServers     *prometheus.Desc
	clusterLeader      *prometheus.Desc
	nodeCount          *prometheus.Desc
	memberStatus       *prometheus.Desc
	serviceCount       *prometheus.Desc
	serviceTag         *prometheus.Desc
	serviceNodeHealthy *prometheus.Desc
	nodeChecks         *prometheus.Desc
	serviceChecks      *prometheus.Desc
	serviceCheckNames  *prometheus.Desc
	keyValues          *prometheus.Desc
}

func newConsulCollector() *consulCollector {
	return &consulCollector{
		up: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "up"),
			"Was the last query of Consul successful.", nil, nil),
		clusterServers: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "raft_peers"),
			"How many peers (servers) are in the Raft cluster.", nil, nil),
		clusterLeader: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "raft_leader"),
			"Does Raft cluster have a leader (according to this node).", nil, nil),
		nodeCount: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "serf_lan_members"),
			"How many members are in the cluster.", nil, nil),
		memberStatus: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "serf_lan_member_status"),
			"Status of member in the cluster. 1=Alive, 2=Leaving, 3=Left, 4=Failed.", []string{"member"}, nil),
		serviceCount: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "catalog_services"),
			"How many services are in the cluster.", nil, nil),
		serviceTag: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "service_tag"),
			"Tags of a service.", []string{"service_id", "node", "tag"}, nil),
		serviceNodeHealthy: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "catalog_service_node_healthy"),
			"Is this service healthy on this node?", []string{"service_id", "node", "service_name"}, nil),
		nodeChecks: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "health_node_status"),
			"Status of health checks associated with a node.", []string{"check", "node", "status"}, nil),
		serviceChecks: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "health_service_status"),
			"Status of health checks associated with a service.", []string{"check", "node", "service_id", "service_name", "status"}, nil),
		serviceCheckNames: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "service_checks"),
			"Link the service id and check name if available.", []string{"service_id", "service_name", "check_id", "check_name", "node"}, nil),
		keyValues: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "catalog_kv"),
			"The values for selected keys in Consul's key/value catalog. Keys with non-numeric values are omitted.", []string{"key"}, nil),
	}
}


func (c *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.clusterServers
	ch <- c.clusterLeader
	ch <- c.nodeCount
	ch <- c.memberStatus
	ch <- c.serviceCount
	ch <- c.serviceNodeHealthy
	ch <- c.nodeChecks
	ch <- c.serviceChecks
	ch <- c.keyValues
	ch <- c.serviceTag
	ch <- c.serviceCheckNames
}

func (c *Exporter) Collect(ch chan<- prometheus.Metric) {
	ok := c.collectPeerMetric(ch)
	ok = c.collectLeaderMertic(ch) && ok
	ok = c.collectNodeMetric(ch) && ok
	ok = c.collectMembersMetric(ch) && ok

	if ok {
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1.0)
	} else {
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0.0)
	}
}

func (c *Exporter) collectPeerMetric(ch chan <- prometheus.Metric) bool {
	peers, err := c.client.Status().Peers()
	if err != nil {
		log.Errorf("can not query consul: %v", err)
		return false
	}
	ch <- prometheus.MustNewConstMetric(c.clusterServers, prometheus.GaugeValue, float64(len(peers)))
	return true
}

func (c *Exporter) collectLeaderMertic(ch chan <- prometheus.Metric) bool {
	leader, err := c.client.Status().Leader()
	if err != nil {
		log.Errorf("can not query consul: %v", err)
		return false
	}
	if len(leader) == 0 {
		ch <- prometheus.MustNewConstMetric(c.clusterLeader, prometheus.GaugeValue, 0)
	} else {
		ch <- prometheus.MustNewConstMetric(c.clusterLeader, prometheus.GaugeValue, 1)
	}
	return true
}

func (c *Exporter) collectNodeMetric(ch chan <- prometheus.Metric) bool {
	nodes, _, err := c.client.Catalog().Nodes(&consul_api.QueryOptions{})
	if err != nil {
		log.Errorf("can not query consul: %v", err)
		return false
	}
	ch <- prometheus.MustNewConstMetric(c.nodeCount, prometheus.GaugeValue, float64(len(nodes)))
	return true
}

func (c *Exporter) collectMembersMetric(ch chan <- prometheus.Metric) bool {
	members, err := c.client.Agent().Members(false)
	if err != nil {
		log.Errorf("can not query consul: %v", err)
		return false
	}

	for _, enty := range members {
		ch <- prometheus.MustNewConstMetric(c.memberStatus, prometheus.GaugeValue, float64(enty.Status), enty.Name)
	}
	return true
}

func (c *Exporter) collectServicesMetric(ch chan <- prometheus.Metric) bool {
	serviceNames, _, err := c.client.Catalog().Services(&consul_api.QueryOptions{})
	if err != nil {
		log.Errorf("can not query consul: %v", err)
		return false
	}
	ch <- prometheus.MustNewConstMetric(c.serviceCount, prometheus.GaugeValue, float64(len(serviceNames)))
	if c.
}

func (c *Exporter) collectHealthSummary(ch chan <- prometheus.Metric, serviceNames map[string][]string) bool {
	ok := make(chan bool)
	for s := range serviceNames {
		go func() {
			if 
		}()
	}
}

func main() {
	exporter, err := NewExporter(ConsulOpt{
		ip:   "192.168.66.100",
		port: 8500,
	}, newConsulCollector())
	if err != nil {
		log.Fatalf("exporter create failed!", err)
	}
	prometheus.MustRegister(exporter)
	log.Info("beging to server on Port: 18081")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":18081", nil))
}
