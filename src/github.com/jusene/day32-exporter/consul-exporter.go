package main

import (
	"fmt"
	consul_api "github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
	"strings"
)

type Exporter struct {
	*consulCollector
	client *consul_api.Client
}

type ConsulOpt struct {
	ip   string
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
	ok = c.collectServicesMetric(ch) && ok
	ok = c.collectHealthSateMetric(ch) && ok
	ok = c.collectKeyValue(ch) && ok

	if ok {
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1.0)
	} else {
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0.0)
	}
}

func (c *Exporter) collectPeerMetric(ch chan<- prometheus.Metric) bool {
	peers, err := c.client.Status().Peers()
	if err != nil {
		log.Errorf("can not query consul: %v", err)
		return false
	}
	ch <- prometheus.MustNewConstMetric(c.clusterServers, prometheus.GaugeValue, float64(len(peers)))
	return true
}

func (c *Exporter) collectLeaderMertic(ch chan<- prometheus.Metric) bool {
	leader, err := c.client.Status().Leader()
	if err != nil {
		log.Errorf("can not query consul: %v", err)
		return false
	}

	// 收集consul leader数据
	if len(leader) == 0 {
		ch <- prometheus.MustNewConstMetric(c.clusterLeader, prometheus.GaugeValue, 0)
	} else {
		ch <- prometheus.MustNewConstMetric(c.clusterLeader, prometheus.GaugeValue, 1)
	}
	return true
}

func (c *Exporter) collectNodeMetric(ch chan<- prometheus.Metric) bool {
	nodes, _, err := c.client.Catalog().Nodes(&consul_api.QueryOptions{})
	if err != nil {
		log.Errorf("can not query consul: %v", err)
		return false
	}

	// 收集consul node数据
	ch <- prometheus.MustNewConstMetric(c.nodeCount, prometheus.GaugeValue, float64(len(nodes)))
	return true
}

func (c *Exporter) collectMembersMetric(ch chan<- prometheus.Metric) bool {
	members, err := c.client.Agent().Members(false)
	if err != nil {
		log.Errorf("can not query consul: %v", err)
		return false
	}

	// 收集consul member数据
	for _, enty := range members {
		ch <- prometheus.MustNewConstMetric(c.memberStatus, prometheus.GaugeValue, float64(enty.Status), enty.Name)
	}
	return true
}

func (c *Exporter) collectServicesMetric(ch chan<- prometheus.Metric) bool {
	serviceNames, _, err := c.client.Catalog().Services(&consul_api.QueryOptions{})
	if err != nil {
		log.Errorf("can not query consul: %v", err)
		return false
	}

	// 收集consul服务数
	ch <- prometheus.MustNewConstMetric(c.serviceCount, prometheus.GaugeValue, float64(len(serviceNames)))

	// 收集服务的健康数据
	if ok := c.collectHealthSummary(ch, serviceNames); !ok {
		return false
	}
	return true
}

func (c *Exporter) collectHealthSummary(ch chan<- prometheus.Metric, serviceNames map[string][]string) bool {
	ok := make(chan bool)
	for s := range serviceNames {
		go func(s string) {
			ok <- c.collectOneHealthSummary(ch, s)
		}(s)
	}

	allOK := true
	for range serviceNames {
		allOK = <-ok && allOK
	}
	close(ok)
	return allOK
}

