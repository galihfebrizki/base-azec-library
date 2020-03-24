package baseModel

type EurekaConfig struct {
	EurekaStatus string `json:"eurekaStatus"`
	PathFile string `json:"pathFile"`
	UrlReg string `json:"urlReg"`
	UrlDeregister string `json:"urlDeregister"`
}