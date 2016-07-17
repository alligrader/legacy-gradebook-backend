package access

import (
	"github.com/mikespook/gorbac"
)

type AssertionFunc func(Role, Permission) bool

func AllGranted(roles []string, perm Permission, assert AssertionFunc) bool {

	rslt := false
	permission := gorbac.Permission(perm)

	for _, role := range roles {
		if !accessControl.IsGranted(role, permission, nil) {
			rslt = true
			break
		}
	}
	return !rslt
}

// TODO Handle the Assert function properly
func AnyGranted(roles []string, perm Permission, assert AssertionFunc) bool {

	rslt := false
	permission := gorbac.Permission(perm)

	for _, role := range roles {
		if accessControl.IsGranted(role, permission, nil) {
			rslt = true
			break
		}
	}
	return rslt
}
