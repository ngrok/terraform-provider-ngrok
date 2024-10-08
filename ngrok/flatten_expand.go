package ngrok

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func flattenEmpty(obj *restapi.Empty) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})

	return []interface{}{m}
}

func flattenEmptySlice(objs *[]restapi.Empty) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEmpty(&v))
	}
	return sl
}

func expandEmpty(in interface{}) *restapi.Empty {
	if in == nil {
		return nil
	}
	var obj restapi.Empty
	return &obj
}

func expandEmptySlice(in interface{}) *[]restapi.Empty {
	var out []restapi.Empty
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEmpty(v))
	}
	return &out
}

func flattenItem(obj *restapi.Item) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID

	return []interface{}{m}
}

func flattenItemSlice(objs *[]restapi.Item) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenItem(&v))
	}
	return sl
}

func expandItem(in interface{}) *restapi.Item {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.Item
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	return &obj
}

func expandItemSlice(in interface{}) *[]restapi.Item {
	var out []restapi.Item
	for _, v := range in.([]interface{}) {
		out = append(out, *expandItem(v))
	}
	return &out
}

func flattenPaging(obj *restapi.Paging) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["before_id"] = obj.BeforeID
	m["limit"] = obj.Limit

	return []interface{}{m}
}

func flattenPagingSlice(objs *[]restapi.Paging) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenPaging(&v))
	}
	return sl
}

func expandPaging(in interface{}) *restapi.Paging {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.Paging
	if v, ok := m["before_id"]; ok {
		obj.BeforeID = expandString(v)
	}
	if v, ok := m["limit"]; ok {
		obj.Limit = expandString(v)
	}
	return &obj
}

func expandPagingSlice(in interface{}) *[]restapi.Paging {
	var out []restapi.Paging
	for _, v := range in.([]interface{}) {
		out = append(out, *expandPaging(v))
	}
	return &out
}

func flattenError(obj *restapi.Error) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["error_code"] = obj.ErrorCode
	m["status_code"] = obj.StatusCode
	m["msg"] = obj.Msg
	m["details"] = obj.Details

	return []interface{}{m}
}

func flattenErrorSlice(objs *[]restapi.Error) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenError(&v))
	}
	return sl
}

func expandError(in interface{}) *restapi.Error {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.Error
	if v, ok := m["error_code"]; ok {
		obj.ErrorCode = *expandString(v)
	}
	if v, ok := m["status_code"]; ok {
		obj.StatusCode = *expandInt32(v)
	}
	if v, ok := m["msg"]; ok {
		obj.Msg = *expandString(v)
	}
	if v, ok := m["details"]; ok {
		obj.Details = *expandStringMap(v)
	}
	return &obj
}

func expandErrorSlice(in interface{}) *[]restapi.Error {
	var out []restapi.Error
	for _, v := range in.([]interface{}) {
		out = append(out, *expandError(v))
	}
	return &out
}

func flattenRef(obj *restapi.Ref) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI

	return []interface{}{m}
}

func flattenRefSlice(objs *[]restapi.Ref) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenRef(&v))
	}
	return sl
}

func expandRef(in interface{}) *restapi.Ref {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.Ref
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	return &obj
}

func expandRefSlice(in interface{}) *[]restapi.Ref {
	var out []restapi.Ref
	for _, v := range in.([]interface{}) {
		out = append(out, *expandRef(v))
	}
	return &out
}

func flattenAbuseReport(obj *restapi.AbuseReport) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["urls"] = obj.URLs
	m["metadata"] = obj.Metadata
	m["status"] = obj.Status
	m["hostnames"] = flattenAbuseReportHostnameSlice(&obj.Hostnames)

	return []interface{}{m}
}

func flattenAbuseReportSlice(objs *[]restapi.AbuseReport) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAbuseReport(&v))
	}
	return sl
}

func expandAbuseReport(in interface{}) *restapi.AbuseReport {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AbuseReport
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["urls"]; ok {
		obj.URLs = *expandStringSlice(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["status"]; ok {
		obj.Status = *expandString(v)
	}
	if v, ok := m["hostnames"]; ok {
		obj.Hostnames = *expandAbuseReportHostnameSlice(v)
	}
	return &obj
}

func expandAbuseReportSlice(in interface{}) *[]restapi.AbuseReport {
	var out []restapi.AbuseReport
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAbuseReport(v))
	}
	return &out
}

func flattenAbuseReportHostname(obj *restapi.AbuseReportHostname) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["hostname"] = obj.Hostname
	m["status"] = obj.Status

	return []interface{}{m}
}

func flattenAbuseReportHostnameSlice(objs *[]restapi.AbuseReportHostname) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAbuseReportHostname(&v))
	}
	return sl
}

func expandAbuseReportHostname(in interface{}) *restapi.AbuseReportHostname {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AbuseReportHostname
	if v, ok := m["hostname"]; ok {
		obj.Hostname = *expandString(v)
	}
	if v, ok := m["status"]; ok {
		obj.Status = *expandString(v)
	}
	return &obj
}

func expandAbuseReportHostnameSlice(in interface{}) *[]restapi.AbuseReportHostname {
	var out []restapi.AbuseReportHostname
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAbuseReportHostname(v))
	}
	return &out
}

func flattenAbuseReportCreate(obj *restapi.AbuseReportCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["urls"] = obj.URLs
	m["metadata"] = obj.Metadata

	return []interface{}{m}
}

func flattenAbuseReportCreateSlice(objs *[]restapi.AbuseReportCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAbuseReportCreate(&v))
	}
	return sl
}

func expandAbuseReportCreate(in interface{}) *restapi.AbuseReportCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AbuseReportCreate
	if v, ok := m["urls"]; ok {
		obj.URLs = *expandStringSlice(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	return &obj
}

func expandAbuseReportCreateSlice(in interface{}) *[]restapi.AbuseReportCreate {
	var out []restapi.AbuseReportCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAbuseReportCreate(v))
	}
	return &out
}

func flattenAgentIngressCreate(obj *restapi.AgentIngressCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["domain"] = obj.Domain
	m["certificate_management_policy"] = flattenAgentIngressCertPolicy(obj.CertificateManagementPolicy)

	return []interface{}{m}
}

func flattenAgentIngressCreateSlice(objs *[]restapi.AgentIngressCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAgentIngressCreate(&v))
	}
	return sl
}

func expandAgentIngressCreate(in interface{}) *restapi.AgentIngressCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AgentIngressCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["domain"]; ok {
		obj.Domain = *expandString(v)
	}
	if v, ok := m["certificate_management_policy"]; ok {
		obj.CertificateManagementPolicy = expandAgentIngressCertPolicy(v)
	}
	return &obj
}

func expandAgentIngressCreateSlice(in interface{}) *[]restapi.AgentIngressCreate {
	var out []restapi.AgentIngressCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAgentIngressCreate(v))
	}
	return &out
}

func flattenAgentIngressUpdate(obj *restapi.AgentIngressUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["certificate_management_policy"] = flattenAgentIngressCertPolicy(obj.CertificateManagementPolicy)

	return []interface{}{m}
}

func flattenAgentIngressUpdateSlice(objs *[]restapi.AgentIngressUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAgentIngressUpdate(&v))
	}
	return sl
}

func expandAgentIngressUpdate(in interface{}) *restapi.AgentIngressUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AgentIngressUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["certificate_management_policy"]; ok {
		obj.CertificateManagementPolicy = expandAgentIngressCertPolicy(v)
	}
	return &obj
}

func expandAgentIngressUpdateSlice(in interface{}) *[]restapi.AgentIngressUpdate {
	var out []restapi.AgentIngressUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAgentIngressUpdate(v))
	}
	return &out
}

func flattenAgentIngress(obj *restapi.AgentIngress) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["domain"] = obj.Domain
	m["ns_targets"] = obj.NSTargets
	m["region_domains"] = obj.RegionDomains
	m["created_at"] = obj.CreatedAt
	m["certificate_management_policy"] = flattenAgentIngressCertPolicy(obj.CertificateManagementPolicy)
	m["certificate_management_status"] = flattenAgentIngressCertStatus(obj.CertificateManagementStatus)

	return []interface{}{m}
}

func flattenAgentIngressSlice(objs *[]restapi.AgentIngress) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAgentIngress(&v))
	}
	return sl
}

func expandAgentIngress(in interface{}) *restapi.AgentIngress {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AgentIngress
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["domain"]; ok {
		obj.Domain = *expandString(v)
	}
	if v, ok := m["ns_targets"]; ok {
		obj.NSTargets = *expandStringSlice(v)
	}
	if v, ok := m["region_domains"]; ok {
		obj.RegionDomains = *expandStringSlice(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["certificate_management_policy"]; ok {
		obj.CertificateManagementPolicy = expandAgentIngressCertPolicy(v)
	}
	if v, ok := m["certificate_management_status"]; ok {
		obj.CertificateManagementStatus = expandAgentIngressCertStatus(v)
	}
	return &obj
}

func expandAgentIngressSlice(in interface{}) *[]restapi.AgentIngress {
	var out []restapi.AgentIngress
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAgentIngress(v))
	}
	return &out
}

func flattenAgentIngressList(obj *restapi.AgentIngressList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["ingresses"] = flattenAgentIngressSlice(&obj.Ingresses)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenAgentIngressListSlice(objs *[]restapi.AgentIngressList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAgentIngressList(&v))
	}
	return sl
}

func expandAgentIngressList(in interface{}) *restapi.AgentIngressList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AgentIngressList
	if v, ok := m["ingresses"]; ok {
		obj.Ingresses = *expandAgentIngressSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandAgentIngressListSlice(in interface{}) *[]restapi.AgentIngressList {
	var out []restapi.AgentIngressList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAgentIngressList(v))
	}
	return &out
}

func flattenAgentIngressCertPolicy(obj *restapi.AgentIngressCertPolicy) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["authority"] = obj.Authority
	m["private_key_type"] = obj.PrivateKeyType

	return []interface{}{m}
}

func flattenAgentIngressCertPolicySlice(objs *[]restapi.AgentIngressCertPolicy) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAgentIngressCertPolicy(&v))
	}
	return sl
}

func expandAgentIngressCertPolicy(in interface{}) *restapi.AgentIngressCertPolicy {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AgentIngressCertPolicy
	if v, ok := m["authority"]; ok {
		obj.Authority = *expandString(v)
	}
	if v, ok := m["private_key_type"]; ok {
		obj.PrivateKeyType = *expandString(v)
	}
	return &obj
}

func expandAgentIngressCertPolicySlice(in interface{}) *[]restapi.AgentIngressCertPolicy {
	var out []restapi.AgentIngressCertPolicy
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAgentIngressCertPolicy(v))
	}
	return &out
}

func flattenAgentIngressCertStatus(obj *restapi.AgentIngressCertStatus) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["renews_at"] = obj.RenewsAt
	m["provisioning_job"] = flattenAgentIngressCertJob(obj.ProvisioningJob)

	return []interface{}{m}
}

func flattenAgentIngressCertStatusSlice(objs *[]restapi.AgentIngressCertStatus) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAgentIngressCertStatus(&v))
	}
	return sl
}

func expandAgentIngressCertStatus(in interface{}) *restapi.AgentIngressCertStatus {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AgentIngressCertStatus
	if v, ok := m["renews_at"]; ok {
		obj.RenewsAt = expandString(v)
	}
	if v, ok := m["provisioning_job"]; ok {
		obj.ProvisioningJob = expandAgentIngressCertJob(v)
	}
	return &obj
}

func expandAgentIngressCertStatusSlice(in interface{}) *[]restapi.AgentIngressCertStatus {
	var out []restapi.AgentIngressCertStatus
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAgentIngressCertStatus(v))
	}
	return &out
}

func flattenAgentIngressCertJob(obj *restapi.AgentIngressCertJob) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["error_code"] = obj.ErrorCode
	m["msg"] = obj.Msg
	m["started_at"] = obj.StartedAt
	m["retries_at"] = obj.RetriesAt

	return []interface{}{m}
}

func flattenAgentIngressCertJobSlice(objs *[]restapi.AgentIngressCertJob) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAgentIngressCertJob(&v))
	}
	return sl
}

func expandAgentIngressCertJob(in interface{}) *restapi.AgentIngressCertJob {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AgentIngressCertJob
	if v, ok := m["error_code"]; ok {
		obj.ErrorCode = expandString(v)
	}
	if v, ok := m["msg"]; ok {
		obj.Msg = *expandString(v)
	}
	if v, ok := m["started_at"]; ok {
		obj.StartedAt = *expandString(v)
	}
	if v, ok := m["retries_at"]; ok {
		obj.RetriesAt = expandString(v)
	}
	return &obj
}

func expandAgentIngressCertJobSlice(in interface{}) *[]restapi.AgentIngressCertJob {
	var out []restapi.AgentIngressCertJob
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAgentIngressCertJob(v))
	}
	return &out
}

func flattenAPIKeyCreate(obj *restapi.APIKeyCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["owner_id"] = obj.OwnerID
	m["owner_email"] = obj.OwnerEmail

	return []interface{}{m}
}

func flattenAPIKeyCreateSlice(objs *[]restapi.APIKeyCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAPIKeyCreate(&v))
	}
	return sl
}

func expandAPIKeyCreate(in interface{}) *restapi.APIKeyCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.APIKeyCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["owner_id"]; ok {
		obj.OwnerID = expandString(v)
	}
	if v, ok := m["owner_email"]; ok {
		obj.OwnerEmail = *expandString(v)
	}
	return &obj
}

func expandAPIKeyCreateSlice(in interface{}) *[]restapi.APIKeyCreate {
	var out []restapi.APIKeyCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAPIKeyCreate(v))
	}
	return &out
}

func flattenAPIKeyUpdate(obj *restapi.APIKeyUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata

	return []interface{}{m}
}

func flattenAPIKeyUpdateSlice(objs *[]restapi.APIKeyUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAPIKeyUpdate(&v))
	}
	return sl
}

func expandAPIKeyUpdate(in interface{}) *restapi.APIKeyUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.APIKeyUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	return &obj
}

func expandAPIKeyUpdateSlice(in interface{}) *[]restapi.APIKeyUpdate {
	var out []restapi.APIKeyUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAPIKeyUpdate(v))
	}
	return &out
}

func flattenAPIKey(obj *restapi.APIKey) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["created_at"] = obj.CreatedAt
	m["token"] = obj.Token
	m["owner_id"] = obj.OwnerID

	return []interface{}{m}
}

func flattenAPIKeySlice(objs *[]restapi.APIKey) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAPIKey(&v))
	}
	return sl
}

func expandAPIKey(in interface{}) *restapi.APIKey {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.APIKey
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["token"]; ok {
		obj.Token = expandString(v)
	}
	if v, ok := m["owner_id"]; ok {
		obj.OwnerID = expandString(v)
	}
	return &obj
}

func expandAPIKeySlice(in interface{}) *[]restapi.APIKey {
	var out []restapi.APIKey
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAPIKey(v))
	}
	return &out
}

func flattenAPIKeyList(obj *restapi.APIKeyList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["keys"] = flattenAPIKeySlice(&obj.Keys)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenAPIKeyListSlice(objs *[]restapi.APIKeyList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAPIKeyList(&v))
	}
	return sl
}

func expandAPIKeyList(in interface{}) *restapi.APIKeyList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.APIKeyList
	if v, ok := m["keys"]; ok {
		obj.Keys = *expandAPIKeySlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandAPIKeyListSlice(in interface{}) *[]restapi.APIKeyList {
	var out []restapi.APIKeyList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAPIKeyList(v))
	}
	return &out
}

func flattenApplicationSession(obj *restapi.ApplicationSession) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["public_url"] = obj.PublicURL
	m["browser_session"] = flattenBrowserSession(&obj.BrowserSession)
	m["application_user"] = flattenRef(obj.ApplicationUser)
	m["created_at"] = obj.CreatedAt
	m["last_active"] = obj.LastActive
	m["expires_at"] = obj.ExpiresAt
	m["endpoint"] = flattenRef(obj.Endpoint)
	m["edge"] = flattenRef(obj.Edge)
	m["route"] = flattenRef(obj.Route)

	return []interface{}{m}
}

func flattenApplicationSessionSlice(objs *[]restapi.ApplicationSession) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenApplicationSession(&v))
	}
	return sl
}

func expandApplicationSession(in interface{}) *restapi.ApplicationSession {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ApplicationSession
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["public_url"]; ok {
		obj.PublicURL = *expandString(v)
	}
	if v, ok := m["browser_session"]; ok {
		obj.BrowserSession = *expandBrowserSession(v)
	}
	if v, ok := m["application_user"]; ok {
		obj.ApplicationUser = expandRef(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["last_active"]; ok {
		obj.LastActive = *expandString(v)
	}
	if v, ok := m["expires_at"]; ok {
		obj.ExpiresAt = *expandString(v)
	}
	if v, ok := m["endpoint"]; ok {
		obj.Endpoint = expandRef(v)
	}
	if v, ok := m["edge"]; ok {
		obj.Edge = expandRef(v)
	}
	if v, ok := m["route"]; ok {
		obj.Route = expandRef(v)
	}
	return &obj
}

func expandApplicationSessionSlice(in interface{}) *[]restapi.ApplicationSession {
	var out []restapi.ApplicationSession
	for _, v := range in.([]interface{}) {
		out = append(out, *expandApplicationSession(v))
	}
	return &out
}

func flattenApplicationSessionList(obj *restapi.ApplicationSessionList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["application_sessions"] = flattenApplicationSessionSlice(&obj.ApplicationSessions)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenApplicationSessionListSlice(objs *[]restapi.ApplicationSessionList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenApplicationSessionList(&v))
	}
	return sl
}

func expandApplicationSessionList(in interface{}) *restapi.ApplicationSessionList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ApplicationSessionList
	if v, ok := m["application_sessions"]; ok {
		obj.ApplicationSessions = *expandApplicationSessionSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandApplicationSessionListSlice(in interface{}) *[]restapi.ApplicationSessionList {
	var out []restapi.ApplicationSessionList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandApplicationSessionList(v))
	}
	return &out
}

func flattenBrowserSession(obj *restapi.BrowserSession) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["user_agent"] = flattenUserAgent(&obj.UserAgent)
	m["ip_address"] = obj.IPAddress
	m["location"] = flattenLocation(obj.Location)

	return []interface{}{m}
}

func flattenBrowserSessionSlice(objs *[]restapi.BrowserSession) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenBrowserSession(&v))
	}
	return sl
}

func expandBrowserSession(in interface{}) *restapi.BrowserSession {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.BrowserSession
	if v, ok := m["user_agent"]; ok {
		obj.UserAgent = *expandUserAgent(v)
	}
	if v, ok := m["ip_address"]; ok {
		obj.IPAddress = *expandString(v)
	}
	if v, ok := m["location"]; ok {
		obj.Location = expandLocation(v)
	}
	return &obj
}

func expandBrowserSessionSlice(in interface{}) *[]restapi.BrowserSession {
	var out []restapi.BrowserSession
	for _, v := range in.([]interface{}) {
		out = append(out, *expandBrowserSession(v))
	}
	return &out
}

func flattenUserAgent(obj *restapi.UserAgent) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["raw"] = obj.Raw
	m["browser_name"] = obj.BrowserName
	m["browser_version"] = obj.BrowserVersion
	m["device_type"] = obj.DeviceType
	m["os_name"] = obj.OSName
	m["os_version"] = obj.OSVersion

	return []interface{}{m}
}

func flattenUserAgentSlice(objs *[]restapi.UserAgent) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenUserAgent(&v))
	}
	return sl
}

func expandUserAgent(in interface{}) *restapi.UserAgent {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.UserAgent
	if v, ok := m["raw"]; ok {
		obj.Raw = *expandString(v)
	}
	if v, ok := m["browser_name"]; ok {
		obj.BrowserName = *expandString(v)
	}
	if v, ok := m["browser_version"]; ok {
		obj.BrowserVersion = *expandString(v)
	}
	if v, ok := m["device_type"]; ok {
		obj.DeviceType = *expandString(v)
	}
	if v, ok := m["os_name"]; ok {
		obj.OSName = *expandString(v)
	}
	if v, ok := m["os_version"]; ok {
		obj.OSVersion = *expandString(v)
	}
	return &obj
}

func expandUserAgentSlice(in interface{}) *[]restapi.UserAgent {
	var out []restapi.UserAgent
	for _, v := range in.([]interface{}) {
		out = append(out, *expandUserAgent(v))
	}
	return &out
}

func flattenLocation(obj *restapi.Location) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["country_code"] = obj.CountryCode
	m["latitude"] = obj.Latitude
	m["longitude"] = obj.Longitude
	m["lat_long_radius_km"] = obj.LatLongRadiusKm

	return []interface{}{m}
}

func flattenLocationSlice(objs *[]restapi.Location) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenLocation(&v))
	}
	return sl
}

func expandLocation(in interface{}) *restapi.Location {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.Location
	if v, ok := m["country_code"]; ok {
		obj.CountryCode = expandString(v)
	}
	if v, ok := m["latitude"]; ok {
		obj.Latitude = expandFloat64(v)
	}
	if v, ok := m["longitude"]; ok {
		obj.Longitude = expandFloat64(v)
	}
	if v, ok := m["lat_long_radius_km"]; ok {
		obj.LatLongRadiusKm = expandUint64(v)
	}
	return &obj
}

func expandLocationSlice(in interface{}) *[]restapi.Location {
	var out []restapi.Location
	for _, v := range in.([]interface{}) {
		out = append(out, *expandLocation(v))
	}
	return &out
}

func flattenApplicationUser(obj *restapi.ApplicationUser) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["identity_provider"] = flattenIdentityProvider(&obj.IdentityProvider)
	m["provider_user_id"] = obj.ProviderUserID
	m["username"] = obj.Username
	m["email"] = obj.Email
	m["name"] = obj.Name
	m["created_at"] = obj.CreatedAt
	m["last_active"] = obj.LastActive
	m["last_login"] = obj.LastLogin

	return []interface{}{m}
}

func flattenApplicationUserSlice(objs *[]restapi.ApplicationUser) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenApplicationUser(&v))
	}
	return sl
}

func expandApplicationUser(in interface{}) *restapi.ApplicationUser {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ApplicationUser
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["identity_provider"]; ok {
		obj.IdentityProvider = *expandIdentityProvider(v)
	}
	if v, ok := m["provider_user_id"]; ok {
		obj.ProviderUserID = *expandString(v)
	}
	if v, ok := m["username"]; ok {
		obj.Username = *expandString(v)
	}
	if v, ok := m["email"]; ok {
		obj.Email = *expandString(v)
	}
	if v, ok := m["name"]; ok {
		obj.Name = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["last_active"]; ok {
		obj.LastActive = *expandString(v)
	}
	if v, ok := m["last_login"]; ok {
		obj.LastLogin = *expandString(v)
	}
	return &obj
}

func expandApplicationUserSlice(in interface{}) *[]restapi.ApplicationUser {
	var out []restapi.ApplicationUser
	for _, v := range in.([]interface{}) {
		out = append(out, *expandApplicationUser(v))
	}
	return &out
}

func flattenApplicationUserList(obj *restapi.ApplicationUserList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["application_users"] = flattenApplicationUserSlice(&obj.ApplicationUsers)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenApplicationUserListSlice(objs *[]restapi.ApplicationUserList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenApplicationUserList(&v))
	}
	return sl
}

func expandApplicationUserList(in interface{}) *restapi.ApplicationUserList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ApplicationUserList
	if v, ok := m["application_users"]; ok {
		obj.ApplicationUsers = *expandApplicationUserSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandApplicationUserListSlice(in interface{}) *[]restapi.ApplicationUserList {
	var out []restapi.ApplicationUserList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandApplicationUserList(v))
	}
	return &out
}

func flattenIdentityProvider(obj *restapi.IdentityProvider) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["name"] = obj.Name
	m["url"] = obj.URL

	return []interface{}{m}
}

func flattenIdentityProviderSlice(objs *[]restapi.IdentityProvider) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIdentityProvider(&v))
	}
	return sl
}

func expandIdentityProvider(in interface{}) *restapi.IdentityProvider {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IdentityProvider
	if v, ok := m["name"]; ok {
		obj.Name = *expandString(v)
	}
	if v, ok := m["url"]; ok {
		obj.URL = *expandString(v)
	}
	return &obj
}

func expandIdentityProviderSlice(in interface{}) *[]restapi.IdentityProvider {
	var out []restapi.IdentityProvider
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIdentityProvider(v))
	}
	return &out
}

func flattenTunnelSession(obj *restapi.TunnelSession) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["agent_version"] = obj.AgentVersion
	m["credential"] = flattenRef(&obj.Credential)
	m["id"] = obj.ID
	m["ip"] = obj.IP
	m["metadata"] = obj.Metadata
	m["os"] = obj.OS
	m["region"] = obj.Region
	m["started_at"] = obj.StartedAt
	m["transport"] = obj.Transport
	m["uri"] = obj.URI

	return []interface{}{m}
}

func flattenTunnelSessionSlice(objs *[]restapi.TunnelSession) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnelSession(&v))
	}
	return sl
}

func expandTunnelSession(in interface{}) *restapi.TunnelSession {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TunnelSession
	if v, ok := m["agent_version"]; ok {
		obj.AgentVersion = *expandString(v)
	}
	if v, ok := m["credential"]; ok {
		obj.Credential = *expandRef(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["ip"]; ok {
		obj.IP = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["os"]; ok {
		obj.OS = *expandString(v)
	}
	if v, ok := m["region"]; ok {
		obj.Region = *expandString(v)
	}
	if v, ok := m["started_at"]; ok {
		obj.StartedAt = *expandString(v)
	}
	if v, ok := m["transport"]; ok {
		obj.Transport = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	return &obj
}

func expandTunnelSessionSlice(in interface{}) *[]restapi.TunnelSession {
	var out []restapi.TunnelSession
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnelSession(v))
	}
	return &out
}

func flattenAgentVersionDeprecated(obj *restapi.AgentVersionDeprecated) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["next_min"] = obj.NextMin
	m["next_date"] = obj.NextDate
	m["msg"] = obj.Msg

	return []interface{}{m}
}

func flattenAgentVersionDeprecatedSlice(objs *[]restapi.AgentVersionDeprecated) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAgentVersionDeprecated(&v))
	}
	return sl
}

func expandAgentVersionDeprecated(in interface{}) *restapi.AgentVersionDeprecated {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AgentVersionDeprecated
	if v, ok := m["next_min"]; ok {
		obj.NextMin = *expandString(v)
	}
	if v, ok := m["next_date"]; ok {
		obj.NextDate = *expandString(v)
	}
	if v, ok := m["msg"]; ok {
		obj.Msg = *expandString(v)
	}
	return &obj
}

func expandAgentVersionDeprecatedSlice(in interface{}) *[]restapi.AgentVersionDeprecated {
	var out []restapi.AgentVersionDeprecated
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAgentVersionDeprecated(v))
	}
	return &out
}

func flattenTunnelSessionUserAgent(obj *restapi.TunnelSessionUserAgent) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["products"] = flattenTunnelSessionProductSlice(&obj.Products)

	return []interface{}{m}
}

func flattenTunnelSessionUserAgentSlice(objs *[]restapi.TunnelSessionUserAgent) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnelSessionUserAgent(&v))
	}
	return sl
}

func expandTunnelSessionUserAgent(in interface{}) *restapi.TunnelSessionUserAgent {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TunnelSessionUserAgent
	if v, ok := m["products"]; ok {
		obj.Products = *expandTunnelSessionProductSlice(v)
	}
	return &obj
}

func expandTunnelSessionUserAgentSlice(in interface{}) *[]restapi.TunnelSessionUserAgent {
	var out []restapi.TunnelSessionUserAgent
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnelSessionUserAgent(v))
	}
	return &out
}

func flattenTunnelSessionProduct(obj *restapi.TunnelSessionProduct) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["name"] = obj.Name
	m["version"] = obj.Version
	m["comment"] = obj.Comment

	return []interface{}{m}
}

func flattenTunnelSessionProductSlice(objs *[]restapi.TunnelSessionProduct) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnelSessionProduct(&v))
	}
	return sl
}

func expandTunnelSessionProduct(in interface{}) *restapi.TunnelSessionProduct {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TunnelSessionProduct
	if v, ok := m["name"]; ok {
		obj.Name = *expandString(v)
	}
	if v, ok := m["version"]; ok {
		obj.Version = *expandString(v)
	}
	if v, ok := m["comment"]; ok {
		obj.Comment = *expandString(v)
	}
	return &obj
}

func expandTunnelSessionProductSlice(in interface{}) *[]restapi.TunnelSessionProduct {
	var out []restapi.TunnelSessionProduct
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnelSessionProduct(v))
	}
	return &out
}

func flattenTunnelSessionList(obj *restapi.TunnelSessionList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["tunnel_sessions"] = flattenTunnelSessionSlice(&obj.TunnelSessions)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenTunnelSessionListSlice(objs *[]restapi.TunnelSessionList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnelSessionList(&v))
	}
	return sl
}

func expandTunnelSessionList(in interface{}) *restapi.TunnelSessionList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TunnelSessionList
	if v, ok := m["tunnel_sessions"]; ok {
		obj.TunnelSessions = *expandTunnelSessionSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandTunnelSessionListSlice(in interface{}) *[]restapi.TunnelSessionList {
	var out []restapi.TunnelSessionList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnelSessionList(v))
	}
	return &out
}

func flattenTunnelSessionsUpdate(obj *restapi.TunnelSessionsUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["version"] = obj.Version

	return []interface{}{m}
}

func flattenTunnelSessionsUpdateSlice(objs *[]restapi.TunnelSessionsUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnelSessionsUpdate(&v))
	}
	return sl
}

func expandTunnelSessionsUpdate(in interface{}) *restapi.TunnelSessionsUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TunnelSessionsUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["version"]; ok {
		obj.Version = *expandString(v)
	}
	return &obj
}

func expandTunnelSessionsUpdateSlice(in interface{}) *[]restapi.TunnelSessionsUpdate {
	var out []restapi.TunnelSessionsUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnelSessionsUpdate(v))
	}
	return &out
}

func flattenAuditEventDashLogin(obj *restapi.AuditEventDashLogin) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["account_id"] = obj.AccountID
	m["user_id"] = obj.UserID
	m["remote_addr"] = obj.RemoteAddr
	m["email"] = obj.Email

	return []interface{}{m}
}

func flattenAuditEventDashLoginSlice(objs *[]restapi.AuditEventDashLogin) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAuditEventDashLogin(&v))
	}
	return sl
}

func expandAuditEventDashLogin(in interface{}) *restapi.AuditEventDashLogin {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AuditEventDashLogin
	if v, ok := m["account_id"]; ok {
		obj.AccountID = *expandString(v)
	}
	if v, ok := m["user_id"]; ok {
		obj.UserID = *expandString(v)
	}
	if v, ok := m["remote_addr"]; ok {
		obj.RemoteAddr = *expandString(v)
	}
	if v, ok := m["email"]; ok {
		obj.Email = *expandString(v)
	}
	return &obj
}

func expandAuditEventDashLoginSlice(in interface{}) *[]restapi.AuditEventDashLogin {
	var out []restapi.AuditEventDashLogin
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAuditEventDashLogin(v))
	}
	return &out
}

