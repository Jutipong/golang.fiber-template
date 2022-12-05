package model

type Config struct {
	Server  *Server_Typd
	Mssql   *Mssql_Type
	Ldap    *Ldap_Type
	Secret  *Secret_Type
	Service *Service_Type
}

type Server_Typd struct {
	Port    string
	Timeout int
}

type Mssql_Type struct {
	Databases Databases_Type `mapstructure:"connection"`
}

type Databases_Type struct {
	GmxDB               string `mapstructure:"gmxDB"`
	TreasuryReportEODDB string `mapstructure:"treasuryReportEODDB"`
}

type Ldap_Type struct {
	Ktb Ktb_Type
}

type Ktb_Type struct {
	Network string
	Address string
	Port    int
	Basedn  struct {
		Search string
	}
	Username string
	Password string
}

type Secret_Type struct {
	Key    string
	JwtKey string
}

type Service_Type struct {
	Endpoint string
	Name     string
}
