/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/criteria/errf"
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/dal/table"
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/kit"
)

// Event defines a event's details info.
type Event struct {
	Spec       *table.EventSpec       `json:"spec"`
	Attachment *table.EventAttachment `json:"attachment"`
	Revision   *table.CreatedRevision `json:"revision"`
}

// Validate an event is valid or not.
func (e Event) Validate(kit *kit.Kit) error {
	if e.Spec == nil {
		return errf.ErrNoSpecF(kit)
	}

	if err := e.Spec.Validate(kit); err != nil {
		return err
	}

	if e.Attachment == nil {
		return errf.ErrNoAttachmentF(kit)
	}

	if err := e.Attachment.Validate(kit); err != nil {
		return err
	}

	if e.Revision == nil {
		return errf.ErrNoRevisionF(kit)
	}

	if err := e.Revision.Validate(kit); err != nil {
		return err
	}

	return nil
}
