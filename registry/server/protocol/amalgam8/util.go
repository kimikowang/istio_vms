package amalgam8

import (
	"time"
	"fmt"
	"strings"
	"net/http"

	"github.com/amalgam8/amalgam8/registry/store"
	"github.com/amalgam8/amalgam8/registry/utils/reflection"
	"github.com/ant0ine/go-json-rest/rest"
)

func statusCodeFromError(err error) int {
	if regerr, ok := err.(*store.Error); ok {
		switch regerr.Code {
		case store.ErrorBadRequest:
			return http.StatusBadRequest
		case store.ErrorNoSuchServiceName:
			return http.StatusNotFound
		case store.ErrorNoSuchServiceInstance:
			return http.StatusGone
		case store.ErrorNamespaceQuotaExceeded:
			return http.StatusForbidden
		case store.ErrorInternalServerError:
			return http.StatusInternalServerError
		case store.ErrorNoInstanceServiceName:
			return http.StatusBadRequest
		case store.ErrorInstanceServiceNameTooLong:
			return http.StatusBadRequest
		case store.ErrorInstanceEndpointValueTooLong:
			return http.StatusBadRequest
		case store.ErrorInstanceStatusLengthTooLong:
			return http.StatusBadRequest
		case store.ErrorInstanceMetaDataTooLong:
			return http.StatusBadRequest
		default:
			return http.StatusInternalServerError
		}
	}
	return http.StatusInternalServerError
}

// Extract and validate filtering fields request. Note that an empty-string request is perfectly valid
func extractFields(r *rest.Request) ([]string, error) {
	if _, filteringRequested := r.URL.Query()["fields"]; !filteringRequested {
		return nil, nil
	}

	fieldsValue := r.URL.Query().Get("fields")
	if fieldsValue == "" {
		return []string{}, nil
	}

	fieldsSplit := strings.Split(fieldsValue, ",")
	fields := make([]string, len(fieldsSplit))
	for i, fld := range fieldsSplit {
		fldName, ok := instanceQueryValuesToFieldNames[fld]
		if !ok {
			return nil, fmt.Errorf("Field %s is not a valid field", fld)
		}

		fields[i] = fldName
	}

	return fields, nil
}

func copyServiceObject(s *store.Service) (*Service, error) {
	svc := &Service{
		ServiceName: s.ServiceName,
		Address: s.Address,
		Ports: CopyPortsList(s.Ports),
		ExternalName: s.ExternalName,
	}

	return svc, nil
}

func copyInstanceWithFilter(sname string, si *store.ServiceInstance, fields []string) (*ServiceInstance, error) {
	inst := &ServiceInstance{
		ID:          si.ID,
		ServiceName: sname,
		Endpoint: &InstanceAddress{
			Type:  si.Endpoint.Type,
			Value: si.Endpoint.Value,
		},
		Status:        si.Status,
		Tags:          si.Tags,
		TTL:           uint32(si.TTL / time.Second),
		Metadata:      si.Metadata,
		LastHeartbeat: &si.LastRenewal,
	}

	if fields != nil {
		filteredInstance := &ServiceInstance{}
		err := reflection.FilterStructByFields(inst, filteredInstance, fields)
		if err != nil {
			return nil, err
		}
		// Add Endpoint because it should always be returned to the user
		filteredInstance.Endpoint = &InstanceAddress{Type: inst.Endpoint.Type, Value: inst.Endpoint.Value}
		filteredInstance.ID = inst.ID
		return filteredInstance, nil
	}

	return inst, nil
}

func CopyPortsList(pl store.PortList) PortList {
	out := make(PortList, len(pl), len(pl))
	for idx, p := range pl {
		out[idx] = CopyPort(p)
	}

	return out
}

func CopyPort(p *store.Port) *Port {
	return &Port{
		Name: p.Name,
		Port: p.Port,
		Protocol: p.Protocol,
	}
}
