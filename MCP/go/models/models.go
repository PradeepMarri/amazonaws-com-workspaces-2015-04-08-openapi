package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeleteConnectClientAddInResult represents the DeleteConnectClientAddInResult schema from the OpenAPI specification
type DeleteConnectClientAddInResult struct {
}

// StartWorkspacesRequest represents the StartWorkspacesRequest schema from the OpenAPI specification
type StartWorkspacesRequest struct {
	Startworkspacerequests interface{} `json:"StartWorkspaceRequests"`
}

// ConnectionAliasAssociation represents the ConnectionAliasAssociation schema from the OpenAPI specification
type ConnectionAliasAssociation struct {
	Connectionidentifier interface{} `json:"ConnectionIdentifier,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Associatedaccountid interface{} `json:"AssociatedAccountId,omitempty"`
	Associationstatus interface{} `json:"AssociationStatus,omitempty"`
}

// RebootWorkspacesResult represents the RebootWorkspacesResult schema from the OpenAPI specification
type RebootWorkspacesResult struct {
	Failedrequests interface{} `json:"FailedRequests,omitempty"`
}

// ModifyAccountResult represents the ModifyAccountResult schema from the OpenAPI specification
type ModifyAccountResult struct {
}

// CreateWorkspacesRequest represents the CreateWorkspacesRequest schema from the OpenAPI specification
type CreateWorkspacesRequest struct {
	Workspaces interface{} `json:"Workspaces"`
}

// DeregisterWorkspaceDirectoryResult represents the DeregisterWorkspaceDirectoryResult schema from the OpenAPI specification
type DeregisterWorkspaceDirectoryResult struct {
}

// ModifyCertificateBasedAuthPropertiesRequest represents the ModifyCertificateBasedAuthPropertiesRequest schema from the OpenAPI specification
type ModifyCertificateBasedAuthPropertiesRequest struct {
	Certificatebasedauthproperties interface{} `json:"CertificateBasedAuthProperties,omitempty"`
	Propertiestodelete interface{} `json:"PropertiesToDelete,omitempty"`
	Resourceid interface{} `json:"ResourceId"`
}

// DescribeTagsResult represents the DescribeTagsResult schema from the OpenAPI specification
type DescribeTagsResult struct {
	Taglist interface{} `json:"TagList,omitempty"`
}

// RevokeIpRulesRequest represents the RevokeIpRulesRequest schema from the OpenAPI specification
type RevokeIpRulesRequest struct {
	Groupid interface{} `json:"GroupId"`
	Userrules interface{} `json:"UserRules"`
}

// DeleteConnectionAliasRequest represents the DeleteConnectionAliasRequest schema from the OpenAPI specification
type DeleteConnectionAliasRequest struct {
	Aliasid interface{} `json:"AliasId"`
}

// UpdateConnectClientAddInResult represents the UpdateConnectClientAddInResult schema from the OpenAPI specification
type UpdateConnectClientAddInResult struct {
}

// ModifyWorkspaceAccessPropertiesResult represents the ModifyWorkspaceAccessPropertiesResult schema from the OpenAPI specification
type ModifyWorkspaceAccessPropertiesResult struct {
}

// ModifyWorkspaceStateRequest represents the ModifyWorkspaceStateRequest schema from the OpenAPI specification
type ModifyWorkspaceStateRequest struct {
	Workspaceid interface{} `json:"WorkspaceId"`
	Workspacestate interface{} `json:"WorkspaceState"`
}

// RebootWorkspacesRequest represents the RebootWorkspacesRequest schema from the OpenAPI specification
type RebootWorkspacesRequest struct {
	Rebootworkspacerequests interface{} `json:"RebootWorkspaceRequests"`
}

// FailedCreateWorkspaceRequest represents the FailedCreateWorkspaceRequest schema from the OpenAPI specification
type FailedCreateWorkspaceRequest struct {
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Workspacerequest interface{} `json:"WorkspaceRequest,omitempty"`
}

// DisassociateConnectionAliasRequest represents the DisassociateConnectionAliasRequest schema from the OpenAPI specification
type DisassociateConnectionAliasRequest struct {
	Aliasid interface{} `json:"AliasId"`
}

// DeleteTagsRequest represents the DeleteTagsRequest schema from the OpenAPI specification
type DeleteTagsRequest struct {
	Resourceid interface{} `json:"ResourceId"`
	Tagkeys interface{} `json:"TagKeys"`
}

// AssociateConnectionAliasResult represents the AssociateConnectionAliasResult schema from the OpenAPI specification
type AssociateConnectionAliasResult struct {
	Connectionidentifier interface{} `json:"ConnectionIdentifier,omitempty"`
}

// DisassociateConnectionAliasResult represents the DisassociateConnectionAliasResult schema from the OpenAPI specification
type DisassociateConnectionAliasResult struct {
}

// DeleteClientBrandingRequest represents the DeleteClientBrandingRequest schema from the OpenAPI specification
type DeleteClientBrandingRequest struct {
	Platforms interface{} `json:"Platforms"`
	Resourceid interface{} `json:"ResourceId"`
}

// RebuildWorkspacesRequest represents the RebuildWorkspacesRequest schema from the OpenAPI specification
type RebuildWorkspacesRequest struct {
	Rebuildworkspacerequests interface{} `json:"RebuildWorkspaceRequests"`
}

// UpdateWorkspaceImagePermissionRequest represents the UpdateWorkspaceImagePermissionRequest schema from the OpenAPI specification
type UpdateWorkspaceImagePermissionRequest struct {
	Allowcopyimage interface{} `json:"AllowCopyImage"`
	Imageid interface{} `json:"ImageId"`
	Sharedaccountid interface{} `json:"SharedAccountId"`
}

// DescribeClientPropertiesResult represents the DescribeClientPropertiesResult schema from the OpenAPI specification
type DescribeClientPropertiesResult struct {
	Clientpropertieslist interface{} `json:"ClientPropertiesList,omitempty"`
}

// CreateUpdatedWorkspaceImageRequest represents the CreateUpdatedWorkspaceImageRequest schema from the OpenAPI specification
type CreateUpdatedWorkspaceImageRequest struct {
	Description interface{} `json:"Description"`
	Name interface{} `json:"Name"`
	Sourceimageid interface{} `json:"SourceImageId"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateIpGroupRequest represents the CreateIpGroupRequest schema from the OpenAPI specification
type CreateIpGroupRequest struct {
	Groupdesc interface{} `json:"GroupDesc,omitempty"`
	Groupname interface{} `json:"GroupName"`
	Tags interface{} `json:"Tags,omitempty"`
	Userrules interface{} `json:"UserRules,omitempty"`
}

// DeleteWorkspaceImageRequest represents the DeleteWorkspaceImageRequest schema from the OpenAPI specification
type DeleteWorkspaceImageRequest struct {
	Imageid interface{} `json:"ImageId"`
}

// IosClientBrandingAttributes represents the IosClientBrandingAttributes schema from the OpenAPI specification
type IosClientBrandingAttributes struct {
	Loginmessage interface{} `json:"LoginMessage,omitempty"`
	Logo2xurl interface{} `json:"Logo2xUrl,omitempty"`
	Logo3xurl interface{} `json:"Logo3xUrl,omitempty"`
	Logourl interface{} `json:"LogoUrl,omitempty"`
	Supportemail interface{} `json:"SupportEmail,omitempty"`
	Supportlink interface{} `json:"SupportLink,omitempty"`
	Forgotpasswordlink interface{} `json:"ForgotPasswordLink,omitempty"`
}

// StopRequest represents the StopRequest schema from the OpenAPI specification
type StopRequest struct {
	Workspaceid interface{} `json:"WorkspaceId,omitempty"`
}

// DeleteConnectClientAddInRequest represents the DeleteConnectClientAddInRequest schema from the OpenAPI specification
type DeleteConnectClientAddInRequest struct {
	Addinid interface{} `json:"AddInId"`
	Resourceid interface{} `json:"ResourceId"`
}

// CopyWorkspaceImageResult represents the CopyWorkspaceImageResult schema from the OpenAPI specification
type CopyWorkspaceImageResult struct {
	Imageid interface{} `json:"ImageId,omitempty"`
}

// DescribeConnectionAliasPermissionsRequest represents the DescribeConnectionAliasPermissionsRequest schema from the OpenAPI specification
type DescribeConnectionAliasPermissionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Aliasid interface{} `json:"AliasId"`
}

