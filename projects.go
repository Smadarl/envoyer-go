package envoyer

import "time"

type Server struct {
	ID                      int         `json:"id"`
	ProjectID               int         `json:"project_id"`
	UserID                  int         `json:"user_id"`
	Name                    string      `json:"name"`
	ConnectAs               string      `json:"connect_as"`
	IPAddress               string      `json:"ip_address"`
	Port                    string      `json:"port"`
	PhpSeven                bool        `json:"php_seven"`
	PhpVersion              string      `json:"php_version"`
	Freebsd                 bool        `json:"freebsd"`
	ReceivesCodeDeployments bool        `json:"receives_code_deployments"`
	ShouldRestartFpm        bool        `json:"should_restart_fpm"`
	DeploymentPath          string  	`json:"deployment_path"`
	PhpPath                 string      `json:"php_path"`
	ComposerPath            string      `json:"composer_path"`
	PublicKey               string      `json:"public_key"`
	ConnectionStatus        string      `json:"connection_status"`
	CurrentActivity         interface{} `json:"current_activity"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
}

type Folder struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Project struct {
	ID                     int           `json:"id"`
	UserID                 int           `json:"user_id"`
	Version                int           `json:"version"`
	Name                   string        `json:"name"`
	Provider               string        `json:"provider"`
	PlainRepository        string        `json:"plain_repository"`
	Repository             string        `json:"repository"`
	Type                   string        `json:"type"`
	Branch                 string        `json:"branch"`
	PushToDeploy           bool          `json:"push_to_deploy"`
	WebhookID              string	     `json:"webhook_id"`
	Status                 interface{}   `json:"status"`
	ShouldDeployAgain      int           `json:"should_deploy_again"`
	DeploymentStartedAt    interface{}   `json:"deployment_started_at"`
	DeploymentFinishedAt   time.Time     `json:"deployment_finished_at"`
	LastDeploymentStatus   string        `json:"last_deployment_status"`
	DailyDeploys           int           `json:"daily_deploys"`
	WeeklyDeploys          int           `json:"weekly_deploys"`
	LastDeploymentTook     int           `json:"last_deployment_took"`
	RetainDeployments      int           `json:"retain_deployments"`
	EnvironmentServers     []int 		 `json:"environment_servers"`
	Folders                []Folder		 `json:"folders"`
	Monitor                string        `json:"monitor"`
	NewYorkStatus          string        `json:"new_york_status"`
	LondonStatus           string        `json:"london_status"`
	SingaporeStatus        string        `json:"singapore_status"`
	Token                  string        `json:"token"`
	CreatedAt              time.Time     `json:"created_at"`
	UpdatedAt              time.Time     `json:"updated_at"`
	InstallDevDependencies bool          `json:"install_dev_dependencies"`
	InstallDependencies    bool          `json:"install_dependencies"`
	QuietComposer          bool          `json:"quiet_composer"`
	Servers                []Server		 `json:"servers"`
	HasEnvironment          bool         `json:"has_environment"`
	HasMonitoringError      bool         `json:"has_monitoring_error"`
	HasMissingHeartbeats    bool         `json:"has_missing_heartbeats"`
	LastDeployedBranch      string		 `json:"last_deployed_branch"`
	LastDeploymentID        int			 `json:"last_deployment_id"`
	LastDeploymentAuthor    string		 `json:"last_deployment_author"`
	LastDeploymentAvatar    string		 `json:"last_deployment_avatar"`
	LastDeploymentHash      string       `json:"last_deployment_hash"`
	LastDeploymentTimestamp string		 `json:"last_deployment_timestamp"`
	CanBeDeployed           bool         `json:"can_be_deployed"`
}

type ProjectResponse struct {
	Project `json:"project"`
}

type Projects []Project

type ProjectsResponse struct {
	Projects `json:"projects"`
}