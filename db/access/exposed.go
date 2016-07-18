package access

import (
	"github.com/mikespook/gorbac"
)

type AssertionFunc func(Role, Permission) bool

func AllGranted(roles []Role, perm Permission, assert AssertionFunc) bool {

	rslt := false
	permission := gorbac.Permission(perm)

	for _, role := range roles {
		var roleName string = roleStr(role)
		if !accessControl.IsGranted(roleName, permission, nil) {
			rslt = true
			break
		}
	}
	return !rslt
}

// TODO Handle the Assert function properly
func AnyGranted(roles []Role, perm Permission, assert AssertionFunc) bool {

	rslt := false
	permission := gorbac.Permission(perm)

	for _, role := range roles {
		var roleName string = roleStr(role)
		if accessControl.IsGranted(roleName, permission, nil) {
			rslt = true
			break
		}
	}
	return rslt
}

func roleStr(r Role) string {
	var str string
	switch r {
	case Admin:
		str = admin
	case Assistant:
		str = assistant
	case Billing:
		str = billing
	case Student:
		str = student
	case Teacher:
		str = teacher
	default:
		str = "unknown role... did you make this role yourself?"
	}
	return str
}
