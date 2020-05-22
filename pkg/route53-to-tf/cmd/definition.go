package cmd

// JsonSource for type definition
type JsonSource struct {
	ResourceRecordSets []struct {
		Name            string `json:"Name"`
		Type            string `json:"Type"`
		TTL             int    `json:"TTL,omitempty"`
		ResourceRecords []struct {
			Value string `json:"Value"`
		} `json:"ResourceRecords,omitempty"`
		AliasTarget struct {
			HostedZoneID         string `json:"HostedZoneId"`
			DNSName              string `json:"DNSName"`
			EvaluateTargetHealth bool   `json:"EvaluateTargetHealth"`
		} `json:"AliasTarget,omitempty"`
	} `json:"ResourceRecordSets"`
}