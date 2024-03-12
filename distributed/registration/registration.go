package registration

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}
type ServiceName string

const (
	LogService = ServiceName("LogService")
)
