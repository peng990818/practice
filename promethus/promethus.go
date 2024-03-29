package promethus

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/prometheus/common/log"
    "math/rand"
    "net/http"
    "os"
)

var (
    cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "cpu_temperature_celsius",
        Help: "Current temperature of the CPU.",
    })
    hdFailures = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "hd_errors_total",
            Help: "Number of hard-disk errors.",
        },
        []string{"device"},
    )
)

type ClusterManager struct {
    Zone         string
    OOMCountDesc *prometheus.Desc
    RAMUsageDesc *prometheus.Desc
}

func init() {
    prometheus.MustRegister(cpuTemp)
    prometheus.MustRegister(hdFailures)
}

func test() {
    workerDB := NewClusterManager("db")
    workerCA := NewClusterManager("ca")

    // Since we are dealing with custom Collector implementations, it might
    // be a good idea to try it out with a pedantic registry.
    reg := prometheus.NewPedanticRegistry()
    reg.MustRegister(workerDB)
    reg.MustRegister(workerCA)

    gatherers := prometheus.Gatherers{
        prometheus.DefaultGatherer,
        reg,
    }

    h := promhttp.HandlerFor(gatherers,
        promhttp.HandlerOpts{
            ErrorLog:      log.NewErrorLogger(),
            ErrorHandling: promhttp.ContinueOnError,
        })
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        h.ServeHTTP(w, r)
    })
    log.Infoln("Start server at :10001")
    if err := http.ListenAndServe(":10001", nil); err != nil {
        log.Errorf("Error occur when start server %v", err)
        os.Exit(1)
    }

}

func (c *ClusterManager) ReallyExpensiveAssessmentOfTheSystemState() (
    oomCountByHost map[string]int, ramUsageByHost map[string]float64,
) {
    oomCountByHost = map[string]int{
        "foo.example.org": int(rand.Int31n(1000)),
        "bar.example.org": int(rand.Int31n(1000)),
    }
    ramUsageByHost = map[string]float64{
        "foo.example.org": rand.Float64() * 100,
        "bar.example.org": rand.Float64() * 100,
    }
    return
}

// Describe simply sends the two Descs in the struct to the channel.
func (c *ClusterManager) Describe(ch chan<- *prometheus.Desc) {
    ch <- c.OOMCountDesc
    ch <- c.RAMUsageDesc
}

func (c *ClusterManager) Collect(ch chan<- prometheus.Metric) {
    oomCountByHost, ramUsageByHost := c.ReallyExpensiveAssessmentOfTheSystemState()
    for host, oomCount := range oomCountByHost {
        ch <- prometheus.MustNewConstMetric(
            c.OOMCountDesc,
            prometheus.CounterValue,
            float64(oomCount),
            host,
        )
    }
    for host, ramUsage := range ramUsageByHost {
        ch <- prometheus.MustNewConstMetric(
            c.RAMUsageDesc,
            prometheus.GaugeValue,
            ramUsage,
            host,
        )
    }
}

func NewClusterManager(zone string) *ClusterManager {
    return &ClusterManager{
        Zone: zone,
        OOMCountDesc: prometheus.NewDesc(
            "clustermanager_oom_crashes_total",
            "Number of OOM crashes.",
            []string{"host"},
            prometheus.Labels{"zone": zone},
        ),
        RAMUsageDesc: prometheus.NewDesc(
            "clustermanager_ram_usage_bytes",
            "RAM usage as reported to the cluster manager.",
            []string{"host"},
            prometheus.Labels{"zone": zone},
        ),
    }
}
