package shared

type RoleUser string

const (
	RoleUserAdmin RoleUser = "admin"
	RoleUserUser  RoleUser = "user"
	RoleUserGuest RoleUser = "guest"
)
