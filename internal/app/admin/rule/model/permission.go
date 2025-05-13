package rule

type AdminPermission struct {
	ID       int64  `json:"id"`
	ModuleID int64  `json:"module_id"`
	Action   string `json:"action"`

	// Optional: For joined queries
	Module AdminModule `json:"module,omitempty"`
}