func flattenAuditEventEndpoint(obj *restapi.AuditEventEndpoint) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["account_id"] = obj.AccountID
	m["region"] = obj.Region
	m["ingress_region"] = obj.IngressRegion
	m["created_at"] = obj.CreatedAt
	m["updated_at"] = obj.UpdatedAt
	m["url"] = obj.URL
	m["proto"] = obj.Proto
	m["domain_id"] = obj.DomainID
	m["tcp_addr_id"] = obj.TCPAddrID
	m["rank"] = obj.Rank
	m["static_tunnel_id"] = obj.StaticTunnelID
	m["static_tunnel_region"] = obj.StaticTunnelRegion
	m["static_tunnel_session_id"] = obj.StaticTunnelSessionID
	m["edge_id"] = obj.EdgeID
	m["type"] = obj.Type

	return []interface{}{m}
}

func flattenAuditEventEndpointSlice(objs *[]restapi.AuditEventEndpoint) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAuditEventEndpoint(&v))
	}
	return sl
}

func expandAuditEventEndpoint(in interface{}) *restapi.AuditEventEndpoint {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AuditEventEndpoint
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["account_id"]; ok {
		obj.AccountID = *expandString(v)
	}
	if v, ok := m["region"]; ok {
		obj.Region = *expandString(v)
	}
	if v, ok := m["ingress_region"]; ok {
		obj.IngressRegion = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["updated_at"]; ok {
		obj.UpdatedAt = *expandString(v)
	}
	if v, ok := m["url"]; ok {
		obj.URL = *expandString(v)
	}
	if v, ok := m["proto"]; ok {
		obj.Proto = *expandString(v)
	}
	if v, ok := m["domain_id"]; ok {
		obj.DomainID = *expandString(v)
	}
	if v, ok := m["tcp_addr_id"]; ok {
		obj.TCPAddrID = *expandString(v)
	}
	if v, ok := m["rank"]; ok {
		obj.Rank = *expandInt32(v)
	}
	if v, ok := m["static_tunnel_id"]; ok {
		obj.StaticTunnelID = *expandString(v)
	}
	if v, ok := m["static_tunnel_region"]; ok {
		obj.StaticTunnelRegion = *expandString(v)
	}
	if v, ok := m["static_tunnel_session_id"]; ok {
		obj.StaticTunnelSessionID = *expandString(v)
	}
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	return &obj
}

func expandAuditEventEndpointSlice(in interface{}) *[]restapi.AuditEventEndpoint {
	var out []restapi.AuditEventEndpoint
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAuditEventEndpoint(v))
	}
	return &out
}

func flattenAuditEventTunnelSession(obj *restapi.AuditEventTunnelSession) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["remote_addr"] = obj.RemoteAddr
	m["created_at"] = obj.CreatedAt
	m["agent_version"] = obj.AgentVersion
	m["transport"] = obj.Transport
	m["os"] = obj.Os
	m["arch"] = obj.Arch
	m["region_id"] = obj.RegionID
	m["cred_id"] = obj.CredID
	m["ssh_cred_id"] = obj.SSHCredID
	m["agent_ingress_hostname"] = obj.AgentIngressHostname
	m["proxy_type"] = obj.ProxyType
	m["mutual_tls"] = obj.MutualTls
	m["service_run"] = obj.ServiceRun
	m["config_version"] = obj.ConfigVersion
	m["custom_cas"] = obj.CustomCas
	m["client_type"] = obj.ClientType
	m["user_agent"] = flattenTunnelSessionUserAgent(&obj.UserAgent)
	m["deprecated"] = flattenAgentVersionDeprecated(&obj.Deprecated)

	return []interface{}{m}
}

func flattenAuditEventTunnelSessionSlice(objs *[]restapi.AuditEventTunnelSession) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAuditEventTunnelSession(&v))
	}
	return sl
}

func expandAuditEventTunnelSession(in interface{}) *restapi.AuditEventTunnelSession {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AuditEventTunnelSession
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["remote_addr"]; ok {
		obj.RemoteAddr = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["agent_version"]; ok {
		obj.AgentVersion = *expandString(v)
	}
	if v, ok := m["transport"]; ok {
		obj.Transport = *expandString(v)
	}
	if v, ok := m["os"]; ok {
		obj.Os = *expandString(v)
	}
	if v, ok := m["arch"]; ok {
		obj.Arch = *expandString(v)
	}
	if v, ok := m["region_id"]; ok {
		obj.RegionID = *expandString(v)
	}
	if v, ok := m["cred_id"]; ok {
		obj.CredID = *expandString(v)
	}
	if v, ok := m["ssh_cred_id"]; ok {
		obj.SSHCredID = *expandString(v)
	}
	if v, ok := m["agent_ingress_hostname"]; ok {
		obj.AgentIngressHostname = *expandString(v)
	}
	if v, ok := m["proxy_type"]; ok {
		obj.ProxyType = *expandString(v)
	}
	if v, ok := m["mutual_tls"]; ok {
		obj.MutualTls = *expandBool(v)
	}
	if v, ok := m["service_run"]; ok {
		obj.ServiceRun = *expandBool(v)
	}
	if v, ok := m["config_version"]; ok {
		obj.ConfigVersion = *expandString(v)
	}
	if v, ok := m["custom_cas"]; ok {
		obj.CustomCas = *expandBool(v)
	}
	if v, ok := m["client_type"]; ok {
		obj.ClientType = *expandString(v)
	}
	if v, ok := m["user_agent"]; ok {
		obj.UserAgent = *expandTunnelSessionUserAgent(v)
	}
	if v, ok := m["deprecated"]; ok {
		obj.Deprecated = *expandAgentVersionDeprecated(v)
	}
	return &obj
}

func expandAuditEventTunnelSessionSlice(in interface{}) *[]restapi.AuditEventTunnelSession {
	var out []restapi.AuditEventTunnelSession
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAuditEventTunnelSession(v))
	}
	return &out
}

func flattenAuditEventTunnel(obj *restapi.AuditEventTunnel) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["created_at"] = obj.CreatedAt
	m["deleted_at"] = obj.DeletedAt
	m["account_id"] = obj.AccountID
	m["session_id"] = obj.SessionID
	m["remote_addr"] = obj.RemoteAddr
	m["region_id"] = obj.RegionID
	m["agent_version"] = obj.AgentVersion
	m["forwards_to"] = obj.ForwardsTo
	m["endpoint_id"] = obj.EndpointID
	m["labels"] = obj.Labels
	m["backend_ids"] = obj.BackendIDs
	m["hostname"] = obj.Hostname
	m["port"] = obj.Port
	m["proto"] = obj.Proto

	return []interface{}{m}
}

func flattenAuditEventTunnelSlice(objs *[]restapi.AuditEventTunnel) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAuditEventTunnel(&v))
	}
	return sl
}

func expandAuditEventTunnel(in interface{}) *restapi.AuditEventTunnel {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AuditEventTunnel
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["deleted_at"]; ok {
		obj.DeletedAt = *expandString(v)
	}
	if v, ok := m["account_id"]; ok {
		obj.AccountID = *expandString(v)
	}
	if v, ok := m["session_id"]; ok {
		obj.SessionID = *expandString(v)
	}
	if v, ok := m["remote_addr"]; ok {
		obj.RemoteAddr = *expandString(v)
	}
	if v, ok := m["region_id"]; ok {
		obj.RegionID = *expandString(v)
	}
	if v, ok := m["agent_version"]; ok {
		obj.AgentVersion = *expandString(v)
	}
	if v, ok := m["forwards_to"]; ok {
		obj.ForwardsTo = *expandString(v)
	}
	if v, ok := m["endpoint_id"]; ok {
		obj.EndpointID = *expandString(v)
	}
	if v, ok := m["labels"]; ok {
		obj.Labels = *expandStringMap(v)
	}
	if v, ok := m["backend_ids"]; ok {
		obj.BackendIDs = *expandStringSlice(v)
	}
	if v, ok := m["hostname"]; ok {
		obj.Hostname = *expandString(v)
	}
	if v, ok := m["port"]; ok {
		obj.Port = *expandInt64(v)
	}
	if v, ok := m["proto"]; ok {
		obj.Proto = *expandString(v)
	}
	return &obj
}

func expandAuditEventTunnelSlice(in interface{}) *[]restapi.AuditEventTunnel {
	var out []restapi.AuditEventTunnel
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAuditEventTunnel(v))
	}
	return &out
}

func flattenFailoverBackend(obj *restapi.FailoverBackend) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["backends"] = obj.Backends

	return []interface{}{m}
}

func flattenFailoverBackendSlice(objs *[]restapi.FailoverBackend) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenFailoverBackend(&v))
	}
	return sl
}

func expandFailoverBackend(in interface{}) *restapi.FailoverBackend {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.FailoverBackend
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandStringSlice(v)
	}
	return &obj
}

func expandFailoverBackendSlice(in interface{}) *[]restapi.FailoverBackend {
	var out []restapi.FailoverBackend
	for _, v := range in.([]interface{}) {
		out = append(out, *expandFailoverBackend(v))
	}
	return &out
}

func flattenFailoverBackendCreate(obj *restapi.FailoverBackendCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["backends"] = obj.Backends

	return []interface{}{m}
}

func flattenFailoverBackendCreateSlice(objs *[]restapi.FailoverBackendCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenFailoverBackendCreate(&v))
	}
	return sl
}

func expandFailoverBackendCreate(in interface{}) *restapi.FailoverBackendCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.FailoverBackendCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandStringSlice(v)
	}
	return &obj
}

func expandFailoverBackendCreateSlice(in interface{}) *[]restapi.FailoverBackendCreate {
	var out []restapi.FailoverBackendCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandFailoverBackendCreate(v))
	}
	return &out
}

func flattenFailoverBackendUpdate(obj *restapi.FailoverBackendUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["backends"] = obj.Backends

	return []interface{}{m}
}

func flattenFailoverBackendUpdateSlice(objs *[]restapi.FailoverBackendUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenFailoverBackendUpdate(&v))
	}
	return sl
}

func expandFailoverBackendUpdate(in interface{}) *restapi.FailoverBackendUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.FailoverBackendUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandStringSlice(v)
	}
	return &obj
}

func expandFailoverBackendUpdateSlice(in interface{}) *[]restapi.FailoverBackendUpdate {
	var out []restapi.FailoverBackendUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandFailoverBackendUpdate(v))
	}
	return &out
}

func flattenFailoverBackendList(obj *restapi.FailoverBackendList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["backends"] = flattenFailoverBackendSlice(&obj.Backends)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenFailoverBackendListSlice(objs *[]restapi.FailoverBackendList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenFailoverBackendList(&v))
	}
	return sl
}

func expandFailoverBackendList(in interface{}) *restapi.FailoverBackendList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.FailoverBackendList
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandFailoverBackendSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandFailoverBackendListSlice(in interface{}) *[]restapi.FailoverBackendList {
	var out []restapi.FailoverBackendList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandFailoverBackendList(v))
	}
	return &out
}

func flattenHTTPResponseBackend(obj *restapi.HTTPResponseBackend) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["body"] = obj.Body
	m["headers"] = obj.Headers
	m["status_code"] = obj.StatusCode

	return []interface{}{m}
}

func flattenHTTPResponseBackendSlice(objs *[]restapi.HTTPResponseBackend) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPResponseBackend(&v))
	}
	return sl
}

func expandHTTPResponseBackend(in interface{}) *restapi.HTTPResponseBackend {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPResponseBackend
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["body"]; ok {
		obj.Body = *expandString(v)
	}
	if v, ok := m["headers"]; ok {
		obj.Headers = *expandStringMap(v)
	}
	if v, ok := m["status_code"]; ok {
		obj.StatusCode = *expandInt32(v)
	}
	return &obj
}

func expandHTTPResponseBackendSlice(in interface{}) *[]restapi.HTTPResponseBackend {
	var out []restapi.HTTPResponseBackend
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPResponseBackend(v))
	}
	return &out
}

func flattenHTTPResponseBackendCreate(obj *restapi.HTTPResponseBackendCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["body"] = obj.Body
	m["headers"] = obj.Headers
	m["status_code"] = obj.StatusCode

	return []interface{}{m}
}

func flattenHTTPResponseBackendCreateSlice(objs *[]restapi.HTTPResponseBackendCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPResponseBackendCreate(&v))
	}
	return sl
}

func expandHTTPResponseBackendCreate(in interface{}) *restapi.HTTPResponseBackendCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPResponseBackendCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["body"]; ok {
		obj.Body = *expandString(v)
	}
	if v, ok := m["headers"]; ok {
		obj.Headers = *expandStringMap(v)
	}
	if v, ok := m["status_code"]; ok {
		obj.StatusCode = expandInt32(v)
	}
	return &obj
}

func expandHTTPResponseBackendCreateSlice(in interface{}) *[]restapi.HTTPResponseBackendCreate {
	var out []restapi.HTTPResponseBackendCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPResponseBackendCreate(v))
	}
	return &out
}

func flattenHTTPResponseBackendUpdate(obj *restapi.HTTPResponseBackendUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["body"] = obj.Body
	m["headers"] = obj.Headers
	m["status_code"] = obj.StatusCode

	return []interface{}{m}
}

func flattenHTTPResponseBackendUpdateSlice(objs *[]restapi.HTTPResponseBackendUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPResponseBackendUpdate(&v))
	}
	return sl
}

func expandHTTPResponseBackendUpdate(in interface{}) *restapi.HTTPResponseBackendUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPResponseBackendUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["body"]; ok {
		obj.Body = expandString(v)
	}
	if v, ok := m["headers"]; ok {
		obj.Headers = expandStringMap(v)
	}
	if v, ok := m["status_code"]; ok {
		obj.StatusCode = expandInt32(v)
	}
	return &obj
}

func expandHTTPResponseBackendUpdateSlice(in interface{}) *[]restapi.HTTPResponseBackendUpdate {
	var out []restapi.HTTPResponseBackendUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPResponseBackendUpdate(v))
	}
	return &out
}

func flattenHTTPResponseBackendList(obj *restapi.HTTPResponseBackendList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["backends"] = flattenHTTPResponseBackendSlice(&obj.Backends)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenHTTPResponseBackendListSlice(objs *[]restapi.HTTPResponseBackendList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPResponseBackendList(&v))
	}
	return sl
}

func expandHTTPResponseBackendList(in interface{}) *restapi.HTTPResponseBackendList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPResponseBackendList
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandHTTPResponseBackendSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandHTTPResponseBackendListSlice(in interface{}) *[]restapi.HTTPResponseBackendList {
	var out []restapi.HTTPResponseBackendList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPResponseBackendList(v))
	}
	return &out
}

func flattenStaticBackend(obj *restapi.StaticBackend) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["address"] = obj.Address
	m["tls"] = flattenStaticBackendTLS(&obj.TLS)

	return []interface{}{m}
}

func flattenStaticBackendSlice(objs *[]restapi.StaticBackend) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenStaticBackend(&v))
	}
	return sl
}

func expandStaticBackend(in interface{}) *restapi.StaticBackend {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.StaticBackend
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["address"]; ok {
		obj.Address = *expandString(v)
	}
	if v, ok := m["tls"]; ok {
		obj.TLS = *expandStaticBackendTLS(v)
	}
	return &obj
}

func expandStaticBackendSlice(in interface{}) *[]restapi.StaticBackend {
	var out []restapi.StaticBackend
	for _, v := range in.([]interface{}) {
		out = append(out, *expandStaticBackend(v))
	}
	return &out
}

func flattenStaticBackendTLS(obj *restapi.StaticBackendTLS) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled

	return []interface{}{m}
}

func flattenStaticBackendTLSSlice(objs *[]restapi.StaticBackendTLS) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenStaticBackendTLS(&v))
	}
	return sl
}

func expandStaticBackendTLS(in interface{}) *restapi.StaticBackendTLS {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.StaticBackendTLS
	if v, ok := m["enabled"]; ok {
		obj.Enabled = *expandBool(v)
	}
	return &obj
}

func expandStaticBackendTLSSlice(in interface{}) *[]restapi.StaticBackendTLS {
	var out []restapi.StaticBackendTLS
	for _, v := range in.([]interface{}) {
		out = append(out, *expandStaticBackendTLS(v))
	}
	return &out
}

func flattenStaticBackendCreate(obj *restapi.StaticBackendCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["address"] = obj.Address
	m["tls"] = flattenStaticBackendTLS(&obj.TLS)

	return []interface{}{m}
}

func flattenStaticBackendCreateSlice(objs *[]restapi.StaticBackendCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenStaticBackendCreate(&v))
	}
	return sl
}

func expandStaticBackendCreate(in interface{}) *restapi.StaticBackendCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.StaticBackendCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["address"]; ok {
		obj.Address = *expandString(v)
	}
	if v, ok := m["tls"]; ok {
		obj.TLS = *expandStaticBackendTLS(v)
	}
	return &obj
}

func expandStaticBackendCreateSlice(in interface{}) *[]restapi.StaticBackendCreate {
	var out []restapi.StaticBackendCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandStaticBackendCreate(v))
	}
	return &out
}

func flattenStaticBackendUpdate(obj *restapi.StaticBackendUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["address"] = obj.Address
	m["tls"] = flattenStaticBackendTLS(&obj.TLS)

	return []interface{}{m}
}

func flattenStaticBackendUpdateSlice(objs *[]restapi.StaticBackendUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenStaticBackendUpdate(&v))
	}
	return sl
}

func expandStaticBackendUpdate(in interface{}) *restapi.StaticBackendUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.StaticBackendUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["address"]; ok {
		obj.Address = *expandString(v)
	}
	if v, ok := m["tls"]; ok {
		obj.TLS = *expandStaticBackendTLS(v)
	}
	return &obj
}

func expandStaticBackendUpdateSlice(in interface{}) *[]restapi.StaticBackendUpdate {
	var out []restapi.StaticBackendUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandStaticBackendUpdate(v))
	}
	return &out
}

func flattenStaticBackendList(obj *restapi.StaticBackendList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["backends"] = flattenStaticBackendSlice(&obj.Backends)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenStaticBackendListSlice(objs *[]restapi.StaticBackendList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenStaticBackendList(&v))
	}
	return sl
}

func expandStaticBackendList(in interface{}) *restapi.StaticBackendList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.StaticBackendList
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandStaticBackendSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandStaticBackendListSlice(in interface{}) *[]restapi.StaticBackendList {
	var out []restapi.StaticBackendList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandStaticBackendList(v))
	}
	return &out
}

func flattenTunnelGroupBackend(obj *restapi.TunnelGroupBackend) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["labels"] = obj.Labels
	m["tunnels"] = flattenRefSlice(&obj.Tunnels)

	return []interface{}{m}
}

func flattenTunnelGroupBackendSlice(objs *[]restapi.TunnelGroupBackend) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnelGroupBackend(&v))
	}
	return sl
}

func expandTunnelGroupBackend(in interface{}) *restapi.TunnelGroupBackend {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TunnelGroupBackend
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["labels"]; ok {
		obj.Labels = *expandStringMap(v)
	}
	if v, ok := m["tunnels"]; ok {
		obj.Tunnels = *expandRefSlice(v)
	}
	return &obj
}

func expandTunnelGroupBackendSlice(in interface{}) *[]restapi.TunnelGroupBackend {
	var out []restapi.TunnelGroupBackend
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnelGroupBackend(v))
	}
	return &out
}

func flattenTunnelGroupBackendCreate(obj *restapi.TunnelGroupBackendCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["labels"] = obj.Labels

	return []interface{}{m}
}

func flattenTunnelGroupBackendCreateSlice(objs *[]restapi.TunnelGroupBackendCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnelGroupBackendCreate(&v))
	}
	return sl
}

func expandTunnelGroupBackendCreate(in interface{}) *restapi.TunnelGroupBackendCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TunnelGroupBackendCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["labels"]; ok {
		obj.Labels = *expandStringMap(v)
	}
	return &obj
}

func expandTunnelGroupBackendCreateSlice(in interface{}) *[]restapi.TunnelGroupBackendCreate {
	var out []restapi.TunnelGroupBackendCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnelGroupBackendCreate(v))
	}
	return &out
}

func flattenTunnelGroupBackendUpdate(obj *restapi.TunnelGroupBackendUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["labels"] = obj.Labels

	return []interface{}{m}
}

func flattenTunnelGroupBackendUpdateSlice(objs *[]restapi.TunnelGroupBackendUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnelGroupBackendUpdate(&v))
	}
	return sl
}

func expandTunnelGroupBackendUpdate(in interface{}) *restapi.TunnelGroupBackendUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TunnelGroupBackendUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["labels"]; ok {
		obj.Labels = *expandStringMap(v)
	}
	return &obj
}

func expandTunnelGroupBackendUpdateSlice(in interface{}) *[]restapi.TunnelGroupBackendUpdate {
	var out []restapi.TunnelGroupBackendUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnelGroupBackendUpdate(v))
	}
	return &out
}

func flattenTunnelGroupBackendList(obj *restapi.TunnelGroupBackendList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["backends"] = flattenTunnelGroupBackendSlice(&obj.Backends)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenTunnelGroupBackendListSlice(objs *[]restapi.TunnelGroupBackendList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnelGroupBackendList(&v))
	}
	return sl
}

func expandTunnelGroupBackendList(in interface{}) *restapi.TunnelGroupBackendList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TunnelGroupBackendList
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandTunnelGroupBackendSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandTunnelGroupBackendListSlice(in interface{}) *[]restapi.TunnelGroupBackendList {
	var out []restapi.TunnelGroupBackendList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnelGroupBackendList(v))
	}
	return &out
}

func flattenWeightedBackend(obj *restapi.WeightedBackend) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["backends"] = obj.Backends

	return []interface{}{m}
}

func flattenWeightedBackendSlice(objs *[]restapi.WeightedBackend) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenWeightedBackend(&v))
	}
	return sl
}

func expandWeightedBackend(in interface{}) *restapi.WeightedBackend {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.WeightedBackend
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandInt64Map(v)
	}
	return &obj
}

func expandWeightedBackendSlice(in interface{}) *[]restapi.WeightedBackend {
	var out []restapi.WeightedBackend
	for _, v := range in.([]interface{}) {
		out = append(out, *expandWeightedBackend(v))
	}
	return &out
}

func flattenWeightedBackendCreate(obj *restapi.WeightedBackendCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["backends"] = obj.Backends

	return []interface{}{m}
}

func flattenWeightedBackendCreateSlice(objs *[]restapi.WeightedBackendCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenWeightedBackendCreate(&v))
	}
	return sl
}

func expandWeightedBackendCreate(in interface{}) *restapi.WeightedBackendCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.WeightedBackendCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandInt64Map(v)
	}
	return &obj
}

func expandWeightedBackendCreateSlice(in interface{}) *[]restapi.WeightedBackendCreate {
	var out []restapi.WeightedBackendCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandWeightedBackendCreate(v))
	}
	return &out
}

func flattenWeightedBackendUpdate(obj *restapi.WeightedBackendUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["backends"] = obj.Backends

	return []interface{}{m}
}

func flattenWeightedBackendUpdateSlice(objs *[]restapi.WeightedBackendUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenWeightedBackendUpdate(&v))
	}
	return sl
}

func expandWeightedBackendUpdate(in interface{}) *restapi.WeightedBackendUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.WeightedBackendUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandInt64Map(v)
	}
	return &obj
}

func expandWeightedBackendUpdateSlice(in interface{}) *[]restapi.WeightedBackendUpdate {
	var out []restapi.WeightedBackendUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandWeightedBackendUpdate(v))
	}
	return &out
}

func flattenWeightedBackendList(obj *restapi.WeightedBackendList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["backends"] = flattenWeightedBackendSlice(&obj.Backends)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenWeightedBackendListSlice(objs *[]restapi.WeightedBackendList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenWeightedBackendList(&v))
	}
	return sl
}

func expandWeightedBackendList(in interface{}) *restapi.WeightedBackendList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.WeightedBackendList
	if v, ok := m["backends"]; ok {
		obj.Backends = *expandWeightedBackendSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandWeightedBackendListSlice(in interface{}) *[]restapi.WeightedBackendList {
	var out []restapi.WeightedBackendList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandWeightedBackendList(v))
	}
	return &out
}

func flattenBotUser(obj *restapi.BotUser) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["name"] = obj.Name
	m["active"] = obj.Active
	m["created_at"] = obj.CreatedAt

	return []interface{}{m}
}

func flattenBotUserSlice(objs *[]restapi.BotUser) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenBotUser(&v))
	}
	return sl
}

func expandBotUser(in interface{}) *restapi.BotUser {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.BotUser
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["name"]; ok {
		obj.Name = *expandString(v)
	}
	if v, ok := m["active"]; ok {
		obj.Active = *expandBool(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	return &obj
}

func expandBotUserSlice(in interface{}) *[]restapi.BotUser {
	var out []restapi.BotUser
	for _, v := range in.([]interface{}) {
		out = append(out, *expandBotUser(v))
	}
	return &out
}

func flattenBotUserCreate(obj *restapi.BotUserCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["name"] = obj.Name
	m["active"] = obj.Active

	return []interface{}{m}
}

func flattenBotUserCreateSlice(objs *[]restapi.BotUserCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenBotUserCreate(&v))
	}
	return sl
}

func expandBotUserCreate(in interface{}) *restapi.BotUserCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.BotUserCreate
	if v, ok := m["name"]; ok {
		obj.Name = *expandString(v)
	}
	if v, ok := m["active"]; ok {
		obj.Active = expandBool(v)
	}
	return &obj
}

func expandBotUserCreateSlice(in interface{}) *[]restapi.BotUserCreate {
	var out []restapi.BotUserCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandBotUserCreate(v))
	}
	return &out
}

func flattenBotUserUpdate(obj *restapi.BotUserUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["name"] = obj.Name
	m["active"] = obj.Active

	return []interface{}{m}
}

func flattenBotUserUpdateSlice(objs *[]restapi.BotUserUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenBotUserUpdate(&v))
	}
	return sl
}

func expandBotUserUpdate(in interface{}) *restapi.BotUserUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.BotUserUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["name"]; ok {
		obj.Name = expandString(v)
	}
	if v, ok := m["active"]; ok {
		obj.Active = expandBool(v)
	}
	return &obj
}

func expandBotUserUpdateSlice(in interface{}) *[]restapi.BotUserUpdate {
	var out []restapi.BotUserUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandBotUserUpdate(v))
	}
	return &out
}

func flattenBotUserList(obj *restapi.BotUserList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["bot_users"] = flattenBotUserSlice(&obj.BotUsers)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenBotUserListSlice(objs *[]restapi.BotUserList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenBotUserList(&v))
	}
	return sl
}

func expandBotUserList(in interface{}) *restapi.BotUserList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.BotUserList
	if v, ok := m["bot_users"]; ok {
		obj.BotUsers = *expandBotUserSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandBotUserListSlice(in interface{}) *[]restapi.BotUserList {
	var out []restapi.BotUserList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandBotUserList(v))
	}
	return &out
}

func flattenCertificateAuthorityCreate(obj *restapi.CertificateAuthorityCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["ca_pem"] = obj.CAPEM

	return []interface{}{m}
}

func flattenCertificateAuthorityCreateSlice(objs *[]restapi.CertificateAuthorityCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenCertificateAuthorityCreate(&v))
	}
	return sl
}

func expandCertificateAuthorityCreate(in interface{}) *restapi.CertificateAuthorityCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.CertificateAuthorityCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["ca_pem"]; ok {
		obj.CAPEM = *expandString(v)
	}
	return &obj
}

func expandCertificateAuthorityCreateSlice(in interface{}) *[]restapi.CertificateAuthorityCreate {
	var out []restapi.CertificateAuthorityCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandCertificateAuthorityCreate(v))
	}
	return &out
}

func flattenCertificateAuthorityUpdate(obj *restapi.CertificateAuthorityUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata

	return []interface{}{m}
}

func flattenCertificateAuthorityUpdateSlice(objs *[]restapi.CertificateAuthorityUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenCertificateAuthorityUpdate(&v))
	}
	return sl
}

func expandCertificateAuthorityUpdate(in interface{}) *restapi.CertificateAuthorityUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.CertificateAuthorityUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	return &obj
}

func expandCertificateAuthorityUpdateSlice(in interface{}) *[]restapi.CertificateAuthorityUpdate {
	var out []restapi.CertificateAuthorityUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandCertificateAuthorityUpdate(v))
	}
	return &out
}

func flattenCertificateAuthority(obj *restapi.CertificateAuthority) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["ca_pem"] = obj.CAPEM
	m["subject_common_name"] = obj.SubjectCommonName
	m["not_before"] = obj.NotBefore
	m["not_after"] = obj.NotAfter
	m["key_usages"] = obj.KeyUsages
	m["extended_key_usages"] = obj.ExtendedKeyUsages

	return []interface{}{m}
}

func flattenCertificateAuthoritySlice(objs *[]restapi.CertificateAuthority) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenCertificateAuthority(&v))
	}
	return sl
}

func expandCertificateAuthority(in interface{}) *restapi.CertificateAuthority {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.CertificateAuthority
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["ca_pem"]; ok {
		obj.CAPEM = *expandString(v)
	}
	if v, ok := m["subject_common_name"]; ok {
		obj.SubjectCommonName = *expandString(v)
	}
	if v, ok := m["not_before"]; ok {
		obj.NotBefore = *expandString(v)
	}
	if v, ok := m["not_after"]; ok {
		obj.NotAfter = *expandString(v)
	}
	if v, ok := m["key_usages"]; ok {
		obj.KeyUsages = *expandStringSlice(v)
	}
	if v, ok := m["extended_key_usages"]; ok {
		obj.ExtendedKeyUsages = *expandStringSlice(v)
	}
	return &obj
}

func expandCertificateAuthoritySlice(in interface{}) *[]restapi.CertificateAuthority {
	var out []restapi.CertificateAuthority
	for _, v := range in.([]interface{}) {
		out = append(out, *expandCertificateAuthority(v))
	}
	return &out
}

func flattenCertificateAuthorityList(obj *restapi.CertificateAuthorityList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["certificate_authorities"] = flattenCertificateAuthoritySlice(&obj.CertificateAuthorities)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenCertificateAuthorityListSlice(objs *[]restapi.CertificateAuthorityList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenCertificateAuthorityList(&v))
	}
	return sl
}

func expandCertificateAuthorityList(in interface{}) *restapi.CertificateAuthorityList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.CertificateAuthorityList
	if v, ok := m["certificate_authorities"]; ok {
		obj.CertificateAuthorities = *expandCertificateAuthoritySlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandCertificateAuthorityListSlice(in interface{}) *[]restapi.CertificateAuthorityList {
	var out []restapi.CertificateAuthorityList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandCertificateAuthorityList(v))
	}
	return &out
}

func flattenCredentialCreate(obj *restapi.CredentialCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["acl"] = obj.ACL
	m["owner_id"] = obj.OwnerID
	m["owner_email"] = obj.OwnerEmail
	m["precomputed_token"] = obj.PrecomputedToken

	return []interface{}{m}
}

func flattenCredentialCreateSlice(objs *[]restapi.CredentialCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenCredentialCreate(&v))
	}
	return sl
}

func expandCredentialCreate(in interface{}) *restapi.CredentialCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.CredentialCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["acl"]; ok {
		obj.ACL = *expandStringSlice(v)
	}
	if v, ok := m["owner_id"]; ok {
		obj.OwnerID = expandString(v)
	}
	if v, ok := m["owner_email"]; ok {
		obj.OwnerEmail = *expandString(v)
	}
	if v, ok := m["precomputed_token"]; ok {
		obj.PrecomputedToken = expandString(v)
	}
	return &obj
}

