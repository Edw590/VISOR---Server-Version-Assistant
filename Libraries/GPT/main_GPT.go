/*******************************************************************************
 * Copyright 2023-2024 Edw590
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 ******************************************************************************/

package main

import (
	"GPT/GPT"
	"Utils"
	"log"
	"time"
)

func main() {
	Utils.PersonalConsts_GL.Init()

	GPT.SetWebsiteInfo(Utils.PersonalConsts_GL.WEBSITE_URL, Utils.PersonalConsts_GL.WEBSITE_PW)

	GPT.SendText("Hi there!")
	for {
		log.Println(GPT.GetEntry(-1))

		time.Sleep(1 * time.Second)
	}
}