package oktaTypes

// Ref: https://developer.okta.com/docs/reference/api/event-types/
const (
	UserAddedtoApplication     string = "application.user_membership.add"
	UserRemovedFromApplication string = "application.user_membership.remove"
	UserAddedToGroup           string = "group.user_membership.add"
	UserRemovedFromGroup       string = "group.user_membership.remove"
)
