package system_setting

import (
	"os"

	"github.com/QuantumNous/new-api/setting/config"
)

type OIDCSettings struct {
	Enabled               bool   `json:"enabled"`
	ClientId              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	WellKnown             string `json:"well_known"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserInfoEndpoint      string `json:"user_info_endpoint"`
}

// 默认配置
var defaultOIDCSettings = OIDCSettings{}

func init() {
	// 注册到全局配置管理器
	config.GlobalConfig.Register("oidc", &defaultOIDCSettings)

	// Bootstrap from environment variables (e.g. K8s secrets)
	if v := os.Getenv("OIDC_ENABLED"); v == "true" {
		defaultOIDCSettings.Enabled = true
	}
	if v := os.Getenv("OIDC_CLIENT_ID"); v != "" {
		defaultOIDCSettings.ClientId = v
	}
	if v := os.Getenv("OIDC_CLIENT_SECRET"); v != "" {
		defaultOIDCSettings.ClientSecret = v
	}
	if v := os.Getenv("OIDC_WELL_KNOWN"); v != "" {
		defaultOIDCSettings.WellKnown = v
	}
	if v := os.Getenv("OIDC_AUTHORIZATION_ENDPOINT"); v != "" {
		defaultOIDCSettings.AuthorizationEndpoint = v
	}
	if v := os.Getenv("OIDC_TOKEN_ENDPOINT"); v != "" {
		defaultOIDCSettings.TokenEndpoint = v
	}
	if v := os.Getenv("OIDC_USERINFO_ENDPOINT"); v != "" {
		defaultOIDCSettings.UserInfoEndpoint = v
	}
}

func GetOIDCSettings() *OIDCSettings {
	return &defaultOIDCSettings
}
