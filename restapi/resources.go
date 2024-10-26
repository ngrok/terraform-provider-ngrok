// Code generated for API Clients. DO NOT EDIT.

package restapi

type Empty struct {
}

type Item struct {
	// a resource identifier
	ID string `json:"id,omitempty"`
}

type Paging struct {
	BeforeID *string `json:"before_id,omitempty"`
	Limit    *string `json:"limit,omitempty"`
}

type Error struct {
	ErrorCode  string            `json:"error_code,omitempty"`
	StatusCode int32             `json:"status_code,omitempty"`
	Msg        string            `json:"msg,omitempty"`
	Details    map[string]string `json:"details,omitempty"`
}

type Ref struct {
	// a resource identifier
	ID string `json:"id,omitempty"`
	// a uri for locating a resource
	URI string `json:"uri,omitempty"`
}

type AbuseReport struct {
	// ID of the abuse report
	ID string `json:"id,omitempty"`
	// URI of the abuse report API resource
	URI string `json:"uri,omitempty"`
	// timestamp that the abuse report record was created in RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// a list of URLs containing suspected abusive content
	URLs []string `json:"urls,omitempty"`
	// arbitrary user-defined data about this abuse report. Optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// Indicates whether ngrok has processed the abuse report. one of PENDING,
	// PROCESSED, or PARTIALLY_PROCESSED
	Status string `json:"status,omitempty"`
	// an array of hostname statuses related to the report
	Hostnames []AbuseReportHostname `json:"hostnames,omitempty"`
}

type AbuseReportHostname struct {
	// the hostname ngrok has parsed out of one of the reported URLs in this abuse
	// report
	Hostname string `json:"hostname,omitempty"`
	// indicates what action ngrok has taken against the hostname. one of PENDING,
	// BANNED, UNBANNED, or IGNORE
	Status string `json:"status,omitempty"`
}

type AbuseReportCreate struct {
	// a list of URLs containing suspected abusive content
	URLs []string `json:"urls,omitempty"`
	// arbitrary user-defined data about this abuse report. Optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
}

type AgentIngressCreate struct {
	// human-readable description of the use of this Agent Ingress. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Agent Ingress. optional,
	// max 4096 bytes
	Metadata string `json:"metadata,omitempty"`
	// the domain that you own to be used as the base domain name to generate regional
	// agent ingress domains.
	Domain string `json:"domain,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled. Optional.
	CertificateManagementPolicy *AgentIngressCertPolicy `json:"certificate_management_policy,omitempty"`
}

type AgentIngressUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of the use of this Agent Ingress. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Agent Ingress. optional,
	// max 4096 bytes
	Metadata *string `json:"metadata,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled. Optional.
	CertificateManagementPolicy *AgentIngressCertPolicy `json:"certificate_management_policy,omitempty"`
}

type AgentIngress struct {
	// unique Agent Ingress resource identifier
	ID string `json:"id,omitempty"`
	// URI to the API resource of this Agent ingress
	URI string `json:"uri,omitempty"`
	// human-readable description of the use of this Agent Ingress. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Agent Ingress. optional,
	// max 4096 bytes
	Metadata string `json:"metadata,omitempty"`
	// the domain that you own to be used as the base domain name to generate regional
	// agent ingress domains.
	Domain string `json:"domain,omitempty"`
	// a list of target values to use as the values of NS records for the domain
	// property these values will delegate control over the domain to ngrok
	NSTargets []string `json:"ns_targets,omitempty"`
	// a list of regional agent ingress domains that are subdomains of the value of
	// domain this value may increase over time as ngrok adds more regions
	RegionDomains []string `json:"region_domains,omitempty"`
	// timestamp when the Agent Ingress was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled
	CertificateManagementPolicy *AgentIngressCertPolicy `json:"certificate_management_policy,omitempty"`
	// status of the automatic certificate management for this domain, or null if
	// automatic management is disabled
	CertificateManagementStatus *AgentIngressCertStatus `json:"certificate_management_status,omitempty"`
}

type AgentIngressList struct {
	// the list of Agent Ingresses owned by this account
	Ingresses []AgentIngress `json:"ingresses,omitempty"`
	// URI of the Agent Ingress list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type AgentIngressCertPolicy struct {
	// certificate authority to request certificates from. The only supported value is
	// letsencrypt.
	Authority string `json:"authority,omitempty"`
	// type of private key to use when requesting certificates. Defaults to rsa, can be
	// either rsa or ecdsa.
	PrivateKeyType string `json:"private_key_type,omitempty"`
}

type AgentIngressCertStatus struct {
	// timestamp when the next renewal will be requested, RFC 3339 format
	RenewsAt *string `json:"renews_at,omitempty"`
	// status of the certificate provisioning job, or null if the certificiate isn't
	// being provisioned or renewed
	ProvisioningJob *AgentIngressCertJob `json:"provisioning_job,omitempty"`
}

type AgentIngressCertJob struct {
	// if present, an error code indicating why provisioning is failing. It may be
	// either a temporary condition (INTERNAL_ERROR), or a permanent one the user must
	// correct (DNS_ERROR).
	ErrorCode *string `json:"error_code,omitempty"`
	// a message describing the current status or error
	Msg string `json:"msg,omitempty"`
	// timestamp when the provisioning job started, RFC 3339 format
	StartedAt string `json:"started_at,omitempty"`
	// timestamp when the provisioning job will be retried
	RetriesAt *string `json:"retries_at,omitempty"`
}

type APIKeyCreate struct {
	// human-readable description of what uses the API key to authenticate. optional,
	// max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined data of this API key. optional, max 4096 bytes
	Metadata string `json:"metadata,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User. Only admins may specify an owner other than themselves. Both owner_id and
	// owner_email may not be specified.
	OwnerEmail string `json:"owner_email,omitempty"`
}

type APIKeyUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of what uses the API key to authenticate. optional,
	// max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined data of this API key. optional, max 4096 bytes
	Metadata *string `json:"metadata,omitempty"`
}

type APIKey struct {
	// unique API key resource identifier
	ID string `json:"id,omitempty"`
	// URI to the API resource of this API key
	URI string `json:"uri,omitempty"`
	// human-readable description of what uses the API key to authenticate. optional,
	// max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined data of this API key. optional, max 4096 bytes
	Metadata string `json:"metadata,omitempty"`
	// timestamp when the api key was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// the bearer token that can be placed into the Authorization header to
	// authenticate request to the ngrok API. This value is only available one time, on
	// the API response from key creation. Otherwise it is null.
	Token *string `json:"token,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
}

type APIKeyList struct {
	// the list of API keys for this account
	Keys []APIKey `json:"keys,omitempty"`
	// URI of the API keys list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type ApplicationSession struct {
	// unique application session resource identifier
	ID string `json:"id,omitempty"`
	// URI of the application session API resource
	URI string `json:"uri,omitempty"`
	// URL of the hostport served by this endpoint
	PublicURL string `json:"public_url,omitempty"`
	// browser session details of the application session
	BrowserSession BrowserSession `json:"browser_session,omitempty"`
	// application user this session is associated with
	ApplicationUser *Ref `json:"application_user,omitempty"`
	// timestamp when the user was created in RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// timestamp when the user was last active in RFC 3339 format
	LastActive string `json:"last_active,omitempty"`
	// timestamp when session expires in RFC 3339 format
	ExpiresAt string `json:"expires_at,omitempty"`
	// ephemeral endpoint this session is associated with
	Endpoint *Ref `json:"endpoint,omitempty"`
	// edge this session is associated with, null if the endpoint is agent-initiated
	Edge *Ref `json:"edge,omitempty"`
	// route this session is associated with, null if the endpoint is agent-initiated
	Route *Ref `json:"route,omitempty"`
}

type ApplicationSessionList struct {
	// list of all application sessions on this account
	ApplicationSessions []ApplicationSession `json:"application_sessions,omitempty"`
	// URI of the application session list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type BrowserSession struct {
	// HTTP User-Agent data
	UserAgent UserAgent `json:"user_agent,omitempty"`
	// IP address
	IPAddress string `json:"ip_address,omitempty"`
	// IP geolocation data
	Location *Location `json:"location,omitempty"`
}

type UserAgent struct {
	// raw User-Agent request header
	Raw string `json:"raw,omitempty"`
	// browser name (e.g. Chrome)
	BrowserName string `json:"browser_name,omitempty"`
	// browser version (e.g. 102)
	BrowserVersion string `json:"browser_version,omitempty"`
	// type of device (e.g. Desktop)
	DeviceType string `json:"device_type,omitempty"`
	// operating system name (e.g. MacOS)
	OSName string `json:"os_name,omitempty"`
	// operating system version (e.g. 10.15.7)
	OSVersion string `json:"os_version,omitempty"`
}

type Location struct {
	// ISO country code
	CountryCode *string `json:"country_code,omitempty"`
	// geographical latitude
	Latitude *float64 `json:"latitude,omitempty"`
	// geographical longitude
	Longitude *float64 `json:"longitude,omitempty"`
	// accuracy radius of the geographical coordinates
	LatLongRadiusKm *uint64 `json:"lat_long_radius_km,omitempty"`
}

type ApplicationUser struct {
	// unique application user resource identifier
	ID string `json:"id,omitempty"`
	// URI of the application user API resource
	URI string `json:"uri,omitempty"`
	// identity provider that the user authenticated with
	IdentityProvider IdentityProvider `json:"identity_provider,omitempty"`
	// unique user identifier
	ProviderUserID string `json:"provider_user_id,omitempty"`
	// user username
	Username string `json:"username,omitempty"`
	// user email
	Email string `json:"email,omitempty"`
	// user common name
	Name string `json:"name,omitempty"`
	// timestamp when the user was created in RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// timestamp when the user was last active in RFC 3339 format
	LastActive string `json:"last_active,omitempty"`
	// timestamp when the user last signed-in in RFC 3339 format
	LastLogin string `json:"last_login,omitempty"`
}

type ApplicationUserList struct {
	// list of all application users on this account
	ApplicationUsers []ApplicationUser `json:"application_users,omitempty"`
	// URI of the application user list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type IdentityProvider struct {
	// name of the identity provider (e.g. Google)
	Name string `json:"name,omitempty"`
	// URL of the identity provider (e.g. https://accounts.google.com
	// (https://accounts.google.com))
	URL string `json:"url,omitempty"`
}

type TunnelSession struct {
	// version of the ngrok agent that started this ngrok tunnel session
	AgentVersion string `json:"agent_version,omitempty"`
	// reference to the tunnel credential or ssh credential used by the ngrok agent to
	// start this tunnel session
	Credential Ref `json:"credential,omitempty"`
	// unique tunnel session resource identifier
	ID string `json:"id,omitempty"`
	// source ip address of the tunnel session
	IP string `json:"ip,omitempty"`
	// arbitrary user-defined data specified in the metadata property in the ngrok
	// configuration file. See the metadata configuration option
	Metadata string `json:"metadata,omitempty"`
	// operating system of the host the ngrok agent is running on
	OS string `json:"os,omitempty"`
	// the ngrok region identifier in which this tunnel session was started
	Region string `json:"region,omitempty"`
	// time when the tunnel session first connected to the ngrok servers
	StartedAt string `json:"started_at,omitempty"`
	// the transport protocol used to start the tunnel session. Either ngrok/v2 or ssh
	Transport string `json:"transport,omitempty"`
	// URI to the API resource of the tunnel session
	URI string `json:"uri,omitempty"`
}

type AgentVersionDeprecated struct {
	NextMin  string `json:"next_min,omitempty"`
	NextDate string `json:"next_date,omitempty"`
	Msg      string `json:"msg,omitempty"`
}

// The User-Agent for a tunnel session
type TunnelSessionUserAgent struct {
	// The list of products making up the User-Agent
	Products []TunnelSessionProduct `json:"products,omitempty"`
}

// A product entry in a tunnel session User-Agent
type TunnelSessionProduct struct {
	// The product name
	Name string `json:"name,omitempty"`
	// The product version
	Version string `json:"version,omitempty"`
	// An optional comment
	Comment string `json:"comment,omitempty"`
}

type TunnelSessionList struct {
	// list of all tunnel sessions on this account
	TunnelSessions []TunnelSession `json:"tunnel_sessions,omitempty"`
	// URI to the API resource of the tunnel session list
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type TunnelSessionsUpdate struct {
	ID string `json:"id,omitempty"`
	// request that the ngrok agent update to this specific version instead of the
	// latest available version
	Version string `json:"version,omitempty"`
}

type AuditEventDashLogin struct {
	AccountID  string `json:"account_id,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	RemoteAddr string `json:"remote_addr,omitempty"`
	Email      string `json:"email,omitempty"`
}

type AuditEventEndpoint struct {
	ID                    string `json:"id,omitempty"`
	AccountID             string `json:"account_id,omitempty"`
	Region                string `json:"region,omitempty"`
	IngressRegion         string `json:"ingress_region,omitempty"`
	CreatedAt             string `json:"created_at,omitempty"`
	UpdatedAt             string `json:"updated_at,omitempty"`
	URL                   string `json:"url,omitempty"`
	Proto                 string `json:"proto,omitempty"`
	DomainID              string `json:"domain_id,omitempty"`
	TCPAddrID             string `json:"tcp_addr_id,omitempty"`
	Rank                  int32  `json:"rank,omitempty"`
	StaticTunnelID        string `json:"static_tunnel_id,omitempty"`
	StaticTunnelRegion    string `json:"static_tunnel_region,omitempty"`
	StaticTunnelSessionID string `json:"static_tunnel_session_id,omitempty"`
	EdgeID                string `json:"edge_id,omitempty"`
	Type                  string `json:"type,omitempty"`
}

type AuditEventTunnelSession struct {
	ID                   string                 `json:"id,omitempty"`
	RemoteAddr           string                 `json:"remote_addr,omitempty"`
	CreatedAt            string                 `json:"created_at,omitempty"`
	AgentVersion         string                 `json:"agent_version,omitempty"`
	Transport            string                 `json:"transport,omitempty"`
	Os                   string                 `json:"os,omitempty"`
	Arch                 string                 `json:"arch,omitempty"`
	RegionID             string                 `json:"region_id,omitempty"`
	CredID               string                 `json:"cred_id,omitempty"`
	SSHCredID            string                 `json:"ssh_cred_id,omitempty"`
	AgentIngressHostname string                 `json:"agent_ingress_hostname,omitempty"`
	ProxyType            string                 `json:"proxy_type,omitempty"`
	MutualTls            bool                   `json:"mutual_tls,omitempty"`
	ServiceRun           bool                   `json:"service_run,omitempty"`
	ConfigVersion        string                 `json:"config_version,omitempty"`
	CustomCas            bool                   `json:"custom_cas,omitempty"`
	ClientType           string                 `json:"client_type,omitempty"`
	UserAgent            TunnelSessionUserAgent `json:"user_agent,omitempty"`
	Deprecated           AgentVersionDeprecated `json:"deprecated,omitempty"`
}

type AuditEventTunnel struct {
	ID           string            `json:"id,omitempty"`
	CreatedAt    string            `json:"created_at,omitempty"`
	DeletedAt    string            `json:"deleted_at,omitempty"`
	AccountID    string            `json:"account_id,omitempty"`
	SessionID    string            `json:"session_id,omitempty"`
	RemoteAddr   string            `json:"remote_addr,omitempty"`
	RegionID     string            `json:"region_id,omitempty"`
	AgentVersion string            `json:"agent_version,omitempty"`
	ForwardsTo   string            `json:"forwards_to,omitempty"`
	EndpointID   string            `json:"endpoint_id,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	BackendIDs   []string          `json:"backend_ids,omitempty"`
	Hostname     string            `json:"hostname,omitempty"`
	Port         int64             `json:"port,omitempty"`
	Proto        string            `json:"proto,omitempty"`
}

type FailoverBackend struct {
	// unique identifier for this Failover backend
	ID string `json:"id,omitempty"`
	// URI of the FailoverBackend API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the backend was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// the ids of the child backends in order
	Backends []string `json:"backends,omitempty"`
}

type FailoverBackendCreate struct {
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// the ids of the child backends in order
	Backends []string `json:"backends,omitempty"`
}

type FailoverBackendUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this backend. Optional
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata *string `json:"metadata,omitempty"`
	// the ids of the child backends in order
	Backends []string `json:"backends,omitempty"`
}

