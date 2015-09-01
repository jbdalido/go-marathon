package gomarathon

// RequestOptions passed for query api
type RequestOptions struct {
	Method string
	Path   string
	Datas  interface{}
	Params *Parameters
}

// Parameters to build url query
type Parameters struct {
	Cmd         string
	Host        string
	Scale       bool
	CallbackURL string
}

// Response representation of a full marathon response
type Response struct {
	Code     int
	Apps     []*Application `json:"apps,omitempty"`
	App      *Application   `json:"app,omitempty"`
	Leader   string         `json:"leader,omitempty"`
	Message  string         `json:"message,omitempty"`
	Versions []string       `json:",omitempty"`
	Tasks    []*Task        `json:"tasks,omitempty"`

	// See ServerInfo below
	FrameworkId    string              `json:"frameworkId,omitempty"`
	EventSub       *InfoEventSub       `json:"event_subscriber,omitempty"`
	HttpConfig     *InfoHttpConfig     `json:"http_config,omitempty"`
	MarathonConfig *InfoMarathonConfig `json:"marathon_config,omitempty"`
	Name           string              `json:"name,omitempty"`
	Version        string              `json:"version,omitempty"`
	ZkConfig       *InfoZkConfig       `json:"zookeeper_config,omitempty"`
}

// Application marathon application see :
// https://github.com/mesosphere/marathon/blob/master/REST.md#apps
type Application struct {
	ID                    string            `json:"id"`
	Cmd                   string            `json:"cmd,omitempty"`
	Constraints           [][]string        `json:"constraints,omitempty"`
	Container             *Container        `json:"container,omitempty"`
	CPUs                  float32           `json:"cpus,omitempty"`
	Deployments           []*Deployment     `json:"deployments,omitempty"`
	Env                   map[string]string `json:"env,omitempty"`
	Executor              string            `json:"executor,omitempty"`
	HealthChecks          []*HealthCheck    `json:"healthChecks,omitempty"`
	Instances             int               `json:"instances,omitemptys"`
	Mem                   float32           `json:"mem,omitempty"`
	Tasks                 []*Task           `json:"tasks,omitempty"`
	Ports                 []int             `json:"ports,omitempty"`
	RequirePorts          bool              `json:"requirePorts,omitempty"`
	BackoffSeconds        float64           `json:"backoffSeconds,omitempty"`
	BackoffFactor         float32           `json:"backoffFactor,omitempty"`
	MaxLaunchDelaySeconds float64           `json:"maxLaunchDelaySeconds,omitempty"`
	TasksRunning          int               `json:"tasksRunning,omitempty"`
	TasksStaged           int               `json:"tasksStaged,omitempty"`
	UpgradeStrategy       *UpgradeStrategy  `json:"upgradeStrategy,omitempty"`
	Uris                  []string          `json:"uris,omitempty"`
	Version               string            `json:"version,omitempty"`
}

// Container is docker parameters
type Container struct {
	Type    string    `json:"type,omitempty"`
	Docker  *Docker   `json:"docker,omitempty"`
	Volumes []*Volume `json:"volumes,omitempty"`
}

// Docker options
type Docker struct {
	Image        string         `json:"image,omitempty"`
	Network      string         `json:"network,omitempty"`
	PortMappings []*PortMapping `json:"portMappings,omitempty"`
}

// Volume is used for mounting a host directory as a container volume
type Volume struct {
	ContainerPath string `json:"containerPath,omitempty"`
	HostPath      string `json:"hostPath,omitempty"`
	Mode          string `json:"mode,omitempty"`
}

// Container PortMappings
type PortMapping struct {
	ContainerPort int    `json:"containerPort,omitempty"`
	HostPort      int    `json:"hostPort,omitempty"`
	ServicePort   int    `json:"servicePort,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

// UpgradeStrategy has a minimumHealthCapacity which defines the minimum number of healty nodes
type UpgradeStrategy struct {
	MinimumHealthCapacity float32 `json:"minimumHealthCapacity,omitempty"`
}

// HealthCheck is described here:
// https://github.com/mesosphere/marathon/blob/master/REST.md#healthchecks
type HealthCheck struct {
	Protocol           string `json:"protocol,omitempty"`
	Path               string `json:"path,omitempty"`
	GracePeriodSeconds int    `json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds    int    `json:"intervalSeconds,omitempty"`
	PortIndex          int    `json:"portIndex,omitempty"`
	TimeoutSeconds     int    `json:"timeoutSeconds,omitempty"`
}

// Task is described here:
// https://github.com/mesosphere/marathon/blob/master/REST.md#tasks
type Task struct {
	AppID     string `json:"appId"`
	Host      string `json:"host"`
	ID        string `json:"id"`
	Ports     []int  `json:"ports"`
	StagedAt  string `json:"stagedAt"`
	StartedAt string `json:"startedAt"`
	Version   string `json:"version"`
}

// Deployment is described here:
// https://mesosphere.github.io/marathon/docs/rest-api.html#get-/v2/deployments
type Deployment struct {
	AffectedApps   []string          `json:"affectedApps"`
	ID             string            `json:"id"`
	Steps          []*DeploymentStep `json:"steps"`
	CurrentActions []*DeploymentStep `json:"currentActions"`
	CurrentStep    int               `json:"currentStep"`
	TotalSteps     int               `json:"totalSteps"`
	Version        string            `json:"version"`
}

// Deployment steps
type DeploymentStep struct {
	Action string `json:"action"`
	App    string `json:"app"`
}

// EventSubscription is described here:
// https://github.com/mesosphere/marathon/blob/master/REST.md#event-subscriptions
type EventSubscription struct {
	CallbackURL  string   `json:"CallbackUrl"`
	ClientIP     string   `json:"ClientIp"`
	EventType    string   `json:"eventType"`
	CallbackURLs []string `json:"CallbackUrls"`
}

// ServerInfo is described here:
// https://mesosphere.github.io/marathon/docs/rest-api.html#server-info

type InfoEventSub struct {
	Type          string   `json:"type"`
	HttpEndpoints []string `json:"http_endpoints,omitempty"`
}

type InfoMarathonConfig struct {
	Checkpoint        bool   `json:"checkpoint"`
	Executor          string `json:"executor"`
	FailoverTimeout   int    `json:"failover_timeout"`
	HA                bool   `json:"ha"`
	Hostname          string `json:"hostname"`
	LocalPortMin      int    `json:"local_port_min"`
	LocalPortMax      int    `json:"local_port_max"`
	Master            string `json:"master"`
	MesosRole         string `json:"mesos_role"`
	MesosUser         string `json:"mesos-user"`
	RecInitialDelay   int    `json:"reconciliation_initial_delay"`
	RecInterval       int    `json:"reconciliation_interval"`
	TaskLaunchTimeout int    `json:"task_launch_timeout"`
}

type InfoHttpConfig struct {
	AssetsPath string `json:"assets_path,omitempty"`
	HttpPort   int    `json:"http_port"`
	HttpsPort  int    `json:"https_port"`
}

type InfoZkConfig struct {
	Zk              string           `json:"zk"`
	ZkFutureTimeout *ZkFutureTimeout `json:"zk_future_timeout"`
	ZkHosts         string           `json:"zk_hosts"`
	ZkPath          string           `json:"zk_path"`
	ZkState         string           `json:"zk_state"`
	ZkTimeout       int              `json:"zk_timeout"`
}

type ZkFutureTimeout struct {
	Duration int `json:"duration"`
}