func expandCredentialCreateSlice(in interface{}) *[]restapi.CredentialCreate {
	var out []restapi.CredentialCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandCredentialCreate(v))
	}
	return &out
}

func flattenCredentialUpdate(obj *restapi.CredentialUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["acl"] = obj.ACL

	return []interface{}{m}
}

func flattenCredentialUpdateSlice(objs *[]restapi.CredentialUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenCredentialUpdate(&v))
	}
	return sl
}

func expandCredentialUpdate(in interface{}) *restapi.CredentialUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.CredentialUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["acl"]; ok {
		obj.ACL = expandStringSlice(v)
	}
	return &obj
}

func expandCredentialUpdateSlice(in interface{}) *[]restapi.CredentialUpdate {
	var out []restapi.CredentialUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandCredentialUpdate(v))
	}
	return &out
}

func flattenCredential(obj *restapi.Credential) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["token"] = obj.Token
	m["acl"] = obj.ACL
	m["owner_id"] = obj.OwnerID

	return []interface{}{m}
}

func flattenCredentialSlice(objs *[]restapi.Credential) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenCredential(&v))
	}
	return sl
}

func expandCredential(in interface{}) *restapi.Credential {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.Credential
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["token"]; ok {
		obj.Token = expandString(v)
	}
	if v, ok := m["acl"]; ok {
		obj.ACL = *expandStringSlice(v)
	}
	if v, ok := m["owner_id"]; ok {
		obj.OwnerID = expandString(v)
	}
	return &obj
}

func expandCredentialSlice(in interface{}) *[]restapi.Credential {
	var out []restapi.Credential
	for _, v := range in.([]interface{}) {
		out = append(out, *expandCredential(v))
	}
	return &out
}

func flattenCredentialList(obj *restapi.CredentialList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["credentials"] = flattenCredentialSlice(&obj.Credentials)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenCredentialListSlice(objs *[]restapi.CredentialList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenCredentialList(&v))
	}
	return sl
}

func expandCredentialList(in interface{}) *restapi.CredentialList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.CredentialList
	if v, ok := m["credentials"]; ok {
		obj.Credentials = *expandCredentialSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandCredentialListSlice(in interface{}) *[]restapi.CredentialList {
	var out []restapi.CredentialList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandCredentialList(v))
	}
	return &out
}

func flattenEndpointConfiguration(obj *restapi.EndpointConfiguration) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["type"] = obj.Type
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["created_at"] = obj.CreatedAt
	m["uri"] = obj.URI
	m["basic_auth"] = flattenEndpointBasicAuth(obj.BasicAuth)
	m["circuit_breaker"] = flattenEndpointCircuitBreaker(obj.CircuitBreaker)
	m["compression"] = flattenEndpointCompression(obj.Compression)
	m["request_headers"] = flattenEndpointRequestHeaders(obj.RequestHeaders)
	m["response_headers"] = flattenEndpointResponseHeaders(obj.ResponseHeaders)
	m["ip_policy"] = flattenEndpointIPPolicy(obj.IPPolicy)
	m["mutual_tls"] = flattenEndpointMutualTLS(obj.MutualTLS)
	m["tls_termination"] = flattenEndpointTLSTermination(obj.TLSTermination)
	m["webhook_validation"] = flattenEndpointWebhookValidation(obj.WebhookValidation)
	m["oauth"] = flattenEndpointOAuth(obj.OAuth)
	m["saml"] = flattenEndpointSAML(obj.SAML)
	m["oidc"] = flattenEndpointOIDC(obj.OIDC)
	m["backend"] = flattenEndpointBackend(obj.Backend)

	return []interface{}{m}
}

func flattenEndpointConfigurationSlice(objs *[]restapi.EndpointConfiguration) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointConfiguration(&v))
	}
	return sl
}

func expandEndpointConfiguration(in interface{}) *restapi.EndpointConfiguration {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointConfiguration
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["basic_auth"]; ok {
		obj.BasicAuth = expandEndpointBasicAuth(v)
	}
	if v, ok := m["circuit_breaker"]; ok {
		obj.CircuitBreaker = expandEndpointCircuitBreaker(v)
	}
	if v, ok := m["compression"]; ok {
		obj.Compression = expandEndpointCompression(v)
	}
	if v, ok := m["request_headers"]; ok {
		obj.RequestHeaders = expandEndpointRequestHeaders(v)
	}
	if v, ok := m["response_headers"]; ok {
		obj.ResponseHeaders = expandEndpointResponseHeaders(v)
	}
	if v, ok := m["ip_policy"]; ok {
		obj.IPPolicy = expandEndpointIPPolicy(v)
	}
	if v, ok := m["mutual_tls"]; ok {
		obj.MutualTLS = expandEndpointMutualTLS(v)
	}
	if v, ok := m["tls_termination"]; ok {
		obj.TLSTermination = expandEndpointTLSTermination(v)
	}
	if v, ok := m["webhook_validation"]; ok {
		obj.WebhookValidation = expandEndpointWebhookValidation(v)
	}
	if v, ok := m["oauth"]; ok {
		obj.OAuth = expandEndpointOAuth(v)
	}
	if v, ok := m["saml"]; ok {
		obj.SAML = expandEndpointSAML(v)
	}
	if v, ok := m["oidc"]; ok {
		obj.OIDC = expandEndpointOIDC(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackend(v)
	}
	return &obj
}

func expandEndpointConfigurationSlice(in interface{}) *[]restapi.EndpointConfiguration {
	var out []restapi.EndpointConfiguration
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointConfiguration(v))
	}
	return &out
}

func flattenEndpointConfigurationList(obj *restapi.EndpointConfigurationList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["endpoint_configurations"] = flattenEndpointConfigurationSlice(&obj.EndpointConfigurations)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenEndpointConfigurationListSlice(objs *[]restapi.EndpointConfigurationList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointConfigurationList(&v))
	}
	return sl
}

func expandEndpointConfigurationList(in interface{}) *restapi.EndpointConfigurationList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointConfigurationList
	if v, ok := m["endpoint_configurations"]; ok {
		obj.EndpointConfigurations = *expandEndpointConfigurationSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandEndpointConfigurationListSlice(in interface{}) *[]restapi.EndpointConfigurationList {
	var out []restapi.EndpointConfigurationList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointConfigurationList(v))
	}
	return &out
}

func flattenEndpointConfigurationUpdate(obj *restapi.EndpointConfigurationUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["basic_auth"] = flattenEndpointBasicAuth(obj.BasicAuth)
	m["circuit_breaker"] = flattenEndpointCircuitBreaker(obj.CircuitBreaker)
	m["compression"] = flattenEndpointCompression(obj.Compression)
	m["request_headers"] = flattenEndpointRequestHeaders(obj.RequestHeaders)
	m["response_headers"] = flattenEndpointResponseHeaders(obj.ResponseHeaders)
	m["ip_policy"] = flattenEndpointIPPolicyMutate(obj.IPPolicy)
	m["mutual_tls"] = flattenEndpointMutualTLSMutate(obj.MutualTLS)
	m["tls_termination"] = flattenEndpointTLSTermination(obj.TLSTermination)
	m["webhook_validation"] = flattenEndpointWebhookValidation(obj.WebhookValidation)
	m["oauth"] = flattenEndpointOAuth(obj.OAuth)
	m["saml"] = flattenEndpointSAMLMutate(obj.SAML)
	m["oidc"] = flattenEndpointOIDC(obj.OIDC)
	m["backend"] = flattenEndpointBackendMutate(obj.Backend)

	return []interface{}{m}
}

func flattenEndpointConfigurationUpdateSlice(objs *[]restapi.EndpointConfigurationUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointConfigurationUpdate(&v))
	}
	return sl
}

func expandEndpointConfigurationUpdate(in interface{}) *restapi.EndpointConfigurationUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointConfigurationUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["basic_auth"]; ok {
		obj.BasicAuth = expandEndpointBasicAuth(v)
	}
	if v, ok := m["circuit_breaker"]; ok {
		obj.CircuitBreaker = expandEndpointCircuitBreaker(v)
	}
	if v, ok := m["compression"]; ok {
		obj.Compression = expandEndpointCompression(v)
	}
	if v, ok := m["request_headers"]; ok {
		obj.RequestHeaders = expandEndpointRequestHeaders(v)
	}
	if v, ok := m["response_headers"]; ok {
		obj.ResponseHeaders = expandEndpointResponseHeaders(v)
	}
	if v, ok := m["ip_policy"]; ok {
		obj.IPPolicy = expandEndpointIPPolicyMutate(v)
	}
	if v, ok := m["mutual_tls"]; ok {
		obj.MutualTLS = expandEndpointMutualTLSMutate(v)
	}
	if v, ok := m["tls_termination"]; ok {
		obj.TLSTermination = expandEndpointTLSTermination(v)
	}
	if v, ok := m["webhook_validation"]; ok {
		obj.WebhookValidation = expandEndpointWebhookValidation(v)
	}
	if v, ok := m["oauth"]; ok {
		obj.OAuth = expandEndpointOAuth(v)
	}
	if v, ok := m["saml"]; ok {
		obj.SAML = expandEndpointSAMLMutate(v)
	}
	if v, ok := m["oidc"]; ok {
		obj.OIDC = expandEndpointOIDC(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackendMutate(v)
	}
	return &obj
}

func expandEndpointConfigurationUpdateSlice(in interface{}) *[]restapi.EndpointConfigurationUpdate {
	var out []restapi.EndpointConfigurationUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointConfigurationUpdate(v))
	}
	return &out
}

func flattenEndpointConfigurationCreate(obj *restapi.EndpointConfigurationCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["type"] = obj.Type
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["basic_auth"] = flattenEndpointBasicAuth(obj.BasicAuth)
	m["circuit_breaker"] = flattenEndpointCircuitBreaker(obj.CircuitBreaker)
	m["compression"] = flattenEndpointCompression(obj.Compression)
	m["request_headers"] = flattenEndpointRequestHeaders(obj.RequestHeaders)
	m["response_headers"] = flattenEndpointResponseHeaders(obj.ResponseHeaders)
	m["ip_policy"] = flattenEndpointIPPolicyMutate(obj.IPPolicy)
	m["mutual_tls"] = flattenEndpointMutualTLSMutate(obj.MutualTLS)
	m["tls_termination"] = flattenEndpointTLSTermination(obj.TLSTermination)
	m["webhook_validation"] = flattenEndpointWebhookValidation(obj.WebhookValidation)
	m["oauth"] = flattenEndpointOAuth(obj.OAuth)
	m["saml"] = flattenEndpointSAMLMutate(obj.SAML)
	m["oidc"] = flattenEndpointOIDC(obj.OIDC)
	m["backend"] = flattenEndpointBackendMutate(obj.Backend)

	return []interface{}{m}
}

func flattenEndpointConfigurationCreateSlice(objs *[]restapi.EndpointConfigurationCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointConfigurationCreate(&v))
	}
	return sl
}

func expandEndpointConfigurationCreate(in interface{}) *restapi.EndpointConfigurationCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointConfigurationCreate
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["basic_auth"]; ok {
		obj.BasicAuth = expandEndpointBasicAuth(v)
	}
	if v, ok := m["circuit_breaker"]; ok {
		obj.CircuitBreaker = expandEndpointCircuitBreaker(v)
	}
	if v, ok := m["compression"]; ok {
		obj.Compression = expandEndpointCompression(v)
	}
	if v, ok := m["request_headers"]; ok {
		obj.RequestHeaders = expandEndpointRequestHeaders(v)
	}
	if v, ok := m["response_headers"]; ok {
		obj.ResponseHeaders = expandEndpointResponseHeaders(v)
	}
	if v, ok := m["ip_policy"]; ok {
		obj.IPPolicy = expandEndpointIPPolicyMutate(v)
	}
	if v, ok := m["mutual_tls"]; ok {
		obj.MutualTLS = expandEndpointMutualTLSMutate(v)
	}
	if v, ok := m["tls_termination"]; ok {
		obj.TLSTermination = expandEndpointTLSTermination(v)
	}
	if v, ok := m["webhook_validation"]; ok {
		obj.WebhookValidation = expandEndpointWebhookValidation(v)
	}
	if v, ok := m["oauth"]; ok {
		obj.OAuth = expandEndpointOAuth(v)
	}
	if v, ok := m["saml"]; ok {
		obj.SAML = expandEndpointSAMLMutate(v)
	}
	if v, ok := m["oidc"]; ok {
		obj.OIDC = expandEndpointOIDC(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackendMutate(v)
	}
	return &obj
}

func expandEndpointConfigurationCreateSlice(in interface{}) *[]restapi.EndpointConfigurationCreate {
	var out []restapi.EndpointConfigurationCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointConfigurationCreate(v))
	}
	return &out
}

func flattenEndpointWebhookValidation(obj *restapi.EndpointWebhookValidation) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["provider"] = obj.Provider
	m["secret"] = obj.Secret

	return []interface{}{m}
}

func flattenEndpointWebhookValidationSlice(objs *[]restapi.EndpointWebhookValidation) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointWebhookValidation(&v))
	}
	return sl
}

func expandEndpointWebhookValidation(in interface{}) *restapi.EndpointWebhookValidation {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointWebhookValidation
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["provider"]; ok {
		obj.Provider = *expandString(v)
	}
	if v, ok := m["secret"]; ok {
		obj.Secret = *expandString(v)
	}
	return &obj
}

func expandEndpointWebhookValidationSlice(in interface{}) *[]restapi.EndpointWebhookValidation {
	var out []restapi.EndpointWebhookValidation
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointWebhookValidation(v))
	}
	return &out
}

func flattenEndpointCompression(obj *restapi.EndpointCompression) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled

	return []interface{}{m}
}

func flattenEndpointCompressionSlice(objs *[]restapi.EndpointCompression) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointCompression(&v))
	}
	return sl
}

func expandEndpointCompression(in interface{}) *restapi.EndpointCompression {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointCompression
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	return &obj
}

func expandEndpointCompressionSlice(in interface{}) *[]restapi.EndpointCompression {
	var out []restapi.EndpointCompression
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointCompression(v))
	}
	return &out
}

func flattenEndpointMutualTLS(obj *restapi.EndpointMutualTLS) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["certificate_authorities"] = flattenRefSlice(&obj.CertificateAuthorities)

	return []interface{}{m}
}

func flattenEndpointMutualTLSSlice(objs *[]restapi.EndpointMutualTLS) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointMutualTLS(&v))
	}
	return sl
}

func expandEndpointMutualTLS(in interface{}) *restapi.EndpointMutualTLS {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointMutualTLS
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["certificate_authorities"]; ok {
		obj.CertificateAuthorities = *expandRefSlice(v)
	}
	return &obj
}

func expandEndpointMutualTLSSlice(in interface{}) *[]restapi.EndpointMutualTLS {
	var out []restapi.EndpointMutualTLS
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointMutualTLS(v))
	}
	return &out
}

func flattenEndpointMutualTLSMutate(obj *restapi.EndpointMutualTLSMutate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["certificate_authority_ids"] = obj.CertificateAuthorityIDs

	return []interface{}{m}
}

func flattenEndpointMutualTLSMutateSlice(objs *[]restapi.EndpointMutualTLSMutate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointMutualTLSMutate(&v))
	}
	return sl
}

func expandEndpointMutualTLSMutate(in interface{}) *restapi.EndpointMutualTLSMutate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointMutualTLSMutate
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["certificate_authority_ids"]; ok {
		obj.CertificateAuthorityIDs = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointMutualTLSMutateSlice(in interface{}) *[]restapi.EndpointMutualTLSMutate {
	var out []restapi.EndpointMutualTLSMutate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointMutualTLSMutate(v))
	}
	return &out
}

func flattenEndpointTLSTermination(obj *restapi.EndpointTLSTermination) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["terminate_at"] = obj.TerminateAt
	m["min_version"] = obj.MinVersion

	return []interface{}{m}
}

func flattenEndpointTLSTerminationSlice(objs *[]restapi.EndpointTLSTermination) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointTLSTermination(&v))
	}
	return sl
}

func expandEndpointTLSTermination(in interface{}) *restapi.EndpointTLSTermination {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointTLSTermination
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["terminate_at"]; ok {
		obj.TerminateAt = *expandString(v)
	}
	if v, ok := m["min_version"]; ok {
		obj.MinVersion = expandString(v)
	}
	return &obj
}

func expandEndpointTLSTerminationSlice(in interface{}) *[]restapi.EndpointTLSTermination {
	var out []restapi.EndpointTLSTermination
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointTLSTermination(v))
	}
	return &out
}

func flattenEndpointTLSTerminationAtEdge(obj *restapi.EndpointTLSTerminationAtEdge) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["min_version"] = obj.MinVersion

	return []interface{}{m}
}

func flattenEndpointTLSTerminationAtEdgeSlice(objs *[]restapi.EndpointTLSTerminationAtEdge) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointTLSTerminationAtEdge(&v))
	}
	return sl
}

func expandEndpointTLSTerminationAtEdge(in interface{}) *restapi.EndpointTLSTerminationAtEdge {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointTLSTerminationAtEdge
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["min_version"]; ok {
		obj.MinVersion = expandString(v)
	}
	return &obj
}

func expandEndpointTLSTerminationAtEdgeSlice(in interface{}) *[]restapi.EndpointTLSTerminationAtEdge {
	var out []restapi.EndpointTLSTerminationAtEdge
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointTLSTerminationAtEdge(v))
	}
	return &out
}

func flattenEndpointBasicAuth(obj *restapi.EndpointBasicAuth) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["auth_provider_id"] = obj.AuthProviderID
	m["realm"] = obj.Realm
	m["allow_options"] = obj.AllowOptions

	return []interface{}{m}
}

func flattenEndpointBasicAuthSlice(objs *[]restapi.EndpointBasicAuth) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointBasicAuth(&v))
	}
	return sl
}

func expandEndpointBasicAuth(in interface{}) *restapi.EndpointBasicAuth {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointBasicAuth
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["auth_provider_id"]; ok {
		obj.AuthProviderID = *expandString(v)
	}
	if v, ok := m["realm"]; ok {
		obj.Realm = *expandString(v)
	}
	if v, ok := m["allow_options"]; ok {
		obj.AllowOptions = *expandBool(v)
	}
	return &obj
}

func expandEndpointBasicAuthSlice(in interface{}) *[]restapi.EndpointBasicAuth {
	var out []restapi.EndpointBasicAuth
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointBasicAuth(v))
	}
	return &out
}

func flattenEndpointRequestHeaders(obj *restapi.EndpointRequestHeaders) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["add"] = obj.Add
	m["remove"] = obj.Remove

	return []interface{}{m}
}

func flattenEndpointRequestHeadersSlice(objs *[]restapi.EndpointRequestHeaders) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointRequestHeaders(&v))
	}
	return sl
}

func expandEndpointRequestHeaders(in interface{}) *restapi.EndpointRequestHeaders {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointRequestHeaders
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["add"]; ok {
		obj.Add = *expandStringMap(v)
	}
	if v, ok := m["remove"]; ok {
		obj.Remove = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointRequestHeadersSlice(in interface{}) *[]restapi.EndpointRequestHeaders {
	var out []restapi.EndpointRequestHeaders
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointRequestHeaders(v))
	}
	return &out
}

func flattenEndpointResponseHeaders(obj *restapi.EndpointResponseHeaders) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["add"] = obj.Add
	m["remove"] = obj.Remove

	return []interface{}{m}
}

func flattenEndpointResponseHeadersSlice(objs *[]restapi.EndpointResponseHeaders) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointResponseHeaders(&v))
	}
	return sl
}

func expandEndpointResponseHeaders(in interface{}) *restapi.EndpointResponseHeaders {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointResponseHeaders
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["add"]; ok {
		obj.Add = *expandStringMap(v)
	}
	if v, ok := m["remove"]; ok {
		obj.Remove = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointResponseHeadersSlice(in interface{}) *[]restapi.EndpointResponseHeaders {
	var out []restapi.EndpointResponseHeaders
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointResponseHeaders(v))
	}
	return &out
}

func flattenEndpointIPPolicy(obj *restapi.EndpointIPPolicy) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["ip_policies"] = flattenRefSlice(&obj.IPPolicies)

	return []interface{}{m}
}

func flattenEndpointIPPolicySlice(objs *[]restapi.EndpointIPPolicy) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointIPPolicy(&v))
	}
	return sl
}

func expandEndpointIPPolicy(in interface{}) *restapi.EndpointIPPolicy {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointIPPolicy
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["ip_policies"]; ok {
		obj.IPPolicies = *expandRefSlice(v)
	}
	return &obj
}

func expandEndpointIPPolicySlice(in interface{}) *[]restapi.EndpointIPPolicy {
	var out []restapi.EndpointIPPolicy
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointIPPolicy(v))
	}
	return &out
}

func flattenEndpointIPPolicyMutate(obj *restapi.EndpointIPPolicyMutate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["ip_policy_ids"] = obj.IPPolicyIDs

	return []interface{}{m}
}

func flattenEndpointIPPolicyMutateSlice(objs *[]restapi.EndpointIPPolicyMutate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointIPPolicyMutate(&v))
	}
	return sl
}

func expandEndpointIPPolicyMutate(in interface{}) *restapi.EndpointIPPolicyMutate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointIPPolicyMutate
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["ip_policy_ids"]; ok {
		obj.IPPolicyIDs = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointIPPolicyMutateSlice(in interface{}) *[]restapi.EndpointIPPolicyMutate {
	var out []restapi.EndpointIPPolicyMutate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointIPPolicyMutate(v))
	}
	return &out
}

func flattenEndpointCircuitBreaker(obj *restapi.EndpointCircuitBreaker) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["tripped_duration"] = obj.TrippedDuration
	m["rolling_window"] = obj.RollingWindow
	m["num_buckets"] = obj.NumBuckets
	m["volume_threshold"] = obj.VolumeThreshold
	m["error_threshold_percentage"] = obj.ErrorThresholdPercentage

	return []interface{}{m}
}

func flattenEndpointCircuitBreakerSlice(objs *[]restapi.EndpointCircuitBreaker) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointCircuitBreaker(&v))
	}
	return sl
}

func expandEndpointCircuitBreaker(in interface{}) *restapi.EndpointCircuitBreaker {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointCircuitBreaker
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["tripped_duration"]; ok {
		obj.TrippedDuration = *expandUint32(v)
	}
	if v, ok := m["rolling_window"]; ok {
		obj.RollingWindow = *expandUint32(v)
	}
	if v, ok := m["num_buckets"]; ok {
		obj.NumBuckets = *expandUint32(v)
	}
	if v, ok := m["volume_threshold"]; ok {
		obj.VolumeThreshold = *expandUint32(v)
	}
	if v, ok := m["error_threshold_percentage"]; ok {
		obj.ErrorThresholdPercentage = *expandFloat64(v)
	}
	return &obj
}

func expandEndpointCircuitBreakerSlice(in interface{}) *[]restapi.EndpointCircuitBreaker {
	var out []restapi.EndpointCircuitBreaker
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointCircuitBreaker(v))
	}
	return &out
}

func flattenEndpointOAuth(obj *restapi.EndpointOAuth) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["provider"] = flattenEndpointOAuthProvider(&obj.Provider)
	m["options_passthrough"] = obj.OptionsPassthrough
	m["cookie_prefix"] = obj.CookiePrefix
	m["inactivity_timeout"] = obj.InactivityTimeout
	m["maximum_duration"] = obj.MaximumDuration
	m["auth_check_interval"] = obj.AuthCheckInterval

	return []interface{}{m}
}

func flattenEndpointOAuthSlice(objs *[]restapi.EndpointOAuth) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuth(&v))
	}
	return sl
}

func expandEndpointOAuth(in interface{}) *restapi.EndpointOAuth {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuth
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["provider"]; ok {
		obj.Provider = *expandEndpointOAuthProvider(v)
	}
	if v, ok := m["options_passthrough"]; ok {
		obj.OptionsPassthrough = *expandBool(v)
	}
	if v, ok := m["cookie_prefix"]; ok {
		obj.CookiePrefix = *expandString(v)
	}
	if v, ok := m["inactivity_timeout"]; ok {
		obj.InactivityTimeout = *expandUint32(v)
	}
	if v, ok := m["maximum_duration"]; ok {
		obj.MaximumDuration = *expandUint32(v)
	}
	if v, ok := m["auth_check_interval"]; ok {
		obj.AuthCheckInterval = *expandUint32(v)
	}
	return &obj
}

func expandEndpointOAuthSlice(in interface{}) *[]restapi.EndpointOAuth {
	var out []restapi.EndpointOAuth
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuth(v))
	}
	return &out
}

func flattenEndpointOAuthProvider(obj *restapi.EndpointOAuthProvider) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["github"] = flattenEndpointOAuthGitHub(obj.Github)
	m["facebook"] = flattenEndpointOAuthFacebook(obj.Facebook)
	m["microsoft"] = flattenEndpointOAuthMicrosoft(obj.Microsoft)
	m["google"] = flattenEndpointOAuthGoogle(obj.Google)
	m["linkedin"] = flattenEndpointOAuthLinkedIn(obj.Linkedin)
	m["gitlab"] = flattenEndpointOAuthGitLab(obj.Gitlab)
	m["twitch"] = flattenEndpointOAuthTwitch(obj.Twitch)
	m["amazon"] = flattenEndpointOAuthAmazon(obj.Amazon)

	return []interface{}{m}
}

func flattenEndpointOAuthProviderSlice(objs *[]restapi.EndpointOAuthProvider) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuthProvider(&v))
	}
	return sl
}

func expandEndpointOAuthProvider(in interface{}) *restapi.EndpointOAuthProvider {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuthProvider
	if v, ok := m["github"]; ok {
		obj.Github = expandEndpointOAuthGitHub(v)
	}
	if v, ok := m["facebook"]; ok {
		obj.Facebook = expandEndpointOAuthFacebook(v)
	}
	if v, ok := m["microsoft"]; ok {
		obj.Microsoft = expandEndpointOAuthMicrosoft(v)
	}
	if v, ok := m["google"]; ok {
		obj.Google = expandEndpointOAuthGoogle(v)
	}
	if v, ok := m["linkedin"]; ok {
		obj.Linkedin = expandEndpointOAuthLinkedIn(v)
	}
	if v, ok := m["gitlab"]; ok {
		obj.Gitlab = expandEndpointOAuthGitLab(v)
	}
	if v, ok := m["twitch"]; ok {
		obj.Twitch = expandEndpointOAuthTwitch(v)
	}
	if v, ok := m["amazon"]; ok {
		obj.Amazon = expandEndpointOAuthAmazon(v)
	}
	return &obj
}

func expandEndpointOAuthProviderSlice(in interface{}) *[]restapi.EndpointOAuthProvider {
	var out []restapi.EndpointOAuthProvider
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuthProvider(v))
	}
	return &out
}

func flattenEndpointOAuthGitHub(obj *restapi.EndpointOAuthGitHub) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["client_id"] = obj.ClientID
	m["client_secret"] = obj.ClientSecret
	m["scopes"] = obj.Scopes
	m["email_addresses"] = obj.EmailAddresses
	m["email_domains"] = obj.EmailDomains
	m["teams"] = obj.Teams
	m["organizations"] = obj.Organizations

	return []interface{}{m}
}

func flattenEndpointOAuthGitHubSlice(objs *[]restapi.EndpointOAuthGitHub) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuthGitHub(&v))
	}
	return sl
}

func expandEndpointOAuthGitHub(in interface{}) *restapi.EndpointOAuthGitHub {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuthGitHub
	if v, ok := m["client_id"]; ok {
		obj.ClientID = expandString(v)
	}
	if v, ok := m["client_secret"]; ok {
		obj.ClientSecret = expandString(v)
	}
	if v, ok := m["scopes"]; ok {
		obj.Scopes = expandStringSlice(v)
	}
	if v, ok := m["email_addresses"]; ok {
		obj.EmailAddresses = expandStringSlice(v)
	}
	if v, ok := m["email_domains"]; ok {
		obj.EmailDomains = expandStringSlice(v)
	}
	if v, ok := m["teams"]; ok {
		obj.Teams = expandStringSlice(v)
	}
	if v, ok := m["organizations"]; ok {
		obj.Organizations = expandStringSlice(v)
	}
	return &obj
}

func expandEndpointOAuthGitHubSlice(in interface{}) *[]restapi.EndpointOAuthGitHub {
	var out []restapi.EndpointOAuthGitHub
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuthGitHub(v))
	}
	return &out
}

func flattenEndpointOAuthFacebook(obj *restapi.EndpointOAuthFacebook) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["client_id"] = obj.ClientID
	m["client_secret"] = obj.ClientSecret
	m["scopes"] = obj.Scopes
	m["email_addresses"] = obj.EmailAddresses
	m["email_domains"] = obj.EmailDomains

	return []interface{}{m}
}

func flattenEndpointOAuthFacebookSlice(objs *[]restapi.EndpointOAuthFacebook) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuthFacebook(&v))
	}
	return sl
}

func expandEndpointOAuthFacebook(in interface{}) *restapi.EndpointOAuthFacebook {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuthFacebook
	if v, ok := m["client_id"]; ok {
		obj.ClientID = expandString(v)
	}
	if v, ok := m["client_secret"]; ok {
		obj.ClientSecret = expandString(v)
	}
	if v, ok := m["scopes"]; ok {
		obj.Scopes = *expandStringSlice(v)
	}
	if v, ok := m["email_addresses"]; ok {
		obj.EmailAddresses = *expandStringSlice(v)
	}
	if v, ok := m["email_domains"]; ok {
		obj.EmailDomains = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointOAuthFacebookSlice(in interface{}) *[]restapi.EndpointOAuthFacebook {
	var out []restapi.EndpointOAuthFacebook
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuthFacebook(v))
	}
	return &out
}

func flattenEndpointOAuthMicrosoft(obj *restapi.EndpointOAuthMicrosoft) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["client_id"] = obj.ClientID
	m["client_secret"] = obj.ClientSecret
	m["scopes"] = obj.Scopes
	m["email_addresses"] = obj.EmailAddresses
	m["email_domains"] = obj.EmailDomains

	return []interface{}{m}
}

func flattenEndpointOAuthMicrosoftSlice(objs *[]restapi.EndpointOAuthMicrosoft) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuthMicrosoft(&v))
	}
	return sl
}

func expandEndpointOAuthMicrosoft(in interface{}) *restapi.EndpointOAuthMicrosoft {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuthMicrosoft
	if v, ok := m["client_id"]; ok {
		obj.ClientID = expandString(v)
	}
	if v, ok := m["client_secret"]; ok {
		obj.ClientSecret = expandString(v)
	}
	if v, ok := m["scopes"]; ok {
		obj.Scopes = *expandStringSlice(v)
	}
	if v, ok := m["email_addresses"]; ok {
		obj.EmailAddresses = *expandStringSlice(v)
	}
	if v, ok := m["email_domains"]; ok {
		obj.EmailDomains = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointOAuthMicrosoftSlice(in interface{}) *[]restapi.EndpointOAuthMicrosoft {
	var out []restapi.EndpointOAuthMicrosoft
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuthMicrosoft(v))
	}
	return &out
}

func flattenEndpointOAuthGoogle(obj *restapi.EndpointOAuthGoogle) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["client_id"] = obj.ClientID
	m["client_secret"] = obj.ClientSecret
	m["scopes"] = obj.Scopes
	m["email_addresses"] = obj.EmailAddresses
	m["email_domains"] = obj.EmailDomains

	return []interface{}{m}
}

