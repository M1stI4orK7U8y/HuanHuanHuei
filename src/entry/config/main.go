package config

// Port port number
func Port() string {
	return getConfig().Port
}

// DBGrpcURL get db grpc url
func DBGrpcURL() string {
	return getConfig().DB.Grpcurl
}

// CoreGrpcURL get core grpc url
func CoreGrpcURL() string {
	return getConfig().Core.Grpcurl
}
