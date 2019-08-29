package config

import "time"

// Port port number
func Port() string {
	return getConfig().Port
}

// IP IP
func IP() string {
	return getConfig().IP
}

// Name returns worker name
func Name() string {
	return getConfig().Name
}

// ServiceName returns service name
func ServiceName() string {
	return getConfig().ServiceName
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