// UserStorage represents the UserStorage schema from the OpenAPI specification
type UserStorage struct {
	Capacity interface{} `json:"Capacity,omitempty"`
}

// CreateIpGroupResult represents the CreateIpGroupResult schema from the OpenAPI specification
type CreateIpGroupResult struct {
	Groupid interface{} `json:"GroupId,omitempty"`
}

// ConnectionAlias represents the ConnectionAlias schema from the OpenAPI specification
type ConnectionAlias struct {
	Aliasid interface{} `json:"AliasId,omitempty"`
	Associations interface{} `json:"Associations,omitempty"`
	Connectionstring interface{} `json:"ConnectionString,omitempty"`
	Owneraccountid interface{} `json:"OwnerAccountId,omitempty"`
	State interface{} `json:"State,omitempty"`
}

// ImagePermission represents the ImagePermission schema from the OpenAPI specification
type ImagePermission struct {
	Sharedaccountid interface{} `json:"SharedAccountId,omitempty"`
}

// ClientProperties represents the ClientProperties schema from the OpenAPI specification
type ClientProperties struct {
	Loguploadenabled interface{} `json:"LogUploadEnabled,omitempty"`
	Reconnectenabled interface{} `json:"ReconnectEnabled,omitempty"`
}

// StandbyWorkspace represents the StandbyWorkspace schema from the OpenAPI specification
type StandbyWorkspace struct {
	Tags interface{} `json:"Tags,omitempty"`
	Volumeencryptionkey interface{} `json:"VolumeEncryptionKey,omitempty"`
	Directoryid interface{} `json:"DirectoryId"`
	Primaryworkspaceid interface{} `json:"PrimaryWorkspaceId"`
}

// WorkspaceCreationProperties represents the WorkspaceCreationProperties schema from the OpenAPI specification
type WorkspaceCreationProperties struct {
	Customsecuritygroupid interface{} `json:"CustomSecurityGroupId,omitempty"`
	Defaultou interface{} `json:"DefaultOu,omitempty"`
	Enableinternetaccess interface{} `json:"EnableInternetAccess,omitempty"`
	Enablemaintenancemode interface{} `json:"EnableMaintenanceMode,omitempty"`
	Enableworkdocs interface{} `json:"EnableWorkDocs,omitempty"`
	Userenabledaslocaladministrator interface{} `json:"UserEnabledAsLocalAdministrator,omitempty"`
}

// RegisterWorkspaceDirectoryResult represents the RegisterWorkspaceDirectoryResult schema from the OpenAPI specification
type RegisterWorkspaceDirectoryResult struct {
}

// ImportWorkspaceImageResult represents the ImportWorkspaceImageResult schema from the OpenAPI specification
type ImportWorkspaceImageResult struct {
	Imageid interface{} `json:"ImageId,omitempty"`
}

// CreateTagsResult represents the CreateTagsResult schema from the OpenAPI specification
type CreateTagsResult struct {
}

// UpdateWorkspaceBundleResult represents the UpdateWorkspaceBundleResult schema from the OpenAPI specification
type UpdateWorkspaceBundleResult struct {
}

// CreateWorkspacesResult represents the CreateWorkspacesResult schema from the OpenAPI specification
type CreateWorkspacesResult struct {
	Failedrequests interface{} `json:"FailedRequests,omitempty"`
	Pendingrequests interface{} `json:"PendingRequests,omitempty"`
}

// DescribeWorkspaceDirectoriesRequest represents the DescribeWorkspaceDirectoriesRequest schema from the OpenAPI specification
type DescribeWorkspaceDirectoriesRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Directoryids interface{} `json:"DirectoryIds,omitempty"`
}

// UpdateWorkspaceBundleRequest represents the UpdateWorkspaceBundleRequest schema from the OpenAPI specification
type UpdateWorkspaceBundleRequest struct {
	Bundleid interface{} `json:"BundleId,omitempty"`
	Imageid interface{} `json:"ImageId,omitempty"`
}

// CreateConnectionAliasRequest represents the CreateConnectionAliasRequest schema from the OpenAPI specification
type CreateConnectionAliasRequest struct {
	Connectionstring interface{} `json:"ConnectionString"`
	Tags interface{} `json:"Tags,omitempty"`
}

// OperatingSystem represents the OperatingSystem schema from the OpenAPI specification
type OperatingSystem struct {
	TypeField interface{} `json:"Type,omitempty"`
}

// DescribeAccountModificationsResult represents the DescribeAccountModificationsResult schema from the OpenAPI specification
type DescribeAccountModificationsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Accountmodifications interface{} `json:"AccountModifications,omitempty"`
}

