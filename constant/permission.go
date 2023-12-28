package constant

const (
	IAMPermissionCreateSubjectRelationTupleDefaultURL                = "https://permission-admin-api.equitize.dev/admin/api/v1/subject-relation-tuples"
	IAMPermissionCreateSubjectRelationTupleMultipleObjectsDefaultURL = "https://permission-admin-api.equitize.dev/admin/api/v1/subject-relation-tuples/multiple-objects"
	IAMPermissionCheckPermissionsDefaultURL                          = "https://permission-public-api.equitize.dev/api/v1/permissions/check"

	IAMPermissionRelationViewer   = "viewer"
	IAMPermissionRelationEditor   = "editor"
	IAMPermissionRelationOwner    = "owner"
	IAMPermissionRelationConsumer = "consumer"

	IAMPermissionCheckPermissionMaxDepthDefault = 3
)