func flattenEndpointOAuthGoogleSlice(objs *[]restapi.EndpointOAuthGoogle) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuthGoogle(&v))
	}
	return sl
}

func expandEndpointOAuthGoogle(in interface{}) *restapi.EndpointOAuthGoogle {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuthGoogle
	if v, ok := m["client_id"]; ok {
		obj.ClientID = expandString(v)
	}
	if v, ok := m["client_secret"]; ok {
		obj.ClientSecret = expandString(v)
	}
	if v, ok := m["scopes"]; ok {
		obj.Scopes = *expandStringSlice(v)
	}
	if v, ok := m["email_addresses"]; ok {
		obj.EmailAddresses = *expandStringSlice(v)
	}
	if v, ok := m["email_domains"]; ok {
		obj.EmailDomains = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointOAuthGoogleSlice(in interface{}) *[]restapi.EndpointOAuthGoogle {
	var out []restapi.EndpointOAuthGoogle
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuthGoogle(v))
	}
	return &out
}

func flattenEndpointOAuthLinkedIn(obj *restapi.EndpointOAuthLinkedIn) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["client_id"] = obj.ClientID
	m["client_secret"] = obj.ClientSecret
	m["scopes"] = obj.Scopes
	m["email_addresses"] = obj.EmailAddresses
	m["email_domains"] = obj.EmailDomains

	return []interface{}{m}
}

func flattenEndpointOAuthLinkedInSlice(objs *[]restapi.EndpointOAuthLinkedIn) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuthLinkedIn(&v))
	}
	return sl
}

func expandEndpointOAuthLinkedIn(in interface{}) *restapi.EndpointOAuthLinkedIn {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuthLinkedIn
	if v, ok := m["client_id"]; ok {
		obj.ClientID = expandString(v)
	}
	if v, ok := m["client_secret"]; ok {
		obj.ClientSecret = expandString(v)
	}
	if v, ok := m["scopes"]; ok {
		obj.Scopes = *expandStringSlice(v)
	}
	if v, ok := m["email_addresses"]; ok {
		obj.EmailAddresses = *expandStringSlice(v)
	}
	if v, ok := m["email_domains"]; ok {
		obj.EmailDomains = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointOAuthLinkedInSlice(in interface{}) *[]restapi.EndpointOAuthLinkedIn {
	var out []restapi.EndpointOAuthLinkedIn
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuthLinkedIn(v))
	}
	return &out
}

func flattenEndpointOAuthGitLab(obj *restapi.EndpointOAuthGitLab) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["client_id"] = obj.ClientID
	m["client_secret"] = obj.ClientSecret
	m["scopes"] = obj.Scopes
	m["email_addresses"] = obj.EmailAddresses
	m["email_domains"] = obj.EmailDomains

	return []interface{}{m}
}

func flattenEndpointOAuthGitLabSlice(objs *[]restapi.EndpointOAuthGitLab) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuthGitLab(&v))
	}
	return sl
}

func expandEndpointOAuthGitLab(in interface{}) *restapi.EndpointOAuthGitLab {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuthGitLab
	if v, ok := m["client_id"]; ok {
		obj.ClientID = expandString(v)
	}
	if v, ok := m["client_secret"]; ok {
		obj.ClientSecret = expandString(v)
	}
	if v, ok := m["scopes"]; ok {
		obj.Scopes = *expandStringSlice(v)
	}
	if v, ok := m["email_addresses"]; ok {
		obj.EmailAddresses = *expandStringSlice(v)
	}
	if v, ok := m["email_domains"]; ok {
		obj.EmailDomains = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointOAuthGitLabSlice(in interface{}) *[]restapi.EndpointOAuthGitLab {
	var out []restapi.EndpointOAuthGitLab
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuthGitLab(v))
	}
	return &out
}

func flattenEndpointOAuthTwitch(obj *restapi.EndpointOAuthTwitch) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["client_id"] = obj.ClientID
	m["client_secret"] = obj.ClientSecret
	m["scopes"] = obj.Scopes
	m["email_addresses"] = obj.EmailAddresses
	m["email_domains"] = obj.EmailDomains

	return []interface{}{m}
}

func flattenEndpointOAuthTwitchSlice(objs *[]restapi.EndpointOAuthTwitch) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuthTwitch(&v))
	}
	return sl
}

func expandEndpointOAuthTwitch(in interface{}) *restapi.EndpointOAuthTwitch {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuthTwitch
	if v, ok := m["client_id"]; ok {
		obj.ClientID = expandString(v)
	}
	if v, ok := m["client_secret"]; ok {
		obj.ClientSecret = expandString(v)
	}
	if v, ok := m["scopes"]; ok {
		obj.Scopes = *expandStringSlice(v)
	}
	if v, ok := m["email_addresses"]; ok {
		obj.EmailAddresses = *expandStringSlice(v)
	}
	if v, ok := m["email_domains"]; ok {
		obj.EmailDomains = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointOAuthTwitchSlice(in interface{}) *[]restapi.EndpointOAuthTwitch {
	var out []restapi.EndpointOAuthTwitch
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuthTwitch(v))
	}
	return &out
}

func flattenEndpointOAuthAmazon(obj *restapi.EndpointOAuthAmazon) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["client_id"] = obj.ClientID
	m["client_secret"] = obj.ClientSecret
	m["scopes"] = obj.Scopes
	m["email_addresses"] = obj.EmailAddresses
	m["email_domains"] = obj.EmailDomains

	return []interface{}{m}
}

func flattenEndpointOAuthAmazonSlice(objs *[]restapi.EndpointOAuthAmazon) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuthAmazon(&v))
	}
	return sl
}

func expandEndpointOAuthAmazon(in interface{}) *restapi.EndpointOAuthAmazon {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuthAmazon
	if v, ok := m["client_id"]; ok {
		obj.ClientID = expandString(v)
	}
	if v, ok := m["client_secret"]; ok {
		obj.ClientSecret = expandString(v)
	}
	if v, ok := m["scopes"]; ok {
		obj.Scopes = *expandStringSlice(v)
	}
	if v, ok := m["email_addresses"]; ok {
		obj.EmailAddresses = *expandStringSlice(v)
	}
	if v, ok := m["email_domains"]; ok {
		obj.EmailDomains = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointOAuthAmazonSlice(in interface{}) *[]restapi.EndpointOAuthAmazon {
	var out []restapi.EndpointOAuthAmazon
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuthAmazon(v))
	}
	return &out
}

func flattenEndpointSAML(obj *restapi.EndpointSAML) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["options_passthrough"] = obj.OptionsPassthrough
	m["cookie_prefix"] = obj.CookiePrefix
	m["inactivity_timeout"] = obj.InactivityTimeout
	m["maximum_duration"] = obj.MaximumDuration
	m["idp_metadata_url"] = obj.IdPMetadataURL
	m["idp_metadata"] = obj.IdPMetadata
	m["force_authn"] = obj.ForceAuthn
	m["allow_idp_initiated"] = obj.AllowIdPInitiated
	m["authorized_groups"] = obj.AuthorizedGroups
	m["entity_id"] = obj.EntityID
	m["assertion_consumer_service_url"] = obj.AssertionConsumerServiceURL
	m["single_logout_url"] = obj.SingleLogoutURL
	m["request_signing_certificate_pem"] = obj.RequestSigningCertificatePEM
	m["metadata_url"] = obj.MetadataURL
	m["nameid_format"] = obj.NameIDFormat

	return []interface{}{m}
}

func flattenEndpointSAMLSlice(objs *[]restapi.EndpointSAML) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointSAML(&v))
	}
	return sl
}

func expandEndpointSAML(in interface{}) *restapi.EndpointSAML {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointSAML
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["options_passthrough"]; ok {
		obj.OptionsPassthrough = *expandBool(v)
	}
	if v, ok := m["cookie_prefix"]; ok {
		obj.CookiePrefix = *expandString(v)
	}
	if v, ok := m["inactivity_timeout"]; ok {
		obj.InactivityTimeout = *expandUint32(v)
	}
	if v, ok := m["maximum_duration"]; ok {
		obj.MaximumDuration = *expandUint32(v)
	}
	if v, ok := m["idp_metadata_url"]; ok {
		obj.IdPMetadataURL = *expandString(v)
	}
	if v, ok := m["idp_metadata"]; ok {
		obj.IdPMetadata = *expandString(v)
	}
	if v, ok := m["force_authn"]; ok {
		obj.ForceAuthn = *expandBool(v)
	}
	if v, ok := m["allow_idp_initiated"]; ok {
		obj.AllowIdPInitiated = expandBool(v)
	}
	if v, ok := m["authorized_groups"]; ok {
		obj.AuthorizedGroups = *expandStringSlice(v)
	}
	if v, ok := m["entity_id"]; ok {
		obj.EntityID = *expandString(v)
	}
	if v, ok := m["assertion_consumer_service_url"]; ok {
		obj.AssertionConsumerServiceURL = *expandString(v)
	}
	if v, ok := m["single_logout_url"]; ok {
		obj.SingleLogoutURL = *expandString(v)
	}
	if v, ok := m["request_signing_certificate_pem"]; ok {
		obj.RequestSigningCertificatePEM = *expandString(v)
	}
	if v, ok := m["metadata_url"]; ok {
		obj.MetadataURL = *expandString(v)
	}
	if v, ok := m["nameid_format"]; ok {
		obj.NameIDFormat = *expandString(v)
	}
	return &obj
}

func expandEndpointSAMLSlice(in interface{}) *[]restapi.EndpointSAML {
	var out []restapi.EndpointSAML
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointSAML(v))
	}
	return &out
}

func flattenEndpointSAMLMutate(obj *restapi.EndpointSAMLMutate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["options_passthrough"] = obj.OptionsPassthrough
	m["cookie_prefix"] = obj.CookiePrefix
	m["inactivity_timeout"] = obj.InactivityTimeout
	m["maximum_duration"] = obj.MaximumDuration
	m["idp_metadata_url"] = obj.IdPMetadataURL
	m["idp_metadata"] = obj.IdPMetadata
	m["force_authn"] = obj.ForceAuthn
	m["allow_idp_initiated"] = obj.AllowIdPInitiated
	m["authorized_groups"] = obj.AuthorizedGroups
	m["nameid_format"] = obj.NameIDFormat

	return []interface{}{m}
}

func flattenEndpointSAMLMutateSlice(objs *[]restapi.EndpointSAMLMutate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointSAMLMutate(&v))
	}
	return sl
}

func expandEndpointSAMLMutate(in interface{}) *restapi.EndpointSAMLMutate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointSAMLMutate
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["options_passthrough"]; ok {
		obj.OptionsPassthrough = *expandBool(v)
	}
	if v, ok := m["cookie_prefix"]; ok {
		obj.CookiePrefix = *expandString(v)
	}
	if v, ok := m["inactivity_timeout"]; ok {
		obj.InactivityTimeout = *expandUint32(v)
	}
	if v, ok := m["maximum_duration"]; ok {
		obj.MaximumDuration = *expandUint32(v)
	}
	if v, ok := m["idp_metadata_url"]; ok {
		obj.IdPMetadataURL = *expandString(v)
	}
	if v, ok := m["idp_metadata"]; ok {
		obj.IdPMetadata = *expandString(v)
	}
	if v, ok := m["force_authn"]; ok {
		obj.ForceAuthn = *expandBool(v)
	}
	if v, ok := m["allow_idp_initiated"]; ok {
		obj.AllowIdPInitiated = expandBool(v)
	}
	if v, ok := m["authorized_groups"]; ok {
		obj.AuthorizedGroups = *expandStringSlice(v)
	}
	if v, ok := m["nameid_format"]; ok {
		obj.NameIDFormat = *expandString(v)
	}
	return &obj
}

func expandEndpointSAMLMutateSlice(in interface{}) *[]restapi.EndpointSAMLMutate {
	var out []restapi.EndpointSAMLMutate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointSAMLMutate(v))
	}
	return &out
}

func flattenEndpointOIDC(obj *restapi.EndpointOIDC) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["options_passthrough"] = obj.OptionsPassthrough
	m["cookie_prefix"] = obj.CookiePrefix
	m["inactivity_timeout"] = obj.InactivityTimeout
	m["maximum_duration"] = obj.MaximumDuration
	m["issuer"] = obj.Issuer
	m["client_id"] = obj.ClientID
	m["client_secret"] = obj.ClientSecret
	m["scopes"] = obj.Scopes

	return []interface{}{m}
}

func flattenEndpointOIDCSlice(objs *[]restapi.EndpointOIDC) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOIDC(&v))
	}
	return sl
}

func expandEndpointOIDC(in interface{}) *restapi.EndpointOIDC {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOIDC
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["options_passthrough"]; ok {
		obj.OptionsPassthrough = *expandBool(v)
	}
	if v, ok := m["cookie_prefix"]; ok {
		obj.CookiePrefix = *expandString(v)
	}
	if v, ok := m["inactivity_timeout"]; ok {
		obj.InactivityTimeout = *expandUint32(v)
	}
	if v, ok := m["maximum_duration"]; ok {
		obj.MaximumDuration = *expandUint32(v)
	}
	if v, ok := m["issuer"]; ok {
		obj.Issuer = *expandString(v)
	}
	if v, ok := m["client_id"]; ok {
		obj.ClientID = *expandString(v)
	}
	if v, ok := m["client_secret"]; ok {
		obj.ClientSecret = *expandString(v)
	}
	if v, ok := m["scopes"]; ok {
		obj.Scopes = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointOIDCSlice(in interface{}) *[]restapi.EndpointOIDC {
	var out []restapi.EndpointOIDC
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOIDC(v))
	}
	return &out
}

func flattenEndpointBackend(obj *restapi.EndpointBackend) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["backend"] = flattenRef(&obj.Backend)

	return []interface{}{m}
}

func flattenEndpointBackendSlice(objs *[]restapi.EndpointBackend) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointBackend(&v))
	}
	return sl
}

func expandEndpointBackend(in interface{}) *restapi.EndpointBackend {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointBackend
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = *expandRef(v)
	}
	return &obj
}

func expandEndpointBackendSlice(in interface{}) *[]restapi.EndpointBackend {
	var out []restapi.EndpointBackend
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointBackend(v))
	}
	return &out
}

func flattenEndpointBackendMutate(obj *restapi.EndpointBackendMutate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["backend_id"] = obj.BackendID

	return []interface{}{m}
}

func flattenEndpointBackendMutateSlice(objs *[]restapi.EndpointBackendMutate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointBackendMutate(&v))
	}
	return sl
}

func expandEndpointBackendMutate(in interface{}) *restapi.EndpointBackendMutate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointBackendMutate
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["backend_id"]; ok {
		obj.BackendID = *expandString(v)
	}
	return &obj
}

func expandEndpointBackendMutateSlice(in interface{}) *[]restapi.EndpointBackendMutate {
	var out []restapi.EndpointBackendMutate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointBackendMutate(v))
	}
	return &out
}

func flattenEndpointWebsocketTCPConverter(obj *restapi.EndpointWebsocketTCPConverter) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled

	return []interface{}{m}
}

func flattenEndpointWebsocketTCPConverterSlice(objs *[]restapi.EndpointWebsocketTCPConverter) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointWebsocketTCPConverter(&v))
	}
	return sl
}

func expandEndpointWebsocketTCPConverter(in interface{}) *restapi.EndpointWebsocketTCPConverter {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointWebsocketTCPConverter
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	return &obj
}

func expandEndpointWebsocketTCPConverterSlice(in interface{}) *[]restapi.EndpointWebsocketTCPConverter {
	var out []restapi.EndpointWebsocketTCPConverter
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointWebsocketTCPConverter(v))
	}
	return &out
}

func flattenEndpointUserAgentFilter(obj *restapi.EndpointUserAgentFilter) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["allow"] = obj.UserAgentFilterAllow
	m["deny"] = obj.UserAgentFilterDeny

	return []interface{}{m}
}

func flattenEndpointUserAgentFilterSlice(objs *[]restapi.EndpointUserAgentFilter) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointUserAgentFilter(&v))
	}
	return sl
}

func expandEndpointUserAgentFilter(in interface{}) *restapi.EndpointUserAgentFilter {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointUserAgentFilter
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["allow"]; ok {
		obj.UserAgentFilterAllow = *expandStringSlice(v)
	}
	if v, ok := m["deny"]; ok {
		obj.UserAgentFilterDeny = *expandStringSlice(v)
	}
	return &obj
}

func expandEndpointUserAgentFilterSlice(in interface{}) *[]restapi.EndpointUserAgentFilter {
	var out []restapi.EndpointUserAgentFilter
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointUserAgentFilter(v))
	}
	return &out
}

func flattenEndpointPolicy(obj *restapi.EndpointPolicy) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["inbound"] = flattenEndpointRuleSlice(&obj.Inbound)
	m["outbound"] = flattenEndpointRuleSlice(&obj.Outbound)

	return []interface{}{m}
}

func flattenEndpointPolicySlice(objs *[]restapi.EndpointPolicy) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointPolicy(&v))
	}
	return sl
}

func expandEndpointPolicy(in interface{}) *restapi.EndpointPolicy {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointPolicy
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["inbound"]; ok {
		obj.Inbound = *expandEndpointRuleSlice(v)
	}
	if v, ok := m["outbound"]; ok {
		obj.Outbound = *expandEndpointRuleSlice(v)
	}
	return &obj
}

func expandEndpointPolicySlice(in interface{}) *[]restapi.EndpointPolicy {
	var out []restapi.EndpointPolicy
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointPolicy(v))
	}
	return &out
}

func flattenEndpointRule(obj *restapi.EndpointRule) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["expressions"] = obj.Expressions
	m["actions"] = flattenEndpointActionSlice(&obj.Actions)
	m["name"] = obj.Name

	return []interface{}{m}
}

func flattenEndpointRuleSlice(objs *[]restapi.EndpointRule) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointRule(&v))
	}
	return sl
}

func expandEndpointRule(in interface{}) *restapi.EndpointRule {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointRule
	if v, ok := m["expressions"]; ok {
		obj.Expressions = *expandStringSlice(v)
	}
	if v, ok := m["actions"]; ok {
		obj.Actions = *expandEndpointActionSlice(v)
	}
	if v, ok := m["name"]; ok {
		obj.Name = *expandString(v)
	}
	return &obj
}

func expandEndpointRuleSlice(in interface{}) *[]restapi.EndpointRule {
	var out []restapi.EndpointRule
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointRule(v))
	}
	return &out
}

func flattenEndpointAction(obj *restapi.EndpointAction) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["type"] = obj.Type
	m["config"] = obj.Config

	return []interface{}{m}
}

func flattenEndpointActionSlice(objs *[]restapi.EndpointAction) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointAction(&v))
	}
	return sl
}

func expandEndpointAction(in interface{}) *restapi.EndpointAction {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointAction
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["config"]; ok {
		obj.Config = *expandAny(v)
	}
	return &obj
}

func expandEndpointActionSlice(in interface{}) *[]restapi.EndpointAction {
	var out []restapi.EndpointAction
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointAction(v))
	}
	return &out
}

func flattenEndpointTrafficPolicy(obj *restapi.EndpointTrafficPolicy) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["enabled"] = obj.Enabled
	m["value"] = obj.Value

	return []interface{}{m}
}

func flattenEndpointTrafficPolicySlice(objs *[]restapi.EndpointTrafficPolicy) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointTrafficPolicy(&v))
	}
	return sl
}

func expandEndpointTrafficPolicy(in interface{}) *restapi.EndpointTrafficPolicy {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointTrafficPolicy
	if v, ok := m["enabled"]; ok {
		obj.Enabled = expandBool(v)
	}
	if v, ok := m["value"]; ok {
		obj.Value = *expandString(v)
	}
	return &obj
}

func expandEndpointTrafficPolicySlice(in interface{}) *[]restapi.EndpointTrafficPolicy {
	var out []restapi.EndpointTrafficPolicy
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointTrafficPolicy(v))
	}
	return &out
}

func flattenEdgeRouteItem(obj *restapi.EdgeRouteItem) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID

	return []interface{}{m}
}

func flattenEdgeRouteItemSlice(objs *[]restapi.EdgeRouteItem) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteItem(&v))
	}
	return sl
}

func expandEdgeRouteItem(in interface{}) *restapi.EdgeRouteItem {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteItem
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	return &obj
}

func expandEdgeRouteItemSlice(in interface{}) *[]restapi.EdgeRouteItem {
	var out []restapi.EdgeRouteItem
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteItem(v))
	}
	return &out
}

func flattenHTTPSEdgeRouteCreate(obj *restapi.HTTPSEdgeRouteCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["match_type"] = obj.MatchType
	m["match"] = obj.Match
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["backend"] = flattenEndpointBackendMutate(obj.Backend)
	m["ip_restriction"] = flattenEndpointIPPolicyMutate(obj.IPRestriction)
	m["circuit_breaker"] = flattenEndpointCircuitBreaker(obj.CircuitBreaker)
	m["compression"] = flattenEndpointCompression(obj.Compression)
	m["request_headers"] = flattenEndpointRequestHeaders(obj.RequestHeaders)
	m["response_headers"] = flattenEndpointResponseHeaders(obj.ResponseHeaders)
	m["webhook_verification"] = flattenEndpointWebhookValidation(obj.WebhookVerification)
	m["oauth"] = flattenEndpointOAuth(obj.OAuth)
	m["saml"] = flattenEndpointSAMLMutate(obj.SAML)
	m["oidc"] = flattenEndpointOIDC(obj.OIDC)
	m["websocket_tcp_converter"] = flattenEndpointWebsocketTCPConverter(obj.WebsocketTCPConverter)
	m["user_agent_filter"] = flattenEndpointUserAgentFilter(obj.UserAgentFilter)
	m["policy"] = flattenEndpointPolicy(obj.Policy)
	m["traffic_policy"] = flattenEndpointTrafficPolicy(obj.TrafficPolicy)

	return []interface{}{m}
}

func flattenHTTPSEdgeRouteCreateSlice(objs *[]restapi.HTTPSEdgeRouteCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPSEdgeRouteCreate(&v))
	}
	return sl
}

func expandHTTPSEdgeRouteCreate(in interface{}) *restapi.HTTPSEdgeRouteCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPSEdgeRouteCreate
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["match_type"]; ok {
		obj.MatchType = *expandString(v)
	}
	if v, ok := m["match"]; ok {
		obj.Match = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackendMutate(v)
	}
	if v, ok := m["ip_restriction"]; ok {
		obj.IPRestriction = expandEndpointIPPolicyMutate(v)
	}
	if v, ok := m["circuit_breaker"]; ok {
		obj.CircuitBreaker = expandEndpointCircuitBreaker(v)
	}
	if v, ok := m["compression"]; ok {
		obj.Compression = expandEndpointCompression(v)
	}
	if v, ok := m["request_headers"]; ok {
		obj.RequestHeaders = expandEndpointRequestHeaders(v)
	}
	if v, ok := m["response_headers"]; ok {
		obj.ResponseHeaders = expandEndpointResponseHeaders(v)
	}
	if v, ok := m["webhook_verification"]; ok {
		obj.WebhookVerification = expandEndpointWebhookValidation(v)
	}
	if v, ok := m["oauth"]; ok {
		obj.OAuth = expandEndpointOAuth(v)
	}
	if v, ok := m["saml"]; ok {
		obj.SAML = expandEndpointSAMLMutate(v)
	}
	if v, ok := m["oidc"]; ok {
		obj.OIDC = expandEndpointOIDC(v)
	}
	if v, ok := m["websocket_tcp_converter"]; ok {
		obj.WebsocketTCPConverter = expandEndpointWebsocketTCPConverter(v)
	}
	if v, ok := m["user_agent_filter"]; ok {
		obj.UserAgentFilter = expandEndpointUserAgentFilter(v)
	}
	if v, ok := m["policy"]; ok {
		obj.Policy = expandEndpointPolicy(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandHTTPSEdgeRouteCreateSlice(in interface{}) *[]restapi.HTTPSEdgeRouteCreate {
	var out []restapi.HTTPSEdgeRouteCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPSEdgeRouteCreate(v))
	}
	return &out
}

func flattenHTTPSEdgeRouteUpdate(obj *restapi.HTTPSEdgeRouteUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["match_type"] = obj.MatchType
	m["match"] = obj.Match
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["backend"] = flattenEndpointBackendMutate(obj.Backend)
	m["ip_restriction"] = flattenEndpointIPPolicyMutate(obj.IPRestriction)
	m["circuit_breaker"] = flattenEndpointCircuitBreaker(obj.CircuitBreaker)
	m["compression"] = flattenEndpointCompression(obj.Compression)
	m["request_headers"] = flattenEndpointRequestHeaders(obj.RequestHeaders)
	m["response_headers"] = flattenEndpointResponseHeaders(obj.ResponseHeaders)
	m["webhook_verification"] = flattenEndpointWebhookValidation(obj.WebhookVerification)
	m["oauth"] = flattenEndpointOAuth(obj.OAuth)
	m["saml"] = flattenEndpointSAMLMutate(obj.SAML)
	m["oidc"] = flattenEndpointOIDC(obj.OIDC)
	m["websocket_tcp_converter"] = flattenEndpointWebsocketTCPConverter(obj.WebsocketTCPConverter)
	m["user_agent_filter"] = flattenEndpointUserAgentFilter(obj.UserAgentFilter)
	m["policy"] = flattenEndpointPolicy(obj.Policy)
	m["traffic_policy"] = flattenEndpointTrafficPolicy(obj.TrafficPolicy)

	return []interface{}{m}
}

func flattenHTTPSEdgeRouteUpdateSlice(objs *[]restapi.HTTPSEdgeRouteUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPSEdgeRouteUpdate(&v))
	}
	return sl
}

func expandHTTPSEdgeRouteUpdate(in interface{}) *restapi.HTTPSEdgeRouteUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPSEdgeRouteUpdate
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["match_type"]; ok {
		obj.MatchType = *expandString(v)
	}
	if v, ok := m["match"]; ok {
		obj.Match = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackendMutate(v)
	}
	if v, ok := m["ip_restriction"]; ok {
		obj.IPRestriction = expandEndpointIPPolicyMutate(v)
	}
	if v, ok := m["circuit_breaker"]; ok {
		obj.CircuitBreaker = expandEndpointCircuitBreaker(v)
	}
	if v, ok := m["compression"]; ok {
		obj.Compression = expandEndpointCompression(v)
	}
	if v, ok := m["request_headers"]; ok {
		obj.RequestHeaders = expandEndpointRequestHeaders(v)
	}
	if v, ok := m["response_headers"]; ok {
		obj.ResponseHeaders = expandEndpointResponseHeaders(v)
	}
	if v, ok := m["webhook_verification"]; ok {
		obj.WebhookVerification = expandEndpointWebhookValidation(v)
	}
	if v, ok := m["oauth"]; ok {
		obj.OAuth = expandEndpointOAuth(v)
	}
	if v, ok := m["saml"]; ok {
		obj.SAML = expandEndpointSAMLMutate(v)
	}
	if v, ok := m["oidc"]; ok {
		obj.OIDC = expandEndpointOIDC(v)
	}
	if v, ok := m["websocket_tcp_converter"]; ok {
		obj.WebsocketTCPConverter = expandEndpointWebsocketTCPConverter(v)
	}
	if v, ok := m["user_agent_filter"]; ok {
		obj.UserAgentFilter = expandEndpointUserAgentFilter(v)
	}
	if v, ok := m["policy"]; ok {
		obj.Policy = expandEndpointPolicy(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandHTTPSEdgeRouteUpdateSlice(in interface{}) *[]restapi.HTTPSEdgeRouteUpdate {
	var out []restapi.HTTPSEdgeRouteUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPSEdgeRouteUpdate(v))
	}
	return &out
}

func flattenHTTPSEdgeRoute(obj *restapi.HTTPSEdgeRoute) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["created_at"] = obj.CreatedAt
	m["match_type"] = obj.MatchType
	m["match"] = obj.Match
	m["uri"] = obj.URI
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["backend"] = flattenEndpointBackend(obj.Backend)
	m["ip_restriction"] = flattenEndpointIPPolicy(obj.IpRestriction)
	m["circuit_breaker"] = flattenEndpointCircuitBreaker(obj.CircuitBreaker)
	m["compression"] = flattenEndpointCompression(obj.Compression)
	m["request_headers"] = flattenEndpointRequestHeaders(obj.RequestHeaders)
	m["response_headers"] = flattenEndpointResponseHeaders(obj.ResponseHeaders)
	m["webhook_verification"] = flattenEndpointWebhookValidation(obj.WebhookVerification)
	m["oauth"] = flattenEndpointOAuth(obj.OAuth)
	m["saml"] = flattenEndpointSAML(obj.SAML)
	m["oidc"] = flattenEndpointOIDC(obj.OIDC)
	m["websocket_tcp_converter"] = flattenEndpointWebsocketTCPConverter(obj.WebsocketTCPConverter)
	m["user_agent_filter"] = flattenEndpointUserAgentFilter(obj.UserAgentFilter)
	m["policy"] = flattenEndpointPolicy(obj.Policy)
	m["traffic_policy"] = flattenEndpointTrafficPolicy(obj.TrafficPolicy)

	return []interface{}{m}
}

func flattenHTTPSEdgeRouteSlice(objs *[]restapi.HTTPSEdgeRoute) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPSEdgeRoute(&v))
	}
	return sl
}

func expandHTTPSEdgeRoute(in interface{}) *restapi.HTTPSEdgeRoute {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPSEdgeRoute
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["match_type"]; ok {
		obj.MatchType = *expandString(v)
	}
	if v, ok := m["match"]; ok {
		obj.Match = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackend(v)
	}
	if v, ok := m["ip_restriction"]; ok {
		obj.IpRestriction = expandEndpointIPPolicy(v)
	}
	if v, ok := m["circuit_breaker"]; ok {
		obj.CircuitBreaker = expandEndpointCircuitBreaker(v)
	}
	if v, ok := m["compression"]; ok {
		obj.Compression = expandEndpointCompression(v)
	}
	if v, ok := m["request_headers"]; ok {
		obj.RequestHeaders = expandEndpointRequestHeaders(v)
	}
	if v, ok := m["response_headers"]; ok {
		obj.ResponseHeaders = expandEndpointResponseHeaders(v)
	}
	if v, ok := m["webhook_verification"]; ok {
		obj.WebhookVerification = expandEndpointWebhookValidation(v)
	}
	if v, ok := m["oauth"]; ok {
		obj.OAuth = expandEndpointOAuth(v)
	}
	if v, ok := m["saml"]; ok {
		obj.SAML = expandEndpointSAML(v)
	}
	if v, ok := m["oidc"]; ok {
		obj.OIDC = expandEndpointOIDC(v)
	}
	if v, ok := m["websocket_tcp_converter"]; ok {
		obj.WebsocketTCPConverter = expandEndpointWebsocketTCPConverter(v)
	}
	if v, ok := m["user_agent_filter"]; ok {
		obj.UserAgentFilter = expandEndpointUserAgentFilter(v)
	}
	if v, ok := m["policy"]; ok {
		obj.Policy = expandEndpointPolicy(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandHTTPSEdgeRouteSlice(in interface{}) *[]restapi.HTTPSEdgeRoute {
	var out []restapi.HTTPSEdgeRoute
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPSEdgeRoute(v))
	}
	return &out
}

func flattenHTTPSEdgeList(obj *restapi.HTTPSEdgeList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["https_edges"] = flattenHTTPSEdgeSlice(&obj.HTTPSEdges)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenHTTPSEdgeListSlice(objs *[]restapi.HTTPSEdgeList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPSEdgeList(&v))
	}
	return sl
}