func (c *Exporter) collectOneHealthSummary(ch chan<- prometheus.Metric, serviceName string) bool {

	if strings.HasPrefix(serviceName, "/") {
		log.Warn("Skipping service because it starts with a slash service_name", serviceName)
		return true
	}
	log.Info("Fetching health summary serviceName ", serviceName)

	service, _, err := c.client.Health().Service(serviceName, "", false, &consul_api.QueryOptions{})
	if err != nil {
		log.Error("Failed to query service health err", err)
		return false
	}

	for _, entry := range service {
		passing := 1.
		for _, hc := range entry.Checks {
			if hc.Status != consul_api.HealthPassing {
				passing = 0.
				break
			}
		}
		ch <- prometheus.MustNewConstMetric(
			c.serviceNodeHealthy, prometheus.GaugeValue, passing, entry.Service.ID, entry.Node.Node, entry.Service.Service,
		)
		tags := make(map[string]struct{})
		for _, tag := range entry.Service.Tags {
			if _, ok := tags[tag]; ok {
				continue
			}
			ch <- prometheus.MustNewConstMetric(c.serviceTag, prometheus.GaugeValue, 1, entry.Service.ID, entry.Node.Node, tag)
			tags[tag] = struct{}{}
		}
	}
	return true
}

func (c *Exporter) collectHealthSateMetric(ch chan<- prometheus.Metric) bool {
	checks, _, err := c.client.Health().State("any", &consul_api.QueryOptions{})
	if err != nil {
		log.Error("Failed to query service health err: ", err)
		return false
	}

	for _, hc := range checks {
		var passing, warning, critical, maintenance float64
		switch hc.Status {
		case consul_api.HealthPassing:
			passing = 1.
		case consul_api.HealthWarning:
			warning = 1.
		case consul_api.HealthCritical:
			critical = 1.
		case consul_api.HealthMaint:
			maintenance = 1.
		}

		if hc.ServiceID == "" {
			ch <- prometheus.MustNewConstMetric(
				c.nodeChecks, prometheus.GaugeValue, passing, hc.CheckID, hc.Node, consul_api.HealthPassing,
			)
			ch <- prometheus.MustNewConstMetric(
				c.nodeChecks, prometheus.GaugeValue, warning, hc.CheckID, hc.Node, consul_api.HealthWarning,
			)
			ch <- prometheus.MustNewConstMetric(
				c.nodeChecks, prometheus.GaugeValue, critical, hc.CheckID, hc.Node, consul_api.HealthCritical,
			)
			ch <- prometheus.MustNewConstMetric(
				c.nodeChecks, prometheus.GaugeValue, maintenance, hc.CheckID, hc.Node, consul_api.HealthMaint,
			)
		} else {
			ch <- prometheus.MustNewConstMetric(
				c.serviceChecks, prometheus.GaugeValue, passing, hc.CheckID, hc.Node, hc.ServiceID, hc.ServiceName, consul_api.HealthPassing,
			)
			ch <- prometheus.MustNewConstMetric(
				c.serviceChecks, prometheus.GaugeValue, warning, hc.CheckID, hc.Node, hc.ServiceID, hc.ServiceName, consul_api.HealthWarning,
			)
			ch <- prometheus.MustNewConstMetric(
				c.serviceChecks, prometheus.GaugeValue, critical, hc.CheckID, hc.Node, hc.ServiceID, hc.ServiceName, consul_api.HealthCritical,
			)
			ch <- prometheus.MustNewConstMetric(
				c.serviceChecks, prometheus.GaugeValue, maintenance, hc.CheckID, hc.Node, hc.ServiceID, hc.ServiceName, consul_api.HealthMaint,
			)
			ch <- prometheus.MustNewConstMetric(
				c.serviceCheckNames, prometheus.GaugeValue, 1, hc.ServiceID, hc.ServiceName, hc.CheckID, hc.Name, hc.Node,
			)
		}
	}
	return true
}

func (c *Exporter) collectKeyValue(ch chan<- prometheus.Metric) bool {
	kv := c.client.KV()
	pairs, _, err := kv.List("/oneci/template/arch", &consul_api.QueryOptions{})
	if err != nil {
		log.Error("Error fetching key/values err: ", err)
		return false
	}

	for _, pair := range pairs {
		ch <- prometheus.MustNewConstMetric(
			c.keyValues, prometheus.GaugeValue, 1, pair.Key)

	}
	return true
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
