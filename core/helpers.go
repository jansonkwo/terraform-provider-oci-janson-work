package core

import (
	"fmt"

	"github.com/MustWin/baremetal-sdk-go"
)

type resourceProvider interface {
	GetOk(string) (interface{}, bool)
}

func getOptionsWithNextPageID(nextPage string, opts []baremetal.Options) (optsWithNextPage []baremetal.Options, hasNextPage bool) {
	optsWithNextPage = opts

	if nextPage == "" {
		hasNextPage = false
	} else {
		hasNextPage = true

		if len(optsWithNextPage) == 0 {
			optsWithNextPage = append(optsWithNextPage, baremetal.Options{
				Page: nextPage,
			})
		} else {
			optsWithNextPage[0].Page = nextPage
		}
	}

	return optsWithNextPage, hasNextPage
}

func getCoreOptionsFromResourceData(resource resourceProvider, keys ...string) (opts []baremetal.Options) {
	opts = []baremetal.Options{}

	for _, key := range keys {
		if val, ok := resource.GetOk(key); ok {

			if len(opts) == 0 {
				opts = append(opts, baremetal.Options{})
			}

			switch key {
			case "availability_domain":
				opts[0].AvailabilityDomain = val.(string)
			case "image_id":
				opts[0].ImageID = val.(string)
			case "instance_id":
				opts[0].InstanceID = val.(string)
			case "limit":
				opts[0].Limit = uint64(val.(int))
			case "page":
				opts[0].Page = val.(string)
			case "vnic_id":
				opts[0].VnicID = val.(string)
			case "volume_id":
				opts[0].VolumeID = val.(string)
			case "drg_id":
				opts[0].DrgID = val.(string)
			case "cpe_id":
				opts[0].CpeID = val.(string)
			case "vcn_id":
				opts[0].VcnID = val.(string)
			case "operating_system":
				opts[0].OperatingSystem = val.(string)
			case "operating_system_version":
				opts[0].OperatingSystemVersion = val.(string)
			default:
				panic(fmt.Sprintf("Unknown key '%s' supplied for Options", key))
			}
		}
	}

	return
}