type FailoverBackendList struct {
	// the list of all Failover backends on this account
	Backends []FailoverBackend `json:"backends,omitempty"`
	// URI of the Failover backends list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type HTTPResponseBackend struct {
	ID string `json:"id,omitempty"`
	// URI of the HTTPResponseBackend API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the backend was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// body to return as fixed content
	Body string `json:"body,omitempty"`
	// headers to return
	Headers map[string]string `json:"headers,omitempty"`
	// status code to return
	StatusCode int32 `json:"status_code,omitempty"`
}

type HTTPResponseBackendCreate struct {
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// body to return as fixed content
	Body string `json:"body,omitempty"`
	// headers to return
	Headers map[string]string `json:"headers,omitempty"`
	// status code to return
	StatusCode *int32 `json:"status_code,omitempty"`
}

type HTTPResponseBackendUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this backend. Optional
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata *string `json:"metadata,omitempty"`
	// body to return as fixed content
	Body *string `json:"body,omitempty"`
	// headers to return
	Headers *map[string]string `json:"headers,omitempty"`
	// status code to return
	StatusCode *int32 `json:"status_code,omitempty"`
}

type HTTPResponseBackendList struct {
	Backends    []HTTPResponseBackend `json:"backends,omitempty"`
	URI         string                `json:"uri,omitempty"`
	NextPageURI *string               `json:"next_page_uri,omitempty"`
}

type StaticBackend struct {
	// unique identifier for this static backend
	ID string `json:"id,omitempty"`
	// URI of the StaticBackend API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the backend was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// the address to forward to
	Address string `json:"address,omitempty"`
	// tls configuration to use
	TLS StaticBackendTLS `json:"tls,omitempty"`
}

type StaticBackendTLS struct {
	// if TLS is checked
	Enabled bool `json:"enabled,omitempty"`
}

type StaticBackendCreate struct {
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// the address to forward to
	Address string `json:"address,omitempty"`
	// tls configuration to use
	TLS StaticBackendTLS `json:"tls,omitempty"`
}

type StaticBackendUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this backend. Optional
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata *string `json:"metadata,omitempty"`
	// the address to forward to
	Address string `json:"address,omitempty"`
	// tls configuration to use
	TLS StaticBackendTLS `json:"tls,omitempty"`
}

type StaticBackendList struct {
	// the list of all static backends on this account
	Backends []StaticBackend `json:"backends,omitempty"`
	// URI of the static backends list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type TunnelGroupBackend struct {
	// unique identifier for this TunnelGroup backend
	ID string `json:"id,omitempty"`
	// URI of the TunnelGroupBackend API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the backend was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// labels to watch for tunnels on, e.g. app->foo, dc->bar
	Labels map[string]string `json:"labels,omitempty"`
	// tunnels matching this backend
	Tunnels []Ref `json:"tunnels,omitempty"`
}

type TunnelGroupBackendCreate struct {
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// labels to watch for tunnels on, e.g. app->foo, dc->bar
	Labels map[string]string `json:"labels,omitempty"`
}

type TunnelGroupBackendUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this backend. Optional
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata *string `json:"metadata,omitempty"`
	// labels to watch for tunnels on, e.g. app->foo, dc->bar
	Labels map[string]string `json:"labels,omitempty"`
}

type TunnelGroupBackendList struct {
	// the list of all TunnelGroup backends on this account
	Backends []TunnelGroupBackend `json:"backends,omitempty"`
	// URI of the TunnelGroup backends list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type WeightedBackend struct {
	// unique identifier for this Weighted backend
	ID string `json:"id,omitempty"`
	// URI of the WeightedBackend API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the backend was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// the ids of the child backends to their weights [0-10000]
	Backends map[string]int64 `json:"backends,omitempty"`
}

type WeightedBackendCreate struct {
	// human-readable description of this backend. Optional
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata string `json:"metadata,omitempty"`
	// the ids of the child backends to their weights [0-10000]
	Backends map[string]int64 `json:"backends,omitempty"`
}

type WeightedBackendUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this backend. Optional
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this backend. Optional
	Metadata *string `json:"metadata,omitempty"`
	// the ids of the child backends to their weights [0-10000]
	Backends map[string]int64 `json:"backends,omitempty"`
}

type WeightedBackendList struct {
	// the list of all Weighted backends on this account
	Backends []WeightedBackend `json:"backends,omitempty"`
	// URI of the Weighted backends list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type BotUser struct {
	// unique API key resource identifier
	ID string `json:"id,omitempty"`
	// URI to the API resource of this bot user
	URI string `json:"uri,omitempty"`
	// human-readable name used to identify the bot
	Name string `json:"name,omitempty"`
	// whether or not the bot is active
	Active bool `json:"active,omitempty"`
	// timestamp when the api key was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
}

type BotUserCreate struct {
	// human-readable name used to identify the bot
	Name string `json:"name,omitempty"`
	// whether or not the bot is active
	Active *bool `json:"active,omitempty"`
}

type BotUserUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable name used to identify the bot
	Name *string `json:"name,omitempty"`
	// whether or not the bot is active
	Active *bool `json:"active,omitempty"`
}

type BotUserList struct {
	// the list of all bot users on this account
	BotUsers []BotUser `json:"bot_users,omitempty"`
	// URI of the bot users list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type CertificateAuthorityCreate struct {
	// human-readable description of this Certificate Authority. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Certificate Authority.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// raw PEM of the Certificate Authority
	CAPEM string `json:"ca_pem,omitempty"`
}

type CertificateAuthorityUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this Certificate Authority. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Certificate Authority.
	// optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

type CertificateAuthority struct {
	// unique identifier for this Certificate Authority
	ID string `json:"id,omitempty"`
	// URI of the Certificate Authority API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the Certificate Authority was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this Certificate Authority. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Certificate Authority.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// raw PEM of the Certificate Authority
	CAPEM string `json:"ca_pem,omitempty"`
	// subject common name of the Certificate Authority
	SubjectCommonName string `json:"subject_common_name,omitempty"`
	// timestamp when this Certificate Authority becomes valid, RFC 3339 format
	NotBefore string `json:"not_before,omitempty"`
	// timestamp when this Certificate Authority becomes invalid, RFC 3339 format
	NotAfter string `json:"not_after,omitempty"`
	// set of actions the private key of this Certificate Authority can be used for
	KeyUsages []string `json:"key_usages,omitempty"`
	// extended set of actions the private key of this Certificate Authority can be
	// used for
	ExtendedKeyUsages []string `json:"extended_key_usages,omitempty"`
}

type CertificateAuthorityList struct {
	// the list of all certificate authorities on this account
	CertificateAuthorities []CertificateAuthority `json:"certificate_authorities,omitempty"`
	// URI of the certificates authorities list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type CredentialCreate struct {
	// human-readable description of who or what will use the credential to
	// authenticate. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this credential. Optional, max
	// 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL []string `json:"acl,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User. Only admins may specify an owner other than themselves. Both owner_id and
	// owner_email may not be specified.
	OwnerEmail string `json:"owner_email,omitempty"`
	// Only authorized accounts may supply a pre-computed token that will be associated
	// with the created credentials.
	PrecomputedToken *string `json:"precomputed_token,omitempty"`
}

type CredentialUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of who or what will use the credential to
	// authenticate. Optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this credential. Optional, max
	// 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL *[]string `json:"acl,omitempty"`
}

type Credential struct {
	// unique tunnel credential resource identifier
	ID string `json:"id,omitempty"`
	// URI of the tunnel credential API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the tunnel credential was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of who or what will use the credential to
	// authenticate. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this credential. Optional, max
	// 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// the credential's authtoken that can be used to authenticate an ngrok agent. This
	// value is only available one time, on the API response from credential creation,
	// otherwise it is null.
	Token *string `json:"token,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL []string `json:"acl,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
}

type CredentialList struct {
	// the list of all tunnel credentials on this account
	Credentials []Credential `json:"credentials,omitempty"`
	// URI of the tunnel credential list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type EndpointConfiguration struct {
	// unique identifier of this endpoint configuration
	ID string `json:"id,omitempty"`
	// they type of traffic this endpoint configuration can be applied to. one of:
	// http, https, tcp
	Type string `json:"type,omitempty"`
	// human-readable description of what this endpoint configuration will be do when
	// applied or what traffic it will be applied to. Optional, max 255 bytes
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this endpoint configuration.
	// Optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// timestamp when the endpoint configuration was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// URI of the endpoint configuration API resource
	URI string `json:"uri,omitempty"`
	// basic auth module configuration or null
	BasicAuth *EndpointBasicAuth `json:"basic_auth,omitempty"`
	// circuit breaker module configuration or null
	CircuitBreaker *EndpointCircuitBreaker `json:"circuit_breaker,omitempty"`
	// compression module configuration or null
	Compression *EndpointCompression `json:"compression,omitempty"`
	// request headers module configuration or null
	RequestHeaders *EndpointRequestHeaders `json:"request_headers,omitempty"`
	// response headers module configuration or null
	ResponseHeaders *EndpointResponseHeaders `json:"response_headers,omitempty"`
	// ip policy module configuration or null
	IPPolicy *EndpointIPPolicy `json:"ip_policy,omitempty"`
	// mutual TLS module configuration or null
	MutualTLS *EndpointMutualTLS `json:"mutual_tls,omitempty"`
	// TLS termination module configuration or null
	TLSTermination *EndpointTLSTermination `json:"tls_termination,omitempty"`
	// webhook validation module configuration or null
	WebhookValidation *EndpointWebhookValidation `json:"webhook_validation,omitempty"`
	// oauth module configuration or null
	OAuth *EndpointOAuth `json:"oauth,omitempty"`
	// saml module configuration or null
	SAML *EndpointSAML `json:"saml,omitempty"`
	// oidc module configuration or null
	OIDC *EndpointOIDC `json:"oidc,omitempty"`
	// backend module configuration or null
	Backend *EndpointBackend `json:"backend,omitempty"`
}

type EndpointConfigurationList struct {
	// the list of all endpoint configurations on this account
	EndpointConfigurations []EndpointConfiguration `json:"endpoint_configurations,omitempty"`
	// URI of the endpoint configurations list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type EndpointConfigurationUpdate struct {
	// unique identifier of this endpoint configuration
	ID string `json:"id,omitempty"`
	// human-readable description of what this endpoint configuration will be do when
	// applied or what traffic it will be applied to. Optional, max 255 bytes
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this endpoint configuration.
	// Optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// basic auth module configuration or null
	BasicAuth *EndpointBasicAuth `json:"basic_auth,omitempty"`
	// circuit breaker module configuration or null
	CircuitBreaker *EndpointCircuitBreaker `json:"circuit_breaker,omitempty"`
	// compression module configuration or null
	Compression *EndpointCompression `json:"compression,omitempty"`
	// request headers module configuration or null
	RequestHeaders *EndpointRequestHeaders `json:"request_headers,omitempty"`
	// response headers module configuration or null
	ResponseHeaders *EndpointResponseHeaders `json:"response_headers,omitempty"`
	// ip policy module configuration or null
	IPPolicy *EndpointIPPolicyMutate `json:"ip_policy,omitempty"`
	// mutual TLS module configuration or null
	MutualTLS *EndpointMutualTLSMutate `json:"mutual_tls,omitempty"`
	// TLS termination module configuration or null
	TLSTermination *EndpointTLSTermination `json:"tls_termination,omitempty"`
	// webhook validation module configuration or null
	WebhookValidation *EndpointWebhookValidation `json:"webhook_validation,omitempty"`
	// oauth module configuration or null
	OAuth *EndpointOAuth `json:"oauth,omitempty"`
	// saml module configuration or null
	SAML *EndpointSAMLMutate `json:"saml,omitempty"`
	// oidc module configuration or null
	OIDC *EndpointOIDC `json:"oidc,omitempty"`
	// backend module configuration or null
	Backend *EndpointBackendMutate `json:"backend,omitempty"`
}

type EndpointConfigurationCreate struct {
	// they type of traffic this endpoint configuration can be applied to. one of:
	// http, https, tcp
	Type string `json:"type,omitempty"`
	// human-readable description of what this endpoint configuration will be do when
	// applied or what traffic it will be applied to. Optional, max 255 bytes
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this endpoint configuration.
	// Optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// basic auth module configuration or null
	BasicAuth *EndpointBasicAuth `json:"basic_auth,omitempty"`
	// circuit breaker module configuration or null
	CircuitBreaker *EndpointCircuitBreaker `json:"circuit_breaker,omitempty"`
	// compression module configuration or null
	Compression *EndpointCompression `json:"compression,omitempty"`
	// request headers module configuration or null
	RequestHeaders *EndpointRequestHeaders `json:"request_headers,omitempty"`
	// response headers module configuration or null
	ResponseHeaders *EndpointResponseHeaders `json:"response_headers,omitempty"`
	// ip policy module configuration or null
	IPPolicy *EndpointIPPolicyMutate `json:"ip_policy,omitempty"`
	// mutual TLS module configuration or null
	MutualTLS *EndpointMutualTLSMutate `json:"mutual_tls,omitempty"`
	// TLS termination module configuration or null
	TLSTermination *EndpointTLSTermination `json:"tls_termination,omitempty"`
	// webhook validation module configuration or null
	WebhookValidation *EndpointWebhookValidation `json:"webhook_validation,omitempty"`
	// oauth module configuration or null
	OAuth *EndpointOAuth `json:"oauth,omitempty"`
	// saml module configuration or null
	SAML *EndpointSAMLMutate `json:"saml,omitempty"`
	// oidc module configuration or null
	OIDC *EndpointOIDC `json:"oidc,omitempty"`
	// backend module configuration or null
	Backend *EndpointBackendMutate `json:"backend,omitempty"`
}

type EndpointWebhookValidation struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// a string indicating which webhook provider will be sending webhooks to this
	// endpoint. Value must be one of the supported providers defined at
	// https://ngrok.com/docs/cloud-edge/modules/webhook-verification
	// (https://ngrok.com/docs/cloud-edge/modules/webhook-verification)
	Provider string `json:"provider,omitempty"`
	// a string secret used to validate requests from the given provider. All providers
	// except AWS SNS require a secret
	Secret string `json:"secret,omitempty"`
}

type EndpointCompression struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
}

type EndpointMutualTLS struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// PEM-encoded CA certificates that will be used to validate. Multiple CAs may be
	// provided by concatenating them together.
	CertificateAuthorities []Ref `json:"certificate_authorities,omitempty"`
}