func expandHTTPSEdgeList(in interface{}) *restapi.HTTPSEdgeList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPSEdgeList
	if v, ok := m["https_edges"]; ok {
		obj.HTTPSEdges = *expandHTTPSEdgeSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandHTTPSEdgeListSlice(in interface{}) *[]restapi.HTTPSEdgeList {
	var out []restapi.HTTPSEdgeList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPSEdgeList(v))
	}
	return &out
}

func flattenHTTPSEdgeCreate(obj *restapi.HTTPSEdgeCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["hostports"] = obj.Hostports
	m["mutual_tls"] = flattenEndpointMutualTLSMutate(obj.MutualTLS)
	m["tls_termination"] = flattenEndpointTLSTerminationAtEdge(obj.TLSTermination)

	return []interface{}{m}
}

func flattenHTTPSEdgeCreateSlice(objs *[]restapi.HTTPSEdgeCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPSEdgeCreate(&v))
	}
	return sl
}

func expandHTTPSEdgeCreate(in interface{}) *restapi.HTTPSEdgeCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPSEdgeCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["hostports"]; ok {
		obj.Hostports = expandStringSlice(v)
	}
	if v, ok := m["mutual_tls"]; ok {
		obj.MutualTLS = expandEndpointMutualTLSMutate(v)
	}
	if v, ok := m["tls_termination"]; ok {
		obj.TLSTermination = expandEndpointTLSTerminationAtEdge(v)
	}
	return &obj
}

func expandHTTPSEdgeCreateSlice(in interface{}) *[]restapi.HTTPSEdgeCreate {
	var out []restapi.HTTPSEdgeCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPSEdgeCreate(v))
	}
	return &out
}

func flattenHTTPSEdgeUpdate(obj *restapi.HTTPSEdgeUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["hostports"] = obj.Hostports
	m["mutual_tls"] = flattenEndpointMutualTLSMutate(obj.MutualTLS)
	m["tls_termination"] = flattenEndpointTLSTerminationAtEdge(obj.TLSTermination)

	return []interface{}{m}
}

func flattenHTTPSEdgeUpdateSlice(objs *[]restapi.HTTPSEdgeUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPSEdgeUpdate(&v))
	}
	return sl
}

func expandHTTPSEdgeUpdate(in interface{}) *restapi.HTTPSEdgeUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPSEdgeUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["hostports"]; ok {
		obj.Hostports = expandStringSlice(v)
	}
	if v, ok := m["mutual_tls"]; ok {
		obj.MutualTLS = expandEndpointMutualTLSMutate(v)
	}
	if v, ok := m["tls_termination"]; ok {
		obj.TLSTermination = expandEndpointTLSTerminationAtEdge(v)
	}
	return &obj
}

func expandHTTPSEdgeUpdateSlice(in interface{}) *[]restapi.HTTPSEdgeUpdate {
	var out []restapi.HTTPSEdgeUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPSEdgeUpdate(v))
	}
	return &out
}

func flattenHTTPSEdge(obj *restapi.HTTPSEdge) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["created_at"] = obj.CreatedAt
	m["uri"] = obj.URI
	m["hostports"] = obj.Hostports
	m["mutual_tls"] = flattenEndpointMutualTLS(obj.MutualTls)
	m["tls_termination"] = flattenEndpointTLSTermination(obj.TlsTermination)
	m["routes"] = flattenHTTPSEdgeRouteSlice(&obj.Routes)

	return []interface{}{m}
}

func flattenHTTPSEdgeSlice(objs *[]restapi.HTTPSEdge) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenHTTPSEdge(&v))
	}
	return sl
}

func expandHTTPSEdge(in interface{}) *restapi.HTTPSEdge {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.HTTPSEdge
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["hostports"]; ok {
		obj.Hostports = expandStringSlice(v)
	}
	if v, ok := m["mutual_tls"]; ok {
		obj.MutualTls = expandEndpointMutualTLS(v)
	}
	if v, ok := m["tls_termination"]; ok {
		obj.TlsTermination = expandEndpointTLSTermination(v)
	}
	if v, ok := m["routes"]; ok {
		obj.Routes = *expandHTTPSEdgeRouteSlice(v)
	}
	return &obj
}

func expandHTTPSEdgeSlice(in interface{}) *[]restapi.HTTPSEdge {
	var out []restapi.HTTPSEdge
	for _, v := range in.([]interface{}) {
		out = append(out, *expandHTTPSEdge(v))
	}
	return &out
}

func flattenEdgeBackendReplace(obj *restapi.EdgeBackendReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointBackendMutate(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeBackendReplaceSlice(objs *[]restapi.EdgeBackendReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeBackendReplace(&v))
	}
	return sl
}

func expandEdgeBackendReplace(in interface{}) *restapi.EdgeBackendReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeBackendReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointBackendMutate(v)
	}
	return &obj
}

func expandEdgeBackendReplaceSlice(in interface{}) *[]restapi.EdgeBackendReplace {
	var out []restapi.EdgeBackendReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeBackendReplace(v))
	}
	return &out
}

func flattenEdgeIPRestrictionReplace(obj *restapi.EdgeIPRestrictionReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointIPPolicyMutate(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeIPRestrictionReplaceSlice(objs *[]restapi.EdgeIPRestrictionReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeIPRestrictionReplace(&v))
	}
	return sl
}

func expandEdgeIPRestrictionReplace(in interface{}) *restapi.EdgeIPRestrictionReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeIPRestrictionReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointIPPolicyMutate(v)
	}
	return &obj
}

func expandEdgeIPRestrictionReplaceSlice(in interface{}) *[]restapi.EdgeIPRestrictionReplace {
	var out []restapi.EdgeIPRestrictionReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeIPRestrictionReplace(v))
	}
	return &out
}

func flattenEdgeMutualTLSReplace(obj *restapi.EdgeMutualTLSReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointMutualTLSMutate(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeMutualTLSReplaceSlice(objs *[]restapi.EdgeMutualTLSReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeMutualTLSReplace(&v))
	}
	return sl
}

func expandEdgeMutualTLSReplace(in interface{}) *restapi.EdgeMutualTLSReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeMutualTLSReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointMutualTLSMutate(v)
	}
	return &obj
}

func expandEdgeMutualTLSReplaceSlice(in interface{}) *[]restapi.EdgeMutualTLSReplace {
	var out []restapi.EdgeMutualTLSReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeMutualTLSReplace(v))
	}
	return &out
}

func flattenEdgeTLSTerminationReplace(obj *restapi.EdgeTLSTerminationReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointTLSTermination(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeTLSTerminationReplaceSlice(objs *[]restapi.EdgeTLSTerminationReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeTLSTerminationReplace(&v))
	}
	return sl
}

func expandEdgeTLSTerminationReplace(in interface{}) *restapi.EdgeTLSTerminationReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeTLSTerminationReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointTLSTermination(v)
	}
	return &obj
}

func expandEdgeTLSTerminationReplaceSlice(in interface{}) *[]restapi.EdgeTLSTerminationReplace {
	var out []restapi.EdgeTLSTerminationReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeTLSTerminationReplace(v))
	}
	return &out
}

func flattenEdgeTLSTerminationAtEdgeReplace(obj *restapi.EdgeTLSTerminationAtEdgeReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointTLSTerminationAtEdge(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeTLSTerminationAtEdgeReplaceSlice(objs *[]restapi.EdgeTLSTerminationAtEdgeReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeTLSTerminationAtEdgeReplace(&v))
	}
	return sl
}

func expandEdgeTLSTerminationAtEdgeReplace(in interface{}) *restapi.EdgeTLSTerminationAtEdgeReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeTLSTerminationAtEdgeReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointTLSTerminationAtEdge(v)
	}
	return &obj
}

func expandEdgeTLSTerminationAtEdgeReplaceSlice(in interface{}) *[]restapi.EdgeTLSTerminationAtEdgeReplace {
	var out []restapi.EdgeTLSTerminationAtEdgeReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeTLSTerminationAtEdgeReplace(v))
	}
	return &out
}

func flattenEdgePolicyReplace(obj *restapi.EdgePolicyReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointPolicy(&obj.Module)

	return []interface{}{m}
}

func flattenEdgePolicyReplaceSlice(objs *[]restapi.EdgePolicyReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgePolicyReplace(&v))
	}
	return sl
}

func expandEdgePolicyReplace(in interface{}) *restapi.EdgePolicyReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgePolicyReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointPolicy(v)
	}
	return &obj
}

func expandEdgePolicyReplaceSlice(in interface{}) *[]restapi.EdgePolicyReplace {
	var out []restapi.EdgePolicyReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgePolicyReplace(v))
	}
	return &out
}

func flattenEdgeTrafficPolicyReplace(obj *restapi.EdgeTrafficPolicyReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointTrafficPolicy(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeTrafficPolicyReplaceSlice(objs *[]restapi.EdgeTrafficPolicyReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeTrafficPolicyReplace(&v))
	}
	return sl
}

func expandEdgeTrafficPolicyReplace(in interface{}) *restapi.EdgeTrafficPolicyReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeTrafficPolicyReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandEdgeTrafficPolicyReplaceSlice(in interface{}) *[]restapi.EdgeTrafficPolicyReplace {
	var out []restapi.EdgeTrafficPolicyReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeTrafficPolicyReplace(v))
	}
	return &out
}

func flattenEdgeRouteBackendReplace(obj *restapi.EdgeRouteBackendReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointBackendMutate(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteBackendReplaceSlice(objs *[]restapi.EdgeRouteBackendReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteBackendReplace(&v))
	}
	return sl
}

func expandEdgeRouteBackendReplace(in interface{}) *restapi.EdgeRouteBackendReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteBackendReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointBackendMutate(v)
	}
	return &obj
}

func expandEdgeRouteBackendReplaceSlice(in interface{}) *[]restapi.EdgeRouteBackendReplace {
	var out []restapi.EdgeRouteBackendReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteBackendReplace(v))
	}
	return &out
}

func flattenEdgeRouteIPRestrictionReplace(obj *restapi.EdgeRouteIPRestrictionReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointIPPolicyMutate(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteIPRestrictionReplaceSlice(objs *[]restapi.EdgeRouteIPRestrictionReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteIPRestrictionReplace(&v))
	}
	return sl
}

func expandEdgeRouteIPRestrictionReplace(in interface{}) *restapi.EdgeRouteIPRestrictionReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteIPRestrictionReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointIPPolicyMutate(v)
	}
	return &obj
}

func expandEdgeRouteIPRestrictionReplaceSlice(in interface{}) *[]restapi.EdgeRouteIPRestrictionReplace {
	var out []restapi.EdgeRouteIPRestrictionReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteIPRestrictionReplace(v))
	}
	return &out
}

func flattenEdgeRouteRequestHeadersReplace(obj *restapi.EdgeRouteRequestHeadersReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointRequestHeaders(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteRequestHeadersReplaceSlice(objs *[]restapi.EdgeRouteRequestHeadersReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteRequestHeadersReplace(&v))
	}
	return sl
}

func expandEdgeRouteRequestHeadersReplace(in interface{}) *restapi.EdgeRouteRequestHeadersReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteRequestHeadersReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointRequestHeaders(v)
	}
	return &obj
}

func expandEdgeRouteRequestHeadersReplaceSlice(in interface{}) *[]restapi.EdgeRouteRequestHeadersReplace {
	var out []restapi.EdgeRouteRequestHeadersReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteRequestHeadersReplace(v))
	}
	return &out
}

func flattenEdgeRouteResponseHeadersReplace(obj *restapi.EdgeRouteResponseHeadersReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointResponseHeaders(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteResponseHeadersReplaceSlice(objs *[]restapi.EdgeRouteResponseHeadersReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteResponseHeadersReplace(&v))
	}
	return sl
}

func expandEdgeRouteResponseHeadersReplace(in interface{}) *restapi.EdgeRouteResponseHeadersReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteResponseHeadersReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointResponseHeaders(v)
	}
	return &obj
}

func expandEdgeRouteResponseHeadersReplaceSlice(in interface{}) *[]restapi.EdgeRouteResponseHeadersReplace {
	var out []restapi.EdgeRouteResponseHeadersReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteResponseHeadersReplace(v))
	}
	return &out
}

func flattenEdgeRouteCompressionReplace(obj *restapi.EdgeRouteCompressionReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointCompression(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteCompressionReplaceSlice(objs *[]restapi.EdgeRouteCompressionReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteCompressionReplace(&v))
	}
	return sl
}

func expandEdgeRouteCompressionReplace(in interface{}) *restapi.EdgeRouteCompressionReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteCompressionReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointCompression(v)
	}
	return &obj
}

func expandEdgeRouteCompressionReplaceSlice(in interface{}) *[]restapi.EdgeRouteCompressionReplace {
	var out []restapi.EdgeRouteCompressionReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteCompressionReplace(v))
	}
	return &out
}

func flattenEdgeRouteCircuitBreakerReplace(obj *restapi.EdgeRouteCircuitBreakerReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointCircuitBreaker(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteCircuitBreakerReplaceSlice(objs *[]restapi.EdgeRouteCircuitBreakerReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteCircuitBreakerReplace(&v))
	}
	return sl
}

func expandEdgeRouteCircuitBreakerReplace(in interface{}) *restapi.EdgeRouteCircuitBreakerReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteCircuitBreakerReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointCircuitBreaker(v)
	}
	return &obj
}

func expandEdgeRouteCircuitBreakerReplaceSlice(in interface{}) *[]restapi.EdgeRouteCircuitBreakerReplace {
	var out []restapi.EdgeRouteCircuitBreakerReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteCircuitBreakerReplace(v))
	}
	return &out
}

func flattenEdgeRouteWebhookVerificationReplace(obj *restapi.EdgeRouteWebhookVerificationReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointWebhookValidation(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteWebhookVerificationReplaceSlice(objs *[]restapi.EdgeRouteWebhookVerificationReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteWebhookVerificationReplace(&v))
	}
	return sl
}

func expandEdgeRouteWebhookVerificationReplace(in interface{}) *restapi.EdgeRouteWebhookVerificationReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteWebhookVerificationReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointWebhookValidation(v)
	}
	return &obj
}

func expandEdgeRouteWebhookVerificationReplaceSlice(in interface{}) *[]restapi.EdgeRouteWebhookVerificationReplace {
	var out []restapi.EdgeRouteWebhookVerificationReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteWebhookVerificationReplace(v))
	}
	return &out
}

func flattenEdgeRouteOAuthReplace(obj *restapi.EdgeRouteOAuthReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointOAuth(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteOAuthReplaceSlice(objs *[]restapi.EdgeRouteOAuthReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteOAuthReplace(&v))
	}
	return sl
}

func expandEdgeRouteOAuthReplace(in interface{}) *restapi.EdgeRouteOAuthReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteOAuthReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointOAuth(v)
	}
	return &obj
}

func expandEdgeRouteOAuthReplaceSlice(in interface{}) *[]restapi.EdgeRouteOAuthReplace {
	var out []restapi.EdgeRouteOAuthReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteOAuthReplace(v))
	}
	return &out
}

func flattenEdgeRouteSAMLReplace(obj *restapi.EdgeRouteSAMLReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointSAMLMutate(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteSAMLReplaceSlice(objs *[]restapi.EdgeRouteSAMLReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteSAMLReplace(&v))
	}
	return sl
}

func expandEdgeRouteSAMLReplace(in interface{}) *restapi.EdgeRouteSAMLReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteSAMLReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointSAMLMutate(v)
	}
	return &obj
}

func expandEdgeRouteSAMLReplaceSlice(in interface{}) *[]restapi.EdgeRouteSAMLReplace {
	var out []restapi.EdgeRouteSAMLReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteSAMLReplace(v))
	}
	return &out
}

func flattenEdgeRouteOIDCReplace(obj *restapi.EdgeRouteOIDCReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointOIDC(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteOIDCReplaceSlice(objs *[]restapi.EdgeRouteOIDCReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteOIDCReplace(&v))
	}
	return sl
}

func expandEdgeRouteOIDCReplace(in interface{}) *restapi.EdgeRouteOIDCReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteOIDCReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointOIDC(v)
	}
	return &obj
}

func expandEdgeRouteOIDCReplaceSlice(in interface{}) *[]restapi.EdgeRouteOIDCReplace {
	var out []restapi.EdgeRouteOIDCReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteOIDCReplace(v))
	}
	return &out
}

func flattenEdgeRouteWebsocketTCPConverterReplace(obj *restapi.EdgeRouteWebsocketTCPConverterReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointWebsocketTCPConverter(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteWebsocketTCPConverterReplaceSlice(objs *[]restapi.EdgeRouteWebsocketTCPConverterReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteWebsocketTCPConverterReplace(&v))
	}
	return sl
}

func expandEdgeRouteWebsocketTCPConverterReplace(in interface{}) *restapi.EdgeRouteWebsocketTCPConverterReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteWebsocketTCPConverterReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointWebsocketTCPConverter(v)
	}
	return &obj
}

func expandEdgeRouteWebsocketTCPConverterReplaceSlice(in interface{}) *[]restapi.EdgeRouteWebsocketTCPConverterReplace {
	var out []restapi.EdgeRouteWebsocketTCPConverterReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteWebsocketTCPConverterReplace(v))
	}
	return &out
}

func flattenEdgeRouteUserAgentFilterReplace(obj *restapi.EdgeRouteUserAgentFilterReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointUserAgentFilter(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteUserAgentFilterReplaceSlice(objs *[]restapi.EdgeRouteUserAgentFilterReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteUserAgentFilterReplace(&v))
	}
	return sl
}

func expandEdgeRouteUserAgentFilterReplace(in interface{}) *restapi.EdgeRouteUserAgentFilterReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteUserAgentFilterReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointUserAgentFilter(v)
	}
	return &obj
}

func expandEdgeRouteUserAgentFilterReplaceSlice(in interface{}) *[]restapi.EdgeRouteUserAgentFilterReplace {
	var out []restapi.EdgeRouteUserAgentFilterReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteUserAgentFilterReplace(v))
	}
	return &out
}

func flattenEdgeRoutePolicyReplace(obj *restapi.EdgeRoutePolicyReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointPolicy(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRoutePolicyReplaceSlice(objs *[]restapi.EdgeRoutePolicyReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRoutePolicyReplace(&v))
	}
	return sl
}

func expandEdgeRoutePolicyReplace(in interface{}) *restapi.EdgeRoutePolicyReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRoutePolicyReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointPolicy(v)
	}
	return &obj
}

func expandEdgeRoutePolicyReplaceSlice(in interface{}) *[]restapi.EdgeRoutePolicyReplace {
	var out []restapi.EdgeRoutePolicyReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRoutePolicyReplace(v))
	}
	return &out
}

func flattenEdgeRouteTrafficPolicyReplace(obj *restapi.EdgeRouteTrafficPolicyReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["edge_id"] = obj.EdgeID
	m["id"] = obj.ID
	m["module"] = flattenEndpointTrafficPolicy(&obj.Module)

	return []interface{}{m}
}

func flattenEdgeRouteTrafficPolicyReplaceSlice(objs *[]restapi.EdgeRouteTrafficPolicyReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEdgeRouteTrafficPolicyReplace(&v))
	}
	return sl
}

func expandEdgeRouteTrafficPolicyReplace(in interface{}) *restapi.EdgeRouteTrafficPolicyReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EdgeRouteTrafficPolicyReplace
	if v, ok := m["edge_id"]; ok {
		obj.EdgeID = *expandString(v)
	}
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandEdgeRouteTrafficPolicyReplaceSlice(in interface{}) *[]restapi.EdgeRouteTrafficPolicyReplace {
	var out []restapi.EdgeRouteTrafficPolicyReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEdgeRouteTrafficPolicyReplace(v))
	}
	return &out
}

func flattenTCPEdgeList(obj *restapi.TCPEdgeList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["tcp_edges"] = flattenTCPEdgeSlice(&obj.TCPEdges)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenTCPEdgeListSlice(objs *[]restapi.TCPEdgeList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTCPEdgeList(&v))
	}
	return sl
}

func expandTCPEdgeList(in interface{}) *restapi.TCPEdgeList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TCPEdgeList
	if v, ok := m["tcp_edges"]; ok {
		obj.TCPEdges = *expandTCPEdgeSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandTCPEdgeListSlice(in interface{}) *[]restapi.TCPEdgeList {
	var out []restapi.TCPEdgeList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTCPEdgeList(v))
	}
	return &out
}

func flattenTCPEdgeCreate(obj *restapi.TCPEdgeCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["hostports"] = obj.Hostports
	m["backend"] = flattenEndpointBackendMutate(obj.Backend)
	m["ip_restriction"] = flattenEndpointIPPolicyMutate(obj.IPRestriction)
	m["policy"] = flattenEndpointPolicy(obj.Policy)
	m["traffic_policy"] = flattenEndpointTrafficPolicy(obj.TrafficPolicy)

	return []interface{}{m}
}

func flattenTCPEdgeCreateSlice(objs *[]restapi.TCPEdgeCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTCPEdgeCreate(&v))
	}
	return sl
}

func expandTCPEdgeCreate(in interface{}) *restapi.TCPEdgeCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TCPEdgeCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["hostports"]; ok {
		obj.Hostports = expandStringSlice(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackendMutate(v)
	}
	if v, ok := m["ip_restriction"]; ok {
		obj.IPRestriction = expandEndpointIPPolicyMutate(v)
	}
	if v, ok := m["policy"]; ok {
		obj.Policy = expandEndpointPolicy(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandTCPEdgeCreateSlice(in interface{}) *[]restapi.TCPEdgeCreate {
	var out []restapi.TCPEdgeCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTCPEdgeCreate(v))
	}
	return &out
}

func flattenTCPEdgeUpdate(obj *restapi.TCPEdgeUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["hostports"] = obj.Hostports
	m["backend"] = flattenEndpointBackendMutate(obj.Backend)
	m["ip_restriction"] = flattenEndpointIPPolicyMutate(obj.IPRestriction)
	m["policy"] = flattenEndpointPolicy(obj.Policy)
	m["traffic_policy"] = flattenEndpointTrafficPolicy(obj.TrafficPolicy)

	return []interface{}{m}
}

func flattenTCPEdgeUpdateSlice(objs *[]restapi.TCPEdgeUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTCPEdgeUpdate(&v))
	}
	return sl
}

func expandTCPEdgeUpdate(in interface{}) *restapi.TCPEdgeUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TCPEdgeUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["hostports"]; ok {
		obj.Hostports = expandStringSlice(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackendMutate(v)
	}
	if v, ok := m["ip_restriction"]; ok {
		obj.IPRestriction = expandEndpointIPPolicyMutate(v)
	}
	if v, ok := m["policy"]; ok {
		obj.Policy = expandEndpointPolicy(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandTCPEdgeUpdateSlice(in interface{}) *[]restapi.TCPEdgeUpdate {
	var out []restapi.TCPEdgeUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTCPEdgeUpdate(v))
	}
	return &out
}

func flattenTCPEdge(obj *restapi.TCPEdge) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["created_at"] = obj.CreatedAt
	m["uri"] = obj.URI
	m["hostports"] = obj.Hostports
	m["backend"] = flattenEndpointBackend(obj.Backend)
	m["ip_restriction"] = flattenEndpointIPPolicy(obj.IpRestriction)
	m["policy"] = flattenEndpointPolicy(obj.Policy)
	m["traffic_policy"] = flattenEndpointTrafficPolicy(obj.TrafficPolicy)

	return []interface{}{m}
}

func flattenTCPEdgeSlice(objs *[]restapi.TCPEdge) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTCPEdge(&v))
	}
	return sl
}

func expandTCPEdge(in interface{}) *restapi.TCPEdge {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TCPEdge
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["hostports"]; ok {
		obj.Hostports = expandStringSlice(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackend(v)
	}
	if v, ok := m["ip_restriction"]; ok {
		obj.IpRestriction = expandEndpointIPPolicy(v)
	}
	if v, ok := m["policy"]; ok {
		obj.Policy = expandEndpointPolicy(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandTCPEdgeSlice(in interface{}) *[]restapi.TCPEdge {
	var out []restapi.TCPEdge
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTCPEdge(v))
	}
	return &out
}

func flattenTLSEdgeList(obj *restapi.TLSEdgeList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["tls_edges"] = flattenTLSEdgeSlice(&obj.TLSEdges)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenTLSEdgeListSlice(objs *[]restapi.TLSEdgeList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTLSEdgeList(&v))
	}
	return sl
}

func expandTLSEdgeList(in interface{}) *restapi.TLSEdgeList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TLSEdgeList
	if v, ok := m["tls_edges"]; ok {
		obj.TLSEdges = *expandTLSEdgeSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandTLSEdgeListSlice(in interface{}) *[]restapi.TLSEdgeList {
	var out []restapi.TLSEdgeList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTLSEdgeList(v))
	}
	return &out
}

func flattenTLSEdgeCreate(obj *restapi.TLSEdgeCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["hostports"] = obj.Hostports
	m["backend"] = flattenEndpointBackendMutate(obj.Backend)
	m["ip_restriction"] = flattenEndpointIPPolicyMutate(obj.IPRestriction)
	m["mutual_tls"] = flattenEndpointMutualTLSMutate(obj.MutualTLS)
	m["tls_termination"] = flattenEndpointTLSTermination(obj.TLSTermination)
	m["policy"] = flattenEndpointPolicy(obj.Policy)
	m["traffic_policy"] = flattenEndpointTrafficPolicy(obj.TrafficPolicy)

	return []interface{}{m}
}

func flattenTLSEdgeCreateSlice(objs *[]restapi.TLSEdgeCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTLSEdgeCreate(&v))
	}
	return sl
}

func expandTLSEdgeCreate(in interface{}) *restapi.TLSEdgeCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TLSEdgeCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["hostports"]; ok {
		obj.Hostports = expandStringSlice(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackendMutate(v)
	}
	if v, ok := m["ip_restriction"]; ok {
		obj.IPRestriction = expandEndpointIPPolicyMutate(v)
	}
	if v, ok := m["mutual_tls"]; ok {
		obj.MutualTLS = expandEndpointMutualTLSMutate(v)
	}
	if v, ok := m["tls_termination"]; ok {
		obj.TLSTermination = expandEndpointTLSTermination(v)
	}
	if v, ok := m["policy"]; ok {
		obj.Policy = expandEndpointPolicy(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandTLSEdgeCreateSlice(in interface{}) *[]restapi.TLSEdgeCreate {
	var out []restapi.TLSEdgeCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTLSEdgeCreate(v))
	}
	return &out
}

func flattenTLSEdgeUpdate(obj *restapi.TLSEdgeUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["hostports"] = obj.Hostports
	m["backend"] = flattenEndpointBackendMutate(obj.Backend)
	m["ip_restriction"] = flattenEndpointIPPolicyMutate(obj.IPRestriction)
	m["mutual_tls"] = flattenEndpointMutualTLSMutate(obj.MutualTLS)
	m["tls_termination"] = flattenEndpointTLSTermination(obj.TLSTermination)
	m["policy"] = flattenEndpointPolicy(obj.Policy)
	m["traffic_policy"] = flattenEndpointTrafficPolicy(obj.TrafficPolicy)

	return []interface{}{m}
}

func flattenTLSEdgeUpdateSlice(objs *[]restapi.TLSEdgeUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTLSEdgeUpdate(&v))
	}
	return sl
}

func expandTLSEdgeUpdate(in interface{}) *restapi.TLSEdgeUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TLSEdgeUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["hostports"]; ok {
		obj.Hostports = expandStringSlice(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackendMutate(v)
	}
	if v, ok := m["ip_restriction"]; ok {
		obj.IPRestriction = expandEndpointIPPolicyMutate(v)
	}
	if v, ok := m["mutual_tls"]; ok {
		obj.MutualTLS = expandEndpointMutualTLSMutate(v)
	}
	if v, ok := m["tls_termination"]; ok {
		obj.TLSTermination = expandEndpointTLSTermination(v)
	}
	if v, ok := m["policy"]; ok {
		obj.Policy = expandEndpointPolicy(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandTLSEdgeUpdateSlice(in interface{}) *[]restapi.TLSEdgeUpdate {
	var out []restapi.TLSEdgeUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTLSEdgeUpdate(v))
	}
	return &out
}

func flattenTLSEdge(obj *restapi.TLSEdge) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["created_at"] = obj.CreatedAt
	m["uri"] = obj.URI
	m["hostports"] = obj.Hostports
	m["backend"] = flattenEndpointBackend(obj.Backend)
	m["ip_restriction"] = flattenEndpointIPPolicy(obj.IpRestriction)
	m["mutual_tls"] = flattenEndpointMutualTLS(obj.MutualTls)
	m["tls_termination"] = flattenEndpointTLSTermination(obj.TlsTermination)
	m["policy"] = flattenEndpointPolicy(obj.Policy)
	m["traffic_policy"] = flattenEndpointTrafficPolicy(obj.TrafficPolicy)

	return []interface{}{m}
}

func flattenTLSEdgeSlice(objs *[]restapi.TLSEdge) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTLSEdge(&v))
	}
	return sl
}

func expandTLSEdge(in interface{}) *restapi.TLSEdge {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TLSEdge
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["hostports"]; ok {
		obj.Hostports = expandStringSlice(v)
	}
	if v, ok := m["backend"]; ok {
		obj.Backend = expandEndpointBackend(v)
	}
	if v, ok := m["ip_restriction"]; ok {
		obj.IpRestriction = expandEndpointIPPolicy(v)
	}
	if v, ok := m["mutual_tls"]; ok {
		obj.MutualTls = expandEndpointMutualTLS(v)
	}
	if v, ok := m["tls_termination"]; ok {
		obj.TlsTermination = expandEndpointTLSTermination(v)
	}
	if v, ok := m["policy"]; ok {
		obj.Policy = expandEndpointPolicy(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = expandEndpointTrafficPolicy(v)
	}
	return &obj
}

func expandTLSEdgeSlice(in interface{}) *[]restapi.TLSEdge {
	var out []restapi.TLSEdge
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTLSEdge(v))
	}
	return &out
}

func flattenEndpoint(obj *restapi.Endpoint) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["region"] = obj.Region
	m["created_at"] = obj.CreatedAt
	m["updated_at"] = obj.UpdatedAt
	m["public_url"] = obj.PublicURL
	m["proto"] = obj.Proto
	m["scheme"] = obj.Scheme
	m["hostport"] = obj.Hostport
	m["host"] = obj.Host
	m["port"] = obj.Port
	m["type"] = obj.Type
	m["metadata"] = obj.Metadata
	m["description"] = obj.Description
	m["domain"] = flattenRef(obj.Domain)
	m["tcp_addr"] = flattenRef(obj.TCPAddr)
	m["tunnel"] = flattenRef(obj.Tunnel)
	m["edge"] = flattenRef(obj.Edge)
	m["upstream_url"] = obj.UpstreamURL
	m["upstream_proto"] = obj.UpstreamProto
	m["url"] = obj.URL
	m["principal_id"] = flattenRef(obj.PrincipalID)
	m["traffic_policy"] = obj.TrafficPolicy
	m["bindings"] = obj.Bindings
	m["tunnel_session"] = flattenRef(obj.TunnelSession)
	m["uri"] = obj.URI

	return []interface{}{m}
}