// DescribeWorkspacesConnectionStatusRequest represents the DescribeWorkspacesConnectionStatusRequest schema from the OpenAPI specification
type DescribeWorkspacesConnectionStatusRequest struct {
	Workspaceids interface{} `json:"WorkspaceIds,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// WorkspaceBundle represents the WorkspaceBundle schema from the OpenAPI specification
type WorkspaceBundle struct {
	Rootstorage interface{} `json:"RootStorage,omitempty"`
	State interface{} `json:"State,omitempty"`
	Userstorage interface{} `json:"UserStorage,omitempty"`
	Computetype interface{} `json:"ComputeType,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
	Bundleid interface{} `json:"BundleId,omitempty"`
	Bundletype interface{} `json:"BundleType,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Imageid interface{} `json:"ImageId,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
}

// DescribeWorkspacesConnectionStatusResult represents the DescribeWorkspacesConnectionStatusResult schema from the OpenAPI specification
type DescribeWorkspacesConnectionStatusResult struct {
	Workspacesconnectionstatus interface{} `json:"WorkspacesConnectionStatus,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateWorkspaceImageRequest represents the CreateWorkspaceImageRequest schema from the OpenAPI specification
type CreateWorkspaceImageRequest struct {
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Workspaceid interface{} `json:"WorkspaceId"`
	Description interface{} `json:"Description"`
}

// CreateConnectClientAddInRequest represents the CreateConnectClientAddInRequest schema from the OpenAPI specification
type CreateConnectClientAddInRequest struct {
	Name interface{} `json:"Name"`
	Resourceid interface{} `json:"ResourceId"`
	Url interface{} `json:"URL"`
}

// RebootRequest represents the RebootRequest schema from the OpenAPI specification
type RebootRequest struct {
	Workspaceid interface{} `json:"WorkspaceId"`
}

// DescribeAccountRequest represents the DescribeAccountRequest schema from the OpenAPI specification
type DescribeAccountRequest struct {
}

// DescribeWorkspaceImagesRequest represents the DescribeWorkspaceImagesRequest schema from the OpenAPI specification
type DescribeWorkspaceImagesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Imageids interface{} `json:"ImageIds,omitempty"`
	Imagetype interface{} `json:"ImageType,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// ModifySelfservicePermissionsResult represents the ModifySelfservicePermissionsResult schema from the OpenAPI specification
type ModifySelfservicePermissionsResult struct {
}

// StartRequest represents the StartRequest schema from the OpenAPI specification
type StartRequest struct {
	Workspaceid interface{} `json:"WorkspaceId,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Value interface{} `json:"Value,omitempty"`
	Key interface{} `json:"Key"`
}

// DefaultImportClientBrandingAttributes represents the DefaultImportClientBrandingAttributes schema from the OpenAPI specification
type DefaultImportClientBrandingAttributes struct {
	Forgotpasswordlink interface{} `json:"ForgotPasswordLink,omitempty"`
	Loginmessage interface{} `json:"LoginMessage,omitempty"`
	Logo interface{} `json:"Logo,omitempty"`
	Supportemail interface{} `json:"SupportEmail,omitempty"`
	Supportlink interface{} `json:"SupportLink,omitempty"`
}

// ModifyWorkspacePropertiesResult represents the ModifyWorkspacePropertiesResult schema from the OpenAPI specification
type ModifyWorkspacePropertiesResult struct {
}

// DefaultClientBrandingAttributes represents the DefaultClientBrandingAttributes schema from the OpenAPI specification
type DefaultClientBrandingAttributes struct {
	Forgotpasswordlink interface{} `json:"ForgotPasswordLink,omitempty"`
	Loginmessage interface{} `json:"LoginMessage,omitempty"`
	Logourl interface{} `json:"LogoUrl,omitempty"`
	Supportemail interface{} `json:"SupportEmail,omitempty"`
	Supportlink interface{} `json:"SupportLink,omitempty"`
}

// DisassociateIpGroupsResult represents the DisassociateIpGroupsResult schema from the OpenAPI specification
type DisassociateIpGroupsResult struct {
}

// DescribeConnectClientAddInsResult represents the DescribeConnectClientAddInsResult schema from the OpenAPI specification
type DescribeConnectClientAddInsResult struct {
	Addins interface{} `json:"AddIns,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ModifyWorkspaceAccessPropertiesRequest represents the ModifyWorkspaceAccessPropertiesRequest schema from the OpenAPI specification
type ModifyWorkspaceAccessPropertiesRequest struct {
	Workspaceaccessproperties interface{} `json:"WorkspaceAccessProperties"`
	Resourceid interface{} `json:"ResourceId"`
}

// ComputeType represents the ComputeType schema from the OpenAPI specification
type ComputeType struct {
	Name interface{} `json:"Name,omitempty"`
}

// ModificationState represents the ModificationState schema from the OpenAPI specification
type ModificationState struct {
	State interface{} `json:"State,omitempty"`
	Resource interface{} `json:"Resource,omitempty"`
}

// DeleteClientBrandingResult represents the DeleteClientBrandingResult schema from the OpenAPI specification
type DeleteClientBrandingResult struct {
}

// DescribeWorkspaceImagePermissionsResult represents the DescribeWorkspaceImagePermissionsResult schema from the OpenAPI specification
type DescribeWorkspaceImagePermissionsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Imageid interface{} `json:"ImageId,omitempty"`
	Imagepermissions interface{} `json:"ImagePermissions,omitempty"`
}

// UpdateResult represents the UpdateResult schema from the OpenAPI specification
type UpdateResult struct {
	Description interface{} `json:"Description,omitempty"`
	Updateavailable interface{} `json:"UpdateAvailable,omitempty"`
}

// RebuildRequest represents the RebuildRequest schema from the OpenAPI specification
type RebuildRequest struct {
	Workspaceid interface{} `json:"WorkspaceId"`
}

// ConnectClientAddIn represents the ConnectClientAddIn schema from the OpenAPI specification
type ConnectClientAddIn struct {
	Addinid interface{} `json:"AddInId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Url interface{} `json:"URL,omitempty"`
}

// DescribeIpGroupsResult represents the DescribeIpGroupsResult schema from the OpenAPI specification
type DescribeIpGroupsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Result interface{} `json:"Result,omitempty"`
}

// CreateStandbyWorkspacesRequest represents the CreateStandbyWorkspacesRequest schema from the OpenAPI specification
type CreateStandbyWorkspacesRequest struct {
	Primaryregion interface{} `json:"PrimaryRegion"`
	Standbyworkspaces interface{} `json:"StandbyWorkspaces"`
}

// CertificateBasedAuthProperties represents the CertificateBasedAuthProperties schema from the OpenAPI specification
type CertificateBasedAuthProperties struct {
	Status interface{} `json:"Status,omitempty"`
	Certificateauthorityarn interface{} `json:"CertificateAuthorityArn,omitempty"`
}

// DeleteIpGroupResult represents the DeleteIpGroupResult schema from the OpenAPI specification
type DeleteIpGroupResult struct {
}

// LoginMessage represents the LoginMessage schema from the OpenAPI specification
type LoginMessage struct {
}

// ModifyWorkspaceStateResult represents the ModifyWorkspaceStateResult schema from the OpenAPI specification
type ModifyWorkspaceStateResult struct {
}

// IpRuleItem represents the IpRuleItem schema from the OpenAPI specification
type IpRuleItem struct {
	Ruledesc interface{} `json:"ruleDesc,omitempty"`
	Iprule interface{} `json:"ipRule,omitempty"`
}

// SamlProperties represents the SamlProperties schema from the OpenAPI specification
type SamlProperties struct {
	Status interface{} `json:"Status,omitempty"`
	Useraccessurl interface{} `json:"UserAccessUrl,omitempty"`
	Relaystateparametername interface{} `json:"RelayStateParameterName,omitempty"`
}

// DeregisterWorkspaceDirectoryRequest represents the DeregisterWorkspaceDirectoryRequest schema from the OpenAPI specification
type DeregisterWorkspaceDirectoryRequest struct {
	Directoryid interface{} `json:"DirectoryId"`
}

// ModifySamlPropertiesRequest represents the ModifySamlPropertiesRequest schema from the OpenAPI specification
type ModifySamlPropertiesRequest struct {
	Samlproperties interface{} `json:"SamlProperties,omitempty"`
	Propertiestodelete interface{} `json:"PropertiesToDelete,omitempty"`
	Resourceid interface{} `json:"ResourceId"`
}

// ModifySamlPropertiesResult represents the ModifySamlPropertiesResult schema from the OpenAPI specification
type ModifySamlPropertiesResult struct {
}

// DeleteConnectionAliasResult represents the DeleteConnectionAliasResult schema from the OpenAPI specification
type DeleteConnectionAliasResult struct {
}

// ModifySelfservicePermissionsRequest represents the ModifySelfservicePermissionsRequest schema from the OpenAPI specification
type ModifySelfservicePermissionsRequest struct {
	Selfservicepermissions interface{} `json:"SelfservicePermissions"`
	Resourceid interface{} `json:"ResourceId"`
}

// DescribeWorkspaceSnapshotsResult represents the DescribeWorkspaceSnapshotsResult schema from the OpenAPI specification
type DescribeWorkspaceSnapshotsResult struct {
	Rebuildsnapshots interface{} `json:"RebuildSnapshots,omitempty"`
	Restoresnapshots interface{} `json:"RestoreSnapshots,omitempty"`
}

// UpdateConnectClientAddInRequest represents the UpdateConnectClientAddInRequest schema from the OpenAPI specification
type UpdateConnectClientAddInRequest struct {
	Addinid interface{} `json:"AddInId"`
	Name interface{} `json:"Name,omitempty"`
	Resourceid interface{} `json:"ResourceId"`
	Url interface{} `json:"URL,omitempty"`
}

// AssociateIpGroupsRequest represents the AssociateIpGroupsRequest schema from the OpenAPI specification
type AssociateIpGroupsRequest struct {
	Groupids interface{} `json:"GroupIds"`
	Directoryid interface{} `json:"DirectoryId"`
}

// DescribeConnectionAliasPermissionsResult represents the DescribeConnectionAliasPermissionsResult schema from the OpenAPI specification
type DescribeConnectionAliasPermissionsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Aliasid interface{} `json:"AliasId,omitempty"`
	Connectionaliaspermissions interface{} `json:"ConnectionAliasPermissions,omitempty"`
}

// AuthorizeIpRulesResult represents the AuthorizeIpRulesResult schema from the OpenAPI specification
type AuthorizeIpRulesResult struct {
}

// ImportWorkspaceImageRequest represents the ImportWorkspaceImageRequest schema from the OpenAPI specification
type ImportWorkspaceImageRequest struct {
	Imagename interface{} `json:"ImageName"`
	Ingestionprocess interface{} `json:"IngestionProcess"`
	Tags interface{} `json:"Tags,omitempty"`
	Applications interface{} `json:"Applications,omitempty"`
	Ec2imageid interface{} `json:"Ec2ImageId"`
	Imagedescription interface{} `json:"ImageDescription"`
}

// ListAvailableManagementCidrRangesRequest represents the ListAvailableManagementCidrRangesRequest schema from the OpenAPI specification
type ListAvailableManagementCidrRangesRequest struct {
	Managementcidrrangeconstraint interface{} `json:"ManagementCidrRangeConstraint"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeWorkspacesRequest represents the DescribeWorkspacesRequest schema from the OpenAPI specification
type DescribeWorkspacesRequest struct {
	Username interface{} `json:"UserName,omitempty"`
	Workspaceids interface{} `json:"WorkspaceIds,omitempty"`
	Bundleid interface{} `json:"BundleId,omitempty"`
	Directoryid interface{} `json:"DirectoryId,omitempty"`
	Limit interface{} `json:"Limit,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ClientPropertiesResult represents the ClientPropertiesResult schema from the OpenAPI specification
type ClientPropertiesResult struct {
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Clientproperties interface{} `json:"ClientProperties,omitempty"`
}

// DescribeTagsRequest represents the DescribeTagsRequest schema from the OpenAPI specification
type DescribeTagsRequest struct {
	Resourceid interface{} `json:"ResourceId"`
}

// CopyWorkspaceImageRequest represents the CopyWorkspaceImageRequest schema from the OpenAPI specification
type CopyWorkspaceImageRequest struct {
	Name interface{} `json:"Name"`
	Sourceimageid interface{} `json:"SourceImageId"`
	Sourceregion interface{} `json:"SourceRegion"`
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// TerminateWorkspacesRequest represents the TerminateWorkspacesRequest schema from the OpenAPI specification
type TerminateWorkspacesRequest struct {
	Terminateworkspacerequests interface{} `json:"TerminateWorkspaceRequests"`
}

// ImportClientBrandingResult represents the ImportClientBrandingResult schema from the OpenAPI specification
type ImportClientBrandingResult struct {
	Devicetypeios interface{} `json:"DeviceTypeIos,omitempty"`
	Devicetypelinux interface{} `json:"DeviceTypeLinux,omitempty"`
	Devicetypeosx interface{} `json:"DeviceTypeOsx,omitempty"`
	Devicetypeweb interface{} `json:"DeviceTypeWeb,omitempty"`
	Devicetypewindows interface{} `json:"DeviceTypeWindows,omitempty"`
	Devicetypeandroid interface{} `json:"DeviceTypeAndroid,omitempty"`
}

// DescribeConnectClientAddInsRequest represents the DescribeConnectClientAddInsRequest schema from the OpenAPI specification
type DescribeConnectClientAddInsRequest struct {
	Resourceid interface{} `json:"ResourceId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ImportClientBrandingRequest represents the ImportClientBrandingRequest schema from the OpenAPI specification
type ImportClientBrandingRequest struct {
	Devicetypeweb interface{} `json:"DeviceTypeWeb,omitempty"`
	Devicetypewindows interface{} `json:"DeviceTypeWindows,omitempty"`
	Resourceid interface{} `json:"ResourceId"`
	Devicetypeandroid interface{} `json:"DeviceTypeAndroid,omitempty"`
	Devicetypeios interface{} `json:"DeviceTypeIos,omitempty"`
	Devicetypelinux interface{} `json:"DeviceTypeLinux,omitempty"`
	Devicetypeosx interface{} `json:"DeviceTypeOsx,omitempty"`
}

// DeleteIpGroupRequest represents the DeleteIpGroupRequest schema from the OpenAPI specification
type DeleteIpGroupRequest struct {
	Groupid interface{} `json:"GroupId"`
}

// DeleteTagsResult represents the DeleteTagsResult schema from the OpenAPI specification
type DeleteTagsResult struct {
}

// UpdateConnectionAliasPermissionResult represents the UpdateConnectionAliasPermissionResult schema from the OpenAPI specification
type UpdateConnectionAliasPermissionResult struct {
}

// DescribeClientBrandingResult represents the DescribeClientBrandingResult schema from the OpenAPI specification
type DescribeClientBrandingResult struct {
	Devicetypeandroid interface{} `json:"DeviceTypeAndroid,omitempty"`
	Devicetypeios interface{} `json:"DeviceTypeIos,omitempty"`
	Devicetypelinux interface{} `json:"DeviceTypeLinux,omitempty"`
	Devicetypeosx interface{} `json:"DeviceTypeOsx,omitempty"`
	Devicetypeweb interface{} `json:"DeviceTypeWeb,omitempty"`
	Devicetypewindows interface{} `json:"DeviceTypeWindows,omitempty"`
}

// DescribeWorkspaceImagesResult represents the DescribeWorkspaceImagesResult schema from the OpenAPI specification
type DescribeWorkspaceImagesResult struct {
	Images interface{} `json:"Images,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RelatedWorkspaceProperties represents the RelatedWorkspaceProperties schema from the OpenAPI specification
type RelatedWorkspaceProperties struct {
	Region interface{} `json:"Region,omitempty"`
	State interface{} `json:"State,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Workspaceid interface{} `json:"WorkspaceId,omitempty"`
}

// RestoreWorkspaceResult represents the RestoreWorkspaceResult schema from the OpenAPI specification
type RestoreWorkspaceResult struct {
}

// ModifyWorkspacePropertiesRequest represents the ModifyWorkspacePropertiesRequest schema from the OpenAPI specification
type ModifyWorkspacePropertiesRequest struct {
	Workspaceid interface{} `json:"WorkspaceId"`
	Workspaceproperties interface{} `json:"WorkspaceProperties"`
}

// DescribeClientPropertiesRequest represents the DescribeClientPropertiesRequest schema from the OpenAPI specification
type DescribeClientPropertiesRequest struct {
	Resourceids interface{} `json:"ResourceIds"`
}

// CreateTagsRequest represents the CreateTagsRequest schema from the OpenAPI specification
type CreateTagsRequest struct {
	Resourceid interface{} `json:"ResourceId"`
	Tags interface{} `json:"Tags"`
}

// CreateStandbyWorkspacesResult represents the CreateStandbyWorkspacesResult schema from the OpenAPI specification
type CreateStandbyWorkspacesResult struct {
	Failedstandbyrequests interface{} `json:"FailedStandbyRequests,omitempty"`
	Pendingstandbyrequests interface{} `json:"PendingStandbyRequests,omitempty"`
}

// TerminateWorkspacesResult represents the TerminateWorkspacesResult schema from the OpenAPI specification
type TerminateWorkspacesResult struct {
	Failedrequests interface{} `json:"FailedRequests,omitempty"`
}

// AccountModification represents the AccountModification schema from the OpenAPI specification
type AccountModification struct {
	Dedicatedtenancymanagementcidrrange interface{} `json:"DedicatedTenancyManagementCidrRange,omitempty"`
	Dedicatedtenancysupport interface{} `json:"DedicatedTenancySupport,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Modificationstate interface{} `json:"ModificationState,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
}

// DeleteWorkspaceBundleResult represents the DeleteWorkspaceBundleResult schema from the OpenAPI specification
type DeleteWorkspaceBundleResult struct {
}

// DescribeWorkspaceImagePermissionsRequest represents the DescribeWorkspaceImagePermissionsRequest schema from the OpenAPI specification
type DescribeWorkspaceImagePermissionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Imageid interface{} `json:"ImageId"`
}

// DescribeAccountResult represents the DescribeAccountResult schema from the OpenAPI specification
type DescribeAccountResult struct {
	Dedicatedtenancymanagementcidrrange interface{} `json:"DedicatedTenancyManagementCidrRange,omitempty"`
	Dedicatedtenancysupport interface{} `json:"DedicatedTenancySupport,omitempty"`
}

// MigrateWorkspaceResult represents the MigrateWorkspaceResult schema from the OpenAPI specification
type MigrateWorkspaceResult struct {
	Sourceworkspaceid interface{} `json:"SourceWorkspaceId,omitempty"`
	Targetworkspaceid interface{} `json:"TargetWorkspaceId,omitempty"`
}

// DescribeWorkspaceBundlesRequest represents the DescribeWorkspaceBundlesRequest schema from the OpenAPI specification
type DescribeWorkspaceBundlesRequest struct {
	Owner interface{} `json:"Owner,omitempty"`
	Bundleids interface{} `json:"BundleIds,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// WorkspaceAccessProperties represents the WorkspaceAccessProperties schema from the OpenAPI specification
type WorkspaceAccessProperties struct {
	Devicetypeios interface{} `json:"DeviceTypeIos,omitempty"`
	Devicetypelinux interface{} `json:"DeviceTypeLinux,omitempty"`
	Devicetypeosx interface{} `json:"DeviceTypeOsx,omitempty"`
	Devicetypeweb interface{} `json:"DeviceTypeWeb,omitempty"`
	Devicetypewindows interface{} `json:"DeviceTypeWindows,omitempty"`
	Devicetypezeroclient interface{} `json:"DeviceTypeZeroClient,omitempty"`
	Devicetypeandroid interface{} `json:"DeviceTypeAndroid,omitempty"`
	Devicetypechromeos interface{} `json:"DeviceTypeChromeOs,omitempty"`
}

// UpdateWorkspaceImagePermissionResult represents the UpdateWorkspaceImagePermissionResult schema from the OpenAPI specification
type UpdateWorkspaceImagePermissionResult struct {
}

// AuthorizeIpRulesRequest represents the AuthorizeIpRulesRequest schema from the OpenAPI specification
type AuthorizeIpRulesRequest struct {
	Userrules interface{} `json:"UserRules"`
	Groupid interface{} `json:"GroupId"`
}

// RestoreWorkspaceRequest represents the RestoreWorkspaceRequest schema from the OpenAPI specification
type RestoreWorkspaceRequest struct {
	Workspaceid interface{} `json:"WorkspaceId"`
}

// PendingCreateStandbyWorkspacesRequest represents the PendingCreateStandbyWorkspacesRequest schema from the OpenAPI specification
type PendingCreateStandbyWorkspacesRequest struct {
	Directoryid interface{} `json:"DirectoryId,omitempty"`
	State interface{} `json:"State,omitempty"`
	Username interface{} `json:"UserName,omitempty"`
	Workspaceid interface{} `json:"WorkspaceId,omitempty"`
}

// RootStorage represents the RootStorage schema from the OpenAPI specification
type RootStorage struct {
	Capacity interface{} `json:"Capacity,omitempty"`
}

// StopWorkspacesRequest represents the StopWorkspacesRequest schema from the OpenAPI specification
type StopWorkspacesRequest struct {
	Stopworkspacerequests interface{} `json:"StopWorkspaceRequests"`
}

// ModifyAccountRequest represents the ModifyAccountRequest schema from the OpenAPI specification
type ModifyAccountRequest struct {
	Dedicatedtenancymanagementcidrrange interface{} `json:"DedicatedTenancyManagementCidrRange,omitempty"`
	Dedicatedtenancysupport interface{} `json:"DedicatedTenancySupport,omitempty"`
}

// DeleteWorkspaceBundleRequest represents the DeleteWorkspaceBundleRequest schema from the OpenAPI specification
type DeleteWorkspaceBundleRequest struct {
	Bundleid interface{} `json:"BundleId,omitempty"`
}

// DescribeWorkspacesResult represents the DescribeWorkspacesResult schema from the OpenAPI specification
type DescribeWorkspacesResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Workspaces interface{} `json:"Workspaces,omitempty"`
}

// IosImportClientBrandingAttributes represents the IosImportClientBrandingAttributes schema from the OpenAPI specification
type IosImportClientBrandingAttributes struct {
	Forgotpasswordlink interface{} `json:"ForgotPasswordLink,omitempty"`
	Loginmessage interface{} `json:"LoginMessage,omitempty"`
	Logo interface{} `json:"Logo,omitempty"`
	Logo2x interface{} `json:"Logo2x,omitempty"`
	Logo3x interface{} `json:"Logo3x,omitempty"`
	Supportemail interface{} `json:"SupportEmail,omitempty"`
	Supportlink interface{} `json:"SupportLink,omitempty"`
}

// FailedWorkspaceChangeRequest represents the FailedWorkspaceChangeRequest schema from the OpenAPI specification
type FailedWorkspaceChangeRequest struct {
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Workspaceid interface{} `json:"WorkspaceId,omitempty"`
}

// ModifyClientPropertiesRequest represents the ModifyClientPropertiesRequest schema from the OpenAPI specification
type ModifyClientPropertiesRequest struct {
	Clientproperties interface{} `json:"ClientProperties"`
	Resourceid interface{} `json:"ResourceId"`
}

// ConnectionAliasPermission represents the ConnectionAliasPermission schema from the OpenAPI specification
type ConnectionAliasPermission struct {
	Allowassociation interface{} `json:"AllowAssociation"`
	Sharedaccountid interface{} `json:"SharedAccountId"`
}

// CreateUpdatedWorkspaceImageResult represents the CreateUpdatedWorkspaceImageResult schema from the OpenAPI specification
type CreateUpdatedWorkspaceImageResult struct {
	Imageid interface{} `json:"ImageId,omitempty"`
}

// CreateConnectClientAddInResult represents the CreateConnectClientAddInResult schema from the OpenAPI specification
type CreateConnectClientAddInResult struct {
	Addinid interface{} `json:"AddInId,omitempty"`
}

// SelfservicePermissions represents the SelfservicePermissions schema from the OpenAPI specification
type SelfservicePermissions struct {
	Changecomputetype interface{} `json:"ChangeComputeType,omitempty"`
	Increasevolumesize interface{} `json:"IncreaseVolumeSize,omitempty"`
	Rebuildworkspace interface{} `json:"RebuildWorkspace,omitempty"`
	Restartworkspace interface{} `json:"RestartWorkspace,omitempty"`
	Switchrunningmode interface{} `json:"SwitchRunningMode,omitempty"`
}

// ModifyClientPropertiesResult represents the ModifyClientPropertiesResult schema from the OpenAPI specification
type ModifyClientPropertiesResult struct {
}

// WorkspaceProperties represents the WorkspaceProperties schema from the OpenAPI specification
type WorkspaceProperties struct {
	Rootvolumesizegib interface{} `json:"RootVolumeSizeGib,omitempty"`
	Runningmode interface{} `json:"RunningMode,omitempty"`
	Runningmodeautostoptimeoutinminutes interface{} `json:"RunningModeAutoStopTimeoutInMinutes,omitempty"`
	Uservolumesizegib interface{} `json:"UserVolumeSizeGib,omitempty"`
	Computetypename interface{} `json:"ComputeTypeName,omitempty"`
	Protocols interface{} `json:"Protocols,omitempty"`
}

// ModifyWorkspaceCreationPropertiesRequest represents the ModifyWorkspaceCreationPropertiesRequest schema from the OpenAPI specification
type ModifyWorkspaceCreationPropertiesRequest struct {
	Resourceid interface{} `json:"ResourceId"`
	Workspacecreationproperties interface{} `json:"WorkspaceCreationProperties"`
}

// ModifyCertificateBasedAuthPropertiesResult represents the ModifyCertificateBasedAuthPropertiesResult schema from the OpenAPI specification
type ModifyCertificateBasedAuthPropertiesResult struct {
}

// DefaultWorkspaceCreationProperties represents the DefaultWorkspaceCreationProperties schema from the OpenAPI specification
type DefaultWorkspaceCreationProperties struct {
	Enableinternetaccess interface{} `json:"EnableInternetAccess,omitempty"`
	Enablemaintenancemode interface{} `json:"EnableMaintenanceMode,omitempty"`
	Enableworkdocs interface{} `json:"EnableWorkDocs,omitempty"`
	Userenabledaslocaladministrator interface{} `json:"UserEnabledAsLocalAdministrator,omitempty"`
	Customsecuritygroupid interface{} `json:"CustomSecurityGroupId,omitempty"`
	Defaultou interface{} `json:"DefaultOu,omitempty"`
}

// DisassociateIpGroupsRequest represents the DisassociateIpGroupsRequest schema from the OpenAPI specification
type DisassociateIpGroupsRequest struct {
	Directoryid interface{} `json:"DirectoryId"`
	Groupids interface{} `json:"GroupIds"`
}

// RevokeIpRulesResult represents the RevokeIpRulesResult schema from the OpenAPI specification
type RevokeIpRulesResult struct {
}

// DescribeConnectionAliasesResult represents the DescribeConnectionAliasesResult schema from the OpenAPI specification
type DescribeConnectionAliasesResult struct {
	Connectionaliases interface{} `json:"ConnectionAliases,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeAccountModificationsRequest represents the DescribeAccountModificationsRequest schema from the OpenAPI specification
type DescribeAccountModificationsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MigrateWorkspaceRequest represents the MigrateWorkspaceRequest schema from the OpenAPI specification
type MigrateWorkspaceRequest struct {
	Bundleid interface{} `json:"BundleId"`
	Sourceworkspaceid interface{} `json:"SourceWorkspaceId"`
}

// AssociateIpGroupsResult represents the AssociateIpGroupsResult schema from the OpenAPI specification
type AssociateIpGroupsResult struct {
}

// DescribeClientBrandingRequest represents the DescribeClientBrandingRequest schema from the OpenAPI specification
type DescribeClientBrandingRequest struct {
	Resourceid interface{} `json:"ResourceId"`
}

// Workspace represents the Workspace schema from the OpenAPI specification
type Workspace struct {
	State interface{} `json:"State,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Relatedworkspaces interface{} `json:"RelatedWorkspaces,omitempty"`
	Workspaceid interface{} `json:"WorkspaceId,omitempty"`
	Volumeencryptionkey interface{} `json:"VolumeEncryptionKey,omitempty"`
	Bundleid interface{} `json:"BundleId,omitempty"`
	Computername interface{} `json:"ComputerName,omitempty"`
	Directoryid interface{} `json:"DirectoryId,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Modificationstates interface{} `json:"ModificationStates,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Uservolumeencryptionenabled interface{} `json:"UserVolumeEncryptionEnabled,omitempty"`
	Ipaddress interface{} `json:"IpAddress,omitempty"`
	Username interface{} `json:"UserName,omitempty"`
	Workspaceproperties interface{} `json:"WorkspaceProperties,omitempty"`
	Rootvolumeencryptionenabled interface{} `json:"RootVolumeEncryptionEnabled,omitempty"`
}

// DescribeIpGroupsRequest represents the DescribeIpGroupsRequest schema from the OpenAPI specification
type DescribeIpGroupsRequest struct {
	Groupids interface{} `json:"GroupIds,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// WorkspaceDirectory represents the WorkspaceDirectory schema from the OpenAPI specification
type WorkspaceDirectory struct {
	Workspacecreationproperties interface{} `json:"WorkspaceCreationProperties,omitempty"`
	Alias interface{} `json:"Alias,omitempty"`
	Ipgroupids interface{} `json:"ipGroupIds,omitempty"`
	Directoryid interface{} `json:"DirectoryId,omitempty"`
	Registrationcode interface{} `json:"RegistrationCode,omitempty"`
	Workspaceaccessproperties interface{} `json:"WorkspaceAccessProperties,omitempty"`
	Selfservicepermissions interface{} `json:"SelfservicePermissions,omitempty"`
	Certificatebasedauthproperties interface{} `json:"CertificateBasedAuthProperties,omitempty"`
	Iamroleid interface{} `json:"IamRoleId,omitempty"`
	Tenancy interface{} `json:"Tenancy,omitempty"`
	Directoryname interface{} `json:"DirectoryName,omitempty"`
	Samlproperties interface{} `json:"SamlProperties,omitempty"`
	Directorytype interface{} `json:"DirectoryType,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Workspacesecuritygroupid interface{} `json:"WorkspaceSecurityGroupId,omitempty"`
	State interface{} `json:"State,omitempty"`
	Customerusername interface{} `json:"CustomerUserName,omitempty"`
	Dnsipaddresses interface{} `json:"DnsIpAddresses,omitempty"`
}

// StartWorkspacesResult represents the StartWorkspacesResult schema from the OpenAPI specification
type StartWorkspacesResult struct {
	Failedrequests interface{} `json:"FailedRequests,omitempty"`
}

// DescribeConnectionAliasesRequest represents the DescribeConnectionAliasesRequest schema from the OpenAPI specification
type DescribeConnectionAliasesRequest struct {
	Aliasids interface{} `json:"AliasIds,omitempty"`
	Limit interface{} `json:"Limit,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
}

// CreateWorkspaceBundleRequest represents the CreateWorkspaceBundleRequest schema from the OpenAPI specification
type CreateWorkspaceBundleRequest struct {
	Rootstorage RootStorage `json:"RootStorage,omitempty"` // Describes the root volume for a WorkSpace bundle.
	Tags interface{} `json:"Tags,omitempty"`
	Userstorage UserStorage `json:"UserStorage"` // Describes the user volume for a WorkSpace bundle.
	Bundledescription interface{} `json:"BundleDescription"`
	Bundlename interface{} `json:"BundleName"`
	Computetype ComputeType `json:"ComputeType"` // Describes the compute type of the bundle.
	Imageid interface{} `json:"ImageId"`
}

// WorkspacesIpGroup represents the WorkspacesIpGroup schema from the OpenAPI specification
type WorkspacesIpGroup struct {
	Groupid interface{} `json:"groupId,omitempty"`
	Groupname interface{} `json:"groupName,omitempty"`
	Userrules interface{} `json:"userRules,omitempty"`
	Groupdesc interface{} `json:"groupDesc,omitempty"`
}

// UpdateRulesOfIpGroupRequest represents the UpdateRulesOfIpGroupRequest schema from the OpenAPI specification
type UpdateRulesOfIpGroupRequest struct {
	Userrules interface{} `json:"UserRules"`
	Groupid interface{} `json:"GroupId"`
}

// UpdateConnectionAliasPermissionRequest represents the UpdateConnectionAliasPermissionRequest schema from the OpenAPI specification
type UpdateConnectionAliasPermissionRequest struct {
	Connectionaliaspermission interface{} `json:"ConnectionAliasPermission"`
	Aliasid interface{} `json:"AliasId"`
}

// StopWorkspacesResult represents the StopWorkspacesResult schema from the OpenAPI specification
type StopWorkspacesResult struct {
	Failedrequests interface{} `json:"FailedRequests,omitempty"`
}

// FailedCreateStandbyWorkspacesRequest represents the FailedCreateStandbyWorkspacesRequest schema from the OpenAPI specification
type FailedCreateStandbyWorkspacesRequest struct {
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Standbyworkspacerequest interface{} `json:"StandbyWorkspaceRequest,omitempty"`
}

// TerminateRequest represents the TerminateRequest schema from the OpenAPI specification
type TerminateRequest struct {
	Workspaceid interface{} `json:"WorkspaceId"`
}

// DescribeWorkspaceBundlesResult represents the DescribeWorkspaceBundlesResult schema from the OpenAPI specification
type DescribeWorkspaceBundlesResult struct {
	Bundles interface{} `json:"Bundles,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// WorkspaceRequest represents the WorkspaceRequest schema from the OpenAPI specification
type WorkspaceRequest struct {
	Uservolumeencryptionenabled interface{} `json:"UserVolumeEncryptionEnabled,omitempty"`
	Volumeencryptionkey interface{} `json:"VolumeEncryptionKey,omitempty"`
	Workspaceproperties interface{} `json:"WorkspaceProperties,omitempty"`
	Bundleid interface{} `json:"BundleId"`
	Directoryid interface{} `json:"DirectoryId"`
	Rootvolumeencryptionenabled interface{} `json:"RootVolumeEncryptionEnabled,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Username interface{} `json:"UserName"`
}

// RegisterWorkspaceDirectoryRequest represents the RegisterWorkspaceDirectoryRequest schema from the OpenAPI specification
type RegisterWorkspaceDirectoryRequest struct {
	Directoryid interface{} `json:"DirectoryId"`
	Enableselfservice interface{} `json:"EnableSelfService,omitempty"`
	Enableworkdocs interface{} `json:"EnableWorkDocs"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Tenancy interface{} `json:"Tenancy,omitempty"`
}

// Snapshot represents the Snapshot schema from the OpenAPI specification
type Snapshot struct {
	Snapshottime interface{} `json:"SnapshotTime,omitempty"`
}

// DescribeWorkspaceDirectoriesResult represents the DescribeWorkspaceDirectoriesResult schema from the OpenAPI specification
type DescribeWorkspaceDirectoriesResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Directories interface{} `json:"Directories,omitempty"`
}

// WorkspaceImage represents the WorkspaceImage schema from the OpenAPI specification
type WorkspaceImage struct {
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	State interface{} `json:"State,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Imageid interface{} `json:"ImageId,omitempty"`
	Updates interface{} `json:"Updates,omitempty"`
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
	Owneraccountid interface{} `json:"OwnerAccountId,omitempty"`
	Requiredtenancy interface{} `json:"RequiredTenancy,omitempty"`
	Created interface{} `json:"Created,omitempty"`
}

// CreateConnectionAliasResult represents the CreateConnectionAliasResult schema from the OpenAPI specification
type CreateConnectionAliasResult struct {
	Aliasid interface{} `json:"AliasId,omitempty"`
}

// AssociateConnectionAliasRequest represents the AssociateConnectionAliasRequest schema from the OpenAPI specification
type AssociateConnectionAliasRequest struct {
	Aliasid interface{} `json:"AliasId"`
	Resourceid interface{} `json:"ResourceId"`
}

// ListAvailableManagementCidrRangesResult represents the ListAvailableManagementCidrRangesResult schema from the OpenAPI specification
type ListAvailableManagementCidrRangesResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Managementcidrranges interface{} `json:"ManagementCidrRanges,omitempty"`
}

// RebuildWorkspacesResult represents the RebuildWorkspacesResult schema from the OpenAPI specification
type RebuildWorkspacesResult struct {
	Failedrequests interface{} `json:"FailedRequests,omitempty"`
}

// DescribeWorkspaceSnapshotsRequest represents the DescribeWorkspaceSnapshotsRequest schema from the OpenAPI specification
type DescribeWorkspaceSnapshotsRequest struct {
	Workspaceid interface{} `json:"WorkspaceId"`
}

// UpdateRulesOfIpGroupResult represents the UpdateRulesOfIpGroupResult schema from the OpenAPI specification
type UpdateRulesOfIpGroupResult struct {
}

// CreateWorkspaceBundleResult represents the CreateWorkspaceBundleResult schema from the OpenAPI specification
type CreateWorkspaceBundleResult struct {
	Workspacebundle WorkspaceBundle `json:"WorkspaceBundle,omitempty"` // Describes a WorkSpace bundle.
}

// ModifyWorkspaceCreationPropertiesResult represents the ModifyWorkspaceCreationPropertiesResult schema from the OpenAPI specification
type ModifyWorkspaceCreationPropertiesResult struct {
}

// WorkspaceConnectionStatus represents the WorkspaceConnectionStatus schema from the OpenAPI specification
type WorkspaceConnectionStatus struct {
	Lastknownuserconnectiontimestamp interface{} `json:"LastKnownUserConnectionTimestamp,omitempty"`
	Workspaceid interface{} `json:"WorkspaceId,omitempty"`
	Connectionstate interface{} `json:"ConnectionState,omitempty"`
	Connectionstatechecktimestamp interface{} `json:"ConnectionStateCheckTimestamp,omitempty"`
}

// CreateWorkspaceImageResult represents the CreateWorkspaceImageResult schema from the OpenAPI specification
type CreateWorkspaceImageResult struct {
	Imageid interface{} `json:"ImageId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
	Owneraccountid interface{} `json:"OwnerAccountId,omitempty"`
	Requiredtenancy interface{} `json:"RequiredTenancy,omitempty"`
	State interface{} `json:"State,omitempty"`
	Created interface{} `json:"Created,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// DeleteWorkspaceImageResult represents the DeleteWorkspaceImageResult schema from the OpenAPI specification
type DeleteWorkspaceImageResult struct {
}
