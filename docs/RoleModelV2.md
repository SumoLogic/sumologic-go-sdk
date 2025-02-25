# RoleModelV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the role. | 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**LogAnalyticsFilter** | Pointer to **string** | A search filter which would be applied on partitions which belong to Log Analytics product area. | [optional] 
**AuditDataFilter** | Pointer to **string** | A search filter which would be applied on partitions which belong to Audit Data product area. Help Doc : (https://help.sumologic.com/docs/manage/security/audit-index/). | [optional] 
**SecurityDataFilter** | Pointer to **string** | A search filter which would be applied on partitions which belong to Security Data product area. | [optional] 
**SelectionType** | Pointer to **string** | Describes the Permission Construct for the list of views in \&quot;selectedViews\&quot; parameter.  ### Valid Values are :    - &#x60;All&#x60; selectionType would allow access to all views in the org.   - &#x60;Allow&#x60; selectionType would allow access to specific views mentioned in \&quot;selectedViews\&quot; parameter.   - &#x60;Deny&#x60; selectionType would deny access to specific views mentioned in \&quot;selectedViews\&quot; parameter. | [optional] 
**SelectedViews** | Pointer to [**[]ViewFilterDefinition**](ViewFilterDefinition.md) | List of views which with specific view level filters in accordance to the selectionType chosen. | [optional] 
**Users** | Pointer to **[]string** | List of user identifiers to assign the role to. | [optional] 
**Capabilities** | Pointer to **[]string** | List of [capabilities](https://help.sumologic.com/docs/manage/users-roles/roles/role-capabilities/) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - downloadSearchResults   - manageIndexes   - manageDataStreams   - viewParsers   - viewDataStreams  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - ipWhitelisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist   - shareDashboardWhitelist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse   - cseViewAutomations   - cseManageContextActions   - cseViewNetworkBlocks   - cseManageInsightTags   - cseViewRules   - cseViewThreatIntelligence   - cseCommentOnInsights   - cseViewEntityGroups   - cseManageEntityConfiguration   - cseManageNetworkBlocks   - cseManageMatchLists   - cseViewCustomInsights   - cseManageActions   - cseManageAutomations   - cseManageMappings   - cseManageThreatIntelligence   - cseViewActions   - cseCreateInsights   - cseManageTagSchemas   - cseInvokeInsights   - cseManageCustomEntityType   - cseViewTagSchemas   - cseDeleteInsights   - cseManageCustomInsights   - cseViewFileAnalysis   - cseManageFileAnalysis   - cseManageEntityCriticality   - cseViewEntityCriticality   - cseViewEntity   - cseManageCustomInsightStatuses   - cseViewContextActions   - cseViewMappings   - cseViewCustomEntityType   - cseManageEntityGroups   - cseViewCustomInsightStatuses   - cseViewEnrichments   - cseManageInsightSignals   - cseManageRules   - cseManageArtifacts   - cseViewMatchLists   - cseManageInsightPolicy   - cseManageEnrichments   - cseViewEntityConfiguration   - cseManageEntity   - cseExecuteAutomations   - cseManageSuppressedEntities   - cseManageInsightStatus     - cseManageInsightAssignee   - cseManageFavoriteFields   - cseViewSuppressedEntities  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts   - viewMutingSchedules   - manageMutingSchedules   - adminMonitorsV2  ### SLO   - viewSlos   - manageSlos  ### CloudSoar   - cloudSoarPlaybooksAccess   - cloudSoarNotificationConfigure   - cloudSoarReportAll   - cloudSoarIncidentTriageAccess   - cloudSoarIncidentTaskView   - cloudSoarIncidentChangeOwnership   - cloudSoarIncidentNotesEdit   - cloudSoarAPIEmailEdit   - cloudSoarIncidentTemplatesAccess   - cloudSoarIncidentPlaybooksManage   - cloudSoarGeneralConfigure   - cloudSoarEntitiesAccess   - cloudSoarEntitiesBulkPhysicalDelete   - cloudSoarIncidentAttachmentsAccess   - cloudSoarAppCentralAccess   - cloudSoarBridgeMonitoringAccess   - viewCloudSoar   - cloudSoarIncidentView   - cloudSoarObservabilityAccess   - cloudSoarAPIEmailRead   - cloudSoarAppCentralExport   - cloudSoarWidgetsAll   - cloudSoarIncidentTaskReassign   - cloudSoarIntegrationsAccess   - cloudSoarCustomizationIncidentLabels   - cloudSoarAutomationRulesConfigure   - cloudSoarIncidentTaskAccessAll   - cloudSoarAuditAndInformationConfigureAuditTrail   - cloudSoarIncidentTriageEdit   - cloudSoarIncidentEdit   - cloudSoarNotificationTriage   - cloudSoarIncidentTriageBulkPhysicalDelete   - cloudSoarIncidentNotesAccess   - cloudSoarAPIUse   - cloudSoarIncidentPlaybooksEdit   - cloudSoarDashboardAll   - cloudSoarEntitiesManage   - cloudSoarIncidentTemplatesConfigure   - cloudSoarIncidentTriageAccessAll   - cloudSoarPlaybooksConfigure   - cloudSoarIncidentAccessAll   - cloudSoarCustomizationLogo   - cloudSoarIncidentTaskAccess   - cloudSoarIncidentTriageView   - cloudSoarIntegrationsConfigure   - cloudSoarIncidentManageInvestigators   - cloudSoarIncidentAccess   - cloudSoarAuditAndInformationLicenseInformation   - cloudSoarIncidentBulkOperations   - cloudSoarCustomizationFields   - cloudSoarIncidentTaskEdit   - cloudSoarDashboardAccess   - cloudSoarIncidentAttachmentsEdit   - cloudSoarIncidentFoldersEdit   - cloudSoarUserManagementGroups   - cloudSoarIncidentPlaybooksAccess   - cloudSoarIncidentWarRoomUse   - cloudSoarReportAccess   - cloudSoarAuditAndInformationAuditTrail   - cloudSoarAutomationRulesAccess   - cloudSoarIncidentTriageChangeOwnership   - cloudSoarObservabilityManagement | [optional] 
**AutofillDependencies** | Pointer to **bool** | Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies. | [optional] [default to true]
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**Id** | **string** | Unique identifier for the role. | 
**SystemDefined** | Pointer to **bool** | Role is system or user defined. | [optional] 

## Methods

### NewRoleModelV2

`func NewRoleModelV2(name string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string, ) *RoleModelV2`

NewRoleModelV2 instantiates a new RoleModelV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleModelV2WithDefaults

`func NewRoleModelV2WithDefaults() *RoleModelV2`

NewRoleModelV2WithDefaults instantiates a new RoleModelV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleModelV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleModelV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleModelV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RoleModelV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleModelV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleModelV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleModelV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogAnalyticsFilter

`func (o *RoleModelV2) GetLogAnalyticsFilter() string`

GetLogAnalyticsFilter returns the LogAnalyticsFilter field if non-nil, zero value otherwise.

### GetLogAnalyticsFilterOk

`func (o *RoleModelV2) GetLogAnalyticsFilterOk() (*string, bool)`

GetLogAnalyticsFilterOk returns a tuple with the LogAnalyticsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAnalyticsFilter

`func (o *RoleModelV2) SetLogAnalyticsFilter(v string)`

SetLogAnalyticsFilter sets LogAnalyticsFilter field to given value.

### HasLogAnalyticsFilter

`func (o *RoleModelV2) HasLogAnalyticsFilter() bool`

HasLogAnalyticsFilter returns a boolean if a field has been set.

### GetAuditDataFilter

`func (o *RoleModelV2) GetAuditDataFilter() string`

GetAuditDataFilter returns the AuditDataFilter field if non-nil, zero value otherwise.

### GetAuditDataFilterOk

`func (o *RoleModelV2) GetAuditDataFilterOk() (*string, bool)`

GetAuditDataFilterOk returns a tuple with the AuditDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditDataFilter

`func (o *RoleModelV2) SetAuditDataFilter(v string)`

SetAuditDataFilter sets AuditDataFilter field to given value.

### HasAuditDataFilter

`func (o *RoleModelV2) HasAuditDataFilter() bool`

HasAuditDataFilter returns a boolean if a field has been set.

### GetSecurityDataFilter

`func (o *RoleModelV2) GetSecurityDataFilter() string`

GetSecurityDataFilter returns the SecurityDataFilter field if non-nil, zero value otherwise.

### GetSecurityDataFilterOk

`func (o *RoleModelV2) GetSecurityDataFilterOk() (*string, bool)`

GetSecurityDataFilterOk returns a tuple with the SecurityDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityDataFilter

`func (o *RoleModelV2) SetSecurityDataFilter(v string)`

SetSecurityDataFilter sets SecurityDataFilter field to given value.

### HasSecurityDataFilter

`func (o *RoleModelV2) HasSecurityDataFilter() bool`

HasSecurityDataFilter returns a boolean if a field has been set.

### GetSelectionType

`func (o *RoleModelV2) GetSelectionType() string`

GetSelectionType returns the SelectionType field if non-nil, zero value otherwise.

### GetSelectionTypeOk

`func (o *RoleModelV2) GetSelectionTypeOk() (*string, bool)`

GetSelectionTypeOk returns a tuple with the SelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionType

`func (o *RoleModelV2) SetSelectionType(v string)`

SetSelectionType sets SelectionType field to given value.

### HasSelectionType

`func (o *RoleModelV2) HasSelectionType() bool`

HasSelectionType returns a boolean if a field has been set.

### GetSelectedViews

`func (o *RoleModelV2) GetSelectedViews() []ViewFilterDefinition`

GetSelectedViews returns the SelectedViews field if non-nil, zero value otherwise.

### GetSelectedViewsOk

`func (o *RoleModelV2) GetSelectedViewsOk() (*[]ViewFilterDefinition, bool)`

GetSelectedViewsOk returns a tuple with the SelectedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedViews

`func (o *RoleModelV2) SetSelectedViews(v []ViewFilterDefinition)`

SetSelectedViews sets SelectedViews field to given value.

### HasSelectedViews

`func (o *RoleModelV2) HasSelectedViews() bool`

HasSelectedViews returns a boolean if a field has been set.

### GetUsers

`func (o *RoleModelV2) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RoleModelV2) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RoleModelV2) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *RoleModelV2) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetCapabilities

`func (o *RoleModelV2) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *RoleModelV2) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *RoleModelV2) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *RoleModelV2) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAutofillDependencies

