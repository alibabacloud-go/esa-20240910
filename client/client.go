// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type WafBatchRuleShared struct {
	Action      *string                    `json:"Action,omitempty" xml:"Action,omitempty"`
	Actions     *WafBatchRuleSharedActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Struct"`
	CrossSiteId *int64                     `json:"CrossSiteId,omitempty" xml:"CrossSiteId,omitempty"`
	Expression  *string                    `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Match       *WafRuleMatch2             `json:"Match,omitempty" xml:"Match,omitempty"`
	Mode        *string                    `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Name        *string                    `json:"Name,omitempty" xml:"Name,omitempty"`
	Target      *string                    `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s WafBatchRuleShared) String() string {
	return tea.Prettify(s)
}

func (s WafBatchRuleShared) GoString() string {
	return s.String()
}

func (s *WafBatchRuleShared) SetAction(v string) *WafBatchRuleShared {
	s.Action = &v
	return s
}

func (s *WafBatchRuleShared) SetActions(v *WafBatchRuleSharedActions) *WafBatchRuleShared {
	s.Actions = v
	return s
}

func (s *WafBatchRuleShared) SetCrossSiteId(v int64) *WafBatchRuleShared {
	s.CrossSiteId = &v
	return s
}

func (s *WafBatchRuleShared) SetExpression(v string) *WafBatchRuleShared {
	s.Expression = &v
	return s
}

func (s *WafBatchRuleShared) SetMatch(v *WafRuleMatch2) *WafBatchRuleShared {
	s.Match = v
	return s
}

func (s *WafBatchRuleShared) SetMode(v string) *WafBatchRuleShared {
	s.Mode = &v
	return s
}

func (s *WafBatchRuleShared) SetName(v string) *WafBatchRuleShared {
	s.Name = &v
	return s
}

func (s *WafBatchRuleShared) SetTarget(v string) *WafBatchRuleShared {
	s.Target = &v
	return s
}

type WafBatchRuleSharedActions struct {
	Response *WafBatchRuleSharedActionsResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
}

func (s WafBatchRuleSharedActions) String() string {
	return tea.Prettify(s)
}

func (s WafBatchRuleSharedActions) GoString() string {
	return s.String()
}

func (s *WafBatchRuleSharedActions) SetResponse(v *WafBatchRuleSharedActionsResponse) *WafBatchRuleSharedActions {
	s.Response = v
	return s
}

type WafBatchRuleSharedActionsResponse struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	Id   *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s WafBatchRuleSharedActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s WafBatchRuleSharedActionsResponse) GoString() string {
	return s.String()
}

func (s *WafBatchRuleSharedActionsResponse) SetCode(v int32) *WafBatchRuleSharedActionsResponse {
	s.Code = &v
	return s
}

func (s *WafBatchRuleSharedActionsResponse) SetId(v int64) *WafBatchRuleSharedActionsResponse {
	s.Id = &v
	return s
}

type WafQuotaInteger struct {
	Equal              *int32 `json:"Equal,omitempty" xml:"Equal,omitempty"`
	GreaterThan        *int32 `json:"GreaterThan,omitempty" xml:"GreaterThan,omitempty"`
	GreaterThanOrEqual *int32 `json:"GreaterThanOrEqual,omitempty" xml:"GreaterThanOrEqual,omitempty"`
	LessThan           *int32 `json:"LessThan,omitempty" xml:"LessThan,omitempty"`
	LessThanOrEqual    *int32 `json:"LessThanOrEqual,omitempty" xml:"LessThanOrEqual,omitempty"`
}

func (s WafQuotaInteger) String() string {
	return tea.Prettify(s)
}

func (s WafQuotaInteger) GoString() string {
	return s.String()
}

func (s *WafQuotaInteger) SetEqual(v int32) *WafQuotaInteger {
	s.Equal = &v
	return s
}

func (s *WafQuotaInteger) SetGreaterThan(v int32) *WafQuotaInteger {
	s.GreaterThan = &v
	return s
}

func (s *WafQuotaInteger) SetGreaterThanOrEqual(v int32) *WafQuotaInteger {
	s.GreaterThanOrEqual = &v
	return s
}

func (s *WafQuotaInteger) SetLessThan(v int32) *WafQuotaInteger {
	s.LessThan = &v
	return s
}

func (s *WafQuotaInteger) SetLessThanOrEqual(v int32) *WafQuotaInteger {
	s.LessThanOrEqual = &v
	return s
}

type WafQuotaString struct {
	Regexp *string `json:"Regexp,omitempty" xml:"Regexp,omitempty"`
}

func (s WafQuotaString) String() string {
	return tea.Prettify(s)
}

func (s WafQuotaString) GoString() string {
	return s.String()
}

func (s *WafQuotaString) SetRegexp(v string) *WafQuotaString {
	s.Regexp = &v
	return s
}

type WafRuleConfig struct {
	Action          *string                         `json:"Action,omitempty" xml:"Action,omitempty"`
	Actions         *WafRuleConfigActions           `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Struct"`
	AppPackage      *WafRuleConfigAppPackage        `json:"AppPackage,omitempty" xml:"AppPackage,omitempty" type:"Struct"`
	AppSdk          *WafRuleConfigAppSdk            `json:"AppSdk,omitempty" xml:"AppSdk,omitempty" type:"Struct"`
	Expression      *string                         `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Id              *int64                          `json:"Id,omitempty" xml:"Id,omitempty"`
	ManagedGroupId  *int64                          `json:"ManagedGroupId,omitempty" xml:"ManagedGroupId,omitempty"`
	ManagedList     *string                         `json:"ManagedList,omitempty" xml:"ManagedList,omitempty"`
	ManagedRulesets []*WafRuleConfigManagedRulesets `json:"ManagedRulesets,omitempty" xml:"ManagedRulesets,omitempty" type:"Repeated"`
	Match           *WafRuleMatch                   `json:"Match,omitempty" xml:"Match,omitempty"`
	Name            *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	RateLimit       *WafRuleConfigRateLimit         `json:"RateLimit,omitempty" xml:"RateLimit,omitempty" type:"Struct"`
	Sigchl          []*string                       `json:"Sigchl,omitempty" xml:"Sigchl,omitempty" type:"Repeated"`
	Status          *string                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Timer           *WafTimer                       `json:"Timer,omitempty" xml:"Timer,omitempty"`
	Type            *string                         `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s WafRuleConfig) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfig) GoString() string {
	return s.String()
}

func (s *WafRuleConfig) SetAction(v string) *WafRuleConfig {
	s.Action = &v
	return s
}

func (s *WafRuleConfig) SetActions(v *WafRuleConfigActions) *WafRuleConfig {
	s.Actions = v
	return s
}

func (s *WafRuleConfig) SetAppPackage(v *WafRuleConfigAppPackage) *WafRuleConfig {
	s.AppPackage = v
	return s
}

func (s *WafRuleConfig) SetAppSdk(v *WafRuleConfigAppSdk) *WafRuleConfig {
	s.AppSdk = v
	return s
}

func (s *WafRuleConfig) SetExpression(v string) *WafRuleConfig {
	s.Expression = &v
	return s
}

func (s *WafRuleConfig) SetId(v int64) *WafRuleConfig {
	s.Id = &v
	return s
}

func (s *WafRuleConfig) SetManagedGroupId(v int64) *WafRuleConfig {
	s.ManagedGroupId = &v
	return s
}

func (s *WafRuleConfig) SetManagedList(v string) *WafRuleConfig {
	s.ManagedList = &v
	return s
}

func (s *WafRuleConfig) SetManagedRulesets(v []*WafRuleConfigManagedRulesets) *WafRuleConfig {
	s.ManagedRulesets = v
	return s
}

func (s *WafRuleConfig) SetMatch(v *WafRuleMatch) *WafRuleConfig {
	s.Match = v
	return s
}

func (s *WafRuleConfig) SetName(v string) *WafRuleConfig {
	s.Name = &v
	return s
}

func (s *WafRuleConfig) SetRateLimit(v *WafRuleConfigRateLimit) *WafRuleConfig {
	s.RateLimit = v
	return s
}

func (s *WafRuleConfig) SetSigchl(v []*string) *WafRuleConfig {
	s.Sigchl = v
	return s
}

func (s *WafRuleConfig) SetStatus(v string) *WafRuleConfig {
	s.Status = &v
	return s
}

func (s *WafRuleConfig) SetTimer(v *WafTimer) *WafRuleConfig {
	s.Timer = v
	return s
}

func (s *WafRuleConfig) SetType(v string) *WafRuleConfig {
	s.Type = &v
	return s
}

type WafRuleConfigActions struct {
	Bypass   *WafRuleConfigActionsBypass   `json:"Bypass,omitempty" xml:"Bypass,omitempty" type:"Struct"`
	Response *WafRuleConfigActionsResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
}

func (s WafRuleConfigActions) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigActions) GoString() string {
	return s.String()
}

func (s *WafRuleConfigActions) SetBypass(v *WafRuleConfigActionsBypass) *WafRuleConfigActions {
	s.Bypass = v
	return s
}

func (s *WafRuleConfigActions) SetResponse(v *WafRuleConfigActionsResponse) *WafRuleConfigActions {
	s.Response = v
	return s
}

type WafRuleConfigActionsBypass struct {
	RegularRules []*int64  `json:"RegularRules,omitempty" xml:"RegularRules,omitempty" type:"Repeated"`
	RegularTypes []*string `json:"RegularTypes,omitempty" xml:"RegularTypes,omitempty" type:"Repeated"`
	Skip         *string   `json:"Skip,omitempty" xml:"Skip,omitempty"`
	Tags         []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s WafRuleConfigActionsBypass) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigActionsBypass) GoString() string {
	return s.String()
}

func (s *WafRuleConfigActionsBypass) SetRegularRules(v []*int64) *WafRuleConfigActionsBypass {
	s.RegularRules = v
	return s
}

func (s *WafRuleConfigActionsBypass) SetRegularTypes(v []*string) *WafRuleConfigActionsBypass {
	s.RegularTypes = v
	return s
}

func (s *WafRuleConfigActionsBypass) SetSkip(v string) *WafRuleConfigActionsBypass {
	s.Skip = &v
	return s
}

func (s *WafRuleConfigActionsBypass) SetTags(v []*string) *WafRuleConfigActionsBypass {
	s.Tags = v
	return s
}

type WafRuleConfigActionsResponse struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	Id   *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s WafRuleConfigActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigActionsResponse) GoString() string {
	return s.String()
}

func (s *WafRuleConfigActionsResponse) SetCode(v int32) *WafRuleConfigActionsResponse {
	s.Code = &v
	return s
}

func (s *WafRuleConfigActionsResponse) SetId(v int64) *WafRuleConfigActionsResponse {
	s.Id = &v
	return s
}

type WafRuleConfigAppPackage struct {
	PackageSigns []*WafRuleConfigAppPackagePackageSigns `json:"PackageSigns,omitempty" xml:"PackageSigns,omitempty" type:"Repeated"`
}

func (s WafRuleConfigAppPackage) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigAppPackage) GoString() string {
	return s.String()
}

func (s *WafRuleConfigAppPackage) SetPackageSigns(v []*WafRuleConfigAppPackagePackageSigns) *WafRuleConfigAppPackage {
	s.PackageSigns = v
	return s
}

type WafRuleConfigAppPackagePackageSigns struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s WafRuleConfigAppPackagePackageSigns) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigAppPackagePackageSigns) GoString() string {
	return s.String()
}

func (s *WafRuleConfigAppPackagePackageSigns) SetName(v string) *WafRuleConfigAppPackagePackageSigns {
	s.Name = &v
	return s
}

func (s *WafRuleConfigAppPackagePackageSigns) SetSign(v string) *WafRuleConfigAppPackagePackageSigns {
	s.Sign = &v
	return s
}

type WafRuleConfigAppSdk struct {
	CustomSign       *WafRuleConfigAppSdkCustomSign `json:"CustomSign,omitempty" xml:"CustomSign,omitempty" type:"Struct"`
	CustomSignStatus *string                        `json:"CustomSignStatus,omitempty" xml:"CustomSignStatus,omitempty"`
	FeatureAbnormal  []*string                      `json:"FeatureAbnormal,omitempty" xml:"FeatureAbnormal,omitempty" type:"Repeated"`
}

func (s WafRuleConfigAppSdk) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigAppSdk) GoString() string {
	return s.String()
}

func (s *WafRuleConfigAppSdk) SetCustomSign(v *WafRuleConfigAppSdkCustomSign) *WafRuleConfigAppSdk {
	s.CustomSign = v
	return s
}

func (s *WafRuleConfigAppSdk) SetCustomSignStatus(v string) *WafRuleConfigAppSdk {
	s.CustomSignStatus = &v
	return s
}

func (s *WafRuleConfigAppSdk) SetFeatureAbnormal(v []*string) *WafRuleConfigAppSdk {
	s.FeatureAbnormal = v
	return s
}

type WafRuleConfigAppSdkCustomSign struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s WafRuleConfigAppSdkCustomSign) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigAppSdkCustomSign) GoString() string {
	return s.String()
}

func (s *WafRuleConfigAppSdkCustomSign) SetKey(v string) *WafRuleConfigAppSdkCustomSign {
	s.Key = &v
	return s
}

func (s *WafRuleConfigAppSdkCustomSign) SetValue(v string) *WafRuleConfigAppSdkCustomSign {
	s.Value = &v
	return s
}

type WafRuleConfigManagedRulesets struct {
	Action          *string                                     `json:"Action,omitempty" xml:"Action,omitempty"`
	AttackType      *int32                                      `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	ManagedRules    []*WafRuleConfigManagedRulesetsManagedRules `json:"ManagedRules,omitempty" xml:"ManagedRules,omitempty" type:"Repeated"`
	NumberEnabled   *int32                                      `json:"NumberEnabled,omitempty" xml:"NumberEnabled,omitempty"`
	NumberTotal     *int32                                      `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
	ProtectionLevel *int32                                      `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
}

func (s WafRuleConfigManagedRulesets) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigManagedRulesets) GoString() string {
	return s.String()
}

func (s *WafRuleConfigManagedRulesets) SetAction(v string) *WafRuleConfigManagedRulesets {
	s.Action = &v
	return s
}

func (s *WafRuleConfigManagedRulesets) SetAttackType(v int32) *WafRuleConfigManagedRulesets {
	s.AttackType = &v
	return s
}

func (s *WafRuleConfigManagedRulesets) SetManagedRules(v []*WafRuleConfigManagedRulesetsManagedRules) *WafRuleConfigManagedRulesets {
	s.ManagedRules = v
	return s
}

func (s *WafRuleConfigManagedRulesets) SetNumberEnabled(v int32) *WafRuleConfigManagedRulesets {
	s.NumberEnabled = &v
	return s
}

func (s *WafRuleConfigManagedRulesets) SetNumberTotal(v int32) *WafRuleConfigManagedRulesets {
	s.NumberTotal = &v
	return s
}

func (s *WafRuleConfigManagedRulesets) SetProtectionLevel(v int32) *WafRuleConfigManagedRulesets {
	s.ProtectionLevel = &v
	return s
}

type WafRuleConfigManagedRulesetsManagedRules struct {
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s WafRuleConfigManagedRulesetsManagedRules) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigManagedRulesetsManagedRules) GoString() string {
	return s.String()
}

func (s *WafRuleConfigManagedRulesetsManagedRules) SetAction(v string) *WafRuleConfigManagedRulesetsManagedRules {
	s.Action = &v
	return s
}

func (s *WafRuleConfigManagedRulesetsManagedRules) SetId(v int64) *WafRuleConfigManagedRulesetsManagedRules {
	s.Id = &v
	return s
}

func (s *WafRuleConfigManagedRulesetsManagedRules) SetStatus(v string) *WafRuleConfigManagedRulesetsManagedRules {
	s.Status = &v
	return s
}

type WafRuleConfigRateLimit struct {
	Characteristics *WafRuleMatch2                   `json:"Characteristics,omitempty" xml:"Characteristics,omitempty"`
	Interval        *int32                           `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OnHit           *bool                            `json:"OnHit,omitempty" xml:"OnHit,omitempty"`
	TTL             *int32                           `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Threshold       *WafRuleConfigRateLimitThreshold `json:"Threshold,omitempty" xml:"Threshold,omitempty" type:"Struct"`
}

func (s WafRuleConfigRateLimit) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigRateLimit) GoString() string {
	return s.String()
}

func (s *WafRuleConfigRateLimit) SetCharacteristics(v *WafRuleMatch2) *WafRuleConfigRateLimit {
	s.Characteristics = v
	return s
}

func (s *WafRuleConfigRateLimit) SetInterval(v int32) *WafRuleConfigRateLimit {
	s.Interval = &v
	return s
}

func (s *WafRuleConfigRateLimit) SetOnHit(v bool) *WafRuleConfigRateLimit {
	s.OnHit = &v
	return s
}

func (s *WafRuleConfigRateLimit) SetTTL(v int32) *WafRuleConfigRateLimit {
	s.TTL = &v
	return s
}

func (s *WafRuleConfigRateLimit) SetThreshold(v *WafRuleConfigRateLimitThreshold) *WafRuleConfigRateLimit {
	s.Threshold = v
	return s
}

type WafRuleConfigRateLimitThreshold struct {
	DistinctManagedRules *int32                                         `json:"DistinctManagedRules,omitempty" xml:"DistinctManagedRules,omitempty"`
	ManagedRulesBlocked  *int32                                         `json:"ManagedRulesBlocked,omitempty" xml:"ManagedRulesBlocked,omitempty"`
	Request              *int32                                         `json:"Request,omitempty" xml:"Request,omitempty"`
	ResponseStatus       *WafRuleConfigRateLimitThresholdResponseStatus `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty" type:"Struct"`
	Traffic              *string                                        `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s WafRuleConfigRateLimitThreshold) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigRateLimitThreshold) GoString() string {
	return s.String()
}

func (s *WafRuleConfigRateLimitThreshold) SetDistinctManagedRules(v int32) *WafRuleConfigRateLimitThreshold {
	s.DistinctManagedRules = &v
	return s
}

func (s *WafRuleConfigRateLimitThreshold) SetManagedRulesBlocked(v int32) *WafRuleConfigRateLimitThreshold {
	s.ManagedRulesBlocked = &v
	return s
}

func (s *WafRuleConfigRateLimitThreshold) SetRequest(v int32) *WafRuleConfigRateLimitThreshold {
	s.Request = &v
	return s
}

func (s *WafRuleConfigRateLimitThreshold) SetResponseStatus(v *WafRuleConfigRateLimitThresholdResponseStatus) *WafRuleConfigRateLimitThreshold {
	s.ResponseStatus = v
	return s
}

func (s *WafRuleConfigRateLimitThreshold) SetTraffic(v string) *WafRuleConfigRateLimitThreshold {
	s.Traffic = &v
	return s
}

type WafRuleConfigRateLimitThresholdResponseStatus struct {
	Code  *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s WafRuleConfigRateLimitThresholdResponseStatus) String() string {
	return tea.Prettify(s)
}

func (s WafRuleConfigRateLimitThresholdResponseStatus) GoString() string {
	return s.String()
}

func (s *WafRuleConfigRateLimitThresholdResponseStatus) SetCode(v int32) *WafRuleConfigRateLimitThresholdResponseStatus {
	s.Code = &v
	return s
}

func (s *WafRuleConfigRateLimitThresholdResponseStatus) SetCount(v int32) *WafRuleConfigRateLimitThresholdResponseStatus {
	s.Count = &v
	return s
}

func (s *WafRuleConfigRateLimitThresholdResponseStatus) SetRatio(v int32) *WafRuleConfigRateLimitThresholdResponseStatus {
	s.Ratio = &v
	return s
}

type WafRuleMatch struct {
	ConvertToLower *bool           `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	Criteria       []*WafRuleMatch `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic          *string         `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchOperator  *string         `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	MatchType      *string         `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MatchValue     interface{}     `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	Negate         *bool           `json:"Negate,omitempty" xml:"Negate,omitempty"`
}

func (s WafRuleMatch) String() string {
	return tea.Prettify(s)
}

func (s WafRuleMatch) GoString() string {
	return s.String()
}

func (s *WafRuleMatch) SetConvertToLower(v bool) *WafRuleMatch {
	s.ConvertToLower = &v
	return s
}

func (s *WafRuleMatch) SetCriteria(v []*WafRuleMatch) *WafRuleMatch {
	s.Criteria = v
	return s
}

func (s *WafRuleMatch) SetLogic(v string) *WafRuleMatch {
	s.Logic = &v
	return s
}

func (s *WafRuleMatch) SetMatchOperator(v string) *WafRuleMatch {
	s.MatchOperator = &v
	return s
}

func (s *WafRuleMatch) SetMatchType(v string) *WafRuleMatch {
	s.MatchType = &v
	return s
}

func (s *WafRuleMatch) SetMatchValue(v interface{}) *WafRuleMatch {
	s.MatchValue = v
	return s
}

func (s *WafRuleMatch) SetNegate(v bool) *WafRuleMatch {
	s.Negate = &v
	return s
}

type WafRuleMatch2 struct {
	ConvertToLower *bool                    `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	Criteria       []*WafRuleMatch2Criteria `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic          *string                  `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchOperator  *string                  `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	MatchType      *string                  `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MatchValue     interface{}              `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	Negate         *bool                    `json:"Negate,omitempty" xml:"Negate,omitempty"`
}

func (s WafRuleMatch2) String() string {
	return tea.Prettify(s)
}

func (s WafRuleMatch2) GoString() string {
	return s.String()
}

func (s *WafRuleMatch2) SetConvertToLower(v bool) *WafRuleMatch2 {
	s.ConvertToLower = &v
	return s
}

func (s *WafRuleMatch2) SetCriteria(v []*WafRuleMatch2Criteria) *WafRuleMatch2 {
	s.Criteria = v
	return s
}

func (s *WafRuleMatch2) SetLogic(v string) *WafRuleMatch2 {
	s.Logic = &v
	return s
}

func (s *WafRuleMatch2) SetMatchOperator(v string) *WafRuleMatch2 {
	s.MatchOperator = &v
	return s
}

func (s *WafRuleMatch2) SetMatchType(v string) *WafRuleMatch2 {
	s.MatchType = &v
	return s
}

func (s *WafRuleMatch2) SetMatchValue(v interface{}) *WafRuleMatch2 {
	s.MatchValue = v
	return s
}

func (s *WafRuleMatch2) SetNegate(v bool) *WafRuleMatch2 {
	s.Negate = &v
	return s
}

type WafRuleMatch2Criteria struct {
	ConvertToLower *bool                            `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	Criteria       []*WafRuleMatch2CriteriaCriteria `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic          *string                          `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchOperator  *string                          `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	MatchType      *string                          `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MatchValue     interface{}                      `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	Negate         *bool                            `json:"Negate,omitempty" xml:"Negate,omitempty"`
}

func (s WafRuleMatch2Criteria) String() string {
	return tea.Prettify(s)
}

func (s WafRuleMatch2Criteria) GoString() string {
	return s.String()
}

func (s *WafRuleMatch2Criteria) SetConvertToLower(v bool) *WafRuleMatch2Criteria {
	s.ConvertToLower = &v
	return s
}

func (s *WafRuleMatch2Criteria) SetCriteria(v []*WafRuleMatch2CriteriaCriteria) *WafRuleMatch2Criteria {
	s.Criteria = v
	return s
}

func (s *WafRuleMatch2Criteria) SetLogic(v string) *WafRuleMatch2Criteria {
	s.Logic = &v
	return s
}

func (s *WafRuleMatch2Criteria) SetMatchOperator(v string) *WafRuleMatch2Criteria {
	s.MatchOperator = &v
	return s
}

func (s *WafRuleMatch2Criteria) SetMatchType(v string) *WafRuleMatch2Criteria {
	s.MatchType = &v
	return s
}

func (s *WafRuleMatch2Criteria) SetMatchValue(v interface{}) *WafRuleMatch2Criteria {
	s.MatchValue = v
	return s
}

func (s *WafRuleMatch2Criteria) SetNegate(v bool) *WafRuleMatch2Criteria {
	s.Negate = &v
	return s
}

type WafRuleMatch2CriteriaCriteria struct {
	ConvertToLower *bool                                    `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	Criteria       []*WafRuleMatch2CriteriaCriteriaCriteria `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic          *string                                  `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchOperator  *string                                  `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	MatchType      *string                                  `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MatchValue     interface{}                              `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	Negate         *bool                                    `json:"Negate,omitempty" xml:"Negate,omitempty"`
}

func (s WafRuleMatch2CriteriaCriteria) String() string {
	return tea.Prettify(s)
}

func (s WafRuleMatch2CriteriaCriteria) GoString() string {
	return s.String()
}

func (s *WafRuleMatch2CriteriaCriteria) SetConvertToLower(v bool) *WafRuleMatch2CriteriaCriteria {
	s.ConvertToLower = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetCriteria(v []*WafRuleMatch2CriteriaCriteriaCriteria) *WafRuleMatch2CriteriaCriteria {
	s.Criteria = v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetLogic(v string) *WafRuleMatch2CriteriaCriteria {
	s.Logic = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetMatchOperator(v string) *WafRuleMatch2CriteriaCriteria {
	s.MatchOperator = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetMatchType(v string) *WafRuleMatch2CriteriaCriteria {
	s.MatchType = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetMatchValue(v interface{}) *WafRuleMatch2CriteriaCriteria {
	s.MatchValue = v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetNegate(v bool) *WafRuleMatch2CriteriaCriteria {
	s.Negate = &v
	return s
}

type WafRuleMatch2CriteriaCriteriaCriteria struct {
	ConvertToLower *bool       `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	MatchOperator  *string     `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	MatchType      *string     `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MatchValue     interface{} `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	Negate         *bool       `json:"Negate,omitempty" xml:"Negate,omitempty"`
}

func (s WafRuleMatch2CriteriaCriteriaCriteria) String() string {
	return tea.Prettify(s)
}

func (s WafRuleMatch2CriteriaCriteriaCriteria) GoString() string {
	return s.String()
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) SetConvertToLower(v bool) *WafRuleMatch2CriteriaCriteriaCriteria {
	s.ConvertToLower = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) SetMatchOperator(v string) *WafRuleMatch2CriteriaCriteriaCriteria {
	s.MatchOperator = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) SetMatchType(v string) *WafRuleMatch2CriteriaCriteriaCriteria {
	s.MatchType = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) SetMatchValue(v interface{}) *WafRuleMatch2CriteriaCriteriaCriteria {
	s.MatchValue = v
	return s
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) SetNegate(v bool) *WafRuleMatch2CriteriaCriteriaCriteria {
	s.Negate = &v
	return s
}

type WafSiteSettings struct {
	AddBotProtectionHeaders *WafSiteSettingsAddBotProtectionHeaders `json:"AddBotProtectionHeaders,omitempty" xml:"AddBotProtectionHeaders,omitempty" type:"Struct"`
	AddSecurityHeaders      *WafSiteSettingsAddSecurityHeaders      `json:"AddSecurityHeaders,omitempty" xml:"AddSecurityHeaders,omitempty" type:"Struct"`
	ClientIpIdentifier      *WafSiteSettingsClientIpIdentifier      `json:"ClientIpIdentifier,omitempty" xml:"ClientIpIdentifier,omitempty" type:"Struct"`
	SecurityLevel           *WafSiteSettingsSecurityLevel           `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty" type:"Struct"`
}

func (s WafSiteSettings) String() string {
	return tea.Prettify(s)
}

func (s WafSiteSettings) GoString() string {
	return s.String()
}

func (s *WafSiteSettings) SetAddBotProtectionHeaders(v *WafSiteSettingsAddBotProtectionHeaders) *WafSiteSettings {
	s.AddBotProtectionHeaders = v
	return s
}

func (s *WafSiteSettings) SetAddSecurityHeaders(v *WafSiteSettingsAddSecurityHeaders) *WafSiteSettings {
	s.AddSecurityHeaders = v
	return s
}

func (s *WafSiteSettings) SetClientIpIdentifier(v *WafSiteSettingsClientIpIdentifier) *WafSiteSettings {
	s.ClientIpIdentifier = v
	return s
}

func (s *WafSiteSettings) SetSecurityLevel(v *WafSiteSettingsSecurityLevel) *WafSiteSettings {
	s.SecurityLevel = v
	return s
}

type WafSiteSettingsAddBotProtectionHeaders struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s WafSiteSettingsAddBotProtectionHeaders) String() string {
	return tea.Prettify(s)
}

func (s WafSiteSettingsAddBotProtectionHeaders) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsAddBotProtectionHeaders) SetEnable(v bool) *WafSiteSettingsAddBotProtectionHeaders {
	s.Enable = &v
	return s
}

type WafSiteSettingsAddSecurityHeaders struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s WafSiteSettingsAddSecurityHeaders) String() string {
	return tea.Prettify(s)
}

func (s WafSiteSettingsAddSecurityHeaders) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsAddSecurityHeaders) SetEnable(v bool) *WafSiteSettingsAddSecurityHeaders {
	s.Enable = &v
	return s
}

type WafSiteSettingsClientIpIdentifier struct {
	Headers []*string `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	Mode    *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s WafSiteSettingsClientIpIdentifier) String() string {
	return tea.Prettify(s)
}

func (s WafSiteSettingsClientIpIdentifier) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsClientIpIdentifier) SetHeaders(v []*string) *WafSiteSettingsClientIpIdentifier {
	s.Headers = v
	return s
}

func (s *WafSiteSettingsClientIpIdentifier) SetMode(v string) *WafSiteSettingsClientIpIdentifier {
	s.Mode = &v
	return s
}

type WafSiteSettingsSecurityLevel struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s WafSiteSettingsSecurityLevel) String() string {
	return tea.Prettify(s)
}

func (s WafSiteSettingsSecurityLevel) GoString() string {
	return s.String()
}

func (s *WafSiteSettingsSecurityLevel) SetValue(v string) *WafSiteSettingsSecurityLevel {
	s.Value = &v
	return s
}

type WafTimer struct {
	Periods       []*WafTimerPeriods       `json:"Periods,omitempty" xml:"Periods,omitempty" type:"Repeated"`
	Scopes        *string                  `json:"Scopes,omitempty" xml:"Scopes,omitempty"`
	WeeklyPeriods []*WafTimerWeeklyPeriods `json:"WeeklyPeriods,omitempty" xml:"WeeklyPeriods,omitempty" type:"Repeated"`
	Zone          *int32                   `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s WafTimer) String() string {
	return tea.Prettify(s)
}

func (s WafTimer) GoString() string {
	return s.String()
}

func (s *WafTimer) SetPeriods(v []*WafTimerPeriods) *WafTimer {
	s.Periods = v
	return s
}

func (s *WafTimer) SetScopes(v string) *WafTimer {
	s.Scopes = &v
	return s
}

func (s *WafTimer) SetWeeklyPeriods(v []*WafTimerWeeklyPeriods) *WafTimer {
	s.WeeklyPeriods = v
	return s
}

func (s *WafTimer) SetZone(v int32) *WafTimer {
	s.Zone = &v
	return s
}

type WafTimerPeriods struct {
	End   *string `json:"End,omitempty" xml:"End,omitempty"`
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s WafTimerPeriods) String() string {
	return tea.Prettify(s)
}

func (s WafTimerPeriods) GoString() string {
	return s.String()
}

func (s *WafTimerPeriods) SetEnd(v string) *WafTimerPeriods {
	s.End = &v
	return s
}

func (s *WafTimerPeriods) SetStart(v string) *WafTimerPeriods {
	s.Start = &v
	return s
}

type WafTimerWeeklyPeriods struct {
	DailyPeriods []*WafTimerWeeklyPeriodsDailyPeriods `json:"DailyPeriods,omitempty" xml:"DailyPeriods,omitempty" type:"Repeated"`
	Days         *string                              `json:"Days,omitempty" xml:"Days,omitempty"`
}

func (s WafTimerWeeklyPeriods) String() string {
	return tea.Prettify(s)
}

func (s WafTimerWeeklyPeriods) GoString() string {
	return s.String()
}

func (s *WafTimerWeeklyPeriods) SetDailyPeriods(v []*WafTimerWeeklyPeriodsDailyPeriods) *WafTimerWeeklyPeriods {
	s.DailyPeriods = v
	return s
}

func (s *WafTimerWeeklyPeriods) SetDays(v string) *WafTimerWeeklyPeriods {
	s.Days = &v
	return s
}

type WafTimerWeeklyPeriodsDailyPeriods struct {
	End   *string `json:"End,omitempty" xml:"End,omitempty"`
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s WafTimerWeeklyPeriodsDailyPeriods) String() string {
	return tea.Prettify(s)
}

func (s WafTimerWeeklyPeriodsDailyPeriods) GoString() string {
	return s.String()
}

func (s *WafTimerWeeklyPeriodsDailyPeriods) SetEnd(v string) *WafTimerWeeklyPeriodsDailyPeriods {
	s.End = &v
	return s
}

func (s *WafTimerWeeklyPeriodsDailyPeriods) SetStart(v string) *WafTimerWeeklyPeriodsDailyPeriods {
	s.Start = &v
	return s
}

type HttpDeliveryHeaderParamValue struct {
	StaticValue *string `json:"StaticValue,omitempty" xml:"StaticValue,omitempty"`
}

func (s HttpDeliveryHeaderParamValue) String() string {
	return tea.Prettify(s)
}

func (s HttpDeliveryHeaderParamValue) GoString() string {
	return s.String()
}

func (s *HttpDeliveryHeaderParamValue) SetStaticValue(v string) *HttpDeliveryHeaderParamValue {
	s.StaticValue = &v
	return s
}

type HttpDeliveryQueryParamValue struct {
	StaticValue *string `json:"StaticValue,omitempty" xml:"StaticValue,omitempty"`
}

func (s HttpDeliveryQueryParamValue) String() string {
	return tea.Prettify(s)
}

func (s HttpDeliveryQueryParamValue) GoString() string {
	return s.String()
}

func (s *HttpDeliveryQueryParamValue) SetStaticValue(v string) *HttpDeliveryQueryParamValue {
	s.StaticValue = &v
	return s
}

type FieldContentValue struct {
	// example:
	//
	// 1
	SortOrder *int64                        `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	FieldList []*FieldContentValueFieldList `json:"FieldList,omitempty" xml:"FieldList,omitempty" type:"Repeated"`
}

func (s FieldContentValue) String() string {
	return tea.Prettify(s)
}

func (s FieldContentValue) GoString() string {
	return s.String()
}

func (s *FieldContentValue) SetSortOrder(v int64) *FieldContentValue {
	s.SortOrder = &v
	return s
}

func (s *FieldContentValue) SetFieldList(v []*FieldContentValueFieldList) *FieldContentValue {
	s.FieldList = v
	return s
}

type FieldContentValueFieldList struct {
	// example:
	//
	// ClientIp
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// example:
	//
	// IP address of the client.
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DescriptionCn *string `json:"DescriptionCn,omitempty" xml:"DescriptionCn,omitempty"`
	// example:
	//
	// Client
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// String
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// 1
	SortOrder *int64 `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
}

func (s FieldContentValueFieldList) String() string {
	return tea.Prettify(s)
}

func (s FieldContentValueFieldList) GoString() string {
	return s.String()
}

func (s *FieldContentValueFieldList) SetFieldName(v string) *FieldContentValueFieldList {
	s.FieldName = &v
	return s
}

func (s *FieldContentValueFieldList) SetDescription(v string) *FieldContentValueFieldList {
	s.Description = &v
	return s
}

func (s *FieldContentValueFieldList) SetDescriptionCn(v string) *FieldContentValueFieldList {
	s.DescriptionCn = &v
	return s
}

func (s *FieldContentValueFieldList) SetCategory(v string) *FieldContentValueFieldList {
	s.Category = &v
	return s
}

func (s *FieldContentValueFieldList) SetDataType(v string) *FieldContentValueFieldList {
	s.DataType = &v
	return s
}

func (s *FieldContentValueFieldList) SetSortOrder(v int64) *FieldContentValueFieldList {
	s.SortOrder = &v
	return s
}

func (s *FieldContentValueFieldList) SetIsDefault(v bool) *FieldContentValueFieldList {
	s.IsDefault = &v
	return s
}

type QuotaListItemsValue struct {
	Enable *bool           `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Value  *WafQuotaString `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QuotaListItemsValue) String() string {
	return tea.Prettify(s)
}

func (s QuotaListItemsValue) GoString() string {
	return s.String()
}

func (s *QuotaListItemsValue) SetEnable(v bool) *QuotaListItemsValue {
	s.Enable = &v
	return s
}

func (s *QuotaListItemsValue) SetValue(v *WafQuotaString) *QuotaListItemsValue {
	s.Value = v
	return s
}

type QuotaPageContentTypesValue struct {
	Enable        *bool            `json:"Enable,omitempty" xml:"Enable,omitempty"`
	ContentLength *WafQuotaInteger `json:"ContentLength,omitempty" xml:"ContentLength,omitempty"`
}

func (s QuotaPageContentTypesValue) String() string {
	return tea.Prettify(s)
}

func (s QuotaPageContentTypesValue) GoString() string {
	return s.String()
}

func (s *QuotaPageContentTypesValue) SetEnable(v bool) *QuotaPageContentTypesValue {
	s.Enable = &v
	return s
}

func (s *QuotaPageContentTypesValue) SetContentLength(v *WafQuotaInteger) *QuotaPageContentTypesValue {
	s.ContentLength = v
	return s
}

type ActivateClientCertificateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ActivateClientCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *ActivateClientCertificateRequest) SetId(v string) *ActivateClientCertificateRequest {
	s.Id = &v
	return s
}

func (s *ActivateClientCertificateRequest) SetSiteId(v int64) *ActivateClientCertificateRequest {
	s.SiteId = &v
	return s
}

type ActivateClientCertificateResponseBody struct {
	// example:
	//
	// babaded901474b9693acf530e0fb****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s ActivateClientCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateClientCertificateResponseBody) SetId(v string) *ActivateClientCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *ActivateClientCertificateResponseBody) SetRequestId(v string) *ActivateClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateClientCertificateResponseBody) SetSiteId(v int64) *ActivateClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *ActivateClientCertificateResponseBody) SetSiteName(v string) *ActivateClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

type ActivateClientCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateClientCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *ActivateClientCertificateResponse) SetHeaders(v map[string]*string) *ActivateClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *ActivateClientCertificateResponse) SetStatusCode(v int32) *ActivateClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateClientCertificateResponse) SetBody(v *ActivateClientCertificateResponseBody) *ActivateClientCertificateResponse {
	s.Body = v
	return s
}

type AddUserBusinessFormRequest struct {
	// This parameter is required.
	Company *string `json:"Company,omitempty" xml:"Company,omitempty"`
	// This parameter is required.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// 记录名称
	//
	// This parameter is required.
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// 记录类型
	//
	// This parameter is required.
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AddUserBusinessForm
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 业务场景
	Website *string `json:"Website,omitempty" xml:"Website,omitempty"`
}

func (s AddUserBusinessFormRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserBusinessFormRequest) GoString() string {
	return s.String()
}

func (s *AddUserBusinessFormRequest) SetCompany(v string) *AddUserBusinessFormRequest {
	s.Company = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetEmail(v string) *AddUserBusinessFormRequest {
	s.Email = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetPhoneNumber(v string) *AddUserBusinessFormRequest {
	s.PhoneNumber = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetPosition(v string) *AddUserBusinessFormRequest {
	s.Position = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetRemark(v string) *AddUserBusinessFormRequest {
	s.Remark = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetUserName(v string) *AddUserBusinessFormRequest {
	s.UserName = &v
	return s
}

func (s *AddUserBusinessFormRequest) SetWebsite(v string) *AddUserBusinessFormRequest {
	s.Website = &v
	return s
}

type AddUserBusinessFormResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserBusinessFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserBusinessFormResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserBusinessFormResponseBody) SetRequestId(v string) *AddUserBusinessFormResponseBody {
	s.RequestId = &v
	return s
}

type AddUserBusinessFormResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserBusinessFormResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserBusinessFormResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserBusinessFormResponse) GoString() string {
	return s.String()
}

func (s *AddUserBusinessFormResponse) SetHeaders(v map[string]*string) *AddUserBusinessFormResponse {
	s.Headers = v
	return s
}

func (s *AddUserBusinessFormResponse) SetStatusCode(v int32) *AddUserBusinessFormResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserBusinessFormResponse) SetBody(v *AddUserBusinessFormResponseBody) *AddUserBusinessFormResponse {
	s.Body = v
	return s
}

type AdvancePurgeObjectCacheRequest struct {
	Area    *string                                `json:"Area,omitempty" xml:"Area,omitempty"`
	Content *AdvancePurgeObjectCacheRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Force   *bool                                  `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId         *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	Stations       *string `json:"Stations,omitempty" xml:"Stations,omitempty"`
	TimeRangeBegin *int32  `json:"TimeRangeBegin,omitempty" xml:"TimeRangeBegin,omitempty"`
	TimeRangeEnd   *int32  `json:"TimeRangeEnd,omitempty" xml:"TimeRangeEnd,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AdvancePurgeObjectCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s AdvancePurgeObjectCacheRequest) GoString() string {
	return s.String()
}

func (s *AdvancePurgeObjectCacheRequest) SetArea(v string) *AdvancePurgeObjectCacheRequest {
	s.Area = &v
	return s
}

func (s *AdvancePurgeObjectCacheRequest) SetContent(v *AdvancePurgeObjectCacheRequestContent) *AdvancePurgeObjectCacheRequest {
	s.Content = v
	return s
}

func (s *AdvancePurgeObjectCacheRequest) SetForce(v bool) *AdvancePurgeObjectCacheRequest {
	s.Force = &v
	return s
}

func (s *AdvancePurgeObjectCacheRequest) SetSiteId(v int64) *AdvancePurgeObjectCacheRequest {
	s.SiteId = &v
	return s
}

func (s *AdvancePurgeObjectCacheRequest) SetStations(v string) *AdvancePurgeObjectCacheRequest {
	s.Stations = &v
	return s
}

func (s *AdvancePurgeObjectCacheRequest) SetTimeRangeBegin(v int32) *AdvancePurgeObjectCacheRequest {
	s.TimeRangeBegin = &v
	return s
}

func (s *AdvancePurgeObjectCacheRequest) SetTimeRangeEnd(v int32) *AdvancePurgeObjectCacheRequest {
	s.TimeRangeEnd = &v
	return s
}

func (s *AdvancePurgeObjectCacheRequest) SetType(v string) *AdvancePurgeObjectCacheRequest {
	s.Type = &v
	return s
}

type AdvancePurgeObjectCacheRequestContent struct {
	CacheTags    []*string     `json:"CacheTags,omitempty" xml:"CacheTags,omitempty" type:"Repeated"`
	Directories  []*string     `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	Files        []interface{} `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	Hostnames    []*string     `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	IgnoreParams []*string     `json:"IgnoreParams,omitempty" xml:"IgnoreParams,omitempty" type:"Repeated"`
	PurgeAll     *bool         `json:"PurgeAll,omitempty" xml:"PurgeAll,omitempty"`
}

func (s AdvancePurgeObjectCacheRequestContent) String() string {
	return tea.Prettify(s)
}

func (s AdvancePurgeObjectCacheRequestContent) GoString() string {
	return s.String()
}

func (s *AdvancePurgeObjectCacheRequestContent) SetCacheTags(v []*string) *AdvancePurgeObjectCacheRequestContent {
	s.CacheTags = v
	return s
}

func (s *AdvancePurgeObjectCacheRequestContent) SetDirectories(v []*string) *AdvancePurgeObjectCacheRequestContent {
	s.Directories = v
	return s
}

func (s *AdvancePurgeObjectCacheRequestContent) SetFiles(v []interface{}) *AdvancePurgeObjectCacheRequestContent {
	s.Files = v
	return s
}

func (s *AdvancePurgeObjectCacheRequestContent) SetHostnames(v []*string) *AdvancePurgeObjectCacheRequestContent {
	s.Hostnames = v
	return s
}

func (s *AdvancePurgeObjectCacheRequestContent) SetIgnoreParams(v []*string) *AdvancePurgeObjectCacheRequestContent {
	s.IgnoreParams = v
	return s
}

func (s *AdvancePurgeObjectCacheRequestContent) SetPurgeAll(v bool) *AdvancePurgeObjectCacheRequestContent {
	s.PurgeAll = &v
	return s
}

type AdvancePurgeObjectCacheShrinkRequest struct {
	Area          *string `json:"Area,omitempty" xml:"Area,omitempty"`
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Force         *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId         *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	Stations       *string `json:"Stations,omitempty" xml:"Stations,omitempty"`
	TimeRangeBegin *int32  `json:"TimeRangeBegin,omitempty" xml:"TimeRangeBegin,omitempty"`
	TimeRangeEnd   *int32  `json:"TimeRangeEnd,omitempty" xml:"TimeRangeEnd,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AdvancePurgeObjectCacheShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AdvancePurgeObjectCacheShrinkRequest) GoString() string {
	return s.String()
}

func (s *AdvancePurgeObjectCacheShrinkRequest) SetArea(v string) *AdvancePurgeObjectCacheShrinkRequest {
	s.Area = &v
	return s
}

func (s *AdvancePurgeObjectCacheShrinkRequest) SetContentShrink(v string) *AdvancePurgeObjectCacheShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *AdvancePurgeObjectCacheShrinkRequest) SetForce(v bool) *AdvancePurgeObjectCacheShrinkRequest {
	s.Force = &v
	return s
}

func (s *AdvancePurgeObjectCacheShrinkRequest) SetSiteId(v int64) *AdvancePurgeObjectCacheShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *AdvancePurgeObjectCacheShrinkRequest) SetStations(v string) *AdvancePurgeObjectCacheShrinkRequest {
	s.Stations = &v
	return s
}

func (s *AdvancePurgeObjectCacheShrinkRequest) SetTimeRangeBegin(v int32) *AdvancePurgeObjectCacheShrinkRequest {
	s.TimeRangeBegin = &v
	return s
}

func (s *AdvancePurgeObjectCacheShrinkRequest) SetTimeRangeEnd(v int32) *AdvancePurgeObjectCacheShrinkRequest {
	s.TimeRangeEnd = &v
	return s
}

func (s *AdvancePurgeObjectCacheShrinkRequest) SetType(v string) *AdvancePurgeObjectCacheShrinkRequest {
	s.Type = &v
	return s
}

type AdvancePurgeObjectCacheResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AdvancePurgeObjectCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AdvancePurgeObjectCacheResponseBody) GoString() string {
	return s.String()
}

func (s *AdvancePurgeObjectCacheResponseBody) SetRequestId(v string) *AdvancePurgeObjectCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *AdvancePurgeObjectCacheResponseBody) SetTaskId(v string) *AdvancePurgeObjectCacheResponseBody {
	s.TaskId = &v
	return s
}

type AdvancePurgeObjectCacheResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AdvancePurgeObjectCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AdvancePurgeObjectCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s AdvancePurgeObjectCacheResponse) GoString() string {
	return s.String()
}

func (s *AdvancePurgeObjectCacheResponse) SetHeaders(v map[string]*string) *AdvancePurgeObjectCacheResponse {
	s.Headers = v
	return s
}

func (s *AdvancePurgeObjectCacheResponse) SetStatusCode(v int32) *AdvancePurgeObjectCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *AdvancePurgeObjectCacheResponse) SetBody(v *AdvancePurgeObjectCacheResponseBody) *AdvancePurgeObjectCacheResponse {
	s.Body = v
	return s
}

type BatchCreateRecordsRequest struct {
	// This parameter is required.
	RecordList []*BatchCreateRecordsRequestRecordList `json:"RecordList,omitempty" xml:"RecordList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s BatchCreateRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsRequest) SetRecordList(v []*BatchCreateRecordsRequestRecordList) *BatchCreateRecordsRequest {
	s.RecordList = v
	return s
}

func (s *BatchCreateRecordsRequest) SetSiteId(v int64) *BatchCreateRecordsRequest {
	s.SiteId = &v
	return s
}

type BatchCreateRecordsRequestRecordList struct {
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "value":"2.2.2.2"
	//
	// }
	Data *BatchCreateRecordsRequestRecordListData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// A/AAAA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BatchCreateRecordsRequestRecordList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsRequestRecordList) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsRequestRecordList) SetBizName(v string) *BatchCreateRecordsRequestRecordList {
	s.BizName = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetData(v *BatchCreateRecordsRequestRecordListData) *BatchCreateRecordsRequestRecordList {
	s.Data = v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetProxied(v bool) *BatchCreateRecordsRequestRecordList {
	s.Proxied = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetRecordName(v string) *BatchCreateRecordsRequestRecordList {
	s.RecordName = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetSourceType(v string) *BatchCreateRecordsRequestRecordList {
	s.SourceType = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetTtl(v int32) *BatchCreateRecordsRequestRecordList {
	s.Ttl = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordList) SetType(v string) *BatchCreateRecordsRequestRecordList {
	s.Type = &v
	return s
}

type BatchCreateRecordsRequestRecordListData struct {
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// example:
	//
	// 0
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s BatchCreateRecordsRequestRecordListData) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsRequestRecordListData) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsRequestRecordListData) SetAlgorithm(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Algorithm = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetCertificate(v string) *BatchCreateRecordsRequestRecordListData {
	s.Certificate = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetFingerprint(v string) *BatchCreateRecordsRequestRecordListData {
	s.Fingerprint = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetFlag(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Flag = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetKeyTag(v int32) *BatchCreateRecordsRequestRecordListData {
	s.KeyTag = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetMatchingType(v int32) *BatchCreateRecordsRequestRecordListData {
	s.MatchingType = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetPort(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Port = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetPriority(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Priority = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetSelector(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Selector = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetTag(v string) *BatchCreateRecordsRequestRecordListData {
	s.Tag = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetType(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Type = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetUsage(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Usage = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetValue(v string) *BatchCreateRecordsRequestRecordListData {
	s.Value = &v
	return s
}

func (s *BatchCreateRecordsRequestRecordListData) SetWeight(v int32) *BatchCreateRecordsRequestRecordListData {
	s.Weight = &v
	return s
}

type BatchCreateRecordsShrinkRequest struct {
	// This parameter is required.
	RecordListShrink *string `json:"RecordList,omitempty" xml:"RecordList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s BatchCreateRecordsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsShrinkRequest) SetRecordListShrink(v string) *BatchCreateRecordsShrinkRequest {
	s.RecordListShrink = &v
	return s
}

func (s *BatchCreateRecordsShrinkRequest) SetSiteId(v int64) *BatchCreateRecordsShrinkRequest {
	s.SiteId = &v
	return s
}

type BatchCreateRecordsResponseBody struct {
	RecordResultList *BatchCreateRecordsResponseBodyRecordResultList `json:"RecordResultList,omitempty" xml:"RecordResultList,omitempty" type:"Struct"`
	// example:
	//
	// 2430E05E-1340-5773-B5E1-B743929F46F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchCreateRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBody) SetRecordResultList(v *BatchCreateRecordsResponseBodyRecordResultList) *BatchCreateRecordsResponseBody {
	s.RecordResultList = v
	return s
}

func (s *BatchCreateRecordsResponseBody) SetRequestId(v string) *BatchCreateRecordsResponseBody {
	s.RequestId = &v
	return s
}

type BatchCreateRecordsResponseBodyRecordResultList struct {
	Failed  []*BatchCreateRecordsResponseBodyRecordResultListFailed  `json:"Failed,omitempty" xml:"Failed,omitempty" type:"Repeated"`
	Success []*BatchCreateRecordsResponseBodyRecordResultListSuccess `json:"Success,omitempty" xml:"Success,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultList) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) SetFailed(v []*BatchCreateRecordsResponseBodyRecordResultListFailed) *BatchCreateRecordsResponseBodyRecordResultList {
	s.Failed = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) SetSuccess(v []*BatchCreateRecordsResponseBodyRecordResultListSuccess) *BatchCreateRecordsResponseBodyRecordResultList {
	s.Success = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultList) SetTotalCount(v int32) *BatchCreateRecordsResponseBodyRecordResultList {
	s.TotalCount = &v
	return s
}

type BatchCreateRecordsResponseBodyRecordResultListFailed struct {
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// {"value":"2.2.2.2"}
	Data        *BatchCreateRecordsResponseBodyRecordResultListFailedData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Description *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// example:
	//
	// A/AAAA
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailed) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailed) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetBizName(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.BizName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetData(v *BatchCreateRecordsResponseBodyRecordResultListFailedData) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Data = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetDescription(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Description = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetProxied(v bool) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Proxied = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetRecordId(v int64) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.RecordId = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetRecordName(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.RecordName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetRecordType(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.RecordType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetSourceType(v string) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.SourceType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailed) SetTtl(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailed {
	s.Ttl = &v
	return s
}

type BatchCreateRecordsResponseBodyRecordResultListFailedData struct {
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// example:
	//
	// RSA
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailedData) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListFailedData) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetAlgorithm(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Algorithm = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetCertificate(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Certificate = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetFingerprint(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Fingerprint = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetFlag(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Flag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetKeyTag(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.KeyTag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetMatchingType(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.MatchingType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetPort(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Port = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetPriority(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Priority = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetSelector(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Selector = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetTag(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Tag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetType(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Type = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetUsage(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Usage = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetValue(v string) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Value = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListFailedData) SetWeight(v int32) *BatchCreateRecordsResponseBodyRecordResultListFailedData {
	s.Weight = &v
	return s
}

type BatchCreateRecordsResponseBodyRecordResultListSuccess struct {
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// {"value":"1.1.1.1"}
	Data *BatchCreateRecordsResponseBodyRecordResultListSuccessData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// example:
	//
	// A/AAAA
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccess) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccess) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetBizName(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.BizName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetData(v *BatchCreateRecordsResponseBodyRecordResultListSuccessData) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Data = v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetDescription(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Description = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetProxied(v bool) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Proxied = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetRecordId(v int64) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.RecordId = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetRecordName(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.RecordName = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetRecordType(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.RecordType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetSourceType(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.SourceType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccess) SetTtl(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccess {
	s.Ttl = &v
	return s
}

type BatchCreateRecordsResponseBodyRecordResultListSuccessData struct {
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// example:
	//
	// 0
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccessData) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsResponseBodyRecordResultListSuccessData) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetAlgorithm(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Algorithm = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetCertificate(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Certificate = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetFingerprint(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Fingerprint = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetFlag(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Flag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetKeyTag(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.KeyTag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetMatchingType(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.MatchingType = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetPort(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Port = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetPriority(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Priority = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetSelector(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Selector = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetTag(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Tag = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetType(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Type = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetUsage(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Usage = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetValue(v string) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Value = &v
	return s
}

func (s *BatchCreateRecordsResponseBodyRecordResultListSuccessData) SetWeight(v int32) *BatchCreateRecordsResponseBodyRecordResultListSuccessData {
	s.Weight = &v
	return s
}

type BatchCreateRecordsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRecordsResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponse) SetHeaders(v map[string]*string) *BatchCreateRecordsResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateRecordsResponse) SetStatusCode(v int32) *BatchCreateRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateRecordsResponse) SetBody(v *BatchCreateRecordsResponseBody) *BatchCreateRecordsResponse {
	s.Body = v
	return s
}

type BatchCreateWafRulesRequest struct {
	Configs []*WafRuleConfig `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// example:
	//
	// http_custom
	Phase  *string             `json:"Phase,omitempty" xml:"Phase,omitempty"`
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s BatchCreateWafRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateWafRulesRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateWafRulesRequest) SetConfigs(v []*WafRuleConfig) *BatchCreateWafRulesRequest {
	s.Configs = v
	return s
}

func (s *BatchCreateWafRulesRequest) SetPhase(v string) *BatchCreateWafRulesRequest {
	s.Phase = &v
	return s
}

func (s *BatchCreateWafRulesRequest) SetShared(v *WafBatchRuleShared) *BatchCreateWafRulesRequest {
	s.Shared = v
	return s
}

func (s *BatchCreateWafRulesRequest) SetSiteId(v int64) *BatchCreateWafRulesRequest {
	s.SiteId = &v
	return s
}

func (s *BatchCreateWafRulesRequest) SetSiteVersion(v int32) *BatchCreateWafRulesRequest {
	s.SiteVersion = &v
	return s
}

type BatchCreateWafRulesShrinkRequest struct {
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// example:
	//
	// http_custom
	Phase        *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	SharedShrink *string `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s BatchCreateWafRulesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateWafRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateWafRulesShrinkRequest) SetConfigsShrink(v string) *BatchCreateWafRulesShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetPhase(v string) *BatchCreateWafRulesShrinkRequest {
	s.Phase = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetSharedShrink(v string) *BatchCreateWafRulesShrinkRequest {
	s.SharedShrink = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetSiteId(v int64) *BatchCreateWafRulesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetSiteVersion(v int32) *BatchCreateWafRulesShrinkRequest {
	s.SiteVersion = &v
	return s
}

type BatchCreateWafRulesResponseBody struct {
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RulesetId *int64  `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
}

func (s BatchCreateWafRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateWafRulesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateWafRulesResponseBody) SetIds(v []*int64) *BatchCreateWafRulesResponseBody {
	s.Ids = v
	return s
}

func (s *BatchCreateWafRulesResponseBody) SetRequestId(v string) *BatchCreateWafRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateWafRulesResponseBody) SetRulesetId(v int64) *BatchCreateWafRulesResponseBody {
	s.RulesetId = &v
	return s
}

type BatchCreateWafRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateWafRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateWafRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateWafRulesResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateWafRulesResponse) SetHeaders(v map[string]*string) *BatchCreateWafRulesResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateWafRulesResponse) SetStatusCode(v int32) *BatchCreateWafRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateWafRulesResponse) SetBody(v *BatchCreateWafRulesResponseBody) *BatchCreateWafRulesResponse {
	s.Body = v
	return s
}

type BatchDeleteKvRequest struct {
	// This parameter is required.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchDeleteKvRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteKvRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvRequest) SetKeys(v []*string) *BatchDeleteKvRequest {
	s.Keys = v
	return s
}

func (s *BatchDeleteKvRequest) SetNamespace(v string) *BatchDeleteKvRequest {
	s.Namespace = &v
	return s
}

type BatchDeleteKvShrinkRequest struct {
	// This parameter is required.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchDeleteKvShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteKvShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvShrinkRequest) SetKeysShrink(v string) *BatchDeleteKvShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *BatchDeleteKvShrinkRequest) SetNamespace(v string) *BatchDeleteKvShrinkRequest {
	s.Namespace = &v
	return s
}

type BatchDeleteKvResponseBody struct {
	FailKeys []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchDeleteKvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteKvResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvResponseBody) SetFailKeys(v []*string) *BatchDeleteKvResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchDeleteKvResponseBody) SetRequestId(v string) *BatchDeleteKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteKvResponseBody) SetSuccessKeys(v []*string) *BatchDeleteKvResponseBody {
	s.SuccessKeys = v
	return s
}

type BatchDeleteKvResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteKvResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteKvResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvResponse) SetHeaders(v map[string]*string) *BatchDeleteKvResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteKvResponse) SetStatusCode(v int32) *BatchDeleteKvResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteKvResponse) SetBody(v *BatchDeleteKvResponseBody) *BatchDeleteKvResponse {
	s.Body = v
	return s
}

type BatchDeleteKvWithHighCapacityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchDeleteKvWithHighCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvWithHighCapacityRequest) SetNamespace(v string) *BatchDeleteKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityRequest) SetUrl(v string) *BatchDeleteKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

type BatchDeleteKvWithHighCapacityAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchDeleteKvWithHighCapacityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *BatchDeleteKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *BatchDeleteKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

type BatchDeleteKvWithHighCapacityResponseBody struct {
	FailKeys []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchDeleteKvWithHighCapacityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteKvWithHighCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvWithHighCapacityResponseBody) SetFailKeys(v []*string) *BatchDeleteKvWithHighCapacityResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchDeleteKvWithHighCapacityResponseBody) SetRequestId(v string) *BatchDeleteKvWithHighCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityResponseBody) SetSuccessKeys(v []*string) *BatchDeleteKvWithHighCapacityResponseBody {
	s.SuccessKeys = v
	return s
}

type BatchDeleteKvWithHighCapacityResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteKvWithHighCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteKvWithHighCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteKvWithHighCapacityResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvWithHighCapacityResponse) SetHeaders(v map[string]*string) *BatchDeleteKvWithHighCapacityResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteKvWithHighCapacityResponse) SetStatusCode(v int32) *BatchDeleteKvWithHighCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityResponse) SetBody(v *BatchDeleteKvWithHighCapacityResponseBody) *BatchDeleteKvWithHighCapacityResponse {
	s.Body = v
	return s
}

type BatchGetExpressionFieldsRequest struct {
	// example:
	//
	// http_bot
	Expressions []*BatchGetExpressionFieldsRequestExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s BatchGetExpressionFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetExpressionFieldsRequest) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsRequest) SetExpressions(v []*BatchGetExpressionFieldsRequestExpressions) *BatchGetExpressionFieldsRequest {
	s.Expressions = v
	return s
}

func (s *BatchGetExpressionFieldsRequest) SetPhase(v string) *BatchGetExpressionFieldsRequest {
	s.Phase = &v
	return s
}

func (s *BatchGetExpressionFieldsRequest) SetSiteId(v int64) *BatchGetExpressionFieldsRequest {
	s.SiteId = &v
	return s
}

type BatchGetExpressionFieldsRequestExpressions struct {
	// example:
	//
	// ip.src eq 1.1.1.1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s BatchGetExpressionFieldsRequestExpressions) String() string {
	return tea.Prettify(s)
}

func (s BatchGetExpressionFieldsRequestExpressions) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsRequestExpressions) SetExpression(v string) *BatchGetExpressionFieldsRequestExpressions {
	s.Expression = &v
	return s
}

func (s *BatchGetExpressionFieldsRequestExpressions) SetId(v int64) *BatchGetExpressionFieldsRequestExpressions {
	s.Id = &v
	return s
}

type BatchGetExpressionFieldsShrinkRequest struct {
	// example:
	//
	// http_bot
	ExpressionsShrink *string `json:"Expressions,omitempty" xml:"Expressions,omitempty"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s BatchGetExpressionFieldsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetExpressionFieldsShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetExpressionsShrink(v string) *BatchGetExpressionFieldsShrinkRequest {
	s.ExpressionsShrink = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetPhase(v string) *BatchGetExpressionFieldsShrinkRequest {
	s.Phase = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetSiteId(v int64) *BatchGetExpressionFieldsShrinkRequest {
	s.SiteId = &v
	return s
}

type BatchGetExpressionFieldsResponseBody struct {
	Fields []*BatchGetExpressionFieldsResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetExpressionFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetExpressionFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsResponseBody) SetFields(v []*BatchGetExpressionFieldsResponseBodyFields) *BatchGetExpressionFieldsResponseBody {
	s.Fields = v
	return s
}

func (s *BatchGetExpressionFieldsResponseBody) SetRequestId(v string) *BatchGetExpressionFieldsResponseBody {
	s.RequestId = &v
	return s
}

type BatchGetExpressionFieldsResponseBodyFields struct {
	Fields []*string `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s BatchGetExpressionFieldsResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s BatchGetExpressionFieldsResponseBodyFields) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsResponseBodyFields) SetFields(v []*string) *BatchGetExpressionFieldsResponseBodyFields {
	s.Fields = v
	return s
}

func (s *BatchGetExpressionFieldsResponseBodyFields) SetId(v string) *BatchGetExpressionFieldsResponseBodyFields {
	s.Id = &v
	return s
}

type BatchGetExpressionFieldsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetExpressionFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetExpressionFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetExpressionFieldsResponse) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsResponse) SetHeaders(v map[string]*string) *BatchGetExpressionFieldsResponse {
	s.Headers = v
	return s
}

func (s *BatchGetExpressionFieldsResponse) SetStatusCode(v int32) *BatchGetExpressionFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetExpressionFieldsResponse) SetBody(v *BatchGetExpressionFieldsResponseBody) *BatchGetExpressionFieldsResponse {
	s.Body = v
	return s
}

type BatchPutKvRequest struct {
	// This parameter is required.
	KvList []*BatchPutKvRequestKvList `json:"KvList,omitempty" xml:"KvList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchPutKvRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchPutKvRequest) GoString() string {
	return s.String()
}

func (s *BatchPutKvRequest) SetKvList(v []*BatchPutKvRequestKvList) *BatchPutKvRequest {
	s.KvList = v
	return s
}

func (s *BatchPutKvRequest) SetNamespace(v string) *BatchPutKvRequest {
	s.Namespace = &v
	return s
}

type BatchPutKvRequestKvList struct {
	// example:
	//
	// 1690081381
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// example:
	//
	// 3600
	ExpirationTtl *int64 `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s BatchPutKvRequestKvList) String() string {
	return tea.Prettify(s)
}

func (s BatchPutKvRequestKvList) GoString() string {
	return s.String()
}

func (s *BatchPutKvRequestKvList) SetExpiration(v int64) *BatchPutKvRequestKvList {
	s.Expiration = &v
	return s
}

func (s *BatchPutKvRequestKvList) SetExpirationTtl(v int64) *BatchPutKvRequestKvList {
	s.ExpirationTtl = &v
	return s
}

func (s *BatchPutKvRequestKvList) SetKey(v string) *BatchPutKvRequestKvList {
	s.Key = &v
	return s
}

func (s *BatchPutKvRequestKvList) SetValue(v string) *BatchPutKvRequestKvList {
	s.Value = &v
	return s
}

type BatchPutKvShrinkRequest struct {
	// This parameter is required.
	KvListShrink *string `json:"KvList,omitempty" xml:"KvList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchPutKvShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchPutKvShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchPutKvShrinkRequest) SetKvListShrink(v string) *BatchPutKvShrinkRequest {
	s.KvListShrink = &v
	return s
}

func (s *BatchPutKvShrinkRequest) SetNamespace(v string) *BatchPutKvShrinkRequest {
	s.Namespace = &v
	return s
}

type BatchPutKvResponseBody struct {
	FailKeys []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchPutKvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchPutKvResponseBody) GoString() string {
	return s.String()
}

func (s *BatchPutKvResponseBody) SetFailKeys(v []*string) *BatchPutKvResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchPutKvResponseBody) SetRequestId(v string) *BatchPutKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchPutKvResponseBody) SetSuccessKeys(v []*string) *BatchPutKvResponseBody {
	s.SuccessKeys = v
	return s
}

type BatchPutKvResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchPutKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchPutKvResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchPutKvResponse) GoString() string {
	return s.String()
}

func (s *BatchPutKvResponse) SetHeaders(v map[string]*string) *BatchPutKvResponse {
	s.Headers = v
	return s
}

func (s *BatchPutKvResponse) SetStatusCode(v int32) *BatchPutKvResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchPutKvResponse) SetBody(v *BatchPutKvResponseBody) *BatchPutKvResponse {
	s.Body = v
	return s
}

type BatchPutKvWithHighCapacityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchPutKvWithHighCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchPutKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *BatchPutKvWithHighCapacityRequest) SetNamespace(v string) *BatchPutKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutKvWithHighCapacityRequest) SetUrl(v string) *BatchPutKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

type BatchPutKvWithHighCapacityAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchPutKvWithHighCapacityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchPutKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *BatchPutKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *BatchPutKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

type BatchPutKvWithHighCapacityResponseBody struct {
	FailKeys []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchPutKvWithHighCapacityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchPutKvWithHighCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *BatchPutKvWithHighCapacityResponseBody) SetFailKeys(v []*string) *BatchPutKvWithHighCapacityResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchPutKvWithHighCapacityResponseBody) SetRequestId(v string) *BatchPutKvWithHighCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchPutKvWithHighCapacityResponseBody) SetSuccessKeys(v []*string) *BatchPutKvWithHighCapacityResponseBody {
	s.SuccessKeys = v
	return s
}

type BatchPutKvWithHighCapacityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchPutKvWithHighCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchPutKvWithHighCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchPutKvWithHighCapacityResponse) GoString() string {
	return s.String()
}

func (s *BatchPutKvWithHighCapacityResponse) SetHeaders(v map[string]*string) *BatchPutKvWithHighCapacityResponse {
	s.Headers = v
	return s
}

func (s *BatchPutKvWithHighCapacityResponse) SetStatusCode(v int32) *BatchPutKvWithHighCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchPutKvWithHighCapacityResponse) SetBody(v *BatchPutKvWithHighCapacityResponseBody) *BatchPutKvWithHighCapacityResponse {
	s.Body = v
	return s
}

type BatchUpdateWafRulesRequest struct {
	Configs []*WafRuleConfig `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 10000001
	RulesetId *int64              `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	Shared    *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s BatchUpdateWafRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateWafRulesRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateWafRulesRequest) SetConfigs(v []*WafRuleConfig) *BatchUpdateWafRulesRequest {
	s.Configs = v
	return s
}

func (s *BatchUpdateWafRulesRequest) SetPhase(v string) *BatchUpdateWafRulesRequest {
	s.Phase = &v
	return s
}

func (s *BatchUpdateWafRulesRequest) SetRulesetId(v int64) *BatchUpdateWafRulesRequest {
	s.RulesetId = &v
	return s
}

func (s *BatchUpdateWafRulesRequest) SetShared(v *WafBatchRuleShared) *BatchUpdateWafRulesRequest {
	s.Shared = v
	return s
}

func (s *BatchUpdateWafRulesRequest) SetSiteId(v int64) *BatchUpdateWafRulesRequest {
	s.SiteId = &v
	return s
}

func (s *BatchUpdateWafRulesRequest) SetSiteVersion(v int32) *BatchUpdateWafRulesRequest {
	s.SiteVersion = &v
	return s
}

type BatchUpdateWafRulesShrinkRequest struct {
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 10000001
	RulesetId    *int64  `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	SharedShrink *string `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s BatchUpdateWafRulesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateWafRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateWafRulesShrinkRequest) SetConfigsShrink(v string) *BatchUpdateWafRulesShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetPhase(v string) *BatchUpdateWafRulesShrinkRequest {
	s.Phase = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetRulesetId(v int64) *BatchUpdateWafRulesShrinkRequest {
	s.RulesetId = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetSharedShrink(v string) *BatchUpdateWafRulesShrinkRequest {
	s.SharedShrink = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetSiteId(v int64) *BatchUpdateWafRulesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetSiteVersion(v int32) *BatchUpdateWafRulesShrinkRequest {
	s.SiteVersion = &v
	return s
}

type BatchUpdateWafRulesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchUpdateWafRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateWafRulesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateWafRulesResponseBody) SetRequestId(v string) *BatchUpdateWafRulesResponseBody {
	s.RequestId = &v
	return s
}

type BatchUpdateWafRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateWafRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateWafRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateWafRulesResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateWafRulesResponse) SetHeaders(v map[string]*string) *BatchUpdateWafRulesResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateWafRulesResponse) SetStatusCode(v int32) *BatchUpdateWafRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateWafRulesResponse) SetBody(v *BatchUpdateWafRulesResponseBody) *BatchUpdateWafRulesResponse {
	s.Body = v
	return s
}

type BlockObjectRequest struct {
	// This parameter is required.
	Content   []*string `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	Extension *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Maxage    *int32    `json:"Maxage,omitempty" xml:"Maxage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BlockObject
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BlockObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s BlockObjectRequest) GoString() string {
	return s.String()
}

func (s *BlockObjectRequest) SetContent(v []*string) *BlockObjectRequest {
	s.Content = v
	return s
}

func (s *BlockObjectRequest) SetExtension(v string) *BlockObjectRequest {
	s.Extension = &v
	return s
}

func (s *BlockObjectRequest) SetMaxage(v int32) *BlockObjectRequest {
	s.Maxage = &v
	return s
}

func (s *BlockObjectRequest) SetSiteId(v int64) *BlockObjectRequest {
	s.SiteId = &v
	return s
}

func (s *BlockObjectRequest) SetType(v string) *BlockObjectRequest {
	s.Type = &v
	return s
}

type BlockObjectShrinkRequest struct {
	// This parameter is required.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Extension     *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Maxage        *int32  `json:"Maxage,omitempty" xml:"Maxage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BlockObject
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BlockObjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BlockObjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *BlockObjectShrinkRequest) SetContentShrink(v string) *BlockObjectShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *BlockObjectShrinkRequest) SetExtension(v string) *BlockObjectShrinkRequest {
	s.Extension = &v
	return s
}

func (s *BlockObjectShrinkRequest) SetMaxage(v int32) *BlockObjectShrinkRequest {
	s.Maxage = &v
	return s
}

func (s *BlockObjectShrinkRequest) SetSiteId(v int64) *BlockObjectShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BlockObjectShrinkRequest) SetType(v string) *BlockObjectShrinkRequest {
	s.Type = &v
	return s
}

type BlockObjectResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BlockObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BlockObjectResponseBody) GoString() string {
	return s.String()
}

func (s *BlockObjectResponseBody) SetRequestId(v string) *BlockObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *BlockObjectResponseBody) SetTaskId(v string) *BlockObjectResponseBody {
	s.TaskId = &v
	return s
}

type BlockObjectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BlockObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BlockObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s BlockObjectResponse) GoString() string {
	return s.String()
}

func (s *BlockObjectResponse) SetHeaders(v map[string]*string) *BlockObjectResponse {
	s.Headers = v
	return s
}

func (s *BlockObjectResponse) SetStatusCode(v int32) *BlockObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *BlockObjectResponse) SetBody(v *BlockObjectResponseBody) *BlockObjectResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rg-axxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetOwnerId(v int64) *ChangeResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetSecurityToken(v string) *ChangeResourceGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetSiteId(v int64) *ChangeResourceGroupRequest {
	s.SiteId = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F8AA0364-0FDB-4AD5-AC74-D69FAB8924ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CheckSiteNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s CheckSiteNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSiteNameRequest) GoString() string {
	return s.String()
}

func (s *CheckSiteNameRequest) SetSiteName(v string) *CheckSiteNameRequest {
	s.SiteName = &v
	return s
}

type CheckSiteNameResponseBody struct {
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	IsSubSite *bool   `json:"IsSubSite,omitempty" xml:"IsSubSite,omitempty"`
	Messeage  *string `json:"Messeage,omitempty" xml:"Messeage,omitempty"`
	// example:
	//
	// true
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckSiteNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSiteNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSiteNameResponseBody) SetDescription(v string) *CheckSiteNameResponseBody {
	s.Description = &v
	return s
}

func (s *CheckSiteNameResponseBody) SetIsSubSite(v bool) *CheckSiteNameResponseBody {
	s.IsSubSite = &v
	return s
}

func (s *CheckSiteNameResponseBody) SetMesseage(v string) *CheckSiteNameResponseBody {
	s.Messeage = &v
	return s
}

func (s *CheckSiteNameResponseBody) SetPassed(v bool) *CheckSiteNameResponseBody {
	s.Passed = &v
	return s
}

func (s *CheckSiteNameResponseBody) SetRequestId(v string) *CheckSiteNameResponseBody {
	s.RequestId = &v
	return s
}

type CheckSiteNameResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSiteNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSiteNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSiteNameResponse) GoString() string {
	return s.String()
}

func (s *CheckSiteNameResponse) SetHeaders(v map[string]*string) *CheckSiteNameResponse {
	s.Headers = v
	return s
}

func (s *CheckSiteNameResponse) SetStatusCode(v int32) *CheckSiteNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSiteNameResponse) SetBody(v *CheckSiteNameResponseBody) *CheckSiteNameResponse {
	s.Body = v
	return s
}

type CheckSiteProjectNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// user_log
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 12312312213212
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CheckSiteProjectNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSiteProjectNameRequest) GoString() string {
	return s.String()
}

func (s *CheckSiteProjectNameRequest) SetProjectName(v string) *CheckSiteProjectNameRequest {
	s.ProjectName = &v
	return s
}

func (s *CheckSiteProjectNameRequest) SetSiteId(v int64) *CheckSiteProjectNameRequest {
	s.SiteId = &v
	return s
}

type CheckSiteProjectNameResponseBody struct {
	// example:
	//
	// true
	Check *bool `json:"Check,omitempty" xml:"Check,omitempty"`
	// example:
	//
	// project name pass the check
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dcdn_waf_userAccount_log
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckSiteProjectNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSiteProjectNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSiteProjectNameResponseBody) SetCheck(v bool) *CheckSiteProjectNameResponseBody {
	s.Check = &v
	return s
}

func (s *CheckSiteProjectNameResponseBody) SetDescription(v string) *CheckSiteProjectNameResponseBody {
	s.Description = &v
	return s
}

func (s *CheckSiteProjectNameResponseBody) SetProjectName(v string) *CheckSiteProjectNameResponseBody {
	s.ProjectName = &v
	return s
}

func (s *CheckSiteProjectNameResponseBody) SetRequestId(v string) *CheckSiteProjectNameResponseBody {
	s.RequestId = &v
	return s
}

type CheckSiteProjectNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSiteProjectNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSiteProjectNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSiteProjectNameResponse) GoString() string {
	return s.String()
}

func (s *CheckSiteProjectNameResponse) SetHeaders(v map[string]*string) *CheckSiteProjectNameResponse {
	s.Headers = v
	return s
}

func (s *CheckSiteProjectNameResponse) SetStatusCode(v int32) *CheckSiteProjectNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSiteProjectNameResponse) SetBody(v *CheckSiteProjectNameResponseBody) *CheckSiteProjectNameResponse {
	s.Body = v
	return s
}

type CheckUserProjectNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ali-dcdn-log-56
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CheckUserProjectNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUserProjectNameRequest) GoString() string {
	return s.String()
}

func (s *CheckUserProjectNameRequest) SetProjectName(v string) *CheckUserProjectNameRequest {
	s.ProjectName = &v
	return s
}

type CheckUserProjectNameResponseBody struct {
	// example:
	//
	// true
	Check *bool `json:"Check,omitempty" xml:"Check,omitempty"`
	// example:
	//
	// project name pass the check
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ali-dcdn-log-56
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckUserProjectNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUserProjectNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserProjectNameResponseBody) SetCheck(v bool) *CheckUserProjectNameResponseBody {
	s.Check = &v
	return s
}

func (s *CheckUserProjectNameResponseBody) SetDescription(v string) *CheckUserProjectNameResponseBody {
	s.Description = &v
	return s
}

func (s *CheckUserProjectNameResponseBody) SetProjectName(v string) *CheckUserProjectNameResponseBody {
	s.ProjectName = &v
	return s
}

func (s *CheckUserProjectNameResponseBody) SetRequestId(v string) *CheckUserProjectNameResponseBody {
	s.RequestId = &v
	return s
}

type CheckUserProjectNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserProjectNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUserProjectNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserProjectNameResponse) GoString() string {
	return s.String()
}

func (s *CheckUserProjectNameResponse) SetHeaders(v map[string]*string) *CheckUserProjectNameResponse {
	s.Headers = v
	return s
}

func (s *CheckUserProjectNameResponse) SetStatusCode(v int32) *CheckUserProjectNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserProjectNameResponse) SetBody(v *CheckUserProjectNameResponseBody) *CheckUserProjectNameResponse {
	s.Body = v
	return s
}

type CreateCustomScenePolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-07T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 7096621098****
	Objects *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-07T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s CreateCustomScenePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomScenePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomScenePolicyRequest) SetEndTime(v string) *CreateCustomScenePolicyRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCustomScenePolicyRequest) SetName(v string) *CreateCustomScenePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomScenePolicyRequest) SetObjects(v string) *CreateCustomScenePolicyRequest {
	s.Objects = &v
	return s
}

func (s *CreateCustomScenePolicyRequest) SetStartTime(v string) *CreateCustomScenePolicyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCustomScenePolicyRequest) SetTemplate(v string) *CreateCustomScenePolicyRequest {
	s.Template = &v
	return s
}

type CreateCustomScenePolicyResponseBody struct {
	// example:
	//
	// 2021-11-07T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// test
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Objects []*string `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-11-07T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s CreateCustomScenePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomScenePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomScenePolicyResponseBody) SetEndTime(v string) *CreateCustomScenePolicyResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetName(v string) *CreateCustomScenePolicyResponseBody {
	s.Name = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetObjects(v []*string) *CreateCustomScenePolicyResponseBody {
	s.Objects = v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetPolicyId(v int64) *CreateCustomScenePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetRequestId(v string) *CreateCustomScenePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetStartTime(v string) *CreateCustomScenePolicyResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetTemplate(v string) *CreateCustomScenePolicyResponseBody {
	s.Template = &v
	return s
}

type CreateCustomScenePolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomScenePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomScenePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomScenePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomScenePolicyResponse) SetHeaders(v map[string]*string) *CreateCustomScenePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomScenePolicyResponse) SetStatusCode(v int32) *CreateCustomScenePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomScenePolicyResponse) SetBody(v *CreateCustomScenePolicyResponseBody) *CreateCustomScenePolicyResponse {
	s.Body = v
	return s
}

type CreateKvNamespaceRequest struct {
	// example:
	//
	// this is a test namespace.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s CreateKvNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKvNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateKvNamespaceRequest) SetDescription(v string) *CreateKvNamespaceRequest {
	s.Description = &v
	return s
}

func (s *CreateKvNamespaceRequest) SetNamespace(v string) *CreateKvNamespaceRequest {
	s.Namespace = &v
	return s
}

type CreateKvNamespaceResponseBody struct {
	// example:
	//
	// this is a test namespace.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 657717877171818496
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateKvNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKvNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKvNamespaceResponseBody) SetDescription(v string) *CreateKvNamespaceResponseBody {
	s.Description = &v
	return s
}

func (s *CreateKvNamespaceResponseBody) SetNamespace(v string) *CreateKvNamespaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *CreateKvNamespaceResponseBody) SetNamespaceId(v string) *CreateKvNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *CreateKvNamespaceResponseBody) SetRequestId(v string) *CreateKvNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKvNamespaceResponseBody) SetStatus(v string) *CreateKvNamespaceResponseBody {
	s.Status = &v
	return s
}

type CreateKvNamespaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKvNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKvNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKvNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateKvNamespaceResponse) SetHeaders(v map[string]*string) *CreateKvNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateKvNamespaceResponse) SetStatusCode(v int32) *CreateKvNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKvNamespaceResponse) SetBody(v *CreateKvNamespaceResponseBody) *CreateKvNamespaceResponse {
	s.Body = v
	return s
}

type CreateListRequest struct {
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// a custom list
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// ip
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateListRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateListRequest) GoString() string {
	return s.String()
}

func (s *CreateListRequest) SetDescription(v string) *CreateListRequest {
	s.Description = &v
	return s
}

func (s *CreateListRequest) SetItems(v []*string) *CreateListRequest {
	s.Items = v
	return s
}

func (s *CreateListRequest) SetKind(v string) *CreateListRequest {
	s.Kind = &v
	return s
}

func (s *CreateListRequest) SetName(v string) *CreateListRequest {
	s.Name = &v
	return s
}

type CreateListShrinkRequest struct {
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// a custom list
	ItemsShrink *string `json:"Items,omitempty" xml:"Items,omitempty"`
	// example:
	//
	// ip
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateListShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateListShrinkRequest) SetDescription(v string) *CreateListShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateListShrinkRequest) SetItemsShrink(v string) *CreateListShrinkRequest {
	s.ItemsShrink = &v
	return s
}

func (s *CreateListShrinkRequest) SetKind(v string) *CreateListShrinkRequest {
	s.Kind = &v
	return s
}

func (s *CreateListShrinkRequest) SetName(v string) *CreateListShrinkRequest {
	s.Name = &v
	return s
}

type CreateListResponseBody struct {
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateListResponseBody) SetId(v int64) *CreateListResponseBody {
	s.Id = &v
	return s
}

func (s *CreateListResponseBody) SetRequestId(v string) *CreateListResponseBody {
	s.RequestId = &v
	return s
}

type CreateListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateListResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateListResponse) GoString() string {
	return s.String()
}

func (s *CreateListResponse) SetHeaders(v map[string]*string) *CreateListResponse {
	s.Headers = v
	return s
}

func (s *CreateListResponse) SetStatusCode(v int32) *CreateListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateListResponse) SetBody(v *CreateListResponseBody) *CreateListResponse {
	s.Body = v
	return s
}

type CreatePageRequest struct {
	// example:
	//
	// PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text/html
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePageRequest) GoString() string {
	return s.String()
}

func (s *CreatePageRequest) SetContent(v string) *CreatePageRequest {
	s.Content = &v
	return s
}

func (s *CreatePageRequest) SetContentType(v string) *CreatePageRequest {
	s.ContentType = &v
	return s
}

func (s *CreatePageRequest) SetDescription(v string) *CreatePageRequest {
	s.Description = &v
	return s
}

func (s *CreatePageRequest) SetName(v string) *CreatePageRequest {
	s.Name = &v
	return s
}

type CreatePageResponseBody struct {
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePageResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePageResponseBody) SetId(v int64) *CreatePageResponseBody {
	s.Id = &v
	return s
}

func (s *CreatePageResponseBody) SetRequestId(v string) *CreatePageResponseBody {
	s.RequestId = &v
	return s
}

type CreatePageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePageResponse) GoString() string {
	return s.String()
}

func (s *CreatePageResponse) SetHeaders(v map[string]*string) *CreatePageResponse {
	s.Headers = v
	return s
}

func (s *CreatePageResponse) SetStatusCode(v int32) *CreatePageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePageResponse) SetBody(v *CreatePageResponseBody) *CreatePageResponse {
	s.Body = v
	return s
}

type CreateRecordRequest struct {
	AuthConf *CreateRecordRequestAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// 业务场景
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "value":"2.2.2.2"
	//
	// }
	Data       *CreateRecordRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HostPolicy *string                  `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// 是否代理加速
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// 记录名称
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// 记录类型
	//
	// This parameter is required.
	//
	// example:
	//
	// A/AAAA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordRequest) SetAuthConf(v *CreateRecordRequestAuthConf) *CreateRecordRequest {
	s.AuthConf = v
	return s
}

func (s *CreateRecordRequest) SetBizName(v string) *CreateRecordRequest {
	s.BizName = &v
	return s
}

func (s *CreateRecordRequest) SetComment(v string) *CreateRecordRequest {
	s.Comment = &v
	return s
}

func (s *CreateRecordRequest) SetData(v *CreateRecordRequestData) *CreateRecordRequest {
	s.Data = v
	return s
}

func (s *CreateRecordRequest) SetHostPolicy(v string) *CreateRecordRequest {
	s.HostPolicy = &v
	return s
}

func (s *CreateRecordRequest) SetProxied(v bool) *CreateRecordRequest {
	s.Proxied = &v
	return s
}

func (s *CreateRecordRequest) SetRecordName(v string) *CreateRecordRequest {
	s.RecordName = &v
	return s
}

func (s *CreateRecordRequest) SetSiteId(v int64) *CreateRecordRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRecordRequest) SetSourceType(v string) *CreateRecordRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRecordRequest) SetTtl(v int32) *CreateRecordRequest {
	s.Ttl = &v
	return s
}

func (s *CreateRecordRequest) SetType(v string) *CreateRecordRequest {
	s.Type = &v
	return s
}

type CreateRecordRequestAuthConf struct {
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// example:
	//
	// v4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateRecordRequestAuthConf) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordRequestAuthConf) GoString() string {
	return s.String()
}

func (s *CreateRecordRequestAuthConf) SetAccessKey(v string) *CreateRecordRequestAuthConf {
	s.AccessKey = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetAuthType(v string) *CreateRecordRequestAuthConf {
	s.AuthType = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetRegion(v string) *CreateRecordRequestAuthConf {
	s.Region = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetSecretKey(v string) *CreateRecordRequestAuthConf {
	s.SecretKey = &v
	return s
}

func (s *CreateRecordRequestAuthConf) SetVersion(v string) *CreateRecordRequestAuthConf {
	s.Version = &v
	return s
}

type CreateRecordRequestData struct {
	// example:
	//
	// 1
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// example:
	//
	// 1
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// RSA
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateRecordRequestData) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordRequestData) GoString() string {
	return s.String()
}

func (s *CreateRecordRequestData) SetAlgorithm(v int32) *CreateRecordRequestData {
	s.Algorithm = &v
	return s
}

func (s *CreateRecordRequestData) SetCertificate(v string) *CreateRecordRequestData {
	s.Certificate = &v
	return s
}

func (s *CreateRecordRequestData) SetFingerprint(v string) *CreateRecordRequestData {
	s.Fingerprint = &v
	return s
}

func (s *CreateRecordRequestData) SetFlag(v int32) *CreateRecordRequestData {
	s.Flag = &v
	return s
}

func (s *CreateRecordRequestData) SetKeyTag(v int32) *CreateRecordRequestData {
	s.KeyTag = &v
	return s
}

func (s *CreateRecordRequestData) SetMatchingType(v int32) *CreateRecordRequestData {
	s.MatchingType = &v
	return s
}

func (s *CreateRecordRequestData) SetPort(v int32) *CreateRecordRequestData {
	s.Port = &v
	return s
}

func (s *CreateRecordRequestData) SetPriority(v int32) *CreateRecordRequestData {
	s.Priority = &v
	return s
}

func (s *CreateRecordRequestData) SetSelector(v int32) *CreateRecordRequestData {
	s.Selector = &v
	return s
}

func (s *CreateRecordRequestData) SetTag(v string) *CreateRecordRequestData {
	s.Tag = &v
	return s
}

func (s *CreateRecordRequestData) SetType(v int32) *CreateRecordRequestData {
	s.Type = &v
	return s
}

func (s *CreateRecordRequestData) SetUsage(v int32) *CreateRecordRequestData {
	s.Usage = &v
	return s
}

func (s *CreateRecordRequestData) SetValue(v string) *CreateRecordRequestData {
	s.Value = &v
	return s
}

func (s *CreateRecordRequestData) SetWeight(v int32) *CreateRecordRequestData {
	s.Weight = &v
	return s
}

type CreateRecordShrinkRequest struct {
	AuthConfShrink *string `json:"AuthConf,omitempty" xml:"AuthConf,omitempty"`
	// 业务场景
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "value":"2.2.2.2"
	//
	// }
	DataShrink *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// 是否代理加速
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// 记录名称
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// 记录类型
	//
	// This parameter is required.
	//
	// example:
	//
	// A/AAAA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRecordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordShrinkRequest) SetAuthConfShrink(v string) *CreateRecordShrinkRequest {
	s.AuthConfShrink = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetBizName(v string) *CreateRecordShrinkRequest {
	s.BizName = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetComment(v string) *CreateRecordShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetDataShrink(v string) *CreateRecordShrinkRequest {
	s.DataShrink = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetHostPolicy(v string) *CreateRecordShrinkRequest {
	s.HostPolicy = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetProxied(v bool) *CreateRecordShrinkRequest {
	s.Proxied = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetRecordName(v string) *CreateRecordShrinkRequest {
	s.RecordName = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetSiteId(v int64) *CreateRecordShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetSourceType(v string) *CreateRecordShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetTtl(v int32) *CreateRecordShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetType(v string) *CreateRecordShrinkRequest {
	s.Type = &v
	return s
}

type CreateRecordResponseBody struct {
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecordResponseBody) SetRecordId(v int64) *CreateRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *CreateRecordResponseBody) SetRequestId(v string) *CreateRecordResponseBody {
	s.RequestId = &v
	return s
}

type CreateRecordResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateRecordResponse) SetHeaders(v map[string]*string) *CreateRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateRecordResponse) SetStatusCode(v int32) *CreateRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecordResponse) SetBody(v *CreateRecordResponseBody) *CreateRecordResponse {
	s.Body = v
	return s
}

type CreateScheduledPreloadExecutionsRequest struct {
	// This parameter is required.
	Executions []*CreateScheduledPreloadExecutionsRequestExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	// example:
	//
	// CreateScheduledPreloadExecutions
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateScheduledPreloadExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsRequest) SetExecutions(v []*CreateScheduledPreloadExecutionsRequestExecutions) *CreateScheduledPreloadExecutionsRequest {
	s.Executions = v
	return s
}

func (s *CreateScheduledPreloadExecutionsRequest) SetId(v string) *CreateScheduledPreloadExecutionsRequest {
	s.Id = &v
	return s
}

type CreateScheduledPreloadExecutionsRequestExecutions struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateScheduledPreloadExecutionsRequestExecutions) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsRequestExecutions) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) SetEndTime(v string) *CreateScheduledPreloadExecutionsRequestExecutions {
	s.EndTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) SetInterval(v int32) *CreateScheduledPreloadExecutionsRequestExecutions {
	s.Interval = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) SetSliceLen(v int32) *CreateScheduledPreloadExecutionsRequestExecutions {
	s.SliceLen = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) SetStartTime(v string) *CreateScheduledPreloadExecutionsRequestExecutions {
	s.StartTime = &v
	return s
}

type CreateScheduledPreloadExecutionsShrinkRequest struct {
	// This parameter is required.
	ExecutionsShrink *string `json:"Executions,omitempty" xml:"Executions,omitempty"`
	// example:
	//
	// CreateScheduledPreloadExecutions
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateScheduledPreloadExecutionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsShrinkRequest) SetExecutionsShrink(v string) *CreateScheduledPreloadExecutionsShrinkRequest {
	s.ExecutionsShrink = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsShrinkRequest) SetId(v string) *CreateScheduledPreloadExecutionsShrinkRequest {
	s.Id = &v
	return s
}

type CreateScheduledPreloadExecutionsResponseBody struct {
	FailedExecutions []*CreateScheduledPreloadExecutionsResponseBodyFailedExecutions `json:"FailedExecutions,omitempty" xml:"FailedExecutions,omitempty" type:"Repeated"`
	FailedMessages   []*string                                                       `json:"FailedMessages,omitempty" xml:"FailedMessages,omitempty" type:"Repeated"`
	// Id of the request
	RequestId         *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessCount      *int32                                                           `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	SuccessExecutions []*CreateScheduledPreloadExecutionsResponseBodySuccessExecutions `json:"SuccessExecutions,omitempty" xml:"SuccessExecutions,omitempty" type:"Repeated"`
	TotalCount        *int32                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateScheduledPreloadExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetFailedExecutions(v []*CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) *CreateScheduledPreloadExecutionsResponseBody {
	s.FailedExecutions = v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetFailedMessages(v []*string) *CreateScheduledPreloadExecutionsResponseBody {
	s.FailedMessages = v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetRequestId(v string) *CreateScheduledPreloadExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetSuccessCount(v int32) *CreateScheduledPreloadExecutionsResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetSuccessExecutions(v []*CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) *CreateScheduledPreloadExecutionsResponseBody {
	s.SuccessExecutions = v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetTotalCount(v int32) *CreateScheduledPreloadExecutionsResponseBody {
	s.TotalCount = &v
	return s
}

type CreateScheduledPreloadExecutionsResponseBodyFailedExecutions struct {
	AliUid    *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetAliUid(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.AliUid = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetEndTime(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.EndTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetId(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.Id = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetInterval(v int32) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.Interval = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetJobId(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.JobId = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetSliceLen(v int32) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.SliceLen = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetStartTime(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.StartTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetStatus(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.Status = &v
	return s
}

type CreateScheduledPreloadExecutionsResponseBodySuccessExecutions struct {
	AliUid    *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetAliUid(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.AliUid = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetEndTime(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.EndTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetId(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.Id = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetInterval(v int32) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.Interval = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetJobId(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.JobId = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetSliceLen(v int32) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.SliceLen = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetStartTime(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.StartTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetStatus(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.Status = &v
	return s
}

type CreateScheduledPreloadExecutionsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledPreloadExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledPreloadExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsResponse) SetHeaders(v map[string]*string) *CreateScheduledPreloadExecutionsResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponse) SetStatusCode(v int32) *CreateScheduledPreloadExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponse) SetBody(v *CreateScheduledPreloadExecutionsResponseBody) *CreateScheduledPreloadExecutionsResponse {
	s.Body = v
	return s
}

type CreateScheduledPreloadJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// oss
	InsertWay *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	// This parameter is required.
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// This parameter is required.
	SiteId  *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	UrlList *string `json:"UrlList,omitempty" xml:"UrlList,omitempty"`
}

func (s CreateScheduledPreloadJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPreloadJobRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadJobRequest) SetInsertWay(v string) *CreateScheduledPreloadJobRequest {
	s.InsertWay = &v
	return s
}

func (s *CreateScheduledPreloadJobRequest) SetName(v string) *CreateScheduledPreloadJobRequest {
	s.Name = &v
	return s
}

func (s *CreateScheduledPreloadJobRequest) SetOssUrl(v string) *CreateScheduledPreloadJobRequest {
	s.OssUrl = &v
	return s
}

func (s *CreateScheduledPreloadJobRequest) SetSiteId(v int64) *CreateScheduledPreloadJobRequest {
	s.SiteId = &v
	return s
}

func (s *CreateScheduledPreloadJobRequest) SetUrlList(v string) *CreateScheduledPreloadJobRequest {
	s.UrlList = &v
	return s
}

type CreateScheduledPreloadJobResponseBody struct {
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CreatedAt     *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Domains       *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	ErrorInfo     *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	FailedFileOss *string `json:"FailedFileOss,omitempty" xml:"FailedFileOss,omitempty"`
	FileId        *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InsertWay     *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId        *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TaskSubmitted *int32  `json:"TaskSubmitted,omitempty" xml:"TaskSubmitted,omitempty"`
	TaskType      *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UrlCount      *int32  `json:"UrlCount,omitempty" xml:"UrlCount,omitempty"`
	UrlSubmitted  *int32  `json:"UrlSubmitted,omitempty" xml:"UrlSubmitted,omitempty"`
}

func (s CreateScheduledPreloadJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPreloadJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadJobResponseBody) SetAliUid(v string) *CreateScheduledPreloadJobResponseBody {
	s.AliUid = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetCreatedAt(v string) *CreateScheduledPreloadJobResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetDomains(v string) *CreateScheduledPreloadJobResponseBody {
	s.Domains = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetErrorInfo(v string) *CreateScheduledPreloadJobResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetFailedFileOss(v string) *CreateScheduledPreloadJobResponseBody {
	s.FailedFileOss = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetFileId(v string) *CreateScheduledPreloadJobResponseBody {
	s.FileId = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetId(v string) *CreateScheduledPreloadJobResponseBody {
	s.Id = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetInsertWay(v string) *CreateScheduledPreloadJobResponseBody {
	s.InsertWay = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetName(v string) *CreateScheduledPreloadJobResponseBody {
	s.Name = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetRequestId(v string) *CreateScheduledPreloadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetSiteId(v int64) *CreateScheduledPreloadJobResponseBody {
	s.SiteId = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetTaskSubmitted(v int32) *CreateScheduledPreloadJobResponseBody {
	s.TaskSubmitted = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetTaskType(v string) *CreateScheduledPreloadJobResponseBody {
	s.TaskType = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetUrlCount(v int32) *CreateScheduledPreloadJobResponseBody {
	s.UrlCount = &v
	return s
}

func (s *CreateScheduledPreloadJobResponseBody) SetUrlSubmitted(v int32) *CreateScheduledPreloadJobResponseBody {
	s.UrlSubmitted = &v
	return s
}

type CreateScheduledPreloadJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledPreloadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledPreloadJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPreloadJobResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadJobResponse) SetHeaders(v map[string]*string) *CreateScheduledPreloadJobResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledPreloadJobResponse) SetStatusCode(v int32) *CreateScheduledPreloadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledPreloadJobResponse) SetBody(v *CreateScheduledPreloadJobResponseBody) *CreateScheduledPreloadJobResponse {
	s.Body = v
	return s
}

type CreateSiteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// NS
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-nwy349jdb03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// rg-acfmw4znnok****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 记录名称
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateSite
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s CreateSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteRequest) SetAccessType(v string) *CreateSiteRequest {
	s.AccessType = &v
	return s
}

func (s *CreateSiteRequest) SetCoverage(v string) *CreateSiteRequest {
	s.Coverage = &v
	return s
}

func (s *CreateSiteRequest) SetInstanceId(v string) *CreateSiteRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSiteRequest) SetResourceGroupId(v string) *CreateSiteRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSiteRequest) SetSiteName(v string) *CreateSiteRequest {
	s.SiteName = &v
	return s
}

type CreateSiteResponseBody struct {
	// example:
	//
	// ns1.example.com,ns2.example.com
	NameServerList *string `json:"NameServerList,omitempty" xml:"NameServerList,omitempty"`
	// example:
	//
	// CB1A380B-09F0-41BB-3C82-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// verify_aah9dioasmov****
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s CreateSiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSiteResponseBody) SetNameServerList(v string) *CreateSiteResponseBody {
	s.NameServerList = &v
	return s
}

func (s *CreateSiteResponseBody) SetRequestId(v string) *CreateSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSiteResponseBody) SetSiteId(v int64) *CreateSiteResponseBody {
	s.SiteId = &v
	return s
}

func (s *CreateSiteResponseBody) SetVerifyCode(v string) *CreateSiteResponseBody {
	s.VerifyCode = &v
	return s
}

type CreateSiteResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteResponse) GoString() string {
	return s.String()
}

func (s *CreateSiteResponse) SetHeaders(v map[string]*string) *CreateSiteResponse {
	s.Headers = v
	return s
}

func (s *CreateSiteResponse) SetStatusCode(v int32) *CreateSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSiteResponse) SetBody(v *CreateSiteResponseBody) *CreateSiteResponse {
	s.Body = v
	return s
}

type CreateSiteCustomLogRequest struct {
	Cookies         []*string `json:"Cookies,omitempty" xml:"Cookies,omitempty" type:"Repeated"`
	RequestHeaders  []*string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	ResponseHeaders []*string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
	// example:
	//
	// 11223
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateSiteCustomLogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteCustomLogRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteCustomLogRequest) SetCookies(v []*string) *CreateSiteCustomLogRequest {
	s.Cookies = v
	return s
}

func (s *CreateSiteCustomLogRequest) SetRequestHeaders(v []*string) *CreateSiteCustomLogRequest {
	s.RequestHeaders = v
	return s
}

func (s *CreateSiteCustomLogRequest) SetResponseHeaders(v []*string) *CreateSiteCustomLogRequest {
	s.ResponseHeaders = v
	return s
}

func (s *CreateSiteCustomLogRequest) SetSiteId(v int64) *CreateSiteCustomLogRequest {
	s.SiteId = &v
	return s
}

type CreateSiteCustomLogShrinkRequest struct {
	CookiesShrink         *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	RequestHeadersShrink  *string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	ResponseHeadersShrink *string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty"`
	// example:
	//
	// 11223
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateSiteCustomLogShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteCustomLogShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteCustomLogShrinkRequest) SetCookiesShrink(v string) *CreateSiteCustomLogShrinkRequest {
	s.CookiesShrink = &v
	return s
}

func (s *CreateSiteCustomLogShrinkRequest) SetRequestHeadersShrink(v string) *CreateSiteCustomLogShrinkRequest {
	s.RequestHeadersShrink = &v
	return s
}

func (s *CreateSiteCustomLogShrinkRequest) SetResponseHeadersShrink(v string) *CreateSiteCustomLogShrinkRequest {
	s.ResponseHeadersShrink = &v
	return s
}

func (s *CreateSiteCustomLogShrinkRequest) SetSiteId(v int64) *CreateSiteCustomLogShrinkRequest {
	s.SiteId = &v
	return s
}

type CreateSiteCustomLogResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 073bd613-6e72-4461-b6bc-19326dfc6a9c
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSiteCustomLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteCustomLogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSiteCustomLogResponseBody) SetRequestId(v string) *CreateSiteCustomLogResponseBody {
	s.RequestId = &v
	return s
}

type CreateSiteCustomLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSiteCustomLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSiteCustomLogResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteCustomLogResponse) GoString() string {
	return s.String()
}

func (s *CreateSiteCustomLogResponse) SetHeaders(v map[string]*string) *CreateSiteCustomLogResponse {
	s.Headers = v
	return s
}

func (s *CreateSiteCustomLogResponse) SetStatusCode(v int32) *CreateSiteCustomLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSiteCustomLogResponse) SetBody(v *CreateSiteCustomLogResponseBody) *CreateSiteCustomLogResponse {
	s.Body = v
	return s
}

type CreateSiteDeliveryTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// 0.0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_agent,ip_adress,ip_port
	FieldName     *string                                     `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	HttpDelivery  *CreateSiteDeliveryTaskRequestHttpDelivery  `json:"HttpDelivery,omitempty" xml:"HttpDelivery,omitempty" type:"Struct"`
	KafkaDelivery *CreateSiteDeliveryTaskRequestKafkaDelivery `json:"KafkaDelivery,omitempty" xml:"KafkaDelivery,omitempty" type:"Struct"`
	OssDelivery   *CreateSiteDeliveryTaskRequestOssDelivery   `json:"OssDelivery,omitempty" xml:"OssDelivery,omitempty" type:"Struct"`
	S3Delivery    *CreateSiteDeliveryTaskRequestS3Delivery    `json:"S3Delivery,omitempty" xml:"S3Delivery,omitempty" type:"Struct"`
	// example:
	//
	// 12312312112***
	SiteId      *int64                                    `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SlsDelivery *CreateSiteDeliveryTaskRequestSlsDelivery `json:"SlsDelivery,omitempty" xml:"SlsDelivery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// dcdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateSiteDeliveryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequest) SetBusinessType(v string) *CreateSiteDeliveryTaskRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetDataCenter(v string) *CreateSiteDeliveryTaskRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetDeliveryType(v string) *CreateSiteDeliveryTaskRequest {
	s.DeliveryType = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetDiscardRate(v float32) *CreateSiteDeliveryTaskRequest {
	s.DiscardRate = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetFieldName(v string) *CreateSiteDeliveryTaskRequest {
	s.FieldName = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetHttpDelivery(v *CreateSiteDeliveryTaskRequestHttpDelivery) *CreateSiteDeliveryTaskRequest {
	s.HttpDelivery = v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetKafkaDelivery(v *CreateSiteDeliveryTaskRequestKafkaDelivery) *CreateSiteDeliveryTaskRequest {
	s.KafkaDelivery = v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetOssDelivery(v *CreateSiteDeliveryTaskRequestOssDelivery) *CreateSiteDeliveryTaskRequest {
	s.OssDelivery = v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetS3Delivery(v *CreateSiteDeliveryTaskRequestS3Delivery) *CreateSiteDeliveryTaskRequest {
	s.S3Delivery = v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetSiteId(v int64) *CreateSiteDeliveryTaskRequest {
	s.SiteId = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetSlsDelivery(v *CreateSiteDeliveryTaskRequestSlsDelivery) *CreateSiteDeliveryTaskRequest {
	s.SlsDelivery = v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetTaskName(v string) *CreateSiteDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

type CreateSiteDeliveryTaskRequestHttpDelivery struct {
	Compress          *string                                                     `json:"Compress,omitempty" xml:"Compress,omitempty"`
	DestUrl           *string                                                     `json:"DestUrl,omitempty" xml:"DestUrl,omitempty"`
	HeaderParam       map[string]*HttpDeliveryHeaderParamValue                    `json:"HeaderParam,omitempty" xml:"HeaderParam,omitempty"`
	LogBodyPrefix     *string                                                     `json:"LogBodyPrefix,omitempty" xml:"LogBodyPrefix,omitempty"`
	LogBodySuffix     *string                                                     `json:"LogBodySuffix,omitempty" xml:"LogBodySuffix,omitempty"`
	MaxBatchMB        *int64                                                      `json:"MaxBatchMB,omitempty" xml:"MaxBatchMB,omitempty"`
	MaxBatchSize      *int64                                                      `json:"MaxBatchSize,omitempty" xml:"MaxBatchSize,omitempty"`
	MaxRetry          *int64                                                      `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	QueryParam        map[string]*HttpDeliveryQueryParamValue                     `json:"QueryParam,omitempty" xml:"QueryParam,omitempty"`
	StandardAuthOn    *bool                                                       `json:"StandardAuthOn,omitempty" xml:"StandardAuthOn,omitempty"`
	StandardAuthParam *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam `json:"StandardAuthParam,omitempty" xml:"StandardAuthParam,omitempty" type:"Struct"`
	TransformTimeout  *int64                                                      `json:"TransformTimeout,omitempty" xml:"TransformTimeout,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestHttpDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestHttpDelivery) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetCompress(v string) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.Compress = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetDestUrl(v string) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.DestUrl = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetHeaderParam(v map[string]*HttpDeliveryHeaderParamValue) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.HeaderParam = v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetLogBodyPrefix(v string) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.LogBodyPrefix = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetLogBodySuffix(v string) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.LogBodySuffix = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetMaxBatchMB(v int64) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.MaxBatchMB = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetMaxBatchSize(v int64) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.MaxBatchSize = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetMaxRetry(v int64) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.MaxRetry = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetQueryParam(v map[string]*HttpDeliveryQueryParamValue) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.QueryParam = v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetStandardAuthOn(v bool) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.StandardAuthOn = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetStandardAuthParam(v *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.StandardAuthParam = v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetTransformTimeout(v int64) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.TransformTimeout = &v
	return s
}

type CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam struct {
	ExpiredTime *int32  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	PrivateKey  *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	UrlPath     *string `json:"UrlPath,omitempty" xml:"UrlPath,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetExpiredTime(v int32) *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.ExpiredTime = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetPrivateKey(v string) *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.PrivateKey = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetUrlPath(v string) *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.UrlPath = &v
	return s
}

type CreateSiteDeliveryTaskRequestKafkaDelivery struct {
	Balancer *string   `json:"Balancer,omitempty" xml:"Balancer,omitempty"`
	Brokers  []*string `json:"Brokers,omitempty" xml:"Brokers,omitempty" type:"Repeated"`
	// example:
	//
	// gzip
	Compress      *string `json:"Compress,omitempty" xml:"Compress,omitempty"`
	MachanismType *string `json:"MachanismType,omitempty" xml:"MachanismType,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	UserAuth      *bool   `json:"UserAuth,omitempty" xml:"UserAuth,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestKafkaDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestKafkaDelivery) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetBalancer(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.Balancer = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetBrokers(v []*string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.Brokers = v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetCompress(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.Compress = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetMachanismType(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.MachanismType = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetPassword(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.Password = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetTopic(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.Topic = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetUserAuth(v bool) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.UserAuth = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetUserName(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.UserName = &v
	return s
}

type CreateSiteDeliveryTaskRequestOssDelivery struct {
	Aliuid     *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// logriver-test/log
	PrefixPath *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestOssDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestOssDelivery) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) SetAliuid(v string) *CreateSiteDeliveryTaskRequestOssDelivery {
	s.Aliuid = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) SetBucketName(v string) *CreateSiteDeliveryTaskRequestOssDelivery {
	s.BucketName = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) SetPrefixPath(v string) *CreateSiteDeliveryTaskRequestOssDelivery {
	s.PrefixPath = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) SetRegion(v string) *CreateSiteDeliveryTaskRequestOssDelivery {
	s.Region = &v
	return s
}

type CreateSiteDeliveryTaskRequestS3Delivery struct {
	AccessKey            *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	BucketPath           *string `json:"BucketPath,omitempty" xml:"BucketPath,omitempty"`
	Endpoint             *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	PrefixPath           *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	S3Cmpt               *bool   `json:"S3Cmpt,omitempty" xml:"S3Cmpt,omitempty"`
	SecretKey            *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	ServerSideEncryption *bool   `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	VertifyType          *string `json:"VertifyType,omitempty" xml:"VertifyType,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestS3Delivery) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestS3Delivery) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetAccessKey(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.AccessKey = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetBucketPath(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.BucketPath = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetEndpoint(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.Endpoint = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetPrefixPath(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.PrefixPath = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetRegion(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.Region = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetS3Cmpt(v bool) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.S3Cmpt = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetSecretKey(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.SecretKey = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetServerSideEncryption(v bool) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.ServerSideEncryption = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetVertifyType(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.VertifyType = &v
	return s
}

type CreateSiteDeliveryTaskRequestSlsDelivery struct {
	SLSLogStore *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	SLSProject  *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	SLSRegion   *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestSlsDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestSlsDelivery) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestSlsDelivery) SetSLSLogStore(v string) *CreateSiteDeliveryTaskRequestSlsDelivery {
	s.SLSLogStore = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestSlsDelivery) SetSLSProject(v string) *CreateSiteDeliveryTaskRequestSlsDelivery {
	s.SLSProject = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestSlsDelivery) SetSLSRegion(v string) *CreateSiteDeliveryTaskRequestSlsDelivery {
	s.SLSRegion = &v
	return s
}

type CreateSiteDeliveryTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// 0.0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_agent,ip_adress,ip_port
	FieldName           *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	HttpDeliveryShrink  *string `json:"HttpDelivery,omitempty" xml:"HttpDelivery,omitempty"`
	KafkaDeliveryShrink *string `json:"KafkaDelivery,omitempty" xml:"KafkaDelivery,omitempty"`
	OssDeliveryShrink   *string `json:"OssDelivery,omitempty" xml:"OssDelivery,omitempty"`
	S3DeliveryShrink    *string `json:"S3Delivery,omitempty" xml:"S3Delivery,omitempty"`
	// example:
	//
	// 12312312112***
	SiteId            *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SlsDeliveryShrink *string `json:"SlsDelivery,omitempty" xml:"SlsDelivery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dcdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateSiteDeliveryTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteDeliveryTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetBusinessType(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetDataCenter(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetDeliveryType(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.DeliveryType = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetDiscardRate(v float32) *CreateSiteDeliveryTaskShrinkRequest {
	s.DiscardRate = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetFieldName(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.FieldName = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetHttpDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.HttpDeliveryShrink = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetKafkaDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.KafkaDeliveryShrink = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetOssDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.OssDeliveryShrink = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetS3DeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.S3DeliveryShrink = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetSiteId(v int64) *CreateSiteDeliveryTaskShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetSlsDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.SlsDeliveryShrink = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetTaskName(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.TaskName = &v
	return s
}

type CreateSiteDeliveryTaskResponseBody struct {
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9358E852-992D-5BC7-8BD7-975CA02773A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456****
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// er-oss
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateSiteDeliveryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskResponseBody) SetDataCenter(v string) *CreateSiteDeliveryTaskResponseBody {
	s.DataCenter = &v
	return s
}

func (s *CreateSiteDeliveryTaskResponseBody) SetRequestId(v string) *CreateSiteDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSiteDeliveryTaskResponseBody) SetSiteId(v string) *CreateSiteDeliveryTaskResponseBody {
	s.SiteId = &v
	return s
}

func (s *CreateSiteDeliveryTaskResponseBody) SetTaskName(v string) *CreateSiteDeliveryTaskResponseBody {
	s.TaskName = &v
	return s
}

type CreateSiteDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSiteDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSiteDeliveryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskResponse) SetHeaders(v map[string]*string) *CreateSiteDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSiteDeliveryTaskResponse) SetStatusCode(v int32) *CreateSiteDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSiteDeliveryTaskResponse) SetBody(v *CreateSiteDeliveryTaskResponseBody) *CreateSiteDeliveryTaskResponse {
	s.Body = v
	return s
}

type CreateUserDeliveryTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// 0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_agent,ip_address,ip_port
	FieldName     *string                                     `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	HttpDelivery  *CreateUserDeliveryTaskRequestHttpDelivery  `json:"HttpDelivery,omitempty" xml:"HttpDelivery,omitempty" type:"Struct"`
	KafkaDelivery *CreateUserDeliveryTaskRequestKafkaDelivery `json:"KafkaDelivery,omitempty" xml:"KafkaDelivery,omitempty" type:"Struct"`
	OssDelivery   *CreateUserDeliveryTaskRequestOssDelivery   `json:"OssDelivery,omitempty" xml:"OssDelivery,omitempty" type:"Struct"`
	S3Delivery    *CreateUserDeliveryTaskRequestS3Delivery    `json:"S3Delivery,omitempty" xml:"S3Delivery,omitempty" type:"Struct"`
	SlsDelivery   *CreateUserDeliveryTaskRequestSlsDelivery   `json:"SlsDelivery,omitempty" xml:"SlsDelivery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateUserDeliveryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequest) SetBusinessType(v string) *CreateUserDeliveryTaskRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetDataCenter(v string) *CreateUserDeliveryTaskRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetDeliveryType(v string) *CreateUserDeliveryTaskRequest {
	s.DeliveryType = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetDiscardRate(v float32) *CreateUserDeliveryTaskRequest {
	s.DiscardRate = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetFieldName(v string) *CreateUserDeliveryTaskRequest {
	s.FieldName = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetHttpDelivery(v *CreateUserDeliveryTaskRequestHttpDelivery) *CreateUserDeliveryTaskRequest {
	s.HttpDelivery = v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetKafkaDelivery(v *CreateUserDeliveryTaskRequestKafkaDelivery) *CreateUserDeliveryTaskRequest {
	s.KafkaDelivery = v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetOssDelivery(v *CreateUserDeliveryTaskRequestOssDelivery) *CreateUserDeliveryTaskRequest {
	s.OssDelivery = v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetS3Delivery(v *CreateUserDeliveryTaskRequestS3Delivery) *CreateUserDeliveryTaskRequest {
	s.S3Delivery = v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetSlsDelivery(v *CreateUserDeliveryTaskRequestSlsDelivery) *CreateUserDeliveryTaskRequest {
	s.SlsDelivery = v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetTaskName(v string) *CreateUserDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

type CreateUserDeliveryTaskRequestHttpDelivery struct {
	Compress          *string                                                     `json:"Compress,omitempty" xml:"Compress,omitempty"`
	DestUrl           *string                                                     `json:"DestUrl,omitempty" xml:"DestUrl,omitempty"`
	HeaderParam       map[string]*HttpDeliveryHeaderParamValue                    `json:"HeaderParam,omitempty" xml:"HeaderParam,omitempty"`
	LastLogSplit      *string                                                     `json:"LastLogSplit,omitempty" xml:"LastLogSplit,omitempty"`
	LogBodyPrefix     *string                                                     `json:"LogBodyPrefix,omitempty" xml:"LogBodyPrefix,omitempty"`
	LogBodySuffix     *string                                                     `json:"LogBodySuffix,omitempty" xml:"LogBodySuffix,omitempty"`
	LogSplit          *string                                                     `json:"LogSplit,omitempty" xml:"LogSplit,omitempty"`
	LogSplitWords     *string                                                     `json:"LogSplitWords,omitempty" xml:"LogSplitWords,omitempty"`
	MaxBackoffMS      *int64                                                      `json:"MaxBackoffMS,omitempty" xml:"MaxBackoffMS,omitempty"`
	MaxBatchMB        *int64                                                      `json:"MaxBatchMB,omitempty" xml:"MaxBatchMB,omitempty"`
	MaxBatchSize      *int64                                                      `json:"MaxBatchSize,omitempty" xml:"MaxBatchSize,omitempty"`
	MaxRetry          *int64                                                      `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	MinBackoffMS      *int64                                                      `json:"MinBackoffMS,omitempty" xml:"MinBackoffMS,omitempty"`
	QueryParam        map[string]*HttpDeliveryQueryParamValue                     `json:"QueryParam,omitempty" xml:"QueryParam,omitempty"`
	ResponseBodyKey   *string                                                     `json:"ResponseBodyKey,omitempty" xml:"ResponseBodyKey,omitempty"`
	StandardAuthOn    *bool                                                       `json:"StandardAuthOn,omitempty" xml:"StandardAuthOn,omitempty"`
	StandardAuthParam *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam `json:"StandardAuthParam,omitempty" xml:"StandardAuthParam,omitempty" type:"Struct"`
	SuccessCode       *int64                                                      `json:"SuccessCode,omitempty" xml:"SuccessCode,omitempty"`
	TransformTimeout  *int64                                                      `json:"TransformTimeout,omitempty" xml:"TransformTimeout,omitempty"`
}

func (s CreateUserDeliveryTaskRequestHttpDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestHttpDelivery) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetCompress(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.Compress = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetDestUrl(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.DestUrl = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetHeaderParam(v map[string]*HttpDeliveryHeaderParamValue) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.HeaderParam = v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetLastLogSplit(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.LastLogSplit = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetLogBodyPrefix(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.LogBodyPrefix = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetLogBodySuffix(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.LogBodySuffix = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetLogSplit(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.LogSplit = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetLogSplitWords(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.LogSplitWords = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetMaxBackoffMS(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.MaxBackoffMS = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetMaxBatchMB(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.MaxBatchMB = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetMaxBatchSize(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.MaxBatchSize = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetMaxRetry(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.MaxRetry = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetMinBackoffMS(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.MinBackoffMS = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetQueryParam(v map[string]*HttpDeliveryQueryParamValue) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.QueryParam = v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetResponseBodyKey(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.ResponseBodyKey = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetStandardAuthOn(v bool) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.StandardAuthOn = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetStandardAuthParam(v *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.StandardAuthParam = v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetSuccessCode(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.SuccessCode = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetTransformTimeout(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.TransformTimeout = &v
	return s
}

type CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam struct {
	ExpiredTime *int32  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	PrivateKey  *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	UrlPath     *string `json:"UrlPath,omitempty" xml:"UrlPath,omitempty"`
}

func (s CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetExpiredTime(v int32) *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.ExpiredTime = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetPrivateKey(v string) *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.PrivateKey = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetUrlPath(v string) *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.UrlPath = &v
	return s
}

type CreateUserDeliveryTaskRequestKafkaDelivery struct {
	Balancer *string   `json:"Balancer,omitempty" xml:"Balancer,omitempty"`
	Brokers  []*string `json:"Brokers,omitempty" xml:"Brokers,omitempty" type:"Repeated"`
	// example:
	//
	// gzip
	Compress      *string `json:"Compress,omitempty" xml:"Compress,omitempty"`
	MachanismType *string `json:"MachanismType,omitempty" xml:"MachanismType,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	UserAuth      *bool   `json:"UserAuth,omitempty" xml:"UserAuth,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserDeliveryTaskRequestKafkaDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestKafkaDelivery) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetBalancer(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.Balancer = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetBrokers(v []*string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.Brokers = v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetCompress(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.Compress = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetMachanismType(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.MachanismType = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetPassword(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.Password = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetTopic(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.Topic = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetUserAuth(v bool) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.UserAuth = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetUserName(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.UserName = &v
	return s
}

type CreateUserDeliveryTaskRequestOssDelivery struct {
	Aliuid     *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// logriver-test/log
	PrefixPath *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateUserDeliveryTaskRequestOssDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestOssDelivery) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) SetAliuid(v string) *CreateUserDeliveryTaskRequestOssDelivery {
	s.Aliuid = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) SetBucketName(v string) *CreateUserDeliveryTaskRequestOssDelivery {
	s.BucketName = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) SetPrefixPath(v string) *CreateUserDeliveryTaskRequestOssDelivery {
	s.PrefixPath = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) SetRegion(v string) *CreateUserDeliveryTaskRequestOssDelivery {
	s.Region = &v
	return s
}

type CreateUserDeliveryTaskRequestS3Delivery struct {
	AccessKey            *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	BucketPath           *string `json:"BucketPath,omitempty" xml:"BucketPath,omitempty"`
	Endpoint             *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	PrefixPath           *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	S3Cmpt               *bool   `json:"S3Cmpt,omitempty" xml:"S3Cmpt,omitempty"`
	SecretKey            *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	ServerSideEncryption *bool   `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	VertifyType          *string `json:"VertifyType,omitempty" xml:"VertifyType,omitempty"`
}

func (s CreateUserDeliveryTaskRequestS3Delivery) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestS3Delivery) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetAccessKey(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.AccessKey = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetBucketPath(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.BucketPath = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetEndpoint(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.Endpoint = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetPrefixPath(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.PrefixPath = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetRegion(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.Region = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetS3Cmpt(v bool) *CreateUserDeliveryTaskRequestS3Delivery {
	s.S3Cmpt = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetSecretKey(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.SecretKey = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetServerSideEncryption(v bool) *CreateUserDeliveryTaskRequestS3Delivery {
	s.ServerSideEncryption = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetVertifyType(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.VertifyType = &v
	return s
}

type CreateUserDeliveryTaskRequestSlsDelivery struct {
	SLSLogStore *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	SLSProject  *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	SLSRegion   *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
}

func (s CreateUserDeliveryTaskRequestSlsDelivery) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestSlsDelivery) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestSlsDelivery) SetSLSLogStore(v string) *CreateUserDeliveryTaskRequestSlsDelivery {
	s.SLSLogStore = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestSlsDelivery) SetSLSProject(v string) *CreateUserDeliveryTaskRequestSlsDelivery {
	s.SLSProject = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestSlsDelivery) SetSLSRegion(v string) *CreateUserDeliveryTaskRequestSlsDelivery {
	s.SLSRegion = &v
	return s
}

type CreateUserDeliveryTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// 0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_agent,ip_address,ip_port
	FieldName           *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	HttpDeliveryShrink  *string `json:"HttpDelivery,omitempty" xml:"HttpDelivery,omitempty"`
	KafkaDeliveryShrink *string `json:"KafkaDelivery,omitempty" xml:"KafkaDelivery,omitempty"`
	OssDeliveryShrink   *string `json:"OssDelivery,omitempty" xml:"OssDelivery,omitempty"`
	S3DeliveryShrink    *string `json:"S3Delivery,omitempty" xml:"S3Delivery,omitempty"`
	SlsDeliveryShrink   *string `json:"SlsDelivery,omitempty" xml:"SlsDelivery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateUserDeliveryTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDeliveryTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetBusinessType(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetDataCenter(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetDeliveryType(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.DeliveryType = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetDiscardRate(v float32) *CreateUserDeliveryTaskShrinkRequest {
	s.DiscardRate = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetFieldName(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.FieldName = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetHttpDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.HttpDeliveryShrink = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetKafkaDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.KafkaDeliveryShrink = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetOssDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.OssDeliveryShrink = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetS3DeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.S3DeliveryShrink = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetSlsDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.SlsDeliveryShrink = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetTaskName(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.TaskName = &v
	return s
}

type CreateUserDeliveryTaskResponseBody struct {
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2CCD40B1-3F20-5FF0-8A67-E3F34B87744F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// er-http
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateUserDeliveryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskResponseBody) SetDataCenter(v string) *CreateUserDeliveryTaskResponseBody {
	s.DataCenter = &v
	return s
}

func (s *CreateUserDeliveryTaskResponseBody) SetRequestId(v string) *CreateUserDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserDeliveryTaskResponseBody) SetStatus(v string) *CreateUserDeliveryTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateUserDeliveryTaskResponseBody) SetTaskName(v string) *CreateUserDeliveryTaskResponseBody {
	s.TaskName = &v
	return s
}

type CreateUserDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserDeliveryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskResponse) SetHeaders(v map[string]*string) *CreateUserDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateUserDeliveryTaskResponse) SetStatusCode(v int32) *CreateUserDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserDeliveryTaskResponse) SetBody(v *CreateUserDeliveryTaskResponseBody) *CreateUserDeliveryTaskResponse {
	s.Body = v
	return s
}

type CreateWafRuleRequest struct {
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateWafRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWafRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWafRuleRequest) SetConfig(v *WafRuleConfig) *CreateWafRuleRequest {
	s.Config = v
	return s
}

func (s *CreateWafRuleRequest) SetPhase(v string) *CreateWafRuleRequest {
	s.Phase = &v
	return s
}

func (s *CreateWafRuleRequest) SetSiteId(v int64) *CreateWafRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWafRuleRequest) SetSiteVersion(v int32) *CreateWafRuleRequest {
	s.SiteVersion = &v
	return s
}

type CreateWafRuleShrinkRequest struct {
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateWafRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWafRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWafRuleShrinkRequest) SetConfigShrink(v string) *CreateWafRuleShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *CreateWafRuleShrinkRequest) SetPhase(v string) *CreateWafRuleShrinkRequest {
	s.Phase = &v
	return s
}

func (s *CreateWafRuleShrinkRequest) SetSiteId(v int64) *CreateWafRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWafRuleShrinkRequest) SetSiteVersion(v int32) *CreateWafRuleShrinkRequest {
	s.SiteVersion = &v
	return s
}

type CreateWafRuleResponseBody struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RulesetId *int64  `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
}

func (s CreateWafRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWafRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWafRuleResponseBody) SetId(v int64) *CreateWafRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateWafRuleResponseBody) SetRequestId(v string) *CreateWafRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWafRuleResponseBody) SetRulesetId(v int64) *CreateWafRuleResponseBody {
	s.RulesetId = &v
	return s
}

type CreateWafRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWafRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWafRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWafRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateWafRuleResponse) SetHeaders(v map[string]*string) *CreateWafRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateWafRuleResponse) SetStatusCode(v int32) *CreateWafRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWafRuleResponse) SetBody(v *CreateWafRuleResponseBody) *CreateWafRuleResponse {
	s.Body = v
	return s
}

type CreateWaitingRoomRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// __aliwaitingroom_example
	CookieName *string `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	// example:
	//
	// Hello%20world!
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// on
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	HostNameAndPath []*CreateWaitingRoomRequestHostNameAndPath `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty" type:"Repeated"`
	// example:
	//
	// on
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// example:
	//
	// enus
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waitingroom_example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 200
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// example:
	//
	// on
	QueueAllEnable *string `json:"QueueAllEnable,omitempty" xml:"QueueAllEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fifo
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 300
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s CreateWaitingRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRequest) SetCookieName(v string) *CreateWaitingRoomRequest {
	s.CookieName = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetCustomPageHtml(v string) *CreateWaitingRoomRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetDescription(v string) *CreateWaitingRoomRequest {
	s.Description = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetDisableSessionRenewalEnable(v string) *CreateWaitingRoomRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetEnable(v string) *CreateWaitingRoomRequest {
	s.Enable = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetHostNameAndPath(v []*CreateWaitingRoomRequestHostNameAndPath) *CreateWaitingRoomRequest {
	s.HostNameAndPath = v
	return s
}

func (s *CreateWaitingRoomRequest) SetJsonResponseEnable(v string) *CreateWaitingRoomRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetLanguage(v string) *CreateWaitingRoomRequest {
	s.Language = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetName(v string) *CreateWaitingRoomRequest {
	s.Name = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetNewUsersPerMinute(v string) *CreateWaitingRoomRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetQueueAllEnable(v string) *CreateWaitingRoomRequest {
	s.QueueAllEnable = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetQueuingMethod(v string) *CreateWaitingRoomRequest {
	s.QueuingMethod = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetQueuingStatusCode(v string) *CreateWaitingRoomRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetSessionDuration(v string) *CreateWaitingRoomRequest {
	s.SessionDuration = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetSiteId(v int64) *CreateWaitingRoomRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetTotalActiveUsers(v string) *CreateWaitingRoomRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetWaitingRoomType(v string) *CreateWaitingRoomRequest {
	s.WaitingRoomType = &v
	return s
}

type CreateWaitingRoomRequestHostNameAndPath struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.
	Subdomain *string `json:"Subdomain,omitempty" xml:"Subdomain,omitempty"`
}

func (s CreateWaitingRoomRequestHostNameAndPath) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomRequestHostNameAndPath) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRequestHostNameAndPath) SetDomain(v string) *CreateWaitingRoomRequestHostNameAndPath {
	s.Domain = &v
	return s
}

func (s *CreateWaitingRoomRequestHostNameAndPath) SetPath(v string) *CreateWaitingRoomRequestHostNameAndPath {
	s.Path = &v
	return s
}

func (s *CreateWaitingRoomRequestHostNameAndPath) SetSubdomain(v string) *CreateWaitingRoomRequestHostNameAndPath {
	s.Subdomain = &v
	return s
}

type CreateWaitingRoomShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// __aliwaitingroom_example
	CookieName *string `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	// example:
	//
	// Hello%20world!
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// on
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	HostNameAndPathShrink *string `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty"`
	// example:
	//
	// on
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// example:
	//
	// enus
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waitingroom_example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 200
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// example:
	//
	// on
	QueueAllEnable *string `json:"QueueAllEnable,omitempty" xml:"QueueAllEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fifo
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 300
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s CreateWaitingRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomShrinkRequest) SetCookieName(v string) *CreateWaitingRoomShrinkRequest {
	s.CookieName = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetCustomPageHtml(v string) *CreateWaitingRoomShrinkRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetDescription(v string) *CreateWaitingRoomShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetDisableSessionRenewalEnable(v string) *CreateWaitingRoomShrinkRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetEnable(v string) *CreateWaitingRoomShrinkRequest {
	s.Enable = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetHostNameAndPathShrink(v string) *CreateWaitingRoomShrinkRequest {
	s.HostNameAndPathShrink = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetJsonResponseEnable(v string) *CreateWaitingRoomShrinkRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetLanguage(v string) *CreateWaitingRoomShrinkRequest {
	s.Language = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetName(v string) *CreateWaitingRoomShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetNewUsersPerMinute(v string) *CreateWaitingRoomShrinkRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetQueueAllEnable(v string) *CreateWaitingRoomShrinkRequest {
	s.QueueAllEnable = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetQueuingMethod(v string) *CreateWaitingRoomShrinkRequest {
	s.QueuingMethod = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetQueuingStatusCode(v string) *CreateWaitingRoomShrinkRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetSessionDuration(v string) *CreateWaitingRoomShrinkRequest {
	s.SessionDuration = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetSiteId(v int64) *CreateWaitingRoomShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetTotalActiveUsers(v string) *CreateWaitingRoomShrinkRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetWaitingRoomType(v string) *CreateWaitingRoomShrinkRequest {
	s.WaitingRoomType = &v
	return s
}

type CreateWaitingRoomResponseBody struct {
	// example:
	//
	// 85H66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWaitingRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomResponseBody) SetRequestId(v string) *CreateWaitingRoomResponseBody {
	s.RequestId = &v
	return s
}

type CreateWaitingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWaitingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWaitingRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomResponse) SetHeaders(v map[string]*string) *CreateWaitingRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateWaitingRoomResponse) SetStatusCode(v int32) *CreateWaitingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWaitingRoomResponse) SetBody(v *CreateWaitingRoomResponseBody) *CreateWaitingRoomResponse {
	s.Body = v
	return s
}

type CreateWaitingRoomEventRequest struct {
	// example:
	//
	// Hello%20world!
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// on
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1719849600
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// on
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// example:
	//
	// zhcn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waitingroom_example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// example:
	//
	// on
	PreQueueEnable *string `json:"PreQueueEnable,omitempty" xml:"PreQueueEnable,omitempty"`
	// example:
	//
	// 1719763200
	PreQueueStartTime *string `json:"PreQueueStartTime,omitempty" xml:"PreQueueStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// random
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// example:
	//
	// on
	RandomPreQueueEnable *string `json:"RandomPreQueueEnable,omitempty" xml:"RandomPreQueueEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1719763200
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// example:
	//
	// 6a51d5bc6460887abd1291dc7d4db28b
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s CreateWaitingRoomEventRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomEventRequest) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomEventRequest) SetCustomPageHtml(v string) *CreateWaitingRoomEventRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetDescription(v string) *CreateWaitingRoomEventRequest {
	s.Description = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetDisableSessionRenewalEnable(v string) *CreateWaitingRoomEventRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetEnable(v string) *CreateWaitingRoomEventRequest {
	s.Enable = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetEndTime(v string) *CreateWaitingRoomEventRequest {
	s.EndTime = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetJsonResponseEnable(v string) *CreateWaitingRoomEventRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetLanguage(v string) *CreateWaitingRoomEventRequest {
	s.Language = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetName(v string) *CreateWaitingRoomEventRequest {
	s.Name = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetNewUsersPerMinute(v string) *CreateWaitingRoomEventRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetPreQueueEnable(v string) *CreateWaitingRoomEventRequest {
	s.PreQueueEnable = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetPreQueueStartTime(v string) *CreateWaitingRoomEventRequest {
	s.PreQueueStartTime = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetQueuingMethod(v string) *CreateWaitingRoomEventRequest {
	s.QueuingMethod = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetQueuingStatusCode(v string) *CreateWaitingRoomEventRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetRandomPreQueueEnable(v string) *CreateWaitingRoomEventRequest {
	s.RandomPreQueueEnable = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetSessionDuration(v string) *CreateWaitingRoomEventRequest {
	s.SessionDuration = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetSiteId(v int64) *CreateWaitingRoomEventRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetStartTime(v string) *CreateWaitingRoomEventRequest {
	s.StartTime = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetTotalActiveUsers(v string) *CreateWaitingRoomEventRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetWaitingRoomId(v string) *CreateWaitingRoomEventRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetWaitingRoomType(v string) *CreateWaitingRoomEventRequest {
	s.WaitingRoomType = &v
	return s
}

type CreateWaitingRoomEventResponseBody struct {
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A123425345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWaitingRoomEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomEventResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomEventResponseBody) SetRequestId(v string) *CreateWaitingRoomEventResponseBody {
	s.RequestId = &v
	return s
}

type CreateWaitingRoomEventResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWaitingRoomEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWaitingRoomEventResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomEventResponse) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomEventResponse) SetHeaders(v map[string]*string) *CreateWaitingRoomEventResponse {
	s.Headers = v
	return s
}

func (s *CreateWaitingRoomEventResponse) SetStatusCode(v int32) *CreateWaitingRoomEventResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWaitingRoomEventResponse) SetBody(v *CreateWaitingRoomEventResponseBody) *CreateWaitingRoomEventResponse {
	s.Body = v
	return s
}

type CreateWaitingRoomRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// (http.request.uri.path.file_name eq \\"jpg\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waitingroom_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 25133f536f1b1f6b6091f6a92c614dd4
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
}

func (s CreateWaitingRoomRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRuleRequest) SetRule(v string) *CreateWaitingRoomRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetRuleEnable(v string) *CreateWaitingRoomRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetRuleName(v string) *CreateWaitingRoomRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetSiteId(v int64) *CreateWaitingRoomRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetWaitingRoomId(v string) *CreateWaitingRoomRuleRequest {
	s.WaitingRoomId = &v
	return s
}

type CreateWaitingRoomRuleResponseBody struct {
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWaitingRoomRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRuleResponseBody) SetRequestId(v string) *CreateWaitingRoomRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateWaitingRoomRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWaitingRoomRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWaitingRoomRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWaitingRoomRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRuleResponse) SetHeaders(v map[string]*string) *CreateWaitingRoomRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateWaitingRoomRuleResponse) SetStatusCode(v int32) *CreateWaitingRoomRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWaitingRoomRuleResponse) SetBody(v *CreateWaitingRoomRuleResponseBody) *CreateWaitingRoomRuleResponse {
	s.Body = v
	return s
}

type DeleteCustomScenePolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeleteCustomScenePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomScenePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomScenePolicyRequest) SetPolicyId(v int64) *DeleteCustomScenePolicyRequest {
	s.PolicyId = &v
	return s
}

type DeleteCustomScenePolicyResponseBody struct {
	// example:
	//
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5CC228B4-7A67-4016-9C9F-4A4133494A91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomScenePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomScenePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomScenePolicyResponseBody) SetPolicyId(v int64) *DeleteCustomScenePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *DeleteCustomScenePolicyResponseBody) SetRequestId(v string) *DeleteCustomScenePolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomScenePolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomScenePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomScenePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomScenePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomScenePolicyResponse) SetHeaders(v map[string]*string) *DeleteCustomScenePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomScenePolicyResponse) SetStatusCode(v int32) *DeleteCustomScenePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomScenePolicyResponse) SetBody(v *DeleteCustomScenePolicyResponseBody) *DeleteCustomScenePolicyResponse {
	s.Body = v
	return s
}

type DeleteListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteListRequest) GoString() string {
	return s.String()
}

func (s *DeleteListRequest) SetId(v int64) *DeleteListRequest {
	s.Id = &v
	return s
}

type DeleteListResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteListResponseBody) SetRequestId(v string) *DeleteListResponseBody {
	s.RequestId = &v
	return s
}

type DeleteListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteListResponse) GoString() string {
	return s.String()
}

func (s *DeleteListResponse) SetHeaders(v map[string]*string) *DeleteListResponse {
	s.Headers = v
	return s
}

func (s *DeleteListResponse) SetStatusCode(v int32) *DeleteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteListResponse) SetBody(v *DeleteListResponseBody) *DeleteListResponse {
	s.Body = v
	return s
}

type DeletePageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeletePageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePageRequest) GoString() string {
	return s.String()
}

func (s *DeletePageRequest) SetId(v int64) *DeletePageRequest {
	s.Id = &v
	return s
}

type DeletePageResponseBody struct {
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePageResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePageResponseBody) SetId(v int64) *DeletePageResponseBody {
	s.Id = &v
	return s
}

func (s *DeletePageResponseBody) SetRequestId(v string) *DeletePageResponseBody {
	s.RequestId = &v
	return s
}

type DeletePageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePageResponse) GoString() string {
	return s.String()
}

func (s *DeletePageResponse) SetHeaders(v map[string]*string) *DeletePageResponse {
	s.Headers = v
	return s
}

func (s *DeletePageResponse) SetStatusCode(v int32) *DeletePageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePageResponse) SetBody(v *DeletePageResponseBody) *DeletePageResponse {
	s.Body = v
	return s
}

type DeleteRecordRequest struct {
	// This parameter is required.
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DeleteRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordRequest) SetRecordId(v int64) *DeleteRecordRequest {
	s.RecordId = &v
	return s
}

type DeleteRecordResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordResponseBody) SetRequestId(v string) *DeleteRecordResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRecordResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordResponse) SetHeaders(v map[string]*string) *DeleteRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordResponse) SetStatusCode(v int32) *DeleteRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecordResponse) SetBody(v *DeleteRecordResponseBody) *DeleteRecordResponse {
	s.Body = v
	return s
}

type DeleteScheduledPreloadExecutionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DeleteScheduledPreloadExecution
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteScheduledPreloadExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledPreloadExecutionRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadExecutionRequest) SetId(v string) *DeleteScheduledPreloadExecutionRequest {
	s.Id = &v
	return s
}

type DeleteScheduledPreloadExecutionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduledPreloadExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledPreloadExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadExecutionResponseBody) SetRequestId(v string) *DeleteScheduledPreloadExecutionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScheduledPreloadExecutionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduledPreloadExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduledPreloadExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledPreloadExecutionResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadExecutionResponse) SetHeaders(v map[string]*string) *DeleteScheduledPreloadExecutionResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledPreloadExecutionResponse) SetStatusCode(v int32) *DeleteScheduledPreloadExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledPreloadExecutionResponse) SetBody(v *DeleteScheduledPreloadExecutionResponseBody) *DeleteScheduledPreloadExecutionResponse {
	s.Body = v
	return s
}

type DeleteScheduledPreloadJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DeleteScheduledPreloadJob
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteScheduledPreloadJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledPreloadJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadJobRequest) SetId(v string) *DeleteScheduledPreloadJobRequest {
	s.Id = &v
	return s
}

type DeleteScheduledPreloadJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduledPreloadJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledPreloadJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadJobResponseBody) SetRequestId(v string) *DeleteScheduledPreloadJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScheduledPreloadJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduledPreloadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduledPreloadJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledPreloadJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadJobResponse) SetHeaders(v map[string]*string) *DeleteScheduledPreloadJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledPreloadJobResponse) SetStatusCode(v int32) *DeleteScheduledPreloadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledPreloadJobResponse) SetBody(v *DeleteScheduledPreloadJobResponseBody) *DeleteScheduledPreloadJobResponse {
	s.Body = v
	return s
}

type DeleteSiteRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSiteRequest) GoString() string {
	return s.String()
}

func (s *DeleteSiteRequest) SetOwnerId(v int64) *DeleteSiteRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSiteRequest) SetSecurityToken(v string) *DeleteSiteRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteSiteRequest) SetSiteId(v int64) *DeleteSiteRequest {
	s.SiteId = &v
	return s
}

type DeleteSiteResponseBody struct {
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSiteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSiteResponseBody) SetRequestId(v string) *DeleteSiteResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSiteResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSiteResponse) GoString() string {
	return s.String()
}

func (s *DeleteSiteResponse) SetHeaders(v map[string]*string) *DeleteSiteResponse {
	s.Headers = v
	return s
}

func (s *DeleteSiteResponse) SetStatusCode(v int32) *DeleteSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSiteResponse) SetBody(v *DeleteSiteResponseBody) *DeleteSiteResponse {
	s.Body = v
	return s
}

type DeleteSiteDeliveryTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456******
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DeleteSiteDeliveryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSiteDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteSiteDeliveryTaskRequest) SetSiteId(v int64) *DeleteSiteDeliveryTaskRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteSiteDeliveryTaskRequest) SetTaskName(v string) *DeleteSiteDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

type DeleteSiteDeliveryTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSiteDeliveryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSiteDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSiteDeliveryTaskResponseBody) SetRequestId(v string) *DeleteSiteDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSiteDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSiteDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSiteDeliveryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSiteDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteSiteDeliveryTaskResponse) SetHeaders(v map[string]*string) *DeleteSiteDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteSiteDeliveryTaskResponse) SetStatusCode(v int32) *DeleteSiteDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSiteDeliveryTaskResponse) SetBody(v *DeleteSiteDeliveryTaskResponseBody) *DeleteSiteDeliveryTaskResponse {
	s.Body = v
	return s
}

type DeleteUserDeliveryTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DeleteUserDeliveryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserDeliveryTaskRequest) SetTaskName(v string) *DeleteUserDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

type DeleteUserDeliveryTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 952ea16b-1f05-4a76-bb32-420282d8aeb9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserDeliveryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserDeliveryTaskResponseBody) SetRequestId(v string) *DeleteUserDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserDeliveryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserDeliveryTaskResponse) SetHeaders(v map[string]*string) *DeleteUserDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserDeliveryTaskResponse) SetStatusCode(v int32) *DeleteUserDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserDeliveryTaskResponse) SetBody(v *DeleteUserDeliveryTaskResponseBody) *DeleteUserDeliveryTaskResponse {
	s.Body = v
	return s
}

type DeleteWafRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s DeleteWafRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWafRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWafRuleRequest) SetId(v int64) *DeleteWafRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteWafRuleRequest) SetSiteId(v int64) *DeleteWafRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteWafRuleRequest) SetSiteVersion(v int32) *DeleteWafRuleRequest {
	s.SiteVersion = &v
	return s
}

type DeleteWafRuleResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWafRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWafRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWafRuleResponseBody) SetRequestId(v string) *DeleteWafRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWafRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWafRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWafRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWafRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWafRuleResponse) SetHeaders(v map[string]*string) *DeleteWafRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWafRuleResponse) SetStatusCode(v int32) *DeleteWafRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWafRuleResponse) SetBody(v *DeleteWafRuleResponseBody) *DeleteWafRuleResponse {
	s.Body = v
	return s
}

type DeleteWafRulesetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s DeleteWafRulesetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *DeleteWafRulesetRequest) SetId(v int64) *DeleteWafRulesetRequest {
	s.Id = &v
	return s
}

func (s *DeleteWafRulesetRequest) SetSiteId(v int64) *DeleteWafRulesetRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteWafRulesetRequest) SetSiteVersion(v int32) *DeleteWafRulesetRequest {
	s.SiteVersion = &v
	return s
}

type DeleteWafRulesetResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWafRulesetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWafRulesetResponseBody) SetRequestId(v string) *DeleteWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWafRulesetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWafRulesetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *DeleteWafRulesetResponse) SetHeaders(v map[string]*string) *DeleteWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *DeleteWafRulesetResponse) SetStatusCode(v int32) *DeleteWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWafRulesetResponse) SetBody(v *DeleteWafRulesetResponseBody) *DeleteWafRulesetResponse {
	s.Body = v
	return s
}

type DeleteWaitingRoomRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 25133f536f1b1f6b6091f6a92c614dd4
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
}

func (s DeleteWaitingRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaitingRoomRequest) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomRequest) SetSiteId(v int64) *DeleteWaitingRoomRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteWaitingRoomRequest) SetWaitingRoomId(v string) *DeleteWaitingRoomRequest {
	s.WaitingRoomId = &v
	return s
}

type DeleteWaitingRoomResponseBody struct {
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A123425345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWaitingRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaitingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomResponseBody) SetRequestId(v string) *DeleteWaitingRoomResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWaitingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWaitingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWaitingRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaitingRoomResponse) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomResponse) SetHeaders(v map[string]*string) *DeleteWaitingRoomResponse {
	s.Headers = v
	return s
}

func (s *DeleteWaitingRoomResponse) SetStatusCode(v int32) *DeleteWaitingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWaitingRoomResponse) SetBody(v *DeleteWaitingRoomResponseBody) *DeleteWaitingRoomResponse {
	s.Body = v
	return s
}

type DeleteWaitingRoomEventRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 302909890***
	WaitingRoomEventId *int64 `json:"WaitingRoomEventId,omitempty" xml:"WaitingRoomEventId,omitempty"`
}

func (s DeleteWaitingRoomEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaitingRoomEventRequest) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomEventRequest) SetSiteId(v int64) *DeleteWaitingRoomEventRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteWaitingRoomEventRequest) SetWaitingRoomEventId(v int64) *DeleteWaitingRoomEventRequest {
	s.WaitingRoomEventId = &v
	return s
}

type DeleteWaitingRoomEventResponseBody struct {
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWaitingRoomEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaitingRoomEventResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomEventResponseBody) SetRequestId(v string) *DeleteWaitingRoomEventResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWaitingRoomEventResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWaitingRoomEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWaitingRoomEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaitingRoomEventResponse) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomEventResponse) SetHeaders(v map[string]*string) *DeleteWaitingRoomEventResponse {
	s.Headers = v
	return s
}

func (s *DeleteWaitingRoomEventResponse) SetStatusCode(v int32) *DeleteWaitingRoomEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWaitingRoomEventResponse) SetBody(v *DeleteWaitingRoomEventResponseBody) *DeleteWaitingRoomEventResponse {
	s.Body = v
	return s
}

type DeleteWaitingRoomRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 3672886****
	WaitingRoomRuleId *int64 `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s DeleteWaitingRoomRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaitingRoomRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomRuleRequest) SetSiteId(v int64) *DeleteWaitingRoomRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteWaitingRoomRuleRequest) SetWaitingRoomRuleId(v int64) *DeleteWaitingRoomRuleRequest {
	s.WaitingRoomRuleId = &v
	return s
}

type DeleteWaitingRoomRuleResponseBody struct {
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWaitingRoomRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaitingRoomRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomRuleResponseBody) SetRequestId(v string) *DeleteWaitingRoomRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWaitingRoomRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWaitingRoomRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWaitingRoomRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaitingRoomRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomRuleResponse) SetHeaders(v map[string]*string) *DeleteWaitingRoomRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWaitingRoomRuleResponse) SetStatusCode(v int32) *DeleteWaitingRoomRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWaitingRoomRuleResponse) SetBody(v *DeleteWaitingRoomRuleResponseBody) *DeleteWaitingRoomRuleResponse {
	s.Body = v
	return s
}

type DescribeCustomScenePoliciesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1234****
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DescribeCustomScenePoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomScenePoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomScenePoliciesRequest) SetPageNumber(v int32) *DescribeCustomScenePoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomScenePoliciesRequest) SetPageSize(v int32) *DescribeCustomScenePoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomScenePoliciesRequest) SetPolicyId(v int64) *DescribeCustomScenePoliciesRequest {
	s.PolicyId = &v
	return s
}

type DescribeCustomScenePoliciesResponseBody struct {
	DataModule []*DescribeCustomScenePoliciesResponseBodyDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 85H66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCustomScenePoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomScenePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomScenePoliciesResponseBody) SetDataModule(v []*DescribeCustomScenePoliciesResponseBodyDataModule) *DescribeCustomScenePoliciesResponseBody {
	s.DataModule = v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) SetPageNumber(v int32) *DescribeCustomScenePoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) SetPageSize(v int32) *DescribeCustomScenePoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) SetQuota(v int32) *DescribeCustomScenePoliciesResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) SetRequestId(v string) *DescribeCustomScenePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) SetTotalCount(v int32) *DescribeCustomScenePoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCustomScenePoliciesResponseBodyDataModule struct {
	// example:
	//
	// 2023-03-06T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// test
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Objects []*string `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	// example:
	//
	// 1234****
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 2023-03-04T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Expired
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s DescribeCustomScenePoliciesResponseBodyDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomScenePoliciesResponseBodyDataModule) GoString() string {
	return s.String()
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetEndTime(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.EndTime = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetName(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.Name = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetObjects(v []*string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.Objects = v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetPolicyId(v int64) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.PolicyId = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetStartTime(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetStatus(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.Status = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetTemplate(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.Template = &v
	return s
}

type DescribeCustomScenePoliciesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomScenePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomScenePoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomScenePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomScenePoliciesResponse) SetHeaders(v map[string]*string) *DescribeCustomScenePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomScenePoliciesResponse) SetStatusCode(v int32) *DescribeCustomScenePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponse) SetBody(v *DescribeCustomScenePoliciesResponseBody) *DescribeCustomScenePoliciesResponse {
	s.Body = v
	return s
}

type DescribeDDoSAllEventListRequest struct {
	// example:
	//
	// 2023-02-22T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// web-cc
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// A short description of struct
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-12T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSAllEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSAllEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSAllEventListRequest) SetEndTime(v string) *DescribeDDoSAllEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) SetEventType(v string) *DescribeDDoSAllEventListRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) SetPageNumber(v int32) *DescribeDDoSAllEventListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) SetPageSize(v int32) *DescribeDDoSAllEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) SetSiteId(v int64) *DescribeDDoSAllEventListRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) SetStartTime(v string) *DescribeDDoSAllEventListRequest {
	s.StartTime = &v
	return s
}

type DescribeDDoSAllEventListResponseBody struct {
	DataList []*DescribeDDoSAllEventListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D73A4243-CFBD-5110-876F-09237E77ECBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDDoSAllEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSAllEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSAllEventListResponseBody) SetDataList(v []*DescribeDDoSAllEventListResponseBodyDataList) *DescribeDDoSAllEventListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) SetPageNumber(v int32) *DescribeDDoSAllEventListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) SetPageSize(v int32) *DescribeDDoSAllEventListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) SetRequestId(v string) *DescribeDDoSAllEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) SetSiteId(v int64) *DescribeDDoSAllEventListResponseBody {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) SetTotalCount(v int32) *DescribeDDoSAllEventListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDDoSAllEventListResponseBodyDataList struct {
	// example:
	//
	// 800
	Bps *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// example:
	//
	// 50
	Cps *int64 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// example:
	//
	// 2023-02-12T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// web-cc_1
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// web-cc
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// 12000
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// example:
	//
	// 7692
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// example:
	//
	// 2023-02-12T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// example.com
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// example:
	//
	// 000000000155****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s DescribeDDoSAllEventListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSAllEventListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetBps(v int64) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Bps = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetCps(v int64) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Cps = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetEndTime(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetEventId(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.EventId = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetEventType(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.EventType = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetPps(v int64) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Pps = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetQps(v int64) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Qps = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetStartTime(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetTarget(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Target = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetTargetId(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.TargetId = &v
	return s
}

type DescribeDDoSAllEventListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSAllEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSAllEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSAllEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSAllEventListResponse) SetHeaders(v map[string]*string) *DescribeDDoSAllEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSAllEventListResponse) SetStatusCode(v int32) *DescribeDDoSAllEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSAllEventListResponse) SetBody(v *DescribeDDoSAllEventListResponseBody) *DescribeDDoSAllEventListResponse {
	s.Body = v
	return s
}

type DescribeHttpDDoSAttackIntelligentProtectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSAttackIntelligentProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpDDoSAttackIntelligentProtectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionRequest) SetSiteId(v int64) *DescribeHttpDDoSAttackIntelligentProtectionRequest {
	s.SiteId = &v
	return s
}

type DescribeHttpDDoSAttackIntelligentProtectionResponseBody struct {
	// example:
	//
	// defense
	AiMode *string `json:"AiMode,omitempty" xml:"AiMode,omitempty"`
	// example:
	//
	// level60
	AiTemplate *string `json:"AiTemplate,omitempty" xml:"AiTemplate,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CB1A380B-09F0-41BB-3C82-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSAttackIntelligentProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpDDoSAttackIntelligentProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) SetAiMode(v string) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody {
	s.AiMode = &v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) SetAiTemplate(v string) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody {
	s.AiTemplate = &v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) SetRequestId(v string) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) SetSiteId(v int64) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody {
	s.SiteId = &v
	return s
}

type DescribeHttpDDoSAttackIntelligentProtectionResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHttpDDoSAttackIntelligentProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHttpDDoSAttackIntelligentProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpDDoSAttackIntelligentProtectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponse) SetHeaders(v map[string]*string) *DescribeHttpDDoSAttackIntelligentProtectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponse) SetStatusCode(v int32) *DescribeHttpDDoSAttackIntelligentProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponse) SetBody(v *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) *DescribeHttpDDoSAttackIntelligentProtectionResponse {
	s.Body = v
	return s
}

type DescribeHttpDDoSAttackProtectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSAttackProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpDDoSAttackProtectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackProtectionRequest) SetSiteId(v int64) *DescribeHttpDDoSAttackProtectionRequest {
	s.SiteId = &v
	return s
}

type DescribeHttpDDoSAttackProtectionResponseBody struct {
	// example:
	//
	// default
	GlobalMode *string `json:"GlobalMode,omitempty" xml:"GlobalMode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 35C66C7B-671H-4297-9187-2C4477247A78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSAttackProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpDDoSAttackProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackProtectionResponseBody) SetGlobalMode(v string) *DescribeHttpDDoSAttackProtectionResponseBody {
	s.GlobalMode = &v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionResponseBody) SetRequestId(v string) *DescribeHttpDDoSAttackProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionResponseBody) SetSiteId(v int64) *DescribeHttpDDoSAttackProtectionResponseBody {
	s.SiteId = &v
	return s
}

type DescribeHttpDDoSAttackProtectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHttpDDoSAttackProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHttpDDoSAttackProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpDDoSAttackProtectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackProtectionResponse) SetHeaders(v map[string]*string) *DescribeHttpDDoSAttackProtectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionResponse) SetStatusCode(v int32) *DescribeHttpDDoSAttackProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionResponse) SetBody(v *DescribeHttpDDoSAttackProtectionResponseBody) *DescribeHttpDDoSAttackProtectionResponse {
	s.Body = v
	return s
}

type DescribeIPRangeListResponseBody struct {
	Content []*DescribeIPRangeListResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIPRangeListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIPRangeListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIPRangeListResponseBody) SetContent(v []*DescribeIPRangeListResponseBodyContent) *DescribeIPRangeListResponseBody {
	s.Content = v
	return s
}

func (s *DescribeIPRangeListResponseBody) SetRequestId(v string) *DescribeIPRangeListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIPRangeListResponseBodyContent struct {
	// example:
	//
	// 172.16.0.0/12
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// IPv4
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
}

func (s DescribeIPRangeListResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeIPRangeListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeIPRangeListResponseBodyContent) SetCidr(v string) *DescribeIPRangeListResponseBodyContent {
	s.Cidr = &v
	return s
}

func (s *DescribeIPRangeListResponseBodyContent) SetIpType(v string) *DescribeIPRangeListResponseBodyContent {
	s.IpType = &v
	return s
}

type DescribeIPRangeListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIPRangeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIPRangeListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIPRangeListResponse) GoString() string {
	return s.String()
}

func (s *DescribeIPRangeListResponse) SetHeaders(v map[string]*string) *DescribeIPRangeListResponse {
	s.Headers = v
	return s
}

func (s *DescribeIPRangeListResponse) SetStatusCode(v int32) *DescribeIPRangeListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIPRangeListResponse) SetBody(v *DescribeIPRangeListResponseBody) *DescribeIPRangeListResponse {
	s.Body = v
	return s
}

type DescribePreloadTasksRequest struct {
	// example:
	//
	// http://a.com/1.jpg?b=2
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2023-03-23T06:23:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 2023-03-22T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePreloadTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePreloadTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribePreloadTasksRequest) SetContent(v string) *DescribePreloadTasksRequest {
	s.Content = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetEndTime(v string) *DescribePreloadTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetPageNumber(v int32) *DescribePreloadTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetPageSize(v int32) *DescribePreloadTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetSiteId(v int64) *DescribePreloadTasksRequest {
	s.SiteId = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetStartTime(v string) *DescribePreloadTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetStatus(v string) *DescribePreloadTasksRequest {
	s.Status = &v
	return s
}

type DescribePreloadTasksResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*DescribePreloadTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 83
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePreloadTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePreloadTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreloadTasksResponseBody) SetPageNumber(v int64) *DescribePreloadTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePreloadTasksResponseBody) SetPageSize(v int64) *DescribePreloadTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePreloadTasksResponseBody) SetRequestId(v string) *DescribePreloadTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreloadTasksResponseBody) SetTasks(v []*DescribePreloadTasksResponseBodyTasks) *DescribePreloadTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribePreloadTasksResponseBody) SetTotalCount(v int64) *DescribePreloadTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePreloadTasksResponseBodyTasks struct {
	// example:
	//
	// http://a.com/1.jpg?b=2
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2023-03-28 14:28:57
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// Internal Error
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 100%
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1597854579687428
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribePreloadTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribePreloadTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribePreloadTasksResponseBodyTasks) SetContent(v string) *DescribePreloadTasksResponseBodyTasks {
	s.Content = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) SetCreateTime(v string) *DescribePreloadTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) SetDescription(v string) *DescribePreloadTasksResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) SetProcess(v string) *DescribePreloadTasksResponseBodyTasks {
	s.Process = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) SetStatus(v string) *DescribePreloadTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribePreloadTasksResponseBodyTasks) SetTaskId(v string) *DescribePreloadTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

type DescribePreloadTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePreloadTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePreloadTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePreloadTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribePreloadTasksResponse) SetHeaders(v map[string]*string) *DescribePreloadTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribePreloadTasksResponse) SetStatusCode(v int32) *DescribePreloadTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePreloadTasksResponse) SetBody(v *DescribePreloadTasksResponseBody) *DescribePreloadTasksResponse {
	s.Body = v
	return s
}

type DescribePurgeTasksRequest struct {
	// example:
	//
	// http://a.com/1.jpg?b=1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2022-11-18T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 2022-11-16T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePurgeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePurgeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribePurgeTasksRequest) SetContent(v string) *DescribePurgeTasksRequest {
	s.Content = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetEndTime(v string) *DescribePurgeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetPageNumber(v int32) *DescribePurgeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetPageSize(v int32) *DescribePurgeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetSiteId(v int64) *DescribePurgeTasksRequest {
	s.SiteId = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetStartTime(v string) *DescribePurgeTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetStatus(v string) *DescribePurgeTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetType(v string) *DescribePurgeTasksRequest {
	s.Type = &v
	return s
}

type DescribePurgeTasksResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A123425345
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*DescribePurgeTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePurgeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePurgeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurgeTasksResponseBody) SetPageNumber(v int64) *DescribePurgeTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePurgeTasksResponseBody) SetPageSize(v int64) *DescribePurgeTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePurgeTasksResponseBody) SetRequestId(v string) *DescribePurgeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurgeTasksResponseBody) SetTasks(v []*DescribePurgeTasksResponseBodyTasks) *DescribePurgeTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribePurgeTasksResponseBody) SetTotalCount(v int64) *DescribePurgeTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePurgeTasksResponseBodyTasks struct {
	// example:
	//
	// http://a.com/1.jpg?b=1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2023-07-26T01:56:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// Internal Error
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 100%
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 16346513304
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePurgeTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribePurgeTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribePurgeTasksResponseBodyTasks) SetContent(v string) *DescribePurgeTasksResponseBodyTasks {
	s.Content = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetCreateTime(v string) *DescribePurgeTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetDescription(v string) *DescribePurgeTasksResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetProcess(v string) *DescribePurgeTasksResponseBodyTasks {
	s.Process = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetStatus(v string) *DescribePurgeTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetTaskId(v string) *DescribePurgeTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribePurgeTasksResponseBodyTasks) SetType(v string) *DescribePurgeTasksResponseBodyTasks {
	s.Type = &v
	return s
}

type DescribePurgeTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurgeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurgeTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePurgeTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribePurgeTasksResponse) SetHeaders(v map[string]*string) *DescribePurgeTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribePurgeTasksResponse) SetStatusCode(v int32) *DescribePurgeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurgeTasksResponse) SetBody(v *DescribePurgeTasksResponseBody) *DescribePurgeTasksResponse {
	s.Body = v
	return s
}

type DisableCustomScenePolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DisableCustomScenePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableCustomScenePolicyRequest) GoString() string {
	return s.String()
}

func (s *DisableCustomScenePolicyRequest) SetPolicyId(v int64) *DisableCustomScenePolicyRequest {
	s.PolicyId = &v
	return s
}

type DisableCustomScenePolicyResponseBody struct {
	// example:
	//
	// 100001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCustomScenePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableCustomScenePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCustomScenePolicyResponseBody) SetPolicyId(v int64) *DisableCustomScenePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *DisableCustomScenePolicyResponseBody) SetRequestId(v string) *DisableCustomScenePolicyResponseBody {
	s.RequestId = &v
	return s
}

type DisableCustomScenePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCustomScenePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCustomScenePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableCustomScenePolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableCustomScenePolicyResponse) SetHeaders(v map[string]*string) *DisableCustomScenePolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableCustomScenePolicyResponse) SetStatusCode(v int32) *DisableCustomScenePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCustomScenePolicyResponse) SetBody(v *DisableCustomScenePolicyResponseBody) *DisableCustomScenePolicyResponse {
	s.Body = v
	return s
}

type EditSiteWafSettingsRequest struct {
	Settings *WafSiteSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s EditSiteWafSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s EditSiteWafSettingsRequest) GoString() string {
	return s.String()
}

func (s *EditSiteWafSettingsRequest) SetSettings(v *WafSiteSettings) *EditSiteWafSettingsRequest {
	s.Settings = v
	return s
}

func (s *EditSiteWafSettingsRequest) SetSiteId(v int64) *EditSiteWafSettingsRequest {
	s.SiteId = &v
	return s
}

func (s *EditSiteWafSettingsRequest) SetSiteVersion(v int32) *EditSiteWafSettingsRequest {
	s.SiteVersion = &v
	return s
}

type EditSiteWafSettingsShrinkRequest struct {
	SettingsShrink *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s EditSiteWafSettingsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EditSiteWafSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *EditSiteWafSettingsShrinkRequest) SetSettingsShrink(v string) *EditSiteWafSettingsShrinkRequest {
	s.SettingsShrink = &v
	return s
}

func (s *EditSiteWafSettingsShrinkRequest) SetSiteId(v int64) *EditSiteWafSettingsShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *EditSiteWafSettingsShrinkRequest) SetSiteVersion(v int32) *EditSiteWafSettingsShrinkRequest {
	s.SiteVersion = &v
	return s
}

type EditSiteWafSettingsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditSiteWafSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditSiteWafSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *EditSiteWafSettingsResponseBody) SetRequestId(v string) *EditSiteWafSettingsResponseBody {
	s.RequestId = &v
	return s
}

type EditSiteWafSettingsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditSiteWafSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditSiteWafSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s EditSiteWafSettingsResponse) GoString() string {
	return s.String()
}

func (s *EditSiteWafSettingsResponse) SetHeaders(v map[string]*string) *EditSiteWafSettingsResponse {
	s.Headers = v
	return s
}

func (s *EditSiteWafSettingsResponse) SetStatusCode(v int32) *EditSiteWafSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *EditSiteWafSettingsResponse) SetBody(v *EditSiteWafSettingsResponseBody) *EditSiteWafSettingsResponse {
	s.Body = v
	return s
}

type EnableCustomScenePolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s EnableCustomScenePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableCustomScenePolicyRequest) GoString() string {
	return s.String()
}

func (s *EnableCustomScenePolicyRequest) SetPolicyId(v int64) *EnableCustomScenePolicyRequest {
	s.PolicyId = &v
	return s
}

type EnableCustomScenePolicyResponseBody struct {
	// example:
	//
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 65C66B7B-671A-8297-9187-2R5477247B76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCustomScenePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableCustomScenePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *EnableCustomScenePolicyResponseBody) SetPolicyId(v int64) *EnableCustomScenePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *EnableCustomScenePolicyResponseBody) SetRequestId(v string) *EnableCustomScenePolicyResponseBody {
	s.RequestId = &v
	return s
}

type EnableCustomScenePolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableCustomScenePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCustomScenePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableCustomScenePolicyResponse) GoString() string {
	return s.String()
}

func (s *EnableCustomScenePolicyResponse) SetHeaders(v map[string]*string) *EnableCustomScenePolicyResponse {
	s.Headers = v
	return s
}

func (s *EnableCustomScenePolicyResponse) SetStatusCode(v int32) *EnableCustomScenePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableCustomScenePolicyResponse) SetBody(v *EnableCustomScenePolicyResponseBody) *EnableCustomScenePolicyResponse {
	s.Body = v
	return s
}

type ExportRecordsRequest struct {
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ExportRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportRecordsRequest) GoString() string {
	return s.String()
}

func (s *ExportRecordsRequest) SetSiteId(v int64) *ExportRecordsRequest {
	s.SiteId = &v
	return s
}

type ExportRecordsResponseBody struct {
	// example:
	//
	// ;; site:example.com.\\n;; Exported:2024-01-24 15:54:35\\n\\n;; A Records\\na1.example.com. 30 IN A 1.1.1.1 direct\\na2.example.com. 30 IN A 1.1.1.1 direct\\na3.example.com. 30 IN A 1.1.1.1 direct\\n
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C69B5894-D1BA-592C-95D0-DADBE7AEAC63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ExportRecordsResponseBody) SetContent(v string) *ExportRecordsResponseBody {
	s.Content = &v
	return s
}

func (s *ExportRecordsResponseBody) SetRequestId(v string) *ExportRecordsResponseBody {
	s.RequestId = &v
	return s
}

type ExportRecordsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportRecordsResponse) GoString() string {
	return s.String()
}

func (s *ExportRecordsResponse) SetHeaders(v map[string]*string) *ExportRecordsResponse {
	s.Headers = v
	return s
}

func (s *ExportRecordsResponse) SetStatusCode(v int32) *ExportRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportRecordsResponse) SetBody(v *ExportRecordsResponseBody) *ExportRecordsResponse {
	s.Body = v
	return s
}

type GetKvNamespaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetKvNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKvNamespaceRequest) GoString() string {
	return s.String()
}

func (s *GetKvNamespaceRequest) SetNamespace(v string) *GetKvNamespaceRequest {
	s.Namespace = &v
	return s
}

type GetKvNamespaceResponseBody struct {
	// example:
	//
	// 1073741824
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// 1 GB
	CapacityString *string `json:"CapacityString,omitempty" xml:"CapacityString,omitempty"`
	// example:
	//
	// 10048576
	CapacityUsed *int64 `json:"CapacityUsed,omitempty" xml:"CapacityUsed,omitempty"`
	// example:
	//
	// 100 MB
	CapacityUsedString *string `json:"CapacityUsedString,omitempty" xml:"CapacityUsedString,omitempty"`
	// example:
	//
	// this is a test namespace.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 643355322374688768
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetKvNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKvNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetKvNamespaceResponseBody) SetCapacity(v int64) *GetKvNamespaceResponseBody {
	s.Capacity = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetCapacityString(v string) *GetKvNamespaceResponseBody {
	s.CapacityString = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetCapacityUsed(v int64) *GetKvNamespaceResponseBody {
	s.CapacityUsed = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetCapacityUsedString(v string) *GetKvNamespaceResponseBody {
	s.CapacityUsedString = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetDescription(v string) *GetKvNamespaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetNamespace(v string) *GetKvNamespaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetNamespaceId(v string) *GetKvNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetRequestId(v string) *GetKvNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetStatus(v string) *GetKvNamespaceResponseBody {
	s.Status = &v
	return s
}

type GetKvNamespaceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKvNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKvNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKvNamespaceResponse) GoString() string {
	return s.String()
}

func (s *GetKvNamespaceResponse) SetHeaders(v map[string]*string) *GetKvNamespaceResponse {
	s.Headers = v
	return s
}

func (s *GetKvNamespaceResponse) SetStatusCode(v int32) *GetKvNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKvNamespaceResponse) SetBody(v *GetKvNamespaceResponseBody) *GetKvNamespaceResponse {
	s.Body = v
	return s
}

type GetListRequest struct {
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetListRequest) GoString() string {
	return s.String()
}

func (s *GetListRequest) SetId(v int64) *GetListRequest {
	s.Id = &v
	return s
}

type GetListResponseBody struct {
	// 自定义响应页面描述
	//
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 自定义响应页面ID
	//
	// example:
	//
	// 40000001
	Id    *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// ip
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// 自定义响应页面名称
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetListResponseBody) GoString() string {
	return s.String()
}

func (s *GetListResponseBody) SetDescription(v string) *GetListResponseBody {
	s.Description = &v
	return s
}

func (s *GetListResponseBody) SetId(v int64) *GetListResponseBody {
	s.Id = &v
	return s
}

func (s *GetListResponseBody) SetItems(v []*string) *GetListResponseBody {
	s.Items = v
	return s
}

func (s *GetListResponseBody) SetKind(v string) *GetListResponseBody {
	s.Kind = &v
	return s
}

func (s *GetListResponseBody) SetName(v string) *GetListResponseBody {
	s.Name = &v
	return s
}

func (s *GetListResponseBody) SetRequestId(v string) *GetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListResponseBody) SetUpdateTime(v string) *GetListResponseBody {
	s.UpdateTime = &v
	return s
}

type GetListResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetListResponse) GoString() string {
	return s.String()
}

func (s *GetListResponse) SetHeaders(v map[string]*string) *GetListResponse {
	s.Headers = v
	return s
}

func (s *GetListResponse) SetStatusCode(v int32) *GetListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListResponse) SetBody(v *GetListResponseBody) *GetListResponse {
	s.Body = v
	return s
}

type GetPageRequest struct {
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPageRequest) GoString() string {
	return s.String()
}

func (s *GetPageRequest) SetId(v int64) *GetPageRequest {
	s.Id = &v
	return s
}

type GetPageResponseBody struct {
	// 自定义响应页面内容BASE64编码
	//
	// This parameter is required.
	//
	// example:
	//
	// PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 自定义响应页面内容类型
	//
	// This parameter is required.
	//
	// example:
	//
	// text/html
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// 自定义响应页面描述
	//
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 自定义响应页面ID
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// custom
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// 自定义响应页面名称
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetPageResponseBody) SetContent(v string) *GetPageResponseBody {
	s.Content = &v
	return s
}

func (s *GetPageResponseBody) SetContentType(v string) *GetPageResponseBody {
	s.ContentType = &v
	return s
}

func (s *GetPageResponseBody) SetDescription(v string) *GetPageResponseBody {
	s.Description = &v
	return s
}

func (s *GetPageResponseBody) SetId(v int64) *GetPageResponseBody {
	s.Id = &v
	return s
}

func (s *GetPageResponseBody) SetKind(v string) *GetPageResponseBody {
	s.Kind = &v
	return s
}

func (s *GetPageResponseBody) SetName(v string) *GetPageResponseBody {
	s.Name = &v
	return s
}

func (s *GetPageResponseBody) SetRequestId(v string) *GetPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPageResponseBody) SetUpdateTime(v string) *GetPageResponseBody {
	s.UpdateTime = &v
	return s
}

type GetPageResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPageResponse) GoString() string {
	return s.String()
}

func (s *GetPageResponse) SetHeaders(v map[string]*string) *GetPageResponse {
	s.Headers = v
	return s
}

func (s *GetPageResponse) SetStatusCode(v int32) *GetPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPageResponse) SetBody(v *GetPageResponseBody) *GetPageResponse {
	s.Body = v
	return s
}

type GetPurgeQuotaRequest struct {
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPurgeQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPurgeQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetPurgeQuotaRequest) SetSiteId(v int64) *GetPurgeQuotaRequest {
	s.SiteId = &v
	return s
}

func (s *GetPurgeQuotaRequest) SetType(v string) *GetPurgeQuotaRequest {
	s.Type = &v
	return s
}

type GetPurgeQuotaResponseBody struct {
	// example:
	//
	// 100000
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s GetPurgeQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPurgeQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetPurgeQuotaResponseBody) SetQuota(v string) *GetPurgeQuotaResponseBody {
	s.Quota = &v
	return s
}

func (s *GetPurgeQuotaResponseBody) SetRequestId(v string) *GetPurgeQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPurgeQuotaResponseBody) SetUsage(v string) *GetPurgeQuotaResponseBody {
	s.Usage = &v
	return s
}

type GetPurgeQuotaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPurgeQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPurgeQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPurgeQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetPurgeQuotaResponse) SetHeaders(v map[string]*string) *GetPurgeQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetPurgeQuotaResponse) SetStatusCode(v int32) *GetPurgeQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPurgeQuotaResponse) SetBody(v *GetPurgeQuotaResponseBody) *GetPurgeQuotaResponse {
	s.Body = v
	return s
}

type GetRealtimeDeliveryFieldRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_er
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s GetRealtimeDeliveryFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeDeliveryFieldRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeDeliveryFieldRequest) SetBusinessType(v string) *GetRealtimeDeliveryFieldRequest {
	s.BusinessType = &v
	return s
}

type GetRealtimeDeliveryFieldResponseBody struct {
	FieldContent map[string]*FieldContentValue `json:"FieldContent,omitempty" xml:"FieldContent,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRealtimeDeliveryFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeDeliveryFieldResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeDeliveryFieldResponseBody) SetFieldContent(v map[string]*FieldContentValue) *GetRealtimeDeliveryFieldResponseBody {
	s.FieldContent = v
	return s
}

func (s *GetRealtimeDeliveryFieldResponseBody) SetRequestId(v string) *GetRealtimeDeliveryFieldResponseBody {
	s.RequestId = &v
	return s
}

type GetRealtimeDeliveryFieldResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealtimeDeliveryFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealtimeDeliveryFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeDeliveryFieldResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeDeliveryFieldResponse) SetHeaders(v map[string]*string) *GetRealtimeDeliveryFieldResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeDeliveryFieldResponse) SetStatusCode(v int32) *GetRealtimeDeliveryFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealtimeDeliveryFieldResponse) SetBody(v *GetRealtimeDeliveryFieldResponseBody) *GetRealtimeDeliveryFieldResponse {
	s.Body = v
	return s
}

type GetRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s GetRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecordRequest) GoString() string {
	return s.String()
}

func (s *GetRecordRequest) SetRecordId(v int64) *GetRecordRequest {
	s.RecordId = &v
	return s
}

type GetRecordResponseBody struct {
	RecordModel *GetRecordResponseBodyRecordModel `json:"RecordModel,omitempty" xml:"RecordModel,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBody) SetRecordModel(v *GetRecordResponseBodyRecordModel) *GetRecordResponseBody {
	s.RecordModel = v
	return s
}

func (s *GetRecordResponseBody) SetRequestId(v string) *GetRecordResponseBody {
	s.RequestId = &v
	return s
}

type GetRecordResponseBodyRecordModel struct {
	AuthConf *GetRecordResponseBodyRecordModelAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// example:
	//
	// image_video
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 2023-03-10T13:30:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// {"value":"1.1.1.1"}
	Data       *GetRecordResponseBodyRecordModelData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HostPolicy *string                               `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// example:
	//
	// a.example.com.cnamezone.com
	RecordCname *string `json:"RecordCname,omitempty" xml:"RecordCname,omitempty"`
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// example:
	//
	// OSS
	RecordSourceType *string `json:"RecordSourceType,omitempty" xml:"RecordSourceType,omitempty"`
	// example:
	//
	// A/AAAA
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// example:
	//
	// 20
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// example:
	//
	// 2023-01-27T02:26:22Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetRecordResponseBodyRecordModel) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResponseBodyRecordModel) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBodyRecordModel) SetAuthConf(v *GetRecordResponseBodyRecordModelAuthConf) *GetRecordResponseBodyRecordModel {
	s.AuthConf = v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetBizName(v string) *GetRecordResponseBodyRecordModel {
	s.BizName = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetComment(v string) *GetRecordResponseBodyRecordModel {
	s.Comment = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetCreateTime(v string) *GetRecordResponseBodyRecordModel {
	s.CreateTime = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetData(v *GetRecordResponseBodyRecordModelData) *GetRecordResponseBodyRecordModel {
	s.Data = v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetHostPolicy(v string) *GetRecordResponseBodyRecordModel {
	s.HostPolicy = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetProxied(v bool) *GetRecordResponseBodyRecordModel {
	s.Proxied = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordCname(v string) *GetRecordResponseBodyRecordModel {
	s.RecordCname = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordId(v int64) *GetRecordResponseBodyRecordModel {
	s.RecordId = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordName(v string) *GetRecordResponseBodyRecordModel {
	s.RecordName = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordSourceType(v string) *GetRecordResponseBodyRecordModel {
	s.RecordSourceType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetRecordType(v string) *GetRecordResponseBodyRecordModel {
	s.RecordType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetSiteId(v int64) *GetRecordResponseBodyRecordModel {
	s.SiteId = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetSiteName(v string) *GetRecordResponseBodyRecordModel {
	s.SiteName = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetTtl(v int32) *GetRecordResponseBodyRecordModel {
	s.Ttl = &v
	return s
}

func (s *GetRecordResponseBodyRecordModel) SetUpdateTime(v string) *GetRecordResponseBodyRecordModel {
	s.UpdateTime = &v
	return s
}

type GetRecordResponseBodyRecordModelAuthConf struct {
	// example:
	//
	// VIxuvJSA2S03f******kp208dy5w7
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// example:
	//
	// public
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// u0Nkg5gBK*******QF5wvKMM504JUHt
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// example:
	//
	// v2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetRecordResponseBodyRecordModelAuthConf) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResponseBodyRecordModelAuthConf) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetAccessKey(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.AccessKey = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetAuthType(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.AuthType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetRegion(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.Region = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetSecretKey(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.SecretKey = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelAuthConf) SetVersion(v string) *GetRecordResponseBodyRecordModelAuthConf {
	s.Version = &v
	return s
}

type GetRecordResponseBodyRecordModelData struct {
	// example:
	//
	// 1
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// 1
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// example:
	//
	// 1
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// example:
	//
	// 8707
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// RSA
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// example:
	//
	// example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s GetRecordResponseBodyRecordModelData) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResponseBodyRecordModelData) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBodyRecordModelData) SetAlgorithm(v int32) *GetRecordResponseBodyRecordModelData {
	s.Algorithm = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetCertificate(v string) *GetRecordResponseBodyRecordModelData {
	s.Certificate = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetFingerprint(v string) *GetRecordResponseBodyRecordModelData {
	s.Fingerprint = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetFlag(v int32) *GetRecordResponseBodyRecordModelData {
	s.Flag = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetKeyTag(v int32) *GetRecordResponseBodyRecordModelData {
	s.KeyTag = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetMatchingType(v int32) *GetRecordResponseBodyRecordModelData {
	s.MatchingType = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetPort(v int32) *GetRecordResponseBodyRecordModelData {
	s.Port = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetPriority(v int32) *GetRecordResponseBodyRecordModelData {
	s.Priority = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetSelector(v int32) *GetRecordResponseBodyRecordModelData {
	s.Selector = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetTag(v string) *GetRecordResponseBodyRecordModelData {
	s.Tag = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetType(v int32) *GetRecordResponseBodyRecordModelData {
	s.Type = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetUsage(v int32) *GetRecordResponseBodyRecordModelData {
	s.Usage = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetValue(v string) *GetRecordResponseBodyRecordModelData {
	s.Value = &v
	return s
}

func (s *GetRecordResponseBodyRecordModelData) SetWeight(v int32) *GetRecordResponseBodyRecordModelData {
	s.Weight = &v
	return s
}

type GetRecordResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResponse) GoString() string {
	return s.String()
}

func (s *GetRecordResponse) SetHeaders(v map[string]*string) *GetRecordResponse {
	s.Headers = v
	return s
}

func (s *GetRecordResponse) SetStatusCode(v int32) *GetRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordResponse) SetBody(v *GetRecordResponseBody) *GetRecordResponse {
	s.Body = v
	return s
}

type GetScheduledPreloadJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// GetScheduledPreloadJob
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetScheduledPreloadJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScheduledPreloadJobRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledPreloadJobRequest) SetId(v string) *GetScheduledPreloadJobRequest {
	s.Id = &v
	return s
}

type GetScheduledPreloadJobResponseBody struct {
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CreatedAt     *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Domains       *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	ErrorInfo     *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	FailedFileOss *string `json:"FailedFileOss,omitempty" xml:"FailedFileOss,omitempty"`
	FileId        *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InsertWay     *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId        *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TaskSubmitted *int32  `json:"TaskSubmitted,omitempty" xml:"TaskSubmitted,omitempty"`
	TaskType      *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UrlCount      *int32  `json:"UrlCount,omitempty" xml:"UrlCount,omitempty"`
	UrlSubmitted  *int32  `json:"UrlSubmitted,omitempty" xml:"UrlSubmitted,omitempty"`
}

func (s GetScheduledPreloadJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScheduledPreloadJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduledPreloadJobResponseBody) SetAliUid(v string) *GetScheduledPreloadJobResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetCreatedAt(v string) *GetScheduledPreloadJobResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetDomains(v string) *GetScheduledPreloadJobResponseBody {
	s.Domains = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetErrorInfo(v string) *GetScheduledPreloadJobResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetFailedFileOss(v string) *GetScheduledPreloadJobResponseBody {
	s.FailedFileOss = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetFileId(v string) *GetScheduledPreloadJobResponseBody {
	s.FileId = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetId(v string) *GetScheduledPreloadJobResponseBody {
	s.Id = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetInsertWay(v string) *GetScheduledPreloadJobResponseBody {
	s.InsertWay = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetName(v string) *GetScheduledPreloadJobResponseBody {
	s.Name = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetRequestId(v string) *GetScheduledPreloadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetSiteId(v int64) *GetScheduledPreloadJobResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetTaskSubmitted(v int32) *GetScheduledPreloadJobResponseBody {
	s.TaskSubmitted = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetTaskType(v string) *GetScheduledPreloadJobResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetUrlCount(v int32) *GetScheduledPreloadJobResponseBody {
	s.UrlCount = &v
	return s
}

func (s *GetScheduledPreloadJobResponseBody) SetUrlSubmitted(v int32) *GetScheduledPreloadJobResponseBody {
	s.UrlSubmitted = &v
	return s
}

type GetScheduledPreloadJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduledPreloadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledPreloadJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScheduledPreloadJobResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledPreloadJobResponse) SetHeaders(v map[string]*string) *GetScheduledPreloadJobResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledPreloadJobResponse) SetStatusCode(v int32) *GetScheduledPreloadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledPreloadJobResponse) SetBody(v *GetScheduledPreloadJobResponseBody) *GetScheduledPreloadJobResponse {
	s.Body = v
	return s
}

type GetSiteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSiteRequest) GoString() string {
	return s.String()
}

func (s *GetSiteRequest) SetSiteId(v int64) *GetSiteRequest {
	s.SiteId = &v
	return s
}

type GetSiteResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteModel *GetSiteResponseBodySiteModel `json:"SiteModel,omitempty" xml:"SiteModel,omitempty" type:"Struct"`
}

func (s GetSiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSiteResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteResponseBody) SetRequestId(v string) *GetSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteResponseBody) SetSiteModel(v *GetSiteResponseBodySiteModel) *GetSiteResponseBody {
	s.SiteModel = v
	return s
}

type GetSiteResponseBodySiteModel struct {
	// example:
	//
	// NS
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// example.cname.com
	CnameZone *string `json:"CnameZone,omitempty" xml:"CnameZone,omitempty"`
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// cas-merge-q6h0bv
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// male1-1.ialicdn.com,female1-1.ialicdn.com
	NameServerList *string `json:"NameServerList,omitempty" xml:"NameServerList,omitempty"`
	// example:
	//
	// plan-168777532****
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// example:
	//
	// normal
	PlanSpecName *string `json:"PlanSpecName,omitempty" xml:"PlanSpecName,omitempty"`
	// example:
	//
	// rg-aek26g6i6se****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {"tag1":"value1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 2023-12-24T02:01:11Z
	UpdateTime   *string            `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VanityNSList map[string]*string `json:"VanityNSList,omitempty" xml:"VanityNSList,omitempty"`
	// example:
	//
	// verify_d516cb3740f81f0cef77d162edd1****
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	// example:
	//
	// true
	VersionManagement *bool `json:"VersionManagement,omitempty" xml:"VersionManagement,omitempty"`
}

func (s GetSiteResponseBodySiteModel) String() string {
	return tea.Prettify(s)
}

func (s GetSiteResponseBodySiteModel) GoString() string {
	return s.String()
}

func (s *GetSiteResponseBodySiteModel) SetAccessType(v string) *GetSiteResponseBodySiteModel {
	s.AccessType = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetCnameZone(v string) *GetSiteResponseBodySiteModel {
	s.CnameZone = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetCoverage(v string) *GetSiteResponseBodySiteModel {
	s.Coverage = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetCreateTime(v string) *GetSiteResponseBodySiteModel {
	s.CreateTime = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetInstanceId(v string) *GetSiteResponseBodySiteModel {
	s.InstanceId = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetNameServerList(v string) *GetSiteResponseBodySiteModel {
	s.NameServerList = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetPlanName(v string) *GetSiteResponseBodySiteModel {
	s.PlanName = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetPlanSpecName(v string) *GetSiteResponseBodySiteModel {
	s.PlanSpecName = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetResourceGroupId(v string) *GetSiteResponseBodySiteModel {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetSiteId(v int64) *GetSiteResponseBodySiteModel {
	s.SiteId = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetSiteName(v string) *GetSiteResponseBodySiteModel {
	s.SiteName = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetStatus(v string) *GetSiteResponseBodySiteModel {
	s.Status = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetTags(v map[string]interface{}) *GetSiteResponseBodySiteModel {
	s.Tags = v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetUpdateTime(v string) *GetSiteResponseBodySiteModel {
	s.UpdateTime = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetVanityNSList(v map[string]*string) *GetSiteResponseBodySiteModel {
	s.VanityNSList = v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetVerifyCode(v string) *GetSiteResponseBodySiteModel {
	s.VerifyCode = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetVersionManagement(v bool) *GetSiteResponseBodySiteModel {
	s.VersionManagement = &v
	return s
}

type GetSiteResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSiteResponse) GoString() string {
	return s.String()
}

func (s *GetSiteResponse) SetHeaders(v map[string]*string) *GetSiteResponse {
	s.Headers = v
	return s
}

func (s *GetSiteResponse) SetStatusCode(v int32) *GetSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteResponse) SetBody(v *GetSiteResponseBody) *GetSiteResponse {
	s.Body = v
	return s
}

type GetSiteCurrentNSRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteCurrentNSRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSiteCurrentNSRequest) GoString() string {
	return s.String()
}

func (s *GetSiteCurrentNSRequest) SetSiteId(v int64) *GetSiteCurrentNSRequest {
	s.SiteId = &v
	return s
}

type GetSiteCurrentNSResponseBody struct {
	NSList []*string `json:"NSList,omitempty" xml:"NSList,omitempty" type:"Repeated"`
	// example:
	//
	// 2430E05E-1340-5773-B5E1-B743929F46F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSiteCurrentNSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSiteCurrentNSResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteCurrentNSResponseBody) SetNSList(v []*string) *GetSiteCurrentNSResponseBody {
	s.NSList = v
	return s
}

func (s *GetSiteCurrentNSResponseBody) SetRequestId(v string) *GetSiteCurrentNSResponseBody {
	s.RequestId = &v
	return s
}

type GetSiteCurrentNSResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteCurrentNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteCurrentNSResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSiteCurrentNSResponse) GoString() string {
	return s.String()
}

func (s *GetSiteCurrentNSResponse) SetHeaders(v map[string]*string) *GetSiteCurrentNSResponse {
	s.Headers = v
	return s
}

func (s *GetSiteCurrentNSResponse) SetStatusCode(v int32) *GetSiteCurrentNSResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteCurrentNSResponse) SetBody(v *GetSiteCurrentNSResponseBody) *GetSiteCurrentNSResponse {
	s.Body = v
	return s
}

type GetSiteCustomLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11223***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteCustomLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSiteCustomLogRequest) GoString() string {
	return s.String()
}

func (s *GetSiteCustomLogRequest) SetSiteId(v int64) *GetSiteCustomLogRequest {
	s.SiteId = &v
	return s
}

type GetSiteCustomLogResponseBody struct {
	// example:
	//
	// 6befa4aa-2a94-4f51-a245-295787192d2c
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// true
	IsExist        *bool                                       `json:"IsExist,omitempty" xml:"IsExist,omitempty"`
	LogCustomField *GetSiteCustomLogResponseBodyLogCustomField `json:"LogCustomField,omitempty" xml:"LogCustomField,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6befa4aa-2a94-4f51-a245-295787192d2c
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 167026711***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteCustomLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSiteCustomLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteCustomLogResponseBody) SetConfigId(v int64) *GetSiteCustomLogResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetSiteCustomLogResponseBody) SetIsExist(v bool) *GetSiteCustomLogResponseBody {
	s.IsExist = &v
	return s
}

func (s *GetSiteCustomLogResponseBody) SetLogCustomField(v *GetSiteCustomLogResponseBodyLogCustomField) *GetSiteCustomLogResponseBody {
	s.LogCustomField = v
	return s
}

func (s *GetSiteCustomLogResponseBody) SetRequestId(v string) *GetSiteCustomLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteCustomLogResponseBody) SetSiteId(v int64) *GetSiteCustomLogResponseBody {
	s.SiteId = &v
	return s
}

type GetSiteCustomLogResponseBodyLogCustomField struct {
	Cookies         []*string `json:"Cookies,omitempty" xml:"Cookies,omitempty" type:"Repeated"`
	RequestHeaders  []*string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	ResponseHeaders []*string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
}

func (s GetSiteCustomLogResponseBodyLogCustomField) String() string {
	return tea.Prettify(s)
}

func (s GetSiteCustomLogResponseBodyLogCustomField) GoString() string {
	return s.String()
}

func (s *GetSiteCustomLogResponseBodyLogCustomField) SetCookies(v []*string) *GetSiteCustomLogResponseBodyLogCustomField {
	s.Cookies = v
	return s
}

func (s *GetSiteCustomLogResponseBodyLogCustomField) SetRequestHeaders(v []*string) *GetSiteCustomLogResponseBodyLogCustomField {
	s.RequestHeaders = v
	return s
}

func (s *GetSiteCustomLogResponseBodyLogCustomField) SetResponseHeaders(v []*string) *GetSiteCustomLogResponseBodyLogCustomField {
	s.ResponseHeaders = v
	return s
}

type GetSiteCustomLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteCustomLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteCustomLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSiteCustomLogResponse) GoString() string {
	return s.String()
}

func (s *GetSiteCustomLogResponse) SetHeaders(v map[string]*string) *GetSiteCustomLogResponse {
	s.Headers = v
	return s
}

func (s *GetSiteCustomLogResponse) SetStatusCode(v int32) *GetSiteCustomLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteCustomLogResponse) SetBody(v *GetSiteCustomLogResponseBody) *GetSiteCustomLogResponse {
	s.Body = v
	return s
}

type GetSiteDeliveryTaskRequest struct {
	// example:
	//
	// 123456***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetSiteDeliveryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSiteDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *GetSiteDeliveryTaskRequest) SetSiteId(v int64) *GetSiteDeliveryTaskRequest {
	s.SiteId = &v
	return s
}

func (s *GetSiteDeliveryTaskRequest) SetTaskName(v string) *GetSiteDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

type GetSiteDeliveryTaskResponseBody struct {
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// 0.0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// example:
	//
	// Client,UserAgent
	FieldList *string `json:"FieldList,omitempty" xml:"FieldList,omitempty"`
	// example:
	//
	// []
	FilterRules *string `json:"FilterRules,omitempty" xml:"FilterRules,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {\\"Region\\": \\"cn-hangzhou\\", \\"Endpoint\\": \\"https://***.oss-cn-hangzhou.aliyuncs.com\\", \\"BucketPath\\": \\"hjy-test002/online-logs\\"}
	SinkConfig interface{} `json:"SinkConfig,omitempty" xml:"SinkConfig,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// test.***.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetSiteDeliveryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSiteDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteDeliveryTaskResponseBody) SetBusinessType(v string) *GetSiteDeliveryTaskResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetDataCenter(v string) *GetSiteDeliveryTaskResponseBody {
	s.DataCenter = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetDeliveryType(v string) *GetSiteDeliveryTaskResponseBody {
	s.DeliveryType = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetDiscardRate(v float32) *GetSiteDeliveryTaskResponseBody {
	s.DiscardRate = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetFieldList(v string) *GetSiteDeliveryTaskResponseBody {
	s.FieldList = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetFilterRules(v string) *GetSiteDeliveryTaskResponseBody {
	s.FilterRules = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetRequestId(v string) *GetSiteDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetSinkConfig(v interface{}) *GetSiteDeliveryTaskResponseBody {
	s.SinkConfig = v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetSiteId(v int64) *GetSiteDeliveryTaskResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetSiteName(v string) *GetSiteDeliveryTaskResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetStatus(v string) *GetSiteDeliveryTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetTaskName(v string) *GetSiteDeliveryTaskResponseBody {
	s.TaskName = &v
	return s
}

type GetSiteDeliveryTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteDeliveryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSiteDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *GetSiteDeliveryTaskResponse) SetHeaders(v map[string]*string) *GetSiteDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *GetSiteDeliveryTaskResponse) SetStatusCode(v int32) *GetSiteDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteDeliveryTaskResponse) SetBody(v *GetSiteDeliveryTaskResponseBody) *GetSiteDeliveryTaskResponse {
	s.Body = v
	return s
}

type GetSiteLogDeliveryQuotaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteLogDeliveryQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSiteLogDeliveryQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetSiteLogDeliveryQuotaRequest) SetBusinessType(v string) *GetSiteLogDeliveryQuotaRequest {
	s.BusinessType = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaRequest) SetSiteId(v int64) *GetSiteLogDeliveryQuotaRequest {
	s.SiteId = &v
	return s
}

type GetSiteLogDeliveryQuotaResponseBody struct {
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 3
	FreeQuota *int64 `json:"FreeQuota,omitempty" xml:"FreeQuota,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4C14840EF0EAAB6D97CDE0C5F6554ACE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteLogDeliveryQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSiteLogDeliveryQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteLogDeliveryQuotaResponseBody) SetBusinessType(v string) *GetSiteLogDeliveryQuotaResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponseBody) SetFreeQuota(v int64) *GetSiteLogDeliveryQuotaResponseBody {
	s.FreeQuota = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponseBody) SetRequestId(v string) *GetSiteLogDeliveryQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponseBody) SetSiteId(v int64) *GetSiteLogDeliveryQuotaResponseBody {
	s.SiteId = &v
	return s
}

type GetSiteLogDeliveryQuotaResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteLogDeliveryQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteLogDeliveryQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSiteLogDeliveryQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetSiteLogDeliveryQuotaResponse) SetHeaders(v map[string]*string) *GetSiteLogDeliveryQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponse) SetStatusCode(v int32) *GetSiteLogDeliveryQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponse) SetBody(v *GetSiteLogDeliveryQuotaResponseBody) *GetSiteLogDeliveryQuotaResponse {
	s.Body = v
	return s
}

type GetSiteWafSettingsRequest struct {
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetSiteWafSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSiteWafSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetSiteWafSettingsRequest) SetSiteId(v int64) *GetSiteWafSettingsRequest {
	s.SiteId = &v
	return s
}

func (s *GetSiteWafSettingsRequest) SetSiteVersion(v int32) *GetSiteWafSettingsRequest {
	s.SiteVersion = &v
	return s
}

type GetSiteWafSettingsResponseBody struct {
	// Id of the request
	RequestId *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Settings  *WafSiteSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
}

func (s GetSiteWafSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSiteWafSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteWafSettingsResponseBody) SetRequestId(v string) *GetSiteWafSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteWafSettingsResponseBody) SetSettings(v *WafSiteSettings) *GetSiteWafSettingsResponseBody {
	s.Settings = v
	return s
}

type GetSiteWafSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteWafSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteWafSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSiteWafSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetSiteWafSettingsResponse) SetHeaders(v map[string]*string) *GetSiteWafSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetSiteWafSettingsResponse) SetStatusCode(v int32) *GetSiteWafSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteWafSettingsResponse) SetBody(v *GetSiteWafSettingsResponseBody) *GetSiteWafSettingsResponse {
	s.Body = v
	return s
}

type GetUploadTaskRequest struct {
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 159253299357****
	UploadId *int64 `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s GetUploadTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTaskRequest) GoString() string {
	return s.String()
}

func (s *GetUploadTaskRequest) SetSiteId(v int64) *GetUploadTaskRequest {
	s.SiteId = &v
	return s
}

func (s *GetUploadTaskRequest) SetUploadId(v int64) *GetUploadTaskRequest {
	s.UploadId = &v
	return s
}

type GetUploadTaskResponseBody struct {
	// example:
	//
	// invalid url
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ET5BF670-09D5-4D0B-BEBY-D96A2A52****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUploadTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadTaskResponseBody) SetDescription(v string) *GetUploadTaskResponseBody {
	s.Description = &v
	return s
}

func (s *GetUploadTaskResponseBody) SetRequestId(v string) *GetUploadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadTaskResponseBody) SetStatus(v string) *GetUploadTaskResponseBody {
	s.Status = &v
	return s
}

type GetUploadTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadTaskResponse) GoString() string {
	return s.String()
}

func (s *GetUploadTaskResponse) SetHeaders(v map[string]*string) *GetUploadTaskResponse {
	s.Headers = v
	return s
}

func (s *GetUploadTaskResponse) SetStatusCode(v int32) *GetUploadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadTaskResponse) SetBody(v *GetUploadTaskResponseBody) *GetUploadTaskResponse {
	s.Body = v
	return s
}

type GetUserDeliveryTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetUserDeliveryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *GetUserDeliveryTaskRequest) SetTaskName(v string) *GetUserDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

type GetUserDeliveryTaskResponseBody struct {
	// example:
	//
	// dcdn_log_er
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// example:
	//
	// oss
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// 0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// example:
	//
	// ClientRequestID,ClientRequestHost
	FieldList *string `json:"FieldList,omitempty" xml:"FieldList,omitempty"`
	// example:
	//
	// [{"ClientSSLProtocol": {"equals": ["TLSv1.3"]}}]
	FilterRules *string `json:"FilterRules,omitempty" xml:"FilterRules,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7072132a-bd3c-46a6-9e81-aba3e0e3f861
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {\\"Project\\": \\"er-online-hjy-pro\\", \\"Logstore\\": \\"er-online-hjy-log\\", \\"Region\\": \\"cn-hangzhou\\", \\"Endpoint\\": \\"cn-hangzhou.log.aliyuncs.com\\", \\"Aliuid\\": \\"1077912128805410\\"}
	SinkConfig interface{} `json:"SinkConfig,omitempty" xml:"SinkConfig,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// testoss11
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetUserDeliveryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserDeliveryTaskResponseBody) SetBusinessType(v string) *GetUserDeliveryTaskResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetDataCenter(v string) *GetUserDeliveryTaskResponseBody {
	s.DataCenter = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetDeliveryType(v string) *GetUserDeliveryTaskResponseBody {
	s.DeliveryType = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetDiscardRate(v float32) *GetUserDeliveryTaskResponseBody {
	s.DiscardRate = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetFieldList(v string) *GetUserDeliveryTaskResponseBody {
	s.FieldList = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetFilterRules(v string) *GetUserDeliveryTaskResponseBody {
	s.FilterRules = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetRequestId(v string) *GetUserDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetSinkConfig(v interface{}) *GetUserDeliveryTaskResponseBody {
	s.SinkConfig = v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetStatus(v string) *GetUserDeliveryTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetTaskName(v string) *GetUserDeliveryTaskResponseBody {
	s.TaskName = &v
	return s
}

type GetUserDeliveryTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserDeliveryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *GetUserDeliveryTaskResponse) SetHeaders(v map[string]*string) *GetUserDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *GetUserDeliveryTaskResponse) SetStatusCode(v int32) *GetUserDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserDeliveryTaskResponse) SetBody(v *GetUserDeliveryTaskResponseBody) *GetUserDeliveryTaskResponse {
	s.Body = v
	return s
}

type GetUserLogDeliveryQuotaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s GetUserLogDeliveryQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserLogDeliveryQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetUserLogDeliveryQuotaRequest) SetBusinessType(v string) *GetUserLogDeliveryQuotaRequest {
	s.BusinessType = &v
	return s
}

type GetUserLogDeliveryQuotaResponseBody struct {
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 3
	FreeQuota *int64 `json:"FreeQuota,omitempty" xml:"FreeQuota,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserLogDeliveryQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserLogDeliveryQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserLogDeliveryQuotaResponseBody) SetBusinessType(v string) *GetUserLogDeliveryQuotaResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetUserLogDeliveryQuotaResponseBody) SetFreeQuota(v int64) *GetUserLogDeliveryQuotaResponseBody {
	s.FreeQuota = &v
	return s
}

func (s *GetUserLogDeliveryQuotaResponseBody) SetRequestId(v string) *GetUserLogDeliveryQuotaResponseBody {
	s.RequestId = &v
	return s
}

type GetUserLogDeliveryQuotaResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserLogDeliveryQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserLogDeliveryQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserLogDeliveryQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetUserLogDeliveryQuotaResponse) SetHeaders(v map[string]*string) *GetUserLogDeliveryQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetUserLogDeliveryQuotaResponse) SetStatusCode(v int32) *GetUserLogDeliveryQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserLogDeliveryQuotaResponse) SetBody(v *GetUserLogDeliveryQuotaResponseBody) *GetUserLogDeliveryQuotaResponse {
	s.Body = v
	return s
}

type GetWafBotAppKeyResponseBody struct {
	// APP key
	//
	// example:
	//
	// example_appkey
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWafBotAppKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWafBotAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetWafBotAppKeyResponseBody) SetAppKey(v string) *GetWafBotAppKeyResponseBody {
	s.AppKey = &v
	return s
}

func (s *GetWafBotAppKeyResponseBody) SetRequestId(v string) *GetWafBotAppKeyResponseBody {
	s.RequestId = &v
	return s
}

type GetWafBotAppKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWafBotAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWafBotAppKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWafBotAppKeyResponse) GoString() string {
	return s.String()
}

func (s *GetWafBotAppKeyResponse) SetHeaders(v map[string]*string) *GetWafBotAppKeyResponse {
	s.Headers = v
	return s
}

func (s *GetWafBotAppKeyResponse) SetStatusCode(v int32) *GetWafBotAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWafBotAppKeyResponse) SetBody(v *GetWafBotAppKeyResponseBody) *GetWafBotAppKeyResponse {
	s.Body = v
	return s
}

type GetWafFilterRequest struct {
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// characteristics
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// example:
	//
	// http_custom_cc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWafFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWafFilterRequest) GoString() string {
	return s.String()
}

func (s *GetWafFilterRequest) SetPhase(v string) *GetWafFilterRequest {
	s.Phase = &v
	return s
}

func (s *GetWafFilterRequest) SetSiteId(v int64) *GetWafFilterRequest {
	s.SiteId = &v
	return s
}

func (s *GetWafFilterRequest) SetTarget(v string) *GetWafFilterRequest {
	s.Target = &v
	return s
}

func (s *GetWafFilterRequest) SetType(v string) *GetWafFilterRequest {
	s.Type = &v
	return s
}

type GetWafFilterResponseBody struct {
	Filter *GetWafFilterResponseBodyFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWafFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWafFilterResponseBody) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBody) SetFilter(v *GetWafFilterResponseBodyFilter) *GetWafFilterResponseBody {
	s.Filter = v
	return s
}

func (s *GetWafFilterResponseBody) SetRequestId(v string) *GetWafFilterResponseBody {
	s.RequestId = &v
	return s
}

type GetWafFilterResponseBodyFilter struct {
	Fields []*GetWafFilterResponseBodyFilterFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// characteristics
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// example:
	//
	// http_custom_cc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWafFilterResponseBodyFilter) String() string {
	return tea.Prettify(s)
}

func (s GetWafFilterResponseBodyFilter) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilter) SetFields(v []*GetWafFilterResponseBodyFilterFields) *GetWafFilterResponseBodyFilter {
	s.Fields = v
	return s
}

func (s *GetWafFilterResponseBodyFilter) SetPhase(v string) *GetWafFilterResponseBodyFilter {
	s.Phase = &v
	return s
}

func (s *GetWafFilterResponseBodyFilter) SetTarget(v string) *GetWafFilterResponseBodyFilter {
	s.Target = &v
	return s
}

func (s *GetWafFilterResponseBodyFilter) SetType(v string) *GetWafFilterResponseBodyFilter {
	s.Type = &v
	return s
}

type GetWafFilterResponseBodyFilterFields struct {
	// example:
	//
	// http.request.headers
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// Header
	Label    *string                                       `json:"Label,omitempty" xml:"Label,omitempty"`
	Logics   []*GetWafFilterResponseBodyFilterFieldsLogics `json:"Logics,omitempty" xml:"Logics,omitempty" type:"Repeated"`
	Selector *GetWafFilterResponseBodyFilterFieldsSelector `json:"Selector,omitempty" xml:"Selector,omitempty" type:"Struct"`
	Sub      *bool                                         `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// example:
	//
	// e.g. Content-Type
	SubTip *string `json:"SubTip,omitempty" xml:"SubTip,omitempty"`
}

func (s GetWafFilterResponseBodyFilterFields) String() string {
	return tea.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFields) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFields) SetKey(v string) *GetWafFilterResponseBodyFilterFields {
	s.Key = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetLabel(v string) *GetWafFilterResponseBodyFilterFields {
	s.Label = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetLogics(v []*GetWafFilterResponseBodyFilterFieldsLogics) *GetWafFilterResponseBodyFilterFields {
	s.Logics = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetSelector(v *GetWafFilterResponseBodyFilterFieldsSelector) *GetWafFilterResponseBodyFilterFields {
	s.Selector = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetSub(v bool) *GetWafFilterResponseBodyFilterFields {
	s.Sub = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetSubTip(v string) *GetWafFilterResponseBodyFilterFields {
	s.SubTip = &v
	return s
}

type GetWafFilterResponseBodyFilterFieldsLogics struct {
	// example:
	//
	// 1
	Attributes *int32 `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// input:single
	Kind     *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Negative *bool   `json:"Negative,omitempty" xml:"Negative,omitempty"`
	// example:
	//
	// Does not equal
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// eq
	Symbol *string `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	// example:
	//
	// e.g. image/jpeg
	Tip *string `json:"Tip,omitempty" xml:"Tip,omitempty"`
	// example:
	//
	// string
	Type      *string                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	Validator *GetWafFilterResponseBodyFilterFieldsLogicsValidator `json:"Validator,omitempty" xml:"Validator,omitempty" type:"Struct"`
}

func (s GetWafFilterResponseBodyFilterFieldsLogics) String() string {
	return tea.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsLogics) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetAttributes(v int32) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Attributes = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetKind(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Kind = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetNegative(v bool) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Negative = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetOperator(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Operator = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetSymbol(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Symbol = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetTip(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Tip = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetType(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Type = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetValidator(v *GetWafFilterResponseBodyFilterFieldsLogicsValidator) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Validator = v
	return s
}

type GetWafFilterResponseBodyFilterFieldsLogicsValidator struct {
	// example:
	//
	// Enter a valid expression
	ErrMsg *string          `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	Length *WafQuotaInteger `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// ^example$
	Pattern *string          `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Range   *WafQuotaInteger `json:"Range,omitempty" xml:"Range,omitempty"`
}

func (s GetWafFilterResponseBodyFilterFieldsLogicsValidator) String() string {
	return tea.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsLogicsValidator) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) SetErrMsg(v string) *GetWafFilterResponseBodyFilterFieldsLogicsValidator {
	s.ErrMsg = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) SetLength(v *WafQuotaInteger) *GetWafFilterResponseBodyFilterFieldsLogicsValidator {
	s.Length = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) SetPattern(v string) *GetWafFilterResponseBodyFilterFieldsLogicsValidator {
	s.Pattern = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) SetRange(v *WafQuotaInteger) *GetWafFilterResponseBodyFilterFieldsLogicsValidator {
	s.Range = v
	return s
}

type GetWafFilterResponseBodyFilterFieldsSelector struct {
	Data []*GetWafFilterResponseBodyFilterFieldsSelectorData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// data
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
}

func (s GetWafFilterResponseBodyFilterFieldsSelector) String() string {
	return tea.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsSelector) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsSelector) SetData(v []*GetWafFilterResponseBodyFilterFieldsSelectorData) *GetWafFilterResponseBodyFilterFieldsSelector {
	s.Data = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSelector) SetKind(v string) *GetWafFilterResponseBodyFilterFieldsSelector {
	s.Kind = &v
	return s
}

type GetWafFilterResponseBodyFilterFieldsSelectorData struct {
	// example:
	//
	// China
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// CN
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetWafFilterResponseBodyFilterFieldsSelectorData) String() string {
	return tea.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsSelectorData) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsSelectorData) SetLabel(v string) *GetWafFilterResponseBodyFilterFieldsSelectorData {
	s.Label = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSelectorData) SetValue(v string) *GetWafFilterResponseBodyFilterFieldsSelectorData {
	s.Value = &v
	return s
}

type GetWafFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWafFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWafFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWafFilterResponse) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponse) SetHeaders(v map[string]*string) *GetWafFilterResponse {
	s.Headers = v
	return s
}

func (s *GetWafFilterResponse) SetStatusCode(v int32) *GetWafFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWafFilterResponse) SetBody(v *GetWafFilterResponseBody) *GetWafFilterResponse {
	s.Body = v
	return s
}

type GetWafQuotaRequest struct {
	// example:
	//
	// page
	Paths *string `json:"Paths,omitempty" xml:"Paths,omitempty"`
}

func (s GetWafQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWafQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetWafQuotaRequest) SetPaths(v string) *GetWafQuotaRequest {
	s.Paths = &v
	return s
}

type GetWafQuotaResponseBody struct {
	Quota *GetWafQuotaResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWafQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWafQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBody) SetQuota(v *GetWafQuotaResponseBodyQuota) *GetWafQuotaResponseBody {
	s.Quota = v
	return s
}

func (s *GetWafQuotaResponseBody) SetRequestId(v string) *GetWafQuotaResponseBody {
	s.RequestId = &v
	return s
}

type GetWafQuotaResponseBodyQuota struct {
	List              *GetWafQuotaResponseBodyQuotaList              `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	ManagedRulesGroup *GetWafQuotaResponseBodyQuotaManagedRulesGroup `json:"ManagedRulesGroup,omitempty" xml:"ManagedRulesGroup,omitempty" type:"Struct"`
	Page              *GetWafQuotaResponseBodyQuotaPage              `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	ScenePolicy       *GetWafQuotaResponseBodyQuotaScenePolicy       `json:"ScenePolicy,omitempty" xml:"ScenePolicy,omitempty" type:"Struct"`
}

func (s GetWafQuotaResponseBodyQuota) String() string {
	return tea.Prettify(s)
}

func (s GetWafQuotaResponseBodyQuota) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBodyQuota) SetList(v *GetWafQuotaResponseBodyQuotaList) *GetWafQuotaResponseBodyQuota {
	s.List = v
	return s
}

func (s *GetWafQuotaResponseBodyQuota) SetManagedRulesGroup(v *GetWafQuotaResponseBodyQuotaManagedRulesGroup) *GetWafQuotaResponseBodyQuota {
	s.ManagedRulesGroup = v
	return s
}

func (s *GetWafQuotaResponseBodyQuota) SetPage(v *GetWafQuotaResponseBodyQuotaPage) *GetWafQuotaResponseBodyQuota {
	s.Page = v
	return s
}

func (s *GetWafQuotaResponseBodyQuota) SetScenePolicy(v *GetWafQuotaResponseBodyQuotaScenePolicy) *GetWafQuotaResponseBodyQuota {
	s.ScenePolicy = v
	return s
}

type GetWafQuotaResponseBodyQuotaList struct {
	Enable             *bool                           `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Items              map[string]*QuotaListItemsValue `json:"Items,omitempty" xml:"Items,omitempty"`
	NumberItemsPerList *WafQuotaInteger                `json:"NumberItemsPerList,omitempty" xml:"NumberItemsPerList,omitempty"`
	NumberItemsTotal   *WafQuotaInteger                `json:"NumberItemsTotal,omitempty" xml:"NumberItemsTotal,omitempty"`
	NumberTotal        *WafQuotaInteger                `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
}

func (s GetWafQuotaResponseBodyQuotaList) String() string {
	return tea.Prettify(s)
}

func (s GetWafQuotaResponseBodyQuotaList) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBodyQuotaList) SetEnable(v bool) *GetWafQuotaResponseBodyQuotaList {
	s.Enable = &v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaList) SetItems(v map[string]*QuotaListItemsValue) *GetWafQuotaResponseBodyQuotaList {
	s.Items = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaList) SetNumberItemsPerList(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaList {
	s.NumberItemsPerList = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaList) SetNumberItemsTotal(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaList {
	s.NumberItemsTotal = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaList) SetNumberTotal(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaList {
	s.NumberTotal = v
	return s
}

type GetWafQuotaResponseBodyQuotaManagedRulesGroup struct {
	Enable      *bool            `json:"Enable,omitempty" xml:"Enable,omitempty"`
	NumberTotal *WafQuotaInteger `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
}

func (s GetWafQuotaResponseBodyQuotaManagedRulesGroup) String() string {
	return tea.Prettify(s)
}

func (s GetWafQuotaResponseBodyQuotaManagedRulesGroup) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBodyQuotaManagedRulesGroup) SetEnable(v bool) *GetWafQuotaResponseBodyQuotaManagedRulesGroup {
	s.Enable = &v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaManagedRulesGroup) SetNumberTotal(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaManagedRulesGroup {
	s.NumberTotal = v
	return s
}

type GetWafQuotaResponseBodyQuotaPage struct {
	ContentTypes map[string]*QuotaPageContentTypesValue `json:"ContentTypes,omitempty" xml:"ContentTypes,omitempty"`
	Enable       *bool                                  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	NumberTotal  *WafQuotaInteger                       `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
}

func (s GetWafQuotaResponseBodyQuotaPage) String() string {
	return tea.Prettify(s)
}

func (s GetWafQuotaResponseBodyQuotaPage) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBodyQuotaPage) SetContentTypes(v map[string]*QuotaPageContentTypesValue) *GetWafQuotaResponseBodyQuotaPage {
	s.ContentTypes = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaPage) SetEnable(v bool) *GetWafQuotaResponseBodyQuotaPage {
	s.Enable = &v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaPage) SetNumberTotal(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaPage {
	s.NumberTotal = v
	return s
}

type GetWafQuotaResponseBodyQuotaScenePolicy struct {
	Enable      *bool            `json:"Enable,omitempty" xml:"Enable,omitempty"`
	NumberTotal *WafQuotaInteger `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
}

func (s GetWafQuotaResponseBodyQuotaScenePolicy) String() string {
	return tea.Prettify(s)
}

func (s GetWafQuotaResponseBodyQuotaScenePolicy) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBodyQuotaScenePolicy) SetEnable(v bool) *GetWafQuotaResponseBodyQuotaScenePolicy {
	s.Enable = &v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaScenePolicy) SetNumberTotal(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaScenePolicy {
	s.NumberTotal = v
	return s
}

type GetWafQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWafQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWafQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWafQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponse) SetHeaders(v map[string]*string) *GetWafQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetWafQuotaResponse) SetStatusCode(v int32) *GetWafQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWafQuotaResponse) SetBody(v *GetWafQuotaResponseBody) *GetWafQuotaResponse {
	s.Body = v
	return s
}

type GetWafRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetWafRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWafRuleRequest) GoString() string {
	return s.String()
}

func (s *GetWafRuleRequest) SetId(v int64) *GetWafRuleRequest {
	s.Id = &v
	return s
}

func (s *GetWafRuleRequest) SetSiteId(v int64) *GetWafRuleRequest {
	s.SiteId = &v
	return s
}

type GetWafRuleResponseBody struct {
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// 自定义响应页面ID
	//
	// example:
	//
	// 2000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 自定义响应页面名称
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 自定义响应页面内容类型
	//
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetWafRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWafRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetWafRuleResponseBody) SetConfig(v *WafRuleConfig) *GetWafRuleResponseBody {
	s.Config = v
	return s
}

func (s *GetWafRuleResponseBody) SetId(v int64) *GetWafRuleResponseBody {
	s.Id = &v
	return s
}

func (s *GetWafRuleResponseBody) SetName(v string) *GetWafRuleResponseBody {
	s.Name = &v
	return s
}

func (s *GetWafRuleResponseBody) SetPhase(v string) *GetWafRuleResponseBody {
	s.Phase = &v
	return s
}

func (s *GetWafRuleResponseBody) SetPosition(v int64) *GetWafRuleResponseBody {
	s.Position = &v
	return s
}

func (s *GetWafRuleResponseBody) SetRequestId(v string) *GetWafRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWafRuleResponseBody) SetStatus(v string) *GetWafRuleResponseBody {
	s.Status = &v
	return s
}

func (s *GetWafRuleResponseBody) SetUpdateTime(v string) *GetWafRuleResponseBody {
	s.UpdateTime = &v
	return s
}

type GetWafRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWafRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWafRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWafRuleResponse) GoString() string {
	return s.String()
}

func (s *GetWafRuleResponse) SetHeaders(v map[string]*string) *GetWafRuleResponse {
	s.Headers = v
	return s
}

func (s *GetWafRuleResponse) SetStatusCode(v int32) *GetWafRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWafRuleResponse) SetBody(v *GetWafRuleResponseBody) *GetWafRuleResponse {
	s.Body = v
	return s
}

type GetWafRulesetRequest struct {
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetWafRulesetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *GetWafRulesetRequest) SetId(v int64) *GetWafRulesetRequest {
	s.Id = &v
	return s
}

func (s *GetWafRulesetRequest) SetPhase(v string) *GetWafRulesetRequest {
	s.Phase = &v
	return s
}

func (s *GetWafRulesetRequest) SetSiteId(v int64) *GetWafRulesetRequest {
	s.SiteId = &v
	return s
}

type GetWafRulesetResponseBody struct {
	// 自定义响应页面ID
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 自定义响应页面名称
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 自定义响应页面内容类型
	//
	// This parameter is required.
	//
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*WafRuleConfig    `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Shared    *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetWafRulesetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *GetWafRulesetResponseBody) SetId(v int64) *GetWafRulesetResponseBody {
	s.Id = &v
	return s
}

func (s *GetWafRulesetResponseBody) SetName(v string) *GetWafRulesetResponseBody {
	s.Name = &v
	return s
}

func (s *GetWafRulesetResponseBody) SetPhase(v string) *GetWafRulesetResponseBody {
	s.Phase = &v
	return s
}

func (s *GetWafRulesetResponseBody) SetRequestId(v string) *GetWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWafRulesetResponseBody) SetRules(v []*WafRuleConfig) *GetWafRulesetResponseBody {
	s.Rules = v
	return s
}

func (s *GetWafRulesetResponseBody) SetShared(v *WafBatchRuleShared) *GetWafRulesetResponseBody {
	s.Shared = v
	return s
}

func (s *GetWafRulesetResponseBody) SetStatus(v string) *GetWafRulesetResponseBody {
	s.Status = &v
	return s
}

func (s *GetWafRulesetResponseBody) SetUpdateTime(v string) *GetWafRulesetResponseBody {
	s.UpdateTime = &v
	return s
}

type GetWafRulesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWafRulesetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *GetWafRulesetResponse) SetHeaders(v map[string]*string) *GetWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *GetWafRulesetResponse) SetStatusCode(v int32) *GetWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWafRulesetResponse) SetBody(v *GetWafRulesetResponseBody) *GetWafRulesetResponse {
	s.Body = v
	return s
}

type ListEdgeContainerAppRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// CreateTime
	OrderKey *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ver-1006157458290860032
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s ListEdgeContainerAppRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeContainerAppRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppRecordsRequest) SetAppId(v string) *ListEdgeContainerAppRecordsRequest {
	s.AppId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) SetOrderKey(v string) *ListEdgeContainerAppRecordsRequest {
	s.OrderKey = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) SetOrderType(v string) *ListEdgeContainerAppRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) SetPageNumber(v int32) *ListEdgeContainerAppRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) SetPageSize(v int32) *ListEdgeContainerAppRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerAppRecordsRequest) SetSearchKey(v string) *ListEdgeContainerAppRecordsRequest {
	s.SearchKey = &v
	return s
}

type ListEdgeContainerAppRecordsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*ListEdgeContainerAppRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEdgeContainerAppRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeContainerAppRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppRecordsResponseBody) SetPageNumber(v int32) *ListEdgeContainerAppRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBody) SetPageSize(v int32) *ListEdgeContainerAppRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBody) SetRecords(v []*ListEdgeContainerAppRecordsResponseBodyRecords) *ListEdgeContainerAppRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBody) SetRequestId(v string) *ListEdgeContainerAppRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBody) SetTotalCount(v int32) *ListEdgeContainerAppRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEdgeContainerAppRecordsResponseBodyRecords struct {
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// kdxceo****.yun****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// example:
	//
	// 27522948436****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 266****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// example:
	//
	// 123456
	SchemdId *int32 `json:"SchemdId,omitempty" xml:"SchemdId,omitempty"`
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 2021-12-22T08:32:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEdgeContainerAppRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeContainerAppRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetAppId(v string) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.AppId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetCname(v string) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.Cname = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetConfigId(v int64) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.ConfigId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetCreateTime(v string) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetRecordId(v int64) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.RecordId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetRecordName(v string) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.RecordName = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetSchemdId(v int32) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.SchemdId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetSiteId(v int64) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.SiteId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetUpdateTime(v string) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.UpdateTime = &v
	return s
}

type ListEdgeContainerAppRecordsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeContainerAppRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeContainerAppRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeContainerAppRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppRecordsResponse) SetHeaders(v map[string]*string) *ListEdgeContainerAppRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeContainerAppRecordsResponse) SetStatusCode(v int32) *ListEdgeContainerAppRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponse) SetBody(v *ListEdgeContainerAppRecordsResponseBody) *ListEdgeContainerAppRecordsResponse {
	s.Body = v
	return s
}

type ListEdgeContainerRecordsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// fuzzy
	RecordMatchType *string `json:"RecordMatchType,omitempty" xml:"RecordMatchType,omitempty"`
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListEdgeContainerRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeContainerRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerRecordsRequest) SetPageNumber(v int32) *ListEdgeContainerRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerRecordsRequest) SetPageSize(v int32) *ListEdgeContainerRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerRecordsRequest) SetRecordMatchType(v string) *ListEdgeContainerRecordsRequest {
	s.RecordMatchType = &v
	return s
}

func (s *ListEdgeContainerRecordsRequest) SetRecordName(v string) *ListEdgeContainerRecordsRequest {
	s.RecordName = &v
	return s
}

func (s *ListEdgeContainerRecordsRequest) SetSiteId(v int64) *ListEdgeContainerRecordsRequest {
	s.SiteId = &v
	return s
}

type ListEdgeContainerRecordsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*ListEdgeContainerRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEdgeContainerRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeContainerRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerRecordsResponseBody) SetPageNumber(v int32) *ListEdgeContainerRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBody) SetPageSize(v int32) *ListEdgeContainerRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBody) SetRecords(v []*ListEdgeContainerRecordsResponseBodyRecords) *ListEdgeContainerRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListEdgeContainerRecordsResponseBody) SetRequestId(v string) *ListEdgeContainerRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBody) SetTotalCount(v int32) *ListEdgeContainerRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEdgeContainerRecordsResponseBodyRecords struct {
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// a.example.com.cnamezone.com
	RecordCname *string `json:"RecordCname,omitempty" xml:"RecordCname,omitempty"`
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// example:
	//
	// 2021-12-22T08:32:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEdgeContainerRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeContainerRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetCreateTime(v string) *ListEdgeContainerRecordsResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetRecordCname(v string) *ListEdgeContainerRecordsResponseBodyRecords {
	s.RecordCname = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetRecordName(v string) *ListEdgeContainerRecordsResponseBodyRecords {
	s.RecordName = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetSiteId(v int64) *ListEdgeContainerRecordsResponseBodyRecords {
	s.SiteId = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetSiteName(v string) *ListEdgeContainerRecordsResponseBodyRecords {
	s.SiteName = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetUpdateTime(v string) *ListEdgeContainerRecordsResponseBodyRecords {
	s.UpdateTime = &v
	return s
}

type ListEdgeContainerRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeContainerRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeContainerRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeContainerRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerRecordsResponse) SetHeaders(v map[string]*string) *ListEdgeContainerRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeContainerRecordsResponse) SetStatusCode(v int32) *ListEdgeContainerRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeContainerRecordsResponse) SetBody(v *ListEdgeContainerRecordsResponseBody) *ListEdgeContainerRecordsResponse {
	s.Body = v
	return s
}

type ListEdgeRoutineRecordsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// fuzzy
	RecordMatchType *string `json:"RecordMatchType,omitempty" xml:"RecordMatchType,omitempty"`
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListEdgeRoutineRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeRoutineRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutineRecordsRequest) SetPageNumber(v int32) *ListEdgeRoutineRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeRoutineRecordsRequest) SetPageSize(v int32) *ListEdgeRoutineRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeRoutineRecordsRequest) SetRecordMatchType(v string) *ListEdgeRoutineRecordsRequest {
	s.RecordMatchType = &v
	return s
}

func (s *ListEdgeRoutineRecordsRequest) SetRecordName(v string) *ListEdgeRoutineRecordsRequest {
	s.RecordName = &v
	return s
}

func (s *ListEdgeRoutineRecordsRequest) SetSiteId(v int64) *ListEdgeRoutineRecordsRequest {
	s.SiteId = &v
	return s
}

type ListEdgeRoutineRecordsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*ListEdgeRoutineRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 121
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEdgeRoutineRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeRoutineRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutineRecordsResponseBody) SetPageNumber(v int32) *ListEdgeRoutineRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBody) SetPageSize(v int32) *ListEdgeRoutineRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBody) SetRecords(v []*ListEdgeRoutineRecordsResponseBodyRecords) *ListEdgeRoutineRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBody) SetRequestId(v string) *ListEdgeRoutineRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBody) SetTotalCount(v int32) *ListEdgeRoutineRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEdgeRoutineRecordsResponseBodyRecords struct {
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// a.example.com.cnamezone.com
	RecordCname *string `json:"RecordCname,omitempty" xml:"RecordCname,omitempty"`
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// example:
	//
	// 2023-12-22T08:32:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEdgeRoutineRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeRoutineRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetCreateTime(v string) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetRecordCname(v string) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.RecordCname = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetRecordName(v string) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.RecordName = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetSiteId(v int64) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.SiteId = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetSiteName(v string) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.SiteName = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetUpdateTime(v string) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.UpdateTime = &v
	return s
}

type ListEdgeRoutineRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeRoutineRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeRoutineRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEdgeRoutineRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutineRecordsResponse) SetHeaders(v map[string]*string) *ListEdgeRoutineRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeRoutineRecordsResponse) SetStatusCode(v int32) *ListEdgeRoutineRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponse) SetBody(v *ListEdgeRoutineRecordsResponseBody) *ListEdgeRoutineRecordsResponse {
	s.Body = v
	return s
}

type ListKvsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// prefix-
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListKvsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKvsRequest) GoString() string {
	return s.String()
}

func (s *ListKvsRequest) SetNamespace(v string) *ListKvsRequest {
	s.Namespace = &v
	return s
}

func (s *ListKvsRequest) SetPageNumber(v int32) *ListKvsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKvsRequest) SetPageSize(v int32) *ListKvsRequest {
	s.PageSize = &v
	return s
}

func (s *ListKvsRequest) SetPrefix(v string) *ListKvsRequest {
	s.Prefix = &v
	return s
}

type ListKvsResponseBody struct {
	Keys []*ListKvsResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1024
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKvsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKvsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKvsResponseBody) SetKeys(v []*ListKvsResponseBodyKeys) *ListKvsResponseBody {
	s.Keys = v
	return s
}

func (s *ListKvsResponseBody) SetPageNumber(v int32) *ListKvsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKvsResponseBody) SetPageSize(v int32) *ListKvsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKvsResponseBody) SetRequestId(v string) *ListKvsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKvsResponseBody) SetTotalCount(v int32) *ListKvsResponseBody {
	s.TotalCount = &v
	return s
}

type ListKvsResponseBodyKeys struct {
	// example:
	//
	// Key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2021-12-13T07:46:03Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListKvsResponseBodyKeys) String() string {
	return tea.Prettify(s)
}

func (s ListKvsResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListKvsResponseBodyKeys) SetName(v string) *ListKvsResponseBodyKeys {
	s.Name = &v
	return s
}

func (s *ListKvsResponseBodyKeys) SetUpdateTime(v string) *ListKvsResponseBodyKeys {
	s.UpdateTime = &v
	return s
}

type ListKvsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKvsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKvsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKvsResponse) GoString() string {
	return s.String()
}

func (s *ListKvsResponse) SetHeaders(v map[string]*string) *ListKvsResponse {
	s.Headers = v
	return s
}

func (s *ListKvsResponse) SetStatusCode(v int32) *ListKvsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKvsResponse) SetBody(v *ListKvsResponseBody) *ListKvsResponse {
	s.Body = v
	return s
}

type ListListsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ListLists
	QueryArgs *ListListsRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
}

func (s ListListsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListsRequest) GoString() string {
	return s.String()
}

func (s *ListListsRequest) SetPageNumber(v int32) *ListListsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListListsRequest) SetPageSize(v int32) *ListListsRequest {
	s.PageSize = &v
	return s
}

func (s *ListListsRequest) SetQueryArgs(v *ListListsRequestQueryArgs) *ListListsRequest {
	s.QueryArgs = v
	return s
}

type ListListsRequestQueryArgs struct {
	Desc *bool `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// a custom list
	DescriptionLike *string `json:"DescriptionLike,omitempty" xml:"DescriptionLike,omitempty"`
	// example:
	//
	// 40000001
	IdLike *string `json:"IdLike,omitempty" xml:"IdLike,omitempty"`
	// example:
	//
	// 10.1.1.1
	ItemLike *string `json:"ItemLike,omitempty" xml:"ItemLike,omitempty"`
	// example:
	//
	// ip
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// example:
	//
	// 10.1.1.1
	NameItemLike *string `json:"NameItemLike,omitempty" xml:"NameItemLike,omitempty"`
	// example:
	//
	// example
	NameLike *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	// example:
	//
	// id
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
}

func (s ListListsRequestQueryArgs) String() string {
	return tea.Prettify(s)
}

func (s ListListsRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListListsRequestQueryArgs) SetDesc(v bool) *ListListsRequestQueryArgs {
	s.Desc = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetDescriptionLike(v string) *ListListsRequestQueryArgs {
	s.DescriptionLike = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetIdLike(v string) *ListListsRequestQueryArgs {
	s.IdLike = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetItemLike(v string) *ListListsRequestQueryArgs {
	s.ItemLike = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetKind(v string) *ListListsRequestQueryArgs {
	s.Kind = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetNameItemLike(v string) *ListListsRequestQueryArgs {
	s.NameItemLike = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetNameLike(v string) *ListListsRequestQueryArgs {
	s.NameLike = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetOrderBy(v string) *ListListsRequestQueryArgs {
	s.OrderBy = &v
	return s
}

type ListListsShrinkRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ListLists
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s ListListsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListListsShrinkRequest) SetPageNumber(v int32) *ListListsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListListsShrinkRequest) SetPageSize(v int32) *ListListsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListListsShrinkRequest) SetQueryArgsShrink(v string) *ListListsShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

type ListListsResponseBody struct {
	Lists []*ListListsResponseBodyLists `json:"Lists,omitempty" xml:"Lists,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 10
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListListsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListsResponseBody) GoString() string {
	return s.String()
}

func (s *ListListsResponseBody) SetLists(v []*ListListsResponseBodyLists) *ListListsResponseBody {
	s.Lists = v
	return s
}

func (s *ListListsResponseBody) SetPageNumber(v int32) *ListListsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListListsResponseBody) SetPageSize(v int32) *ListListsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListListsResponseBody) SetRequestId(v string) *ListListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListsResponseBody) SetTotalCount(v int32) *ListListsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListListsResponseBody) SetUsage(v int64) *ListListsResponseBody {
	s.Usage = &v
	return s
}

type ListListsResponseBodyLists struct {
	// 自定义响应页面描述
	//
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ip
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// example:
	//
	// 100
	Length *int64 `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListListsResponseBodyLists) String() string {
	return tea.Prettify(s)
}

func (s ListListsResponseBodyLists) GoString() string {
	return s.String()
}

func (s *ListListsResponseBodyLists) SetDescription(v string) *ListListsResponseBodyLists {
	s.Description = &v
	return s
}

func (s *ListListsResponseBodyLists) SetId(v int64) *ListListsResponseBodyLists {
	s.Id = &v
	return s
}

func (s *ListListsResponseBodyLists) SetKind(v string) *ListListsResponseBodyLists {
	s.Kind = &v
	return s
}

func (s *ListListsResponseBodyLists) SetLength(v int64) *ListListsResponseBodyLists {
	s.Length = &v
	return s
}

func (s *ListListsResponseBodyLists) SetName(v string) *ListListsResponseBodyLists {
	s.Name = &v
	return s
}

func (s *ListListsResponseBodyLists) SetUpdateTime(v string) *ListListsResponseBodyLists {
	s.UpdateTime = &v
	return s
}

type ListListsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListListsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListsResponse) GoString() string {
	return s.String()
}

func (s *ListListsResponse) SetHeaders(v map[string]*string) *ListListsResponse {
	s.Headers = v
	return s
}

func (s *ListListsResponse) SetStatusCode(v int32) *ListListsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListsResponse) SetBody(v *ListListsResponseBody) *ListListsResponse {
	s.Body = v
	return s
}

type ListLoadBalancerRegionsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1024
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLoadBalancerRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancerRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerRegionsRequest) SetPageNumber(v int32) *ListLoadBalancerRegionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLoadBalancerRegionsRequest) SetPageSize(v int32) *ListLoadBalancerRegionsRequest {
	s.PageSize = &v
	return s
}

type ListLoadBalancerRegionsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1024
	PageSize *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Regions  []*ListLoadBalancerRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 81A5E222-24BF-17EF-9E80-A68D9B8F363D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListLoadBalancerRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancerRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerRegionsResponseBody) SetPageNumber(v int32) *ListLoadBalancerRegionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) SetPageSize(v int32) *ListLoadBalancerRegionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) SetRegions(v []*ListLoadBalancerRegionsResponseBodyRegions) *ListLoadBalancerRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) SetRequestId(v string) *ListLoadBalancerRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) SetTotalCount(v int32) *ListLoadBalancerRegionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBody) SetTotalPage(v int32) *ListLoadBalancerRegionsResponseBody {
	s.TotalPage = &v
	return s
}

type ListLoadBalancerRegionsResponseBodyRegions struct {
	RegionCnName *string `json:"RegionCnName,omitempty" xml:"RegionCnName,omitempty"`
	// example:
	//
	// SEAS
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// example:
	//
	// South East Asia
	RegionEnName *string                                                 `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	SubRegions   []*ListLoadBalancerRegionsResponseBodyRegionsSubRegions `json:"SubRegions,omitempty" xml:"SubRegions,omitempty" type:"Repeated"`
}

func (s ListLoadBalancerRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancerRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) SetRegionCnName(v string) *ListLoadBalancerRegionsResponseBodyRegions {
	s.RegionCnName = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) SetRegionCode(v string) *ListLoadBalancerRegionsResponseBodyRegions {
	s.RegionCode = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) SetRegionEnName(v string) *ListLoadBalancerRegionsResponseBodyRegions {
	s.RegionEnName = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegions) SetSubRegions(v []*ListLoadBalancerRegionsResponseBodyRegionsSubRegions) *ListLoadBalancerRegionsResponseBodyRegions {
	s.SubRegions = v
	return s
}

type ListLoadBalancerRegionsResponseBodyRegionsSubRegions struct {
	SubRegionCnName *string `json:"SubRegionCnName,omitempty" xml:"SubRegionCnName,omitempty"`
	// example:
	//
	// ID
	SubRegionCode *string `json:"SubRegionCode,omitempty" xml:"SubRegionCode,omitempty"`
	// example:
	//
	// Indonesia
	SubRegionEnName *string `json:"SubRegionEnName,omitempty" xml:"SubRegionEnName,omitempty"`
}

func (s ListLoadBalancerRegionsResponseBodyRegionsSubRegions) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancerRegionsResponseBodyRegionsSubRegions) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerRegionsResponseBodyRegionsSubRegions) SetSubRegionCnName(v string) *ListLoadBalancerRegionsResponseBodyRegionsSubRegions {
	s.SubRegionCnName = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegionsSubRegions) SetSubRegionCode(v string) *ListLoadBalancerRegionsResponseBodyRegionsSubRegions {
	s.SubRegionCode = &v
	return s
}

func (s *ListLoadBalancerRegionsResponseBodyRegionsSubRegions) SetSubRegionEnName(v string) *ListLoadBalancerRegionsResponseBodyRegionsSubRegions {
	s.SubRegionEnName = &v
	return s
}

type ListLoadBalancerRegionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLoadBalancerRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLoadBalancerRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancerRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerRegionsResponse) SetHeaders(v map[string]*string) *ListLoadBalancerRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListLoadBalancerRegionsResponse) SetStatusCode(v int32) *ListLoadBalancerRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLoadBalancerRegionsResponse) SetBody(v *ListLoadBalancerRegionsResponseBody) *ListLoadBalancerRegionsResponse {
	s.Body = v
	return s
}

type ListManagedRulesGroupsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListManagedRulesGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListManagedRulesGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListManagedRulesGroupsRequest) SetPageNumber(v int32) *ListManagedRulesGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListManagedRulesGroupsRequest) SetPageSize(v int32) *ListManagedRulesGroupsRequest {
	s.PageSize = &v
	return s
}

type ListManagedRulesGroupsResponseBody struct {
	ManagedRulesGroups []*ListManagedRulesGroupsResponseBodyManagedRulesGroups `json:"ManagedRulesGroups,omitempty" xml:"ManagedRulesGroups,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListManagedRulesGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListManagedRulesGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListManagedRulesGroupsResponseBody) SetManagedRulesGroups(v []*ListManagedRulesGroupsResponseBodyManagedRulesGroups) *ListManagedRulesGroupsResponseBody {
	s.ManagedRulesGroups = v
	return s
}

func (s *ListManagedRulesGroupsResponseBody) SetPageNumber(v int32) *ListManagedRulesGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListManagedRulesGroupsResponseBody) SetPageSize(v int32) *ListManagedRulesGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListManagedRulesGroupsResponseBody) SetRequestId(v string) *ListManagedRulesGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListManagedRulesGroupsResponseBody) SetTotalCount(v int32) *ListManagedRulesGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListManagedRulesGroupsResponseBodyManagedRulesGroups struct {
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1000
	RuleCount *int64 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s ListManagedRulesGroupsResponseBodyManagedRulesGroups) String() string {
	return tea.Prettify(s)
}

func (s ListManagedRulesGroupsResponseBodyManagedRulesGroups) GoString() string {
	return s.String()
}

func (s *ListManagedRulesGroupsResponseBodyManagedRulesGroups) SetName(v string) *ListManagedRulesGroupsResponseBodyManagedRulesGroups {
	s.Name = &v
	return s
}

func (s *ListManagedRulesGroupsResponseBodyManagedRulesGroups) SetRuleCount(v int64) *ListManagedRulesGroupsResponseBodyManagedRulesGroups {
	s.RuleCount = &v
	return s
}

type ListManagedRulesGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManagedRulesGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManagedRulesGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListManagedRulesGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListManagedRulesGroupsResponse) SetHeaders(v map[string]*string) *ListManagedRulesGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListManagedRulesGroupsResponse) SetStatusCode(v int32) *ListManagedRulesGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManagedRulesGroupsResponse) SetBody(v *ListManagedRulesGroupsResponseBody) *ListManagedRulesGroupsResponse {
	s.Body = v
	return s
}

type ListPagesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPagesRequest) GoString() string {
	return s.String()
}

func (s *ListPagesRequest) SetPageNumber(v int32) *ListPagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPagesRequest) SetPageSize(v int32) *ListPagesRequest {
	s.PageSize = &v
	return s
}

type ListPagesResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pages    []*ListPagesResponseBodyPages `json:"Pages,omitempty" xml:"Pages,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 10
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListPagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPagesResponseBody) SetPageNumber(v int32) *ListPagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPagesResponseBody) SetPageSize(v int32) *ListPagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPagesResponseBody) SetPages(v []*ListPagesResponseBodyPages) *ListPagesResponseBody {
	s.Pages = v
	return s
}

func (s *ListPagesResponseBody) SetRequestId(v string) *ListPagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPagesResponseBody) SetTotalCount(v int32) *ListPagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPagesResponseBody) SetUsage(v int64) *ListPagesResponseBody {
	s.Usage = &v
	return s
}

type ListPagesResponseBodyPages struct {
	// 自定义响应页面内容BASE64编码
	//
	// This parameter is required.
	//
	// example:
	//
	// PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 自定义响应页面内容类型
	//
	// This parameter is required.
	//
	// example:
	//
	// text/html
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// 自定义响应页面描述
	//
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// custom
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPagesResponseBodyPages) String() string {
	return tea.Prettify(s)
}

func (s ListPagesResponseBodyPages) GoString() string {
	return s.String()
}

func (s *ListPagesResponseBodyPages) SetContent(v string) *ListPagesResponseBodyPages {
	s.Content = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetContentType(v string) *ListPagesResponseBodyPages {
	s.ContentType = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetDescription(v string) *ListPagesResponseBodyPages {
	s.Description = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetId(v int64) *ListPagesResponseBodyPages {
	s.Id = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetKind(v string) *ListPagesResponseBodyPages {
	s.Kind = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetName(v string) *ListPagesResponseBodyPages {
	s.Name = &v
	return s
}

func (s *ListPagesResponseBodyPages) SetUpdateTime(v string) *ListPagesResponseBodyPages {
	s.UpdateTime = &v
	return s
}

type ListPagesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPagesResponse) GoString() string {
	return s.String()
}

func (s *ListPagesResponse) SetHeaders(v map[string]*string) *ListPagesResponse {
	s.Headers = v
	return s
}

func (s *ListPagesResponse) SetStatusCode(v int32) *ListPagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPagesResponse) SetBody(v *ListPagesResponseBody) *ListPagesResponse {
	s.Body = v
	return s
}

type ListRecordsRequest struct {
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	Proxied *string `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// example:
	//
	// fuzzy
	RecordMatchType *string `json:"RecordMatchType,omitempty" xml:"RecordMatchType,omitempty"`
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// CNAME
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListRecordsRequest) SetBizName(v string) *ListRecordsRequest {
	s.BizName = &v
	return s
}

func (s *ListRecordsRequest) SetPageNumber(v int32) *ListRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecordsRequest) SetPageSize(v int32) *ListRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecordsRequest) SetProxied(v string) *ListRecordsRequest {
	s.Proxied = &v
	return s
}

func (s *ListRecordsRequest) SetRecordMatchType(v string) *ListRecordsRequest {
	s.RecordMatchType = &v
	return s
}

func (s *ListRecordsRequest) SetRecordName(v string) *ListRecordsRequest {
	s.RecordName = &v
	return s
}

func (s *ListRecordsRequest) SetSiteId(v int64) *ListRecordsRequest {
	s.SiteId = &v
	return s
}

func (s *ListRecordsRequest) SetSourceType(v string) *ListRecordsRequest {
	s.SourceType = &v
	return s
}

func (s *ListRecordsRequest) SetType(v string) *ListRecordsRequest {
	s.Type = &v
	return s
}

type ListRecordsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*ListRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecordsResponseBody) SetPageNumber(v int32) *ListRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRecordsResponseBody) SetPageSize(v int32) *ListRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRecordsResponseBody) SetRecords(v []*ListRecordsResponseBodyRecords) *ListRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListRecordsResponseBody) SetRequestId(v string) *ListRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecordsResponseBody) SetTotalCount(v int32) *ListRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type ListRecordsResponseBodyRecords struct {
	AuthConf *ListRecordsResponseBodyRecordsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// this is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// {"value":"1.1.1.1"}
	Data       *ListRecordsResponseBodyRecordsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HostPolicy *string                             `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// example:
	//
	// a.example.com.cnamezone.com
	RecordCname *string `json:"RecordCname,omitempty" xml:"RecordCname,omitempty"`
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// example:
	//
	// OSS
	RecordSourceType *string `json:"RecordSourceType,omitempty" xml:"RecordSourceType,omitempty"`
	// example:
	//
	// A/AAAA
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// example:
	//
	// 30
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// example:
	//
	// 2023-06-07T10:02:59Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListRecordsResponseBodyRecords) SetAuthConf(v *ListRecordsResponseBodyRecordsAuthConf) *ListRecordsResponseBodyRecords {
	s.AuthConf = v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetBizName(v string) *ListRecordsResponseBodyRecords {
	s.BizName = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetComment(v string) *ListRecordsResponseBodyRecords {
	s.Comment = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetCreateTime(v string) *ListRecordsResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetData(v *ListRecordsResponseBodyRecordsData) *ListRecordsResponseBodyRecords {
	s.Data = v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetHostPolicy(v string) *ListRecordsResponseBodyRecords {
	s.HostPolicy = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetProxied(v bool) *ListRecordsResponseBodyRecords {
	s.Proxied = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetRecordCname(v string) *ListRecordsResponseBodyRecords {
	s.RecordCname = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetRecordId(v int64) *ListRecordsResponseBodyRecords {
	s.RecordId = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetRecordName(v string) *ListRecordsResponseBodyRecords {
	s.RecordName = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetRecordSourceType(v string) *ListRecordsResponseBodyRecords {
	s.RecordSourceType = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetRecordType(v string) *ListRecordsResponseBodyRecords {
	s.RecordType = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetSiteId(v int64) *ListRecordsResponseBodyRecords {
	s.SiteId = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetSiteName(v string) *ListRecordsResponseBodyRecords {
	s.SiteName = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetTtl(v int64) *ListRecordsResponseBodyRecords {
	s.Ttl = &v
	return s
}

func (s *ListRecordsResponseBodyRecords) SetUpdateTime(v string) *ListRecordsResponseBodyRecords {
	s.UpdateTime = &v
	return s
}

type ListRecordsResponseBodyRecordsAuthConf struct {
	// example:
	//
	// u0Nkg5gBK***QF5wvKMM504JUHt
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// example:
	//
	// private
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// VIxuvJSA2S03f***kp208dy5w7
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// example:
	//
	// v4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListRecordsResponseBodyRecordsAuthConf) String() string {
	return tea.Prettify(s)
}

func (s ListRecordsResponseBodyRecordsAuthConf) GoString() string {
	return s.String()
}

func (s *ListRecordsResponseBodyRecordsAuthConf) SetAccessKey(v string) *ListRecordsResponseBodyRecordsAuthConf {
	s.AccessKey = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsAuthConf) SetAuthType(v string) *ListRecordsResponseBodyRecordsAuthConf {
	s.AuthType = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsAuthConf) SetRegion(v string) *ListRecordsResponseBodyRecordsAuthConf {
	s.Region = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsAuthConf) SetSecretKey(v string) *ListRecordsResponseBodyRecordsAuthConf {
	s.SecretKey = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsAuthConf) SetVersion(v string) *ListRecordsResponseBodyRecordsAuthConf {
	s.Version = &v
	return s
}

type ListRecordsResponseBodyRecordsData struct {
	// example:
	//
	// 0
	Algorithm *int32 `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// dGVzdGFkYWxrcw==
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// abcdef1234567890
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// example:
	//
	// 128
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// 0
	KeyTag *int32 `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	// example:
	//
	// 0
	MatchingType *int32 `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 0
	Selector *int32 `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// example:
	//
	// issue
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 0
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// example:
	//
	// CNAME
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListRecordsResponseBodyRecordsData) String() string {
	return tea.Prettify(s)
}

func (s ListRecordsResponseBodyRecordsData) GoString() string {
	return s.String()
}

func (s *ListRecordsResponseBodyRecordsData) SetAlgorithm(v int32) *ListRecordsResponseBodyRecordsData {
	s.Algorithm = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetCertificate(v string) *ListRecordsResponseBodyRecordsData {
	s.Certificate = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetFingerprint(v string) *ListRecordsResponseBodyRecordsData {
	s.Fingerprint = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetFlag(v int32) *ListRecordsResponseBodyRecordsData {
	s.Flag = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetKeyTag(v int32) *ListRecordsResponseBodyRecordsData {
	s.KeyTag = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetMatchingType(v int32) *ListRecordsResponseBodyRecordsData {
	s.MatchingType = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetPort(v int32) *ListRecordsResponseBodyRecordsData {
	s.Port = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetPriority(v int32) *ListRecordsResponseBodyRecordsData {
	s.Priority = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetSelector(v int32) *ListRecordsResponseBodyRecordsData {
	s.Selector = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetTag(v string) *ListRecordsResponseBodyRecordsData {
	s.Tag = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetType(v int32) *ListRecordsResponseBodyRecordsData {
	s.Type = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetUsage(v int32) *ListRecordsResponseBodyRecordsData {
	s.Usage = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetValue(v string) *ListRecordsResponseBodyRecordsData {
	s.Value = &v
	return s
}

func (s *ListRecordsResponseBodyRecordsData) SetWeight(v int32) *ListRecordsResponseBodyRecordsData {
	s.Weight = &v
	return s
}

type ListRecordsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListRecordsResponse) SetHeaders(v map[string]*string) *ListRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListRecordsResponse) SetStatusCode(v int32) *ListRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecordsResponse) SetBody(v *ListRecordsResponseBody) *ListRecordsResponse {
	s.Body = v
	return s
}

type ListScheduledPreloadExecutionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ListScheduledPreloadExecutions
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListScheduledPreloadExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPreloadExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadExecutionsRequest) SetId(v string) *ListScheduledPreloadExecutionsRequest {
	s.Id = &v
	return s
}

type ListScheduledPreloadExecutionsResponseBody struct {
	Executions []*ListScheduledPreloadExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScheduledPreloadExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPreloadExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadExecutionsResponseBody) SetExecutions(v []*ListScheduledPreloadExecutionsResponseBodyExecutions) *ListScheduledPreloadExecutionsResponseBody {
	s.Executions = v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBody) SetRequestId(v string) *ListScheduledPreloadExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBody) SetTotalCount(v int32) *ListScheduledPreloadExecutionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListScheduledPreloadExecutionsResponseBodyExecutions struct {
	AliUid    *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListScheduledPreloadExecutionsResponseBodyExecutions) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPreloadExecutionsResponseBodyExecutions) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetAliUid(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.AliUid = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetEndTime(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.EndTime = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetId(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.Id = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetInterval(v int32) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.Interval = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetJobId(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.JobId = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetSliceLen(v int32) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.SliceLen = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetStartTime(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.StartTime = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetStatus(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.Status = &v
	return s
}

type ListScheduledPreloadExecutionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledPreloadExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduledPreloadExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPreloadExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadExecutionsResponse) SetHeaders(v map[string]*string) *ListScheduledPreloadExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledPreloadExecutionsResponse) SetStatusCode(v int32) *ListScheduledPreloadExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponse) SetBody(v *ListScheduledPreloadExecutionsResponseBody) *ListScheduledPreloadExecutionsResponse {
	s.Body = v
	return s
}

type ListScheduledPreloadJobsRequest struct {
	EndTime    *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ListScheduledPreloadJobs
	SiteId    *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListScheduledPreloadJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPreloadJobsRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadJobsRequest) SetEndTime(v int64) *ListScheduledPreloadJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListScheduledPreloadJobsRequest) SetPageNumber(v int32) *ListScheduledPreloadJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScheduledPreloadJobsRequest) SetPageSize(v int32) *ListScheduledPreloadJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduledPreloadJobsRequest) SetSiteId(v int64) *ListScheduledPreloadJobsRequest {
	s.SiteId = &v
	return s
}

func (s *ListScheduledPreloadJobsRequest) SetStartTime(v int64) *ListScheduledPreloadJobsRequest {
	s.StartTime = &v
	return s
}

type ListScheduledPreloadJobsResponseBody struct {
	Jobs []*ListScheduledPreloadJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScheduledPreloadJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPreloadJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadJobsResponseBody) SetJobs(v []*ListScheduledPreloadJobsResponseBodyJobs) *ListScheduledPreloadJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListScheduledPreloadJobsResponseBody) SetRequestId(v string) *ListScheduledPreloadJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBody) SetTotalCount(v string) *ListScheduledPreloadJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListScheduledPreloadJobsResponseBodyJobs struct {
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CreatedAt     *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Domains       *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	ErrorInfo     *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	FailedFileOss *string `json:"FailedFileOss,omitempty" xml:"FailedFileOss,omitempty"`
	FileId        *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InsertWay     *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SiteId        *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TaskSubmitted *int32  `json:"TaskSubmitted,omitempty" xml:"TaskSubmitted,omitempty"`
	TaskType      *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UrlCount      *int32  `json:"UrlCount,omitempty" xml:"UrlCount,omitempty"`
	UrlSubmitted  *int32  `json:"UrlSubmitted,omitempty" xml:"UrlSubmitted,omitempty"`
}

func (s ListScheduledPreloadJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPreloadJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetAliUid(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.AliUid = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetCreatedAt(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.CreatedAt = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetDomains(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.Domains = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetErrorInfo(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.ErrorInfo = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetFailedFileOss(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.FailedFileOss = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetFileId(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.FileId = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetId(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetInsertWay(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.InsertWay = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetName(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.Name = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetSiteId(v int64) *ListScheduledPreloadJobsResponseBodyJobs {
	s.SiteId = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetTaskSubmitted(v int32) *ListScheduledPreloadJobsResponseBodyJobs {
	s.TaskSubmitted = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetTaskType(v string) *ListScheduledPreloadJobsResponseBodyJobs {
	s.TaskType = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetUrlCount(v int32) *ListScheduledPreloadJobsResponseBodyJobs {
	s.UrlCount = &v
	return s
}

func (s *ListScheduledPreloadJobsResponseBodyJobs) SetUrlSubmitted(v int32) *ListScheduledPreloadJobsResponseBodyJobs {
	s.UrlSubmitted = &v
	return s
}

type ListScheduledPreloadJobsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledPreloadJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduledPreloadJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPreloadJobsResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadJobsResponse) SetHeaders(v map[string]*string) *ListScheduledPreloadJobsResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledPreloadJobsResponse) SetStatusCode(v int32) *ListScheduledPreloadJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledPreloadJobsResponse) SetBody(v *ListScheduledPreloadJobsResponseBody) *ListScheduledPreloadJobsResponse {
	s.Body = v
	return s
}

type ListSiteDeliveryTasksRequest struct {
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListSiteDeliveryTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSiteDeliveryTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSiteDeliveryTasksRequest) SetBusinessType(v string) *ListSiteDeliveryTasksRequest {
	s.BusinessType = &v
	return s
}

func (s *ListSiteDeliveryTasksRequest) SetPageNumber(v int64) *ListSiteDeliveryTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSiteDeliveryTasksRequest) SetPageSize(v int64) *ListSiteDeliveryTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSiteDeliveryTasksRequest) SetSiteId(v int64) *ListSiteDeliveryTasksRequest {
	s.SiteId = &v
	return s
}

type ListSiteDeliveryTasksResponseBody struct {
	// example:
	//
	// 0
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*ListSiteDeliveryTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSiteDeliveryTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSiteDeliveryTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSiteDeliveryTasksResponseBody) SetPageNumber(v int32) *ListSiteDeliveryTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBody) SetPageSize(v int32) *ListSiteDeliveryTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBody) SetRequestId(v string) *ListSiteDeliveryTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBody) SetTasks(v []*ListSiteDeliveryTasksResponseBodyTasks) *ListSiteDeliveryTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListSiteDeliveryTasksResponseBody) SetTotalCount(v int32) *ListSiteDeliveryTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListSiteDeliveryTasksResponseBodyTasks struct {
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListSiteDeliveryTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListSiteDeliveryTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) SetBusinessType(v string) *ListSiteDeliveryTasksResponseBodyTasks {
	s.BusinessType = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) SetDataCenter(v string) *ListSiteDeliveryTasksResponseBodyTasks {
	s.DataCenter = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) SetDeliveryType(v string) *ListSiteDeliveryTasksResponseBodyTasks {
	s.DeliveryType = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) SetStatus(v string) *ListSiteDeliveryTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) SetTaskName(v string) *ListSiteDeliveryTasksResponseBodyTasks {
	s.TaskName = &v
	return s
}

type ListSiteDeliveryTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSiteDeliveryTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSiteDeliveryTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSiteDeliveryTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSiteDeliveryTasksResponse) SetHeaders(v map[string]*string) *ListSiteDeliveryTasksResponse {
	s.Headers = v
	return s
}

func (s *ListSiteDeliveryTasksResponse) SetStatusCode(v int32) *ListSiteDeliveryTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSiteDeliveryTasksResponse) SetBody(v *ListSiteDeliveryTasksResponseBody) *ListSiteDeliveryTasksResponse {
	s.Body = v
	return s
}

type ListSitesRequest struct {
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Coverage   *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// example:
	//
	// false
	OnlyEnterprise *bool `json:"OnlyEnterprise,omitempty" xml:"OnlyEnterprise,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlanSubscribeType *string `json:"PlanSubscribeType,omitempty" xml:"PlanSubscribeType,omitempty"`
	// example:
	//
	// rg-aekzd3styujvyei
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// example:
	//
	// fuzzy
	SiteSearchType *string `json:"SiteSearchType,omitempty" xml:"SiteSearchType,omitempty"`
	// example:
	//
	// pending
	Status    *string                      `json:"Status,omitempty" xml:"Status,omitempty"`
	TagFilter []*ListSitesRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Repeated"`
}

func (s ListSitesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSitesRequest) GoString() string {
	return s.String()
}

func (s *ListSitesRequest) SetAccessType(v string) *ListSitesRequest {
	s.AccessType = &v
	return s
}

func (s *ListSitesRequest) SetCoverage(v string) *ListSitesRequest {
	s.Coverage = &v
	return s
}

func (s *ListSitesRequest) SetOnlyEnterprise(v bool) *ListSitesRequest {
	s.OnlyEnterprise = &v
	return s
}

func (s *ListSitesRequest) SetPageNumber(v int32) *ListSitesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSitesRequest) SetPageSize(v int32) *ListSitesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSitesRequest) SetPlanSubscribeType(v string) *ListSitesRequest {
	s.PlanSubscribeType = &v
	return s
}

func (s *ListSitesRequest) SetResourceGroupId(v string) *ListSitesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSitesRequest) SetSiteName(v string) *ListSitesRequest {
	s.SiteName = &v
	return s
}

func (s *ListSitesRequest) SetSiteSearchType(v string) *ListSitesRequest {
	s.SiteSearchType = &v
	return s
}

func (s *ListSitesRequest) SetStatus(v string) *ListSitesRequest {
	s.Status = &v
	return s
}

func (s *ListSitesRequest) SetTagFilter(v []*ListSitesRequestTagFilter) *ListSitesRequest {
	s.TagFilter = v
	return s
}

type ListSitesRequestTagFilter struct {
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// aaa
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSitesRequestTagFilter) String() string {
	return tea.Prettify(s)
}

func (s ListSitesRequestTagFilter) GoString() string {
	return s.String()
}

func (s *ListSitesRequestTagFilter) SetKey(v string) *ListSitesRequestTagFilter {
	s.Key = &v
	return s
}

func (s *ListSitesRequestTagFilter) SetValue(v string) *ListSitesRequestTagFilter {
	s.Value = &v
	return s
}

type ListSitesShrinkRequest struct {
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Coverage   *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// example:
	//
	// false
	OnlyEnterprise *bool `json:"OnlyEnterprise,omitempty" xml:"OnlyEnterprise,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlanSubscribeType *string `json:"PlanSubscribeType,omitempty" xml:"PlanSubscribeType,omitempty"`
	// example:
	//
	// rg-aekzd3styujvyei
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// example:
	//
	// fuzzy
	SiteSearchType *string `json:"SiteSearchType,omitempty" xml:"SiteSearchType,omitempty"`
	// example:
	//
	// pending
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TagFilterShrink *string `json:"TagFilter,omitempty" xml:"TagFilter,omitempty"`
}

func (s ListSitesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSitesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSitesShrinkRequest) SetAccessType(v string) *ListSitesShrinkRequest {
	s.AccessType = &v
	return s
}

func (s *ListSitesShrinkRequest) SetCoverage(v string) *ListSitesShrinkRequest {
	s.Coverage = &v
	return s
}

func (s *ListSitesShrinkRequest) SetOnlyEnterprise(v bool) *ListSitesShrinkRequest {
	s.OnlyEnterprise = &v
	return s
}

func (s *ListSitesShrinkRequest) SetPageNumber(v int32) *ListSitesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSitesShrinkRequest) SetPageSize(v int32) *ListSitesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSitesShrinkRequest) SetPlanSubscribeType(v string) *ListSitesShrinkRequest {
	s.PlanSubscribeType = &v
	return s
}

func (s *ListSitesShrinkRequest) SetResourceGroupId(v string) *ListSitesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSitesShrinkRequest) SetSiteName(v string) *ListSitesShrinkRequest {
	s.SiteName = &v
	return s
}

func (s *ListSitesShrinkRequest) SetSiteSearchType(v string) *ListSitesShrinkRequest {
	s.SiteSearchType = &v
	return s
}

func (s *ListSitesShrinkRequest) SetStatus(v string) *ListSitesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListSitesShrinkRequest) SetTagFilterShrink(v string) *ListSitesShrinkRequest {
	s.TagFilterShrink = &v
	return s
}

type ListSitesResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sites     []*ListSitesResponseBodySites `json:"Sites,omitempty" xml:"Sites,omitempty" type:"Repeated"`
	// example:
	//
	// 40
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSitesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSitesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSitesResponseBody) SetPageNumber(v int32) *ListSitesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSitesResponseBody) SetPageSize(v int32) *ListSitesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSitesResponseBody) SetRequestId(v string) *ListSitesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSitesResponseBody) SetSites(v []*ListSitesResponseBodySites) *ListSitesResponseBody {
	s.Sites = v
	return s
}

func (s *ListSitesResponseBody) SetTotalCount(v int32) *ListSitesResponseBody {
	s.TotalCount = &v
	return s
}

type ListSitesResponseBodySites struct {
	// example:
	//
	// NS
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// example.cname.com
	CnameZone *string `json:"CnameZone,omitempty" xml:"CnameZone,omitempty"`
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// onBvtlmIyeXLbiDw81F9
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// male1-1.ialicdn.com,female1-1.ialicdn.com
	NameServerList *string `json:"NameServerList,omitempty" xml:"NameServerList,omitempty"`
	// example:
	//
	// plan-168656498****
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// example:
	//
	// normal
	PlanSpecName *string `json:"PlanSpecName,omitempty" xml:"PlanSpecName,omitempty"`
	// example:
	//
	// rg-aek26g6i6se6pna
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {"tag1":"value1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 2023-12-24T02:01:11Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// verify_d516cb3740f81f0cef77d162edd1****
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s ListSitesResponseBodySites) String() string {
	return tea.Prettify(s)
}

func (s ListSitesResponseBodySites) GoString() string {
	return s.String()
}

func (s *ListSitesResponseBodySites) SetAccessType(v string) *ListSitesResponseBodySites {
	s.AccessType = &v
	return s
}

func (s *ListSitesResponseBodySites) SetCnameZone(v string) *ListSitesResponseBodySites {
	s.CnameZone = &v
	return s
}

func (s *ListSitesResponseBodySites) SetCoverage(v string) *ListSitesResponseBodySites {
	s.Coverage = &v
	return s
}

func (s *ListSitesResponseBodySites) SetCreateTime(v string) *ListSitesResponseBodySites {
	s.CreateTime = &v
	return s
}

func (s *ListSitesResponseBodySites) SetInstanceId(v string) *ListSitesResponseBodySites {
	s.InstanceId = &v
	return s
}

func (s *ListSitesResponseBodySites) SetNameServerList(v string) *ListSitesResponseBodySites {
	s.NameServerList = &v
	return s
}

func (s *ListSitesResponseBodySites) SetPlanName(v string) *ListSitesResponseBodySites {
	s.PlanName = &v
	return s
}

func (s *ListSitesResponseBodySites) SetPlanSpecName(v string) *ListSitesResponseBodySites {
	s.PlanSpecName = &v
	return s
}

func (s *ListSitesResponseBodySites) SetResourceGroupId(v string) *ListSitesResponseBodySites {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSitesResponseBodySites) SetSiteId(v int64) *ListSitesResponseBodySites {
	s.SiteId = &v
	return s
}

func (s *ListSitesResponseBodySites) SetSiteName(v string) *ListSitesResponseBodySites {
	s.SiteName = &v
	return s
}

func (s *ListSitesResponseBodySites) SetStatus(v string) *ListSitesResponseBodySites {
	s.Status = &v
	return s
}

func (s *ListSitesResponseBodySites) SetTags(v map[string]interface{}) *ListSitesResponseBodySites {
	s.Tags = v
	return s
}

func (s *ListSitesResponseBodySites) SetUpdateTime(v string) *ListSitesResponseBodySites {
	s.UpdateTime = &v
	return s
}

func (s *ListSitesResponseBodySites) SetVerifyCode(v string) *ListSitesResponseBodySites {
	s.VerifyCode = &v
	return s
}

type ListSitesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSitesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSitesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSitesResponse) GoString() string {
	return s.String()
}

func (s *ListSitesResponse) SetHeaders(v map[string]*string) *ListSitesResponse {
	s.Headers = v
	return s
}

func (s *ListSitesResponse) SetStatusCode(v int32) *ListSitesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSitesResponse) SetBody(v *ListSitesResponseBody) *ListSitesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// example:
	//
	// 20
	MaxItem *int32 `json:"MaxItem,omitempty" xml:"MaxItem,omitempty"`
	// example:
	//
	// AAAAAZjtYxxxxxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 要创建并绑定标签的资源所在的地域ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源ID,最多 50个子项
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// site
	ResourceType  *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string                       `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxItem(v int32) *ListTagResourcesRequest {
	s.MaxItem = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetSecurityToken(v string) *ListTagResourcesRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// 标签键
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// example:
	//
	// AAAAAYwsxxxxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) SetTotalCount(v int32) *ListTagResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// example:
	//
	// example.com
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// site
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListUploadTasksRequest struct {
	// example:
	//
	// 2019-12-06T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 2018-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListUploadTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUploadTasksRequest) GoString() string {
	return s.String()
}

func (s *ListUploadTasksRequest) SetEndTime(v string) *ListUploadTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListUploadTasksRequest) SetSiteId(v int64) *ListUploadTasksRequest {
	s.SiteId = &v
	return s
}

func (s *ListUploadTasksRequest) SetStartTime(v string) *ListUploadTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListUploadTasksRequest) SetType(v string) *ListUploadTasksRequest {
	s.Type = &v
	return s
}

type ListUploadTasksResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6D****
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*ListUploadTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListUploadTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUploadTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListUploadTasksResponseBody) SetRequestId(v string) *ListUploadTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUploadTasksResponseBody) SetTasks(v []*ListUploadTasksResponseBodyTasks) *ListUploadTasksResponseBody {
	s.Tasks = v
	return s
}

type ListUploadTasksResponseBodyTasks struct {
	// example:
	//
	// 2023-07-26T01:56:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// invalid url
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// InvalidUrl,InvalidDomain
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 159253299357****
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	// example:
	//
	// purge_file_task
	UploadTaskName *string `json:"UploadTaskName,omitempty" xml:"UploadTaskName,omitempty"`
}

func (s ListUploadTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListUploadTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListUploadTasksResponseBodyTasks) SetCreateTime(v string) *ListUploadTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetDescription(v string) *ListUploadTasksResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetErrorCode(v string) *ListUploadTasksResponseBodyTasks {
	s.ErrorCode = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetStatus(v string) *ListUploadTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetType(v string) *ListUploadTasksResponseBodyTasks {
	s.Type = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetUploadId(v string) *ListUploadTasksResponseBodyTasks {
	s.UploadId = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetUploadTaskName(v string) *ListUploadTasksResponseBodyTasks {
	s.UploadTaskName = &v
	return s
}

type ListUploadTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUploadTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUploadTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUploadTasksResponse) GoString() string {
	return s.String()
}

func (s *ListUploadTasksResponse) SetHeaders(v map[string]*string) *ListUploadTasksResponse {
	s.Headers = v
	return s
}

func (s *ListUploadTasksResponse) SetStatusCode(v int32) *ListUploadTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUploadTasksResponse) SetBody(v *ListUploadTasksResponseBody) *ListUploadTasksResponse {
	s.Body = v
	return s
}

type ListUserDeliveryTasksRequest struct {
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserDeliveryTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserDeliveryTasksRequest) GoString() string {
	return s.String()
}

func (s *ListUserDeliveryTasksRequest) SetBusinessType(v string) *ListUserDeliveryTasksRequest {
	s.BusinessType = &v
	return s
}

func (s *ListUserDeliveryTasksRequest) SetPageNumber(v int64) *ListUserDeliveryTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserDeliveryTasksRequest) SetPageSize(v int64) *ListUserDeliveryTasksRequest {
	s.PageSize = &v
	return s
}

type ListUserDeliveryTasksResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1f94c47f-3a1a-4f69-8d6c-bfeee1b49aab
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*ListUserDeliveryTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserDeliveryTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserDeliveryTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDeliveryTasksResponseBody) SetPageNumber(v int32) *ListUserDeliveryTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBody) SetPageSize(v int32) *ListUserDeliveryTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBody) SetRequestId(v string) *ListUserDeliveryTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBody) SetTasks(v []*ListUserDeliveryTasksResponseBodyTasks) *ListUserDeliveryTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListUserDeliveryTasksResponseBody) SetTotalCount(v int32) *ListUserDeliveryTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListUserDeliveryTasksResponseBodyTasks struct {
	// example:
	//
	// dcdn_log_er
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// example:
	//
	// oss
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// testoss11
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListUserDeliveryTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListUserDeliveryTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListUserDeliveryTasksResponseBodyTasks) SetBusinessType(v string) *ListUserDeliveryTasksResponseBodyTasks {
	s.BusinessType = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBodyTasks) SetDataCenter(v string) *ListUserDeliveryTasksResponseBodyTasks {
	s.DataCenter = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBodyTasks) SetDeliveryType(v string) *ListUserDeliveryTasksResponseBodyTasks {
	s.DeliveryType = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBodyTasks) SetStatus(v string) *ListUserDeliveryTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBodyTasks) SetTaskName(v string) *ListUserDeliveryTasksResponseBodyTasks {
	s.TaskName = &v
	return s
}

type ListUserDeliveryTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserDeliveryTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserDeliveryTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserDeliveryTasksResponse) GoString() string {
	return s.String()
}

func (s *ListUserDeliveryTasksResponse) SetHeaders(v map[string]*string) *ListUserDeliveryTasksResponse {
	s.Headers = v
	return s
}

func (s *ListUserDeliveryTasksResponse) SetStatusCode(v int32) *ListUserDeliveryTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDeliveryTasksResponse) SetBody(v *ListUserDeliveryTasksResponseBody) *ListUserDeliveryTasksResponse {
	s.Body = v
	return s
}

type ListWafPhasesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListWafPhasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWafPhasesRequest) GoString() string {
	return s.String()
}

func (s *ListWafPhasesRequest) SetSiteId(v int64) *ListWafPhasesRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafPhasesRequest) SetSiteVersion(v int32) *ListWafPhasesRequest {
	s.SiteVersion = &v
	return s
}

type ListWafPhasesResponseBody struct {
	Phases []*ListWafPhasesResponseBodyPhases `json:"Phases,omitempty" xml:"Phases,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWafPhasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWafPhasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafPhasesResponseBody) SetPhases(v []*ListWafPhasesResponseBodyPhases) *ListWafPhasesResponseBody {
	s.Phases = v
	return s
}

func (s *ListWafPhasesResponseBody) SetRequestId(v string) *ListWafPhasesResponseBody {
	s.RequestId = &v
	return s
}

type ListWafPhasesResponseBodyPhases struct {
	Phase    *string                                    `json:"Phase,omitempty" xml:"Phase,omitempty"`
	Rulesets []*ListWafPhasesResponseBodyPhasesRulesets `json:"Rulesets,omitempty" xml:"Rulesets,omitempty" type:"Repeated"`
}

func (s ListWafPhasesResponseBodyPhases) String() string {
	return tea.Prettify(s)
}

func (s ListWafPhasesResponseBodyPhases) GoString() string {
	return s.String()
}

func (s *ListWafPhasesResponseBodyPhases) SetPhase(v string) *ListWafPhasesResponseBodyPhases {
	s.Phase = &v
	return s
}

func (s *ListWafPhasesResponseBodyPhases) SetRulesets(v []*ListWafPhasesResponseBodyPhasesRulesets) *ListWafPhasesResponseBodyPhases {
	s.Rulesets = v
	return s
}

type ListWafPhasesResponseBodyPhasesRulesets struct {
	Id     *int64              `json:"Id,omitempty" xml:"Id,omitempty"`
	Name   *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	Rules  []*WafRuleConfig    `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
}

func (s ListWafPhasesResponseBodyPhasesRulesets) String() string {
	return tea.Prettify(s)
}

func (s ListWafPhasesResponseBodyPhasesRulesets) GoString() string {
	return s.String()
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) SetId(v int64) *ListWafPhasesResponseBodyPhasesRulesets {
	s.Id = &v
	return s
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) SetName(v string) *ListWafPhasesResponseBodyPhasesRulesets {
	s.Name = &v
	return s
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) SetRules(v []*WafRuleConfig) *ListWafPhasesResponseBodyPhasesRulesets {
	s.Rules = v
	return s
}

func (s *ListWafPhasesResponseBodyPhasesRulesets) SetShared(v *WafBatchRuleShared) *ListWafPhasesResponseBodyPhasesRulesets {
	s.Shared = v
	return s
}

type ListWafPhasesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafPhasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafPhasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWafPhasesResponse) GoString() string {
	return s.String()
}

func (s *ListWafPhasesResponse) SetHeaders(v map[string]*string) *ListWafPhasesResponse {
	s.Headers = v
	return s
}

func (s *ListWafPhasesResponse) SetStatusCode(v int32) *ListWafPhasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafPhasesResponse) SetBody(v *ListWafPhasesResponseBody) *ListWafPhasesResponse {
	s.Body = v
	return s
}

type ListWafRulesRequest struct {
	// example:
	//
	// http_custom
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// http_custom
	QueryArgs *ListWafRulesRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListWafRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWafRulesRequest) SetPageNumber(v int32) *ListWafRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesRequest) SetPageSize(v int32) *ListWafRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesRequest) SetPhase(v string) *ListWafRulesRequest {
	s.Phase = &v
	return s
}

func (s *ListWafRulesRequest) SetQueryArgs(v *ListWafRulesRequestQueryArgs) *ListWafRulesRequest {
	s.QueryArgs = v
	return s
}

func (s *ListWafRulesRequest) SetSiteId(v int64) *ListWafRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafRulesRequest) SetSiteVersion(v int32) *ListWafRulesRequest {
	s.SiteVersion = &v
	return s
}

type ListWafRulesRequestQueryArgs struct {
	Desc *bool `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// example
	IdNameLike *string `json:"IdNameLike,omitempty" xml:"IdNameLike,omitempty"`
	// example:
	//
	// example
	NameLike *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	// example:
	//
	// position
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWafRulesRequestQueryArgs) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListWafRulesRequestQueryArgs) SetDesc(v bool) *ListWafRulesRequestQueryArgs {
	s.Desc = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetId(v int64) *ListWafRulesRequestQueryArgs {
	s.Id = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetIdNameLike(v string) *ListWafRulesRequestQueryArgs {
	s.IdNameLike = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetNameLike(v string) *ListWafRulesRequestQueryArgs {
	s.NameLike = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetOrderBy(v string) *ListWafRulesRequestQueryArgs {
	s.OrderBy = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetRulesetId(v int64) *ListWafRulesRequestQueryArgs {
	s.RulesetId = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetStatus(v string) *ListWafRulesRequestQueryArgs {
	s.Status = &v
	return s
}

type ListWafRulesShrinkRequest struct {
	// example:
	//
	// http_custom
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// http_custom
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListWafRulesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWafRulesShrinkRequest) SetPageNumber(v int32) *ListWafRulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetPageSize(v int32) *ListWafRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetPhase(v string) *ListWafRulesShrinkRequest {
	s.Phase = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetQueryArgsShrink(v string) *ListWafRulesShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetSiteId(v int64) *ListWafRulesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetSiteVersion(v int32) *ListWafRulesShrinkRequest {
	s.SiteVersion = &v
	return s
}

type ListWafRulesResponseBody struct {
	// example:
	//
	// 10
	InstanceUsage *int64 `json:"InstanceUsage,omitempty" xml:"InstanceUsage,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*ListWafRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	SiteUsage *int64 `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWafRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafRulesResponseBody) SetInstanceUsage(v int64) *ListWafRulesResponseBody {
	s.InstanceUsage = &v
	return s
}

func (s *ListWafRulesResponseBody) SetPageNumber(v int32) *ListWafRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesResponseBody) SetPageSize(v int32) *ListWafRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesResponseBody) SetRequestId(v string) *ListWafRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWafRulesResponseBody) SetRules(v []*ListWafRulesResponseBodyRules) *ListWafRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListWafRulesResponseBody) SetSiteUsage(v int64) *ListWafRulesResponseBody {
	s.SiteUsage = &v
	return s
}

func (s *ListWafRulesResponseBody) SetTotalCount(v int64) *ListWafRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListWafRulesResponseBodyRules struct {
	// example:
	//
	// deny
	Action                *string        `json:"Action,omitempty" xml:"Action,omitempty"`
	CharacteristicsFields []*string      `json:"CharacteristicsFields,omitempty" xml:"CharacteristicsFields,omitempty" type:"Repeated"`
	Config                *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	Fields                []*string      `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// example:
	//
	// part
	Skip *string `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// example:
	//
	// on
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Timer  *WafTimer `json:"Timer,omitempty" xml:"Timer,omitempty"`
	// example:
	//
	// http_custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListWafRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListWafRulesResponseBodyRules) SetAction(v string) *ListWafRulesResponseBodyRules {
	s.Action = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetCharacteristicsFields(v []*string) *ListWafRulesResponseBodyRules {
	s.CharacteristicsFields = v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetConfig(v *WafRuleConfig) *ListWafRulesResponseBodyRules {
	s.Config = v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetFields(v []*string) *ListWafRulesResponseBodyRules {
	s.Fields = v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetId(v int64) *ListWafRulesResponseBodyRules {
	s.Id = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetName(v string) *ListWafRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetPhase(v string) *ListWafRulesResponseBodyRules {
	s.Phase = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetPosition(v int64) *ListWafRulesResponseBodyRules {
	s.Position = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetRulesetId(v int64) *ListWafRulesResponseBodyRules {
	s.RulesetId = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetSkip(v string) *ListWafRulesResponseBodyRules {
	s.Skip = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetStatus(v string) *ListWafRulesResponseBodyRules {
	s.Status = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetTags(v []*string) *ListWafRulesResponseBodyRules {
	s.Tags = v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetTimer(v *WafTimer) *ListWafRulesResponseBodyRules {
	s.Timer = v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetType(v string) *ListWafRulesResponseBodyRules {
	s.Type = &v
	return s
}

func (s *ListWafRulesResponseBodyRules) SetUpdateTime(v string) *ListWafRulesResponseBodyRules {
	s.UpdateTime = &v
	return s
}

type ListWafRulesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWafRulesResponse) SetHeaders(v map[string]*string) *ListWafRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWafRulesResponse) SetStatusCode(v int32) *ListWafRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafRulesResponse) SetBody(v *ListWafRulesResponseBody) *ListWafRulesResponse {
	s.Body = v
	return s
}

type ListWafRulesetsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// http_bot
	QueryArgs *ListWafRulesetsRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListWafRulesetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesetsRequest) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsRequest) SetPageNumber(v int32) *ListWafRulesetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesetsRequest) SetPageSize(v int32) *ListWafRulesetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesetsRequest) SetPhase(v string) *ListWafRulesetsRequest {
	s.Phase = &v
	return s
}

func (s *ListWafRulesetsRequest) SetQueryArgs(v *ListWafRulesetsRequestQueryArgs) *ListWafRulesetsRequest {
	s.QueryArgs = v
	return s
}

func (s *ListWafRulesetsRequest) SetSiteId(v int64) *ListWafRulesetsRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafRulesetsRequest) SetSiteVersion(v int32) *ListWafRulesetsRequest {
	s.SiteVersion = &v
	return s
}

type ListWafRulesetsRequestQueryArgs struct {
	// example:
	//
	// example
	AnyLike *string `json:"AnyLike,omitempty" xml:"AnyLike,omitempty"`
	Desc    *bool   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// example
	NameLike *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	// example:
	//
	// id
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
}

func (s ListWafRulesetsRequestQueryArgs) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesetsRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsRequestQueryArgs) SetAnyLike(v string) *ListWafRulesetsRequestQueryArgs {
	s.AnyLike = &v
	return s
}

func (s *ListWafRulesetsRequestQueryArgs) SetDesc(v bool) *ListWafRulesetsRequestQueryArgs {
	s.Desc = &v
	return s
}

func (s *ListWafRulesetsRequestQueryArgs) SetNameLike(v string) *ListWafRulesetsRequestQueryArgs {
	s.NameLike = &v
	return s
}

func (s *ListWafRulesetsRequestQueryArgs) SetOrderBy(v string) *ListWafRulesetsRequestQueryArgs {
	s.OrderBy = &v
	return s
}

type ListWafRulesetsShrinkRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// http_bot
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListWafRulesetsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsShrinkRequest) SetPageNumber(v int32) *ListWafRulesetsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) SetPageSize(v int32) *ListWafRulesetsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) SetPhase(v string) *ListWafRulesetsShrinkRequest {
	s.Phase = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) SetQueryArgsShrink(v string) *ListWafRulesetsShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) SetSiteId(v int64) *ListWafRulesetsShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) SetSiteVersion(v int32) *ListWafRulesetsShrinkRequest {
	s.SiteVersion = &v
	return s
}

type ListWafRulesetsResponseBody struct {
	// example:
	//
	// 10
	InstanceUsage *int64 `json:"InstanceUsage,omitempty" xml:"InstanceUsage,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rulesets  []*ListWafRulesetsResponseBodyRulesets `json:"Rulesets,omitempty" xml:"Rulesets,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	SiteUsage *int64 `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty"`
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWafRulesetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsResponseBody) SetInstanceUsage(v int64) *ListWafRulesetsResponseBody {
	s.InstanceUsage = &v
	return s
}

func (s *ListWafRulesetsResponseBody) SetPageNumber(v int32) *ListWafRulesetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesetsResponseBody) SetPageSize(v int32) *ListWafRulesetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesetsResponseBody) SetRequestId(v string) *ListWafRulesetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWafRulesetsResponseBody) SetRulesets(v []*ListWafRulesetsResponseBodyRulesets) *ListWafRulesetsResponseBody {
	s.Rulesets = v
	return s
}

func (s *ListWafRulesetsResponseBody) SetSiteUsage(v int64) *ListWafRulesetsResponseBody {
	s.SiteUsage = &v
	return s
}

func (s *ListWafRulesetsResponseBody) SetTotalCount(v int64) *ListWafRulesetsResponseBody {
	s.TotalCount = &v
	return s
}

type ListWafRulesetsResponseBodyRulesets struct {
	Fields []*string `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// web
	Target *string   `json:"Target,omitempty" xml:"Target,omitempty"`
	Types  []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListWafRulesetsResponseBodyRulesets) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesetsResponseBodyRulesets) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsResponseBodyRulesets) SetFields(v []*string) *ListWafRulesetsResponseBodyRulesets {
	s.Fields = v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetId(v int64) *ListWafRulesetsResponseBodyRulesets {
	s.Id = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetName(v string) *ListWafRulesetsResponseBodyRulesets {
	s.Name = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetPhase(v string) *ListWafRulesetsResponseBodyRulesets {
	s.Phase = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetStatus(v string) *ListWafRulesetsResponseBodyRulesets {
	s.Status = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetTarget(v string) *ListWafRulesetsResponseBodyRulesets {
	s.Target = &v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetTypes(v []*string) *ListWafRulesetsResponseBodyRulesets {
	s.Types = v
	return s
}

func (s *ListWafRulesetsResponseBodyRulesets) SetUpdateTime(v string) *ListWafRulesetsResponseBodyRulesets {
	s.UpdateTime = &v
	return s
}

type ListWafRulesetsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafRulesetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafRulesetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWafRulesetsResponse) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsResponse) SetHeaders(v map[string]*string) *ListWafRulesetsResponse {
	s.Headers = v
	return s
}

func (s *ListWafRulesetsResponse) SetStatusCode(v int32) *ListWafRulesetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafRulesetsResponse) SetBody(v *ListWafRulesetsResponseBody) *ListWafRulesetsResponse {
	s.Body = v
	return s
}

type ListWafTemplateRulesRequest struct {
	// example:
	//
	// http_anti_scan
	Phase     *string                               `json:"Phase,omitempty" xml:"Phase,omitempty"`
	QueryArgs *ListWafTemplateRulesRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
}

func (s ListWafTemplateRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWafTemplateRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesRequest) SetPhase(v string) *ListWafTemplateRulesRequest {
	s.Phase = &v
	return s
}

func (s *ListWafTemplateRulesRequest) SetQueryArgs(v *ListWafTemplateRulesRequestQueryArgs) *ListWafTemplateRulesRequest {
	s.QueryArgs = v
	return s
}

type ListWafTemplateRulesRequestQueryArgs struct {
	// example:
	//
	// http_directory_traversal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWafTemplateRulesRequestQueryArgs) String() string {
	return tea.Prettify(s)
}

func (s ListWafTemplateRulesRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesRequestQueryArgs) SetType(v string) *ListWafTemplateRulesRequestQueryArgs {
	s.Type = &v
	return s
}

type ListWafTemplateRulesShrinkRequest struct {
	// example:
	//
	// http_anti_scan
	Phase           *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s ListWafTemplateRulesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWafTemplateRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesShrinkRequest) SetPhase(v string) *ListWafTemplateRulesShrinkRequest {
	s.Phase = &v
	return s
}

func (s *ListWafTemplateRulesShrinkRequest) SetQueryArgsShrink(v string) *ListWafTemplateRulesShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

type ListWafTemplateRulesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*ListWafTemplateRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ListWafTemplateRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWafTemplateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesResponseBody) SetRequestId(v string) *ListWafTemplateRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWafTemplateRulesResponseBody) SetRules(v []*ListWafTemplateRulesResponseBodyRules) *ListWafTemplateRulesResponseBody {
	s.Rules = v
	return s
}

type ListWafTemplateRulesResponseBodyRules struct {
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// HTTP Directory Traversal Rule [Template]
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// http_anti_scan
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// http_directory_traversal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWafTemplateRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s ListWafTemplateRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesResponseBodyRules) SetConfig(v *WafRuleConfig) *ListWafTemplateRulesResponseBodyRules {
	s.Config = v
	return s
}

func (s *ListWafTemplateRulesResponseBodyRules) SetName(v string) *ListWafTemplateRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *ListWafTemplateRulesResponseBodyRules) SetPhase(v string) *ListWafTemplateRulesResponseBodyRules {
	s.Phase = &v
	return s
}

func (s *ListWafTemplateRulesResponseBodyRules) SetStatus(v string) *ListWafTemplateRulesResponseBodyRules {
	s.Status = &v
	return s
}

func (s *ListWafTemplateRulesResponseBodyRules) SetType(v string) *ListWafTemplateRulesResponseBodyRules {
	s.Type = &v
	return s
}

type ListWafTemplateRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafTemplateRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafTemplateRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWafTemplateRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesResponse) SetHeaders(v map[string]*string) *ListWafTemplateRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWafTemplateRulesResponse) SetStatusCode(v int32) *ListWafTemplateRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafTemplateRulesResponse) SetBody(v *ListWafTemplateRulesResponseBody) *ListWafTemplateRulesResponse {
	s.Body = v
	return s
}

type ListWafUsageOfRulesRequest struct {
	// example:
	//
	// http_anti_scan
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// ListWafUsageOfRules
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListWafUsageOfRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWafUsageOfRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWafUsageOfRulesRequest) SetPhase(v string) *ListWafUsageOfRulesRequest {
	s.Phase = &v
	return s
}

func (s *ListWafUsageOfRulesRequest) SetSiteId(v int64) *ListWafUsageOfRulesRequest {
	s.SiteId = &v
	return s
}

type ListWafUsageOfRulesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sites     []*ListWafUsageOfRulesResponseBodySites `json:"Sites,omitempty" xml:"Sites,omitempty" type:"Repeated"`
}

func (s ListWafUsageOfRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWafUsageOfRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafUsageOfRulesResponseBody) SetRequestId(v string) *ListWafUsageOfRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWafUsageOfRulesResponseBody) SetSites(v []*ListWafUsageOfRulesResponseBodySites) *ListWafUsageOfRulesResponseBody {
	s.Sites = v
	return s
}

type ListWafUsageOfRulesResponseBodySites struct {
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListWafUsageOfRulesResponseBodySites) String() string {
	return tea.Prettify(s)
}

func (s ListWafUsageOfRulesResponseBodySites) GoString() string {
	return s.String()
}

func (s *ListWafUsageOfRulesResponseBodySites) SetId(v int64) *ListWafUsageOfRulesResponseBodySites {
	s.Id = &v
	return s
}

func (s *ListWafUsageOfRulesResponseBodySites) SetName(v string) *ListWafUsageOfRulesResponseBodySites {
	s.Name = &v
	return s
}

func (s *ListWafUsageOfRulesResponseBodySites) SetUsage(v int64) *ListWafUsageOfRulesResponseBodySites {
	s.Usage = &v
	return s
}

type ListWafUsageOfRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafUsageOfRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafUsageOfRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWafUsageOfRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWafUsageOfRulesResponse) SetHeaders(v map[string]*string) *ListWafUsageOfRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWafUsageOfRulesResponse) SetStatusCode(v int32) *ListWafUsageOfRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafUsageOfRulesResponse) SetBody(v *ListWafUsageOfRulesResponseBody) *ListWafUsageOfRulesResponse {
	s.Body = v
	return s
}

type ListWaitingRoomEventsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 89677721098****
	WaitingRoomEventId *int64 `json:"WaitingRoomEventId,omitempty" xml:"WaitingRoomEventId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6a51d5bc6460887abd129****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
}

func (s ListWaitingRoomEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomEventsRequest) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomEventsRequest) SetSiteId(v int64) *ListWaitingRoomEventsRequest {
	s.SiteId = &v
	return s
}

func (s *ListWaitingRoomEventsRequest) SetWaitingRoomEventId(v int64) *ListWaitingRoomEventsRequest {
	s.WaitingRoomEventId = &v
	return s
}

func (s *ListWaitingRoomEventsRequest) SetWaitingRoomId(v string) *ListWaitingRoomEventsRequest {
	s.WaitingRoomId = &v
	return s
}

type ListWaitingRoomEventsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// f3c3700a-4c0f-4a24-b576-fd7dbf9e7c55
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WaitingRoomEvents []*ListWaitingRoomEventsResponseBodyWaitingRoomEvents `json:"WaitingRoomEvents,omitempty" xml:"WaitingRoomEvents,omitempty" type:"Repeated"`
}

func (s ListWaitingRoomEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomEventsResponseBody) SetRequestId(v string) *ListWaitingRoomEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBody) SetWaitingRoomEvents(v []*ListWaitingRoomEventsResponseBodyWaitingRoomEvents) *ListWaitingRoomEventsResponseBody {
	s.WaitingRoomEvents = v
	return s
}

type ListWaitingRoomEventsResponseBodyWaitingRoomEvents struct {
	// example:
	//
	// html-yets-maqi1111
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	// example:
	//
	// terraform-example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// off
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 1719814497
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// off
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// example:
	//
	// zhcn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 11
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// example:
	//
	// on
	PreQueueEnable *string `json:"PreQueueEnable,omitempty" xml:"PreQueueEnable,omitempty"`
	// example:
	//
	// 1719814097
	PreQueueStartTime *string `json:"PreQueueStartTime,omitempty" xml:"PreQueueStartTime,omitempty"`
	// example:
	//
	// fifo
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// example:
	//
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// example:
	//
	// on
	RandomPreQueueEnable *string `json:"RandomPreQueueEnable,omitempty" xml:"RandomPreQueueEnable,omitempty"`
	// example:
	//
	// 3
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// example:
	//
	// 1719814398
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 22
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// example:
	//
	// 89677721098****
	WaitingRoomEventId *int64 `json:"WaitingRoomEventId,omitempty" xml:"WaitingRoomEventId,omitempty"`
	// example:
	//
	// 5c938a045c9ca46607163d34966****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// example:
	//
	// custom
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s ListWaitingRoomEventsResponseBodyWaitingRoomEvents) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomEventsResponseBodyWaitingRoomEvents) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetCustomPageHtml(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.CustomPageHtml = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetDescription(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.Description = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetDisableSessionRenewalEnable(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetEnable(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.Enable = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetEndTime(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.EndTime = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetJsonResponseEnable(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.JsonResponseEnable = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetLanguage(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.Language = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetName(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.Name = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetNewUsersPerMinute(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.NewUsersPerMinute = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetPreQueueEnable(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.PreQueueEnable = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetPreQueueStartTime(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.PreQueueStartTime = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetQueuingMethod(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.QueuingMethod = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetQueuingStatusCode(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.QueuingStatusCode = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetRandomPreQueueEnable(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.RandomPreQueueEnable = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetSessionDuration(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.SessionDuration = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetStartTime(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.StartTime = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetTotalActiveUsers(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.TotalActiveUsers = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetWaitingRoomEventId(v int64) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.WaitingRoomEventId = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetWaitingRoomId(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.WaitingRoomId = &v
	return s
}

func (s *ListWaitingRoomEventsResponseBodyWaitingRoomEvents) SetWaitingRoomType(v string) *ListWaitingRoomEventsResponseBodyWaitingRoomEvents {
	s.WaitingRoomType = &v
	return s
}

type ListWaitingRoomEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWaitingRoomEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWaitingRoomEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomEventsResponse) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomEventsResponse) SetHeaders(v map[string]*string) *ListWaitingRoomEventsResponse {
	s.Headers = v
	return s
}

func (s *ListWaitingRoomEventsResponse) SetStatusCode(v int32) *ListWaitingRoomEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWaitingRoomEventsResponse) SetBody(v *ListWaitingRoomEventsResponseBody) *ListWaitingRoomEventsResponse {
	s.Body = v
	return s
}

type ListWaitingRoomRulesRequest struct {
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6a51d5bc6460887abd129****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// example:
	//
	// 37286782688****
	WaitingRoomRuleId *int64 `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s ListWaitingRoomRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomRulesRequest) SetRuleName(v string) *ListWaitingRoomRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListWaitingRoomRulesRequest) SetSiteId(v int64) *ListWaitingRoomRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListWaitingRoomRulesRequest) SetWaitingRoomId(v string) *ListWaitingRoomRulesRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *ListWaitingRoomRulesRequest) SetWaitingRoomRuleId(v int64) *ListWaitingRoomRulesRequest {
	s.WaitingRoomRuleId = &v
	return s
}

type ListWaitingRoomRulesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A123425345
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WaitingRoomRules []*ListWaitingRoomRulesResponseBodyWaitingRoomRules `json:"WaitingRoomRules,omitempty" xml:"WaitingRoomRules,omitempty" type:"Repeated"`
}

func (s ListWaitingRoomRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomRulesResponseBody) SetRequestId(v string) *ListWaitingRoomRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBody) SetWaitingRoomRules(v []*ListWaitingRoomRulesResponseBodyWaitingRoomRules) *ListWaitingRoomRulesResponseBody {
	s.WaitingRoomRules = v
	return s
}

type ListWaitingRoomRulesResponseBodyWaitingRoomRules struct {
	// example:
	//
	// (http.request.uri.path.file_name eq \\"jpg\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// example:
	//
	// ip
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 37286782688****
	WaitingRoomRuleId *int64 `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s ListWaitingRoomRulesResponseBodyWaitingRoomRules) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomRulesResponseBodyWaitingRoomRules) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetRule(v string) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.Rule = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetRuleEnable(v string) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.RuleEnable = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetRuleName(v string) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.RuleName = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetWaitingRoomRuleId(v int64) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.WaitingRoomRuleId = &v
	return s
}

type ListWaitingRoomRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWaitingRoomRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWaitingRoomRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomRulesResponse) SetHeaders(v map[string]*string) *ListWaitingRoomRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWaitingRoomRulesResponse) SetStatusCode(v int32) *ListWaitingRoomRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWaitingRoomRulesResponse) SetBody(v *ListWaitingRoomRulesResponseBody) *ListWaitingRoomRulesResponse {
	s.Body = v
	return s
}

type ListWaitingRoomsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 120876698010528
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 6a51d5bc6460887abd1291dc7d4d****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
}

func (s ListWaitingRoomsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomsRequest) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomsRequest) SetSiteId(v int64) *ListWaitingRoomsRequest {
	s.SiteId = &v
	return s
}

func (s *ListWaitingRoomsRequest) SetWaitingRoomId(v string) *ListWaitingRoomsRequest {
	s.WaitingRoomId = &v
	return s
}

type ListWaitingRoomsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WaitingRooms []*ListWaitingRoomsResponseBodyWaitingRooms `json:"WaitingRooms,omitempty" xml:"WaitingRooms,omitempty" type:"Repeated"`
}

func (s ListWaitingRoomsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomsResponseBody) SetRequestId(v string) *ListWaitingRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWaitingRoomsResponseBody) SetWaitingRooms(v []*ListWaitingRoomsResponseBodyWaitingRooms) *ListWaitingRoomsResponseBody {
	s.WaitingRooms = v
	return s
}

type ListWaitingRoomsResponseBodyWaitingRooms struct {
	// example:
	//
	// __aliwaitingroom_example
	CookieName     *string `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// on
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// example:
	//
	// on
	Enable          *string                                                    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HostNameAndPath []*ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty" type:"Repeated"`
	// example:
	//
	// on
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// example:
	//
	// zhcn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 200
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// example:
	//
	// on
	QueueAllEnable *string `json:"QueueAllEnable,omitempty" xml:"QueueAllEnable,omitempty"`
	// example:
	//
	// random
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// example:
	//
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// example:
	//
	// 3600
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// example:
	//
	// 300
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// example:
	//
	// 6a51d5bc6460887abd1291dc7d4d****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// example:
	//
	// default
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s ListWaitingRoomsResponseBodyWaitingRooms) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomsResponseBodyWaitingRooms) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetCookieName(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.CookieName = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetCustomPageHtml(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.CustomPageHtml = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetDescription(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.Description = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetDisableSessionRenewalEnable(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetEnable(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.Enable = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetHostNameAndPath(v []*ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.HostNameAndPath = v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetJsonResponseEnable(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.JsonResponseEnable = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetLanguage(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.Language = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetName(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.Name = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetNewUsersPerMinute(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.NewUsersPerMinute = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetQueueAllEnable(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.QueueAllEnable = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetQueuingMethod(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.QueuingMethod = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetQueuingStatusCode(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.QueuingStatusCode = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetSessionDuration(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.SessionDuration = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetTotalActiveUsers(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.TotalActiveUsers = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetWaitingRoomId(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.WaitingRoomId = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetWaitingRoomType(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.WaitingRoomType = &v
	return s
}

type ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath struct {
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// test.
	Subdomain *string `json:"Subdomain,omitempty" xml:"Subdomain,omitempty"`
}

func (s ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) SetDomain(v string) *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath {
	s.Domain = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) SetPath(v string) *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath {
	s.Path = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) SetSubdomain(v string) *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath {
	s.Subdomain = &v
	return s
}

type ListWaitingRoomsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWaitingRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWaitingRoomsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWaitingRoomsResponse) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomsResponse) SetHeaders(v map[string]*string) *ListWaitingRoomsResponse {
	s.Headers = v
	return s
}

func (s *ListWaitingRoomsResponse) SetStatusCode(v int32) *ListWaitingRoomsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWaitingRoomsResponse) SetBody(v *ListWaitingRoomsResponseBody) *ListWaitingRoomsResponse {
	s.Body = v
	return s
}

type PreloadCachesRequest struct {
	Content []*string          `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s PreloadCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s PreloadCachesRequest) GoString() string {
	return s.String()
}

func (s *PreloadCachesRequest) SetContent(v []*string) *PreloadCachesRequest {
	s.Content = v
	return s
}

func (s *PreloadCachesRequest) SetHeaders(v map[string]*string) *PreloadCachesRequest {
	s.Headers = v
	return s
}

func (s *PreloadCachesRequest) SetSiteId(v int64) *PreloadCachesRequest {
	s.SiteId = &v
	return s
}

type PreloadCachesShrinkRequest struct {
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	HeadersShrink *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s PreloadCachesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PreloadCachesShrinkRequest) GoString() string {
	return s.String()
}

func (s *PreloadCachesShrinkRequest) SetContentShrink(v string) *PreloadCachesShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *PreloadCachesShrinkRequest) SetHeadersShrink(v string) *PreloadCachesShrinkRequest {
	s.HeadersShrink = &v
	return s
}

func (s *PreloadCachesShrinkRequest) SetSiteId(v int64) *PreloadCachesShrinkRequest {
	s.SiteId = &v
	return s
}

type PreloadCachesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 16401427840
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PreloadCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreloadCachesResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadCachesResponseBody) SetRequestId(v string) *PreloadCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreloadCachesResponseBody) SetTaskId(v string) *PreloadCachesResponseBody {
	s.TaskId = &v
	return s
}

type PreloadCachesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreloadCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreloadCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s PreloadCachesResponse) GoString() string {
	return s.String()
}

func (s *PreloadCachesResponse) SetHeaders(v map[string]*string) *PreloadCachesResponse {
	s.Headers = v
	return s
}

func (s *PreloadCachesResponse) SetStatusCode(v int32) *PreloadCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *PreloadCachesResponse) SetBody(v *PreloadCachesResponseBody) *PreloadCachesResponse {
	s.Body = v
	return s
}

type PurgeCachesRequest struct {
	Content *PurgeCachesRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// true
	EdgeComputePurge *bool `json:"EdgeComputePurge,omitempty" xml:"EdgeComputePurge,omitempty"`
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PurgeCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s PurgeCachesRequest) GoString() string {
	return s.String()
}

func (s *PurgeCachesRequest) SetContent(v *PurgeCachesRequestContent) *PurgeCachesRequest {
	s.Content = v
	return s
}

func (s *PurgeCachesRequest) SetEdgeComputePurge(v bool) *PurgeCachesRequest {
	s.EdgeComputePurge = &v
	return s
}

func (s *PurgeCachesRequest) SetForce(v bool) *PurgeCachesRequest {
	s.Force = &v
	return s
}

func (s *PurgeCachesRequest) SetSiteId(v int64) *PurgeCachesRequest {
	s.SiteId = &v
	return s
}

func (s *PurgeCachesRequest) SetType(v string) *PurgeCachesRequest {
	s.Type = &v
	return s
}

type PurgeCachesRequestContent struct {
	CacheTags    []*string     `json:"CacheTags,omitempty" xml:"CacheTags,omitempty" type:"Repeated"`
	Directories  []*string     `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	Files        []interface{} `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	Hostnames    []*string     `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	IgnoreParams []*string     `json:"IgnoreParams,omitempty" xml:"IgnoreParams,omitempty" type:"Repeated"`
	// example:
	//
	// true
	PurgeAll *bool `json:"PurgeAll,omitempty" xml:"PurgeAll,omitempty"`
}

func (s PurgeCachesRequestContent) String() string {
	return tea.Prettify(s)
}

func (s PurgeCachesRequestContent) GoString() string {
	return s.String()
}

func (s *PurgeCachesRequestContent) SetCacheTags(v []*string) *PurgeCachesRequestContent {
	s.CacheTags = v
	return s
}

func (s *PurgeCachesRequestContent) SetDirectories(v []*string) *PurgeCachesRequestContent {
	s.Directories = v
	return s
}

func (s *PurgeCachesRequestContent) SetFiles(v []interface{}) *PurgeCachesRequestContent {
	s.Files = v
	return s
}

func (s *PurgeCachesRequestContent) SetHostnames(v []*string) *PurgeCachesRequestContent {
	s.Hostnames = v
	return s
}

func (s *PurgeCachesRequestContent) SetIgnoreParams(v []*string) *PurgeCachesRequestContent {
	s.IgnoreParams = v
	return s
}

func (s *PurgeCachesRequestContent) SetPurgeAll(v bool) *PurgeCachesRequestContent {
	s.PurgeAll = &v
	return s
}

type PurgeCachesShrinkRequest struct {
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// true
	EdgeComputePurge *bool `json:"EdgeComputePurge,omitempty" xml:"EdgeComputePurge,omitempty"`
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PurgeCachesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PurgeCachesShrinkRequest) GoString() string {
	return s.String()
}

func (s *PurgeCachesShrinkRequest) SetContentShrink(v string) *PurgeCachesShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *PurgeCachesShrinkRequest) SetEdgeComputePurge(v bool) *PurgeCachesShrinkRequest {
	s.EdgeComputePurge = &v
	return s
}

func (s *PurgeCachesShrinkRequest) SetForce(v bool) *PurgeCachesShrinkRequest {
	s.Force = &v
	return s
}

func (s *PurgeCachesShrinkRequest) SetSiteId(v int64) *PurgeCachesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *PurgeCachesShrinkRequest) SetType(v string) *PurgeCachesShrinkRequest {
	s.Type = &v
	return s
}

type PurgeCachesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// ET5BF670-09D5-4D0B-BEBY-D96A2A528000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15940956620
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PurgeCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PurgeCachesResponseBody) GoString() string {
	return s.String()
}

func (s *PurgeCachesResponseBody) SetRequestId(v string) *PurgeCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurgeCachesResponseBody) SetTaskId(v string) *PurgeCachesResponseBody {
	s.TaskId = &v
	return s
}

type PurgeCachesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurgeCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurgeCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s PurgeCachesResponse) GoString() string {
	return s.String()
}

func (s *PurgeCachesResponse) SetHeaders(v map[string]*string) *PurgeCachesResponse {
	s.Headers = v
	return s
}

func (s *PurgeCachesResponse) SetStatusCode(v int32) *PurgeCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *PurgeCachesResponse) SetBody(v *PurgeCachesResponseBody) *PurgeCachesResponse {
	s.Body = v
	return s
}

type PutKvRequest struct {
	// example:
	//
	// true
	Base64 *bool `json:"Base64,omitempty" xml:"Base64,omitempty"`
	// example:
	//
	// 1690081381
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// example:
	//
	// 3600
	ExpirationTtl *int64 `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutKvRequest) String() string {
	return tea.Prettify(s)
}

func (s PutKvRequest) GoString() string {
	return s.String()
}

func (s *PutKvRequest) SetBase64(v bool) *PutKvRequest {
	s.Base64 = &v
	return s
}

func (s *PutKvRequest) SetExpiration(v int64) *PutKvRequest {
	s.Expiration = &v
	return s
}

func (s *PutKvRequest) SetExpirationTtl(v int64) *PutKvRequest {
	s.ExpirationTtl = &v
	return s
}

func (s *PutKvRequest) SetKey(v string) *PutKvRequest {
	s.Key = &v
	return s
}

func (s *PutKvRequest) SetNamespace(v string) *PutKvRequest {
	s.Namespace = &v
	return s
}

func (s *PutKvRequest) SetValue(v string) *PutKvRequest {
	s.Value = &v
	return s
}

type PutKvResponseBody struct {
	// example:
	//
	// 4
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutKvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutKvResponseBody) GoString() string {
	return s.String()
}

func (s *PutKvResponseBody) SetLength(v string) *PutKvResponseBody {
	s.Length = &v
	return s
}

func (s *PutKvResponseBody) SetRequestId(v string) *PutKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutKvResponseBody) SetValue(v string) *PutKvResponseBody {
	s.Value = &v
	return s
}

type PutKvResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutKvResponse) String() string {
	return tea.Prettify(s)
}

func (s PutKvResponse) GoString() string {
	return s.String()
}

func (s *PutKvResponse) SetHeaders(v map[string]*string) *PutKvResponse {
	s.Headers = v
	return s
}

func (s *PutKvResponse) SetStatusCode(v int32) *PutKvResponse {
	s.StatusCode = &v
	return s
}

func (s *PutKvResponse) SetBody(v *PutKvResponseBody) *PutKvResponse {
	s.Body = v
	return s
}

type PutKvAccountRequest struct {
	// example:
	//
	// prod
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PutKvAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s PutKvAccountRequest) GoString() string {
	return s.String()
}

func (s *PutKvAccountRequest) SetAccountType(v string) *PutKvAccountRequest {
	s.AccountType = &v
	return s
}

func (s *PutKvAccountRequest) SetStatus(v string) *PutKvAccountRequest {
	s.Status = &v
	return s
}

type PutKvAccountResponseBody struct {
	NamespaceList []*PutKvAccountResponseBodyNamespaceList `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	NamespaceQuota *int32 `json:"NamespaceQuota,omitempty" xml:"NamespaceQuota,omitempty"`
	// example:
	//
	// 1
	NamespaceUsed *int32 `json:"NamespaceUsed,omitempty" xml:"NamespaceUsed,omitempty"`
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PutKvAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutKvAccountResponseBody) GoString() string {
	return s.String()
}

func (s *PutKvAccountResponseBody) SetNamespaceList(v []*PutKvAccountResponseBodyNamespaceList) *PutKvAccountResponseBody {
	s.NamespaceList = v
	return s
}

func (s *PutKvAccountResponseBody) SetNamespaceQuota(v int32) *PutKvAccountResponseBody {
	s.NamespaceQuota = &v
	return s
}

func (s *PutKvAccountResponseBody) SetNamespaceUsed(v int32) *PutKvAccountResponseBody {
	s.NamespaceUsed = &v
	return s
}

func (s *PutKvAccountResponseBody) SetRequestId(v string) *PutKvAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutKvAccountResponseBody) SetStatus(v string) *PutKvAccountResponseBody {
	s.Status = &v
	return s
}

type PutKvAccountResponseBodyNamespaceList struct {
	// example:
	//
	// the first namespace
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 643355322374688768
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PutKvAccountResponseBodyNamespaceList) String() string {
	return tea.Prettify(s)
}

func (s PutKvAccountResponseBodyNamespaceList) GoString() string {
	return s.String()
}

func (s *PutKvAccountResponseBodyNamespaceList) SetDescription(v string) *PutKvAccountResponseBodyNamespaceList {
	s.Description = &v
	return s
}

func (s *PutKvAccountResponseBodyNamespaceList) SetNamespace(v string) *PutKvAccountResponseBodyNamespaceList {
	s.Namespace = &v
	return s
}

func (s *PutKvAccountResponseBodyNamespaceList) SetNamespaceId(v string) *PutKvAccountResponseBodyNamespaceList {
	s.NamespaceId = &v
	return s
}

func (s *PutKvAccountResponseBodyNamespaceList) SetStatus(v string) *PutKvAccountResponseBodyNamespaceList {
	s.Status = &v
	return s
}

type PutKvAccountResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutKvAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutKvAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s PutKvAccountResponse) GoString() string {
	return s.String()
}

func (s *PutKvAccountResponse) SetHeaders(v map[string]*string) *PutKvAccountResponse {
	s.Headers = v
	return s
}

func (s *PutKvAccountResponse) SetStatusCode(v int32) *PutKvAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *PutKvAccountResponse) SetBody(v *PutKvAccountResponseBody) *PutKvAccountResponse {
	s.Body = v
	return s
}

type PutKvWithHighCapacityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_namesapce
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PutKvWithHighCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s PutKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *PutKvWithHighCapacityRequest) SetKey(v string) *PutKvWithHighCapacityRequest {
	s.Key = &v
	return s
}

func (s *PutKvWithHighCapacityRequest) SetNamespace(v string) *PutKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *PutKvWithHighCapacityRequest) SetUrl(v string) *PutKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

type PutKvWithHighCapacityAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_namesapce
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PutKvWithHighCapacityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PutKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *PutKvWithHighCapacityAdvanceRequest) SetKey(v string) *PutKvWithHighCapacityAdvanceRequest {
	s.Key = &v
	return s
}

func (s *PutKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *PutKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *PutKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *PutKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

type PutKvWithHighCapacityResponseBody struct {
	// example:
	//
	// 4
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutKvWithHighCapacityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutKvWithHighCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *PutKvWithHighCapacityResponseBody) SetLength(v string) *PutKvWithHighCapacityResponseBody {
	s.Length = &v
	return s
}

func (s *PutKvWithHighCapacityResponseBody) SetRequestId(v string) *PutKvWithHighCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutKvWithHighCapacityResponseBody) SetValue(v string) *PutKvWithHighCapacityResponseBody {
	s.Value = &v
	return s
}

type PutKvWithHighCapacityResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutKvWithHighCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutKvWithHighCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s PutKvWithHighCapacityResponse) GoString() string {
	return s.String()
}

func (s *PutKvWithHighCapacityResponse) SetHeaders(v map[string]*string) *PutKvWithHighCapacityResponse {
	s.Headers = v
	return s
}

func (s *PutKvWithHighCapacityResponse) SetStatusCode(v int32) *PutKvWithHighCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *PutKvWithHighCapacityResponse) SetBody(v *PutKvWithHighCapacityResponseBody) *PutKvWithHighCapacityResponse {
	s.Body = v
	return s
}

type ResetScheduledPreloadJobRequest struct {
	// example:
	//
	// ResetScheduledPreloadJob
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ResetScheduledPreloadJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetScheduledPreloadJobRequest) GoString() string {
	return s.String()
}

func (s *ResetScheduledPreloadJobRequest) SetId(v string) *ResetScheduledPreloadJobRequest {
	s.Id = &v
	return s
}

type ResetScheduledPreloadJobResponseBody struct {
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CreatedAt     *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Domains       *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	ErrorInfo     *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	FailedFileOss *string `json:"FailedFileOss,omitempty" xml:"FailedFileOss,omitempty"`
	FileId        *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InsertWay     *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId        *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TaskSubmitted *int32  `json:"TaskSubmitted,omitempty" xml:"TaskSubmitted,omitempty"`
	TaskType      *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UrlCount      *int32  `json:"UrlCount,omitempty" xml:"UrlCount,omitempty"`
	UrlSubmitted  *int32  `json:"UrlSubmitted,omitempty" xml:"UrlSubmitted,omitempty"`
}

func (s ResetScheduledPreloadJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetScheduledPreloadJobResponseBody) GoString() string {
	return s.String()
}

func (s *ResetScheduledPreloadJobResponseBody) SetAliUid(v string) *ResetScheduledPreloadJobResponseBody {
	s.AliUid = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetCreatedAt(v string) *ResetScheduledPreloadJobResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetDomains(v string) *ResetScheduledPreloadJobResponseBody {
	s.Domains = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetErrorInfo(v string) *ResetScheduledPreloadJobResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetFailedFileOss(v string) *ResetScheduledPreloadJobResponseBody {
	s.FailedFileOss = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetFileId(v string) *ResetScheduledPreloadJobResponseBody {
	s.FileId = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetId(v string) *ResetScheduledPreloadJobResponseBody {
	s.Id = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetInsertWay(v string) *ResetScheduledPreloadJobResponseBody {
	s.InsertWay = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetName(v string) *ResetScheduledPreloadJobResponseBody {
	s.Name = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetRequestId(v string) *ResetScheduledPreloadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetSiteId(v int64) *ResetScheduledPreloadJobResponseBody {
	s.SiteId = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetTaskSubmitted(v int32) *ResetScheduledPreloadJobResponseBody {
	s.TaskSubmitted = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetTaskType(v string) *ResetScheduledPreloadJobResponseBody {
	s.TaskType = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetUrlCount(v int32) *ResetScheduledPreloadJobResponseBody {
	s.UrlCount = &v
	return s
}

func (s *ResetScheduledPreloadJobResponseBody) SetUrlSubmitted(v int32) *ResetScheduledPreloadJobResponseBody {
	s.UrlSubmitted = &v
	return s
}

type ResetScheduledPreloadJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetScheduledPreloadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetScheduledPreloadJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetScheduledPreloadJobResponse) GoString() string {
	return s.String()
}

func (s *ResetScheduledPreloadJobResponse) SetHeaders(v map[string]*string) *ResetScheduledPreloadJobResponse {
	s.Headers = v
	return s
}

func (s *ResetScheduledPreloadJobResponse) SetStatusCode(v int32) *ResetScheduledPreloadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetScheduledPreloadJobResponse) SetBody(v *ResetScheduledPreloadJobResponseBody) *ResetScheduledPreloadJobResponse {
	s.Body = v
	return s
}

type SetCertificateRequest struct {
	// example:
	//
	// 30000478
	CasId *int64 `json:"CasId,omitempty" xml:"CasId,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// 30001303
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// yourCertName
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// -----BEGIN PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cas
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// true
	Update *bool `json:"Update,omitempty" xml:"Update,omitempty"`
}

func (s SetCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetCertificateRequest) SetCasId(v int64) *SetCertificateRequest {
	s.CasId = &v
	return s
}

func (s *SetCertificateRequest) SetCertificate(v string) *SetCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *SetCertificateRequest) SetId(v string) *SetCertificateRequest {
	s.Id = &v
	return s
}

func (s *SetCertificateRequest) SetName(v string) *SetCertificateRequest {
	s.Name = &v
	return s
}

func (s *SetCertificateRequest) SetOwnerId(v int64) *SetCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCertificateRequest) SetPrivateKey(v string) *SetCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *SetCertificateRequest) SetRegion(v string) *SetCertificateRequest {
	s.Region = &v
	return s
}

func (s *SetCertificateRequest) SetSecurityToken(v string) *SetCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetCertificateRequest) SetSiteId(v int64) *SetCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *SetCertificateRequest) SetType(v string) *SetCertificateRequest {
	s.Type = &v
	return s
}

func (s *SetCertificateRequest) SetUpdate(v bool) *SetCertificateRequest {
	s.Update = &v
	return s
}

type SetCertificateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A666D44F-19D6-490E-97CF-1A64AB962C57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetCertificateResponseBody) SetRequestId(v string) *SetCertificateResponseBody {
	s.RequestId = &v
	return s
}

type SetCertificateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetCertificateResponse) SetHeaders(v map[string]*string) *SetCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetCertificateResponse) SetStatusCode(v int32) *SetCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCertificateResponse) SetBody(v *SetCertificateResponseBody) *SetCertificateResponse {
	s.Body = v
	return s
}

type SetHttpDDoSAttackIntelligentProtectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// defense
	AiMode *string `json:"AiMode,omitempty" xml:"AiMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// level60
	AiTemplate *string `json:"AiTemplate,omitempty" xml:"AiTemplate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetHttpDDoSAttackIntelligentProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetHttpDDoSAttackIntelligentProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackIntelligentProtectionRequest) SetAiMode(v string) *SetHttpDDoSAttackIntelligentProtectionRequest {
	s.AiMode = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionRequest) SetAiTemplate(v string) *SetHttpDDoSAttackIntelligentProtectionRequest {
	s.AiTemplate = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionRequest) SetSiteId(v int64) *SetHttpDDoSAttackIntelligentProtectionRequest {
	s.SiteId = &v
	return s
}

type SetHttpDDoSAttackIntelligentProtectionResponseBody struct {
	// example:
	//
	// defense
	AiMode *string `json:"AiMode,omitempty" xml:"AiMode,omitempty"`
	// example:
	//
	// level60
	AiTemplate *string `json:"AiTemplate,omitempty" xml:"AiTemplate,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 156A6B-677B1A-4297B7-9187B7-2B44792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetHttpDDoSAttackIntelligentProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetHttpDDoSAttackIntelligentProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) SetAiMode(v string) *SetHttpDDoSAttackIntelligentProtectionResponseBody {
	s.AiMode = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) SetAiTemplate(v string) *SetHttpDDoSAttackIntelligentProtectionResponseBody {
	s.AiTemplate = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) SetRequestId(v string) *SetHttpDDoSAttackIntelligentProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) SetSiteId(v int64) *SetHttpDDoSAttackIntelligentProtectionResponseBody {
	s.SiteId = &v
	return s
}

type SetHttpDDoSAttackIntelligentProtectionResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetHttpDDoSAttackIntelligentProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetHttpDDoSAttackIntelligentProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetHttpDDoSAttackIntelligentProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponse) SetHeaders(v map[string]*string) *SetHttpDDoSAttackIntelligentProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponse) SetStatusCode(v int32) *SetHttpDDoSAttackIntelligentProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponse) SetBody(v *SetHttpDDoSAttackIntelligentProtectionResponseBody) *SetHttpDDoSAttackIntelligentProtectionResponse {
	s.Body = v
	return s
}

type SetHttpDDoSAttackProtectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// default
	GlobalMode *string `json:"GlobalMode,omitempty" xml:"GlobalMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetHttpDDoSAttackProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetHttpDDoSAttackProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackProtectionRequest) SetGlobalMode(v string) *SetHttpDDoSAttackProtectionRequest {
	s.GlobalMode = &v
	return s
}

func (s *SetHttpDDoSAttackProtectionRequest) SetSiteId(v int64) *SetHttpDDoSAttackProtectionRequest {
	s.SiteId = &v
	return s
}

type SetHttpDDoSAttackProtectionResponseBody struct {
	// example:
	//
	// default
	GlobalMode *string `json:"GlobalMode,omitempty" xml:"GlobalMode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetHttpDDoSAttackProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetHttpDDoSAttackProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackProtectionResponseBody) SetGlobalMode(v string) *SetHttpDDoSAttackProtectionResponseBody {
	s.GlobalMode = &v
	return s
}

func (s *SetHttpDDoSAttackProtectionResponseBody) SetRequestId(v string) *SetHttpDDoSAttackProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetHttpDDoSAttackProtectionResponseBody) SetSiteId(v int64) *SetHttpDDoSAttackProtectionResponseBody {
	s.SiteId = &v
	return s
}

type SetHttpDDoSAttackProtectionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetHttpDDoSAttackProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetHttpDDoSAttackProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetHttpDDoSAttackProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackProtectionResponse) SetHeaders(v map[string]*string) *SetHttpDDoSAttackProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetHttpDDoSAttackProtectionResponse) SetStatusCode(v int32) *SetHttpDDoSAttackProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetHttpDDoSAttackProtectionResponse) SetBody(v *SetHttpDDoSAttackProtectionResponseBody) *SetHttpDDoSAttackProtectionResponse {
	s.Body = v
	return s
}

type StartScheduledPreloadExecutionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// StartScheduledPreloadExecution
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StartScheduledPreloadExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s StartScheduledPreloadExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartScheduledPreloadExecutionRequest) SetId(v string) *StartScheduledPreloadExecutionRequest {
	s.Id = &v
	return s
}

type StartScheduledPreloadExecutionResponseBody struct {
	AliUid  *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 60
	Interval *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StartScheduledPreloadExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartScheduledPreloadExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartScheduledPreloadExecutionResponseBody) SetAliUid(v string) *StartScheduledPreloadExecutionResponseBody {
	s.AliUid = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetEndTime(v string) *StartScheduledPreloadExecutionResponseBody {
	s.EndTime = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetId(v string) *StartScheduledPreloadExecutionResponseBody {
	s.Id = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetInterval(v int32) *StartScheduledPreloadExecutionResponseBody {
	s.Interval = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetJobId(v string) *StartScheduledPreloadExecutionResponseBody {
	s.JobId = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetRequestId(v string) *StartScheduledPreloadExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetSliceLen(v int32) *StartScheduledPreloadExecutionResponseBody {
	s.SliceLen = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetStartTime(v string) *StartScheduledPreloadExecutionResponseBody {
	s.StartTime = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetStatus(v string) *StartScheduledPreloadExecutionResponseBody {
	s.Status = &v
	return s
}

type StartScheduledPreloadExecutionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartScheduledPreloadExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartScheduledPreloadExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s StartScheduledPreloadExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartScheduledPreloadExecutionResponse) SetHeaders(v map[string]*string) *StartScheduledPreloadExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartScheduledPreloadExecutionResponse) SetStatusCode(v int32) *StartScheduledPreloadExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponse) SetBody(v *StartScheduledPreloadExecutionResponseBody) *StartScheduledPreloadExecutionResponse {
	s.Body = v
	return s
}

type StopScheduledPreloadExecutionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// StopScheduledPreloadExecution
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StopScheduledPreloadExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s StopScheduledPreloadExecutionRequest) GoString() string {
	return s.String()
}

func (s *StopScheduledPreloadExecutionRequest) SetId(v string) *StopScheduledPreloadExecutionRequest {
	s.Id = &v
	return s
}

type StopScheduledPreloadExecutionResponseBody struct {
	AliUid   *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StopScheduledPreloadExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopScheduledPreloadExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StopScheduledPreloadExecutionResponseBody) SetAliUid(v string) *StopScheduledPreloadExecutionResponseBody {
	s.AliUid = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetEndTime(v string) *StopScheduledPreloadExecutionResponseBody {
	s.EndTime = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetId(v string) *StopScheduledPreloadExecutionResponseBody {
	s.Id = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetInterval(v int32) *StopScheduledPreloadExecutionResponseBody {
	s.Interval = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetJobId(v string) *StopScheduledPreloadExecutionResponseBody {
	s.JobId = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetRequestId(v string) *StopScheduledPreloadExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetSliceLen(v int32) *StopScheduledPreloadExecutionResponseBody {
	s.SliceLen = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetStartTime(v string) *StopScheduledPreloadExecutionResponseBody {
	s.StartTime = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetStatus(v string) *StopScheduledPreloadExecutionResponseBody {
	s.Status = &v
	return s
}

type StopScheduledPreloadExecutionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopScheduledPreloadExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopScheduledPreloadExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s StopScheduledPreloadExecutionResponse) GoString() string {
	return s.String()
}

func (s *StopScheduledPreloadExecutionResponse) SetHeaders(v map[string]*string) *StopScheduledPreloadExecutionResponse {
	s.Headers = v
	return s
}

func (s *StopScheduledPreloadExecutionResponse) SetStatusCode(v int32) *StopScheduledPreloadExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponse) SetBody(v *StopScheduledPreloadExecutionResponseBody) *StopScheduledPreloadExecutionResponse {
	s.Body = v
	return s
}

type TransformExpressionToMatchRequest struct {
	// example:
	//
	// http_bot
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s TransformExpressionToMatchRequest) String() string {
	return tea.Prettify(s)
}

func (s TransformExpressionToMatchRequest) GoString() string {
	return s.String()
}

func (s *TransformExpressionToMatchRequest) SetExpression(v string) *TransformExpressionToMatchRequest {
	s.Expression = &v
	return s
}

func (s *TransformExpressionToMatchRequest) SetPhase(v string) *TransformExpressionToMatchRequest {
	s.Phase = &v
	return s
}

func (s *TransformExpressionToMatchRequest) SetSiteId(v int64) *TransformExpressionToMatchRequest {
	s.SiteId = &v
	return s
}

type TransformExpressionToMatchResponseBody struct {
	Match *WafRuleMatch `json:"Match,omitempty" xml:"Match,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransformExpressionToMatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransformExpressionToMatchResponseBody) GoString() string {
	return s.String()
}

func (s *TransformExpressionToMatchResponseBody) SetMatch(v *WafRuleMatch) *TransformExpressionToMatchResponseBody {
	s.Match = v
	return s
}

func (s *TransformExpressionToMatchResponseBody) SetRequestId(v string) *TransformExpressionToMatchResponseBody {
	s.RequestId = &v
	return s
}

type TransformExpressionToMatchResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransformExpressionToMatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransformExpressionToMatchResponse) String() string {
	return tea.Prettify(s)
}

func (s TransformExpressionToMatchResponse) GoString() string {
	return s.String()
}

func (s *TransformExpressionToMatchResponse) SetHeaders(v map[string]*string) *TransformExpressionToMatchResponse {
	s.Headers = v
	return s
}

func (s *TransformExpressionToMatchResponse) SetStatusCode(v int32) *TransformExpressionToMatchResponse {
	s.StatusCode = &v
	return s
}

func (s *TransformExpressionToMatchResponse) SetBody(v *TransformExpressionToMatchResponseBody) *TransformExpressionToMatchResponse {
	s.Body = v
	return s
}

type TransformMatchToExpressionRequest struct {
	// example:
	//
	// http_bot
	Match *WafRuleMatch `json:"Match,omitempty" xml:"Match,omitempty"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s TransformMatchToExpressionRequest) String() string {
	return tea.Prettify(s)
}

func (s TransformMatchToExpressionRequest) GoString() string {
	return s.String()
}

func (s *TransformMatchToExpressionRequest) SetMatch(v *WafRuleMatch) *TransformMatchToExpressionRequest {
	s.Match = v
	return s
}

func (s *TransformMatchToExpressionRequest) SetPhase(v string) *TransformMatchToExpressionRequest {
	s.Phase = &v
	return s
}

func (s *TransformMatchToExpressionRequest) SetSiteId(v int64) *TransformMatchToExpressionRequest {
	s.SiteId = &v
	return s
}

type TransformMatchToExpressionShrinkRequest struct {
	// example:
	//
	// http_bot
	MatchShrink *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s TransformMatchToExpressionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TransformMatchToExpressionShrinkRequest) GoString() string {
	return s.String()
}

func (s *TransformMatchToExpressionShrinkRequest) SetMatchShrink(v string) *TransformMatchToExpressionShrinkRequest {
	s.MatchShrink = &v
	return s
}

func (s *TransformMatchToExpressionShrinkRequest) SetPhase(v string) *TransformMatchToExpressionShrinkRequest {
	s.Phase = &v
	return s
}

func (s *TransformMatchToExpressionShrinkRequest) SetSiteId(v int64) *TransformMatchToExpressionShrinkRequest {
	s.SiteId = &v
	return s
}

type TransformMatchToExpressionResponseBody struct {
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransformMatchToExpressionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransformMatchToExpressionResponseBody) GoString() string {
	return s.String()
}

func (s *TransformMatchToExpressionResponseBody) SetExpression(v string) *TransformMatchToExpressionResponseBody {
	s.Expression = &v
	return s
}

func (s *TransformMatchToExpressionResponseBody) SetRequestId(v string) *TransformMatchToExpressionResponseBody {
	s.RequestId = &v
	return s
}

type TransformMatchToExpressionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransformMatchToExpressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransformMatchToExpressionResponse) String() string {
	return tea.Prettify(s)
}

func (s TransformMatchToExpressionResponse) GoString() string {
	return s.String()
}

func (s *TransformMatchToExpressionResponse) SetHeaders(v map[string]*string) *TransformMatchToExpressionResponse {
	s.Headers = v
	return s
}

func (s *TransformMatchToExpressionResponse) SetStatusCode(v int32) *TransformMatchToExpressionResponse {
	s.StatusCode = &v
	return s
}

func (s *TransformMatchToExpressionResponse) SetBody(v *TransformMatchToExpressionResponseBody) *TransformMatchToExpressionResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// 是否全部删除，只针对TagKey.N为空时有效
	//
	// example:
	//
	// false
	All     *bool  `json:"All,omitempty" xml:"All,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 要创建并绑定标签的资源所在的地域ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源ID,最多 50个子项
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// site
	ResourceType  *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TagKey        []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetSecurityToken(v string) *UntagResourcesRequest {
	s.SecurityToken = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// example:
	//
	// 85H66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateCustomScenePolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-03T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123456****
	Objects *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s UpdateCustomScenePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomScenePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomScenePolicyRequest) SetEndTime(v string) *UpdateCustomScenePolicyRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetName(v string) *UpdateCustomScenePolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetObjects(v string) *UpdateCustomScenePolicyRequest {
	s.Objects = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetPolicyId(v int64) *UpdateCustomScenePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetStartTime(v string) *UpdateCustomScenePolicyRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetTemplate(v string) *UpdateCustomScenePolicyRequest {
	s.Template = &v
	return s
}

type UpdateCustomScenePolicyResponseBody struct {
	// example:
	//
	// 2023-04-03T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// test
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Objects []*string `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2023-04-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s UpdateCustomScenePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomScenePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomScenePolicyResponseBody) SetEndTime(v string) *UpdateCustomScenePolicyResponseBody {
	s.EndTime = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetName(v string) *UpdateCustomScenePolicyResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetObjects(v []*string) *UpdateCustomScenePolicyResponseBody {
	s.Objects = v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetPolicyId(v int64) *UpdateCustomScenePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetRequestId(v string) *UpdateCustomScenePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetStartTime(v string) *UpdateCustomScenePolicyResponseBody {
	s.StartTime = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetTemplate(v string) *UpdateCustomScenePolicyResponseBody {
	s.Template = &v
	return s
}

type UpdateCustomScenePolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomScenePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomScenePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomScenePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomScenePolicyResponse) SetHeaders(v map[string]*string) *UpdateCustomScenePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomScenePolicyResponse) SetStatusCode(v int32) *UpdateCustomScenePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomScenePolicyResponse) SetBody(v *UpdateCustomScenePolicyResponseBody) *UpdateCustomScenePolicyResponse {
	s.Body = v
	return s
}

type UpdateKvNamespaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// new_ns
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateKvNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKvNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateKvNamespaceRequest) SetNamespace(v string) *UpdateKvNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateKvNamespaceRequest) SetTitle(v string) *UpdateKvNamespaceRequest {
	s.Title = &v
	return s
}

type UpdateKvNamespaceResponseBody struct {
	// example:
	//
	// this is a test ns.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// new_ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 643355322374688768
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateKvNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKvNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKvNamespaceResponseBody) SetDescription(v string) *UpdateKvNamespaceResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateKvNamespaceResponseBody) SetNamespace(v string) *UpdateKvNamespaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *UpdateKvNamespaceResponseBody) SetNamespaceId(v string) *UpdateKvNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *UpdateKvNamespaceResponseBody) SetRequestId(v string) *UpdateKvNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKvNamespaceResponseBody) SetStatus(v string) *UpdateKvNamespaceResponseBody {
	s.Status = &v
	return s
}

type UpdateKvNamespaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKvNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKvNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKvNamespaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateKvNamespaceResponse) SetHeaders(v map[string]*string) *UpdateKvNamespaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateKvNamespaceResponse) SetStatusCode(v int32) *UpdateKvNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKvNamespaceResponse) SetBody(v *UpdateKvNamespaceResponseBody) *UpdateKvNamespaceResponse {
	s.Body = v
	return s
}

type UpdateListRequest struct {
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// a custom list
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateListRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListRequest) GoString() string {
	return s.String()
}

func (s *UpdateListRequest) SetDescription(v string) *UpdateListRequest {
	s.Description = &v
	return s
}

func (s *UpdateListRequest) SetId(v int64) *UpdateListRequest {
	s.Id = &v
	return s
}

func (s *UpdateListRequest) SetItems(v []*string) *UpdateListRequest {
	s.Items = v
	return s
}

func (s *UpdateListRequest) SetName(v string) *UpdateListRequest {
	s.Name = &v
	return s
}

type UpdateListShrinkRequest struct {
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// a custom list
	ItemsShrink *string `json:"Items,omitempty" xml:"Items,omitempty"`
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateListShrinkRequest) SetDescription(v string) *UpdateListShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateListShrinkRequest) SetId(v int64) *UpdateListShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateListShrinkRequest) SetItemsShrink(v string) *UpdateListShrinkRequest {
	s.ItemsShrink = &v
	return s
}

func (s *UpdateListShrinkRequest) SetName(v string) *UpdateListShrinkRequest {
	s.Name = &v
	return s
}

type UpdateListResponseBody struct {
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListResponseBody) SetRequestId(v string) *UpdateListResponseBody {
	s.RequestId = &v
	return s
}

type UpdateListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateListResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateListResponse) GoString() string {
	return s.String()
}

func (s *UpdateListResponse) SetHeaders(v map[string]*string) *UpdateListResponse {
	s.Headers = v
	return s
}

func (s *UpdateListResponse) SetStatusCode(v int32) *UpdateListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateListResponse) SetBody(v *UpdateListResponseBody) *UpdateListResponse {
	s.Body = v
	return s
}

type UpdatePageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text/html
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdatePageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePageRequest) GoString() string {
	return s.String()
}

func (s *UpdatePageRequest) SetContent(v string) *UpdatePageRequest {
	s.Content = &v
	return s
}

func (s *UpdatePageRequest) SetContentType(v string) *UpdatePageRequest {
	s.ContentType = &v
	return s
}

func (s *UpdatePageRequest) SetDescription(v string) *UpdatePageRequest {
	s.Description = &v
	return s
}

func (s *UpdatePageRequest) SetId(v int64) *UpdatePageRequest {
	s.Id = &v
	return s
}

func (s *UpdatePageRequest) SetName(v string) *UpdatePageRequest {
	s.Name = &v
	return s
}

type UpdatePageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePageResponseBody) SetRequestId(v string) *UpdatePageResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePageResponse) GoString() string {
	return s.String()
}

func (s *UpdatePageResponse) SetHeaders(v map[string]*string) *UpdatePageResponse {
	s.Headers = v
	return s
}

func (s *UpdatePageResponse) SetStatusCode(v int32) *UpdatePageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePageResponse) SetBody(v *UpdatePageResponseBody) *UpdatePageResponse {
	s.Body = v
	return s
}

type UpdateRecordRequest struct {
	AuthConf *UpdateRecordRequestAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	BizName  *string                      `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Comment  *string                      `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	Data       *UpdateRecordRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HostPolicy *string                  `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// 是否代理加速
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// This parameter is required.
	RecordId   *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Ttl        *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordRequest) SetAuthConf(v *UpdateRecordRequestAuthConf) *UpdateRecordRequest {
	s.AuthConf = v
	return s
}

func (s *UpdateRecordRequest) SetBizName(v string) *UpdateRecordRequest {
	s.BizName = &v
	return s
}

func (s *UpdateRecordRequest) SetComment(v string) *UpdateRecordRequest {
	s.Comment = &v
	return s
}

func (s *UpdateRecordRequest) SetData(v *UpdateRecordRequestData) *UpdateRecordRequest {
	s.Data = v
	return s
}

func (s *UpdateRecordRequest) SetHostPolicy(v string) *UpdateRecordRequest {
	s.HostPolicy = &v
	return s
}

func (s *UpdateRecordRequest) SetProxied(v bool) *UpdateRecordRequest {
	s.Proxied = &v
	return s
}

func (s *UpdateRecordRequest) SetRecordId(v int64) *UpdateRecordRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecordRequest) SetSourceType(v string) *UpdateRecordRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateRecordRequest) SetTtl(v int32) *UpdateRecordRequest {
	s.Ttl = &v
	return s
}

type UpdateRecordRequestAuthConf struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateRecordRequestAuthConf) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordRequestAuthConf) GoString() string {
	return s.String()
}

func (s *UpdateRecordRequestAuthConf) SetAccessKey(v string) *UpdateRecordRequestAuthConf {
	s.AccessKey = &v
	return s
}

func (s *UpdateRecordRequestAuthConf) SetAuthType(v string) *UpdateRecordRequestAuthConf {
	s.AuthType = &v
	return s
}

func (s *UpdateRecordRequestAuthConf) SetRegion(v string) *UpdateRecordRequestAuthConf {
	s.Region = &v
	return s
}

func (s *UpdateRecordRequestAuthConf) SetSecretKey(v string) *UpdateRecordRequestAuthConf {
	s.SecretKey = &v
	return s
}

func (s *UpdateRecordRequestAuthConf) SetVersion(v string) *UpdateRecordRequestAuthConf {
	s.Version = &v
	return s
}

type UpdateRecordRequestData struct {
	Algorithm    *int32  `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Certificate  *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Fingerprint  *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Flag         *int32  `json:"Flag,omitempty" xml:"Flag,omitempty"`
	KeyTag       *int32  `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	MatchingType *int32  `json:"MatchingType,omitempty" xml:"MatchingType,omitempty"`
	Port         *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Selector     *int32  `json:"Selector,omitempty" xml:"Selector,omitempty"`
	Tag          *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Usage        *int32  `json:"Usage,omitempty" xml:"Usage,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Weight       *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateRecordRequestData) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordRequestData) GoString() string {
	return s.String()
}

func (s *UpdateRecordRequestData) SetAlgorithm(v int32) *UpdateRecordRequestData {
	s.Algorithm = &v
	return s
}

func (s *UpdateRecordRequestData) SetCertificate(v string) *UpdateRecordRequestData {
	s.Certificate = &v
	return s
}

func (s *UpdateRecordRequestData) SetFingerprint(v string) *UpdateRecordRequestData {
	s.Fingerprint = &v
	return s
}

func (s *UpdateRecordRequestData) SetFlag(v int32) *UpdateRecordRequestData {
	s.Flag = &v
	return s
}

func (s *UpdateRecordRequestData) SetKeyTag(v int32) *UpdateRecordRequestData {
	s.KeyTag = &v
	return s
}

func (s *UpdateRecordRequestData) SetMatchingType(v int32) *UpdateRecordRequestData {
	s.MatchingType = &v
	return s
}

func (s *UpdateRecordRequestData) SetPort(v int32) *UpdateRecordRequestData {
	s.Port = &v
	return s
}

func (s *UpdateRecordRequestData) SetPriority(v int32) *UpdateRecordRequestData {
	s.Priority = &v
	return s
}

func (s *UpdateRecordRequestData) SetSelector(v int32) *UpdateRecordRequestData {
	s.Selector = &v
	return s
}

func (s *UpdateRecordRequestData) SetTag(v string) *UpdateRecordRequestData {
	s.Tag = &v
	return s
}

func (s *UpdateRecordRequestData) SetType(v int32) *UpdateRecordRequestData {
	s.Type = &v
	return s
}

func (s *UpdateRecordRequestData) SetUsage(v int32) *UpdateRecordRequestData {
	s.Usage = &v
	return s
}

func (s *UpdateRecordRequestData) SetValue(v string) *UpdateRecordRequestData {
	s.Value = &v
	return s
}

func (s *UpdateRecordRequestData) SetWeight(v int32) *UpdateRecordRequestData {
	s.Weight = &v
	return s
}

type UpdateRecordShrinkRequest struct {
	AuthConfShrink *string `json:"AuthConf,omitempty" xml:"AuthConf,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Comment        *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	DataShrink *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// 是否代理加速
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// This parameter is required.
	RecordId   *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Ttl        *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateRecordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordShrinkRequest) SetAuthConfShrink(v string) *UpdateRecordShrinkRequest {
	s.AuthConfShrink = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetBizName(v string) *UpdateRecordShrinkRequest {
	s.BizName = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetComment(v string) *UpdateRecordShrinkRequest {
	s.Comment = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetDataShrink(v string) *UpdateRecordShrinkRequest {
	s.DataShrink = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetHostPolicy(v string) *UpdateRecordShrinkRequest {
	s.HostPolicy = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetProxied(v bool) *UpdateRecordShrinkRequest {
	s.Proxied = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetRecordId(v int64) *UpdateRecordShrinkRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetSourceType(v string) *UpdateRecordShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetTtl(v int32) *UpdateRecordShrinkRequest {
	s.Ttl = &v
	return s
}

type UpdateRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordResponseBody) SetRequestId(v string) *UpdateRecordResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRecordResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordResponse) SetHeaders(v map[string]*string) *UpdateRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordResponse) SetStatusCode(v int32) *UpdateRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecordResponse) SetBody(v *UpdateRecordResponseBody) *UpdateRecordResponse {
	s.Body = v
	return s
}

type UpdateScheduledPreloadExecutionRequest struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UpdateScheduledPreloadExecution
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s UpdateScheduledPreloadExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduledPreloadExecutionRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPreloadExecutionRequest) SetEndTime(v string) *UpdateScheduledPreloadExecutionRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionRequest) SetId(v string) *UpdateScheduledPreloadExecutionRequest {
	s.Id = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionRequest) SetInterval(v int32) *UpdateScheduledPreloadExecutionRequest {
	s.Interval = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionRequest) SetSliceLen(v int32) *UpdateScheduledPreloadExecutionRequest {
	s.SliceLen = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionRequest) SetStartTime(v string) *UpdateScheduledPreloadExecutionRequest {
	s.StartTime = &v
	return s
}

type UpdateScheduledPreloadExecutionResponseBody struct {
	AliUid   *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateScheduledPreloadExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduledPreloadExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetAliUid(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.AliUid = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetEndTime(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.EndTime = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetId(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetInterval(v int32) *UpdateScheduledPreloadExecutionResponseBody {
	s.Interval = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetJobId(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetRequestId(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetSliceLen(v int32) *UpdateScheduledPreloadExecutionResponseBody {
	s.SliceLen = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetStartTime(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.StartTime = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetStatus(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.Status = &v
	return s
}

type UpdateScheduledPreloadExecutionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduledPreloadExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduledPreloadExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduledPreloadExecutionResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPreloadExecutionResponse) SetHeaders(v map[string]*string) *UpdateScheduledPreloadExecutionResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponse) SetStatusCode(v int32) *UpdateScheduledPreloadExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponse) SetBody(v *UpdateScheduledPreloadExecutionResponseBody) *UpdateScheduledPreloadExecutionResponse {
	s.Body = v
	return s
}

type UpdateSiteAccessTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// NS
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSiteAccessTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteAccessTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteAccessTypeRequest) SetAccessType(v string) *UpdateSiteAccessTypeRequest {
	s.AccessType = &v
	return s
}

func (s *UpdateSiteAccessTypeRequest) SetSiteId(v int64) *UpdateSiteAccessTypeRequest {
	s.SiteId = &v
	return s
}

type UpdateSiteAccessTypeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteAccessTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteAccessTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteAccessTypeResponseBody) SetRequestId(v string) *UpdateSiteAccessTypeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSiteAccessTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteAccessTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteAccessTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteAccessTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteAccessTypeResponse) SetHeaders(v map[string]*string) *UpdateSiteAccessTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteAccessTypeResponse) SetStatusCode(v int32) *UpdateSiteAccessTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteAccessTypeResponse) SetBody(v *UpdateSiteAccessTypeResponseBody) *UpdateSiteAccessTypeResponse {
	s.Body = v
	return s
}

type UpdateSiteCoverageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// global
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSiteCoverageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteCoverageRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteCoverageRequest) SetCoverage(v string) *UpdateSiteCoverageRequest {
	s.Coverage = &v
	return s
}

func (s *UpdateSiteCoverageRequest) SetSiteId(v int64) *UpdateSiteCoverageRequest {
	s.SiteId = &v
	return s
}

type UpdateSiteCoverageResponseBody struct {
	// example:
	//
	// 65C66B7B-671A-8297-9187-2R5477247B76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteCoverageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteCoverageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteCoverageResponseBody) SetRequestId(v string) *UpdateSiteCoverageResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSiteCoverageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteCoverageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteCoverageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteCoverageResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteCoverageResponse) SetHeaders(v map[string]*string) *UpdateSiteCoverageResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteCoverageResponse) SetStatusCode(v int32) *UpdateSiteCoverageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteCoverageResponse) SetBody(v *UpdateSiteCoverageResponseBody) *UpdateSiteCoverageResponse {
	s.Body = v
	return s
}

type UpdateSiteCustomLogRequest struct {
	Cookies         []*string `json:"Cookies,omitempty" xml:"Cookies,omitempty" type:"Repeated"`
	RequestHeaders  []*string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	ResponseHeaders []*string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
	// site id
	//
	// example:
	//
	// 11223****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSiteCustomLogRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteCustomLogRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteCustomLogRequest) SetCookies(v []*string) *UpdateSiteCustomLogRequest {
	s.Cookies = v
	return s
}

func (s *UpdateSiteCustomLogRequest) SetRequestHeaders(v []*string) *UpdateSiteCustomLogRequest {
	s.RequestHeaders = v
	return s
}

func (s *UpdateSiteCustomLogRequest) SetResponseHeaders(v []*string) *UpdateSiteCustomLogRequest {
	s.ResponseHeaders = v
	return s
}

func (s *UpdateSiteCustomLogRequest) SetSiteId(v int64) *UpdateSiteCustomLogRequest {
	s.SiteId = &v
	return s
}

type UpdateSiteCustomLogShrinkRequest struct {
	CookiesShrink         *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	RequestHeadersShrink  *string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	ResponseHeadersShrink *string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty"`
	// site id
	//
	// example:
	//
	// 11223****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSiteCustomLogShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteCustomLogShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteCustomLogShrinkRequest) SetCookiesShrink(v string) *UpdateSiteCustomLogShrinkRequest {
	s.CookiesShrink = &v
	return s
}

func (s *UpdateSiteCustomLogShrinkRequest) SetRequestHeadersShrink(v string) *UpdateSiteCustomLogShrinkRequest {
	s.RequestHeadersShrink = &v
	return s
}

func (s *UpdateSiteCustomLogShrinkRequest) SetResponseHeadersShrink(v string) *UpdateSiteCustomLogShrinkRequest {
	s.ResponseHeadersShrink = &v
	return s
}

func (s *UpdateSiteCustomLogShrinkRequest) SetSiteId(v int64) *UpdateSiteCustomLogShrinkRequest {
	s.SiteId = &v
	return s
}

type UpdateSiteCustomLogResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// ET5BF670-09D5-4D0B-BEBY-D96A2A528000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteCustomLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteCustomLogResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteCustomLogResponseBody) SetRequestId(v string) *UpdateSiteCustomLogResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSiteCustomLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteCustomLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteCustomLogResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteCustomLogResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteCustomLogResponse) SetHeaders(v map[string]*string) *UpdateSiteCustomLogResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteCustomLogResponse) SetStatusCode(v int32) *UpdateSiteCustomLogResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteCustomLogResponse) SetBody(v *UpdateSiteCustomLogResponseBody) *UpdateSiteCustomLogResponse {
	s.Body = v
	return s
}

type UpdateSiteDeliveryTaskRequest struct {
	// example:
	//
	// dcdn_log_er
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 0.0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ClientIP,UserAgent
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateSiteDeliveryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskRequest) SetBusinessType(v string) *UpdateSiteDeliveryTaskRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateSiteDeliveryTaskRequest) SetDiscardRate(v float32) *UpdateSiteDeliveryTaskRequest {
	s.DiscardRate = &v
	return s
}

func (s *UpdateSiteDeliveryTaskRequest) SetFieldName(v string) *UpdateSiteDeliveryTaskRequest {
	s.FieldName = &v
	return s
}

func (s *UpdateSiteDeliveryTaskRequest) SetSiteId(v int64) *UpdateSiteDeliveryTaskRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteDeliveryTaskRequest) SetTaskName(v string) *UpdateSiteDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

type UpdateSiteDeliveryTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteDeliveryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskResponseBody) SetRequestId(v string) *UpdateSiteDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSiteDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteDeliveryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskResponse) SetHeaders(v map[string]*string) *UpdateSiteDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteDeliveryTaskResponse) SetStatusCode(v int32) *UpdateSiteDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteDeliveryTaskResponse) SetBody(v *UpdateSiteDeliveryTaskResponseBody) *UpdateSiteDeliveryTaskResponse {
	s.Body = v
	return s
}

type UpdateSiteDeliveryTaskStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// online
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateSiteDeliveryTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteDeliveryTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskStatusRequest) SetMethod(v string) *UpdateSiteDeliveryTaskStatusRequest {
	s.Method = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusRequest) SetSiteId(v int64) *UpdateSiteDeliveryTaskStatusRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusRequest) SetTaskName(v string) *UpdateSiteDeliveryTaskStatusRequest {
	s.TaskName = &v
	return s
}

type UpdateSiteDeliveryTaskStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateSiteDeliveryTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteDeliveryTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskStatusResponseBody) SetRequestId(v string) *UpdateSiteDeliveryTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusResponseBody) SetStatus(v string) *UpdateSiteDeliveryTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusResponseBody) SetTaskName(v string) *UpdateSiteDeliveryTaskStatusResponseBody {
	s.TaskName = &v
	return s
}

type UpdateSiteDeliveryTaskStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteDeliveryTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteDeliveryTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteDeliveryTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskStatusResponse) SetHeaders(v map[string]*string) *UpdateSiteDeliveryTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusResponse) SetStatusCode(v int32) *UpdateSiteDeliveryTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusResponse) SetBody(v *UpdateSiteDeliveryTaskStatusResponseBody) *UpdateSiteDeliveryTaskStatusResponse {
	s.Body = v
	return s
}

type UpdateSiteVanityNSRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// ns1.example.com,ns2.example.com
	VanityNSList *string `json:"VanityNSList,omitempty" xml:"VanityNSList,omitempty"`
}

func (s UpdateSiteVanityNSRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteVanityNSRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteVanityNSRequest) SetSiteId(v int64) *UpdateSiteVanityNSRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteVanityNSRequest) SetVanityNSList(v string) *UpdateSiteVanityNSRequest {
	s.VanityNSList = &v
	return s
}

type UpdateSiteVanityNSResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteVanityNSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteVanityNSResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteVanityNSResponseBody) SetRequestId(v string) *UpdateSiteVanityNSResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSiteVanityNSResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteVanityNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteVanityNSResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSiteVanityNSResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteVanityNSResponse) SetHeaders(v map[string]*string) *UpdateSiteVanityNSResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteVanityNSResponse) SetStatusCode(v int32) *UpdateSiteVanityNSResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteVanityNSResponse) SetBody(v *UpdateSiteVanityNSResponseBody) *UpdateSiteVanityNSResponse {
	s.Body = v
	return s
}

type UpdateUserDeliveryTaskRequest struct {
	// example:
	//
	// dcdn_log_er
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ClientRequestID,ClientRequestHost
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateUserDeliveryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskRequest) SetBusinessType(v string) *UpdateUserDeliveryTaskRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateUserDeliveryTaskRequest) SetDiscardRate(v float32) *UpdateUserDeliveryTaskRequest {
	s.DiscardRate = &v
	return s
}

func (s *UpdateUserDeliveryTaskRequest) SetFieldName(v string) *UpdateUserDeliveryTaskRequest {
	s.FieldName = &v
	return s
}

func (s *UpdateUserDeliveryTaskRequest) SetTaskName(v string) *UpdateUserDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

type UpdateUserDeliveryTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserDeliveryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskResponseBody) SetRequestId(v string) *UpdateUserDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserDeliveryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskResponse) SetHeaders(v map[string]*string) *UpdateUserDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserDeliveryTaskResponse) SetStatusCode(v int32) *UpdateUserDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserDeliveryTaskResponse) SetBody(v *UpdateUserDeliveryTaskResponseBody) *UpdateUserDeliveryTaskResponse {
	s.Body = v
	return s
}

type UpdateUserDeliveryTaskStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// online
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateUserDeliveryTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDeliveryTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskStatusRequest) SetMethod(v string) *UpdateUserDeliveryTaskStatusRequest {
	s.Method = &v
	return s
}

func (s *UpdateUserDeliveryTaskStatusRequest) SetTaskName(v string) *UpdateUserDeliveryTaskStatusRequest {
	s.TaskName = &v
	return s
}

type UpdateUserDeliveryTaskStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateUserDeliveryTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDeliveryTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskStatusResponseBody) SetRequestId(v string) *UpdateUserDeliveryTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserDeliveryTaskStatusResponseBody) SetStatus(v string) *UpdateUserDeliveryTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateUserDeliveryTaskStatusResponseBody) SetTaskName(v string) *UpdateUserDeliveryTaskStatusResponseBody {
	s.TaskName = &v
	return s
}

type UpdateUserDeliveryTaskStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserDeliveryTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserDeliveryTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDeliveryTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskStatusResponse) SetHeaders(v map[string]*string) *UpdateUserDeliveryTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserDeliveryTaskStatusResponse) SetStatusCode(v int32) *UpdateUserDeliveryTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserDeliveryTaskStatusResponse) SetBody(v *UpdateUserDeliveryTaskStatusResponseBody) *UpdateUserDeliveryTaskStatusResponse {
	s.Body = v
	return s
}

type UpdateWafRuleRequest struct {
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateWafRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWafRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWafRuleRequest) SetConfig(v *WafRuleConfig) *UpdateWafRuleRequest {
	s.Config = v
	return s
}

func (s *UpdateWafRuleRequest) SetId(v int64) *UpdateWafRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateWafRuleRequest) SetPosition(v int64) *UpdateWafRuleRequest {
	s.Position = &v
	return s
}

func (s *UpdateWafRuleRequest) SetSiteId(v int64) *UpdateWafRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWafRuleRequest) SetSiteVersion(v int32) *UpdateWafRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateWafRuleRequest) SetStatus(v string) *UpdateWafRuleRequest {
	s.Status = &v
	return s
}

type UpdateWafRuleShrinkRequest struct {
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateWafRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWafRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWafRuleShrinkRequest) SetConfigShrink(v string) *UpdateWafRuleShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) SetId(v int64) *UpdateWafRuleShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) SetPosition(v int64) *UpdateWafRuleShrinkRequest {
	s.Position = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) SetSiteId(v int64) *UpdateWafRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) SetSiteVersion(v int32) *UpdateWafRuleShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) SetStatus(v string) *UpdateWafRuleShrinkRequest {
	s.Status = &v
	return s
}

type UpdateWafRuleResponseBody struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWafRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWafRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWafRuleResponseBody) SetId(v int64) *UpdateWafRuleResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateWafRuleResponseBody) SetRequestId(v string) *UpdateWafRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWafRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWafRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWafRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWafRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWafRuleResponse) SetHeaders(v map[string]*string) *UpdateWafRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWafRuleResponse) SetStatusCode(v int32) *UpdateWafRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWafRuleResponse) SetBody(v *UpdateWafRuleResponseBody) *UpdateWafRuleResponse {
	s.Body = v
	return s
}

type UpdateWafRulesetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateWafRulesetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *UpdateWafRulesetRequest) SetId(v int64) *UpdateWafRulesetRequest {
	s.Id = &v
	return s
}

func (s *UpdateWafRulesetRequest) SetSiteId(v int64) *UpdateWafRulesetRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWafRulesetRequest) SetSiteVersion(v int32) *UpdateWafRulesetRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateWafRulesetRequest) SetStatus(v string) *UpdateWafRulesetRequest {
	s.Status = &v
	return s
}

type UpdateWafRulesetResponseBody struct {
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWafRulesetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWafRulesetResponseBody) SetRequestId(v string) *UpdateWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWafRulesetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWafRulesetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *UpdateWafRulesetResponse) SetHeaders(v map[string]*string) *UpdateWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *UpdateWafRulesetResponse) SetStatusCode(v int32) *UpdateWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWafRulesetResponse) SetBody(v *UpdateWafRulesetResponseBody) *UpdateWafRulesetResponse {
	s.Body = v
	return s
}

type UpdateWaitingRoomRequest struct {
	// example:
	//
	// __aliwaitingroom_example
	CookieName *string `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	// example:
	//
	// Hello%20world!
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// on
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// example:
	//
	// on
	Enable          *string                                    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HostNameAndPath []*UpdateWaitingRoomRequestHostNameAndPath `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty" type:"Repeated"`
	// example:
	//
	// on
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// example:
	//
	// zhcn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 200
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// example:
	//
	// on
	QueueAllEnable *string `json:"QueueAllEnable,omitempty" xml:"QueueAllEnable,omitempty"`
	// example:
	//
	// random
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// example:
	//
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// example:
	//
	// 5
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 300
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6a51d5bc6460887abd129****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// example:
	//
	// default
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s UpdateWaitingRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRequest) SetCookieName(v string) *UpdateWaitingRoomRequest {
	s.CookieName = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetCustomPageHtml(v string) *UpdateWaitingRoomRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetDescription(v string) *UpdateWaitingRoomRequest {
	s.Description = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetDisableSessionRenewalEnable(v string) *UpdateWaitingRoomRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetEnable(v string) *UpdateWaitingRoomRequest {
	s.Enable = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetHostNameAndPath(v []*UpdateWaitingRoomRequestHostNameAndPath) *UpdateWaitingRoomRequest {
	s.HostNameAndPath = v
	return s
}

func (s *UpdateWaitingRoomRequest) SetJsonResponseEnable(v string) *UpdateWaitingRoomRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetLanguage(v string) *UpdateWaitingRoomRequest {
	s.Language = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetName(v string) *UpdateWaitingRoomRequest {
	s.Name = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetNewUsersPerMinute(v string) *UpdateWaitingRoomRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetQueueAllEnable(v string) *UpdateWaitingRoomRequest {
	s.QueueAllEnable = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetQueuingMethod(v string) *UpdateWaitingRoomRequest {
	s.QueuingMethod = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetQueuingStatusCode(v string) *UpdateWaitingRoomRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetSessionDuration(v string) *UpdateWaitingRoomRequest {
	s.SessionDuration = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetSiteId(v int64) *UpdateWaitingRoomRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetTotalActiveUsers(v string) *UpdateWaitingRoomRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetWaitingRoomId(v string) *UpdateWaitingRoomRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetWaitingRoomType(v string) *UpdateWaitingRoomRequest {
	s.WaitingRoomType = &v
	return s
}

type UpdateWaitingRoomRequestHostNameAndPath struct {
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// test.
	Subdomain *string `json:"Subdomain,omitempty" xml:"Subdomain,omitempty"`
}

func (s UpdateWaitingRoomRequestHostNameAndPath) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomRequestHostNameAndPath) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRequestHostNameAndPath) SetDomain(v string) *UpdateWaitingRoomRequestHostNameAndPath {
	s.Domain = &v
	return s
}

func (s *UpdateWaitingRoomRequestHostNameAndPath) SetPath(v string) *UpdateWaitingRoomRequestHostNameAndPath {
	s.Path = &v
	return s
}

func (s *UpdateWaitingRoomRequestHostNameAndPath) SetSubdomain(v string) *UpdateWaitingRoomRequestHostNameAndPath {
	s.Subdomain = &v
	return s
}

type UpdateWaitingRoomShrinkRequest struct {
	// example:
	//
	// __aliwaitingroom_example
	CookieName *string `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	// example:
	//
	// Hello%20world!
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// on
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// example:
	//
	// on
	Enable                *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HostNameAndPathShrink *string `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty"`
	// example:
	//
	// on
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// example:
	//
	// zhcn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 200
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// example:
	//
	// on
	QueueAllEnable *string `json:"QueueAllEnable,omitempty" xml:"QueueAllEnable,omitempty"`
	// example:
	//
	// random
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// example:
	//
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// example:
	//
	// 5
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 300
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6a51d5bc6460887abd129****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// example:
	//
	// default
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s UpdateWaitingRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomShrinkRequest) SetCookieName(v string) *UpdateWaitingRoomShrinkRequest {
	s.CookieName = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetCustomPageHtml(v string) *UpdateWaitingRoomShrinkRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetDescription(v string) *UpdateWaitingRoomShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetDisableSessionRenewalEnable(v string) *UpdateWaitingRoomShrinkRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetEnable(v string) *UpdateWaitingRoomShrinkRequest {
	s.Enable = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetHostNameAndPathShrink(v string) *UpdateWaitingRoomShrinkRequest {
	s.HostNameAndPathShrink = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetJsonResponseEnable(v string) *UpdateWaitingRoomShrinkRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetLanguage(v string) *UpdateWaitingRoomShrinkRequest {
	s.Language = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetName(v string) *UpdateWaitingRoomShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetNewUsersPerMinute(v string) *UpdateWaitingRoomShrinkRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetQueueAllEnable(v string) *UpdateWaitingRoomShrinkRequest {
	s.QueueAllEnable = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetQueuingMethod(v string) *UpdateWaitingRoomShrinkRequest {
	s.QueuingMethod = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetQueuingStatusCode(v string) *UpdateWaitingRoomShrinkRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetSessionDuration(v string) *UpdateWaitingRoomShrinkRequest {
	s.SessionDuration = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetSiteId(v int64) *UpdateWaitingRoomShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetTotalActiveUsers(v string) *UpdateWaitingRoomShrinkRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetWaitingRoomId(v string) *UpdateWaitingRoomShrinkRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetWaitingRoomType(v string) *UpdateWaitingRoomShrinkRequest {
	s.WaitingRoomType = &v
	return s
}

type UpdateWaitingRoomResponseBody struct {
	// example:
	//
	// 0195619f-eab3-4a66-ac00-ed53d913e72e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWaitingRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomResponseBody) SetRequestId(v string) *UpdateWaitingRoomResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWaitingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWaitingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWaitingRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomResponse) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomResponse) SetHeaders(v map[string]*string) *UpdateWaitingRoomResponse {
	s.Headers = v
	return s
}

func (s *UpdateWaitingRoomResponse) SetStatusCode(v int32) *UpdateWaitingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWaitingRoomResponse) SetBody(v *UpdateWaitingRoomResponseBody) *UpdateWaitingRoomResponse {
	s.Body = v
	return s
}

type UpdateWaitingRoomEventRequest struct {
	// example:
	//
	// html-yets-maqi1111
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	// example:
	//
	// http://yywyyw.com
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// off
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 1719849600
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// off
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// example:
	//
	// enus
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 300
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// example:
	//
	// on
	PreQueueEnable *string `json:"PreQueueEnable,omitempty" xml:"PreQueueEnable,omitempty"`
	// example:
	//
	// 1719763200
	PreQueueStartTime *string `json:"PreQueueStartTime,omitempty" xml:"PreQueueStartTime,omitempty"`
	// example:
	//
	// fifo
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// example:
	//
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// example:
	//
	// on
	RandomPreQueueEnable *string `json:"RandomPreQueueEnable,omitempty" xml:"RandomPreQueueEnable,omitempty"`
	// example:
	//
	// 5
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 1719763200
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 200
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// example:
	//
	// 89677721098****
	WaitingRoomEventId *int64 `json:"WaitingRoomEventId,omitempty" xml:"WaitingRoomEventId,omitempty"`
	// example:
	//
	// custom
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s UpdateWaitingRoomEventRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomEventRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomEventRequest) SetCustomPageHtml(v string) *UpdateWaitingRoomEventRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetDescription(v string) *UpdateWaitingRoomEventRequest {
	s.Description = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetDisableSessionRenewalEnable(v string) *UpdateWaitingRoomEventRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetEnable(v string) *UpdateWaitingRoomEventRequest {
	s.Enable = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetEndTime(v string) *UpdateWaitingRoomEventRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetJsonResponseEnable(v string) *UpdateWaitingRoomEventRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetLanguage(v string) *UpdateWaitingRoomEventRequest {
	s.Language = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetName(v string) *UpdateWaitingRoomEventRequest {
	s.Name = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetNewUsersPerMinute(v string) *UpdateWaitingRoomEventRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetPreQueueEnable(v string) *UpdateWaitingRoomEventRequest {
	s.PreQueueEnable = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetPreQueueStartTime(v string) *UpdateWaitingRoomEventRequest {
	s.PreQueueStartTime = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetQueuingMethod(v string) *UpdateWaitingRoomEventRequest {
	s.QueuingMethod = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetQueuingStatusCode(v string) *UpdateWaitingRoomEventRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetRandomPreQueueEnable(v string) *UpdateWaitingRoomEventRequest {
	s.RandomPreQueueEnable = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetSessionDuration(v string) *UpdateWaitingRoomEventRequest {
	s.SessionDuration = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetSiteId(v int64) *UpdateWaitingRoomEventRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetStartTime(v string) *UpdateWaitingRoomEventRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetTotalActiveUsers(v string) *UpdateWaitingRoomEventRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetWaitingRoomEventId(v int64) *UpdateWaitingRoomEventRequest {
	s.WaitingRoomEventId = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetWaitingRoomType(v string) *UpdateWaitingRoomEventRequest {
	s.WaitingRoomType = &v
	return s
}

type UpdateWaitingRoomEventResponseBody struct {
	// example:
	//
	// 0195619f-eab3-4a66-ac00-ed53d913e72e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWaitingRoomEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomEventResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomEventResponseBody) SetRequestId(v string) *UpdateWaitingRoomEventResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWaitingRoomEventResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWaitingRoomEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWaitingRoomEventResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomEventResponse) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomEventResponse) SetHeaders(v map[string]*string) *UpdateWaitingRoomEventResponse {
	s.Headers = v
	return s
}

func (s *UpdateWaitingRoomEventResponse) SetStatusCode(v int32) *UpdateWaitingRoomEventResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWaitingRoomEventResponse) SetBody(v *UpdateWaitingRoomEventResponseBody) *UpdateWaitingRoomEventResponse {
	s.Body = v
	return s
}

type UpdateWaitingRoomRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// (http.request.uri.path.file_name eq \\"jpg\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 8987739839****
	WaitingRoomRuleId *int64 `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s UpdateWaitingRoomRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRuleRequest) SetRule(v string) *UpdateWaitingRoomRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetRuleEnable(v string) *UpdateWaitingRoomRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetRuleName(v string) *UpdateWaitingRoomRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetSiteId(v int64) *UpdateWaitingRoomRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetWaitingRoomRuleId(v int64) *UpdateWaitingRoomRuleRequest {
	s.WaitingRoomRuleId = &v
	return s
}

type UpdateWaitingRoomRuleResponseBody struct {
	// example:
	//
	// 9bfe9d95-7bf6-469d-a628-ed7bc9f25073
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWaitingRoomRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRuleResponseBody) SetRequestId(v string) *UpdateWaitingRoomRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWaitingRoomRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWaitingRoomRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWaitingRoomRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaitingRoomRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRuleResponse) SetHeaders(v map[string]*string) *UpdateWaitingRoomRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWaitingRoomRuleResponse) SetStatusCode(v int32) *UpdateWaitingRoomRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWaitingRoomRuleResponse) SetBody(v *UpdateWaitingRoomRuleResponseBody) *UpdateWaitingRoomRuleResponse {
	s.Body = v
	return s
}

type UploadFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// purge_task_2024_11_11
	UploadTaskName *string `json:"UploadTaskName,omitempty" xml:"UploadTaskName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxxx.oss-cn-shenzhen.aliyuncs.com/test_oss_file?Expires=1708659191&OSSAccessKeyId=**********&Signature=**********
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UploadFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadFileRequest) GoString() string {
	return s.String()
}

func (s *UploadFileRequest) SetSiteId(v int64) *UploadFileRequest {
	s.SiteId = &v
	return s
}

func (s *UploadFileRequest) SetType(v string) *UploadFileRequest {
	s.Type = &v
	return s
}

func (s *UploadFileRequest) SetUploadTaskName(v string) *UploadFileRequest {
	s.UploadTaskName = &v
	return s
}

func (s *UploadFileRequest) SetUrl(v string) *UploadFileRequest {
	s.Url = &v
	return s
}

type UploadFileAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// purge_task_2024_11_11
	UploadTaskName *string `json:"UploadTaskName,omitempty" xml:"UploadTaskName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxxx.oss-cn-shenzhen.aliyuncs.com/test_oss_file?Expires=1708659191&OSSAccessKeyId=**********&Signature=**********
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UploadFileAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadFileAdvanceRequest) SetSiteId(v int64) *UploadFileAdvanceRequest {
	s.SiteId = &v
	return s
}

func (s *UploadFileAdvanceRequest) SetType(v string) *UploadFileAdvanceRequest {
	s.Type = &v
	return s
}

func (s *UploadFileAdvanceRequest) SetUploadTaskName(v string) *UploadFileAdvanceRequest {
	s.UploadTaskName = &v
	return s
}

func (s *UploadFileAdvanceRequest) SetUrlObject(v io.Reader) *UploadFileAdvanceRequest {
	s.UrlObject = v
	return s
}

type UploadFileResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 159253299357****
	UploadId *int64 `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s UploadFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadFileResponseBody) SetRequestId(v string) *UploadFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadFileResponseBody) SetUploadId(v int64) *UploadFileResponseBody {
	s.UploadId = &v
	return s
}

type UploadFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadFileResponse) GoString() string {
	return s.String()
}

func (s *UploadFileResponse) SetHeaders(v map[string]*string) *UploadFileResponse {
	s.Headers = v
	return s
}

func (s *UploadFileResponse) SetStatusCode(v int32) *UploadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadFileResponse) SetBody(v *UploadFileResponseBody) *UploadFileResponse {
	s.Body = v
	return s
}

type VerifySiteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s VerifySiteRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifySiteRequest) GoString() string {
	return s.String()
}

func (s *VerifySiteRequest) SetSiteId(v int64) *VerifySiteRequest {
	s.SiteId = &v
	return s
}

type VerifySiteResponseBody struct {
	// example:
	//
	// true
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 65C66B7B-671A-8297-9187-2R5477247B76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifySiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifySiteResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySiteResponseBody) SetPassed(v bool) *VerifySiteResponseBody {
	s.Passed = &v
	return s
}

func (s *VerifySiteResponseBody) SetRequestId(v string) *VerifySiteResponseBody {
	s.RequestId = &v
	return s
}

type VerifySiteResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifySiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifySiteResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifySiteResponse) GoString() string {
	return s.String()
}

func (s *VerifySiteResponse) SetHeaders(v map[string]*string) *VerifySiteResponse {
	s.Headers = v
	return s
}

func (s *VerifySiteResponse) SetStatusCode(v int32) *VerifySiteResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifySiteResponse) SetBody(v *VerifySiteResponseBody) *VerifySiteResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("esa"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 激活客户端证书
//
// @param request - ActivateClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateClientCertificateResponse
func (client *Client) ActivateClientCertificateWithOptions(request *ActivateClientCertificateRequest, runtime *util.RuntimeOptions) (_result *ActivateClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateClientCertificate"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 激活客户端证书
//
// @param request - ActivateClientCertificateRequest
//
// @return ActivateClientCertificateResponse
func (client *Client) ActivateClientCertificate(request *ActivateClientCertificateRequest) (_result *ActivateClientCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateClientCertificateResponse{}
	_body, _err := client.ActivateClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增用户商机单
//
// @param request - AddUserBusinessFormRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserBusinessFormResponse
func (client *Client) AddUserBusinessFormWithOptions(request *AddUserBusinessFormRequest, runtime *util.RuntimeOptions) (_result *AddUserBusinessFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Company)) {
		query["Company"] = request.Company
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Position)) {
		query["Position"] = request.Position
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.Website)) {
		query["Website"] = request.Website
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserBusinessForm"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserBusinessFormResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增用户商机单
//
// @param request - AddUserBusinessFormRequest
//
// @return AddUserBusinessFormResponse
func (client *Client) AddUserBusinessForm(request *AddUserBusinessFormRequest) (_result *AddUserBusinessFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserBusinessFormResponse{}
	_body, _err := client.AddUserBusinessFormWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 缓存对象缓存
//
// @param tmpReq - AdvancePurgeObjectCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AdvancePurgeObjectCacheResponse
func (client *Client) AdvancePurgeObjectCacheWithOptions(tmpReq *AdvancePurgeObjectCacheRequest, runtime *util.RuntimeOptions) (_result *AdvancePurgeObjectCacheResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AdvancePurgeObjectCacheShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Content)) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, tea.String("Content"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Area)) {
		query["Area"] = request.Area
	}

	if !tea.BoolValue(util.IsUnset(request.ContentShrink)) {
		query["Content"] = request.ContentShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Stations)) {
		query["Stations"] = request.Stations
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRangeBegin)) {
		query["TimeRangeBegin"] = request.TimeRangeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRangeEnd)) {
		query["TimeRangeEnd"] = request.TimeRangeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AdvancePurgeObjectCache"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AdvancePurgeObjectCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 缓存对象缓存
//
// @param request - AdvancePurgeObjectCacheRequest
//
// @return AdvancePurgeObjectCacheResponse
func (client *Client) AdvancePurgeObjectCache(request *AdvancePurgeObjectCacheRequest) (_result *AdvancePurgeObjectCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AdvancePurgeObjectCacheResponse{}
	_body, _err := client.AdvancePurgeObjectCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建记录
//
// @param tmpReq - BatchCreateRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateRecordsResponse
func (client *Client) BatchCreateRecordsWithOptions(tmpReq *BatchCreateRecordsRequest, runtime *util.RuntimeOptions) (_result *BatchCreateRecordsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchCreateRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RecordList)) {
		request.RecordListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordList, tea.String("RecordList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordListShrink)) {
		query["RecordList"] = request.RecordListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCreateRecords"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCreateRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建记录
//
// @param request - BatchCreateRecordsRequest
//
// @return BatchCreateRecordsResponse
func (client *Client) BatchCreateRecords(request *BatchCreateRecordsRequest) (_result *BatchCreateRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchCreateRecordsResponse{}
	_body, _err := client.BatchCreateRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量创建WAF规则
//
// @param tmpReq - BatchCreateWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateWafRulesResponse
func (client *Client) BatchCreateWafRulesWithOptions(tmpReq *BatchCreateWafRulesRequest, runtime *util.RuntimeOptions) (_result *BatchCreateWafRulesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchCreateWafRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Configs)) {
		request.ConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configs, tea.String("Configs"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Shared)) {
		request.SharedShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Shared, tea.String("Shared"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigsShrink)) {
		body["Configs"] = request.ConfigsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		body["Phase"] = request.Phase
	}

	if !tea.BoolValue(util.IsUnset(request.SharedShrink)) {
		body["Shared"] = request.SharedShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCreateWafRules"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCreateWafRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量创建WAF规则
//
// @param request - BatchCreateWafRulesRequest
//
// @return BatchCreateWafRulesResponse
func (client *Client) BatchCreateWafRules(request *BatchCreateWafRulesRequest) (_result *BatchCreateWafRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchCreateWafRulesResponse{}
	_body, _err := client.BatchCreateWafRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除Namespace的key-value对
//
// @param tmpReq - BatchDeleteKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteKvResponse
func (client *Client) BatchDeleteKvWithOptions(tmpReq *BatchDeleteKvRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteKvResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchDeleteKvShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Keys)) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, tea.String("Keys"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeysShrink)) {
		body["Keys"] = request.KeysShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteKv"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除Namespace的key-value对
//
// @param request - BatchDeleteKvRequest
//
// @return BatchDeleteKvResponse
func (client *Client) BatchDeleteKv(request *BatchDeleteKvRequest) (_result *BatchDeleteKvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteKvResponse{}
	_body, _err := client.BatchDeleteKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除Namespace下的KV队，支持大body的上传，上限100M
//
// @param request - BatchDeleteKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteKvWithHighCapacityResponse
func (client *Client) BatchDeleteKvWithHighCapacityWithOptions(request *BatchDeleteKvWithHighCapacityRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteKvWithHighCapacityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteKvWithHighCapacity"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteKvWithHighCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除Namespace下的KV队，支持大body的上传，上限100M
//
// @param request - BatchDeleteKvWithHighCapacityRequest
//
// @return BatchDeleteKvWithHighCapacityResponse
func (client *Client) BatchDeleteKvWithHighCapacity(request *BatchDeleteKvWithHighCapacityRequest) (_result *BatchDeleteKvWithHighCapacityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteKvWithHighCapacityResponse{}
	_body, _err := client.BatchDeleteKvWithHighCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteKvWithHighCapacityAdvance(request *BatchDeleteKvWithHighCapacityAdvanceRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteKvWithHighCapacityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ESA"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	batchDeleteKvWithHighCapacityReq := &BatchDeleteKvWithHighCapacityRequest{}
	openapiutil.Convert(request, batchDeleteKvWithHighCapacityReq)
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		batchDeleteKvWithHighCapacityReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	batchDeleteKvWithHighCapacityResp, _err := client.BatchDeleteKvWithHighCapacityWithOptions(batchDeleteKvWithHighCapacityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = batchDeleteKvWithHighCapacityResp
	return _result, _err
}

// Summary:
//
// 批量获取表达式的匹配项
//
// @param tmpReq - BatchGetExpressionFieldsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetExpressionFieldsResponse
func (client *Client) BatchGetExpressionFieldsWithOptions(tmpReq *BatchGetExpressionFieldsRequest, runtime *util.RuntimeOptions) (_result *BatchGetExpressionFieldsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchGetExpressionFieldsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Expressions)) {
		request.ExpressionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Expressions, tea.String("Expressions"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpressionsShrink)) {
		body["Expressions"] = request.ExpressionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		body["Phase"] = request.Phase
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGetExpressionFields"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetExpressionFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取表达式的匹配项
//
// @param request - BatchGetExpressionFieldsRequest
//
// @return BatchGetExpressionFieldsResponse
func (client *Client) BatchGetExpressionFields(request *BatchGetExpressionFieldsRequest) (_result *BatchGetExpressionFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchGetExpressionFieldsResponse{}
	_body, _err := client.BatchGetExpressionFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量设置Namespace的key-value对
//
// @param tmpReq - BatchPutKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchPutKvResponse
func (client *Client) BatchPutKvWithOptions(tmpReq *BatchPutKvRequest, runtime *util.RuntimeOptions) (_result *BatchPutKvResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchPutKvShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.KvList)) {
		request.KvListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KvList, tea.String("KvList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KvListShrink)) {
		body["KvList"] = request.KvListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchPutKv"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchPutKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量设置Namespace的key-value对
//
// @param request - BatchPutKvRequest
//
// @return BatchPutKvResponse
func (client *Client) BatchPutKv(request *BatchPutKvRequest) (_result *BatchPutKvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchPutKvResponse{}
	_body, _err := client.BatchPutKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量设置Namespace的key-value对，支持最大100M的请求体
//
// @param request - BatchPutKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchPutKvWithHighCapacityResponse
func (client *Client) BatchPutKvWithHighCapacityWithOptions(request *BatchPutKvWithHighCapacityRequest, runtime *util.RuntimeOptions) (_result *BatchPutKvWithHighCapacityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchPutKvWithHighCapacity"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchPutKvWithHighCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量设置Namespace的key-value对，支持最大100M的请求体
//
// @param request - BatchPutKvWithHighCapacityRequest
//
// @return BatchPutKvWithHighCapacityResponse
func (client *Client) BatchPutKvWithHighCapacity(request *BatchPutKvWithHighCapacityRequest) (_result *BatchPutKvWithHighCapacityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchPutKvWithHighCapacityResponse{}
	_body, _err := client.BatchPutKvWithHighCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchPutKvWithHighCapacityAdvance(request *BatchPutKvWithHighCapacityAdvanceRequest, runtime *util.RuntimeOptions) (_result *BatchPutKvWithHighCapacityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ESA"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	batchPutKvWithHighCapacityReq := &BatchPutKvWithHighCapacityRequest{}
	openapiutil.Convert(request, batchPutKvWithHighCapacityReq)
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		batchPutKvWithHighCapacityReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	batchPutKvWithHighCapacityResp, _err := client.BatchPutKvWithHighCapacityWithOptions(batchPutKvWithHighCapacityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = batchPutKvWithHighCapacityResp
	return _result, _err
}

// Summary:
//
// 批量修改WAF规则
//
// @param tmpReq - BatchUpdateWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateWafRulesResponse
func (client *Client) BatchUpdateWafRulesWithOptions(tmpReq *BatchUpdateWafRulesRequest, runtime *util.RuntimeOptions) (_result *BatchUpdateWafRulesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchUpdateWafRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Configs)) {
		request.ConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configs, tea.String("Configs"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Shared)) {
		request.SharedShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Shared, tea.String("Shared"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigsShrink)) {
		body["Configs"] = request.ConfigsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		body["Phase"] = request.Phase
	}

	if !tea.BoolValue(util.IsUnset(request.RulesetId)) {
		body["RulesetId"] = request.RulesetId
	}

	if !tea.BoolValue(util.IsUnset(request.SharedShrink)) {
		body["Shared"] = request.SharedShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUpdateWafRules"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdateWafRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量修改WAF规则
//
// @param request - BatchUpdateWafRulesRequest
//
// @return BatchUpdateWafRulesResponse
func (client *Client) BatchUpdateWafRules(request *BatchUpdateWafRulesRequest) (_result *BatchUpdateWafRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUpdateWafRulesResponse{}
	_body, _err := client.BatchUpdateWafRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// URL封禁
//
// @param tmpReq - BlockObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BlockObjectResponse
func (client *Client) BlockObjectWithOptions(tmpReq *BlockObjectRequest, runtime *util.RuntimeOptions) (_result *BlockObjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BlockObjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Content)) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, tea.String("Content"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentShrink)) {
		query["Content"] = request.ContentShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		query["Extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.Maxage)) {
		query["Maxage"] = request.Maxage
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BlockObject"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BlockObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// URL封禁
//
// @param request - BlockObjectRequest
//
// @return BlockObjectResponse
func (client *Client) BlockObject(request *BlockObjectRequest) (_result *BlockObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BlockObjectResponse{}
	_body, _err := client.BlockObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改站点的企业资源组
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改站点的企业资源组
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验站点名称是否可用
//
// @param request - CheckSiteNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSiteNameResponse
func (client *Client) CheckSiteNameWithOptions(request *CheckSiteNameRequest, runtime *util.RuntimeOptions) (_result *CheckSiteNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteName)) {
		query["SiteName"] = request.SiteName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckSiteName"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckSiteNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验站点名称是否可用
//
// @param request - CheckSiteNameRequest
//
// @return CheckSiteNameResponse
func (client *Client) CheckSiteName(request *CheckSiteNameRequest) (_result *CheckSiteNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckSiteNameResponse{}
	_body, _err := client.CheckSiteNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实时日志任务投递名检查
//
// @param request - CheckSiteProjectNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSiteProjectNameResponse
func (client *Client) CheckSiteProjectNameWithOptions(request *CheckSiteProjectNameRequest, runtime *util.RuntimeOptions) (_result *CheckSiteProjectNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckSiteProjectName"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckSiteProjectNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时日志任务投递名检查
//
// @param request - CheckSiteProjectNameRequest
//
// @return CheckSiteProjectNameResponse
func (client *Client) CheckSiteProjectName(request *CheckSiteProjectNameRequest) (_result *CheckSiteProjectNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckSiteProjectNameResponse{}
	_body, _err := client.CheckSiteProjectNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实时日志用户任务投递名检查
//
// @param request - CheckUserProjectNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUserProjectNameResponse
func (client *Client) CheckUserProjectNameWithOptions(request *CheckUserProjectNameRequest, runtime *util.RuntimeOptions) (_result *CheckUserProjectNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckUserProjectName"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUserProjectNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时日志用户任务投递名检查
//
// @param request - CheckUserProjectNameRequest
//
// @return CheckUserProjectNameResponse
func (client *Client) CheckUserProjectName(request *CheckUserProjectNameRequest) (_result *CheckUserProjectNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUserProjectNameResponse{}
	_body, _err := client.CheckUserProjectNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建定制场景策略
//
// @param request - CreateCustomScenePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomScenePolicyResponse
func (client *Client) CreateCustomScenePolicyWithOptions(request *CreateCustomScenePolicyRequest, runtime *util.RuntimeOptions) (_result *CreateCustomScenePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Objects)) {
		query["Objects"] = request.Objects
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Template)) {
		query["Template"] = request.Template
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomScenePolicy"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomScenePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建定制场景策略
//
// @param request - CreateCustomScenePolicyRequest
//
// @return CreateCustomScenePolicyResponse
func (client *Client) CreateCustomScenePolicy(request *CreateCustomScenePolicyRequest) (_result *CreateCustomScenePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomScenePolicyResponse{}
	_body, _err := client.CreateCustomScenePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加Namespace
//
// @param request - CreateKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKvNamespaceResponse
func (client *Client) CreateKvNamespaceWithOptions(request *CreateKvNamespaceRequest, runtime *util.RuntimeOptions) (_result *CreateKvNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateKvNamespace"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateKvNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加Namespace
//
// @param request - CreateKvNamespaceRequest
//
// @return CreateKvNamespaceResponse
func (client *Client) CreateKvNamespace(request *CreateKvNamespaceRequest) (_result *CreateKvNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKvNamespaceResponse{}
	_body, _err := client.CreateKvNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建自定义列表
//
// @param tmpReq - CreateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateListResponse
func (client *Client) CreateListWithOptions(tmpReq *CreateListRequest, runtime *util.RuntimeOptions) (_result *CreateListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Items)) {
		request.ItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Items, tea.String("Items"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ItemsShrink)) {
		body["Items"] = request.ItemsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Kind)) {
		body["Kind"] = request.Kind
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateList"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义列表
//
// @param request - CreateListRequest
//
// @return CreateListResponse
func (client *Client) CreateList(request *CreateListRequest) (_result *CreateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateListResponse{}
	_body, _err := client.CreateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用CreatePage创建自定义响应页面
//
// @param request - CreatePageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePageResponse
func (client *Client) CreatePageWithOptions(request *CreatePageRequest, runtime *util.RuntimeOptions) (_result *CreatePageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["ContentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePage"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用CreatePage创建自定义响应页面
//
// @param request - CreatePageRequest
//
// @return CreatePageResponse
func (client *Client) CreatePage(request *CreatePageRequest) (_result *CreatePageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePageResponse{}
	_body, _err := client.CreatePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建记录
//
// @param tmpReq - CreateRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecordResponse
func (client *Client) CreateRecordWithOptions(tmpReq *CreateRecordRequest, runtime *util.RuntimeOptions) (_result *CreateRecordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AuthConf)) {
		request.AuthConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthConf, tea.String("AuthConf"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Data)) {
		request.DataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Data, tea.String("Data"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthConfShrink)) {
		query["AuthConf"] = request.AuthConfShrink
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		query["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.DataShrink)) {
		query["Data"] = request.DataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HostPolicy)) {
		query["HostPolicy"] = request.HostPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Proxied)) {
		query["Proxied"] = request.Proxied
	}

	if !tea.BoolValue(util.IsUnset(request.RecordName)) {
		query["RecordName"] = request.RecordName
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		query["Ttl"] = request.Ttl
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRecord"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建记录
//
// @param request - CreateRecordRequest
//
// @return CreateRecordResponse
func (client *Client) CreateRecord(request *CreateRecordRequest) (_result *CreateRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRecordResponse{}
	_body, _err := client.CreateRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量新增定时预热任务的计划
//
// @param tmpReq - CreateScheduledPreloadExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledPreloadExecutionsResponse
func (client *Client) CreateScheduledPreloadExecutionsWithOptions(tmpReq *CreateScheduledPreloadExecutionsRequest, runtime *util.RuntimeOptions) (_result *CreateScheduledPreloadExecutionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateScheduledPreloadExecutionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Executions)) {
		request.ExecutionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Executions, tea.String("Executions"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutionsShrink)) {
		body["Executions"] = request.ExecutionsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScheduledPreloadExecutions"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScheduledPreloadExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量新增定时预热任务的计划
//
// @param request - CreateScheduledPreloadExecutionsRequest
//
// @return CreateScheduledPreloadExecutionsResponse
func (client *Client) CreateScheduledPreloadExecutions(request *CreateScheduledPreloadExecutionsRequest) (_result *CreateScheduledPreloadExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScheduledPreloadExecutionsResponse{}
	_body, _err := client.CreateScheduledPreloadExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增定时预热任务
//
// @param request - CreateScheduledPreloadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledPreloadJobResponse
func (client *Client) CreateScheduledPreloadJobWithOptions(request *CreateScheduledPreloadJobRequest, runtime *util.RuntimeOptions) (_result *CreateScheduledPreloadJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InsertWay)) {
		body["InsertWay"] = request.InsertWay
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssUrl)) {
		body["OssUrl"] = request.OssUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		body["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.UrlList)) {
		body["UrlList"] = request.UrlList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScheduledPreloadJob"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScheduledPreloadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增定时预热任务
//
// @param request - CreateScheduledPreloadJobRequest
//
// @return CreateScheduledPreloadJobResponse
func (client *Client) CreateScheduledPreloadJob(request *CreateScheduledPreloadJobRequest) (_result *CreateScheduledPreloadJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScheduledPreloadJobResponse{}
	_body, _err := client.CreateScheduledPreloadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建站点
//
// @param request - CreateSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSiteResponse
func (client *Client) CreateSiteWithOptions(request *CreateSiteRequest, runtime *util.RuntimeOptions) (_result *CreateSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.Coverage)) {
		query["Coverage"] = request.Coverage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteName)) {
		query["SiteName"] = request.SiteName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSite"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建站点
//
// @param request - CreateSiteRequest
//
// @return CreateSiteResponse
func (client *Client) CreateSite(request *CreateSiteRequest) (_result *CreateSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSiteResponse{}
	_body, _err := client.CreateSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建自定义字段
//
// @param tmpReq - CreateSiteCustomLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSiteCustomLogResponse
func (client *Client) CreateSiteCustomLogWithOptions(tmpReq *CreateSiteCustomLogRequest, runtime *util.RuntimeOptions) (_result *CreateSiteCustomLogResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSiteCustomLogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Cookies)) {
		request.CookiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cookies, tea.String("Cookies"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RequestHeaders)) {
		request.RequestHeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestHeaders, tea.String("RequestHeaders"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResponseHeaders)) {
		request.ResponseHeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResponseHeaders, tea.String("ResponseHeaders"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CookiesShrink)) {
		body["Cookies"] = request.CookiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RequestHeadersShrink)) {
		body["RequestHeaders"] = request.RequestHeadersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseHeadersShrink)) {
		body["ResponseHeaders"] = request.ResponseHeadersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		body["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSiteCustomLog"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSiteCustomLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建自定义字段
//
// @param request - CreateSiteCustomLogRequest
//
// @return CreateSiteCustomLogResponse
func (client *Client) CreateSiteCustomLog(request *CreateSiteCustomLogRequest) (_result *CreateSiteCustomLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSiteCustomLogResponse{}
	_body, _err := client.CreateSiteCustomLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建一个任务投递
//
// @param tmpReq - CreateSiteDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSiteDeliveryTaskResponse
func (client *Client) CreateSiteDeliveryTaskWithOptions(tmpReq *CreateSiteDeliveryTaskRequest, runtime *util.RuntimeOptions) (_result *CreateSiteDeliveryTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSiteDeliveryTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HttpDelivery)) {
		request.HttpDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HttpDelivery, tea.String("HttpDelivery"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.KafkaDelivery)) {
		request.KafkaDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KafkaDelivery, tea.String("KafkaDelivery"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OssDelivery)) {
		request.OssDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OssDelivery, tea.String("OssDelivery"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.S3Delivery)) {
		request.S3DeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.S3Delivery, tea.String("S3Delivery"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SlsDelivery)) {
		request.SlsDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SlsDelivery, tea.String("SlsDelivery"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		body["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.DataCenter)) {
		body["DataCenter"] = request.DataCenter
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryType)) {
		body["DeliveryType"] = request.DeliveryType
	}

	if !tea.BoolValue(util.IsUnset(request.DiscardRate)) {
		body["DiscardRate"] = request.DiscardRate
	}

	if !tea.BoolValue(util.IsUnset(request.FieldName)) {
		body["FieldName"] = request.FieldName
	}

	if !tea.BoolValue(util.IsUnset(request.HttpDeliveryShrink)) {
		body["HttpDelivery"] = request.HttpDeliveryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.KafkaDeliveryShrink)) {
		body["KafkaDelivery"] = request.KafkaDeliveryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OssDeliveryShrink)) {
		body["OssDelivery"] = request.OssDeliveryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.S3DeliveryShrink)) {
		body["S3Delivery"] = request.S3DeliveryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		body["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SlsDeliveryShrink)) {
		body["SlsDelivery"] = request.SlsDeliveryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSiteDeliveryTask"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSiteDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建一个任务投递
//
// @param request - CreateSiteDeliveryTaskRequest
//
// @return CreateSiteDeliveryTaskResponse
func (client *Client) CreateSiteDeliveryTask(request *CreateSiteDeliveryTaskRequest) (_result *CreateSiteDeliveryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSiteDeliveryTaskResponse{}
	_body, _err := client.CreateSiteDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建一个用户粒度任务投递
//
// @param tmpReq - CreateUserDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserDeliveryTaskResponse
func (client *Client) CreateUserDeliveryTaskWithOptions(tmpReq *CreateUserDeliveryTaskRequest, runtime *util.RuntimeOptions) (_result *CreateUserDeliveryTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateUserDeliveryTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HttpDelivery)) {
		request.HttpDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HttpDelivery, tea.String("HttpDelivery"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.KafkaDelivery)) {
		request.KafkaDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KafkaDelivery, tea.String("KafkaDelivery"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OssDelivery)) {
		request.OssDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OssDelivery, tea.String("OssDelivery"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.S3Delivery)) {
		request.S3DeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.S3Delivery, tea.String("S3Delivery"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SlsDelivery)) {
		request.SlsDeliveryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SlsDelivery, tea.String("SlsDelivery"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		body["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.DataCenter)) {
		body["DataCenter"] = request.DataCenter
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryType)) {
		body["DeliveryType"] = request.DeliveryType
	}

	if !tea.BoolValue(util.IsUnset(request.DiscardRate)) {
		body["DiscardRate"] = request.DiscardRate
	}

	if !tea.BoolValue(util.IsUnset(request.FieldName)) {
		body["FieldName"] = request.FieldName
	}

	if !tea.BoolValue(util.IsUnset(request.HttpDeliveryShrink)) {
		body["HttpDelivery"] = request.HttpDeliveryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.KafkaDeliveryShrink)) {
		body["KafkaDelivery"] = request.KafkaDeliveryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OssDeliveryShrink)) {
		body["OssDelivery"] = request.OssDeliveryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.S3DeliveryShrink)) {
		body["S3Delivery"] = request.S3DeliveryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SlsDeliveryShrink)) {
		body["SlsDelivery"] = request.SlsDeliveryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserDeliveryTask"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建一个用户粒度任务投递
//
// @param request - CreateUserDeliveryTaskRequest
//
// @return CreateUserDeliveryTaskResponse
func (client *Client) CreateUserDeliveryTask(request *CreateUserDeliveryTaskRequest) (_result *CreateUserDeliveryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserDeliveryTaskResponse{}
	_body, _err := client.CreateUserDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建WAF规则
//
// @param tmpReq - CreateWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWafRuleResponse
func (client *Client) CreateWafRuleWithOptions(tmpReq *CreateWafRuleRequest, runtime *util.RuntimeOptions) (_result *CreateWafRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateWafRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Config)) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, tea.String("Config"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigShrink)) {
		body["Config"] = request.ConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		body["Phase"] = request.Phase
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWafRule"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWafRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建WAF规则
//
// @param request - CreateWafRuleRequest
//
// @return CreateWafRuleResponse
func (client *Client) CreateWafRule(request *CreateWafRuleRequest) (_result *CreateWafRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWafRuleResponse{}
	_body, _err := client.CreateWafRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建等候室
//
// @param tmpReq - CreateWaitingRoomRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWaitingRoomResponse
func (client *Client) CreateWaitingRoomWithOptions(tmpReq *CreateWaitingRoomRequest, runtime *util.RuntimeOptions) (_result *CreateWaitingRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateWaitingRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HostNameAndPath)) {
		request.HostNameAndPathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HostNameAndPath, tea.String("HostNameAndPath"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CookieName)) {
		query["CookieName"] = request.CookieName
	}

	if !tea.BoolValue(util.IsUnset(request.CustomPageHtml)) {
		query["CustomPageHtml"] = request.CustomPageHtml
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisableSessionRenewalEnable)) {
		query["DisableSessionRenewalEnable"] = request.DisableSessionRenewalEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.HostNameAndPathShrink)) {
		query["HostNameAndPath"] = request.HostNameAndPathShrink
	}

	if !tea.BoolValue(util.IsUnset(request.JsonResponseEnable)) {
		query["JsonResponseEnable"] = request.JsonResponseEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NewUsersPerMinute)) {
		query["NewUsersPerMinute"] = request.NewUsersPerMinute
	}

	if !tea.BoolValue(util.IsUnset(request.QueueAllEnable)) {
		query["QueueAllEnable"] = request.QueueAllEnable
	}

	if !tea.BoolValue(util.IsUnset(request.QueuingMethod)) {
		query["QueuingMethod"] = request.QueuingMethod
	}

	if !tea.BoolValue(util.IsUnset(request.QueuingStatusCode)) {
		query["QueuingStatusCode"] = request.QueuingStatusCode
	}

	if !tea.BoolValue(util.IsUnset(request.SessionDuration)) {
		query["SessionDuration"] = request.SessionDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalActiveUsers)) {
		query["TotalActiveUsers"] = request.TotalActiveUsers
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomType)) {
		query["WaitingRoomType"] = request.WaitingRoomType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWaitingRoom"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWaitingRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建等候室
//
// @param request - CreateWaitingRoomRequest
//
// @return CreateWaitingRoomResponse
func (client *Client) CreateWaitingRoom(request *CreateWaitingRoomRequest) (_result *CreateWaitingRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWaitingRoomResponse{}
	_body, _err := client.CreateWaitingRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建等候室事件
//
// @param request - CreateWaitingRoomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWaitingRoomEventResponse
func (client *Client) CreateWaitingRoomEventWithOptions(request *CreateWaitingRoomEventRequest, runtime *util.RuntimeOptions) (_result *CreateWaitingRoomEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomPageHtml)) {
		query["CustomPageHtml"] = request.CustomPageHtml
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisableSessionRenewalEnable)) {
		query["DisableSessionRenewalEnable"] = request.DisableSessionRenewalEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.JsonResponseEnable)) {
		query["JsonResponseEnable"] = request.JsonResponseEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NewUsersPerMinute)) {
		query["NewUsersPerMinute"] = request.NewUsersPerMinute
	}

	if !tea.BoolValue(util.IsUnset(request.PreQueueEnable)) {
		query["PreQueueEnable"] = request.PreQueueEnable
	}

	if !tea.BoolValue(util.IsUnset(request.PreQueueStartTime)) {
		query["PreQueueStartTime"] = request.PreQueueStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.QueuingMethod)) {
		query["QueuingMethod"] = request.QueuingMethod
	}

	if !tea.BoolValue(util.IsUnset(request.QueuingStatusCode)) {
		query["QueuingStatusCode"] = request.QueuingStatusCode
	}

	if !tea.BoolValue(util.IsUnset(request.RandomPreQueueEnable)) {
		query["RandomPreQueueEnable"] = request.RandomPreQueueEnable
	}

	if !tea.BoolValue(util.IsUnset(request.SessionDuration)) {
		query["SessionDuration"] = request.SessionDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TotalActiveUsers)) {
		query["TotalActiveUsers"] = request.TotalActiveUsers
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomId)) {
		query["WaitingRoomId"] = request.WaitingRoomId
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomType)) {
		query["WaitingRoomType"] = request.WaitingRoomType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWaitingRoomEvent"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWaitingRoomEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建等候室事件
//
// @param request - CreateWaitingRoomEventRequest
//
// @return CreateWaitingRoomEventResponse
func (client *Client) CreateWaitingRoomEvent(request *CreateWaitingRoomEventRequest) (_result *CreateWaitingRoomEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWaitingRoomEventResponse{}
	_body, _err := client.CreateWaitingRoomEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建等候室规则
//
// @param request - CreateWaitingRoomRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWaitingRoomRuleResponse
func (client *Client) CreateWaitingRoomRuleWithOptions(request *CreateWaitingRoomRuleRequest, runtime *util.RuntimeOptions) (_result *CreateWaitingRoomRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Rule)) {
		query["Rule"] = request.Rule
	}

	if !tea.BoolValue(util.IsUnset(request.RuleEnable)) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomId)) {
		query["WaitingRoomId"] = request.WaitingRoomId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWaitingRoomRule"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWaitingRoomRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建等候室规则
//
// @param request - CreateWaitingRoomRuleRequest
//
// @return CreateWaitingRoomRuleResponse
func (client *Client) CreateWaitingRoomRule(request *CreateWaitingRoomRuleRequest) (_result *CreateWaitingRoomRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWaitingRoomRuleResponse{}
	_body, _err := client.CreateWaitingRoomRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除定制场景策略
//
// @param request - DeleteCustomScenePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomScenePolicyResponse
func (client *Client) DeleteCustomScenePolicyWithOptions(request *DeleteCustomScenePolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomScenePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomScenePolicy"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomScenePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除定制场景策略
//
// @param request - DeleteCustomScenePolicyRequest
//
// @return DeleteCustomScenePolicyResponse
func (client *Client) DeleteCustomScenePolicy(request *DeleteCustomScenePolicyRequest) (_result *DeleteCustomScenePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomScenePolicyResponse{}
	_body, _err := client.DeleteCustomScenePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义列表
//
// @param request - DeleteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteListResponse
func (client *Client) DeleteListWithOptions(request *DeleteListRequest, runtime *util.RuntimeOptions) (_result *DeleteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteList"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义列表
//
// @param request - DeleteListRequest
//
// @return DeleteListResponse
func (client *Client) DeleteList(request *DeleteListRequest) (_result *DeleteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteListResponse{}
	_body, _err := client.DeleteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义响应页面
//
// @param request - DeletePageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePageResponse
func (client *Client) DeletePageWithOptions(request *DeletePageRequest, runtime *util.RuntimeOptions) (_result *DeletePageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePage"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义响应页面
//
// @param request - DeletePageRequest
//
// @return DeletePageResponse
func (client *Client) DeletePage(request *DeletePageRequest) (_result *DeletePageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePageResponse{}
	_body, _err := client.DeletePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除记录
//
// @param request - DeleteRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecordResponse
func (client *Client) DeleteRecordWithOptions(request *DeleteRecordRequest, runtime *util.RuntimeOptions) (_result *DeleteRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRecord"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除记录
//
// @param request - DeleteRecordRequest
//
// @return DeleteRecordResponse
func (client *Client) DeleteRecord(request *DeleteRecordRequest) (_result *DeleteRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRecordResponse{}
	_body, _err := client.DeleteRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除单个定时预热计划
//
// @param request - DeleteScheduledPreloadExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledPreloadExecutionResponse
func (client *Client) DeleteScheduledPreloadExecutionWithOptions(request *DeleteScheduledPreloadExecutionRequest, runtime *util.RuntimeOptions) (_result *DeleteScheduledPreloadExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScheduledPreloadExecution"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScheduledPreloadExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除单个定时预热计划
//
// @param request - DeleteScheduledPreloadExecutionRequest
//
// @return DeleteScheduledPreloadExecutionResponse
func (client *Client) DeleteScheduledPreloadExecution(request *DeleteScheduledPreloadExecutionRequest) (_result *DeleteScheduledPreloadExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScheduledPreloadExecutionResponse{}
	_body, _err := client.DeleteScheduledPreloadExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定定时预热任务
//
// @param request - DeleteScheduledPreloadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledPreloadJobResponse
func (client *Client) DeleteScheduledPreloadJobWithOptions(request *DeleteScheduledPreloadJobRequest, runtime *util.RuntimeOptions) (_result *DeleteScheduledPreloadJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScheduledPreloadJob"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScheduledPreloadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定定时预热任务
//
// @param request - DeleteScheduledPreloadJobRequest
//
// @return DeleteScheduledPreloadJobResponse
func (client *Client) DeleteScheduledPreloadJob(request *DeleteScheduledPreloadJobRequest) (_result *DeleteScheduledPreloadJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScheduledPreloadJobResponse{}
	_body, _err := client.DeleteScheduledPreloadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除站点
//
// @param request - DeleteSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSiteResponse
func (client *Client) DeleteSiteWithOptions(request *DeleteSiteRequest, runtime *util.RuntimeOptions) (_result *DeleteSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSite"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除站点
//
// @param request - DeleteSiteRequest
//
// @return DeleteSiteResponse
func (client *Client) DeleteSite(request *DeleteSiteRequest) (_result *DeleteSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSiteResponse{}
	_body, _err := client.DeleteSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除一个任务投递
//
// @param request - DeleteSiteDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSiteDeliveryTaskResponse
func (client *Client) DeleteSiteDeliveryTaskWithOptions(request *DeleteSiteDeliveryTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteSiteDeliveryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		body["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSiteDeliveryTask"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSiteDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个任务投递
//
// @param request - DeleteSiteDeliveryTaskRequest
//
// @return DeleteSiteDeliveryTaskResponse
func (client *Client) DeleteSiteDeliveryTask(request *DeleteSiteDeliveryTaskRequest) (_result *DeleteSiteDeliveryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSiteDeliveryTaskResponse{}
	_body, _err := client.DeleteSiteDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除一个用户任务投递
//
// @param request - DeleteUserDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserDeliveryTaskResponse
func (client *Client) DeleteUserDeliveryTaskWithOptions(request *DeleteUserDeliveryTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteUserDeliveryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserDeliveryTask"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个用户任务投递
//
// @param request - DeleteUserDeliveryTaskRequest
//
// @return DeleteUserDeliveryTaskResponse
func (client *Client) DeleteUserDeliveryTask(request *DeleteUserDeliveryTaskRequest) (_result *DeleteUserDeliveryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserDeliveryTaskResponse{}
	_body, _err := client.DeleteUserDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除WAF规则
//
// @param request - DeleteWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWafRuleResponse
func (client *Client) DeleteWafRuleWithOptions(request *DeleteWafRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteWafRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWafRule"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWafRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除WAF规则
//
// @param request - DeleteWafRuleRequest
//
// @return DeleteWafRuleResponse
func (client *Client) DeleteWafRule(request *DeleteWafRuleRequest) (_result *DeleteWafRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWafRuleResponse{}
	_body, _err := client.DeleteWafRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除WAF规则集
//
// @param request - DeleteWafRulesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWafRulesetResponse
func (client *Client) DeleteWafRulesetWithOptions(request *DeleteWafRulesetRequest, runtime *util.RuntimeOptions) (_result *DeleteWafRulesetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWafRuleset"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWafRulesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除WAF规则集
//
// @param request - DeleteWafRulesetRequest
//
// @return DeleteWafRulesetResponse
func (client *Client) DeleteWafRuleset(request *DeleteWafRulesetRequest) (_result *DeleteWafRulesetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWafRulesetResponse{}
	_body, _err := client.DeleteWafRulesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除等候室
//
// @param request - DeleteWaitingRoomRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWaitingRoomResponse
func (client *Client) DeleteWaitingRoomWithOptions(request *DeleteWaitingRoomRequest, runtime *util.RuntimeOptions) (_result *DeleteWaitingRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomId)) {
		query["WaitingRoomId"] = request.WaitingRoomId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWaitingRoom"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWaitingRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除等候室
//
// @param request - DeleteWaitingRoomRequest
//
// @return DeleteWaitingRoomResponse
func (client *Client) DeleteWaitingRoom(request *DeleteWaitingRoomRequest) (_result *DeleteWaitingRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWaitingRoomResponse{}
	_body, _err := client.DeleteWaitingRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除等候室事件
//
// @param request - DeleteWaitingRoomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWaitingRoomEventResponse
func (client *Client) DeleteWaitingRoomEventWithOptions(request *DeleteWaitingRoomEventRequest, runtime *util.RuntimeOptions) (_result *DeleteWaitingRoomEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomEventId)) {
		query["WaitingRoomEventId"] = request.WaitingRoomEventId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWaitingRoomEvent"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWaitingRoomEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除等候室事件
//
// @param request - DeleteWaitingRoomEventRequest
//
// @return DeleteWaitingRoomEventResponse
func (client *Client) DeleteWaitingRoomEvent(request *DeleteWaitingRoomEventRequest) (_result *DeleteWaitingRoomEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWaitingRoomEventResponse{}
	_body, _err := client.DeleteWaitingRoomEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除等候室规则
//
// @param request - DeleteWaitingRoomRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWaitingRoomRuleResponse
func (client *Client) DeleteWaitingRoomRuleWithOptions(request *DeleteWaitingRoomRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteWaitingRoomRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomRuleId)) {
		query["WaitingRoomRuleId"] = request.WaitingRoomRuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWaitingRoomRule"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWaitingRoomRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除等候室规则
//
// @param request - DeleteWaitingRoomRuleRequest
//
// @return DeleteWaitingRoomRuleResponse
func (client *Client) DeleteWaitingRoomRule(request *DeleteWaitingRoomRuleRequest) (_result *DeleteWaitingRoomRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWaitingRoomRuleResponse{}
	_body, _err := client.DeleteWaitingRoomRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询定制场景策略配置
//
// @param request - DescribeCustomScenePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomScenePoliciesResponse
func (client *Client) DescribeCustomScenePoliciesWithOptions(request *DescribeCustomScenePoliciesRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomScenePoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomScenePolicies"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomScenePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询定制场景策略配置
//
// @param request - DescribeCustomScenePoliciesRequest
//
// @return DescribeCustomScenePoliciesResponse
func (client *Client) DescribeCustomScenePolicies(request *DescribeCustomScenePoliciesRequest) (_result *DescribeCustomScenePoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomScenePoliciesResponse{}
	_body, _err := client.DescribeCustomScenePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 攻击分析-查询攻击事件列表
//
// @param request - DescribeDDoSAllEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDoSAllEventListResponse
func (client *Client) DescribeDDoSAllEventListWithOptions(request *DescribeDDoSAllEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeDDoSAllEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDDoSAllEventList"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDDoSAllEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 攻击分析-查询攻击事件列表
//
// @param request - DescribeDDoSAllEventListRequest
//
// @return DescribeDDoSAllEventListResponse
func (client *Client) DescribeDDoSAllEventList(request *DescribeDDoSAllEventListRequest) (_result *DescribeDDoSAllEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDoSAllEventListResponse{}
	_body, _err := client.DescribeDDoSAllEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询HTTP DDoS智能防护配置信息
//
// @param request - DescribeHttpDDoSAttackIntelligentProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHttpDDoSAttackIntelligentProtectionResponse
func (client *Client) DescribeHttpDDoSAttackIntelligentProtectionWithOptions(request *DescribeHttpDDoSAttackIntelligentProtectionRequest, runtime *util.RuntimeOptions) (_result *DescribeHttpDDoSAttackIntelligentProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHttpDDoSAttackIntelligentProtection"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHttpDDoSAttackIntelligentProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询HTTP DDoS智能防护配置信息
//
// @param request - DescribeHttpDDoSAttackIntelligentProtectionRequest
//
// @return DescribeHttpDDoSAttackIntelligentProtectionResponse
func (client *Client) DescribeHttpDDoSAttackIntelligentProtection(request *DescribeHttpDDoSAttackIntelligentProtectionRequest) (_result *DescribeHttpDDoSAttackIntelligentProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHttpDDoSAttackIntelligentProtectionResponse{}
	_body, _err := client.DescribeHttpDDoSAttackIntelligentProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询HTTP DDoS攻击防护配置信息
//
// @param request - DescribeHttpDDoSAttackProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHttpDDoSAttackProtectionResponse
func (client *Client) DescribeHttpDDoSAttackProtectionWithOptions(request *DescribeHttpDDoSAttackProtectionRequest, runtime *util.RuntimeOptions) (_result *DescribeHttpDDoSAttackProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHttpDDoSAttackProtection"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHttpDDoSAttackProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询HTTP DDoS攻击防护配置信息
//
// @param request - DescribeHttpDDoSAttackProtectionRequest
//
// @return DescribeHttpDDoSAttackProtectionResponse
func (client *Client) DescribeHttpDDoSAttackProtection(request *DescribeHttpDDoSAttackProtectionRequest) (_result *DescribeHttpDDoSAttackProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHttpDDoSAttackProtectionResponse{}
	_body, _err := client.DescribeHttpDDoSAttackProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询加速服务节点IP段列表
//
// @param request - DescribeIPRangeListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIPRangeListResponse
func (client *Client) DescribeIPRangeListWithOptions(runtime *util.RuntimeOptions) (_result *DescribeIPRangeListResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeIPRangeList"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIPRangeListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询加速服务节点IP段列表
//
// @return DescribeIPRangeListResponse
func (client *Client) DescribeIPRangeList() (_result *DescribeIPRangeListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIPRangeListResponse{}
	_body, _err := client.DescribeIPRangeListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预热任务查询接口
//
// @param request - DescribePreloadTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePreloadTasksResponse
func (client *Client) DescribePreloadTasksWithOptions(request *DescribePreloadTasksRequest, runtime *util.RuntimeOptions) (_result *DescribePreloadTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePreloadTasks"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePreloadTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预热任务查询接口
//
// @param request - DescribePreloadTasksRequest
//
// @return DescribePreloadTasksResponse
func (client *Client) DescribePreloadTasks(request *DescribePreloadTasksRequest) (_result *DescribePreloadTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePreloadTasksResponse{}
	_body, _err := client.DescribePreloadTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 刷新任务查询接口
//
// @param request - DescribePurgeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurgeTasksResponse
func (client *Client) DescribePurgeTasksWithOptions(request *DescribePurgeTasksRequest, runtime *util.RuntimeOptions) (_result *DescribePurgeTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurgeTasks"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePurgeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 刷新任务查询接口
//
// @param request - DescribePurgeTasksRequest
//
// @return DescribePurgeTasksResponse
func (client *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (_result *DescribePurgeTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePurgeTasksResponse{}
	_body, _err := client.DescribePurgeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 禁用定制场景策略
//
// @param request - DisableCustomScenePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCustomScenePolicyResponse
func (client *Client) DisableCustomScenePolicyWithOptions(request *DisableCustomScenePolicyRequest, runtime *util.RuntimeOptions) (_result *DisableCustomScenePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableCustomScenePolicy"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableCustomScenePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用定制场景策略
//
// @param request - DisableCustomScenePolicyRequest
//
// @return DisableCustomScenePolicyResponse
func (client *Client) DisableCustomScenePolicy(request *DisableCustomScenePolicyRequest) (_result *DisableCustomScenePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableCustomScenePolicyResponse{}
	_body, _err := client.DisableCustomScenePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑站点WAF配置
//
// @param tmpReq - EditSiteWafSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditSiteWafSettingsResponse
func (client *Client) EditSiteWafSettingsWithOptions(tmpReq *EditSiteWafSettingsRequest, runtime *util.RuntimeOptions) (_result *EditSiteWafSettingsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &EditSiteWafSettingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Settings)) {
		request.SettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Settings, tea.String("Settings"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SettingsShrink)) {
		body["Settings"] = request.SettingsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EditSiteWafSettings"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditSiteWafSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑站点WAF配置
//
// @param request - EditSiteWafSettingsRequest
//
// @return EditSiteWafSettingsResponse
func (client *Client) EditSiteWafSettings(request *EditSiteWafSettingsRequest) (_result *EditSiteWafSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditSiteWafSettingsResponse{}
	_body, _err := client.EditSiteWafSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动定制场景策略
//
// @param request - EnableCustomScenePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCustomScenePolicyResponse
func (client *Client) EnableCustomScenePolicyWithOptions(request *EnableCustomScenePolicyRequest, runtime *util.RuntimeOptions) (_result *EnableCustomScenePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableCustomScenePolicy"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableCustomScenePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动定制场景策略
//
// @param request - EnableCustomScenePolicyRequest
//
// @return EnableCustomScenePolicyResponse
func (client *Client) EnableCustomScenePolicy(request *EnableCustomScenePolicyRequest) (_result *EnableCustomScenePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableCustomScenePolicyResponse{}
	_body, _err := client.EnableCustomScenePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出记录
//
// @param request - ExportRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportRecordsResponse
func (client *Client) ExportRecordsWithOptions(request *ExportRecordsRequest, runtime *util.RuntimeOptions) (_result *ExportRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportRecords"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出记录
//
// @param request - ExportRecordsRequest
//
// @return ExportRecordsResponse
func (client *Client) ExportRecords(request *ExportRecordsRequest) (_result *ExportRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportRecordsResponse{}
	_body, _err := client.ExportRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Namespace信息
//
// @param request - GetKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKvNamespaceResponse
func (client *Client) GetKvNamespaceWithOptions(request *GetKvNamespaceRequest, runtime *util.RuntimeOptions) (_result *GetKvNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetKvNamespace"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetKvNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Namespace信息
//
// @param request - GetKvNamespaceRequest
//
// @return GetKvNamespaceResponse
func (client *Client) GetKvNamespace(request *GetKvNamespaceRequest) (_result *GetKvNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKvNamespaceResponse{}
	_body, _err := client.GetKvNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个自定义列表
//
// @param request - GetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetListResponse
func (client *Client) GetListWithOptions(request *GetListRequest, runtime *util.RuntimeOptions) (_result *GetListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetList"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个自定义列表
//
// @param request - GetListRequest
//
// @return GetListResponse
func (client *Client) GetList(request *GetListRequest) (_result *GetListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetListResponse{}
	_body, _err := client.GetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个自定义响应页面详情
//
// @param request - GetPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPageResponse
func (client *Client) GetPageWithOptions(request *GetPageRequest, runtime *util.RuntimeOptions) (_result *GetPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPage"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个自定义响应页面详情
//
// @param request - GetPageRequest
//
// @return GetPageResponse
func (client *Client) GetPage(request *GetPageRequest) (_result *GetPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPageResponse{}
	_body, _err := client.GetPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取刷新Quota
//
// @param request - GetPurgeQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPurgeQuotaResponse
func (client *Client) GetPurgeQuotaWithOptions(request *GetPurgeQuotaRequest, runtime *util.RuntimeOptions) (_result *GetPurgeQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPurgeQuota"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPurgeQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取刷新Quota
//
// @param request - GetPurgeQuotaRequest
//
// @return GetPurgeQuotaResponse
func (client *Client) GetPurgeQuota(request *GetPurgeQuotaRequest) (_result *GetPurgeQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPurgeQuotaResponse{}
	_body, _err := client.GetPurgeQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ub日志字段列表接口
//
// @param request - GetRealtimeDeliveryFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealtimeDeliveryFieldResponse
func (client *Client) GetRealtimeDeliveryFieldWithOptions(request *GetRealtimeDeliveryFieldRequest, runtime *util.RuntimeOptions) (_result *GetRealtimeDeliveryFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRealtimeDeliveryField"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRealtimeDeliveryFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ub日志字段列表接口
//
// @param request - GetRealtimeDeliveryFieldRequest
//
// @return GetRealtimeDeliveryFieldResponse
func (client *Client) GetRealtimeDeliveryField(request *GetRealtimeDeliveryFieldRequest) (_result *GetRealtimeDeliveryFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRealtimeDeliveryFieldResponse{}
	_body, _err := client.GetRealtimeDeliveryFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个记录信息
//
// @param request - GetRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecordResponse
func (client *Client) GetRecordWithOptions(request *GetRecordRequest, runtime *util.RuntimeOptions) (_result *GetRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecord"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个记录信息
//
// @param request - GetRecordRequest
//
// @return GetRecordResponse
func (client *Client) GetRecord(request *GetRecordRequest) (_result *GetRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecordResponse{}
	_body, _err := client.GetRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个定时预热任务
//
// @param request - GetScheduledPreloadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduledPreloadJobResponse
func (client *Client) GetScheduledPreloadJobWithOptions(request *GetScheduledPreloadJobRequest, runtime *util.RuntimeOptions) (_result *GetScheduledPreloadJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScheduledPreloadJob"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScheduledPreloadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个定时预热任务
//
// @param request - GetScheduledPreloadJobRequest
//
// @return GetScheduledPreloadJobResponse
func (client *Client) GetScheduledPreloadJob(request *GetScheduledPreloadJobRequest) (_result *GetScheduledPreloadJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScheduledPreloadJobResponse{}
	_body, _err := client.GetScheduledPreloadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个站点信息
//
// @param request - GetSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteResponse
func (client *Client) GetSiteWithOptions(request *GetSiteRequest, runtime *util.RuntimeOptions) (_result *GetSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSite"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个站点信息
//
// @param request - GetSiteRequest
//
// @return GetSiteResponse
func (client *Client) GetSite(request *GetSiteRequest) (_result *GetSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSiteResponse{}
	_body, _err := client.GetSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询当前NS列表
//
// @param request - GetSiteCurrentNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteCurrentNSResponse
func (client *Client) GetSiteCurrentNSWithOptions(request *GetSiteCurrentNSRequest, runtime *util.RuntimeOptions) (_result *GetSiteCurrentNSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSiteCurrentNS"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSiteCurrentNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前NS列表
//
// @param request - GetSiteCurrentNSRequest
//
// @return GetSiteCurrentNSResponse
func (client *Client) GetSiteCurrentNS(request *GetSiteCurrentNSRequest) (_result *GetSiteCurrentNSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSiteCurrentNSResponse{}
	_body, _err := client.GetSiteCurrentNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取自定义字段
//
// @param request - GetSiteCustomLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteCustomLogResponse
func (client *Client) GetSiteCustomLogWithOptions(request *GetSiteCustomLogRequest, runtime *util.RuntimeOptions) (_result *GetSiteCustomLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSiteCustomLog"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSiteCustomLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义字段
//
// @param request - GetSiteCustomLogRequest
//
// @return GetSiteCustomLogResponse
func (client *Client) GetSiteCustomLog(request *GetSiteCustomLogRequest) (_result *GetSiteCustomLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSiteCustomLogResponse{}
	_body, _err := client.GetSiteCustomLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一个实时日志任务投递
//
// @param request - GetSiteDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteDeliveryTaskResponse
func (client *Client) GetSiteDeliveryTaskWithOptions(request *GetSiteDeliveryTaskRequest, runtime *util.RuntimeOptions) (_result *GetSiteDeliveryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSiteDeliveryTask"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSiteDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一个实时日志任务投递
//
// @param request - GetSiteDeliveryTaskRequest
//
// @return GetSiteDeliveryTaskResponse
func (client *Client) GetSiteDeliveryTask(request *GetSiteDeliveryTaskRequest) (_result *GetSiteDeliveryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSiteDeliveryTaskResponse{}
	_body, _err := client.GetSiteDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取日志投递任务quota数
//
// @param request - GetSiteLogDeliveryQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteLogDeliveryQuotaResponse
func (client *Client) GetSiteLogDeliveryQuotaWithOptions(request *GetSiteLogDeliveryQuotaRequest, runtime *util.RuntimeOptions) (_result *GetSiteLogDeliveryQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSiteLogDeliveryQuota"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSiteLogDeliveryQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志投递任务quota数
//
// @param request - GetSiteLogDeliveryQuotaRequest
//
// @return GetSiteLogDeliveryQuotaResponse
func (client *Client) GetSiteLogDeliveryQuota(request *GetSiteLogDeliveryQuotaRequest) (_result *GetSiteLogDeliveryQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSiteLogDeliveryQuotaResponse{}
	_body, _err := client.GetSiteLogDeliveryQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取站点WAF配置
//
// @param request - GetSiteWafSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSiteWafSettingsResponse
func (client *Client) GetSiteWafSettingsWithOptions(request *GetSiteWafSettingsRequest, runtime *util.RuntimeOptions) (_result *GetSiteWafSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSiteWafSettings"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSiteWafSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取站点WAF配置
//
// @param request - GetSiteWafSettingsRequest
//
// @return GetSiteWafSettingsResponse
func (client *Client) GetSiteWafSettings(request *GetSiteWafSettingsRequest) (_result *GetSiteWafSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSiteWafSettingsResponse{}
	_body, _err := client.GetSiteWafSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文件上传任务查询接口
//
// @param request - GetUploadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadTaskResponse
func (client *Client) GetUploadTaskWithOptions(request *GetUploadTaskRequest, runtime *util.RuntimeOptions) (_result *GetUploadTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUploadTask"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUploadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文件上传任务查询接口
//
// @param request - GetUploadTaskRequest
//
// @return GetUploadTaskResponse
func (client *Client) GetUploadTask(request *GetUploadTaskRequest) (_result *GetUploadTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUploadTaskResponse{}
	_body, _err := client.GetUploadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一个用户粒度任务投递
//
// @param request - GetUserDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserDeliveryTaskResponse
func (client *Client) GetUserDeliveryTaskWithOptions(request *GetUserDeliveryTaskRequest, runtime *util.RuntimeOptions) (_result *GetUserDeliveryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserDeliveryTask"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一个用户粒度任务投递
//
// @param request - GetUserDeliveryTaskRequest
//
// @return GetUserDeliveryTaskResponse
func (client *Client) GetUserDeliveryTask(request *GetUserDeliveryTaskRequest) (_result *GetUserDeliveryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserDeliveryTaskResponse{}
	_body, _err := client.GetUserDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取日志投递任务用户quota数
//
// @param request - GetUserLogDeliveryQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserLogDeliveryQuotaResponse
func (client *Client) GetUserLogDeliveryQuotaWithOptions(request *GetUserLogDeliveryQuotaRequest, runtime *util.RuntimeOptions) (_result *GetUserLogDeliveryQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserLogDeliveryQuota"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserLogDeliveryQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志投递任务用户quota数
//
// @param request - GetUserLogDeliveryQuotaRequest
//
// @return GetUserLogDeliveryQuotaResponse
func (client *Client) GetUserLogDeliveryQuota(request *GetUserLogDeliveryQuotaRequest) (_result *GetUserLogDeliveryQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserLogDeliveryQuotaResponse{}
	_body, _err := client.GetUserLogDeliveryQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetWafBotAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWafBotAppKeyResponse
func (client *Client) GetWafBotAppKeyWithOptions(runtime *util.RuntimeOptions) (_result *GetWafBotAppKeyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetWafBotAppKey"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWafBotAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return GetWafBotAppKeyResponse
func (client *Client) GetWafBotAppKey() (_result *GetWafBotAppKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWafBotAppKeyResponse{}
	_body, _err := client.GetWafBotAppKeyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将匹配项转换为表达式
//
// @param request - GetWafFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWafFilterResponse
func (client *Client) GetWafFilterWithOptions(request *GetWafFilterRequest, runtime *util.RuntimeOptions) (_result *GetWafFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		query["Phase"] = request.Phase
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWafFilter"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWafFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将匹配项转换为表达式
//
// @param request - GetWafFilterRequest
//
// @return GetWafFilterResponse
func (client *Client) GetWafFilter(request *GetWafFilterRequest) (_result *GetWafFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWafFilterResponse{}
	_body, _err := client.GetWafFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取WAF配额详情
//
// @param request - GetWafQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWafQuotaResponse
func (client *Client) GetWafQuotaWithOptions(request *GetWafQuotaRequest, runtime *util.RuntimeOptions) (_result *GetWafQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Paths)) {
		query["Paths"] = request.Paths
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWafQuota"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWafQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取WAF配额详情
//
// @param request - GetWafQuotaRequest
//
// @return GetWafQuotaResponse
func (client *Client) GetWafQuota(request *GetWafQuotaRequest) (_result *GetWafQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWafQuotaResponse{}
	_body, _err := client.GetWafQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个WAF规则详情
//
// @param request - GetWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWafRuleResponse
func (client *Client) GetWafRuleWithOptions(request *GetWafRuleRequest, runtime *util.RuntimeOptions) (_result *GetWafRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWafRule"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWafRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个WAF规则详情
//
// @param request - GetWafRuleRequest
//
// @return GetWafRuleResponse
func (client *Client) GetWafRule(request *GetWafRuleRequest) (_result *GetWafRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWafRuleResponse{}
	_body, _err := client.GetWafRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取WAF规则集详情
//
// @param request - GetWafRulesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWafRulesetResponse
func (client *Client) GetWafRulesetWithOptions(request *GetWafRulesetRequest, runtime *util.RuntimeOptions) (_result *GetWafRulesetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		query["Phase"] = request.Phase
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWafRuleset"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWafRulesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取WAF规则集详情
//
// @param request - GetWafRulesetRequest
//
// @return GetWafRulesetResponse
func (client *Client) GetWafRuleset(request *GetWafRulesetRequest) (_result *GetWafRulesetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWafRulesetResponse{}
	_body, _err := client.GetWafRulesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一个边缘容器应用的全部域名记录
//
// @param request - ListEdgeContainerAppRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeContainerAppRecordsResponse
func (client *Client) ListEdgeContainerAppRecordsWithOptions(request *ListEdgeContainerAppRecordsRequest, runtime *util.RuntimeOptions) (_result *ListEdgeContainerAppRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEdgeContainerAppRecords"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEdgeContainerAppRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一个边缘容器应用的全部域名记录
//
// @param request - ListEdgeContainerAppRecordsRequest
//
// @return ListEdgeContainerAppRecordsResponse
func (client *Client) ListEdgeContainerAppRecords(request *ListEdgeContainerAppRecordsRequest) (_result *ListEdgeContainerAppRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEdgeContainerAppRecordsResponse{}
	_body, _err := client.ListEdgeContainerAppRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询站点的边缘容器记录
//
// @param request - ListEdgeContainerRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeContainerRecordsResponse
func (client *Client) ListEdgeContainerRecordsWithOptions(request *ListEdgeContainerRecordsRequest, runtime *util.RuntimeOptions) (_result *ListEdgeContainerRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEdgeContainerRecords"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEdgeContainerRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询站点的边缘容器记录
//
// @param request - ListEdgeContainerRecordsRequest
//
// @return ListEdgeContainerRecordsResponse
func (client *Client) ListEdgeContainerRecords(request *ListEdgeContainerRecordsRequest) (_result *ListEdgeContainerRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEdgeContainerRecordsResponse{}
	_body, _err := client.ListEdgeContainerRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询站点的边缘路由记录
//
// @param request - ListEdgeRoutineRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeRoutineRecordsResponse
func (client *Client) ListEdgeRoutineRecordsWithOptions(request *ListEdgeRoutineRecordsRequest, runtime *util.RuntimeOptions) (_result *ListEdgeRoutineRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEdgeRoutineRecords"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEdgeRoutineRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询站点的边缘路由记录
//
// @param request - ListEdgeRoutineRecordsRequest
//
// @return ListEdgeRoutineRecordsResponse
func (client *Client) ListEdgeRoutineRecords(request *ListEdgeRoutineRecordsRequest) (_result *ListEdgeRoutineRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEdgeRoutineRecordsResponse{}
	_body, _err := client.ListEdgeRoutineRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 遍历Namespace的Key值
//
// @param request - ListKvsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKvsResponse
func (client *Client) ListKvsWithOptions(request *ListKvsRequest, runtime *util.RuntimeOptions) (_result *ListKvsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKvs"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListKvsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 遍历Namespace的Key值
//
// @param request - ListKvsRequest
//
// @return ListKvsResponse
func (client *Client) ListKvs(request *ListKvsRequest) (_result *ListKvsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListKvsResponse{}
	_body, _err := client.ListKvsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举自定义列表
//
// @param tmpReq - ListListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListListsResponse
func (client *Client) ListListsWithOptions(tmpReq *ListListsRequest, runtime *util.RuntimeOptions) (_result *ListListsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListListsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.QueryArgs)) {
		request.QueryArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryArgs, tea.String("QueryArgs"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryArgsShrink)) {
		query["QueryArgs"] = request.QueryArgsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLists"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举自定义列表
//
// @param request - ListListsRequest
//
// @return ListListsResponse
func (client *Client) ListLists(request *ListListsRequest) (_result *ListListsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListsResponse{}
	_body, _err := client.ListListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询负载均衡区域列表
//
// @param request - ListLoadBalancerRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLoadBalancerRegionsResponse
func (client *Client) ListLoadBalancerRegionsWithOptions(request *ListLoadBalancerRegionsRequest, runtime *util.RuntimeOptions) (_result *ListLoadBalancerRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLoadBalancerRegions"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLoadBalancerRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询负载均衡区域列表
//
// @param request - ListLoadBalancerRegionsRequest
//
// @return ListLoadBalancerRegionsResponse
func (client *Client) ListLoadBalancerRegions(request *ListLoadBalancerRegionsRequest) (_result *ListLoadBalancerRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLoadBalancerRegionsResponse{}
	_body, _err := client.ListLoadBalancerRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举自定义托管规则组
//
// @param request - ListManagedRulesGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManagedRulesGroupsResponse
func (client *Client) ListManagedRulesGroupsWithOptions(request *ListManagedRulesGroupsRequest, runtime *util.RuntimeOptions) (_result *ListManagedRulesGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListManagedRulesGroups"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListManagedRulesGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举自定义托管规则组
//
// @param request - ListManagedRulesGroupsRequest
//
// @return ListManagedRulesGroupsResponse
func (client *Client) ListManagedRulesGroups(request *ListManagedRulesGroupsRequest) (_result *ListManagedRulesGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListManagedRulesGroupsResponse{}
	_body, _err := client.ListManagedRulesGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举自定义响应页面
//
// @param request - ListPagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPagesResponse
func (client *Client) ListPagesWithOptions(request *ListPagesRequest, runtime *util.RuntimeOptions) (_result *ListPagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPages"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举自定义响应页面
//
// @param request - ListPagesRequest
//
// @return ListPagesResponse
func (client *Client) ListPages(request *ListPagesRequest) (_result *ListPagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPagesResponse{}
	_body, _err := client.ListPagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询站点下的记录列表
//
// @param request - ListRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecordsResponse
func (client *Client) ListRecordsWithOptions(request *ListRecordsRequest, runtime *util.RuntimeOptions) (_result *ListRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecords"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询站点下的记录列表
//
// @param request - ListRecordsRequest
//
// @return ListRecordsResponse
func (client *Client) ListRecords(request *ListRecordsRequest) (_result *ListRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecordsResponse{}
	_body, _err := client.ListRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出指定任务下的执行计划
//
// @param request - ListScheduledPreloadExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledPreloadExecutionsResponse
func (client *Client) ListScheduledPreloadExecutionsWithOptions(request *ListScheduledPreloadExecutionsRequest, runtime *util.RuntimeOptions) (_result *ListScheduledPreloadExecutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScheduledPreloadExecutions"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListScheduledPreloadExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出指定任务下的执行计划
//
// @param request - ListScheduledPreloadExecutionsRequest
//
// @return ListScheduledPreloadExecutionsResponse
func (client *Client) ListScheduledPreloadExecutions(request *ListScheduledPreloadExecutionsRequest) (_result *ListScheduledPreloadExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScheduledPreloadExecutionsResponse{}
	_body, _err := client.ListScheduledPreloadExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出定时预热任务列表
//
// @param request - ListScheduledPreloadJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledPreloadJobsResponse
func (client *Client) ListScheduledPreloadJobsWithOptions(request *ListScheduledPreloadJobsRequest, runtime *util.RuntimeOptions) (_result *ListScheduledPreloadJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScheduledPreloadJobs"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListScheduledPreloadJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出定时预热任务列表
//
// @param request - ListScheduledPreloadJobsRequest
//
// @return ListScheduledPreloadJobsResponse
func (client *Client) ListScheduledPreloadJobs(request *ListScheduledPreloadJobsRequest) (_result *ListScheduledPreloadJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScheduledPreloadJobsResponse{}
	_body, _err := client.ListScheduledPreloadJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出全部任务投递
//
// @param request - ListSiteDeliveryTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSiteDeliveryTasksResponse
func (client *Client) ListSiteDeliveryTasksWithOptions(request *ListSiteDeliveryTasksRequest, runtime *util.RuntimeOptions) (_result *ListSiteDeliveryTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSiteDeliveryTasks"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSiteDeliveryTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出全部任务投递
//
// @param request - ListSiteDeliveryTasksRequest
//
// @return ListSiteDeliveryTasksResponse
func (client *Client) ListSiteDeliveryTasks(request *ListSiteDeliveryTasksRequest) (_result *ListSiteDeliveryTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSiteDeliveryTasksResponse{}
	_body, _err := client.ListSiteDeliveryTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询站点列表
//
// @param tmpReq - ListSitesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSitesResponse
func (client *Client) ListSitesWithOptions(tmpReq *ListSitesRequest, runtime *util.RuntimeOptions) (_result *ListSitesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListSitesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TagFilter)) {
		request.TagFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagFilter, tea.String("TagFilter"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSites"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSitesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询站点列表
//
// @param request - ListSitesRequest
//
// @return ListSitesResponse
func (client *Client) ListSites(request *ListSitesRequest) (_result *ListSitesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSitesResponse{}
	_body, _err := client.ListSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询云资源已经绑定的标签列表
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["MaxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云资源已经绑定的标签列表
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件上传任务
//
// @param request - ListUploadTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUploadTasksResponse
func (client *Client) ListUploadTasksWithOptions(request *ListUploadTasksRequest, runtime *util.RuntimeOptions) (_result *ListUploadTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUploadTasks"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUploadTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件上传任务
//
// @param request - ListUploadTasksRequest
//
// @return ListUploadTasksResponse
func (client *Client) ListUploadTasks(request *ListUploadTasksRequest) (_result *ListUploadTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUploadTasksResponse{}
	_body, _err := client.ListUploadTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出用户全部任务投递
//
// @param request - ListUserDeliveryTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserDeliveryTasksResponse
func (client *Client) ListUserDeliveryTasksWithOptions(request *ListUserDeliveryTasksRequest, runtime *util.RuntimeOptions) (_result *ListUserDeliveryTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserDeliveryTasks"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserDeliveryTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出用户全部任务投递
//
// @param request - ListUserDeliveryTasksRequest
//
// @return ListUserDeliveryTasksResponse
func (client *Client) ListUserDeliveryTasks(request *ListUserDeliveryTasksRequest) (_result *ListUserDeliveryTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserDeliveryTasksResponse{}
	_body, _err := client.ListUserDeliveryTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举WAF阶段
//
// @param request - ListWafPhasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafPhasesResponse
func (client *Client) ListWafPhasesWithOptions(request *ListWafPhasesRequest, runtime *util.RuntimeOptions) (_result *ListWafPhasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWafPhases"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWafPhasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举WAF阶段
//
// @param request - ListWafPhasesRequest
//
// @return ListWafPhasesResponse
func (client *Client) ListWafPhases(request *ListWafPhasesRequest) (_result *ListWafPhasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWafPhasesResponse{}
	_body, _err := client.ListWafPhasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举WAF规则
//
// @param tmpReq - ListWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafRulesResponse
func (client *Client) ListWafRulesWithOptions(tmpReq *ListWafRulesRequest, runtime *util.RuntimeOptions) (_result *ListWafRulesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListWafRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.QueryArgs)) {
		request.QueryArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryArgs, tea.String("QueryArgs"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		query["Phase"] = request.Phase
	}

	if !tea.BoolValue(util.IsUnset(request.QueryArgsShrink)) {
		query["QueryArgs"] = request.QueryArgsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWafRules"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWafRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举WAF规则
//
// @param request - ListWafRulesRequest
//
// @return ListWafRulesResponse
func (client *Client) ListWafRules(request *ListWafRulesRequest) (_result *ListWafRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWafRulesResponse{}
	_body, _err := client.ListWafRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举WAF规则集
//
// @param tmpReq - ListWafRulesetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafRulesetsResponse
func (client *Client) ListWafRulesetsWithOptions(tmpReq *ListWafRulesetsRequest, runtime *util.RuntimeOptions) (_result *ListWafRulesetsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListWafRulesetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.QueryArgs)) {
		request.QueryArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryArgs, tea.String("QueryArgs"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		query["Phase"] = request.Phase
	}

	if !tea.BoolValue(util.IsUnset(request.QueryArgsShrink)) {
		query["QueryArgs"] = request.QueryArgsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWafRulesets"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWafRulesetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举WAF规则集
//
// @param request - ListWafRulesetsRequest
//
// @return ListWafRulesetsResponse
func (client *Client) ListWafRulesets(request *ListWafRulesetsRequest) (_result *ListWafRulesetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWafRulesetsResponse{}
	_body, _err := client.ListWafRulesetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举WAF模板规则
//
// @param tmpReq - ListWafTemplateRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafTemplateRulesResponse
func (client *Client) ListWafTemplateRulesWithOptions(tmpReq *ListWafTemplateRulesRequest, runtime *util.RuntimeOptions) (_result *ListWafTemplateRulesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListWafTemplateRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.QueryArgs)) {
		request.QueryArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryArgs, tea.String("QueryArgs"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		query["Phase"] = request.Phase
	}

	if !tea.BoolValue(util.IsUnset(request.QueryArgsShrink)) {
		query["QueryArgs"] = request.QueryArgsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWafTemplateRules"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWafTemplateRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举WAF模板规则
//
// @param request - ListWafTemplateRulesRequest
//
// @return ListWafTemplateRulesResponse
func (client *Client) ListWafTemplateRules(request *ListWafTemplateRulesRequest) (_result *ListWafTemplateRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWafTemplateRulesResponse{}
	_body, _err := client.ListWafTemplateRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举WAF规则使用情况
//
// @param request - ListWafUsageOfRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWafUsageOfRulesResponse
func (client *Client) ListWafUsageOfRulesWithOptions(request *ListWafUsageOfRulesRequest, runtime *util.RuntimeOptions) (_result *ListWafUsageOfRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		query["Phase"] = request.Phase
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWafUsageOfRules"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWafUsageOfRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举WAF规则使用情况
//
// @param request - ListWafUsageOfRulesRequest
//
// @return ListWafUsageOfRulesResponse
func (client *Client) ListWafUsageOfRules(request *ListWafUsageOfRulesRequest) (_result *ListWafUsageOfRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWafUsageOfRulesResponse{}
	_body, _err := client.ListWafUsageOfRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询等候室事件
//
// @param request - ListWaitingRoomEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWaitingRoomEventsResponse
func (client *Client) ListWaitingRoomEventsWithOptions(request *ListWaitingRoomEventsRequest, runtime *util.RuntimeOptions) (_result *ListWaitingRoomEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWaitingRoomEvents"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWaitingRoomEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询等候室事件
//
// @param request - ListWaitingRoomEventsRequest
//
// @return ListWaitingRoomEventsResponse
func (client *Client) ListWaitingRoomEvents(request *ListWaitingRoomEventsRequest) (_result *ListWaitingRoomEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWaitingRoomEventsResponse{}
	_body, _err := client.ListWaitingRoomEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询等候室绕过规则
//
// @param request - ListWaitingRoomRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWaitingRoomRulesResponse
func (client *Client) ListWaitingRoomRulesWithOptions(request *ListWaitingRoomRulesRequest, runtime *util.RuntimeOptions) (_result *ListWaitingRoomRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWaitingRoomRules"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWaitingRoomRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询等候室绕过规则
//
// @param request - ListWaitingRoomRulesRequest
//
// @return ListWaitingRoomRulesResponse
func (client *Client) ListWaitingRoomRules(request *ListWaitingRoomRulesRequest) (_result *ListWaitingRoomRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWaitingRoomRulesResponse{}
	_body, _err := client.ListWaitingRoomRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询等候室
//
// @param request - ListWaitingRoomsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWaitingRoomsResponse
func (client *Client) ListWaitingRoomsWithOptions(request *ListWaitingRoomsRequest, runtime *util.RuntimeOptions) (_result *ListWaitingRoomsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWaitingRooms"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWaitingRoomsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询等候室
//
// @param request - ListWaitingRoomsRequest
//
// @return ListWaitingRoomsResponse
func (client *Client) ListWaitingRooms(request *ListWaitingRoomsRequest) (_result *ListWaitingRoomsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWaitingRoomsResponse{}
	_body, _err := client.ListWaitingRoomsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 缓存预热
//
// @param tmpReq - PreloadCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreloadCachesResponse
func (client *Client) PreloadCachesWithOptions(tmpReq *PreloadCachesRequest, runtime *util.RuntimeOptions) (_result *PreloadCachesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PreloadCachesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Content)) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, tea.String("Content"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Headers)) {
		request.HeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Headers, tea.String("Headers"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentShrink)) {
		query["Content"] = request.ContentShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HeadersShrink)) {
		query["Headers"] = request.HeadersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PreloadCaches"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PreloadCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 缓存预热
//
// @param request - PreloadCachesRequest
//
// @return PreloadCachesResponse
func (client *Client) PreloadCaches(request *PreloadCachesRequest) (_result *PreloadCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PreloadCachesResponse{}
	_body, _err := client.PreloadCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 缓存刷新
//
// @param tmpReq - PurgeCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurgeCachesResponse
func (client *Client) PurgeCachesWithOptions(tmpReq *PurgeCachesRequest, runtime *util.RuntimeOptions) (_result *PurgeCachesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PurgeCachesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Content)) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, tea.String("Content"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentShrink)) {
		query["Content"] = request.ContentShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EdgeComputePurge)) {
		query["EdgeComputePurge"] = request.EdgeComputePurge
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PurgeCaches"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PurgeCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 缓存刷新
//
// @param request - PurgeCachesRequest
//
// @return PurgeCachesResponse
func (client *Client) PurgeCaches(request *PurgeCachesRequest) (_result *PurgeCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PurgeCachesResponse{}
	_body, _err := client.PurgeCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置Namespace的Key-Value对
//
// @param request - PutKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutKvResponse
func (client *Client) PutKvWithOptions(request *PutKvRequest, runtime *util.RuntimeOptions) (_result *PutKvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Base64)) {
		query["Base64"] = request.Base64
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		query["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.ExpirationTtl)) {
		query["ExpirationTtl"] = request.ExpirationTtl
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutKv"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置Namespace的Key-Value对
//
// @param request - PutKvRequest
//
// @return PutKvResponse
func (client *Client) PutKv(request *PutKvRequest) (_result *PutKvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutKvResponse{}
	_body, _err := client.PutKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置账户
//
// @param request - PutKvAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutKvAccountResponse
func (client *Client) PutKvAccountWithOptions(request *PutKvAccountRequest, runtime *util.RuntimeOptions) (_result *PutKvAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		body["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutKvAccount"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutKvAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置账户
//
// @param request - PutKvAccountRequest
//
// @return PutKvAccountResponse
func (client *Client) PutKvAccount(request *PutKvAccountRequest) (_result *PutKvAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutKvAccountResponse{}
	_body, _err := client.PutKvAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置Namespace的Key-Value对，支持最大25M的Body
//
// @param request - PutKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutKvWithHighCapacityResponse
func (client *Client) PutKvWithHighCapacityWithOptions(request *PutKvWithHighCapacityRequest, runtime *util.RuntimeOptions) (_result *PutKvWithHighCapacityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutKvWithHighCapacity"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutKvWithHighCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置Namespace的Key-Value对，支持最大25M的Body
//
// @param request - PutKvWithHighCapacityRequest
//
// @return PutKvWithHighCapacityResponse
func (client *Client) PutKvWithHighCapacity(request *PutKvWithHighCapacityRequest) (_result *PutKvWithHighCapacityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutKvWithHighCapacityResponse{}
	_body, _err := client.PutKvWithHighCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutKvWithHighCapacityAdvance(request *PutKvWithHighCapacityAdvanceRequest, runtime *util.RuntimeOptions) (_result *PutKvWithHighCapacityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ESA"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	putKvWithHighCapacityReq := &PutKvWithHighCapacityRequest{}
	openapiutil.Convert(request, putKvWithHighCapacityReq)
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		putKvWithHighCapacityReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	putKvWithHighCapacityResp, _err := client.PutKvWithHighCapacityWithOptions(putKvWithHighCapacityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = putKvWithHighCapacityResp
	return _result, _err
}

// Summary:
//
// 重置定时预热任务的进度，从头开始预热
//
// @param request - ResetScheduledPreloadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetScheduledPreloadJobResponse
func (client *Client) ResetScheduledPreloadJobWithOptions(request *ResetScheduledPreloadJobRequest, runtime *util.RuntimeOptions) (_result *ResetScheduledPreloadJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetScheduledPreloadJob"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetScheduledPreloadJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置定时预热任务的进度，从头开始预热
//
// @param request - ResetScheduledPreloadJobRequest
//
// @return ResetScheduledPreloadJobResponse
func (client *Client) ResetScheduledPreloadJob(request *ResetScheduledPreloadJobRequest) (_result *ResetScheduledPreloadJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetScheduledPreloadJobResponse{}
	_body, _err := client.ResetScheduledPreloadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置证书
//
// @param request - SetCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCertificateResponse
func (client *Client) SetCertificateWithOptions(request *SetCertificateRequest, runtime *util.RuntimeOptions) (_result *SetCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CasId)) {
		body["CasId"] = request.CasId
	}

	if !tea.BoolValue(util.IsUnset(request.Certificate)) {
		body["Certificate"] = request.Certificate
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		body["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		body["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Update)) {
		body["Update"] = request.Update
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetCertificate"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置证书
//
// @param request - SetCertificateRequest
//
// @return SetCertificateResponse
func (client *Client) SetCertificate(request *SetCertificateRequest) (_result *SetCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCertificateResponse{}
	_body, _err := client.SetCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置HTTP DDoS智能防护配置信息
//
// @param request - SetHttpDDoSAttackIntelligentProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetHttpDDoSAttackIntelligentProtectionResponse
func (client *Client) SetHttpDDoSAttackIntelligentProtectionWithOptions(request *SetHttpDDoSAttackIntelligentProtectionRequest, runtime *util.RuntimeOptions) (_result *SetHttpDDoSAttackIntelligentProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AiMode)) {
		query["AiMode"] = request.AiMode
	}

	if !tea.BoolValue(util.IsUnset(request.AiTemplate)) {
		query["AiTemplate"] = request.AiTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetHttpDDoSAttackIntelligentProtection"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetHttpDDoSAttackIntelligentProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置HTTP DDoS智能防护配置信息
//
// @param request - SetHttpDDoSAttackIntelligentProtectionRequest
//
// @return SetHttpDDoSAttackIntelligentProtectionResponse
func (client *Client) SetHttpDDoSAttackIntelligentProtection(request *SetHttpDDoSAttackIntelligentProtectionRequest) (_result *SetHttpDDoSAttackIntelligentProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetHttpDDoSAttackIntelligentProtectionResponse{}
	_body, _err := client.SetHttpDDoSAttackIntelligentProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置HTTP DDoS攻击防护配置信息
//
// @param request - SetHttpDDoSAttackProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetHttpDDoSAttackProtectionResponse
func (client *Client) SetHttpDDoSAttackProtectionWithOptions(request *SetHttpDDoSAttackProtectionRequest, runtime *util.RuntimeOptions) (_result *SetHttpDDoSAttackProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GlobalMode)) {
		query["GlobalMode"] = request.GlobalMode
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetHttpDDoSAttackProtection"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetHttpDDoSAttackProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置HTTP DDoS攻击防护配置信息
//
// @param request - SetHttpDDoSAttackProtectionRequest
//
// @return SetHttpDDoSAttackProtectionResponse
func (client *Client) SetHttpDDoSAttackProtection(request *SetHttpDDoSAttackProtectionRequest) (_result *SetHttpDDoSAttackProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetHttpDDoSAttackProtectionResponse{}
	_body, _err := client.SetHttpDDoSAttackProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开始单个定时预热计划
//
// @param request - StartScheduledPreloadExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartScheduledPreloadExecutionResponse
func (client *Client) StartScheduledPreloadExecutionWithOptions(request *StartScheduledPreloadExecutionRequest, runtime *util.RuntimeOptions) (_result *StartScheduledPreloadExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartScheduledPreloadExecution"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartScheduledPreloadExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开始单个定时预热计划
//
// @param request - StartScheduledPreloadExecutionRequest
//
// @return StartScheduledPreloadExecutionResponse
func (client *Client) StartScheduledPreloadExecution(request *StartScheduledPreloadExecutionRequest) (_result *StartScheduledPreloadExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartScheduledPreloadExecutionResponse{}
	_body, _err := client.StartScheduledPreloadExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止单个定时预热计划
//
// @param request - StopScheduledPreloadExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopScheduledPreloadExecutionResponse
func (client *Client) StopScheduledPreloadExecutionWithOptions(request *StopScheduledPreloadExecutionRequest, runtime *util.RuntimeOptions) (_result *StopScheduledPreloadExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopScheduledPreloadExecution"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopScheduledPreloadExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止单个定时预热计划
//
// @param request - StopScheduledPreloadExecutionRequest
//
// @return StopScheduledPreloadExecutionResponse
func (client *Client) StopScheduledPreloadExecution(request *StopScheduledPreloadExecutionRequest) (_result *StopScheduledPreloadExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopScheduledPreloadExecutionResponse{}
	_body, _err := client.StopScheduledPreloadExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将表达式转换为匹配项
//
// @param request - TransformExpressionToMatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransformExpressionToMatchResponse
func (client *Client) TransformExpressionToMatchWithOptions(request *TransformExpressionToMatchRequest, runtime *util.RuntimeOptions) (_result *TransformExpressionToMatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		body["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		body["Phase"] = request.Phase
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TransformExpressionToMatch"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransformExpressionToMatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将表达式转换为匹配项
//
// @param request - TransformExpressionToMatchRequest
//
// @return TransformExpressionToMatchResponse
func (client *Client) TransformExpressionToMatch(request *TransformExpressionToMatchRequest) (_result *TransformExpressionToMatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransformExpressionToMatchResponse{}
	_body, _err := client.TransformExpressionToMatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将匹配项转换为表达式
//
// @param tmpReq - TransformMatchToExpressionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransformMatchToExpressionResponse
func (client *Client) TransformMatchToExpressionWithOptions(tmpReq *TransformMatchToExpressionRequest, runtime *util.RuntimeOptions) (_result *TransformMatchToExpressionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TransformMatchToExpressionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Match)) {
		request.MatchShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Match, tea.String("Match"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchShrink)) {
		body["Match"] = request.MatchShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		body["Phase"] = request.Phase
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TransformMatchToExpression"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransformMatchToExpressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将匹配项转换为表达式
//
// @param request - TransformMatchToExpressionRequest
//
// @return TransformMatchToExpressionResponse
func (client *Client) TransformMatchToExpression(request *TransformMatchToExpressionRequest) (_result *TransformMatchToExpressionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransformMatchToExpressionResponse{}
	_body, _err := client.TransformMatchToExpressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为资源列表统一解绑标签
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为资源列表统一解绑标签
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改定制场景策略
//
// @param request - UpdateCustomScenePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomScenePolicyResponse
func (client *Client) UpdateCustomScenePolicyWithOptions(request *UpdateCustomScenePolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomScenePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Objects)) {
		query["Objects"] = request.Objects
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Template)) {
		query["Template"] = request.Template
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomScenePolicy"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomScenePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改定制场景策略
//
// @param request - UpdateCustomScenePolicyRequest
//
// @return UpdateCustomScenePolicyResponse
func (client *Client) UpdateCustomScenePolicy(request *UpdateCustomScenePolicyRequest) (_result *UpdateCustomScenePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomScenePolicyResponse{}
	_body, _err := client.UpdateCustomScenePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重命名账号下的Namespace
//
// @param request - UpdateKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKvNamespaceResponse
func (client *Client) UpdateKvNamespaceWithOptions(request *UpdateKvNamespaceRequest, runtime *util.RuntimeOptions) (_result *UpdateKvNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateKvNamespace"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateKvNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重命名账号下的Namespace
//
// @param request - UpdateKvNamespaceRequest
//
// @return UpdateKvNamespaceResponse
func (client *Client) UpdateKvNamespace(request *UpdateKvNamespaceRequest) (_result *UpdateKvNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateKvNamespaceResponse{}
	_body, _err := client.UpdateKvNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新自定义列表
//
// @param tmpReq - UpdateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateListResponse
func (client *Client) UpdateListWithOptions(tmpReq *UpdateListRequest, runtime *util.RuntimeOptions) (_result *UpdateListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Items)) {
		request.ItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Items, tea.String("Items"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ItemsShrink)) {
		body["Items"] = request.ItemsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateList"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新自定义列表
//
// @param request - UpdateListRequest
//
// @return UpdateListResponse
func (client *Client) UpdateList(request *UpdateListRequest) (_result *UpdateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateListResponse{}
	_body, _err := client.UpdateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新自定义响应页面
//
// @param request - UpdatePageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePageResponse
func (client *Client) UpdatePageWithOptions(request *UpdatePageRequest, runtime *util.RuntimeOptions) (_result *UpdatePageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["ContentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePage"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新自定义响应页面
//
// @param request - UpdatePageRequest
//
// @return UpdatePageResponse
func (client *Client) UpdatePage(request *UpdatePageRequest) (_result *UpdatePageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePageResponse{}
	_body, _err := client.UpdatePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新记录
//
// @param tmpReq - UpdateRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordResponse
func (client *Client) UpdateRecordWithOptions(tmpReq *UpdateRecordRequest, runtime *util.RuntimeOptions) (_result *UpdateRecordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AuthConf)) {
		request.AuthConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthConf, tea.String("AuthConf"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Data)) {
		request.DataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Data, tea.String("Data"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthConfShrink)) {
		query["AuthConf"] = request.AuthConfShrink
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		query["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.DataShrink)) {
		query["Data"] = request.DataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HostPolicy)) {
		query["HostPolicy"] = request.HostPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Proxied)) {
		query["Proxied"] = request.Proxied
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		query["Ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRecord"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新记录
//
// @param request - UpdateRecordRequest
//
// @return UpdateRecordResponse
func (client *Client) UpdateRecord(request *UpdateRecordRequest) (_result *UpdateRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRecordResponse{}
	_body, _err := client.UpdateRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新单个定时预热计划
//
// @param request - UpdateScheduledPreloadExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduledPreloadExecutionResponse
func (client *Client) UpdateScheduledPreloadExecutionWithOptions(request *UpdateScheduledPreloadExecutionRequest, runtime *util.RuntimeOptions) (_result *UpdateScheduledPreloadExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		body["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.SliceLen)) {
		body["SliceLen"] = request.SliceLen
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateScheduledPreloadExecution"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateScheduledPreloadExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新单个定时预热计划
//
// @param request - UpdateScheduledPreloadExecutionRequest
//
// @return UpdateScheduledPreloadExecutionResponse
func (client *Client) UpdateScheduledPreloadExecution(request *UpdateScheduledPreloadExecutionRequest) (_result *UpdateScheduledPreloadExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateScheduledPreloadExecutionResponse{}
	_body, _err := client.UpdateScheduledPreloadExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改站点接入方式
//
// @param request - UpdateSiteAccessTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteAccessTypeResponse
func (client *Client) UpdateSiteAccessTypeWithOptions(request *UpdateSiteAccessTypeRequest, runtime *util.RuntimeOptions) (_result *UpdateSiteAccessTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSiteAccessType"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSiteAccessTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改站点接入方式
//
// @param request - UpdateSiteAccessTypeRequest
//
// @return UpdateSiteAccessTypeResponse
func (client *Client) UpdateSiteAccessType(request *UpdateSiteAccessTypeRequest) (_result *UpdateSiteAccessTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSiteAccessTypeResponse{}
	_body, _err := client.UpdateSiteAccessTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新站点加速区域
//
// @param request - UpdateSiteCoverageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteCoverageResponse
func (client *Client) UpdateSiteCoverageWithOptions(request *UpdateSiteCoverageRequest, runtime *util.RuntimeOptions) (_result *UpdateSiteCoverageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Coverage)) {
		query["Coverage"] = request.Coverage
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSiteCoverage"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSiteCoverageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新站点加速区域
//
// @param request - UpdateSiteCoverageRequest
//
// @return UpdateSiteCoverageResponse
func (client *Client) UpdateSiteCoverage(request *UpdateSiteCoverageRequest) (_result *UpdateSiteCoverageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSiteCoverageResponse{}
	_body, _err := client.UpdateSiteCoverageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改自定义字段
//
// @param tmpReq - UpdateSiteCustomLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteCustomLogResponse
func (client *Client) UpdateSiteCustomLogWithOptions(tmpReq *UpdateSiteCustomLogRequest, runtime *util.RuntimeOptions) (_result *UpdateSiteCustomLogResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSiteCustomLogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Cookies)) {
		request.CookiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cookies, tea.String("Cookies"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RequestHeaders)) {
		request.RequestHeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestHeaders, tea.String("RequestHeaders"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResponseHeaders)) {
		request.ResponseHeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResponseHeaders, tea.String("ResponseHeaders"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CookiesShrink)) {
		body["Cookies"] = request.CookiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RequestHeadersShrink)) {
		body["RequestHeaders"] = request.RequestHeadersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseHeadersShrink)) {
		body["ResponseHeaders"] = request.ResponseHeadersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		body["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSiteCustomLog"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSiteCustomLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自定义字段
//
// @param request - UpdateSiteCustomLogRequest
//
// @return UpdateSiteCustomLogResponse
func (client *Client) UpdateSiteCustomLog(request *UpdateSiteCustomLogRequest) (_result *UpdateSiteCustomLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSiteCustomLogResponse{}
	_body, _err := client.UpdateSiteCustomLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改一个任务投递
//
// @param request - UpdateSiteDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteDeliveryTaskResponse
func (client *Client) UpdateSiteDeliveryTaskWithOptions(request *UpdateSiteDeliveryTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateSiteDeliveryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		body["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.DiscardRate)) {
		body["DiscardRate"] = request.DiscardRate
	}

	if !tea.BoolValue(util.IsUnset(request.FieldName)) {
		body["FieldName"] = request.FieldName
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		body["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSiteDeliveryTask"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSiteDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改一个任务投递
//
// @param request - UpdateSiteDeliveryTaskRequest
//
// @return UpdateSiteDeliveryTaskResponse
func (client *Client) UpdateSiteDeliveryTask(request *UpdateSiteDeliveryTaskRequest) (_result *UpdateSiteDeliveryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSiteDeliveryTaskResponse{}
	_body, _err := client.UpdateSiteDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上下线一个任务投递
//
// @param request - UpdateSiteDeliveryTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteDeliveryTaskStatusResponse
func (client *Client) UpdateSiteDeliveryTaskStatusWithOptions(request *UpdateSiteDeliveryTaskStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateSiteDeliveryTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSiteDeliveryTaskStatus"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSiteDeliveryTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上下线一个任务投递
//
// @param request - UpdateSiteDeliveryTaskStatusRequest
//
// @return UpdateSiteDeliveryTaskStatusResponse
func (client *Client) UpdateSiteDeliveryTaskStatus(request *UpdateSiteDeliveryTaskStatusRequest) (_result *UpdateSiteDeliveryTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSiteDeliveryTaskStatusResponse{}
	_body, _err := client.UpdateSiteDeliveryTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改站点自定义NS
//
// @param request - UpdateSiteVanityNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSiteVanityNSResponse
func (client *Client) UpdateSiteVanityNSWithOptions(request *UpdateSiteVanityNSRequest, runtime *util.RuntimeOptions) (_result *UpdateSiteVanityNSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.VanityNSList)) {
		query["VanityNSList"] = request.VanityNSList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSiteVanityNS"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSiteVanityNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改站点自定义NS
//
// @param request - UpdateSiteVanityNSRequest
//
// @return UpdateSiteVanityNSResponse
func (client *Client) UpdateSiteVanityNS(request *UpdateSiteVanityNSRequest) (_result *UpdateSiteVanityNSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSiteVanityNSResponse{}
	_body, _err := client.UpdateSiteVanityNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改一个用户粒度任务投递
//
// @param request - UpdateUserDeliveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDeliveryTaskResponse
func (client *Client) UpdateUserDeliveryTaskWithOptions(request *UpdateUserDeliveryTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateUserDeliveryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		body["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.DiscardRate)) {
		body["DiscardRate"] = request.DiscardRate
	}

	if !tea.BoolValue(util.IsUnset(request.FieldName)) {
		body["FieldName"] = request.FieldName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserDeliveryTask"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改一个用户粒度任务投递
//
// @param request - UpdateUserDeliveryTaskRequest
//
// @return UpdateUserDeliveryTaskResponse
func (client *Client) UpdateUserDeliveryTask(request *UpdateUserDeliveryTaskRequest) (_result *UpdateUserDeliveryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserDeliveryTaskResponse{}
	_body, _err := client.UpdateUserDeliveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上下线一个用户任务投递
//
// @param request - UpdateUserDeliveryTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDeliveryTaskStatusResponse
func (client *Client) UpdateUserDeliveryTaskStatusWithOptions(request *UpdateUserDeliveryTaskStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateUserDeliveryTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserDeliveryTaskStatus"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserDeliveryTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上下线一个用户任务投递
//
// @param request - UpdateUserDeliveryTaskStatusRequest
//
// @return UpdateUserDeliveryTaskStatusResponse
func (client *Client) UpdateUserDeliveryTaskStatus(request *UpdateUserDeliveryTaskStatusRequest) (_result *UpdateUserDeliveryTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserDeliveryTaskStatusResponse{}
	_body, _err := client.UpdateUserDeliveryTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新WAF规则页面
//
// @param tmpReq - UpdateWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWafRuleResponse
func (client *Client) UpdateWafRuleWithOptions(tmpReq *UpdateWafRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateWafRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateWafRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Config)) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, tea.String("Config"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigShrink)) {
		body["Config"] = request.ConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Position)) {
		body["Position"] = request.Position
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWafRule"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWafRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新WAF规则页面
//
// @param request - UpdateWafRuleRequest
//
// @return UpdateWafRuleResponse
func (client *Client) UpdateWafRule(request *UpdateWafRuleRequest) (_result *UpdateWafRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWafRuleResponse{}
	_body, _err := client.UpdateWafRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新WAF规则集
//
// @param request - UpdateWafRulesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWafRulesetResponse
func (client *Client) UpdateWafRulesetWithOptions(request *UpdateWafRulesetRequest, runtime *util.RuntimeOptions) (_result *UpdateWafRulesetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SiteVersion)) {
		query["SiteVersion"] = request.SiteVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWafRuleset"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWafRulesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新WAF规则集
//
// @param request - UpdateWafRulesetRequest
//
// @return UpdateWafRulesetResponse
func (client *Client) UpdateWafRuleset(request *UpdateWafRulesetRequest) (_result *UpdateWafRulesetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWafRulesetResponse{}
	_body, _err := client.UpdateWafRulesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改等候室
//
// @param tmpReq - UpdateWaitingRoomRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWaitingRoomResponse
func (client *Client) UpdateWaitingRoomWithOptions(tmpReq *UpdateWaitingRoomRequest, runtime *util.RuntimeOptions) (_result *UpdateWaitingRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateWaitingRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HostNameAndPath)) {
		request.HostNameAndPathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HostNameAndPath, tea.String("HostNameAndPath"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CookieName)) {
		query["CookieName"] = request.CookieName
	}

	if !tea.BoolValue(util.IsUnset(request.CustomPageHtml)) {
		query["CustomPageHtml"] = request.CustomPageHtml
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisableSessionRenewalEnable)) {
		query["DisableSessionRenewalEnable"] = request.DisableSessionRenewalEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.HostNameAndPathShrink)) {
		query["HostNameAndPath"] = request.HostNameAndPathShrink
	}

	if !tea.BoolValue(util.IsUnset(request.JsonResponseEnable)) {
		query["JsonResponseEnable"] = request.JsonResponseEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NewUsersPerMinute)) {
		query["NewUsersPerMinute"] = request.NewUsersPerMinute
	}

	if !tea.BoolValue(util.IsUnset(request.QueueAllEnable)) {
		query["QueueAllEnable"] = request.QueueAllEnable
	}

	if !tea.BoolValue(util.IsUnset(request.QueuingMethod)) {
		query["QueuingMethod"] = request.QueuingMethod
	}

	if !tea.BoolValue(util.IsUnset(request.QueuingStatusCode)) {
		query["QueuingStatusCode"] = request.QueuingStatusCode
	}

	if !tea.BoolValue(util.IsUnset(request.SessionDuration)) {
		query["SessionDuration"] = request.SessionDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalActiveUsers)) {
		query["TotalActiveUsers"] = request.TotalActiveUsers
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomId)) {
		query["WaitingRoomId"] = request.WaitingRoomId
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomType)) {
		query["WaitingRoomType"] = request.WaitingRoomType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWaitingRoom"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWaitingRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改等候室
//
// @param request - UpdateWaitingRoomRequest
//
// @return UpdateWaitingRoomResponse
func (client *Client) UpdateWaitingRoom(request *UpdateWaitingRoomRequest) (_result *UpdateWaitingRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWaitingRoomResponse{}
	_body, _err := client.UpdateWaitingRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改等候室事件
//
// @param request - UpdateWaitingRoomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWaitingRoomEventResponse
func (client *Client) UpdateWaitingRoomEventWithOptions(request *UpdateWaitingRoomEventRequest, runtime *util.RuntimeOptions) (_result *UpdateWaitingRoomEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomPageHtml)) {
		query["CustomPageHtml"] = request.CustomPageHtml
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisableSessionRenewalEnable)) {
		query["DisableSessionRenewalEnable"] = request.DisableSessionRenewalEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.JsonResponseEnable)) {
		query["JsonResponseEnable"] = request.JsonResponseEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NewUsersPerMinute)) {
		query["NewUsersPerMinute"] = request.NewUsersPerMinute
	}

	if !tea.BoolValue(util.IsUnset(request.PreQueueEnable)) {
		query["PreQueueEnable"] = request.PreQueueEnable
	}

	if !tea.BoolValue(util.IsUnset(request.PreQueueStartTime)) {
		query["PreQueueStartTime"] = request.PreQueueStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.QueuingMethod)) {
		query["QueuingMethod"] = request.QueuingMethod
	}

	if !tea.BoolValue(util.IsUnset(request.QueuingStatusCode)) {
		query["QueuingStatusCode"] = request.QueuingStatusCode
	}

	if !tea.BoolValue(util.IsUnset(request.RandomPreQueueEnable)) {
		query["RandomPreQueueEnable"] = request.RandomPreQueueEnable
	}

	if !tea.BoolValue(util.IsUnset(request.SessionDuration)) {
		query["SessionDuration"] = request.SessionDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TotalActiveUsers)) {
		query["TotalActiveUsers"] = request.TotalActiveUsers
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomEventId)) {
		query["WaitingRoomEventId"] = request.WaitingRoomEventId
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomType)) {
		query["WaitingRoomType"] = request.WaitingRoomType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWaitingRoomEvent"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWaitingRoomEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改等候室事件
//
// @param request - UpdateWaitingRoomEventRequest
//
// @return UpdateWaitingRoomEventResponse
func (client *Client) UpdateWaitingRoomEvent(request *UpdateWaitingRoomEventRequest) (_result *UpdateWaitingRoomEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWaitingRoomEventResponse{}
	_body, _err := client.UpdateWaitingRoomEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改等候室规则
//
// @param request - UpdateWaitingRoomRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWaitingRoomRuleResponse
func (client *Client) UpdateWaitingRoomRuleWithOptions(request *UpdateWaitingRoomRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateWaitingRoomRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Rule)) {
		query["Rule"] = request.Rule
	}

	if !tea.BoolValue(util.IsUnset(request.RuleEnable)) {
		query["RuleEnable"] = request.RuleEnable
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.WaitingRoomRuleId)) {
		query["WaitingRoomRuleId"] = request.WaitingRoomRuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWaitingRoomRule"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWaitingRoomRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改等候室规则
//
// @param request - UpdateWaitingRoomRuleRequest
//
// @return UpdateWaitingRoomRuleResponse
func (client *Client) UpdateWaitingRoomRule(request *UpdateWaitingRoomRuleRequest) (_result *UpdateWaitingRoomRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWaitingRoomRuleResponse{}
	_body, _err := client.UpdateWaitingRoomRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 缓存刷新文件上传
//
// @param request - UploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadFileResponse
func (client *Client) UploadFileWithOptions(request *UploadFileRequest, runtime *util.RuntimeOptions) (_result *UploadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UploadTaskName)) {
		query["UploadTaskName"] = request.UploadTaskName
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadFile"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 缓存刷新文件上传
//
// @param request - UploadFileRequest
//
// @return UploadFileResponse
func (client *Client) UploadFile(request *UploadFileRequest) (_result *UploadFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadFileResponse{}
	_body, _err := client.UploadFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadFileAdvance(request *UploadFileAdvanceRequest, runtime *util.RuntimeOptions) (_result *UploadFileResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ESA"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	uploadFileReq := &UploadFileRequest{}
	openapiutil.Convert(request, uploadFileReq)
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		uploadFileReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	uploadFileResp, _err := client.UploadFileWithOptions(uploadFileReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = uploadFileResp
	return _result, _err
}

// Summary:
//
// 校验站点的归属
//
// @param request - VerifySiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifySiteResponse
func (client *Client) VerifySiteWithOptions(request *VerifySiteRequest, runtime *util.RuntimeOptions) (_result *VerifySiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		query["SiteId"] = request.SiteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifySite"),
		Version:     tea.String("2024-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifySiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验站点的归属
//
// @param request - VerifySiteRequest
//
// @return VerifySiteResponse
func (client *Client) VerifySite(request *VerifySiteRequest) (_result *VerifySiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifySiteResponse{}
	_body, _err := client.VerifySiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
