package entity

type Deployment struct {
	Code     string
	Replicas int
	Image    string
	Port     int
}
