package diagnosis

import (
	"github.com/rudderlabs/analytics-go"
	"github.com/rudderlabs/rudder-server/config"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	StartTime           = "diagnosis_start_time"
	ServerStart         = "server_start"
	ConfigProcessed     = "config_processed"
	SourcesCount        = "no_of_sources"
	DesitanationCount   = "no_of_destinations"
	ServerStarted       = "server_started"
	ConfigIdentify      = "identify"
	GatewayEvents       = "gateway_events"
	GatewaySuccess      = "gateway_success"
	GatewayFailure      = "gateway_failure"
	RouterEvents        = "router_events"
	RouterType          = "router_type"
	RouterAborted       = "router_aborted"
	RouterRetries       = "router_retries"
	RouterSuccess       = "router_success"
	RouterCompletedTime = "router_completed_time"
	BatchRouterEvents   = "batch_router_events"
	BatchRouterType     = "batch_router_type"
	BatchRouterSuccess  = "batch_router_success"
	BatchRouterFailed   = "batch_router_failed"
)

var (
	EnableDiagnosis bool
	rudderEndpoint  string
)

var diagnosis Diagnosis

type Diagnosis struct {
	Client    analytics.Client
	StartTime time.Time
	serverId  string
}

func init() {
	EnableDiagnosis = config.GetBool("Diagnosis.enableDiagnosis", true)
	rudderEndpoint = config.GetString("Diagnosis.endpoint", "http://localhost:8080")
	config := analytics.Config{
		Endpoint: rudderEndpoint,
	}
	client, _ := analytics.NewWithConfig("1TnQwbNV2QBdOsVlZIeKsvP2cez", config)
	diagnosis.Client = client
	diagnosis.StartTime = time.Now()
	diagnosis.serverId = uuid.NewV4().String()
}

func Track(event string, properties map[string]interface{}) {
	if EnableDiagnosis {
		properties[StartTime] = diagnosis.StartTime
		diagnosis.Client.Enqueue(
			analytics.Track{
				Event:      event,
				Properties: properties,
				UserId:     diagnosis.serverId,
			},
		)
	}
}

func Identify(event string, properties map[string]interface{}) {
	if EnableDiagnosis {
		properties[StartTime] = diagnosis.StartTime
		diagnosis.Client.Enqueue(
			analytics.Track{
				Event:      event,
				Properties: properties,
				UserId:     diagnosis.serverId,
			},
		)
	}
}
