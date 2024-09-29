package flow

type CreatePullConfigCommand struct {
	PullFlowGateway PullFlowGateway
	ConnectionId    any
	Name            string
	IntegrationType IntegrationType
	OriginPath      string
	TargetPath      string
	PrefixFilter    string
	SuffixFilter    string
	RegexFilter     string
}

type UpdatePullConfigCommand struct {
	PullFlowGateway PullFlowGateway
	ConnectionId    any
	Name            string
	IntegrationType IntegrationType
	OriginPath      string
	TargetPath      string
	PrefixFilter    string
	SuffixFilter    string
	RegexFilter     string
}
