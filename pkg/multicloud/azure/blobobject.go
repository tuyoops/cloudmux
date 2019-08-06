// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azure

import (
	"yunion.io/x/onecloud/pkg/cloudprovider"
)

type SObject struct {
	container *SContainer

	cloudprovider.SBaseCloudObject
}

func (o *SObject) GetIBucket() cloudprovider.ICloudBucket {
	return o.container.storageaccount
}

func (o *SObject) GetAcl() cloudprovider.TBucketACLType {
	return cloudprovider.ACLDefault
}

func (o *SObject) SetAcl(aclStr cloudprovider.TBucketACLType) error {
	return nil
}
