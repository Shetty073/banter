package enums

type SupportTicketStatus string

const (
	SupportTicketOpen       SupportTicketStatus = "open"
	SupportTicketClosed     SupportTicketStatus = "closed"
	SupportTicketInProgress SupportTicketStatus = "in_progress"
)

type SupportTicketIssueType string

const (
	SupportTicketIssueTechnical SupportTicketIssueType = "technical"
	SupportTicketIssueBilling   SupportTicketIssueType = "billing"
	SupportTicketIssueOther     SupportTicketIssueType = "other"
)