type EndpointMutualTLSMutate struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// list of certificate authorities that will be used to validate the TLS client
	// certificate presented by the initiator of the TLS connection
	CertificateAuthorityIDs []string `json:"certificate_authority_ids,omitempty"`
}

type EndpointTLSTermination struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// edge if the ngrok edge should terminate TLS traffic, upstream if TLS traffic
	// should be passed through to the upstream ngrok agent / application server for
	// termination. if upstream is chosen, most other modules will be disallowed
	// because they rely on the ngrok edge being able to access the underlying traffic.
	TerminateAt string `json:"terminate_at,omitempty"`
	// The minimum TLS version used for termination and advertised to the client during
	// the TLS handshake. if unspecified, ngrok will choose an industry-safe default.
	// This value must be null if terminate_at is set to upstream.
	MinVersion *string `json:"min_version,omitempty"`
}

type EndpointTLSTerminationAtEdge struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// The minimum TLS version used for termination and advertised to the client during
	// the TLS handshake. if unspecified, ngrok will choose an industry-safe default.
	// This value must be null if terminate_at is set to upstream.
	MinVersion *string `json:"min_version,omitempty"`
}

type EndpointBasicAuth struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// determines how the basic auth credentials are validated. Currently only the
	// value agent is supported which means that credentials will be validated against
	// the username and password specified by the ngrok agent's --basic-auth flag, if
	// any.
	AuthProviderID string `json:"auth_provider_id,omitempty"`
	// an arbitrary string to be specified in as the 'realm' value in the
	// WWW-Authenticate header. default is ngrok
	Realm string `json:"realm,omitempty"`
	// true or false indicating whether to allow OPTIONS requests through without
	// authentication which is necessary for CORS. default is false
	AllowOptions bool `json:"allow_options,omitempty"`
}

type EndpointRequestHeaders struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// a map of header key to header value that will be injected into the HTTP Request
	// before being sent to the upstream application server
	Add map[string]string `json:"add,omitempty"`
	// a list of header names that will be removed from the HTTP Request before being
	// sent to the upstream application server
	Remove []string `json:"remove,omitempty"`
}

type EndpointResponseHeaders struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// a map of header key to header value that will be injected into the HTTP Response
	// returned to the HTTP client
	Add map[string]string `json:"add,omitempty"`
	// a list of header names that will be removed from the HTTP Response returned to
	// the HTTP client
	Remove []string `json:"remove,omitempty"`
}

type EndpointIPPolicy struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// list of all IP policies that will be used to check if a source IP is allowed
	// access to the endpoint
	IPPolicies []Ref `json:"ip_policies,omitempty"`
}

type EndpointIPPolicyMutate struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// list of all IP policies that will be used to check if a source IP is allowed
	// access to the endpoint
	IPPolicyIDs []string `json:"ip_policy_ids,omitempty"`
}

type EndpointCircuitBreaker struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// Integer number of seconds after which the circuit is tripped to wait before
	// re-evaluating upstream health
	TrippedDuration uint32 `json:"tripped_duration,omitempty"`
	// Integer number of seconds in the statistical rolling window that metrics are
	// retained for.
	RollingWindow uint32 `json:"rolling_window,omitempty"`
	// Integer number of buckets into which metrics are retained. Max 128.
	NumBuckets uint32 `json:"num_buckets,omitempty"`
	// Integer number of requests in a rolling window that will trip the circuit.
	// Helpful if traffic volume is low.
	VolumeThreshold uint32 `json:"volume_threshold,omitempty"`
	// Error threshold percentage should be between 0 - 1.0, not 0-100.0
	ErrorThresholdPercentage float64 `json:"error_threshold_percentage,omitempty"`
}

type EndpointOAuth struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// an object which defines the identity provider to use for authentication and
	// configuration for who may access the endpoint
	Provider EndpointOAuthProvider `json:"provider,omitempty"`
	// Do not enforce authentication on HTTP OPTIONS requests. necessary if you are
	// supporting CORS.
	OptionsPassthrough bool `json:"options_passthrough,omitempty"`
	// the prefix of the session cookie that ngrok sets on the http client to cache
	// authentication. default is 'ngrok.'
	CookiePrefix string `json:"cookie_prefix,omitempty"`
	// Integer number of seconds of inactivity after which if the user has not accessed
	// the endpoint, their session will time out and they will be forced to
	// reauthenticate.
	InactivityTimeout uint32 `json:"inactivity_timeout,omitempty"`
	// Integer number of seconds of the maximum duration of an authenticated session.
	// After this period is exceeded, a user must reauthenticate.
	MaximumDuration uint32 `json:"maximum_duration,omitempty"`
	// Integer number of seconds after which ngrok guarantees it will refresh user
	// state from the identity provider and recheck whether the user is still
	// authorized to access the endpoint. This is the preferred tunable to use to
	// enforce a minimum amount of time after which a revoked user will no longer be
	// able to access the resource.
	AuthCheckInterval uint32 `json:"auth_check_interval,omitempty"`
}

type EndpointOAuthProvider struct {
	// configuration for using github as the identity provider
	Github *EndpointOAuthGitHub `json:"github,omitempty"`
	// configuration for using facebook as the identity provider
	Facebook *EndpointOAuthFacebook `json:"facebook,omitempty"`
	// configuration for using microsoft as the identity provider
	Microsoft *EndpointOAuthMicrosoft `json:"microsoft,omitempty"`
	// configuration for using google as the identity provider
	Google *EndpointOAuthGoogle `json:"google,omitempty"`
	// configuration for using linkedin as the identity provider
	Linkedin *EndpointOAuthLinkedIn `json:"linkedin,omitempty"`
	// configuration for using gitlab as the identity provider
	Gitlab *EndpointOAuthGitLab `json:"gitlab,omitempty"`
	// configuration for using twitch as the identity provider
	Twitch *EndpointOAuthTwitch `json:"twitch,omitempty"`
	// configuration for using amazon as the identity provider
	Amazon *EndpointOAuthAmazon `json:"amazon,omitempty"`
}

type EndpointOAuthGitHub struct {
	// the OAuth app client ID. retrieve it from the identity provider's dashboard
	// where you created your own OAuth app. optional. if unspecified, ngrok will use
	// its own managed oauth application which has additional restrictions. see the
	// OAuth module docs for more details. if present, client_secret must be present as
	// well.
	ClientID *string `json:"client_id,omitempty"`
	// the OAuth app client secret. retrieve if from the identity provider's dashboard
	// where you created your own OAuth app. optional, see all of the caveats in the
	// docs for client_id.
	ClientSecret *string `json:"client_secret,omitempty"`
	// a list of provider-specific OAuth scopes with the permissions your OAuth app
	// would like to ask for. these may not be set if you are using the ngrok-managed
	// oauth app (i.e. you must pass both client_id and client_secret to set scopes)
	Scopes *[]string `json:"scopes,omitempty"`
	// a list of email addresses of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailAddresses *[]string `json:"email_addresses,omitempty"`
	// a list of email domains of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailDomains *[]string `json:"email_domains,omitempty"`
	// a list of github teams identifiers. users will be allowed access to the endpoint
	// if they are a member of any of these teams. identifiers should be in the 'slug'
	// format qualified with the org name, e.g. org-name/team-name
	Teams *[]string `json:"teams,omitempty"`
	// a list of github org identifiers. users who are members of any of the listed
	// organizations will be allowed access. identifiers should be the organization's
	// 'slug'
	Organizations *[]string `json:"organizations,omitempty"`
}

type EndpointOAuthFacebook struct {
	// the OAuth app client ID. retrieve it from the identity provider's dashboard
	// where you created your own OAuth app. optional. if unspecified, ngrok will use
	// its own managed oauth application which has additional restrictions. see the
	// OAuth module docs for more details. if present, client_secret must be present as
	// well.
	ClientID *string `json:"client_id,omitempty"`
	// the OAuth app client secret. retrieve if from the identity provider's dashboard
	// where you created your own OAuth app. optional, see all of the caveats in the
	// docs for client_id.
	ClientSecret *string `json:"client_secret,omitempty"`
	// a list of provider-specific OAuth scopes with the permissions your OAuth app
	// would like to ask for. these may not be set if you are using the ngrok-managed
	// oauth app (i.e. you must pass both client_id and client_secret to set scopes)
	Scopes []string `json:"scopes,omitempty"`
	// a list of email addresses of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailAddresses []string `json:"email_addresses,omitempty"`
	// a list of email domains of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailDomains []string `json:"email_domains,omitempty"`
}

type EndpointOAuthMicrosoft struct {
	// the OAuth app client ID. retrieve it from the identity provider's dashboard
	// where you created your own OAuth app. optional. if unspecified, ngrok will use
	// its own managed oauth application which has additional restrictions. see the
	// OAuth module docs for more details. if present, client_secret must be present as
	// well.
	ClientID *string `json:"client_id,omitempty"`
	// the OAuth app client secret. retrieve if from the identity provider's dashboard
	// where you created your own OAuth app. optional, see all of the caveats in the
	// docs for client_id.
	ClientSecret *string `json:"client_secret,omitempty"`
	// a list of provider-specific OAuth scopes with the permissions your OAuth app
	// would like to ask for. these may not be set if you are using the ngrok-managed
	// oauth app (i.e. you must pass both client_id and client_secret to set scopes)
	Scopes []string `json:"scopes,omitempty"`
	// a list of email addresses of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailAddresses []string `json:"email_addresses,omitempty"`
	// a list of email domains of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailDomains []string `json:"email_domains,omitempty"`
}

type EndpointOAuthGoogle struct {
	// the OAuth app client ID. retrieve it from the identity provider's dashboard
	// where you created your own OAuth app. optional. if unspecified, ngrok will use
	// its own managed oauth application which has additional restrictions. see the
	// OAuth module docs for more details. if present, client_secret must be present as
	// well.
	ClientID *string `json:"client_id,omitempty"`
	// the OAuth app client secret. retrieve if from the identity provider's dashboard
	// where you created your own OAuth app. optional, see all of the caveats in the
	// docs for client_id.
	ClientSecret *string `json:"client_secret,omitempty"`
	// a list of provider-specific OAuth scopes with the permissions your OAuth app
	// would like to ask for. these may not be set if you are using the ngrok-managed
	// oauth app (i.e. you must pass both client_id and client_secret to set scopes)
	Scopes []string `json:"scopes,omitempty"`
	// a list of email addresses of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailAddresses []string `json:"email_addresses,omitempty"`
	// a list of email domains of users authenticated by identity provider who are
	// allowed access to the endpoint
	EmailDomains []string `json:"email_domains,omitempty"`
}

type EndpointOAuthLinkedIn struct {
	ClientID       *string  `json:"client_id,omitempty"`
	ClientSecret   *string  `json:"client_secret,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	EmailDomains   []string `json:"email_domains,omitempty"`
}

type EndpointOAuthGitLab struct {
	ClientID       *string  `json:"client_id,omitempty"`
	ClientSecret   *string  `json:"client_secret,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	EmailDomains   []string `json:"email_domains,omitempty"`
}

type EndpointOAuthTwitch struct {
	ClientID       *string  `json:"client_id,omitempty"`
	ClientSecret   *string  `json:"client_secret,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	EmailDomains   []string `json:"email_domains,omitempty"`
}

type EndpointOAuthAmazon struct {
	ClientID       *string  `json:"client_id,omitempty"`
	ClientSecret   *string  `json:"client_secret,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	EmailDomains   []string `json:"email_domains,omitempty"`
}

type EndpointSAML struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// Do not enforce authentication on HTTP OPTIONS requests. necessary if you are
	// supporting CORS.
	OptionsPassthrough bool `json:"options_passthrough,omitempty"`
	// the prefix of the session cookie that ngrok sets on the http client to cache
	// authentication. default is 'ngrok.'
	CookiePrefix string `json:"cookie_prefix,omitempty"`
	// Integer number of seconds of inactivity after which if the user has not accessed
	// the endpoint, their session will time out and they will be forced to
	// reauthenticate.
	InactivityTimeout uint32 `json:"inactivity_timeout,omitempty"`
	// Integer number of seconds of the maximum duration of an authenticated session.
	// After this period is exceeded, a user must reauthenticate.
	MaximumDuration uint32 `json:"maximum_duration,omitempty"`
	// The IdP's metadata URL which returns the XML IdP EntityDescriptor. The IdP's
	// metadata URL specifies how to connect to the IdP as well as its public key which
	// is then used to validate the signature on incoming SAML assertions to the ACS
	// endpoint.
	IdPMetadataURL string `json:"idp_metadata_url,omitempty"`
	// The full XML IdP EntityDescriptor. Your IdP may provide this to you as a a file
	// to download or as a URL.
	IdPMetadata string `json:"idp_metadata,omitempty"`
	// If true, indicates that whenever we redirect a user to the IdP for
	// authentication that the IdP must prompt the user for authentication credentials
	// even if the user already has a valid session with the IdP.
	ForceAuthn bool `json:"force_authn,omitempty"`
	// If true, the IdP may initiate a login directly (e.g. the user does not need to
	// visit the endpoint first and then be redirected). The IdP should set the
	// RelayState parameter to the target URL of the resource they want the user to be
	// redirected to after the SAML login assertion has been processed.
	AllowIdPInitiated *bool `json:"allow_idp_initiated,omitempty"`
	// If present, only users who are a member of one of the listed groups may access
	// the target endpoint.
	AuthorizedGroups []string `json:"authorized_groups,omitempty"`
	// The SP Entity's unique ID. This always takes the form of a URL. In ngrok's
	// implementation, this URL is the same as the metadata URL. This will need to be
	// specified to the IdP as configuration.
	EntityID string `json:"entity_id,omitempty"`
	// The public URL of the SP's Assertion Consumer Service. This is where the IdP
	// will redirect to during an authentication flow. This will need to be specified
	// to the IdP as configuration.
	AssertionConsumerServiceURL string `json:"assertion_consumer_service_url,omitempty"`
	// The public URL of the SP's Single Logout Service. This is where the IdP will
	// redirect to during a single logout flow. This will optionally need to be
	// specified to the IdP as configuration.
	SingleLogoutURL string `json:"single_logout_url,omitempty"`
	// PEM-encoded x.509 certificate of the key pair that is used to sign all SAML
	// requests that the ngrok SP makes to the IdP. Many IdPs do not support request
	// signing verification, but we highly recommend specifying this in the IdP's
	// configuration if it is supported.
	RequestSigningCertificatePEM string `json:"request_signing_certificate_pem,omitempty"`
	// A public URL where the SP's metadata is hosted. If an IdP supports dynamic
	// configuration, this is the URL it can use to retrieve the SP metadata.
	MetadataURL string `json:"metadata_url,omitempty"`
	// Defines the name identifier format the SP expects the IdP to use in its
	// assertions to identify subjects. If unspecified, a default value of
	// urn:oasis:names:tc:SAML:2.0:nameid-format:persistent will be used. A subset of
	// the allowed values enumerated by the SAML specification are supported.
	NameIDFormat string `json:"nameid_format,omitempty"`
}

