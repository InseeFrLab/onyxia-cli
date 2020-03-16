package api

type CreateNamespaceDTO struct {
	Owner     Owner     `json:"owner"`
	Namespace Namespace `json:"namespace"`
}

type Owner struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Namespace struct {
	Id string `json:"id"`
}