func flattenEndpointSlice(objs *[]restapi.Endpoint) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpoint(&v))
	}
	return sl
}

func expandEndpoint(in interface{}) *restapi.Endpoint {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.Endpoint
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["region"]; ok {
		obj.Region = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["updated_at"]; ok {
		obj.UpdatedAt = *expandString(v)
	}
	if v, ok := m["public_url"]; ok {
		obj.PublicURL = *expandString(v)
	}
	if v, ok := m["proto"]; ok {
		obj.Proto = *expandString(v)
	}
	if v, ok := m["scheme"]; ok {
		obj.Scheme = *expandString(v)
	}
	if v, ok := m["hostport"]; ok {
		obj.Hostport = *expandString(v)
	}
	if v, ok := m["host"]; ok {
		obj.Host = *expandString(v)
	}
	if v, ok := m["port"]; ok {
		obj.Port = *expandInt64(v)
	}
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["domain"]; ok {
		obj.Domain = expandRef(v)
	}
	if v, ok := m["tcp_addr"]; ok {
		obj.TCPAddr = expandRef(v)
	}
	if v, ok := m["tunnel"]; ok {
		obj.Tunnel = expandRef(v)
	}
	if v, ok := m["edge"]; ok {
		obj.Edge = expandRef(v)
	}
	if v, ok := m["upstream_url"]; ok {
		obj.UpstreamURL = *expandString(v)
	}
	if v, ok := m["upstream_proto"]; ok {
		obj.UpstreamProto = *expandString(v)
	}
	if v, ok := m["url"]; ok {
		obj.URL = *expandString(v)
	}
	if v, ok := m["principal_id"]; ok {
		obj.PrincipalID = expandRef(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = *expandString(v)
	}
	if v, ok := m["bindings"]; ok {
		obj.Bindings = expandStringSlice(v)
	}
	if v, ok := m["tunnel_session"]; ok {
		obj.TunnelSession = expandRef(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	return &obj
}

func expandEndpointSlice(in interface{}) *[]restapi.Endpoint {
	var out []restapi.Endpoint
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpoint(v))
	}
	return &out
}

func flattenEndpointList(obj *restapi.EndpointList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["endpoints"] = flattenEndpointSlice(&obj.Endpoints)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenEndpointListSlice(objs *[]restapi.EndpointList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointList(&v))
	}
	return sl
}

func expandEndpointList(in interface{}) *restapi.EndpointList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointList
	if v, ok := m["endpoints"]; ok {
		obj.Endpoints = *expandEndpointSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandEndpointListSlice(in interface{}) *[]restapi.EndpointList {
	var out []restapi.EndpointList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointList(v))
	}
	return &out
}

func flattenEndpointCreate(obj *restapi.EndpointCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["url"] = obj.URL
	m["type"] = obj.Type
	m["traffic_policy"] = obj.TrafficPolicy
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["bindings"] = obj.Bindings

	return []interface{}{m}
}

func flattenEndpointCreateSlice(objs *[]restapi.EndpointCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointCreate(&v))
	}
	return sl
}

func expandEndpointCreate(in interface{}) *restapi.EndpointCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointCreate
	if v, ok := m["url"]; ok {
		obj.URL = *expandString(v)
	}
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["bindings"]; ok {
		obj.Bindings = expandStringSlice(v)
	}
	return &obj
}

func expandEndpointCreateSlice(in interface{}) *[]restapi.EndpointCreate {
	var out []restapi.EndpointCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointCreate(v))
	}
	return &out
}

func flattenEndpointUpdate(obj *restapi.EndpointUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["url"] = obj.Url
	m["traffic_policy"] = obj.TrafficPolicy
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["bindings"] = obj.Bindings

	return []interface{}{m}
}

func flattenEndpointUpdateSlice(objs *[]restapi.EndpointUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointUpdate(&v))
	}
	return sl
}

func expandEndpointUpdate(in interface{}) *restapi.EndpointUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["url"]; ok {
		obj.Url = expandString(v)
	}
	if v, ok := m["traffic_policy"]; ok {
		obj.TrafficPolicy = expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["bindings"]; ok {
		obj.Bindings = expandStringSlice(v)
	}
	return &obj
}

func expandEndpointUpdateSlice(in interface{}) *[]restapi.EndpointUpdate {
	var out []restapi.EndpointUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointUpdate(v))
	}
	return &out
}

func flattenAgentSessionEvent(obj *restapi.AgentSessionEvent) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["session"] = flattenRef(&obj.Session)
	m["credential"] = flattenRef(obj.Credential)
	m["agent_ip"] = obj.AgentIP
	m["ingress_server_ip"] = obj.IngressServerIP
	m["region"] = obj.Region
	m["ingress_hostname"] = obj.IngressHostname
	m["user_agent"] = obj.UserAgent
	m["metadata"] = obj.Metadata
	m["os"] = obj.OS
	m["arch"] = obj.Arch
	m["transport"] = obj.Transport
	m["started_at"] = obj.StartedAt
	m["expires_at"] = obj.ExpiresAt
	m["stopped_at"] = obj.StoppedAt
	m["deprecated"] = flattenAgentDeprecated(obj.Deprecated)
	m["error"] = obj.Error

	return []interface{}{m}
}

func flattenAgentSessionEventSlice(objs *[]restapi.AgentSessionEvent) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAgentSessionEvent(&v))
	}
	return sl
}

func expandAgentSessionEvent(in interface{}) *restapi.AgentSessionEvent {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AgentSessionEvent
	if v, ok := m["session"]; ok {
		obj.Session = *expandRef(v)
	}
	if v, ok := m["credential"]; ok {
		obj.Credential = expandRef(v)
	}
	if v, ok := m["agent_ip"]; ok {
		obj.AgentIP = *expandString(v)
	}
	if v, ok := m["ingress_server_ip"]; ok {
		obj.IngressServerIP = *expandString(v)
	}
	if v, ok := m["region"]; ok {
		obj.Region = *expandString(v)
	}
	if v, ok := m["ingress_hostname"]; ok {
		obj.IngressHostname = *expandString(v)
	}
	if v, ok := m["user_agent"]; ok {
		obj.UserAgent = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["os"]; ok {
		obj.OS = *expandString(v)
	}
	if v, ok := m["arch"]; ok {
		obj.Arch = *expandString(v)
	}
	if v, ok := m["transport"]; ok {
		obj.Transport = *expandString(v)
	}
	if v, ok := m["started_at"]; ok {
		obj.StartedAt = *expandString(v)
	}
	if v, ok := m["expires_at"]; ok {
		obj.ExpiresAt = expandString(v)
	}
	if v, ok := m["stopped_at"]; ok {
		obj.StoppedAt = expandString(v)
	}
	if v, ok := m["deprecated"]; ok {
		obj.Deprecated = expandAgentDeprecated(v)
	}
	if v, ok := m["error"]; ok {
		obj.Error = *expandString(v)
	}
	return &obj
}

func expandAgentSessionEventSlice(in interface{}) *[]restapi.AgentSessionEvent {
	var out []restapi.AgentSessionEvent
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAgentSessionEvent(v))
	}
	return &out
}

func flattenAgentDeprecated(obj *restapi.AgentDeprecated) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["upcoming_minimum_version"] = obj.UpcomingMinimumVersion
	m["upcoming_enforcement_date"] = obj.UpcomingEnforcementDate
	m["message"] = obj.Message

	return []interface{}{m}
}

func flattenAgentDeprecatedSlice(objs *[]restapi.AgentDeprecated) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAgentDeprecated(&v))
	}
	return sl
}

func expandAgentDeprecated(in interface{}) *restapi.AgentDeprecated {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AgentDeprecated
	if v, ok := m["upcoming_minimum_version"]; ok {
		obj.UpcomingMinimumVersion = *expandString(v)
	}
	if v, ok := m["upcoming_enforcement_date"]; ok {
		obj.UpcomingEnforcementDate = *expandString(v)
	}
	if v, ok := m["message"]; ok {
		obj.Message = *expandString(v)
	}
	return &obj
}

func expandAgentDeprecatedSlice(in interface{}) *[]restapi.AgentDeprecated {
	var out []restapi.AgentDeprecated
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAgentDeprecated(v))
	}
	return &out
}

func flattenEventDestinationCreate(obj *restapi.EventDestinationCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["metadata"] = obj.Metadata
	m["description"] = obj.Description
	m["format"] = obj.Format
	m["target"] = flattenEventTarget(&obj.Target)
	m["verify_with_test_event"] = obj.VerifyWithTestEvent

	return []interface{}{m}
}

func flattenEventDestinationCreateSlice(objs *[]restapi.EventDestinationCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventDestinationCreate(&v))
	}
	return sl
}

func expandEventDestinationCreate(in interface{}) *restapi.EventDestinationCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventDestinationCreate
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["format"]; ok {
		obj.Format = *expandString(v)
	}
	if v, ok := m["target"]; ok {
		obj.Target = *expandEventTarget(v)
	}
	if v, ok := m["verify_with_test_event"]; ok {
		obj.VerifyWithTestEvent = expandBool(v)
	}
	return &obj
}

func expandEventDestinationCreateSlice(in interface{}) *[]restapi.EventDestinationCreate {
	var out []restapi.EventDestinationCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventDestinationCreate(v))
	}
	return &out
}

func flattenEventDestinationUpdate(obj *restapi.EventDestinationUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["metadata"] = obj.Metadata
	m["description"] = obj.Description
	m["format"] = obj.Format
	m["target"] = flattenEventTarget(obj.Target)
	m["verify_with_test_event"] = obj.VerifyWithTestEvent

	return []interface{}{m}
}

func flattenEventDestinationUpdateSlice(objs *[]restapi.EventDestinationUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventDestinationUpdate(&v))
	}
	return sl
}

func expandEventDestinationUpdate(in interface{}) *restapi.EventDestinationUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventDestinationUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["format"]; ok {
		obj.Format = expandString(v)
	}
	if v, ok := m["target"]; ok {
		obj.Target = expandEventTarget(v)
	}
	if v, ok := m["verify_with_test_event"]; ok {
		obj.VerifyWithTestEvent = expandBool(v)
	}
	return &obj
}

func expandEventDestinationUpdateSlice(in interface{}) *[]restapi.EventDestinationUpdate {
	var out []restapi.EventDestinationUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventDestinationUpdate(v))
	}
	return &out
}

func flattenEventDestination(obj *restapi.EventDestination) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["metadata"] = obj.Metadata
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["format"] = obj.Format
	m["target"] = flattenEventTarget(&obj.Target)
	m["uri"] = obj.URI

	return []interface{}{m}
}

func flattenEventDestinationSlice(objs *[]restapi.EventDestination) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventDestination(&v))
	}
	return sl
}

func expandEventDestination(in interface{}) *restapi.EventDestination {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventDestination
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["format"]; ok {
		obj.Format = *expandString(v)
	}
	if v, ok := m["target"]; ok {
		obj.Target = *expandEventTarget(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	return &obj
}

func expandEventDestinationSlice(in interface{}) *[]restapi.EventDestination {
	var out []restapi.EventDestination
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventDestination(v))
	}
	return &out
}

func flattenEventDestinationList(obj *restapi.EventDestinationList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["event_destinations"] = flattenEventDestinationSlice(&obj.EventDestinations)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenEventDestinationListSlice(objs *[]restapi.EventDestinationList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventDestinationList(&v))
	}
	return sl
}

func expandEventDestinationList(in interface{}) *restapi.EventDestinationList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventDestinationList
	if v, ok := m["event_destinations"]; ok {
		obj.EventDestinations = *expandEventDestinationSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandEventDestinationListSlice(in interface{}) *[]restapi.EventDestinationList {
	var out []restapi.EventDestinationList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventDestinationList(v))
	}
	return &out
}

func flattenEventTarget(obj *restapi.EventTarget) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["firehose"] = flattenEventTargetFirehose(obj.Firehose)
	m["kinesis"] = flattenEventTargetKinesis(obj.Kinesis)
	m["cloudwatch_logs"] = flattenEventTargetCloudwatchLogs(obj.CloudwatchLogs)
	m["debug"] = flattenEventTargetDebug(obj.Debug)
	m["datadog"] = flattenEventTargetDatadog(obj.Datadog)
	m["azure_logs_ingestion"] = flattenEventTargetAzureLogsIngestion(obj.AzureLogsIngestion)

	return []interface{}{m}
}

func flattenEventTargetSlice(objs *[]restapi.EventTarget) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventTarget(&v))
	}
	return sl
}

func expandEventTarget(in interface{}) *restapi.EventTarget {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventTarget
	if v, ok := m["firehose"]; ok {
		obj.Firehose = expandEventTargetFirehose(v)
	}
	if v, ok := m["kinesis"]; ok {
		obj.Kinesis = expandEventTargetKinesis(v)
	}
	if v, ok := m["cloudwatch_logs"]; ok {
		obj.CloudwatchLogs = expandEventTargetCloudwatchLogs(v)
	}
	if v, ok := m["debug"]; ok {
		obj.Debug = expandEventTargetDebug(v)
	}
	if v, ok := m["datadog"]; ok {
		obj.Datadog = expandEventTargetDatadog(v)
	}
	if v, ok := m["azure_logs_ingestion"]; ok {
		obj.AzureLogsIngestion = expandEventTargetAzureLogsIngestion(v)
	}
	return &obj
}

func expandEventTargetSlice(in interface{}) *[]restapi.EventTarget {
	var out []restapi.EventTarget
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventTarget(v))
	}
	return &out
}

func flattenEventTargetFirehose(obj *restapi.EventTargetFirehose) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["auth"] = flattenAWSAuth(&obj.Auth)
	m["delivery_stream_arn"] = obj.DeliveryStreamARN

	return []interface{}{m}
}

func flattenEventTargetFirehoseSlice(objs *[]restapi.EventTargetFirehose) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventTargetFirehose(&v))
	}
	return sl
}

func expandEventTargetFirehose(in interface{}) *restapi.EventTargetFirehose {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventTargetFirehose
	if v, ok := m["auth"]; ok {
		obj.Auth = *expandAWSAuth(v)
	}
	if v, ok := m["delivery_stream_arn"]; ok {
		obj.DeliveryStreamARN = *expandString(v)
	}
	return &obj
}

func expandEventTargetFirehoseSlice(in interface{}) *[]restapi.EventTargetFirehose {
	var out []restapi.EventTargetFirehose
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventTargetFirehose(v))
	}
	return &out
}

func flattenEventTargetKinesis(obj *restapi.EventTargetKinesis) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["auth"] = flattenAWSAuth(&obj.Auth)
	m["stream_arn"] = obj.StreamARN

	return []interface{}{m}
}

func flattenEventTargetKinesisSlice(objs *[]restapi.EventTargetKinesis) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventTargetKinesis(&v))
	}
	return sl
}

func expandEventTargetKinesis(in interface{}) *restapi.EventTargetKinesis {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventTargetKinesis
	if v, ok := m["auth"]; ok {
		obj.Auth = *expandAWSAuth(v)
	}
	if v, ok := m["stream_arn"]; ok {
		obj.StreamARN = *expandString(v)
	}
	return &obj
}

func expandEventTargetKinesisSlice(in interface{}) *[]restapi.EventTargetKinesis {
	var out []restapi.EventTargetKinesis
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventTargetKinesis(v))
	}
	return &out
}

func flattenEventTargetCloudwatchLogs(obj *restapi.EventTargetCloudwatchLogs) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["auth"] = flattenAWSAuth(&obj.Auth)
	m["log_group_arn"] = obj.LogGroupARN

	return []interface{}{m}
}

func flattenEventTargetCloudwatchLogsSlice(objs *[]restapi.EventTargetCloudwatchLogs) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventTargetCloudwatchLogs(&v))
	}
	return sl
}

func expandEventTargetCloudwatchLogs(in interface{}) *restapi.EventTargetCloudwatchLogs {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventTargetCloudwatchLogs
	if v, ok := m["auth"]; ok {
		obj.Auth = *expandAWSAuth(v)
	}
	if v, ok := m["log_group_arn"]; ok {
		obj.LogGroupARN = *expandString(v)
	}
	return &obj
}

func expandEventTargetCloudwatchLogsSlice(in interface{}) *[]restapi.EventTargetCloudwatchLogs {
	var out []restapi.EventTargetCloudwatchLogs
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventTargetCloudwatchLogs(v))
	}
	return &out
}

func flattenEventTargetS3(obj *restapi.EventTargetS3) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["auth"] = flattenAWSAuth(&obj.Auth)
	m["bucket_arn"] = obj.BucketARN
	m["object_prefix"] = obj.ObjectPrefix
	m["compression"] = obj.Compression
	m["max_file_size"] = obj.MaxFileSize
	m["max_file_age"] = obj.MaxFileAge

	return []interface{}{m}
}

func flattenEventTargetS3Slice(objs *[]restapi.EventTargetS3) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventTargetS3(&v))
	}
	return sl
}

func expandEventTargetS3(in interface{}) *restapi.EventTargetS3 {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventTargetS3
	if v, ok := m["auth"]; ok {
		obj.Auth = *expandAWSAuth(v)
	}
	if v, ok := m["bucket_arn"]; ok {
		obj.BucketARN = *expandString(v)
	}
	if v, ok := m["object_prefix"]; ok {
		obj.ObjectPrefix = *expandString(v)
	}
	if v, ok := m["compression"]; ok {
		obj.Compression = *expandBool(v)
	}
	if v, ok := m["max_file_size"]; ok {
		obj.MaxFileSize = *expandInt64(v)
	}
	if v, ok := m["max_file_age"]; ok {
		obj.MaxFileAge = *expandInt64(v)
	}
	return &obj
}

func expandEventTargetS3Slice(in interface{}) *[]restapi.EventTargetS3 {
	var out []restapi.EventTargetS3
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventTargetS3(v))
	}
	return &out
}

func flattenEventTargetDebug(obj *restapi.EventTargetDebug) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["log"] = obj.Log
	m["callback_url"] = obj.CallbackURL

	return []interface{}{m}
}

func flattenEventTargetDebugSlice(objs *[]restapi.EventTargetDebug) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventTargetDebug(&v))
	}
	return sl
}

func expandEventTargetDebug(in interface{}) *restapi.EventTargetDebug {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventTargetDebug
	if v, ok := m["log"]; ok {
		obj.Log = *expandBool(v)
	}
	if v, ok := m["callback_url"]; ok {
		obj.CallbackURL = *expandString(v)
	}
	return &obj
}

func expandEventTargetDebugSlice(in interface{}) *[]restapi.EventTargetDebug {
	var out []restapi.EventTargetDebug
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventTargetDebug(v))
	}
	return &out
}

func flattenEventTargetDatadog(obj *restapi.EventTargetDatadog) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["api_key"] = obj.ApiKey
	m["ddtags"] = obj.Ddtags
	m["service"] = obj.Service
	m["ddsite"] = obj.Ddsite

	return []interface{}{m}
}

func flattenEventTargetDatadogSlice(objs *[]restapi.EventTargetDatadog) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventTargetDatadog(&v))
	}
	return sl
}

func expandEventTargetDatadog(in interface{}) *restapi.EventTargetDatadog {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventTargetDatadog
	if v, ok := m["api_key"]; ok {
		obj.ApiKey = expandString(v)
	}
	if v, ok := m["ddtags"]; ok {
		obj.Ddtags = expandString(v)
	}
	if v, ok := m["service"]; ok {
		obj.Service = expandString(v)
	}
	if v, ok := m["ddsite"]; ok {
		obj.Ddsite = expandString(v)
	}
	return &obj
}

func expandEventTargetDatadogSlice(in interface{}) *[]restapi.EventTargetDatadog {
	var out []restapi.EventTargetDatadog
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventTargetDatadog(v))
	}
	return &out
}

func flattenEventTargetAzureLogsIngestion(obj *restapi.EventTargetAzureLogsIngestion) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["tenant_id"] = obj.TenantId
	m["client_id"] = obj.ClientId
	m["client_secret"] = obj.ClientSecret
	m["logs_ingestion_uri"] = obj.LogsIngestionURI
	m["data_collection_rule_id"] = obj.DataCollectionRuleId
	m["data_collection_stream_name"] = obj.DataCollectionStreamName

	return []interface{}{m}
}

func flattenEventTargetAzureLogsIngestionSlice(objs *[]restapi.EventTargetAzureLogsIngestion) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventTargetAzureLogsIngestion(&v))
	}
	return sl
}

func expandEventTargetAzureLogsIngestion(in interface{}) *restapi.EventTargetAzureLogsIngestion {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventTargetAzureLogsIngestion
	if v, ok := m["tenant_id"]; ok {
		obj.TenantId = *expandString(v)
	}
	if v, ok := m["client_id"]; ok {
		obj.ClientId = *expandString(v)
	}
	if v, ok := m["client_secret"]; ok {
		obj.ClientSecret = *expandString(v)
	}
	if v, ok := m["logs_ingestion_uri"]; ok {
		obj.LogsIngestionURI = *expandString(v)
	}
	if v, ok := m["data_collection_rule_id"]; ok {
		obj.DataCollectionRuleId = *expandString(v)
	}
	if v, ok := m["data_collection_stream_name"]; ok {
		obj.DataCollectionStreamName = *expandString(v)
	}
	return &obj
}

func expandEventTargetAzureLogsIngestionSlice(in interface{}) *[]restapi.EventTargetAzureLogsIngestion {
	var out []restapi.EventTargetAzureLogsIngestion
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventTargetAzureLogsIngestion(v))
	}
	return &out
}

func flattenAWSAuth(obj *restapi.AWSAuth) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["role"] = flattenAWSRole(obj.Role)
	m["creds"] = flattenAWSCredentials(obj.Creds)

	return []interface{}{m}
}

func flattenAWSAuthSlice(objs *[]restapi.AWSAuth) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAWSAuth(&v))
	}
	return sl
}

func expandAWSAuth(in interface{}) *restapi.AWSAuth {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AWSAuth
	if v, ok := m["role"]; ok {
		obj.Role = expandAWSRole(v)
	}
	if v, ok := m["creds"]; ok {
		obj.Creds = expandAWSCredentials(v)
	}
	return &obj
}

func expandAWSAuthSlice(in interface{}) *[]restapi.AWSAuth {
	var out []restapi.AWSAuth
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAWSAuth(v))
	}
	return &out
}

func flattenAWSRole(obj *restapi.AWSRole) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["role_arn"] = obj.RoleARN

	return []interface{}{m}
}

func flattenAWSRoleSlice(objs *[]restapi.AWSRole) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAWSRole(&v))
	}
	return sl
}

func expandAWSRole(in interface{}) *restapi.AWSRole {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AWSRole
	if v, ok := m["role_arn"]; ok {
		obj.RoleARN = *expandString(v)
	}
	return &obj
}

func expandAWSRoleSlice(in interface{}) *[]restapi.AWSRole {
	var out []restapi.AWSRole
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAWSRole(v))
	}
	return &out
}

func flattenAWSCredentials(obj *restapi.AWSCredentials) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["aws_access_key_id"] = obj.AWSAccessKeyID
	m["aws_secret_access_key"] = obj.AWSSecretAccessKey

	return []interface{}{m}
}

func flattenAWSCredentialsSlice(objs *[]restapi.AWSCredentials) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenAWSCredentials(&v))
	}
	return sl
}

func expandAWSCredentials(in interface{}) *restapi.AWSCredentials {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.AWSCredentials
	if v, ok := m["aws_access_key_id"]; ok {
		obj.AWSAccessKeyID = *expandString(v)
	}
	if v, ok := m["aws_secret_access_key"]; ok {
		obj.AWSSecretAccessKey = expandString(v)
	}
	return &obj
}

func expandAWSCredentialsSlice(in interface{}) *[]restapi.AWSCredentials {
	var out []restapi.AWSCredentials
	for _, v := range in.([]interface{}) {
		out = append(out, *expandAWSCredentials(v))
	}
	return &out
}

func flattenSentEvent(obj *restapi.SentEvent) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["event_id"] = obj.EventID

	return []interface{}{m}
}

func flattenSentEventSlice(objs *[]restapi.SentEvent) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSentEvent(&v))
	}
	return sl
}

func expandSentEvent(in interface{}) *restapi.SentEvent {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SentEvent
	if v, ok := m["event_id"]; ok {
		obj.EventID = *expandString(v)
	}
	return &obj
}

func expandSentEventSlice(in interface{}) *[]restapi.SentEvent {
	var out []restapi.SentEvent
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSentEvent(v))
	}
	return &out
}

func flattenEventSubscriptionCreate(obj *restapi.EventSubscriptionCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["metadata"] = obj.Metadata
	m["description"] = obj.Description
	m["sources"] = flattenEventSourceReplaceSlice(&obj.Sources)
	m["destination_ids"] = obj.DestinationIDs

	return []interface{}{m}
}

func flattenEventSubscriptionCreateSlice(objs *[]restapi.EventSubscriptionCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSubscriptionCreate(&v))
	}
	return sl
}

func expandEventSubscriptionCreate(in interface{}) *restapi.EventSubscriptionCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSubscriptionCreate
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["sources"]; ok {
		obj.Sources = *expandEventSourceReplaceSlice(v)
	}
	if v, ok := m["destination_ids"]; ok {
		obj.DestinationIDs = *expandStringSlice(v)
	}
	return &obj
}

func expandEventSubscriptionCreateSlice(in interface{}) *[]restapi.EventSubscriptionCreate {
	var out []restapi.EventSubscriptionCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSubscriptionCreate(v))
	}
	return &out
}

func flattenEventSubscriptionUpdate(obj *restapi.EventSubscriptionUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["metadata"] = obj.Metadata
	m["description"] = obj.Description
	m["sources"] = flattenEventSourceReplaceSlice(obj.Sources)
	m["destination_ids"] = obj.DestinationIDs

	return []interface{}{m}
}

func flattenEventSubscriptionUpdateSlice(objs *[]restapi.EventSubscriptionUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSubscriptionUpdate(&v))
	}
	return sl
}

func expandEventSubscriptionUpdate(in interface{}) *restapi.EventSubscriptionUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSubscriptionUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["sources"]; ok {
		obj.Sources = expandEventSourceReplaceSlice(v)
	}
	if v, ok := m["destination_ids"]; ok {
		obj.DestinationIDs = expandStringSlice(v)
	}
	return &obj
}

func expandEventSubscriptionUpdateSlice(in interface{}) *[]restapi.EventSubscriptionUpdate {
	var out []restapi.EventSubscriptionUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSubscriptionUpdate(v))
	}
	return &out
}

func flattenEventSubscriptionList(obj *restapi.EventSubscriptionList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["event_subscriptions"] = flattenEventSubscriptionSlice(&obj.EventSubscriptions)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenEventSubscriptionListSlice(objs *[]restapi.EventSubscriptionList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSubscriptionList(&v))
	}
	return sl
}

func expandEventSubscriptionList(in interface{}) *restapi.EventSubscriptionList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSubscriptionList
	if v, ok := m["event_subscriptions"]; ok {
		obj.EventSubscriptions = *expandEventSubscriptionSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandEventSubscriptionListSlice(in interface{}) *[]restapi.EventSubscriptionList {
	var out []restapi.EventSubscriptionList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSubscriptionList(v))
	}
	return &out
}

func flattenEventSubscription(obj *restapi.EventSubscription) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["metadata"] = obj.Metadata
	m["description"] = obj.Description
	m["sources"] = flattenEventSourceSlice(&obj.Sources)
	m["destinations"] = flattenRefSlice(&obj.Destinations)

	return []interface{}{m}
}

func flattenEventSubscriptionSlice(objs *[]restapi.EventSubscription) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSubscription(&v))
	}
	return sl
}

func expandEventSubscription(in interface{}) *restapi.EventSubscription {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSubscription
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["sources"]; ok {
		obj.Sources = *expandEventSourceSlice(v)
	}
	if v, ok := m["destinations"]; ok {
		obj.Destinations = *expandRefSlice(v)
	}
	return &obj
}

func expandEventSubscriptionSlice(in interface{}) *[]restapi.EventSubscription {
	var out []restapi.EventSubscription
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSubscription(v))
	}
	return &out
}

func flattenEventSourceReplace(obj *restapi.EventSourceReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["type"] = obj.Type
	m["filter"] = obj.Filter
	m["fields"] = obj.Fields

	return []interface{}{m}
}

func flattenEventSourceReplaceSlice(objs *[]restapi.EventSourceReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSourceReplace(&v))
	}
	return sl
}

func expandEventSourceReplace(in interface{}) *restapi.EventSourceReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSourceReplace
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["filter"]; ok {
		obj.Filter = *expandString(v)
	}
	if v, ok := m["fields"]; ok {
		obj.Fields = *expandStringSlice(v)
	}
	return &obj
}

func expandEventSourceReplaceSlice(in interface{}) *[]restapi.EventSourceReplace {
	var out []restapi.EventSourceReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSourceReplace(v))
	}
	return &out
}

func flattenEventSource(obj *restapi.EventSource) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["type"] = obj.Type
	m["filter"] = obj.Filter
	m["fields"] = obj.Fields
	m["uri"] = obj.URI

	return []interface{}{m}
}

func flattenEventSourceSlice(objs *[]restapi.EventSource) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSource(&v))
	}
	return sl
}

func expandEventSource(in interface{}) *restapi.EventSource {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSource
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["filter"]; ok {
		obj.Filter = *expandString(v)
	}
	if v, ok := m["fields"]; ok {
		obj.Fields = *expandStringSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	return &obj
}

func expandEventSourceSlice(in interface{}) *[]restapi.EventSource {
	var out []restapi.EventSource
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSource(v))
	}
	return &out
}

func flattenEventSourceList(obj *restapi.EventSourceList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["sources"] = flattenEventSourceSlice(&obj.Sources)
	m["uri"] = obj.URI

	return []interface{}{m}
}

func flattenEventSourceListSlice(objs *[]restapi.EventSourceList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSourceList(&v))
	}
	return sl
}

func expandEventSourceList(in interface{}) *restapi.EventSourceList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSourceList
	if v, ok := m["sources"]; ok {
		obj.Sources = *expandEventSourceSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	return &obj
}

func expandEventSourceListSlice(in interface{}) *[]restapi.EventSourceList {
	var out []restapi.EventSourceList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSourceList(v))
	}
	return &out
}

func flattenEventSourceCreate(obj *restapi.EventSourceCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["subscription_id"] = obj.SubscriptionID
	m["type"] = obj.Type
	m["filter"] = obj.Filter
	m["fields"] = obj.Fields

	return []interface{}{m}
}

func flattenEventSourceCreateSlice(objs *[]restapi.EventSourceCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSourceCreate(&v))
	}
	return sl
}

func expandEventSourceCreate(in interface{}) *restapi.EventSourceCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSourceCreate
	if v, ok := m["subscription_id"]; ok {
		obj.SubscriptionID = *expandString(v)
	}
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["filter"]; ok {
		obj.Filter = *expandString(v)
	}
	if v, ok := m["fields"]; ok {
		obj.Fields = *expandStringSlice(v)
	}
	return &obj
}