`func (o *RoleModelV2) GetAutofillDependencies() bool`

GetAutofillDependencies returns the AutofillDependencies field if non-nil, zero value otherwise.

### GetAutofillDependenciesOk

`func (o *RoleModelV2) GetAutofillDependenciesOk() (*bool, bool)`

GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillDependencies

`func (o *RoleModelV2) SetAutofillDependencies(v bool)`

SetAutofillDependencies sets AutofillDependencies field to given value.

### HasAutofillDependencies

`func (o *RoleModelV2) HasAutofillDependencies() bool`

HasAutofillDependencies returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RoleModelV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoleModelV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoleModelV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *RoleModelV2) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RoleModelV2) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RoleModelV2) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *RoleModelV2) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RoleModelV2) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RoleModelV2) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *RoleModelV2) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *RoleModelV2) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *RoleModelV2) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetId

`func (o *RoleModelV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleModelV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleModelV2) SetId(v string)`

SetId sets Id field to given value.


### GetSystemDefined

`func (o *RoleModelV2) GetSystemDefined() bool`

GetSystemDefined returns the SystemDefined field if non-nil, zero value otherwise.

### GetSystemDefinedOk

`func (o *RoleModelV2) GetSystemDefinedOk() (*bool, bool)`

GetSystemDefinedOk returns a tuple with the SystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDefined

`func (o *RoleModelV2) SetSystemDefined(v bool)`

SetSystemDefined sets SystemDefined field to given value.

### HasSystemDefined

`func (o *RoleModelV2) HasSystemDefined() bool`

HasSystemDefined returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