type EndpointSAMLMutate struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// Do not enforce authentication on HTTP OPTIONS requests. necessary if you are
	// supporting CORS.
	OptionsPassthrough bool `json:"options_passthrough,omitempty"`
	// the prefix of the session cookie that ngrok sets on the http client to cache
	// authentication. default is 'ngrok.'
	CookiePrefix string `json:"cookie_prefix,omitempty"`
	// Integer number of seconds of inactivity after which if the user has not accessed
	// the endpoint, their session will time out and they will be forced to
	// reauthenticate.
	InactivityTimeout uint32 `json:"inactivity_timeout,omitempty"`
	// Integer number of seconds of the maximum duration of an authenticated session.
	// After this period is exceeded, a user must reauthenticate.
	MaximumDuration uint32 `json:"maximum_duration,omitempty"`
	// The IdP's metadata URL which returns the XML IdP EntityDescriptor. The IdP's
	// metadata URL specifies how to connect to the IdP as well as its public key which
	// is then used to validate the signature on incoming SAML assertions to the ACS
	// endpoint.
	IdPMetadataURL string `json:"idp_metadata_url,omitempty"`
	// The full XML IdP EntityDescriptor. Your IdP may provide this to you as a a file
	// to download or as a URL.
	IdPMetadata string `json:"idp_metadata,omitempty"`
	// If true, indicates that whenever we redirect a user to the IdP for
	// authentication that the IdP must prompt the user for authentication credentials
	// even if the user already has a valid session with the IdP.
	ForceAuthn bool `json:"force_authn,omitempty"`
	// If true, the IdP may initiate a login directly (e.g. the user does not need to
	// visit the endpoint first and then be redirected). The IdP should set the
	// RelayState parameter to the target URL of the resource they want the user to be
	// redirected to after the SAML login assertion has been processed.
	AllowIdPInitiated *bool `json:"allow_idp_initiated,omitempty"`
	// If present, only users who are a member of one of the listed groups may access
	// the target endpoint.
	AuthorizedGroups []string `json:"authorized_groups,omitempty"`
	// Defines the name identifier format the SP expects the IdP to use in its
	// assertions to identify subjects. If unspecified, a default value of
	// urn:oasis:names:tc:SAML:2.0:nameid-format:persistent will be used. A subset of
	// the allowed values enumerated by the SAML specification are supported.
	NameIDFormat string `json:"nameid_format,omitempty"`
}

type EndpointOIDC struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// Do not enforce authentication on HTTP OPTIONS requests. necessary if you are
	// supporting CORS.
	OptionsPassthrough bool `json:"options_passthrough,omitempty"`
	// the prefix of the session cookie that ngrok sets on the http client to cache
	// authentication. default is 'ngrok.'
	CookiePrefix string `json:"cookie_prefix,omitempty"`
	// Integer number of seconds of inactivity after which if the user has not accessed
	// the endpoint, their session will time out and they will be forced to
	// reauthenticate.
	InactivityTimeout uint32 `json:"inactivity_timeout,omitempty"`
	// Integer number of seconds of the maximum duration of an authenticated session.
	// After this period is exceeded, a user must reauthenticate.
	MaximumDuration uint32 `json:"maximum_duration,omitempty"`
	// URL of the OIDC "OpenID provider". This is the base URL used for discovery.
	Issuer string `json:"issuer,omitempty"`
	// The OIDC app's client ID and OIDC audience.
	ClientID string `json:"client_id,omitempty"`
	// The OIDC app's client secret.
	ClientSecret string `json:"client_secret,omitempty"`
	// The set of scopes to request from the OIDC identity provider.
	Scopes []string `json:"scopes,omitempty"`
}

type EndpointBackend struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// backend to be used to back this endpoint
	Backend Ref `json:"backend,omitempty"`
}

type EndpointBackendMutate struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// backend to be used to back this endpoint
	BackendID string `json:"backend_id,omitempty"`
}

type EndpointWebsocketTCPConverter struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
}

type EndpointUserAgentFilter struct {
	Enabled              *bool    `json:"enabled,omitempty"`
	UserAgentFilterAllow []string `json:"allow,omitempty"`
	UserAgentFilterDeny  []string `json:"deny,omitempty"`
}

type EndpointPolicy struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// the inbound rules of the traffic policy.
	Inbound []EndpointRule `json:"inbound,omitempty"`
	// the outbound rules on the traffic policy.
	Outbound []EndpointRule `json:"outbound,omitempty"`
}

type EndpointRule struct {
	// cel expressions that filter traffic the policy rule applies to.
	Expressions []string `json:"expressions,omitempty"`
	// the set of actions on a policy rule.
	Actions []EndpointAction `json:"actions,omitempty"`
	// the name of the rule that is part of the traffic policy.
	Name string `json:"name,omitempty"`
}

type EndpointAction struct {
	// the type of action on the policy rule.
	Type string `json:"type,omitempty"`
	// the configuration for the action on the policy rule.
	Config any `json:"config,omitempty"`
}

type EndpointTrafficPolicy struct {
	// true if the module will be applied to traffic, false to disable. default true if
	// unspecified
	Enabled *bool `json:"enabled,omitempty"`
	// the traffic policy that should be applied to the traffic on your endpoint.
	Value string `json:"value,omitempty"`
}

type EdgeRouteItem struct {
	// unique identifier of this edge
	EdgeID string `json:"edge_id,omitempty"`
	// unique identifier of this edge route
	ID string `json:"id,omitempty"`
}

type HTTPSEdgeRouteCreate struct {
	// unique identifier of this edge
	EdgeID string `json:"edge_id,omitempty"`
	// Type of match to use for this route. Valid values are "exact_path" and
	// "path_prefix".
	MatchType string `json:"match_type,omitempty"`
	// Route selector: "/blog" or "example.com" or "example.com/blog"
	Match string `json:"match,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// backend module configuration or null
	Backend *EndpointBackendMutate `json:"backend,omitempty"`
	// ip restriction module configuration or null
	IPRestriction *EndpointIPPolicyMutate `json:"ip_restriction,omitempty"`
	// circuit breaker module configuration or null
	CircuitBreaker *EndpointCircuitBreaker `json:"circuit_breaker,omitempty"`
	// compression module configuration or null
	Compression *EndpointCompression `json:"compression,omitempty"`
	// request headers module configuration or null
	RequestHeaders *EndpointRequestHeaders `json:"request_headers,omitempty"`
	// response headers module configuration or null
	ResponseHeaders *EndpointResponseHeaders `json:"response_headers,omitempty"`
	// webhook verification module configuration or null
	WebhookVerification *EndpointWebhookValidation `json:"webhook_verification,omitempty"`
	// oauth module configuration or null
	OAuth *EndpointOAuth `json:"oauth,omitempty"`
	// saml module configuration or null
	SAML *EndpointSAMLMutate `json:"saml,omitempty"`
	// oidc module configuration or null
	OIDC *EndpointOIDC `json:"oidc,omitempty"`
	// websocket to tcp adapter configuration or null
	WebsocketTCPConverter *EndpointWebsocketTCPConverter `json:"websocket_tcp_converter,omitempty"`
	UserAgentFilter       *EndpointUserAgentFilter       `json:"user_agent_filter,omitempty"`
	// the traffic policy associated with this edge or null
	Policy *EndpointPolicy `json:"policy,omitempty"`
	// the traffic policy associated with this edge or null
	TrafficPolicy *EndpointTrafficPolicy `json:"traffic_policy,omitempty"`
}

type HTTPSEdgeRouteUpdate struct {
	// unique identifier of this edge
	EdgeID string `json:"edge_id,omitempty"`
	// unique identifier of this edge route
	ID string `json:"id,omitempty"`
	// Type of match to use for this route. Valid values are "exact_path" and
	// "path_prefix".
	MatchType string `json:"match_type,omitempty"`
	// Route selector: "/blog" or "example.com" or "example.com/blog"
	Match string `json:"match,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// backend module configuration or null
	Backend *EndpointBackendMutate `json:"backend,omitempty"`
	// ip restriction module configuration or null
	IPRestriction *EndpointIPPolicyMutate `json:"ip_restriction,omitempty"`
	// circuit breaker module configuration or null
	CircuitBreaker *EndpointCircuitBreaker `json:"circuit_breaker,omitempty"`
	// compression module configuration or null
	Compression *EndpointCompression `json:"compression,omitempty"`
	// request headers module configuration or null
	RequestHeaders *EndpointRequestHeaders `json:"request_headers,omitempty"`
	// response headers module configuration or null
	ResponseHeaders *EndpointResponseHeaders `json:"response_headers,omitempty"`
	// webhook verification module configuration or null
	WebhookVerification *EndpointWebhookValidation `json:"webhook_verification,omitempty"`
	// oauth module configuration or null
	OAuth *EndpointOAuth `json:"oauth,omitempty"`
	// saml module configuration or null
	SAML *EndpointSAMLMutate `json:"saml,omitempty"`
	// oidc module configuration or null
	OIDC *EndpointOIDC `json:"oidc,omitempty"`
	// websocket to tcp adapter configuration or null
	WebsocketTCPConverter *EndpointWebsocketTCPConverter `json:"websocket_tcp_converter,omitempty"`
	UserAgentFilter       *EndpointUserAgentFilter       `json:"user_agent_filter,omitempty"`
	// the traffic policy associated with this edge or null
	Policy *EndpointPolicy `json:"policy,omitempty"`
	// the traffic policy associated with this edge or null
	TrafficPolicy *EndpointTrafficPolicy `json:"traffic_policy,omitempty"`
}

