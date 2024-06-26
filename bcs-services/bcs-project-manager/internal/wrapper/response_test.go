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

package wrapper

import (
	"fmt"
	"testing"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/internal/util/errorx"
	proto "github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/proto/bcsproject"
)

func TestResponseReflect(t *testing.T) {

	rErr := errorx.NewReadableErr(errorx.PermDeniedErr, "没有权限")

	pingResp := &proto.PingResponse{
		Data: "string",
	}
	err := RenderResponse(pingResp, "test_request_id", rErr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("pingResp, Code: %d, Message: %s, Data: %v", pingResp.Code, pingResp.Message, pingResp.Data)
	fmt.Printf("pingResp, Code: %d, Message: %s, Data: %v", pingResp.Code, pingResp.Message, pingResp.Data)

	healthzResp := &proto.HealthzResponse{
		Data: &proto.HealthzData{
			Status:      "ok",
			MongoStatus: "ok",
		},
	}
	err = RenderResponse(healthzResp, "test_request_id", rErr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("healthzResp, Code: %d, Message: %s, Data: %v", healthzResp.Code, healthzResp.Message, healthzResp.Data)
	fmt.Printf("healthzResp, Code: %d, Message: %s, Data: %v", healthzResp.Code, healthzResp.Message, healthzResp.Data)

	createVariableResp := &proto.CreateVariableResponse{
		Data: &proto.CreateVariableData{
			Id:          "variable-abcde",
			ProjectCode: "test-project",
			Name:        "test",
		},
	}
	err = RenderResponse(createVariableResp, "test_request_id", rErr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("createVariableResp, Code: %d, Message: %s, Data: %v",
		createVariableResp.Code, createVariableResp.Message, createVariableResp.Data)
	fmt.Printf("createVariableResp, Code: %d, Message: %s, Data: %v",
		createVariableResp.Code, createVariableResp.Message, createVariableResp.Data)

	renderVariablesResp := &proto.RenderVariablesResponse{
		Data: []*proto.VariableValue{
			{
				Id:  "test",
				Key: "key",
			},
		},
	}
	err = RenderResponse(renderVariablesResp, "test_request_id", rErr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("renderVariablesResp, Code: %d, Message: %s, Data: %v",
		renderVariablesResp.Code, renderVariablesResp.Message, renderVariablesResp.Data)
	fmt.Printf("renderVariablesResp, Code: %d, Message: %s, Data: %v",
		renderVariablesResp.Code, renderVariablesResp.Message, renderVariablesResp.Data)

	importVariablesResp := &proto.ImportVariablesResponse{}
	err = RenderResponse(importVariablesResp, "test_request_id", rErr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("importVariablesResp, Code: %d, Message: %s", importVariablesResp.Code, importVariablesResp.Message)
	fmt.Printf("importVariablesResp, Code: %d, Message: %s", importVariablesResp.Code, importVariablesResp.Message)
}