func expandEventSourceCreateSlice(in interface{}) *[]restapi.EventSourceCreate {
	var out []restapi.EventSourceCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSourceCreate(v))
	}
	return &out
}

func flattenEventSourceUpdate(obj *restapi.EventSourceUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["subscription_id"] = obj.SubscriptionID
	m["type"] = obj.Type
	m["filter"] = obj.Filter
	m["fields"] = obj.Fields

	return []interface{}{m}
}

func flattenEventSourceUpdateSlice(objs *[]restapi.EventSourceUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSourceUpdate(&v))
	}
	return sl
}

func expandEventSourceUpdate(in interface{}) *restapi.EventSourceUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSourceUpdate
	if v, ok := m["subscription_id"]; ok {
		obj.SubscriptionID = *expandString(v)
	}
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["filter"]; ok {
		obj.Filter = expandString(v)
	}
	if v, ok := m["fields"]; ok {
		obj.Fields = expandStringSlice(v)
	}
	return &obj
}

func expandEventSourceUpdateSlice(in interface{}) *[]restapi.EventSourceUpdate {
	var out []restapi.EventSourceUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSourceUpdate(v))
	}
	return &out
}

func flattenEventSourceItem(obj *restapi.EventSourceItem) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["subscription_id"] = obj.SubscriptionID
	m["type"] = obj.Type

	return []interface{}{m}
}

func flattenEventSourceItemSlice(objs *[]restapi.EventSourceItem) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSourceItem(&v))
	}
	return sl
}

func expandEventSourceItem(in interface{}) *restapi.EventSourceItem {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSourceItem
	if v, ok := m["subscription_id"]; ok {
		obj.SubscriptionID = *expandString(v)
	}
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	return &obj
}

func expandEventSourceItemSlice(in interface{}) *[]restapi.EventSourceItem {
	var out []restapi.EventSourceItem
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSourceItem(v))
	}
	return &out
}

func flattenEventSourcePaging(obj *restapi.EventSourcePaging) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["subscription_id"] = obj.SubscriptionID

	return []interface{}{m}
}

func flattenEventSourcePagingSlice(objs *[]restapi.EventSourcePaging) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEventSourcePaging(&v))
	}
	return sl
}

func expandEventSourcePaging(in interface{}) *restapi.EventSourcePaging {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EventSourcePaging
	if v, ok := m["subscription_id"]; ok {
		obj.SubscriptionID = *expandString(v)
	}
	return &obj
}

func expandEventSourcePagingSlice(in interface{}) *[]restapi.EventSourcePaging {
	var out []restapi.EventSourcePaging
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEventSourcePaging(v))
	}
	return &out
}

func flattenIPPolicyCreate(obj *restapi.IPPolicyCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["action"] = obj.Action

	return []interface{}{m}
}

func flattenIPPolicyCreateSlice(objs *[]restapi.IPPolicyCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPPolicyCreate(&v))
	}
	return sl
}

func expandIPPolicyCreate(in interface{}) *restapi.IPPolicyCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPPolicyCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["action"]; ok {
		obj.Action = expandString(v)
	}
	return &obj
}

func expandIPPolicyCreateSlice(in interface{}) *[]restapi.IPPolicyCreate {
	var out []restapi.IPPolicyCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPPolicyCreate(v))
	}
	return &out
}

func flattenIPPolicyUpdate(obj *restapi.IPPolicyUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata

	return []interface{}{m}
}

func flattenIPPolicyUpdateSlice(objs *[]restapi.IPPolicyUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPPolicyUpdate(&v))
	}
	return sl
}

func expandIPPolicyUpdate(in interface{}) *restapi.IPPolicyUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPPolicyUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	return &obj
}

func expandIPPolicyUpdateSlice(in interface{}) *[]restapi.IPPolicyUpdate {
	var out []restapi.IPPolicyUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPPolicyUpdate(v))
	}
	return &out
}

func flattenIPPolicy(obj *restapi.IPPolicy) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["action"] = obj.Action

	return []interface{}{m}
}

func flattenIPPolicySlice(objs *[]restapi.IPPolicy) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPPolicy(&v))
	}
	return sl
}

func expandIPPolicy(in interface{}) *restapi.IPPolicy {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPPolicy
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["action"]; ok {
		obj.Action = expandString(v)
	}
	return &obj
}

func expandIPPolicySlice(in interface{}) *[]restapi.IPPolicy {
	var out []restapi.IPPolicy
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPPolicy(v))
	}
	return &out
}

func flattenIPPolicyList(obj *restapi.IPPolicyList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["ip_policies"] = flattenIPPolicySlice(&obj.IPPolicies)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenIPPolicyListSlice(objs *[]restapi.IPPolicyList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPPolicyList(&v))
	}
	return sl
}

func expandIPPolicyList(in interface{}) *restapi.IPPolicyList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPPolicyList
	if v, ok := m["ip_policies"]; ok {
		obj.IPPolicies = *expandIPPolicySlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandIPPolicyListSlice(in interface{}) *[]restapi.IPPolicyList {
	var out []restapi.IPPolicyList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPPolicyList(v))
	}
	return &out
}

func flattenIPPolicyRuleCreate(obj *restapi.IPPolicyRuleCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["cidr"] = obj.CIDR
	m["ip_policy_id"] = obj.IPPolicyID
	m["action"] = obj.Action

	return []interface{}{m}
}

func flattenIPPolicyRuleCreateSlice(objs *[]restapi.IPPolicyRuleCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPPolicyRuleCreate(&v))
	}
	return sl
}

func expandIPPolicyRuleCreate(in interface{}) *restapi.IPPolicyRuleCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPPolicyRuleCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["cidr"]; ok {
		obj.CIDR = *expandString(v)
	}
	if v, ok := m["ip_policy_id"]; ok {
		obj.IPPolicyID = *expandString(v)
	}
	if v, ok := m["action"]; ok {
		obj.Action = expandString(v)
	}
	return &obj
}

func expandIPPolicyRuleCreateSlice(in interface{}) *[]restapi.IPPolicyRuleCreate {
	var out []restapi.IPPolicyRuleCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPPolicyRuleCreate(v))
	}
	return &out
}

func flattenIPPolicyRuleUpdate(obj *restapi.IPPolicyRuleUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["cidr"] = obj.CIDR

	return []interface{}{m}
}

func flattenIPPolicyRuleUpdateSlice(objs *[]restapi.IPPolicyRuleUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPPolicyRuleUpdate(&v))
	}
	return sl
}

func expandIPPolicyRuleUpdate(in interface{}) *restapi.IPPolicyRuleUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPPolicyRuleUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["cidr"]; ok {
		obj.CIDR = expandString(v)
	}
	return &obj
}

func expandIPPolicyRuleUpdateSlice(in interface{}) *[]restapi.IPPolicyRuleUpdate {
	var out []restapi.IPPolicyRuleUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPPolicyRuleUpdate(v))
	}
	return &out
}

func flattenIPPolicyRule(obj *restapi.IPPolicyRule) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["cidr"] = obj.CIDR
	m["ip_policy"] = flattenRef(&obj.IPPolicy)
	m["action"] = obj.Action

	return []interface{}{m}
}

func flattenIPPolicyRuleSlice(objs *[]restapi.IPPolicyRule) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPPolicyRule(&v))
	}
	return sl
}

func expandIPPolicyRule(in interface{}) *restapi.IPPolicyRule {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPPolicyRule
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["cidr"]; ok {
		obj.CIDR = *expandString(v)
	}
	if v, ok := m["ip_policy"]; ok {
		obj.IPPolicy = *expandRef(v)
	}
	if v, ok := m["action"]; ok {
		obj.Action = *expandString(v)
	}
	return &obj
}

func expandIPPolicyRuleSlice(in interface{}) *[]restapi.IPPolicyRule {
	var out []restapi.IPPolicyRule
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPPolicyRule(v))
	}
	return &out
}

func flattenIPPolicyRuleList(obj *restapi.IPPolicyRuleList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["ip_policy_rules"] = flattenIPPolicyRuleSlice(&obj.IPPolicyRules)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenIPPolicyRuleListSlice(objs *[]restapi.IPPolicyRuleList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPPolicyRuleList(&v))
	}
	return sl
}

func expandIPPolicyRuleList(in interface{}) *restapi.IPPolicyRuleList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPPolicyRuleList
	if v, ok := m["ip_policy_rules"]; ok {
		obj.IPPolicyRules = *expandIPPolicyRuleSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandIPPolicyRuleListSlice(in interface{}) *[]restapi.IPPolicyRuleList {
	var out []restapi.IPPolicyRuleList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPPolicyRuleList(v))
	}
	return &out
}

func flattenIPRestrictionCreate(obj *restapi.IPRestrictionCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["enforced"] = obj.Enforced
	m["type"] = obj.Type
	m["ip_policy_ids"] = obj.IPPolicyIDs

	return []interface{}{m}
}

func flattenIPRestrictionCreateSlice(objs *[]restapi.IPRestrictionCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPRestrictionCreate(&v))
	}
	return sl
}

func expandIPRestrictionCreate(in interface{}) *restapi.IPRestrictionCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPRestrictionCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["enforced"]; ok {
		obj.Enforced = *expandBool(v)
	}
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["ip_policy_ids"]; ok {
		obj.IPPolicyIDs = *expandStringSlice(v)
	}
	return &obj
}

func expandIPRestrictionCreateSlice(in interface{}) *[]restapi.IPRestrictionCreate {
	var out []restapi.IPRestrictionCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPRestrictionCreate(v))
	}
	return &out
}

func flattenIPRestrictionUpdate(obj *restapi.IPRestrictionUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["enforced"] = obj.Enforced
	m["ip_policy_ids"] = obj.IPPolicyIDs

	return []interface{}{m}
}

func flattenIPRestrictionUpdateSlice(objs *[]restapi.IPRestrictionUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPRestrictionUpdate(&v))
	}
	return sl
}

func expandIPRestrictionUpdate(in interface{}) *restapi.IPRestrictionUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPRestrictionUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["enforced"]; ok {
		obj.Enforced = expandBool(v)
	}
	if v, ok := m["ip_policy_ids"]; ok {
		obj.IPPolicyIDs = *expandStringSlice(v)
	}
	return &obj
}

func expandIPRestrictionUpdateSlice(in interface{}) *[]restapi.IPRestrictionUpdate {
	var out []restapi.IPRestrictionUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPRestrictionUpdate(v))
	}
	return &out
}

func flattenIPRestriction(obj *restapi.IPRestriction) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["enforced"] = obj.Enforced
	m["type"] = obj.Type
	m["ip_policies"] = flattenRefSlice(&obj.IPPolicies)

	return []interface{}{m}
}

func flattenIPRestrictionSlice(objs *[]restapi.IPRestriction) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPRestriction(&v))
	}
	return sl
}

func expandIPRestriction(in interface{}) *restapi.IPRestriction {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPRestriction
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["enforced"]; ok {
		obj.Enforced = *expandBool(v)
	}
	if v, ok := m["type"]; ok {
		obj.Type = *expandString(v)
	}
	if v, ok := m["ip_policies"]; ok {
		obj.IPPolicies = *expandRefSlice(v)
	}
	return &obj
}

func expandIPRestrictionSlice(in interface{}) *[]restapi.IPRestriction {
	var out []restapi.IPRestriction
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPRestriction(v))
	}
	return &out
}

func flattenIPRestrictionList(obj *restapi.IPRestrictionList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["ip_restrictions"] = flattenIPRestrictionSlice(&obj.IPRestrictions)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenIPRestrictionListSlice(objs *[]restapi.IPRestrictionList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenIPRestrictionList(&v))
	}
	return sl
}

func expandIPRestrictionList(in interface{}) *restapi.IPRestrictionList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.IPRestrictionList
	if v, ok := m["ip_restrictions"]; ok {
		obj.IPRestrictions = *expandIPRestrictionSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandIPRestrictionListSlice(in interface{}) *[]restapi.IPRestrictionList {
	var out []restapi.IPRestrictionList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandIPRestrictionList(v))
	}
	return &out
}

func flattenEndpointBasicAuthReplace(obj *restapi.EndpointBasicAuthReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointBasicAuth(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointBasicAuthReplaceSlice(objs *[]restapi.EndpointBasicAuthReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointBasicAuthReplace(&v))
	}
	return sl
}

func expandEndpointBasicAuthReplace(in interface{}) *restapi.EndpointBasicAuthReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointBasicAuthReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointBasicAuth(v)
	}
	return &obj
}

func expandEndpointBasicAuthReplaceSlice(in interface{}) *[]restapi.EndpointBasicAuthReplace {
	var out []restapi.EndpointBasicAuthReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointBasicAuthReplace(v))
	}
	return &out
}

func flattenEndpointCircuitBreakerReplace(obj *restapi.EndpointCircuitBreakerReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointCircuitBreaker(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointCircuitBreakerReplaceSlice(objs *[]restapi.EndpointCircuitBreakerReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointCircuitBreakerReplace(&v))
	}
	return sl
}

func expandEndpointCircuitBreakerReplace(in interface{}) *restapi.EndpointCircuitBreakerReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointCircuitBreakerReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointCircuitBreaker(v)
	}
	return &obj
}

func expandEndpointCircuitBreakerReplaceSlice(in interface{}) *[]restapi.EndpointCircuitBreakerReplace {
	var out []restapi.EndpointCircuitBreakerReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointCircuitBreakerReplace(v))
	}
	return &out
}

func flattenEndpointCompressionReplace(obj *restapi.EndpointCompressionReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointCompression(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointCompressionReplaceSlice(objs *[]restapi.EndpointCompressionReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointCompressionReplace(&v))
	}
	return sl
}

func expandEndpointCompressionReplace(in interface{}) *restapi.EndpointCompressionReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointCompressionReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointCompression(v)
	}
	return &obj
}

func expandEndpointCompressionReplaceSlice(in interface{}) *[]restapi.EndpointCompressionReplace {
	var out []restapi.EndpointCompressionReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointCompressionReplace(v))
	}
	return &out
}

func flattenEndpointTLSTerminationReplace(obj *restapi.EndpointTLSTerminationReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointTLSTermination(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointTLSTerminationReplaceSlice(objs *[]restapi.EndpointTLSTerminationReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointTLSTerminationReplace(&v))
	}
	return sl
}

func expandEndpointTLSTerminationReplace(in interface{}) *restapi.EndpointTLSTerminationReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointTLSTerminationReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointTLSTermination(v)
	}
	return &obj
}

func expandEndpointTLSTerminationReplaceSlice(in interface{}) *[]restapi.EndpointTLSTerminationReplace {
	var out []restapi.EndpointTLSTerminationReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointTLSTerminationReplace(v))
	}
	return &out
}

func flattenEndpointIPPolicyReplace(obj *restapi.EndpointIPPolicyReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointIPPolicyMutate(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointIPPolicyReplaceSlice(objs *[]restapi.EndpointIPPolicyReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointIPPolicyReplace(&v))
	}
	return sl
}

func expandEndpointIPPolicyReplace(in interface{}) *restapi.EndpointIPPolicyReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointIPPolicyReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointIPPolicyMutate(v)
	}
	return &obj
}

func expandEndpointIPPolicyReplaceSlice(in interface{}) *[]restapi.EndpointIPPolicyReplace {
	var out []restapi.EndpointIPPolicyReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointIPPolicyReplace(v))
	}
	return &out
}

func flattenEndpointMutualTLSReplace(obj *restapi.EndpointMutualTLSReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointMutualTLSMutate(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointMutualTLSReplaceSlice(objs *[]restapi.EndpointMutualTLSReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointMutualTLSReplace(&v))
	}
	return sl
}

func expandEndpointMutualTLSReplace(in interface{}) *restapi.EndpointMutualTLSReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointMutualTLSReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointMutualTLSMutate(v)
	}
	return &obj
}

func expandEndpointMutualTLSReplaceSlice(in interface{}) *[]restapi.EndpointMutualTLSReplace {
	var out []restapi.EndpointMutualTLSReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointMutualTLSReplace(v))
	}
	return &out
}

func flattenEndpointRequestHeadersReplace(obj *restapi.EndpointRequestHeadersReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointRequestHeaders(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointRequestHeadersReplaceSlice(objs *[]restapi.EndpointRequestHeadersReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointRequestHeadersReplace(&v))
	}
	return sl
}

func expandEndpointRequestHeadersReplace(in interface{}) *restapi.EndpointRequestHeadersReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointRequestHeadersReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointRequestHeaders(v)
	}
	return &obj
}

func expandEndpointRequestHeadersReplaceSlice(in interface{}) *[]restapi.EndpointRequestHeadersReplace {
	var out []restapi.EndpointRequestHeadersReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointRequestHeadersReplace(v))
	}
	return &out
}

func flattenEndpointResponseHeadersReplace(obj *restapi.EndpointResponseHeadersReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointResponseHeaders(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointResponseHeadersReplaceSlice(objs *[]restapi.EndpointResponseHeadersReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointResponseHeadersReplace(&v))
	}
	return sl
}

func expandEndpointResponseHeadersReplace(in interface{}) *restapi.EndpointResponseHeadersReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointResponseHeadersReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointResponseHeaders(v)
	}
	return &obj
}

func expandEndpointResponseHeadersReplaceSlice(in interface{}) *[]restapi.EndpointResponseHeadersReplace {
	var out []restapi.EndpointResponseHeadersReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointResponseHeadersReplace(v))
	}
	return &out
}

func flattenEndpointOAuthReplace(obj *restapi.EndpointOAuthReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointOAuth(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointOAuthReplaceSlice(objs *[]restapi.EndpointOAuthReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOAuthReplace(&v))
	}
	return sl
}

func expandEndpointOAuthReplace(in interface{}) *restapi.EndpointOAuthReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOAuthReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointOAuth(v)
	}
	return &obj
}

func expandEndpointOAuthReplaceSlice(in interface{}) *[]restapi.EndpointOAuthReplace {
	var out []restapi.EndpointOAuthReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOAuthReplace(v))
	}
	return &out
}

func flattenEndpointWebhookValidationReplace(obj *restapi.EndpointWebhookValidationReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointWebhookValidation(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointWebhookValidationReplaceSlice(objs *[]restapi.EndpointWebhookValidationReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointWebhookValidationReplace(&v))
	}
	return sl
}

func expandEndpointWebhookValidationReplace(in interface{}) *restapi.EndpointWebhookValidationReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointWebhookValidationReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointWebhookValidation(v)
	}
	return &obj
}

func expandEndpointWebhookValidationReplaceSlice(in interface{}) *[]restapi.EndpointWebhookValidationReplace {
	var out []restapi.EndpointWebhookValidationReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointWebhookValidationReplace(v))
	}
	return &out
}

func flattenEndpointSAMLReplace(obj *restapi.EndpointSAMLReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointSAMLMutate(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointSAMLReplaceSlice(objs *[]restapi.EndpointSAMLReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointSAMLReplace(&v))
	}
	return sl
}

func expandEndpointSAMLReplace(in interface{}) *restapi.EndpointSAMLReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointSAMLReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointSAMLMutate(v)
	}
	return &obj
}

func expandEndpointSAMLReplaceSlice(in interface{}) *[]restapi.EndpointSAMLReplace {
	var out []restapi.EndpointSAMLReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointSAMLReplace(v))
	}
	return &out
}

func flattenEndpointOIDCReplace(obj *restapi.EndpointOIDCReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointOIDC(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointOIDCReplaceSlice(objs *[]restapi.EndpointOIDCReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointOIDCReplace(&v))
	}
	return sl
}

func expandEndpointOIDCReplace(in interface{}) *restapi.EndpointOIDCReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointOIDCReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointOIDC(v)
	}
	return &obj
}

func expandEndpointOIDCReplaceSlice(in interface{}) *[]restapi.EndpointOIDCReplace {
	var out []restapi.EndpointOIDCReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointOIDCReplace(v))
	}
	return &out
}

func flattenEndpointBackendReplace(obj *restapi.EndpointBackendReplace) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["module"] = flattenEndpointBackendMutate(&obj.Module)

	return []interface{}{m}
}

func flattenEndpointBackendReplaceSlice(objs *[]restapi.EndpointBackendReplace) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenEndpointBackendReplace(&v))
	}
	return sl
}

func expandEndpointBackendReplace(in interface{}) *restapi.EndpointBackendReplace {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.EndpointBackendReplace
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["module"]; ok {
		obj.Module = *expandEndpointBackendMutate(v)
	}
	return &obj
}

func expandEndpointBackendReplaceSlice(in interface{}) *[]restapi.EndpointBackendReplace {
	var out []restapi.EndpointBackendReplace
	for _, v := range in.([]interface{}) {
		out = append(out, *expandEndpointBackendReplace(v))
	}
	return &out
}

func flattenReservedAddrCreate(obj *restapi.ReservedAddrCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["region"] = obj.Region
	m["endpoint_configuration_id"] = obj.EndpointConfigurationID

	return []interface{}{m}
}

func flattenReservedAddrCreateSlice(objs *[]restapi.ReservedAddrCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedAddrCreate(&v))
	}
	return sl
}

func expandReservedAddrCreate(in interface{}) *restapi.ReservedAddrCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedAddrCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["region"]; ok {
		obj.Region = *expandString(v)
	}
	if v, ok := m["endpoint_configuration_id"]; ok {
		obj.EndpointConfigurationID = expandString(v)
	}
	return &obj
}

func expandReservedAddrCreateSlice(in interface{}) *[]restapi.ReservedAddrCreate {
	var out []restapi.ReservedAddrCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedAddrCreate(v))
	}
	return &out
}

func flattenReservedAddrUpdate(obj *restapi.ReservedAddrUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["endpoint_configuration_id"] = obj.EndpointConfigurationID

	return []interface{}{m}
}

func flattenReservedAddrUpdateSlice(objs *[]restapi.ReservedAddrUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedAddrUpdate(&v))
	}
	return sl
}

func expandReservedAddrUpdate(in interface{}) *restapi.ReservedAddrUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedAddrUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["endpoint_configuration_id"]; ok {
		obj.EndpointConfigurationID = expandString(v)
	}
	return &obj
}

func expandReservedAddrUpdateSlice(in interface{}) *[]restapi.ReservedAddrUpdate {
	var out []restapi.ReservedAddrUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedAddrUpdate(v))
	}
	return &out
}

func flattenReservedAddr(obj *restapi.ReservedAddr) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["addr"] = obj.Addr
	m["region"] = obj.Region
	m["endpoint_configuration"] = flattenRef(obj.EndpointConfiguration)

	return []interface{}{m}
}

func flattenReservedAddrSlice(objs *[]restapi.ReservedAddr) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedAddr(&v))
	}
	return sl
}

func expandReservedAddr(in interface{}) *restapi.ReservedAddr {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedAddr
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["addr"]; ok {
		obj.Addr = *expandString(v)
	}
	if v, ok := m["region"]; ok {
		obj.Region = *expandString(v)
	}
	if v, ok := m["endpoint_configuration"]; ok {
		obj.EndpointConfiguration = expandRef(v)
	}
	return &obj
}

func expandReservedAddrSlice(in interface{}) *[]restapi.ReservedAddr {
	var out []restapi.ReservedAddr
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedAddr(v))
	}
	return &out
}

func flattenReservedAddrList(obj *restapi.ReservedAddrList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["reserved_addrs"] = flattenReservedAddrSlice(&obj.ReservedAddrs)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenReservedAddrListSlice(objs *[]restapi.ReservedAddrList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedAddrList(&v))
	}
	return sl
}

func expandReservedAddrList(in interface{}) *restapi.ReservedAddrList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedAddrList
	if v, ok := m["reserved_addrs"]; ok {
		obj.ReservedAddrs = *expandReservedAddrSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandReservedAddrListSlice(in interface{}) *[]restapi.ReservedAddrList {
	var out []restapi.ReservedAddrList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedAddrList(v))
	}
	return &out
}

func flattenReservedDomainCreate(obj *restapi.ReservedDomainCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["name"] = obj.Name
	m["domain"] = obj.Domain
	m["region"] = obj.Region
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["http_endpoint_configuration_id"] = obj.HTTPEndpointConfigurationID
	m["https_endpoint_configuration_id"] = obj.HTTPSEndpointConfigurationID
	m["certificate_id"] = obj.CertificateID
	m["certificate_management_policy"] = flattenReservedDomainCertPolicy(obj.CertificateManagementPolicy)
	m["error_redirect_url"] = obj.ErrorRedirectUrl

	return []interface{}{m}
}

func flattenReservedDomainCreateSlice(objs *[]restapi.ReservedDomainCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedDomainCreate(&v))
	}
	return sl
}

func expandReservedDomainCreate(in interface{}) *restapi.ReservedDomainCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedDomainCreate
	if v, ok := m["name"]; ok {
		obj.Name = *expandString(v)
	}
	if v, ok := m["domain"]; ok {
		obj.Domain = *expandString(v)
	}
	if v, ok := m["region"]; ok {
		obj.Region = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["http_endpoint_configuration_id"]; ok {
		obj.HTTPEndpointConfigurationID = expandString(v)
	}
	if v, ok := m["https_endpoint_configuration_id"]; ok {
		obj.HTTPSEndpointConfigurationID = expandString(v)
	}
	if v, ok := m["certificate_id"]; ok {
		obj.CertificateID = expandString(v)
	}
	if v, ok := m["certificate_management_policy"]; ok {
		obj.CertificateManagementPolicy = expandReservedDomainCertPolicy(v)
	}
	if v, ok := m["error_redirect_url"]; ok {
		obj.ErrorRedirectUrl = expandString(v)
	}
	return &obj
}

func expandReservedDomainCreateSlice(in interface{}) *[]restapi.ReservedDomainCreate {
	var out []restapi.ReservedDomainCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedDomainCreate(v))
	}
	return &out
}

func flattenReservedDomainUpdate(obj *restapi.ReservedDomainUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["http_endpoint_configuration_id"] = obj.HTTPEndpointConfigurationID
	m["https_endpoint_configuration_id"] = obj.HTTPSEndpointConfigurationID
	m["certificate_id"] = obj.CertificateID
	m["certificate_management_policy"] = flattenReservedDomainCertPolicy(obj.CertificateManagementPolicy)
	m["region"] = obj.Region
	m["error_redirect_url"] = obj.ErrorRedirectUrl

	return []interface{}{m}
}

func flattenReservedDomainUpdateSlice(objs *[]restapi.ReservedDomainUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedDomainUpdate(&v))
	}
	return sl
}

func expandReservedDomainUpdate(in interface{}) *restapi.ReservedDomainUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedDomainUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["http_endpoint_configuration_id"]; ok {
		obj.HTTPEndpointConfigurationID = expandString(v)
	}
	if v, ok := m["https_endpoint_configuration_id"]; ok {
		obj.HTTPSEndpointConfigurationID = expandString(v)
	}
	if v, ok := m["certificate_id"]; ok {
		obj.CertificateID = expandString(v)
	}
	if v, ok := m["certificate_management_policy"]; ok {
		obj.CertificateManagementPolicy = expandReservedDomainCertPolicy(v)
	}
	if v, ok := m["region"]; ok {
		obj.Region = expandString(v)
	}
	if v, ok := m["error_redirect_url"]; ok {
		obj.ErrorRedirectUrl = expandString(v)
	}
	return &obj
}

func expandReservedDomainUpdateSlice(in interface{}) *[]restapi.ReservedDomainUpdate {
	var out []restapi.ReservedDomainUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedDomainUpdate(v))
	}
	return &out
}

func flattenReservedDomain(obj *restapi.ReservedDomain) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["domain"] = obj.Domain
	m["region"] = obj.Region
	m["cname_target"] = obj.CNAMETarget
	m["http_endpoint_configuration"] = flattenRef(obj.HTTPEndpointConfiguration)
	m["https_endpoint_configuration"] = flattenRef(obj.HTTPSEndpointConfiguration)
	m["certificate"] = flattenRef(obj.Certificate)
	m["certificate_management_policy"] = flattenReservedDomainCertPolicy(obj.CertificateManagementPolicy)
	m["certificate_management_status"] = flattenReservedDomainCertStatus(obj.CertificateManagementStatus)
	m["acme_challenge_cname_target"] = obj.ACMEChallengeCNAMETarget
	m["error_redirect_url"] = obj.ErrorRedirectURL

	return []interface{}{m}
}

func flattenReservedDomainSlice(objs *[]restapi.ReservedDomain) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedDomain(&v))
	}
	return sl
}

func expandReservedDomain(in interface{}) *restapi.ReservedDomain {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedDomain
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["domain"]; ok {
		obj.Domain = *expandString(v)
	}
	if v, ok := m["region"]; ok {
		obj.Region = *expandString(v)
	}
	if v, ok := m["cname_target"]; ok {
		obj.CNAMETarget = expandString(v)
	}
	if v, ok := m["http_endpoint_configuration"]; ok {
		obj.HTTPEndpointConfiguration = expandRef(v)
	}
	if v, ok := m["https_endpoint_configuration"]; ok {
		obj.HTTPSEndpointConfiguration = expandRef(v)
	}
	if v, ok := m["certificate"]; ok {
		obj.Certificate = expandRef(v)
	}
	if v, ok := m["certificate_management_policy"]; ok {
		obj.CertificateManagementPolicy = expandReservedDomainCertPolicy(v)
	}
	if v, ok := m["certificate_management_status"]; ok {
		obj.CertificateManagementStatus = expandReservedDomainCertStatus(v)
	}
	if v, ok := m["acme_challenge_cname_target"]; ok {
		obj.ACMEChallengeCNAMETarget = expandString(v)
	}
	if v, ok := m["error_redirect_url"]; ok {
		obj.ErrorRedirectURL = expandString(v)
	}
	return &obj
}

func expandReservedDomainSlice(in interface{}) *[]restapi.ReservedDomain {
	var out []restapi.ReservedDomain
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedDomain(v))
	}
	return &out
}

func flattenReservedDomainList(obj *restapi.ReservedDomainList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["reserved_domains"] = flattenReservedDomainSlice(&obj.ReservedDomains)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenReservedDomainListSlice(objs *[]restapi.ReservedDomainList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedDomainList(&v))
	}
	return sl
}

func expandReservedDomainList(in interface{}) *restapi.ReservedDomainList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedDomainList
	if v, ok := m["reserved_domains"]; ok {
		obj.ReservedDomains = *expandReservedDomainSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandReservedDomainListSlice(in interface{}) *[]restapi.ReservedDomainList {
	var out []restapi.ReservedDomainList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedDomainList(v))
	}
	return &out
}

func flattenReservedDomainCertPolicy(obj *restapi.ReservedDomainCertPolicy) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["authority"] = obj.Authority
	m["private_key_type"] = obj.PrivateKeyType

	return []interface{}{m}
}

func flattenReservedDomainCertPolicySlice(objs *[]restapi.ReservedDomainCertPolicy) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedDomainCertPolicy(&v))
	}
	return sl
}