type HTTPSEdgeRoute struct {
	// unique identifier of this edge
	EdgeID string `json:"edge_id,omitempty"`
	// unique identifier of this edge route
	ID string `json:"id,omitempty"`
	// timestamp when the edge configuration was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// Type of match to use for this route. Valid values are "exact_path" and
	// "path_prefix".
	MatchType string `json:"match_type,omitempty"`
	// Route selector: "/blog" or "example.com" or "example.com/blog"
	Match string `json:"match,omitempty"`
	// URI of the edge API resource
	URI string `json:"uri,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// backend module configuration or null
	Backend *EndpointBackend `json:"backend,omitempty"`
	// ip restriction module configuration or null
	IpRestriction *EndpointIPPolicy `json:"ip_restriction,omitempty"`
	// circuit breaker module configuration or null
	CircuitBreaker *EndpointCircuitBreaker `json:"circuit_breaker,omitempty"`
	// compression module configuration or null
	Compression *EndpointCompression `json:"compression,omitempty"`
	// request headers module configuration or null
	RequestHeaders *EndpointRequestHeaders `json:"request_headers,omitempty"`
	// response headers module configuration or null
	ResponseHeaders *EndpointResponseHeaders `json:"response_headers,omitempty"`
	// webhook verification module configuration or null
	WebhookVerification *EndpointWebhookValidation `json:"webhook_verification,omitempty"`
	// oauth module configuration or null
	OAuth *EndpointOAuth `json:"oauth,omitempty"`
	// saml module configuration or null
	SAML *EndpointSAML `json:"saml,omitempty"`
	// oidc module configuration or null
	OIDC *EndpointOIDC `json:"oidc,omitempty"`
	// websocket to tcp adapter configuration or null
	WebsocketTCPConverter *EndpointWebsocketTCPConverter `json:"websocket_tcp_converter,omitempty"`
	UserAgentFilter       *EndpointUserAgentFilter       `json:"user_agent_filter,omitempty"`
	// the traffic policy associated with this edge or null
	Policy *EndpointPolicy `json:"policy,omitempty"`
	// the traffic policy associated with this edge or null
	TrafficPolicy *EndpointTrafficPolicy `json:"traffic_policy,omitempty"`
}

type HTTPSEdgeList struct {
	// the list of all HTTPS Edges on this account
	HTTPSEdges []HTTPSEdge `json:"https_edges,omitempty"`
	// URI of the HTTPS Edge list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type HTTPSEdgeCreate struct {
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge; optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports *[]string `json:"hostports,omitempty"`
	// edge modules
	MutualTLS      *EndpointMutualTLSMutate      `json:"mutual_tls,omitempty"`
	TLSTermination *EndpointTLSTerminationAtEdge `json:"tls_termination,omitempty"`
}

type HTTPSEdgeUpdate struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge; optional, max 4096
	// bytes.
	Metadata *string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports *[]string `json:"hostports,omitempty"`
	// edge modules
	MutualTLS      *EndpointMutualTLSMutate      `json:"mutual_tls,omitempty"`
	TLSTermination *EndpointTLSTerminationAtEdge `json:"tls_termination,omitempty"`
}

type HTTPSEdge struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge; optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// timestamp when the edge configuration was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// URI of the edge API resource
	URI string `json:"uri,omitempty"`
	// hostports served by this edge
	Hostports *[]string `json:"hostports,omitempty"`
	// edge modules
	MutualTls      *EndpointMutualTLS      `json:"mutual_tls,omitempty"`
	TlsTermination *EndpointTLSTermination `json:"tls_termination,omitempty"`
	// routes
	Routes []HTTPSEdgeRoute `json:"routes,omitempty"`
}

type EdgeBackendReplace struct {
	ID     string                `json:"id,omitempty"`
	Module EndpointBackendMutate `json:"module,omitempty"`
}

type EdgeIPRestrictionReplace struct {
	ID     string                 `json:"id,omitempty"`
	Module EndpointIPPolicyMutate `json:"module,omitempty"`
}

type EdgeMutualTLSReplace struct {
	ID     string                  `json:"id,omitempty"`
	Module EndpointMutualTLSMutate `json:"module,omitempty"`
}

type EdgeTLSTerminationReplace struct {
	ID     string                 `json:"id,omitempty"`
	Module EndpointTLSTermination `json:"module,omitempty"`
}

type EdgeTLSTerminationAtEdgeReplace struct {
	ID     string                       `json:"id,omitempty"`
	Module EndpointTLSTerminationAtEdge `json:"module,omitempty"`
}

type EdgePolicyReplace struct {
	ID     string         `json:"id,omitempty"`
	Module EndpointPolicy `json:"module,omitempty"`
}

type EdgeTrafficPolicyReplace struct {
	ID     string                `json:"id,omitempty"`
	Module EndpointTrafficPolicy `json:"module,omitempty"`
}

type EdgeRouteBackendReplace struct {
	EdgeID string                `json:"edge_id,omitempty"`
	ID     string                `json:"id,omitempty"`
	Module EndpointBackendMutate `json:"module,omitempty"`
}

type EdgeRouteIPRestrictionReplace struct {
	EdgeID string                 `json:"edge_id,omitempty"`
	ID     string                 `json:"id,omitempty"`
	Module EndpointIPPolicyMutate `json:"module,omitempty"`
}

type EdgeRouteRequestHeadersReplace struct {
	EdgeID string                 `json:"edge_id,omitempty"`
	ID     string                 `json:"id,omitempty"`
	Module EndpointRequestHeaders `json:"module,omitempty"`
}

type EdgeRouteResponseHeadersReplace struct {
	EdgeID string                  `json:"edge_id,omitempty"`
	ID     string                  `json:"id,omitempty"`
	Module EndpointResponseHeaders `json:"module,omitempty"`
}

type EdgeRouteCompressionReplace struct {
	EdgeID string              `json:"edge_id,omitempty"`
	ID     string              `json:"id,omitempty"`
	Module EndpointCompression `json:"module,omitempty"`
}

type EdgeRouteCircuitBreakerReplace struct {
	EdgeID string                 `json:"edge_id,omitempty"`
	ID     string                 `json:"id,omitempty"`
	Module EndpointCircuitBreaker `json:"module,omitempty"`
}

type EdgeRouteWebhookVerificationReplace struct {
	EdgeID string                    `json:"edge_id,omitempty"`
	ID     string                    `json:"id,omitempty"`
	Module EndpointWebhookValidation `json:"module,omitempty"`
}

type EdgeRouteOAuthReplace struct {
	EdgeID string        `json:"edge_id,omitempty"`
	ID     string        `json:"id,omitempty"`
	Module EndpointOAuth `json:"module,omitempty"`
}

type EdgeRouteSAMLReplace struct {
	EdgeID string             `json:"edge_id,omitempty"`
	ID     string             `json:"id,omitempty"`
	Module EndpointSAMLMutate `json:"module,omitempty"`
}

type EdgeRouteOIDCReplace struct {
	EdgeID string       `json:"edge_id,omitempty"`
	ID     string       `json:"id,omitempty"`
	Module EndpointOIDC `json:"module,omitempty"`
}

type EdgeRouteWebsocketTCPConverterReplace struct {
	EdgeID string                        `json:"edge_id,omitempty"`
	ID     string                        `json:"id,omitempty"`
	Module EndpointWebsocketTCPConverter `json:"module,omitempty"`
}

type EdgeRouteUserAgentFilterReplace struct {
	EdgeID string                  `json:"edge_id,omitempty"`
	ID     string                  `json:"id,omitempty"`
	Module EndpointUserAgentFilter `json:"module,omitempty"`
}

type EdgeRoutePolicyReplace struct {
	EdgeID string         `json:"edge_id,omitempty"`
	ID     string         `json:"id,omitempty"`
	Module EndpointPolicy `json:"module,omitempty"`
}

type EdgeRouteTrafficPolicyReplace struct {
	EdgeID string                `json:"edge_id,omitempty"`
	ID     string                `json:"id,omitempty"`
	Module EndpointTrafficPolicy `json:"module,omitempty"`
}

type TCPEdgeList struct {
	// the list of all TCP Edges on this account
	TCPEdges []TCPEdge `json:"tcp_edges,omitempty"`
	// URI of the TCP Edge list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type TCPEdgeCreate struct {
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports *[]string `json:"hostports,omitempty"`
	// edge modules
	Backend       *EndpointBackendMutate  `json:"backend,omitempty"`
	IPRestriction *EndpointIPPolicyMutate `json:"ip_restriction,omitempty"`
	// the traffic policy associated with this edge or null
	Policy *EndpointPolicy `json:"policy,omitempty"`
	// the traffic policy associated with this edge or null
	TrafficPolicy *EndpointTrafficPolicy `json:"traffic_policy,omitempty"`
}

type TCPEdgeUpdate struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata *string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports *[]string `json:"hostports,omitempty"`
	// edge modules
	Backend       *EndpointBackendMutate  `json:"backend,omitempty"`
	IPRestriction *EndpointIPPolicyMutate `json:"ip_restriction,omitempty"`
	// the traffic policy associated with this edge or null
	Policy *EndpointPolicy `json:"policy,omitempty"`
	// the traffic policy associated with this edge or null
	TrafficPolicy *EndpointTrafficPolicy `json:"traffic_policy,omitempty"`
}

type TCPEdge struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// timestamp when the edge was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// URI of the edge API resource
	URI string `json:"uri,omitempty"`
	// hostports served by this edge
	Hostports *[]string `json:"hostports,omitempty"`
	// edge modules
	Backend       *EndpointBackend  `json:"backend,omitempty"`
	IpRestriction *EndpointIPPolicy `json:"ip_restriction,omitempty"`
	// the traffic policy associated with this edge or null
	Policy *EndpointPolicy `json:"policy,omitempty"`
	// the traffic policy associated with this edge or null
	TrafficPolicy *EndpointTrafficPolicy `json:"traffic_policy,omitempty"`
}

type TLSEdgeList struct {
	// the list of all TLS Edges on this account
	TLSEdges []TLSEdge `json:"tls_edges,omitempty"`
	// URI of the TLS Edge list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type TLSEdgeCreate struct {
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports *[]string `json:"hostports,omitempty"`
	// edge modules
	Backend        *EndpointBackendMutate   `json:"backend,omitempty"`
	IPRestriction  *EndpointIPPolicyMutate  `json:"ip_restriction,omitempty"`
	MutualTLS      *EndpointMutualTLSMutate `json:"mutual_tls,omitempty"`
	TLSTermination *EndpointTLSTermination  `json:"tls_termination,omitempty"`
	// the traffic policy associated with this edge or null
	Policy *EndpointPolicy `json:"policy,omitempty"`
	// the traffic policy associated with this edge or null
	TrafficPolicy *EndpointTrafficPolicy `json:"traffic_policy,omitempty"`
}

type TLSEdgeUpdate struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata *string `json:"metadata,omitempty"`
	// hostports served by this edge
	Hostports *[]string `json:"hostports,omitempty"`
	// edge modules
	Backend        *EndpointBackendMutate   `json:"backend,omitempty"`
	IPRestriction  *EndpointIPPolicyMutate  `json:"ip_restriction,omitempty"`
	MutualTLS      *EndpointMutualTLSMutate `json:"mutual_tls,omitempty"`
	TLSTermination *EndpointTLSTermination  `json:"tls_termination,omitempty"`
	// the traffic policy associated with this edge or null
	Policy *EndpointPolicy `json:"policy,omitempty"`
	// the traffic policy associated with this edge or null
	TrafficPolicy *EndpointTrafficPolicy `json:"traffic_policy,omitempty"`
}

type TLSEdge struct {
	// unique identifier of this edge
	ID string `json:"id,omitempty"`
	// human-readable description of what this edge will be used for; optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this edge. Optional, max 4096
	// bytes.
	Metadata string `json:"metadata,omitempty"`
	// timestamp when the edge configuration was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// URI of the edge API resource
	URI string `json:"uri,omitempty"`
	// hostports served by this edge
	Hostports *[]string `json:"hostports,omitempty"`
	// edge modules
	Backend        *EndpointBackend        `json:"backend,omitempty"`
	IpRestriction  *EndpointIPPolicy       `json:"ip_restriction,omitempty"`
	MutualTls      *EndpointMutualTLS      `json:"mutual_tls,omitempty"`
	TlsTermination *EndpointTLSTermination `json:"tls_termination,omitempty"`
	// the traffic policy associated with this edge or null
	Policy *EndpointPolicy `json:"policy,omitempty"`
	// the traffic policy associated with this edge or null
	TrafficPolicy *EndpointTrafficPolicy `json:"traffic_policy,omitempty"`
}

type Endpoint struct {
	// unique endpoint resource identifier
	ID string `json:"id,omitempty"`
	// identifier of the region this endpoint belongs to
	Region string `json:"region,omitempty"`
	// timestamp when the endpoint was created in RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// timestamp when the endpoint was updated in RFC 3339 format
	UpdatedAt string `json:"updated_at,omitempty"`
	// URL of the hostport served by this endpoint
	PublicURL string `json:"public_url,omitempty"`
	// protocol served by this endpoint. one of http, https, tcp, or tls
	Proto  string `json:"proto,omitempty"`
	Scheme string `json:"scheme,omitempty"`
	// hostport served by this endpoint (hostname:port) -> soon to be deprecated
	Hostport string `json:"hostport,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int64  `json:"port,omitempty"`
	// whether the endpoint is ephemeral (served directly by an agent-initiated tunnel)
	// or edge (served by an edge) or cloud (represents a cloud endpoint)
	Type string `json:"type,omitempty"`
	// user-supplied metadata of the associated tunnel or edge object
	Metadata string `json:"metadata,omitempty"`
	// user-supplied description of the associated tunnel
	Description string `json:"description,omitempty"`
	// the domain reserved for this endpoint
	Domain *Ref `json:"domain,omitempty"`
	// the address reserved for this endpoint
	TCPAddr *Ref `json:"tcp_addr,omitempty"`
	// the tunnel serving requests to this endpoint, if this is an ephemeral endpoint
	Tunnel *Ref `json:"tunnel,omitempty"`
	// the edge serving requests to this endpoint, if this is an edge endpoint
	Edge *Ref `json:"edge,omitempty"`
	// the local address the tunnel forwards to
	UpstreamURL string `json:"upstream_url,omitempty"`
	// the protocol the agent uses to forward with
	UpstreamProto string `json:"upstream_proto,omitempty"`
	// the url of the endpoint
	URL string `json:"url,omitempty"`
	// The ID of the owner (bot or user) that owns this endpoint
	Principal *Ref `json:"principal,omitempty"`
	// TODO: deprecate me!
	PrincipalID *Ref `json:"principal_id,omitempty"`
	// The traffic policy attached to this endpoint
	TrafficPolicy string `json:"traffic_policy,omitempty"`
	// the bindings associated with this endpoint
	Bindings *[]string `json:"bindings,omitempty"`
	// The tunnel session of the agent for this endpoint
	TunnelSession *Ref `json:"tunnel_session,omitempty"`
	// URI of the clep API resource
	URI string `json:"uri,omitempty"`
	// user supplied name for the endpoint
	Name string `json:"name,omitempty"`
}

type EndpointList struct {
	// the list of all active endpoints on this account
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	// URI of the endpoints list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type EndpointCreate struct {
	// the url of the endpoint
	URL string `json:"url,omitempty"`
	// whether the endpoint is ephemeral (served directly by an agent-initiated tunnel)
	// or edge (served by an edge) or cloud (represents a cloud endpoint)
	Type string `json:"type,omitempty"`
	// The traffic policy attached to this endpoint
	TrafficPolicy string `json:"traffic_policy,omitempty"`
	// user-supplied description of the associated tunnel
	Description *string `json:"description,omitempty"`
	// user-supplied metadata of the associated tunnel or edge object
	Metadata *string `json:"metadata,omitempty"`
	// the bindings associated with this endpoint
	Bindings *[]string `json:"bindings,omitempty"`
}

type EndpointUpdate struct {
	// unique endpoint resource identifier
	ID string `json:"id,omitempty"`
	// the url of the endpoint
	Url *string `json:"url,omitempty"`
	// The traffic policy attached to this endpoint
	TrafficPolicy *string `json:"traffic_policy,omitempty"`
	// user-supplied description of the associated tunnel
	Description *string `json:"description,omitempty"`
	// user-supplied metadata of the associated tunnel or edge object
	Metadata *string `json:"metadata,omitempty"`
	// the bindings associated with this endpoint
	Bindings *[]string `json:"bindings,omitempty"`
}

type AgentSessionEvent struct {
	// a reference to the session to which this event corresponds
	Session Ref `json:"session,omitempty"`
	// a reference to the credential used to authenticate this session
	Credential *Ref `json:"credential,omitempty"`
	// the ip address from which the agent is connecting
	AgentIP string `json:"agent_ip,omitempty"`
	// the ip address of the ingress server to which the agent is connecting
	IngressServerIP string `json:"ingress_server_ip,omitempty"`
	// the region of the tunnel server
	Region string `json:"region,omitempty"`
	// the hostname of the tunnel server
	IngressHostname string `json:"ingress_hostname,omitempty"`
	// the user agent provided to the tunnel server by the agent
	UserAgent string `json:"user_agent,omitempty"`
	// the session metadata provided by the agent on connection
	Metadata string `json:"metadata,omitempty"`
	// the operating system of the machine on which the agent is running
	OS string `json:"os,omitempty"`
	// the CPU architecture of the machine on which the agent is running
	Arch string `json:"arch,omitempty"`
	// the transport protocol used internally by the agent "muxado" for agents and
	// agent libraries, "ssh" for reverse SSH tunnels
	Transport string `json:"transport,omitempty"`
	// the time at which the session started
	StartedAt string `json:"started_at,omitempty"`
	// the time at which the session expires
	ExpiresAt *string `json:"expires_at,omitempty"`
	// the time at which the session stopped
	StoppedAt *string `json:"stopped_at,omitempty"`
	// If the current agent version is deprecated, informs when support will be dropped
	// and the next minimum supported version
	Deprecated *AgentDeprecated `json:"deprecated,omitempty"`
	// on a failed session start, an explanation of the failure on a successful session
	// start, the empty string on a session stop, the reason for the session stop
	Error string `json:"error,omitempty"`
}

type AgentDeprecated struct {
	// the upcoming minimum supported agent version
	UpcomingMinimumVersion string `json:"upcoming_minimum_version,omitempty"`
	// the date by which the current agent must be upgraded to the upcoming minimum
	// version
	UpcomingEnforcementDate string `json:"upcoming_enforcement_date,omitempty"`
	// additional information about the agent deprecation
	Message string `json:"message,omitempty"`
}

type EventDestinationCreate struct {
	// Arbitrary user-defined machine-readable data of this Event Destination.
	// Optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// Human-readable description of the Event Destination. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// The output format you would like to serialize events into when sending to their
	// target. Currently the only accepted value is JSON.
	Format string `json:"format,omitempty"`
	// An object that encapsulates where and how to send your events. An event
	// destination must contain exactly one of the following objects, leaving the rest
	// null: kinesis, firehose, cloudwatch_logs, or s3.
	Target              EventTarget `json:"target,omitempty"`
	VerifyWithTestEvent *bool       `json:"verify_with_test_event,omitempty"`
}

type EventDestinationUpdate struct {
	// Unique identifier for this Event Destination.
	ID string `json:"id,omitempty"`
	// Arbitrary user-defined machine-readable data of this Event Destination.
	// Optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// Human-readable description of the Event Destination. Optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// The output format you would like to serialize events into when sending to their
	// target. Currently the only accepted value is JSON.
	Format *string `json:"format,omitempty"`
	// An object that encapsulates where and how to send your events. An event
	// destination must contain exactly one of the following objects, leaving the rest
	// null: kinesis, firehose, cloudwatch_logs, or s3.
	Target              *EventTarget `json:"target,omitempty"`
	VerifyWithTestEvent *bool        `json:"verify_with_test_event,omitempty"`
}

type EventDestination struct {
	// Unique identifier for this Event Destination.
	ID string `json:"id,omitempty"`
	// Arbitrary user-defined machine-readable data of this Event Destination.
	// Optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// Timestamp when the Event Destination was created, RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// Human-readable description of the Event Destination. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// The output format you would like to serialize events into when sending to their
	// target. Currently the only accepted value is JSON.
	Format string `json:"format,omitempty"`
	// An object that encapsulates where and how to send your events. An event
	// destination must contain exactly one of the following objects, leaving the rest
	// null: kinesis, firehose, cloudwatch_logs, or s3.
	Target EventTarget `json:"target,omitempty"`
	// URI of the Event Destination API resource.
	URI string `json:"uri,omitempty"`
}

