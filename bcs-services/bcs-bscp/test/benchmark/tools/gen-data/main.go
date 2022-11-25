/*
Tencent is pleased to support the open source community by making Basic Service Configuration Platform available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"log"

	"bscp.io/test/client/api"
)

/*
	注：
		1. 生成数据需要调整db隔离级别为读已提交，脚本为并发执行，如果不调整，会死锁
		2. 如果需要导出数据，将 audit 清空
*/

var cli *api.Client

func init() {
	var host string
	flag.StringVar(&host, "host", "http://127.0.0.1:8080", "api server address")
	flag.Parse()

	// init client.
	var err error
	cli, err = api.NewApiClient(host, nil)
	if err != nil {
		log.Printf("new api server client failed, err: %v", err)
		return
	}
}

func main() {
	// batch gen batch data for test.
	if err := genBaseData(); err != nil {
		log.Println(err)
		return
	}

	if err := genSceneData1(); err != nil {
		log.Println(err)
		return
	}

	if err := genSceneData2(); err != nil {
		log.Println(err)
		return
	}

	if err := genSceneData3(); err != nil {
		log.Println(err)
		return
	}

	if err := genSceneData4(); err != nil {
		log.Println(err)
		return
	}

	if err := genSceneData5(); err != nil {
		log.Println(err)
		return
	}

	if err := genSceneData6(); err != nil {
		log.Println(err)
		return
	}
}