package config

import "time"

// IP IP
func IP() string {
	return getConfig().IP
}

// Port port number
func Port() string {
	return getConfig().Port
}

// DBServiceName get db service name
func DBServiceName() string {
	return getConfig().DB.ServiceName
}

// CoreServiceName get core service name
func CoreServiceName() string {
	return getConfig().Core.ServiceName
}

// ETCDHosts returns all etcd hosts address
func ETCDHosts() []string {
	return getConfig().ETCDHosts
}

// ETCDTimeout returns etcd connection timeout
func ETCDTimeout() time.Duration {
	return getConfig().ETCDTimeout * time.Second
}

// Heartbeat returns the heartbeat time to say i am alive
func Heartbeat() time.Duration {
	return getConfig().Heartbeat * time.Second
}
