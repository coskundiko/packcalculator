package messages

const (
	FailedProcessConfigMsg = "Failed to process config"
	EchoServiceStartingMsg = "Echo Service Starting on port: %v"
	ConfigNilErrorMsg      = "Configuration is nil. Ensure config.LoadConfig() is called before server.New()."
	ServerExitedMsg        = "Server Exited Properly"
)
