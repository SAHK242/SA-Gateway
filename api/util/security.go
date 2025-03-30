package apiutil

import (
	grpcauth "gateway/grpc/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

const (
	ContextPermissionKey      = "ContextPermission"
	ContextPrincipalKey       = "ContextPrincipal"
	ContextSubjectEvaluateKey = "ContextSubjectEvaluateKey"
)

func SetPermissions(ctx *fiber.Ctx, permissions []string) {
	if permissions == nil {
		return
	}

	permissionsString, err := json.Marshal(permissions)

	if err != nil {
		return
	}

	ctx.Locals(ContextPermissionKey, string(permissionsString))
}

//func SetEvaluateSubject(ctx *fiber.Ctx, subject []*gcommon.SubjectEvaluation) {
//	if subject == nil {
//		return
//	}
//
//	subjectString, err := json.Marshal(subject)
//
//	if err != nil {
//		return
//	}
//
//	ctx.Locals(ContextSubjectEvaluateKey, string(subjectString))
//}
//
//func toString(value interface{}) string {
//	if str, ok := value.(string); ok {
//		return str
//	}
//	return fmt.Sprintf("%v", value)
//}
//
//func GetEvaluateSubject(ctx *fiber.Ctx) []grpcauth.GrpcACL {
//	subjectString := toString(ctx.Locals(ContextSubjectEvaluateKey))
//
//	if subjectString == "" {
//		return []grpcauth.GrpcACL{}
//	}
//
//	subjects := make([]*gcommon.SubjectEvaluation, 0)
//
//	err := json.Unmarshal([]byte(subjectString), &subjects)
//
//	if err != nil {
//		return []grpcauth.GrpcACL{}
//	}
//
//	grpcAcls := make([]grpcauth.GrpcACL, 0)
//	for _, subject := range subjects {
//		grpcAcls = append(grpcAcls, grpcauth.GrpcACL{
//			Action:   subject.Action.Name,
//			Resource: subject.Resource.Name,
//			Decision: int32(subject.Decision),
//			Scope:    int32(subject.Scope),
//			Audiences: lo.Map(subject.Audiences, func(item *gcommon.Audience, _ int) grpcauth.Audience {
//				return grpcauth.Audience{
//					Type: int32(item.Type),
//					Code: item.Code,
//				}
//			}),
//			Module: int32(subject.Module),
//		})
//	}
//
//	return grpcAcls
//}

//func GetPrincipal(ctx *fiber.Ctx) *grpcauth.GrpcPrincipal {
//	principalString := toString(ctx.Locals(ContextPrincipalKey))
//
//	if principalString == "" {
//		p, e := BuildPrincipal(ctx, gcommon.PrincipalType_PRINCIPAL_TYPE_USER)
//
//		if e != nil {
//			return nil
//		}
//
//		go SetPrincipal(ctx, p)
//
//		return p
//	}
//
//	principal := new(grpcauth.GrpcPrincipal)
//
//	err := json.Unmarshal([]byte(principalString), principal)
//
//	if err != nil {
//		return nil
//	}
//
//	return principal
//}

func SetPrincipal(ctx *fiber.Ctx, principal *grpcauth.GrpcPrincipal) {
	if principal == nil {
		return
	}

	principalString, err := json.Marshal(principal)

	if err != nil {
		return
	}

	ctx.Locals(ContextPrincipalKey, string(principalString))
}

//func BuildPrincipal(ctx *fiber.Ctx, principalType gcommon.PrincipalType) (*grpcauth.GrpcPrincipal, error) {
//	switch principalType {
//	case gcommon.PrincipalType_PRINCIPAL_TYPE_USER:
//		return buildUserPrincipal(ctx)
//	case gcommon.PrincipalType_PRINCIPAL_TYPE_MACHINE:
//		return buildMachinePrincipal(ctx)
//	case gcommon.PrincipalType_PRINCIPAL_TYPE_SERVICE:
//		return buildServicePrincipal(ctx)
//	default:
//		return nil, fmt.Errorf("invalid principal type %s", principalType.String())
//	}
//}
