package registry

type ServiceInstance struct {
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	Version  string            `json:"version"`
	MateData map[string]string `json:"mate_data"`
	Endpoint []string          `json:"endpoint"`
}
