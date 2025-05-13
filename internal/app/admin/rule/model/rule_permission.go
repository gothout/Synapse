package rule

type AdminRulePermission struct {
	RuleID       int64 `json:"rule_id"`
	PermissionID int64 `json:"permission_id"`

	// Optional: For join queries
	Permission AdminPermission `json:"permission,omitempty"`
}

type RuleNamespacePermission struct {
	Namespace string `json:"namespace"` // ex: "admin.enterprise.create"
}
