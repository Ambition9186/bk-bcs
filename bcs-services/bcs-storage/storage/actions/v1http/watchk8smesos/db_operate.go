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

// Package watchk8smesos xxx
package watchk8smesos

import (
	"context"
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/apiserver"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/store"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/types"
)

// PutData 更新/新增
func PutData(ctx context.Context, newObj *types.RawObject, opt *store.UpdateOptions) error {
	// 创建连接
	db := apiserver.GetAPIResource().GetStoreClient(dbConfig)

	// 更新数据
	return db.Update(ctx, newObj, opt)
}

// RemoveDta 删除数据
func RemoveDta(ctx context.Context, newObj *types.RawObject, opt *store.DeleteOptions) error {
	// 创建连接
	db := apiserver.GetAPIResource().GetStoreClient(dbConfig)

	// 删除数据
	return db.Delete(ctx, newObj, opt)
}

// GetData 查询数据
func GetData(ctx context.Context, objectType types.ObjectType, key types.ObjectKey, opt *store.GetOptions) (
	*types.RawObject, error) {
	// 创建连接
	db := apiserver.GetAPIResource().GetStoreClient(dbConfig)

	// 获取数据
	return db.Get(ctx, objectType, key, opt)
}

// GetList 获取list数据
func GetList(ctx context.Context, objectType types.ObjectType, opts *store.ListOptions) (retList []string, err error) {
	// 创建连接
	db := apiserver.GetAPIResource().GetStoreClient(dbConfig)

	// 获取数据
	objList, err := db.List(ctx, objectType, opts)
	if err != nil {
		return nil, err
	}

	for _, obj := range objList {
		retList = append(retList, fmt.Sprintf("%s.%s", obj.GetNamespace(), obj.GetName()))
	}

	return retList, nil
}
