package _5_easyjson

// cmd >> go/bin/easyjson -all struct_def.go
type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}