func expandReservedDomainCertPolicy(in interface{}) *restapi.ReservedDomainCertPolicy {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedDomainCertPolicy
	if v, ok := m["authority"]; ok {
		obj.Authority = *expandString(v)
	}
	if v, ok := m["private_key_type"]; ok {
		obj.PrivateKeyType = *expandString(v)
	}
	return &obj
}

func expandReservedDomainCertPolicySlice(in interface{}) *[]restapi.ReservedDomainCertPolicy {
	var out []restapi.ReservedDomainCertPolicy
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedDomainCertPolicy(v))
	}
	return &out
}

func flattenReservedDomainCertStatus(obj *restapi.ReservedDomainCertStatus) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["renews_at"] = obj.RenewsAt
	m["provisioning_job"] = flattenReservedDomainCertJob(obj.ProvisioningJob)

	return []interface{}{m}
}

func flattenReservedDomainCertStatusSlice(objs *[]restapi.ReservedDomainCertStatus) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedDomainCertStatus(&v))
	}
	return sl
}

func expandReservedDomainCertStatus(in interface{}) *restapi.ReservedDomainCertStatus {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedDomainCertStatus
	if v, ok := m["renews_at"]; ok {
		obj.RenewsAt = expandString(v)
	}
	if v, ok := m["provisioning_job"]; ok {
		obj.ProvisioningJob = expandReservedDomainCertJob(v)
	}
	return &obj
}

func expandReservedDomainCertStatusSlice(in interface{}) *[]restapi.ReservedDomainCertStatus {
	var out []restapi.ReservedDomainCertStatus
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedDomainCertStatus(v))
	}
	return &out
}

func flattenReservedDomainCertJob(obj *restapi.ReservedDomainCertJob) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["error_code"] = obj.ErrorCode
	m["msg"] = obj.Msg
	m["started_at"] = obj.StartedAt
	m["retries_at"] = obj.RetriesAt

	return []interface{}{m}
}

func flattenReservedDomainCertJobSlice(objs *[]restapi.ReservedDomainCertJob) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenReservedDomainCertJob(&v))
	}
	return sl
}

func expandReservedDomainCertJob(in interface{}) *restapi.ReservedDomainCertJob {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.ReservedDomainCertJob
	if v, ok := m["error_code"]; ok {
		obj.ErrorCode = expandString(v)
	}
	if v, ok := m["msg"]; ok {
		obj.Msg = *expandString(v)
	}
	if v, ok := m["started_at"]; ok {
		obj.StartedAt = *expandString(v)
	}
	if v, ok := m["retries_at"]; ok {
		obj.RetriesAt = expandString(v)
	}
	return &obj
}

func expandReservedDomainCertJobSlice(in interface{}) *[]restapi.ReservedDomainCertJob {
	var out []restapi.ReservedDomainCertJob
	for _, v := range in.([]interface{}) {
		out = append(out, *expandReservedDomainCertJob(v))
	}
	return &out
}

func flattenRootResponse(obj *restapi.RootResponse) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["uri"] = obj.URI
	m["subresource_uris"] = obj.SubresourceURIs

	return []interface{}{m}
}

func flattenRootResponseSlice(objs *[]restapi.RootResponse) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenRootResponse(&v))
	}
	return sl
}

func expandRootResponse(in interface{}) *restapi.RootResponse {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.RootResponse
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["subresource_uris"]; ok {
		obj.SubresourceURIs = *expandStringMap(v)
	}
	return &obj
}

func expandRootResponseSlice(in interface{}) *[]restapi.RootResponse {
	var out []restapi.RootResponse
	for _, v := range in.([]interface{}) {
		out = append(out, *expandRootResponse(v))
	}
	return &out
}

func flattenSSHCertificateAuthorityCreate(obj *restapi.SSHCertificateAuthorityCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["private_key_type"] = obj.PrivateKeyType
	m["elliptic_curve"] = obj.EllipticCurve
	m["key_size"] = obj.KeySize

	return []interface{}{m}
}

func flattenSSHCertificateAuthorityCreateSlice(objs *[]restapi.SSHCertificateAuthorityCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHCertificateAuthorityCreate(&v))
	}
	return sl
}

func expandSSHCertificateAuthorityCreate(in interface{}) *restapi.SSHCertificateAuthorityCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHCertificateAuthorityCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["private_key_type"]; ok {
		obj.PrivateKeyType = *expandString(v)
	}
	if v, ok := m["elliptic_curve"]; ok {
		obj.EllipticCurve = *expandString(v)
	}
	if v, ok := m["key_size"]; ok {
		obj.KeySize = *expandInt64(v)
	}
	return &obj
}

func expandSSHCertificateAuthorityCreateSlice(in interface{}) *[]restapi.SSHCertificateAuthorityCreate {
	var out []restapi.SSHCertificateAuthorityCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHCertificateAuthorityCreate(v))
	}
	return &out
}

func flattenSSHCertificateAuthorityUpdate(obj *restapi.SSHCertificateAuthorityUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata

	return []interface{}{m}
}

func flattenSSHCertificateAuthorityUpdateSlice(objs *[]restapi.SSHCertificateAuthorityUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHCertificateAuthorityUpdate(&v))
	}
	return sl
}

func expandSSHCertificateAuthorityUpdate(in interface{}) *restapi.SSHCertificateAuthorityUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHCertificateAuthorityUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	return &obj
}

func expandSSHCertificateAuthorityUpdateSlice(in interface{}) *[]restapi.SSHCertificateAuthorityUpdate {
	var out []restapi.SSHCertificateAuthorityUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHCertificateAuthorityUpdate(v))
	}
	return &out
}

func flattenSSHCertificateAuthority(obj *restapi.SSHCertificateAuthority) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["public_key"] = obj.PublicKey
	m["key_type"] = obj.KeyType

	return []interface{}{m}
}

func flattenSSHCertificateAuthoritySlice(objs *[]restapi.SSHCertificateAuthority) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHCertificateAuthority(&v))
	}
	return sl
}

func expandSSHCertificateAuthority(in interface{}) *restapi.SSHCertificateAuthority {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHCertificateAuthority
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["public_key"]; ok {
		obj.PublicKey = *expandString(v)
	}
	if v, ok := m["key_type"]; ok {
		obj.KeyType = *expandString(v)
	}
	return &obj
}

func expandSSHCertificateAuthoritySlice(in interface{}) *[]restapi.SSHCertificateAuthority {
	var out []restapi.SSHCertificateAuthority
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHCertificateAuthority(v))
	}
	return &out
}

func flattenSSHCertificateAuthorityList(obj *restapi.SSHCertificateAuthorityList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["ssh_certificate_authorities"] = flattenSSHCertificateAuthoritySlice(&obj.SSHCertificateAuthorities)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenSSHCertificateAuthorityListSlice(objs *[]restapi.SSHCertificateAuthorityList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHCertificateAuthorityList(&v))
	}
	return sl
}

func expandSSHCertificateAuthorityList(in interface{}) *restapi.SSHCertificateAuthorityList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHCertificateAuthorityList
	if v, ok := m["ssh_certificate_authorities"]; ok {
		obj.SSHCertificateAuthorities = *expandSSHCertificateAuthoritySlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandSSHCertificateAuthorityListSlice(in interface{}) *[]restapi.SSHCertificateAuthorityList {
	var out []restapi.SSHCertificateAuthorityList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHCertificateAuthorityList(v))
	}
	return &out
}

func flattenSSHCredentialCreate(obj *restapi.SSHCredentialCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["acl"] = obj.ACL
	m["public_key"] = obj.PublicKey
	m["owner_id"] = obj.OwnerID
	m["owner_email"] = obj.OwnerEmail

	return []interface{}{m}
}

func flattenSSHCredentialCreateSlice(objs *[]restapi.SSHCredentialCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHCredentialCreate(&v))
	}
	return sl
}

func expandSSHCredentialCreate(in interface{}) *restapi.SSHCredentialCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHCredentialCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["acl"]; ok {
		obj.ACL = *expandStringSlice(v)
	}
	if v, ok := m["public_key"]; ok {
		obj.PublicKey = *expandString(v)
	}
	if v, ok := m["owner_id"]; ok {
		obj.OwnerID = expandString(v)
	}
	if v, ok := m["owner_email"]; ok {
		obj.OwnerEmail = *expandString(v)
	}
	return &obj
}

func expandSSHCredentialCreateSlice(in interface{}) *[]restapi.SSHCredentialCreate {
	var out []restapi.SSHCredentialCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHCredentialCreate(v))
	}
	return &out
}

func flattenSSHCredentialUpdate(obj *restapi.SSHCredentialUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["acl"] = obj.ACL

	return []interface{}{m}
}

func flattenSSHCredentialUpdateSlice(objs *[]restapi.SSHCredentialUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHCredentialUpdate(&v))
	}
	return sl
}

func expandSSHCredentialUpdate(in interface{}) *restapi.SSHCredentialUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHCredentialUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	if v, ok := m["acl"]; ok {
		obj.ACL = expandStringSlice(v)
	}
	return &obj
}

func expandSSHCredentialUpdateSlice(in interface{}) *[]restapi.SSHCredentialUpdate {
	var out []restapi.SSHCredentialUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHCredentialUpdate(v))
	}
	return &out
}

func flattenSSHCredential(obj *restapi.SSHCredential) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["public_key"] = obj.PublicKey
	m["acl"] = obj.ACL
	m["owner_id"] = obj.OwnerID

	return []interface{}{m}
}

func flattenSSHCredentialSlice(objs *[]restapi.SSHCredential) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHCredential(&v))
	}
	return sl
}

func expandSSHCredential(in interface{}) *restapi.SSHCredential {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHCredential
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["public_key"]; ok {
		obj.PublicKey = *expandString(v)
	}
	if v, ok := m["acl"]; ok {
		obj.ACL = *expandStringSlice(v)
	}
	if v, ok := m["owner_id"]; ok {
		obj.OwnerID = expandString(v)
	}
	return &obj
}

func expandSSHCredentialSlice(in interface{}) *[]restapi.SSHCredential {
	var out []restapi.SSHCredential
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHCredential(v))
	}
	return &out
}

func flattenSSHCredentialList(obj *restapi.SSHCredentialList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["ssh_credentials"] = flattenSSHCredentialSlice(&obj.SSHCredentials)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenSSHCredentialListSlice(objs *[]restapi.SSHCredentialList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHCredentialList(&v))
	}
	return sl
}

func expandSSHCredentialList(in interface{}) *restapi.SSHCredentialList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHCredentialList
	if v, ok := m["ssh_credentials"]; ok {
		obj.SSHCredentials = *expandSSHCredentialSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandSSHCredentialListSlice(in interface{}) *[]restapi.SSHCredentialList {
	var out []restapi.SSHCredentialList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHCredentialList(v))
	}
	return &out
}

func flattenSSHHostCertificateCreate(obj *restapi.SSHHostCertificateCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["ssh_certificate_authority_id"] = obj.SSHCertificateAuthorityID
	m["public_key"] = obj.PublicKey
	m["principals"] = obj.Principals
	m["valid_after"] = obj.ValidAfter
	m["valid_until"] = obj.ValidUntil
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata

	return []interface{}{m}
}

func flattenSSHHostCertificateCreateSlice(objs *[]restapi.SSHHostCertificateCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHHostCertificateCreate(&v))
	}
	return sl
}

func expandSSHHostCertificateCreate(in interface{}) *restapi.SSHHostCertificateCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHHostCertificateCreate
	if v, ok := m["ssh_certificate_authority_id"]; ok {
		obj.SSHCertificateAuthorityID = *expandString(v)
	}
	if v, ok := m["public_key"]; ok {
		obj.PublicKey = *expandString(v)
	}
	if v, ok := m["principals"]; ok {
		obj.Principals = *expandStringSlice(v)
	}
	if v, ok := m["valid_after"]; ok {
		obj.ValidAfter = *expandString(v)
	}
	if v, ok := m["valid_until"]; ok {
		obj.ValidUntil = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	return &obj
}

func expandSSHHostCertificateCreateSlice(in interface{}) *[]restapi.SSHHostCertificateCreate {
	var out []restapi.SSHHostCertificateCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHHostCertificateCreate(v))
	}
	return &out
}

func flattenSSHHostCertificateUpdate(obj *restapi.SSHHostCertificateUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata

	return []interface{}{m}
}

func flattenSSHHostCertificateUpdateSlice(objs *[]restapi.SSHHostCertificateUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHHostCertificateUpdate(&v))
	}
	return sl
}

func expandSSHHostCertificateUpdate(in interface{}) *restapi.SSHHostCertificateUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHHostCertificateUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	return &obj
}

func expandSSHHostCertificateUpdateSlice(in interface{}) *[]restapi.SSHHostCertificateUpdate {
	var out []restapi.SSHHostCertificateUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHHostCertificateUpdate(v))
	}
	return &out
}

func flattenSSHHostCertificate(obj *restapi.SSHHostCertificate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["public_key"] = obj.PublicKey
	m["key_type"] = obj.KeyType
	m["ssh_certificate_authority_id"] = obj.SSHCertificateAuthorityID
	m["principals"] = obj.Principals
	m["valid_after"] = obj.ValidAfter
	m["valid_until"] = obj.ValidUntil
	m["certificate"] = obj.Certificate

	return []interface{}{m}
}

func flattenSSHHostCertificateSlice(objs *[]restapi.SSHHostCertificate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHHostCertificate(&v))
	}
	return sl
}

func expandSSHHostCertificate(in interface{}) *restapi.SSHHostCertificate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHHostCertificate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["public_key"]; ok {
		obj.PublicKey = *expandString(v)
	}
	if v, ok := m["key_type"]; ok {
		obj.KeyType = *expandString(v)
	}
	if v, ok := m["ssh_certificate_authority_id"]; ok {
		obj.SSHCertificateAuthorityID = *expandString(v)
	}
	if v, ok := m["principals"]; ok {
		obj.Principals = *expandStringSlice(v)
	}
	if v, ok := m["valid_after"]; ok {
		obj.ValidAfter = *expandString(v)
	}
	if v, ok := m["valid_until"]; ok {
		obj.ValidUntil = *expandString(v)
	}
	if v, ok := m["certificate"]; ok {
		obj.Certificate = *expandString(v)
	}
	return &obj
}

func expandSSHHostCertificateSlice(in interface{}) *[]restapi.SSHHostCertificate {
	var out []restapi.SSHHostCertificate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHHostCertificate(v))
	}
	return &out
}

func flattenSSHHostCertificateList(obj *restapi.SSHHostCertificateList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["ssh_host_certificates"] = flattenSSHHostCertificateSlice(&obj.SSHHostCertificates)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenSSHHostCertificateListSlice(objs *[]restapi.SSHHostCertificateList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHHostCertificateList(&v))
	}
	return sl
}

func expandSSHHostCertificateList(in interface{}) *restapi.SSHHostCertificateList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHHostCertificateList
	if v, ok := m["ssh_host_certificates"]; ok {
		obj.SSHHostCertificates = *expandSSHHostCertificateSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandSSHHostCertificateListSlice(in interface{}) *[]restapi.SSHHostCertificateList {
	var out []restapi.SSHHostCertificateList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHHostCertificateList(v))
	}
	return &out
}

func flattenSSHUserCertificateCreate(obj *restapi.SSHUserCertificateCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["ssh_certificate_authority_id"] = obj.SSHCertificateAuthorityID
	m["public_key"] = obj.PublicKey
	m["principals"] = obj.Principals
	m["critical_options"] = obj.CriticalOptions
	m["extensions"] = obj.Extensions
	m["valid_after"] = obj.ValidAfter
	m["valid_until"] = obj.ValidUntil
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata

	return []interface{}{m}
}

func flattenSSHUserCertificateCreateSlice(objs *[]restapi.SSHUserCertificateCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHUserCertificateCreate(&v))
	}
	return sl
}

func expandSSHUserCertificateCreate(in interface{}) *restapi.SSHUserCertificateCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHUserCertificateCreate
	if v, ok := m["ssh_certificate_authority_id"]; ok {
		obj.SSHCertificateAuthorityID = *expandString(v)
	}
	if v, ok := m["public_key"]; ok {
		obj.PublicKey = *expandString(v)
	}
	if v, ok := m["principals"]; ok {
		obj.Principals = *expandStringSlice(v)
	}
	if v, ok := m["critical_options"]; ok {
		obj.CriticalOptions = *expandStringMap(v)
	}
	if v, ok := m["extensions"]; ok {
		obj.Extensions = *expandStringMap(v)
	}
	if v, ok := m["valid_after"]; ok {
		obj.ValidAfter = *expandString(v)
	}
	if v, ok := m["valid_until"]; ok {
		obj.ValidUntil = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	return &obj
}

func expandSSHUserCertificateCreateSlice(in interface{}) *[]restapi.SSHUserCertificateCreate {
	var out []restapi.SSHUserCertificateCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHUserCertificateCreate(v))
	}
	return &out
}

func flattenSSHUserCertificateUpdate(obj *restapi.SSHUserCertificateUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata

	return []interface{}{m}
}

func flattenSSHUserCertificateUpdateSlice(objs *[]restapi.SSHUserCertificateUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHUserCertificateUpdate(&v))
	}
	return sl
}

func expandSSHUserCertificateUpdate(in interface{}) *restapi.SSHUserCertificateUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHUserCertificateUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	return &obj
}

func expandSSHUserCertificateUpdateSlice(in interface{}) *[]restapi.SSHUserCertificateUpdate {
	var out []restapi.SSHUserCertificateUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHUserCertificateUpdate(v))
	}
	return &out
}

func flattenSSHUserCertificate(obj *restapi.SSHUserCertificate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["public_key"] = obj.PublicKey
	m["key_type"] = obj.KeyType
	m["ssh_certificate_authority_id"] = obj.SSHCertificateAuthorityID
	m["principals"] = obj.Principals
	m["critical_options"] = obj.CriticalOptions
	m["extensions"] = obj.Extensions
	m["valid_after"] = obj.ValidAfter
	m["valid_until"] = obj.ValidUntil
	m["certificate"] = obj.Certificate

	return []interface{}{m}
}

func flattenSSHUserCertificateSlice(objs *[]restapi.SSHUserCertificate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHUserCertificate(&v))
	}
	return sl
}

func expandSSHUserCertificate(in interface{}) *restapi.SSHUserCertificate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHUserCertificate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["public_key"]; ok {
		obj.PublicKey = *expandString(v)
	}
	if v, ok := m["key_type"]; ok {
		obj.KeyType = *expandString(v)
	}
	if v, ok := m["ssh_certificate_authority_id"]; ok {
		obj.SSHCertificateAuthorityID = *expandString(v)
	}
	if v, ok := m["principals"]; ok {
		obj.Principals = *expandStringSlice(v)
	}
	if v, ok := m["critical_options"]; ok {
		obj.CriticalOptions = *expandStringMap(v)
	}
	if v, ok := m["extensions"]; ok {
		obj.Extensions = *expandStringMap(v)
	}
	if v, ok := m["valid_after"]; ok {
		obj.ValidAfter = *expandString(v)
	}
	if v, ok := m["valid_until"]; ok {
		obj.ValidUntil = *expandString(v)
	}
	if v, ok := m["certificate"]; ok {
		obj.Certificate = *expandString(v)
	}
	return &obj
}

func expandSSHUserCertificateSlice(in interface{}) *[]restapi.SSHUserCertificate {
	var out []restapi.SSHUserCertificate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHUserCertificate(v))
	}
	return &out
}

func flattenSSHUserCertificateList(obj *restapi.SSHUserCertificateList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["ssh_user_certificates"] = flattenSSHUserCertificateSlice(&obj.SSHUserCertificates)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenSSHUserCertificateListSlice(objs *[]restapi.SSHUserCertificateList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenSSHUserCertificateList(&v))
	}
	return sl
}

func expandSSHUserCertificateList(in interface{}) *restapi.SSHUserCertificateList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.SSHUserCertificateList
	if v, ok := m["ssh_user_certificates"]; ok {
		obj.SSHUserCertificates = *expandSSHUserCertificateSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandSSHUserCertificateListSlice(in interface{}) *[]restapi.SSHUserCertificateList {
	var out []restapi.SSHUserCertificateList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandSSHUserCertificateList(v))
	}
	return &out
}

func flattenTLSCertificateCreate(obj *restapi.TLSCertificateCreate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["certificate_pem"] = obj.CertificatePEM
	m["private_key_pem"] = obj.PrivateKeyPEM

	return []interface{}{m}
}

func flattenTLSCertificateCreateSlice(objs *[]restapi.TLSCertificateCreate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTLSCertificateCreate(&v))
	}
	return sl
}

func expandTLSCertificateCreate(in interface{}) *restapi.TLSCertificateCreate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TLSCertificateCreate
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["certificate_pem"]; ok {
		obj.CertificatePEM = *expandString(v)
	}
	if v, ok := m["private_key_pem"]; ok {
		obj.PrivateKeyPEM = *expandString(v)
	}
	return &obj
}

func expandTLSCertificateCreateSlice(in interface{}) *[]restapi.TLSCertificateCreate {
	var out []restapi.TLSCertificateCreate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTLSCertificateCreate(v))
	}
	return &out
}

func flattenTLSCertificateUpdate(obj *restapi.TLSCertificateUpdate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata

	return []interface{}{m}
}

func flattenTLSCertificateUpdateSlice(objs *[]restapi.TLSCertificateUpdate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTLSCertificateUpdate(&v))
	}
	return sl
}

func expandTLSCertificateUpdate(in interface{}) *restapi.TLSCertificateUpdate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TLSCertificateUpdate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = expandString(v)
	}
	return &obj
}

func expandTLSCertificateUpdateSlice(in interface{}) *[]restapi.TLSCertificateUpdate {
	var out []restapi.TLSCertificateUpdate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTLSCertificateUpdate(v))
	}
	return &out
}

func flattenTLSCertificate(obj *restapi.TLSCertificate) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["uri"] = obj.URI
	m["created_at"] = obj.CreatedAt
	m["description"] = obj.Description
	m["metadata"] = obj.Metadata
	m["certificate_pem"] = obj.CertificatePEM
	m["subject_common_name"] = obj.SubjectCommonName
	m["subject_alternative_names"] = flattenTLSCertificateSANs(&obj.SubjectAlternativeNames)
	m["issued_at"] = obj.IssuedAt
	m["not_before"] = obj.NotBefore
	m["not_after"] = obj.NotAfter
	m["key_usages"] = obj.KeyUsages
	m["extended_key_usages"] = obj.ExtendedKeyUsages
	m["private_key_type"] = obj.PrivateKeyType
	m["issuer_common_name"] = obj.IssuerCommonName
	m["serial_number"] = obj.SerialNumber
	m["subject_organization"] = obj.SubjectOrganization
	m["subject_organizational_unit"] = obj.SubjectOrganizationalUnit
	m["subject_locality"] = obj.SubjectLocality
	m["subject_province"] = obj.SubjectProvince
	m["subject_country"] = obj.SubjectCountry

	return []interface{}{m}
}

func flattenTLSCertificateSlice(objs *[]restapi.TLSCertificate) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTLSCertificate(&v))
	}
	return sl
}

func expandTLSCertificate(in interface{}) *restapi.TLSCertificate {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TLSCertificate
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["created_at"]; ok {
		obj.CreatedAt = *expandString(v)
	}
	if v, ok := m["description"]; ok {
		obj.Description = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["certificate_pem"]; ok {
		obj.CertificatePEM = *expandString(v)
	}
	if v, ok := m["subject_common_name"]; ok {
		obj.SubjectCommonName = *expandString(v)
	}
	if v, ok := m["subject_alternative_names"]; ok {
		obj.SubjectAlternativeNames = *expandTLSCertificateSANs(v)
	}
	if v, ok := m["issued_at"]; ok {
		obj.IssuedAt = expandString(v)
	}
	if v, ok := m["not_before"]; ok {
		obj.NotBefore = *expandString(v)
	}
	if v, ok := m["not_after"]; ok {
		obj.NotAfter = *expandString(v)
	}
	if v, ok := m["key_usages"]; ok {
		obj.KeyUsages = *expandStringSlice(v)
	}
	if v, ok := m["extended_key_usages"]; ok {
		obj.ExtendedKeyUsages = *expandStringSlice(v)
	}
	if v, ok := m["private_key_type"]; ok {
		obj.PrivateKeyType = *expandString(v)
	}
	if v, ok := m["issuer_common_name"]; ok {
		obj.IssuerCommonName = *expandString(v)
	}
	if v, ok := m["serial_number"]; ok {
		obj.SerialNumber = *expandString(v)
	}
	if v, ok := m["subject_organization"]; ok {
		obj.SubjectOrganization = *expandString(v)
	}
	if v, ok := m["subject_organizational_unit"]; ok {
		obj.SubjectOrganizationalUnit = *expandString(v)
	}
	if v, ok := m["subject_locality"]; ok {
		obj.SubjectLocality = *expandString(v)
	}
	if v, ok := m["subject_province"]; ok {
		obj.SubjectProvince = *expandString(v)
	}
	if v, ok := m["subject_country"]; ok {
		obj.SubjectCountry = *expandString(v)
	}
	return &obj
}

func expandTLSCertificateSlice(in interface{}) *[]restapi.TLSCertificate {
	var out []restapi.TLSCertificate
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTLSCertificate(v))
	}
	return &out
}

func flattenTLSCertificateList(obj *restapi.TLSCertificateList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["tls_certificates"] = flattenTLSCertificateSlice(&obj.TLSCertificates)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenTLSCertificateListSlice(objs *[]restapi.TLSCertificateList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTLSCertificateList(&v))
	}
	return sl
}

func expandTLSCertificateList(in interface{}) *restapi.TLSCertificateList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TLSCertificateList
	if v, ok := m["tls_certificates"]; ok {
		obj.TLSCertificates = *expandTLSCertificateSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandTLSCertificateListSlice(in interface{}) *[]restapi.TLSCertificateList {
	var out []restapi.TLSCertificateList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTLSCertificateList(v))
	}
	return &out
}

func flattenTLSCertificateSANs(obj *restapi.TLSCertificateSANs) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["dns_names"] = obj.DNSNames
	m["ips"] = obj.IPs

	return []interface{}{m}
}

func flattenTLSCertificateSANsSlice(objs *[]restapi.TLSCertificateSANs) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTLSCertificateSANs(&v))
	}
	return sl
}

func expandTLSCertificateSANs(in interface{}) *restapi.TLSCertificateSANs {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TLSCertificateSANs
	if v, ok := m["dns_names"]; ok {
		obj.DNSNames = *expandStringSlice(v)
	}
	if v, ok := m["ips"]; ok {
		obj.IPs = *expandStringSlice(v)
	}
	return &obj
}

func expandTLSCertificateSANsSlice(in interface{}) *[]restapi.TLSCertificateSANs {
	var out []restapi.TLSCertificateSANs
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTLSCertificateSANs(v))
	}
	return &out
}

func flattenTunnel(obj *restapi.Tunnel) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["id"] = obj.ID
	m["public_url"] = obj.PublicURL
	m["started_at"] = obj.StartedAt
	m["metadata"] = obj.Metadata
	m["proto"] = obj.Proto
	m["region"] = obj.Region
	m["tunnel_session"] = flattenRef(&obj.TunnelSession)
	m["endpoint"] = flattenRef(obj.Endpoint)
	m["labels"] = obj.Labels
	m["backends"] = flattenRefSlice(obj.Backends)
	m["forwards_to"] = obj.ForwardsTo

	return []interface{}{m}
}

func flattenTunnelSlice(objs *[]restapi.Tunnel) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnel(&v))
	}
	return sl
}

func expandTunnel(in interface{}) *restapi.Tunnel {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.Tunnel
	if v, ok := m["id"]; ok {
		obj.ID = *expandString(v)
	}
	if v, ok := m["public_url"]; ok {
		obj.PublicURL = *expandString(v)
	}
	if v, ok := m["started_at"]; ok {
		obj.StartedAt = *expandString(v)
	}
	if v, ok := m["metadata"]; ok {
		obj.Metadata = *expandString(v)
	}
	if v, ok := m["proto"]; ok {
		obj.Proto = *expandString(v)
	}
	if v, ok := m["region"]; ok {
		obj.Region = *expandString(v)
	}
	if v, ok := m["tunnel_session"]; ok {
		obj.TunnelSession = *expandRef(v)
	}
	if v, ok := m["endpoint"]; ok {
		obj.Endpoint = expandRef(v)
	}
	if v, ok := m["labels"]; ok {
		obj.Labels = expandStringMap(v)
	}
	if v, ok := m["backends"]; ok {
		obj.Backends = expandRefSlice(v)
	}
	if v, ok := m["forwards_to"]; ok {
		obj.ForwardsTo = *expandString(v)
	}
	return &obj
}

func expandTunnelSlice(in interface{}) *[]restapi.Tunnel {
	var out []restapi.Tunnel
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnel(v))
	}
	return &out
}

func flattenTunnelList(obj *restapi.TunnelList) interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m["tunnels"] = flattenTunnelSlice(&obj.Tunnels)
	m["uri"] = obj.URI
	m["next_page_uri"] = obj.NextPageURI

	return []interface{}{m}
}

func flattenTunnelListSlice(objs *[]restapi.TunnelList) (sl []interface{}) {
	if objs == nil {
		return nil
	}

	for _, v := range *objs {
		sl = append(sl, flattenTunnelList(&v))
	}
	return sl
}

func expandTunnelList(in interface{}) *restapi.TunnelList {
	if in == nil {
		return nil
	}
	v := in.(*schema.Set)

	if v.Len() == 0 {
		return nil
	}

	m := v.List()[0].(map[string]interface{})
	var obj restapi.TunnelList
	if v, ok := m["tunnels"]; ok {
		obj.Tunnels = *expandTunnelSlice(v)
	}
	if v, ok := m["uri"]; ok {
		obj.URI = *expandString(v)
	}
	if v, ok := m["next_page_uri"]; ok {
		obj.NextPageURI = expandString(v)
	}
	return &obj
}

func expandTunnelListSlice(in interface{}) *[]restapi.TunnelList {
	var out []restapi.TunnelList
	for _, v := range in.([]interface{}) {
		out = append(out, *expandTunnelList(v))
	}
	return &out
}

func expandString(v interface{}) *string {
	x := v.(string)
	return &x
}

func expandInt32(v interface{}) *int32 {
	x := int32(v.(int))
	return &x
}

func expandInt64(v interface{}) *int64 {
	x := int64(v.(int))
	return &x
}

func expandInt64Map(vs interface{}) *map[string]int64 {
	out := make(map[string]int64)
	for k, v := range vs.(map[string]interface{}) {
		out[k] = v.(int64)
	}
	return &out
}

func expandUint32(v interface{}) *uint32 {
	x := uint32(v.(int))
	return &x
}

func expandUint64(v interface{}) *uint64 {
	x := uint64(v.(int))
	return &x
}

func expandBool(v interface{}) *bool {
	x := v.(bool)
	return &x
}

func expandFloat64(v interface{}) *float64 {
	x := v.(float64)
	return &x
}

func expandByteSlice(v interface{}) *[]byte {
	x := []byte(v.(string))
	return &x
}

func expandStringSlice(vs interface{}) *[]string {
	var out []string
	for _, v := range vs.([]interface{}) {
		out = append(out, v.(string))
	}
	return &out
}

func expandStringMap(vs interface{}) *map[string]string {
	out := make(map[string]string)
	for k, v := range vs.(map[string]interface{}) {
		out[k] = v.(string)
	}
	return &out
}

func expandAny(v interface{}) *any {
	return &v
}
