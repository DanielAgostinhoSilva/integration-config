package flow

type PullFlowGateway interface {
	Save(config PullFlowConfig)
	Delete(config PullFlowConfig)
	FindById(id string) (PullFlowConfig, error)
	ExistOriginPathAndPrefixAndIdNot(originPath string, prefix string, id string) bool
	ExistOriginPathAndSuffixAndIdNot(originPath string, prefix string, id string) bool
	ExistOriginPathAndRegexAndIdNot(originPath string, prefix string, id string) bool
}
