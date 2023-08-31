// Code generated for API Clients. DO NOT EDIT.

package transform

import (
	"github.com/ngrok/terraform-provider-ngrok/restapi"
)

func ConvertRefSliceToStringSlice(refs *[]restapi.Ref) (ss []string) {
	for _, ref := range *refs {
		ss = append(ss, ref.ID)
	}
	return ss
}
