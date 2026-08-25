package domain

type StatusCode string

const (
	StatusPending  StatusCode = "pending"
	StatusApproved StatusCode = "approved"
	StatusChanged  StatusCode = "changed"
	StatusArchived StatusCode = "archived"
	StatusRejected StatusCode = "rejected"
)

func IsTerminal(s string) bool { return s == string(StatusArchived) }
func CanEdit(s string) bool {
	return s == string(StatusPending) || s == string(StatusApproved) || s == string(StatusChanged)
}
func StatusLabel(s string) string {
	switch s {
	case "pending":
		return "待审核"
	case "approved":
		return "已批准"
	case "changed":
		return "待复核"
	case "archived":
		return "已归档"
	case "rejected":
		return "已驳回"
	}
	return "未知"
}
