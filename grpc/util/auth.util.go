package grpcutil

import (
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

type (
	GrpcPrincipal struct {
		Type        int32        `json:"type"` // 1: user, 2: machine, 3: service
		User        *GrpcUser    `json:"user"`
		Machine     *GrpcMachine `json:"machine"`
		Service     *GrpcService `json:"service"`
		Permissions []string     `json:"permissions"` // Only for user
		ACLs        []GrpcACL    `json:"acls"`
		Realm       string       `json:"realm"` // check token from which keycloak realm
		jwt.Claims  `json:"-"`
	}

	GrpcUser struct {
		UserId    string `json:"user_id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		TenantId  string `json:"tenant_id"`
		OrgId     string `json:"org_id"`
	}

	GrpcMachine struct {
		MachineId string `json:"machine_id"`
	}

	GrpcService struct {
		ServiceId string `json:"service_id"`
	}

	Audience struct {
		Type int32  `json:"type"`
		Code string `json:"code"` // unique value to determine the audiences (ex: if 'type' is Role then 'code' is role code)
	}

	GrpcACL struct {
		Action      string     `json:"action"`   // action name
		Resource    string     `json:"resource"` // resource name
		Decision    int32      `json:"decision"`
		Scope       int32      `json:"scope"`
		CustomScope string     `json:"custom_scope"` // custom scope code
		Audiences   []Audience `json:"audiences"`
		Module      int32      `json:"module"`
	}
)

func (g GrpcACL) WithScope(scope int32) GrpcACL {
	g.Scope = scope
	return g
}

func (g GrpcACL) WithCustomScope(customScope string) GrpcACL {
	g.Scope = 4
	g.CustomScope = customScope
	return g
}

func (p GrpcPrincipal) HasAcl(acl GrpcACL) bool {
	return HasAcl(p.ACLs, acl)
}

func (p GrpcPrincipal) HasAllAcl(checks ...GrpcACL) bool {
	return HasAllAcl(p.ACLs, checks...)
}

func (p GrpcPrincipal) HasAnyAcl(checks ...GrpcACL) bool {
	return HasAnyAcl(p.ACLs, checks...)
}

func HasAllAcl(acls []GrpcACL, checks ...GrpcACL) bool {
	for _, acl := range checks {
		if !HasAcl(acls, acl) {
			return false
		}
	}
	return true
}

func HasAnyAcl(acls []GrpcACL, checks ...GrpcACL) bool {
	for _, acl := range checks {
		if HasAcl(acls, acl) {
			return true
		}
	}
	return false
}

func HasAcl(acls []GrpcACL, acl GrpcACL) bool {
	for _, item := range acls {
		if item.Decision != 1 {
			continue
		}
		if !strings.EqualFold(item.Resource, acl.Resource) ||
			!strings.EqualFold(item.Action, acl.Action) {
			continue
		}
		// if required acl.Scope == 0 , don't check scope
		if acl.Scope != 0 && item.Scope != acl.Scope {
			continue
		}

		if acl.CustomScope != "" && !strings.EqualFold(item.CustomScope, acl.CustomScope) {
			continue
		}

		return true
	}
	return false
}
