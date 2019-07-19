package kongfig

import "github.com/hbagdi/go-kong/kong"

// Config represents a declarative configuration generated using kongfig of Kong < 0.14
type Config struct {
	Apis         []API         `yaml:",omitempty" json:"apis"`
	Consumers    []Consumer    `yaml:",omitempty" json:"consumers"`
	Plugins      []Plugin      `yaml:",omitempty" json:"plugins"`
	Upstreams    []interface{} `yaml:",omitempty" json:"upstreams"`
	Certificates []interface{} `yaml:",omitempty" json:"certificates"`
}

// API represents an API object in Kong < 0.14
type API struct {
	Name       string        `yaml:",omitempty" json:"name"`
	Plugins    []Plugin      `yaml:",omitempty" json:"plugins"`
	Attributes APIAttributes `yaml:",omitempty" json:"attributes"`
}

// APIAttributes represents the attributes of a API
type APIAttributes struct {
	Hosts                  []string `yaml:",omitempty" json:"hosts"`
	Uris                   []string `yaml:",omitempty" json:"uris"`
	StripURI               bool     `yaml:"strip_uri,omitempty" json:"strip_uri"`
	PreserveHost           bool     `yaml:"preserve_host,omitempty" json:"preserve_host"`
	UpstreamURL            string   `yaml:"upstream_url,omitempty" json:"upstream_url"`
	Retries                int      `yaml:",omitempty" json:"retries"`
	UpstreamConnectTimeout int      `yaml:",omitempty" json:"upstream_connect_timeout"`
	UpstreamReadTimeout    int      `yaml:",omitempty" json:"upstream_read_timeout"`
	UpstreamSendTimeout    int      `yaml:",omitempty" json:"upstream_send_timeout"`
	HTTPSOnly              bool     `yaml:"https_only,omitempty" json:"https_only"`
	HTTPIfTerminated       bool     `yaml:"http_if_terminated,omitempty" json:"http_if_terminated"`
	Methods                []string `yaml:",omitempty" json:"methods"`
}

// Plugin represents a Plugin in Kong < 0.14
type Plugin struct {
	Name       string           `yaml:",omitempty" json:"name"`
	Attributes PluginAttributes `yaml:",omitempty" json:"attributes"`
}

// PluginAttributes represents the attributes of a Kong Plugin < 0.14
type PluginAttributes struct {
	Enabled   bool               `yaml:",omitempty" json:"enabled"`
	Config    kong.Configuration `yaml:",omitempty" json:"config"`
	Protocols []string           `yaml:"protocols,omitempty" json:"protocols"`
	RunOn     *string            `yaml:"run_on,omitempty" json:"run_on,omitempty"`
}

// Consumer represents a consumer in Kong < 0.14
type Consumer struct {
	ID       *string `yaml:",omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Username string  `yaml:",omitempty" json:"username"`
	CustomID *string `yaml:",omitempty" json:"custom_id,omitempty"`
	// Acls        []string     `yaml:",omitempty" json:"acls"`
	Credentials []Credential `yaml:",omitempty" json:"credentials"`
}

// Credential represents a consumer credential
type Credential struct {
	Name       string               `yaml:",omitempty" json:"name"`
	Attributes CredentialAttributes `yaml:",omitempty" json:"attributes"`
}

// CredentialAttributes represents the Credential attributes
type CredentialAttributes struct {
	RSAPublicKey *string `yaml:",omitempty" json:"rsa_public_key,omitempty"`
	Algorithm    string  `yaml:",omitempty" json:"algorithm,omitempty"`
	Key          string  `yaml:",omitempty" json:"key"`
	Secret       *string `yaml:",omitempty" json:"secret,omitempty"`
}
