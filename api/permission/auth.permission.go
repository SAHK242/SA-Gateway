package permission

import grpcauth "gateway/grpc/util"

const (
	ResourceCondition  = "Condition"
	ResourcePermission = "Permission"
	ResourceModules    = "Modules"
	ActionView         = "View"
	ActionUpsert       = "Upsert"
	ActionDelete       = "Delete"
	ActionAccess       = "Access"
)

func ViewConditionPermission() grpcauth.GrpcACL {
	return grpcauth.GrpcACL{
		Resource: ResourceCondition,
		Action:   ActionView,
	}
}

func UpsertConditionPermission() grpcauth.GrpcACL {
	return grpcauth.GrpcACL{
		Resource: ResourceCondition,
		Action:   ActionUpsert,
	}
}

func DeleteConditionPermission() grpcauth.GrpcACL {
	return grpcauth.GrpcACL{
		Resource: ResourceCondition,
		Action:   ActionDelete,
	}
}

func ViewPermissionPermission() grpcauth.GrpcACL {
	return grpcauth.GrpcACL{
		Resource: ResourcePermission,
		Action:   ActionView,
	}
}

func UpsertPermissionPermission() grpcauth.GrpcACL {
	return grpcauth.GrpcACL{
		Resource: ResourcePermission,
		Action:   ActionUpsert,
	}
}

func DeletePermissionPermission() grpcauth.GrpcACL {
	return grpcauth.GrpcACL{
		Resource: ResourcePermission,
		Action:   ActionDelete,
	}
}

func AccessModulesPermission() grpcauth.GrpcACL {
	return grpcauth.GrpcACL{
		Resource: ResourceModules,
		Action:   ActionAccess,
	}
}