type EventDestinationList struct {
	// The list of all Event Destinations on this account.
	EventDestinations []EventDestination `json:"event_destinations,omitempty"`
	// URI of the Event Destinations list API resource.
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page.
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type EventTarget struct {
	// Configuration used to send events to Amazon Kinesis Data Firehose.
	Firehose *EventTargetFirehose `json:"firehose,omitempty"`
	// Configuration used to send events to Amazon Kinesis.
	Kinesis *EventTargetKinesis `json:"kinesis,omitempty"`
	// Configuration used to send events to Amazon CloudWatch Logs.
	CloudwatchLogs *EventTargetCloudwatchLogs `json:"cloudwatch_logs,omitempty"`
	// Configuration used for internal debugging.
	Debug *EventTargetDebug `json:"debug,omitempty"`
	// Configuration used to send events to Datadog.
	Datadog            *EventTargetDatadog            `json:"datadog,omitempty"`
	AzureLogsIngestion *EventTargetAzureLogsIngestion `json:"azure_logs_ingestion,omitempty"`
}

type EventTargetFirehose struct {
	// Configuration for how to authenticate into your AWS account. Exactly one of role
	// or creds should be configured.
	Auth AWSAuth `json:"auth,omitempty"`
	// An Amazon Resource Name specifying the Firehose delivery stream to deposit
	// events into.
	DeliveryStreamARN string `json:"delivery_stream_arn,omitempty"`
}

type EventTargetKinesis struct {
	// Configuration for how to authenticate into your AWS account. Exactly one of role
	// or creds should be configured.
	Auth AWSAuth `json:"auth,omitempty"`
	// An Amazon Resource Name specifying the Kinesis stream to deposit events into.
	StreamARN string `json:"stream_arn,omitempty"`
}

type EventTargetCloudwatchLogs struct {
	// Configuration for how to authenticate into your AWS account. Exactly one of role
	// or creds should be configured.
	Auth AWSAuth `json:"auth,omitempty"`
	// An Amazon Resource Name specifying the CloudWatch Logs group to deposit events
	// into.
	LogGroupARN string `json:"log_group_arn,omitempty"`
}

type EventTargetS3 struct {
	// Configuration for how to authenticate into your AWS account. Exactly one of role
	// or creds should be configured.
	Auth AWSAuth `json:"auth,omitempty"`
	// An Amazon Resource Name specifying the S3 bucket to deposit events into.
	BucketARN string `json:"bucket_arn,omitempty"`
	// An optional prefix to prepend to S3 object keys.
	ObjectPrefix string `json:"object_prefix,omitempty"`
	// Whether or not to compress files with gzip.
	Compression bool `json:"compression,omitempty"`
	// How many bytes we should accumulate into a single file before sending to S3.
	MaxFileSize int64 `json:"max_file_size,omitempty"`
	// How many seconds we should batch up events before sending them to S3.
	MaxFileAge int64 `json:"max_file_age,omitempty"`
}

type EventTargetDebug struct {
	// Whether or not to output to publisher service logs.
	Log bool `json:"log,omitempty"`
	// URL to send events to.
	CallbackURL string `json:"callback_url,omitempty"`
}

type EventTargetDatadog struct {
	// Datadog API key to use.
	ApiKey *string `json:"api_key,omitempty"`
	// Tags to send with the event.
	Ddtags *string `json:"ddtags,omitempty"`
	// Service name to send with the event.
	Service *string `json:"service,omitempty"`
	// Datadog site to send event to.
	Ddsite *string `json:"ddsite,omitempty"`
}

type EventTargetAzureLogsIngestion struct {
	// Tenant ID for the Azure account
	TenantId string `json:"tenant_id,omitempty"`
	// Client ID for the application client
	ClientId string `json:"client_id,omitempty"`
	// Client Secret for the application client
	ClientSecret string `json:"client_secret,omitempty"`
	// Data collection endpoint logs ingestion URI
	LogsIngestionURI string `json:"logs_ingestion_uri,omitempty"`
	// Data collection rule immutable ID
	DataCollectionRuleId string `json:"data_collection_rule_id,omitempty"`
	// Data collection stream name to use as destination, located inside the DCR
	DataCollectionStreamName string `json:"data_collection_stream_name,omitempty"`
}

type AWSAuth struct {
	// A role for ngrok to assume on your behalf to deposit events into your AWS
	// account.
	Role *AWSRole `json:"role,omitempty"`
	// Credentials to your AWS account if you prefer ngrok to sign in with long-term
	// access keys.
	Creds *AWSCredentials `json:"creds,omitempty"`
}

type AWSRole struct {
	// An ARN that specifies the role that ngrok should use to deliver to the
	// configured target.
	RoleARN string `json:"role_arn,omitempty"`
}

type AWSCredentials struct {
	// The ID portion of an AWS access key.
	AWSAccessKeyID string `json:"aws_access_key_id,omitempty"`
	// The secret portion of an AWS access key.
	AWSSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`
}

type SentEvent struct {
	EventID string `json:"event_id,omitempty"`
}

type EventSubscriptionCreate struct {
	// Arbitrary customer supplied information intended to be machine readable.
	// Optional, max 4096 chars.
	Metadata string `json:"metadata,omitempty"`
	// Arbitrary customer supplied information intended to be human readable. Optional,
	// max 255 chars.
	Description string `json:"description,omitempty"`
	// Sources containing the types for which this event subscription will trigger
	Sources []EventSourceReplace `json:"sources,omitempty"`
	// A list of Event Destination IDs which should be used for this Event
	// Subscription.
	DestinationIDs []string `json:"destination_ids,omitempty"`
}

type EventSubscriptionUpdate struct {
	// Unique identifier for this Event Subscription.
	ID string `json:"id,omitempty"`
	// Arbitrary customer supplied information intended to be machine readable.
	// Optional, max 4096 chars.
	Metadata *string `json:"metadata,omitempty"`
	// Arbitrary customer supplied information intended to be human readable. Optional,
	// max 255 chars.
	Description *string `json:"description,omitempty"`
	// Sources containing the types for which this event subscription will trigger
	Sources *[]EventSourceReplace `json:"sources,omitempty"`
	// A list of Event Destination IDs which should be used for this Event
	// Subscription.
	DestinationIDs *[]string `json:"destination_ids,omitempty"`
}

type EventSubscriptionList struct {
	// The list of all Event Subscriptions on this account.
	EventSubscriptions []EventSubscription `json:"event_subscriptions,omitempty"`
	// URI of the Event Subscriptions list API resource.
	URI string `json:"uri,omitempty"`
	// URI of next page, or null if there is no next page.
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type EventSubscription struct {
	// Unique identifier for this Event Subscription.
	ID string `json:"id,omitempty"`
	// URI of the Event Subscription API resource.
	URI string `json:"uri,omitempty"`
	// When the Event Subscription was created (RFC 3339 format).
	CreatedAt string `json:"created_at,omitempty"`
	// Arbitrary customer supplied information intended to be machine readable.
	// Optional, max 4096 chars.
	Metadata string `json:"metadata,omitempty"`
	// Arbitrary customer supplied information intended to be human readable. Optional,
	// max 255 chars.
	Description string `json:"description,omitempty"`
	// Sources containing the types for which this event subscription will trigger
	Sources []EventSource `json:"sources,omitempty"`
	// Destinations to which these events will be sent
	Destinations []Ref `json:"destinations,omitempty"`
}

type EventSourceReplace struct {
	// Type of event for which an event subscription will trigger
	Type string `json:"type,omitempty"`
	// TODO
	Filter string `json:"filter,omitempty"`
	// TODO
	Fields []string `json:"fields,omitempty"`
}

type EventSource struct {
	// Type of event for which an event subscription will trigger
	Type string `json:"type,omitempty"`
	// TODO
	Filter string `json:"filter,omitempty"`
	// TODO
	Fields []string `json:"fields,omitempty"`
	// URI of the Event Source API resource.
	URI string `json:"uri,omitempty"`
}

type EventSourceList struct {
	// The list of all Event Sources for an Event Subscription
	Sources []EventSource `json:"sources,omitempty"`
	// URI of the next page, or null if there is no next page.
	URI string `json:"uri,omitempty"`
}

type EventSourceCreate struct {
	// The unique identifier for the Event Subscription that this Event Source is
	// attached to.
	SubscriptionID string `json:"subscription_id,omitempty"`
	// Type of event for which an event subscription will trigger
	Type string `json:"type,omitempty"`
	// TODO
	Filter string `json:"filter,omitempty"`
	// TODO
	Fields []string `json:"fields,omitempty"`
}

type EventSourceUpdate struct {
	// The unique identifier for the Event Subscription that this Event Source is
	// attached to.
	SubscriptionID string `json:"subscription_id,omitempty"`
	// Type of event for which an event subscription will trigger
	Type string `json:"type,omitempty"`
	// TODO
	Filter *string `json:"filter,omitempty"`
	// TODO
	Fields *[]string `json:"fields,omitempty"`
}

// This is needed instead of Item because the parameters are different.
type EventSourceItem struct {
	// The unique identifier for the Event Subscription that this Event Source is
	// attached to.
	SubscriptionID string `json:"subscription_id,omitempty"`
	// Type of event for which an event subscription will trigger
	Type string `json:"type,omitempty"`
}

// This is needed instead of Paging because the parameters are different. We also don't need the typical pagination params because pagination of this isn't necessary or supported.
type EventSourcePaging struct {
	// The unique identifier for the Event Subscription that this Event Source is
	// attached to.
	SubscriptionID string `json:"subscription_id,omitempty"`
}

type IPPolicyCreate struct {
	// human-readable description of the source IPs of this IP policy. optional, max
	// 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy. optional, max
	// 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// this field is deprecated. Please leave it empty and use the ip policy rule
	// object's "action" field instead. It is temporarily retained for backwards
	// compatibility reasons.
	Action *string `json:"action,omitempty"`
}

type IPPolicyUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of the source IPs of this IP policy. optional, max
	// 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy. optional, max
	// 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

type IPPolicy struct {
	// unique identifier for this IP policy
	ID string `json:"id,omitempty"`
	// URI of the IP Policy API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the IP policy was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of the source IPs of this IP policy. optional, max
	// 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy. optional, max
	// 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// this field is deprecated. Please leave it empty and use the ip policy rule
	// object's "action" field instead. It is temporarily retained for backwards
	// compatibility reasons.
	Action *string `json:"action,omitempty"`
}

type IPPolicyList struct {
	// the list of all IP policies on this account
	IPPolicies []IPPolicy `json:"ip_policies,omitempty"`
	// URI of the IP policy list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type IPPolicyRuleCreate struct {
	// human-readable description of the source IPs of this IP rule. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy rule. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// an IP or IP range specified in CIDR notation. IPv4 and IPv6 are both supported.
	CIDR string `json:"cidr,omitempty"`
	// ID of the IP policy this IP policy rule will be attached to
	IPPolicyID string `json:"ip_policy_id,omitempty"`
	// the action to apply to the policy rule, either allow or deny
	Action *string `json:"action,omitempty"`
}

type IPPolicyRuleUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of the source IPs of this IP rule. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy rule. optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// an IP or IP range specified in CIDR notation. IPv4 and IPv6 are both supported.
	CIDR *string `json:"cidr,omitempty"`
}

type IPPolicyRule struct {
	// unique identifier for this IP policy rule
	ID string `json:"id,omitempty"`
	// URI of the IP policy rule API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the IP policy rule was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of the source IPs of this IP rule. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP policy rule. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// an IP or IP range specified in CIDR notation. IPv4 and IPv6 are both supported.
	CIDR string `json:"cidr,omitempty"`
	// object describing the IP policy this IP Policy Rule belongs to
	IPPolicy Ref `json:"ip_policy,omitempty"`
	// the action to apply to the policy rule, either allow or deny
	Action string `json:"action,omitempty"`
}

type IPPolicyRuleList struct {
	// the list of all IP policy rules on this account
	IPPolicyRules []IPPolicyRule `json:"ip_policy_rules,omitempty"`
	// URI of the IP policy rule list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type IPRestrictionCreate struct {
	// human-readable description of this IP restriction. optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP restriction. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// true if the IP restriction will be enforced. if false, only warnings will be
	// issued
	Enforced bool `json:"enforced,omitempty"`
	// the type of IP restriction. this defines what traffic will be restricted with
	// the attached policies. four values are currently supported: dashboard, api,
	// agent, and endpoints
	Type string `json:"type,omitempty"`
	// the set of IP policy identifiers that are used to enforce the restriction
	IPPolicyIDs []string `json:"ip_policy_ids,omitempty"`
}

type IPRestrictionUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this IP restriction. optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP restriction. optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// true if the IP restriction will be enforced. if false, only warnings will be
	// issued
	Enforced *bool `json:"enforced,omitempty"`
	// the set of IP policy identifiers that are used to enforce the restriction
	IPPolicyIDs []string `json:"ip_policy_ids,omitempty"`
}

type IPRestriction struct {
	// unique identifier for this IP restriction
	ID string `json:"id,omitempty"`
	// URI of the IP restriction API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the IP restriction was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this IP restriction. optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this IP restriction. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// true if the IP restriction will be enforced. if false, only warnings will be
	// issued
	Enforced bool `json:"enforced,omitempty"`
	// the type of IP restriction. this defines what traffic will be restricted with
	// the attached policies. four values are currently supported: dashboard, api,
	// agent, and endpoints
	Type string `json:"type,omitempty"`
	// the set of IP policies that are used to enforce the restriction
	IPPolicies []Ref `json:"ip_policies,omitempty"`
}

type IPRestrictionList struct {
	// the list of all IP restrictions on this account
	IPRestrictions []IPRestriction `json:"ip_restrictions,omitempty"`
	// URI of the IP restrictions list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type KubernetesOperatorCreate struct {
	// human-readable description of this Kubernetes Operator. optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Kubernetes Operator.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// features enabled for this Kubernetes Operator. a subset of {"bindings",
	// "ingress", and "gateway"}
	EnabledFeatures []string `json:"enabled_features,omitempty"`
	// the ngrok region in which the ingress for this operator is served. defaults to
	// "global"
	Region string `json:"region,omitempty"`
	// information about the deployment of this Kubernetes Operator
	Deployment KubernetesOperatorDeployment `json:"deployment,omitempty"`
	// configuration for the Bindings feature of this Kubernetes Operator. set only if
	// enabling the "bindings" feature
	Binding *KubernetesOperatorBindingCreate `json:"binding,omitempty"`
}

type KubernetesOperatorBindingCreate struct {
	// the name by which endpoints can be bound to this Kubernetes Operator. starts
	// with "k8s/"
	Name string `json:"name,omitempty"`
	// the regexes for urls allowed to be bound to this operator
	AllowedURLs []string `json:"allowed_urls,omitempty"`
	// CSR is supplied during initial creation to enable creating a mutual TLS secured
	// connection between ngrok and the operator. This is an internal implementation
	// detail and subject to change.
	CSR string `json:"csr,omitempty"`
	// the public ingress endpoint for this Kubernetes Operator
	IngressEndpoint *string `json:"ingress_endpoint,omitempty"`
}

type KubernetesOperatorUpdate struct {
	// unique identifier for this Kubernetes Operator
	ID string `json:"id,omitempty"`
	// human-readable description of this Kubernetes Operator. optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Kubernetes Operator.
	// optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// features enabled for this Kubernetes Operator. a subset of {"bindings",
	// "ingress", and "gateway"}
	EnabledFeatures *[]string `json:"enabled_features,omitempty"`
	// the ngrok region in which the ingress for this operator is served. defaults to
	// "global"
	Region *string `json:"region,omitempty"`
	// configuration for the Bindings feature of this Kubernetes Operator. set only if
	// enabling the "bindings" feature
	Binding *KubernetesOperatorBindingUpdate `json:"binding,omitempty"`
}

type KubernetesOperatorBindingUpdate struct {
	// the name by which endpoints can be bound to this Kubernetes Operator. starts
	// with "k8s/"
	Name *string `json:"name,omitempty"`
	// the regexes for urls allowed to be bound to this operator
	AllowedURLs *[]string `json:"allowed_urls,omitempty"`
	// CSR is supplied during initial creation to enable creating a mutual TLS secured
	// connection between ngrok and the operator. This is an internal implementation
	// detail and subject to change.
	CSR *string `json:"csr,omitempty"`
	// the public ingress endpoint for this Kubernetes Operator
	IngressEndpoint *string `json:"ingress_endpoint,omitempty"`
}

type KubernetesOperator struct {
	// unique identifier for this Kubernetes Operator
	ID string `json:"id,omitempty"`
	// URI of this Kubernetes Operator API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the Kubernetes Operator was created. RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// timestamp when the Kubernetes Operator was last updated. RFC 3339 format
	UpdatedAt string `json:"updated_at,omitempty"`
	// human-readable description of this Kubernetes Operator. optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this Kubernetes Operator.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// the principal who created this Kubernetes Operator
	Principal Ref `json:"principal,omitempty"`
	// features enabled for this Kubernetes Operator. a subset of {"bindings",
	// "ingress", and "gateway"}
	EnabledFeatures []string `json:"enabled_features,omitempty"`
	// the ngrok region in which the ingress for this operator is served. defaults to
	// "global"
	Region string `json:"region,omitempty"`
	// information about the deployment of this Kubernetes Operator
	Deployment KubernetesOperatorDeployment `json:"deployment,omitempty"`
	// information about the Bindings feature of this Kubernetes Operator, if enabled
	Binding *KubernetesOperatorBinding `json:"binding,omitempty"`
}

type KubernetesOperatorDeployment struct {
	// the deployment name
	Name string `json:"name,omitempty"`
	// the namespace this Kubernetes Operator is deployed to
	Namespace string `json:"namespace,omitempty"`
	// the version of this Kubernetes Operator
	Version string `json:"version,omitempty"`
}

type KubernetesOperatorCert struct {
	// the public client certificate generated for this Kubernetes Operator from the
	// CSR supplied when enabling the Bindings feature
	Cert string `json:"cert,omitempty"`
	// timestamp when the certificate becomes valid. RFC 3339 format
	NotBefore string `json:"not_before,omitempty"`
	// timestamp when the certificate becomes invalid. RFC 3339 format
	NotAfter string `json:"not_after,omitempty"`
}

type KubernetesOperatorBinding struct {
	// the name by which endpoints can be bound to this Kubernetes Operator. starts
	// with "k8s/"
	Name string `json:"name,omitempty"`
	// the regexes for urls allowed to be bound to this operator
	AllowedURLs []string `json:"allowed_urls,omitempty"`
	// the binding certificate information
	Cert KubernetesOperatorCert `json:"cert,omitempty"`
	// the public ingress endpoint for this Kubernetes Operator
	IngressEndpoint string `json:"ingress_endpoint,omitempty"`
}

type KubernetesOperatorList struct {
	// the list of Kubernetes Operators for this account
	Operators []KubernetesOperator `json:"operators,omitempty"`
	URI       string               `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type EndpointBasicAuthReplace struct {
	ID     string            `json:"id,omitempty"`
	Module EndpointBasicAuth `json:"module,omitempty"`
}

type EndpointCircuitBreakerReplace struct {
	ID     string                 `json:"id,omitempty"`
	Module EndpointCircuitBreaker `json:"module,omitempty"`
}

type EndpointCompressionReplace struct {
	ID     string              `json:"id,omitempty"`
	Module EndpointCompression `json:"module,omitempty"`
}

type EndpointTLSTerminationReplace struct {
	ID     string                 `json:"id,omitempty"`
	Module EndpointTLSTermination `json:"module,omitempty"`
}

type EndpointIPPolicyReplace struct {
	ID     string                 `json:"id,omitempty"`
	Module EndpointIPPolicyMutate `json:"module,omitempty"`
}

type EndpointMutualTLSReplace struct {
	ID     string                  `json:"id,omitempty"`
	Module EndpointMutualTLSMutate `json:"module,omitempty"`
}

type EndpointRequestHeadersReplace struct {
	ID     string                 `json:"id,omitempty"`
	Module EndpointRequestHeaders `json:"module,omitempty"`
}

type EndpointResponseHeadersReplace struct {
	ID     string                  `json:"id,omitempty"`
	Module EndpointResponseHeaders `json:"module,omitempty"`
}

type EndpointOAuthReplace struct {
	ID     string        `json:"id,omitempty"`
	Module EndpointOAuth `json:"module,omitempty"`
}

type EndpointWebhookValidationReplace struct {
	ID     string                    `json:"id,omitempty"`
	Module EndpointWebhookValidation `json:"module,omitempty"`
}

type EndpointSAMLReplace struct {
	ID     string             `json:"id,omitempty"`
	Module EndpointSAMLMutate `json:"module,omitempty"`
}

type EndpointOIDCReplace struct {
	ID     string       `json:"id,omitempty"`
	Module EndpointOIDC `json:"module,omitempty"`
}

type EndpointBackendReplace struct {
	ID     string                `json:"id,omitempty"`
	Module EndpointBackendMutate `json:"module,omitempty"`
}

type ReservedAddrCreate struct {
	// human-readable description of what this reserved address will be used for
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved address. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// reserve the address in this geographic ngrok datacenter. Optional, default is
	// us. (au, eu, ap, us, jp, in, sa)
	Region string `json:"region,omitempty"`
	// ID of an endpoint configuration of type tcp that will be used to handle inbound
	// traffic to this address
	EndpointConfigurationID *string `json:"endpoint_configuration_id,omitempty"`
}

type ReservedAddrUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of what this reserved address will be used for
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved address. Optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// ID of an endpoint configuration of type tcp that will be used to handle inbound
	// traffic to this address
	EndpointConfigurationID *string `json:"endpoint_configuration_id,omitempty"`
}

type ReservedAddr struct {
	// unique reserved address resource identifier
	ID string `json:"id,omitempty"`
	// URI of the reserved address API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the reserved address was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of what this reserved address will be used for
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved address. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// hostname:port of the reserved address that was assigned at creation time
	Addr string `json:"addr,omitempty"`
	// reserve the address in this geographic ngrok datacenter. Optional, default is
	// us. (au, eu, ap, us, jp, in, sa)
	Region string `json:"region,omitempty"`
	// object reference to the endpoint configuration that will be applied to traffic
	// to this address
	EndpointConfiguration *Ref `json:"endpoint_configuration,omitempty"`
}

type ReservedAddrList struct {
	// the list of all reserved addresses on this account
	ReservedAddrs []ReservedAddr `json:"reserved_addrs,omitempty"`
	// URI of the reserved address list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type ReservedDomainCreate struct {
	// the domain name to reserve. It may be a full domain name like app.example.com.
	// If the name does not contain a '.' it will reserve that subdomain on ngrok.io.
	Name string `json:"name,omitempty"`
	// hostname of the reserved domain
	Domain string `json:"domain,omitempty"`
	// deprecated: With the launch of the ngrok Global Network domains traffic is now
	// handled globally. This field applied only to endpoints. Note that agents may
	// still connect to specific regions. Optional, null by default. (au, eu, ap, us,
	// jp, in, sa)
	Region string `json:"region,omitempty"`
	// human-readable description of what this reserved domain will be used for
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved domain. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// ID of an endpoint configuration of type http that will be used to handle inbound
	// http traffic to this domain
	HTTPEndpointConfigurationID *string `json:"http_endpoint_configuration_id,omitempty"`
	// ID of an endpoint configuration of type https that will be used to handle
	// inbound https traffic to this domain
	HTTPSEndpointConfigurationID *string `json:"https_endpoint_configuration_id,omitempty"`
	// ID of a user-uploaded TLS certificate to use for connections to targeting this
	// domain. Optional, mutually exclusive with certificate_management_policy.
	CertificateID *string `json:"certificate_id,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled. Optional, mutually exclusive with
	// certificate_id.
	CertificateManagementPolicy *ReservedDomainCertPolicy `json:"certificate_management_policy,omitempty"`
	// Custom URL with CEL Expression Variable support for redirecting when an ngrok
	// error occurs. Max 10000 bytes.
	ErrorRedirectUrl *string `json:"error_redirect_url,omitempty"`
}

type ReservedDomainUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of what this reserved domain will be used for
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved domain. Optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// ID of an endpoint configuration of type http that will be used to handle inbound
	// http traffic to this domain
	HTTPEndpointConfigurationID *string `json:"http_endpoint_configuration_id,omitempty"`
	// ID of an endpoint configuration of type https that will be used to handle
	// inbound https traffic to this domain
	HTTPSEndpointConfigurationID *string `json:"https_endpoint_configuration_id,omitempty"`
	// ID of a user-uploaded TLS certificate to use for connections to targeting this
	// domain. Optional, mutually exclusive with certificate_management_policy.
	CertificateID *string `json:"certificate_id,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled. Optional, mutually exclusive with
	// certificate_id.
	CertificateManagementPolicy *ReservedDomainCertPolicy `json:"certificate_management_policy,omitempty"`
	// deprecated: With the launch of the ngrok Global Network domains traffic is now
	// handled globally. This field applied only to endpoints. Note that agents may
	// still connect to specific regions. Optional, null by default. (au, eu, ap, us,
	// jp, in, sa)
	Region *string `json:"region,omitempty"`
	// Custom URL with CEL Expression Variable support for redirecting when an ngrok
	// error occurs. Max 10000 bytes.
	ErrorRedirectUrl *string `json:"error_redirect_url,omitempty"`
}

type ReservedDomain struct {
	// unique reserved domain resource identifier
	ID string `json:"id,omitempty"`
	// URI of the reserved domain API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the reserved domain was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of what this reserved domain will be used for
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this reserved domain. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// hostname of the reserved domain
	Domain string `json:"domain,omitempty"`
	// deprecated: With the launch of the ngrok Global Network domains traffic is now
	// handled globally. This field applied only to endpoints. Note that agents may
	// still connect to specific regions. Optional, null by default. (au, eu, ap, us,
	// jp, in, sa)
	Region string `json:"region,omitempty"`
	// DNS CNAME target for a custom hostname, or null if the reserved domain is a
	// subdomain of an ngrok owned domain (e.g. *.ngrok.app)
	CNAMETarget *string `json:"cname_target,omitempty"`
	// object referencing the endpoint configuration applied to http traffic on this
	// domain
	HTTPEndpointConfiguration *Ref `json:"http_endpoint_configuration,omitempty"`
	// object referencing the endpoint configuration applied to https traffic on this
	// domain
	HTTPSEndpointConfiguration *Ref `json:"https_endpoint_configuration,omitempty"`
	// object referencing the TLS certificate used for connections to this domain. This
	// can be either a user-uploaded certificate, the most recently issued automatic
	// one, or null otherwise.
	Certificate *Ref `json:"certificate,omitempty"`
	// configuration for automatic management of TLS certificates for this domain, or
	// null if automatic management is disabled
	CertificateManagementPolicy *ReservedDomainCertPolicy `json:"certificate_management_policy,omitempty"`
	// status of the automatic certificate management for this domain, or null if
	// automatic management is disabled
	CertificateManagementStatus *ReservedDomainCertStatus `json:"certificate_management_status,omitempty"`
	// DNS CNAME target for the host _acme-challenge.example.com, where example.com is
	// your reserved domain name. This is required to issue certificates for wildcard,
	// non-ngrok reserved domains. Must be null for non-wildcard domains and ngrok
	// subdomains.
	ACMEChallengeCNAMETarget *string `json:"acme_challenge_cname_target,omitempty"`
	// Custom URL with CEL Expression Variable support for redirecting when an ngrok
	// error occurs. Max 10000 bytes.
	ErrorRedirectURL *string `json:"error_redirect_url,omitempty"`
}

type ReservedDomainList struct {
	// the list of all reserved domains on this account
	ReservedDomains []ReservedDomain `json:"reserved_domains,omitempty"`
	// URI of the reserved domain list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type ReservedDomainCertPolicy struct {
	// certificate authority to request certificates from. The only supported value is
	// letsencrypt.
	Authority string `json:"authority,omitempty"`
	// type of private key to use when requesting certificates. Defaults to ecdsa, can
	// be either rsa or ecdsa.
	PrivateKeyType string `json:"private_key_type,omitempty"`
}

type ReservedDomainCertStatus struct {
	// timestamp when the next renewal will be requested, RFC 3339 format
	RenewsAt *string `json:"renews_at,omitempty"`
	// status of the certificate provisioning job, or null if the certificiate isn't
	// being provisioned or renewed
	ProvisioningJob *ReservedDomainCertJob `json:"provisioning_job,omitempty"`
}

type ReservedDomainCertJob struct {
	// if present, an error code indicating why provisioning is failing. It may be
	// either a temporary condition (INTERNAL_ERROR), or a permanent one the user must
	// correct (DNS_ERROR).
	ErrorCode *string `json:"error_code,omitempty"`
	// a message describing the current status or error
	Msg string `json:"msg,omitempty"`
	// timestamp when the provisioning job started, RFC 3339 format
	StartedAt string `json:"started_at,omitempty"`
	// timestamp when the provisioning job will be retried
	RetriesAt *string `json:"retries_at,omitempty"`
}

type RootResponse struct {
	URI             string            `json:"uri,omitempty"`
	SubresourceURIs map[string]string `json:"subresource_uris,omitempty"`
}

type SSHCertificateAuthorityCreate struct {
	// human-readable description of this SSH Certificate Authority. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Certificate Authority.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// the type of private key to generate. one of rsa, ecdsa, ed25519
	PrivateKeyType string `json:"private_key_type,omitempty"`
	// the type of elliptic curve to use when creating an ECDSA key
	EllipticCurve string `json:"elliptic_curve,omitempty"`
	// the key size to use when creating an RSA key. one of 2048 or 4096
	KeySize int64 `json:"key_size,omitempty"`
}

type SSHCertificateAuthorityUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this SSH Certificate Authority. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Certificate Authority.
	// optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

type SSHCertificateAuthority struct {
	// unique identifier for this SSH Certificate Authority
	ID string `json:"id,omitempty"`
	// URI of the SSH Certificate Authority API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the SSH Certificate Authority API resource was created, RFC 3339
	// format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this SSH Certificate Authority. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Certificate Authority.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// raw public key for this SSH Certificate Authority
	PublicKey string `json:"public_key,omitempty"`
	// the type of private key for this SSH Certificate Authority
	KeyType string `json:"key_type,omitempty"`
}

type SSHCertificateAuthorityList struct {
	// the list of all certificate authorities on this account
	SSHCertificateAuthorities []SSHCertificateAuthority `json:"ssh_certificate_authorities,omitempty"`
	// URI of the certificates authorities list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type SSHCredentialCreate struct {
	// human-readable description of who or what will use the ssh credential to
	// authenticate. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this ssh credential. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL []string `json:"acl,omitempty"`
	// the PEM-encoded public key of the SSH keypair that will be used to authenticate
	PublicKey string `json:"public_key,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User. Only admins may specify an owner other than themselves. Both owner_id and
	// owner_email may not be specified.
	OwnerEmail string `json:"owner_email,omitempty"`
}

type SSHCredentialUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of who or what will use the ssh credential to
	// authenticate. Optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this ssh credential. Optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL *[]string `json:"acl,omitempty"`
}

type SSHCredential struct {
	// unique ssh credential resource identifier
	ID string `json:"id,omitempty"`
	// URI of the ssh credential API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the ssh credential was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of who or what will use the ssh credential to
	// authenticate. Optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this ssh credential. Optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// the PEM-encoded public key of the SSH keypair that will be used to authenticate
	PublicKey string `json:"public_key,omitempty"`
	// optional list of ACL rules. If unspecified, the credential will have no
	// restrictions. The only allowed ACL rule at this time is the bind rule. The bind
	// rule allows the caller to restrict what domains, addresses, and labels the token
	// is allowed to bind. For example, to allow the token to open a tunnel on
	// example.ngrok.io your ACL would include the rule bind:example.ngrok.io. Bind
	// rules for domains may specify a leading wildcard to match multiple domains with
	// a common suffix. For example, you may specify a rule of bind:*.example.com which
	// will allow x.example.com, y.example.com, *.example.com, etc. Bind rules for
	// labels may specify a wildcard key and/or value to match multiple labels. For
	// example, you may specify a rule of bind:*=example which will allow x=example,
	// y=example, etc. A rule of '*' is equivalent to no acl at all and will explicitly
	// permit all actions.
	ACL []string `json:"acl,omitempty"`
	// If supplied at credential creation, ownership will be assigned to the specified
	// User or Bot. Only admins may specify an owner other than themselves. Defaults to
	// the authenticated User or Bot.
	OwnerID *string `json:"owner_id,omitempty"`
}

type SSHCredentialList struct {
	// the list of all ssh credentials on this account
	SSHCredentials []SSHCredential `json:"ssh_credentials,omitempty"`
	// URI of the ssh credential list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type SSHHostCertificateCreate struct {
	// the ssh certificate authority that is used to sign this ssh host certificate
	SSHCertificateAuthorityID string `json:"ssh_certificate_authority_id,omitempty"`
	// a public key in OpenSSH Authorized Keys format that this certificate signs
	PublicKey string `json:"public_key,omitempty"`
	// the list of principals included in the ssh host certificate. This is the list of
	// hostnames and/or IP addresses that are authorized to serve SSH traffic with this
	// certificate. Dangerously, if no principals are specified, this certificate is
	// considered valid for all hosts.
	Principals []string `json:"principals,omitempty"`
	// The time when the host certificate becomes valid, in RFC 3339 format. Defaults
	// to the current time if unspecified.
	ValidAfter string `json:"valid_after,omitempty"`
	// The time when this host certificate becomes invalid, in RFC 3339 format. If
	// unspecified, a default value of one year in the future will be used. The OpenSSH
	// certificates RFC calls this valid_before.
	ValidUntil string `json:"valid_until,omitempty"`
	// human-readable description of this SSH Host Certificate. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Host Certificate.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
}

type SSHHostCertificateUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this SSH Host Certificate. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Host Certificate.
	// optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

type SSHHostCertificate struct {
	// unique identifier for this SSH Host Certificate
	ID string `json:"id,omitempty"`
	// URI of the SSH Host Certificate API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the SSH Host Certificate API resource was created, RFC 3339
	// format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this SSH Host Certificate. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH Host Certificate.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// a public key in OpenSSH Authorized Keys format that this certificate signs
	PublicKey string `json:"public_key,omitempty"`
	// the key type of the public_key, one of rsa, ecdsa or ed25519
	KeyType string `json:"key_type,omitempty"`
	// the ssh certificate authority that is used to sign this ssh host certificate
	SSHCertificateAuthorityID string `json:"ssh_certificate_authority_id,omitempty"`
	// the list of principals included in the ssh host certificate. This is the list of
	// hostnames and/or IP addresses that are authorized to serve SSH traffic with this
	// certificate. Dangerously, if no principals are specified, this certificate is
	// considered valid for all hosts.
	Principals []string `json:"principals,omitempty"`
	// the time when the ssh host certificate becomes valid, in RFC 3339 format.
	ValidAfter string `json:"valid_after,omitempty"`
	// the time after which the ssh host certificate becomes invalid, in RFC 3339
	// format. the OpenSSH certificates RFC calls this valid_before.
	ValidUntil string `json:"valid_until,omitempty"`
	// the signed SSH certificate in OpenSSH Authorized Keys format. this value should
	// be placed in a -cert.pub certificate file on disk that should be referenced in
	// your sshd_config configuration file with a HostCertificate directive
	Certificate string `json:"certificate,omitempty"`
}

type SSHHostCertificateList struct {
	// the list of all ssh host certificates on this account
	SSHHostCertificates []SSHHostCertificate `json:"ssh_host_certificates,omitempty"`
	// URI of the ssh host certificates list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type SSHUserCertificateCreate struct {
	// the ssh certificate authority that is used to sign this ssh user certificate
	SSHCertificateAuthorityID string `json:"ssh_certificate_authority_id,omitempty"`
	// a public key in OpenSSH Authorized Keys format that this certificate signs
	PublicKey string `json:"public_key,omitempty"`
	// the list of principals included in the ssh user certificate. This is the list of
	// usernames that the certificate holder may sign in as on a machine authorizing
	// the signing certificate authority. Dangerously, if no principals are specified,
	// this certificate may be used to log in as any user.
	Principals []string `json:"principals,omitempty"`
	// A map of critical options included in the certificate. Only two critical options
	// are currently defined by OpenSSH: force-command and source-address. See the
	// OpenSSH certificate protocol spec
	// (https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys) for
	// additional details.
	CriticalOptions map[string]string `json:"critical_options,omitempty"`
	// A map of extensions included in the certificate. Extensions are additional
	// metadata that can be interpreted by the SSH server for any purpose. These can be
	// used to permit or deny the ability to open a terminal, do port forwarding, x11
	// forwarding, and more. If unspecified, the certificate will include limited
	// permissions with the following extension map: {"permit-pty": "",
	// "permit-user-rc": ""} OpenSSH understands a number of predefined extensions. See
	// the OpenSSH certificate protocol spec
	// (https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys) for
	// additional details.
	Extensions map[string]string `json:"extensions,omitempty"`
	// The time when the user certificate becomes valid, in RFC 3339 format. Defaults
	// to the current time if unspecified.
	ValidAfter string `json:"valid_after,omitempty"`
	// The time when this host certificate becomes invalid, in RFC 3339 format. If
	// unspecified, a default value of 24 hours will be used. The OpenSSH certificates
	// RFC calls this valid_before.
	ValidUntil string `json:"valid_until,omitempty"`
	// human-readable description of this SSH User Certificate. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH User Certificate.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
}

type SSHUserCertificateUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this SSH User Certificate. optional, max 255
	// bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH User Certificate.
	// optional, max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

type SSHUserCertificate struct {
	// unique identifier for this SSH User Certificate
	ID string `json:"id,omitempty"`
	// URI of the SSH User Certificate API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the SSH User Certificate API resource was created, RFC 3339
	// format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this SSH User Certificate. optional, max 255
	// bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this SSH User Certificate.
	// optional, max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// a public key in OpenSSH Authorized Keys format that this certificate signs
	PublicKey string `json:"public_key,omitempty"`
	// the key type of the public_key, one of rsa, ecdsa or ed25519
	KeyType string `json:"key_type,omitempty"`
	// the ssh certificate authority that is used to sign this ssh user certificate
	SSHCertificateAuthorityID string `json:"ssh_certificate_authority_id,omitempty"`
	// the list of principals included in the ssh user certificate. This is the list of
	// usernames that the certificate holder may sign in as on a machine authorizing
	// the signing certificate authority. Dangerously, if no principals are specified,
	// this certificate may be used to log in as any user.
	Principals []string `json:"principals,omitempty"`
	// A map of critical options included in the certificate. Only two critical options
	// are currently defined by OpenSSH: force-command and source-address. See the
	// OpenSSH certificate protocol spec
	// (https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys) for
	// additional details.
	CriticalOptions map[string]string `json:"critical_options,omitempty"`
	// A map of extensions included in the certificate. Extensions are additional
	// metadata that can be interpreted by the SSH server for any purpose. These can be
	// used to permit or deny the ability to open a terminal, do port forwarding, x11
	// forwarding, and more. If unspecified, the certificate will include limited
	// permissions with the following extension map: {"permit-pty": "",
	// "permit-user-rc": ""} OpenSSH understands a number of predefined extensions. See
	// the OpenSSH certificate protocol spec
	// (https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys) for
	// additional details.
	Extensions map[string]string `json:"extensions,omitempty"`
	// the time when the ssh host certificate becomes valid, in RFC 3339 format.
	ValidAfter string `json:"valid_after,omitempty"`
	// the time after which the ssh host certificate becomes invalid, in RFC 3339
	// format. the OpenSSH certificates RFC calls this valid_before.
	ValidUntil string `json:"valid_until,omitempty"`
	// the signed SSH certificate in OpenSSH Authorized Keys Format. this value should
	// be placed in a -cert.pub certificate file on disk that should be referenced in
	// your sshd_config configuration file with a HostCertificate directive
	Certificate string `json:"certificate,omitempty"`
}

type SSHUserCertificateList struct {
	// the list of all ssh user certificates on this account
	SSHUserCertificates []SSHUserCertificate `json:"ssh_user_certificates,omitempty"`
	// URI of the ssh user certificates list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type TLSCertificateCreate struct {
	// human-readable description of this TLS certificate. optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this TLS certificate. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// chain of PEM-encoded certificates, leaf first. See Certificate Bundles
	// (https://ngrok.com/docs/cloud-edge/endpoints#certificate-chains).
	CertificatePEM string `json:"certificate_pem,omitempty"`
	// private key for the TLS certificate, PEM-encoded. See Private Keys
	// (https://ngrok.com/docs/cloud-edge/endpoints#private-keys).
	PrivateKeyPEM string `json:"private_key_pem,omitempty"`
}

type TLSCertificateUpdate struct {
	ID string `json:"id,omitempty"`
	// human-readable description of this TLS certificate. optional, max 255 bytes.
	Description *string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this TLS certificate. optional,
	// max 4096 bytes.
	Metadata *string `json:"metadata,omitempty"`
}

type TLSCertificate struct {
	// unique identifier for this TLS certificate
	ID string `json:"id,omitempty"`
	// URI of the TLS certificate API resource
	URI string `json:"uri,omitempty"`
	// timestamp when the TLS certificate was created, RFC 3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// human-readable description of this TLS certificate. optional, max 255 bytes.
	Description string `json:"description,omitempty"`
	// arbitrary user-defined machine-readable data of this TLS certificate. optional,
	// max 4096 bytes.
	Metadata string `json:"metadata,omitempty"`
	// chain of PEM-encoded certificates, leaf first. See Certificate Bundles
	// (https://ngrok.com/docs/cloud-edge/endpoints#certificate-chains).
	CertificatePEM string `json:"certificate_pem,omitempty"`
	// subject common name from the leaf of this TLS certificate
	SubjectCommonName string `json:"subject_common_name,omitempty"`
	// subject alternative names (SANs) from the leaf of this TLS certificate
	SubjectAlternativeNames TLSCertificateSANs `json:"subject_alternative_names,omitempty"`
	// timestamp (in RFC 3339 format) when this TLS certificate was issued
	// automatically, or null if this certificate was user-uploaded
	IssuedAt *string `json:"issued_at,omitempty"`
	// timestamp when this TLS certificate becomes valid, RFC 3339 format
	NotBefore string `json:"not_before,omitempty"`
	// timestamp when this TLS certificate becomes invalid, RFC 3339 format
	NotAfter string `json:"not_after,omitempty"`
	// set of actions the private key of this TLS certificate can be used for
	KeyUsages []string `json:"key_usages,omitempty"`
	// extended set of actions the private key of this TLS certificate can be used for
	ExtendedKeyUsages []string `json:"extended_key_usages,omitempty"`
	// type of the private key of this TLS certificate. One of rsa, ecdsa, or ed25519.
	PrivateKeyType string `json:"private_key_type,omitempty"`
	// issuer common name from the leaf of this TLS certificate
	IssuerCommonName string `json:"issuer_common_name,omitempty"`
	// serial number of the leaf of this TLS certificate
	SerialNumber string `json:"serial_number,omitempty"`
	// subject organization from the leaf of this TLS certificate
	SubjectOrganization string `json:"subject_organization,omitempty"`
	// subject organizational unit from the leaf of this TLS certificate
	SubjectOrganizationalUnit string `json:"subject_organizational_unit,omitempty"`
	// subject locality from the leaf of this TLS certificate
	SubjectLocality string `json:"subject_locality,omitempty"`
	// subject province from the leaf of this TLS certificate
	SubjectProvince string `json:"subject_province,omitempty"`
	// subject country from the leaf of this TLS certificate
	SubjectCountry string `json:"subject_country,omitempty"`
}

type TLSCertificateList struct {
	// the list of all TLS certificates on this account
	TLSCertificates []TLSCertificate `json:"tls_certificates,omitempty"`
	// URI of the TLS certificates list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}

type TLSCertificateSANs struct {
	// set of additional domains (including wildcards) this TLS certificate is valid
	// for
	DNSNames []string `json:"dns_names,omitempty"`
	// set of IP addresses this TLS certificate is also valid for
	IPs []string `json:"ips,omitempty"`
}

type Tunnel struct {
	// unique tunnel resource identifier
	ID string `json:"id,omitempty"`
	// URL of the ephemeral tunnel's public endpoint
	PublicURL string `json:"public_url,omitempty"`
	// timestamp when the tunnel was initiated in RFC 3339 format
	StartedAt string `json:"started_at,omitempty"`
	// user-supplied metadata for the tunnel defined in the ngrok configuration file.
	// See the tunnel metadata configuration option
	// (https://ngrok.com/docs/secure-tunnels/ngrok-agent/reference/config#common-tunnel-configuration-properties)
	// In API version 0, this value was instead pulled from the top-level metadata
	// configuration option
	// (https://ngrok.com/docs/secure-tunnels/ngrok-agent/reference/config#metadata).
	Metadata string `json:"metadata,omitempty"`
	// tunnel protocol for ephemeral tunnels. one of http, https, tcp or tls
	Proto string `json:"proto,omitempty"`
	// identifier of tune region where the tunnel is running
	Region string `json:"region,omitempty"`
	// reference object pointing to the tunnel session on which this tunnel was started
	TunnelSession Ref `json:"tunnel_session,omitempty"`
	// the ephemeral endpoint this tunnel is associated with, if this is an
	// agent-initiated tunnel
	Endpoint *Ref `json:"endpoint,omitempty"`
	// the labels the tunnel group backends will match against, if this is a backend
	// tunnel
	Labels *map[string]string `json:"labels,omitempty"`
	// tunnel group backends served by this backend tunnel
	Backends *[]Ref `json:"backends,omitempty"`
	// upstream address the ngrok agent forwards traffic over this tunnel to. this may
	// be expressed as a URL or a network address.
	ForwardsTo string `json:"forwards_to,omitempty"`
}

type TunnelList struct {
	// the list of all online tunnels on this account
	Tunnels []Tunnel `json:"tunnels,omitempty"`
	// URI of the tunnels list API resource
	URI string `json:"uri,omitempty"`
	// URI of the next page, or null if there is no next page
	NextPageURI *string `json:"next_page_uri,omitempty"`
}
