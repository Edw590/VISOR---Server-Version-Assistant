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

package MOD_1

import (
	MOD_5 "EmailSender"
	MOD_7 "GPTCommunicator"
	MOD_6 "OnlineInfoChk"
	MOD_4 "RssFeedNotifier"
	MOD_2 "SMARTChecker"
	MOD_3 "Speech"
	MOD_9 "UserLocator"
	"Utils"
	MOD_8 "WebsiteBackend"
)

// Modules Manager //

const _TIME_SLEEP_S int = 5

var modules_GL []Utils.Module

var _MAP_MOD_NUM_START = map[int]func(modules *Utils.Module){
	Utils.NUM_MOD_SMARTChecker:    MOD_2.Start,
	Utils.NUM_MOD_Speech:          MOD_3.Start,
	Utils.NUM_MOD_RssFeedNotifier: MOD_4.Start,
	Utils.NUM_MOD_EmailSender:     MOD_5.Start,
	Utils.NUM_MOD_OnlineInfoChk:   MOD_6.Start,
	Utils.NUM_MOD_GPTCommunicator: MOD_7.Start,
	Utils.NUM_MOD_WebsiteBackend:  MOD_8.Start,
	Utils.NUM_MOD_UserLocator:     MOD_9.Start,
}

type _MGIModSpecInfo any
var (
	realMain Utils.RealMain = nil
	moduleInfo_GL Utils.ModuleInfo[_MGIModSpecInfo]
)
func Start(modules []Utils.Module) {
	modules_GL = modules
	Utils.ModStartup[_MGIModSpecInfo](realMain, &modules_GL[Utils.NUM_MOD_ModManager])
}
func init() {realMain =
	func(module_stop *bool, moduleInfo_any any) {
		//moduleInfo_GL = moduleInfo_any.(Utils.ModuleInfo[_MGIModSpecInfo])

		// Check all modules' support and put on a list to later warn if there were changes of support or not.
		var mod_support_list [Utils.MODS_ARRAY_SIZE]bool
		for mod_num := 0; mod_num < Utils.MODS_ARRAY_SIZE; mod_num++ {
			mod_support_list[mod_num] = Utils.IsModSupportedMODULES(mod_num)
		}

		for {
			var modules_to_start [Utils.MODS_ARRAY_SIZE]bool
			var modules_to_stop [Utils.MODS_ARRAY_SIZE]bool

			for mod_num := 0; mod_num < Utils.MODS_ARRAY_SIZE; mod_num++ {
				if mod_num == Utils.NUM_MOD_VISOR || mod_num == Utils.NUM_MOD_ModManager {
					continue
				}

				// Only start the modules supported by the server or client depending on the VISOR_SERVER constant.
				if Utils.PersonalConsts_GL.VISOR_SERVER && (Utils.MOD_NUMS_SUPPORT[mod_num] & Utils.MOD_SERVER == 0) {
					continue
				} else if !Utils.PersonalConsts_GL.VISOR_SERVER && (Utils.MOD_NUMS_SUPPORT[mod_num] & Utils.MOD_CLIENT == 0) {
					continue
				}

				var module_supported bool = Utils.IsModSupportedMODULES(mod_num)

				if module_supported {
					if !mod_support_list[mod_num] {
						// Module was not supported and now it is.
						//log.Println("The following module is now supported on this machine: " + mod_name)

						mod_support_list[mod_num] = true
					}
				} else {
					if mod_support_list[mod_num] {
						// Module was not supported and now it is.
						//log.Println("The following module stopped being supported on this machine: " + mod_name)

						mod_support_list[mod_num] = false
					}
				}

				//log.Println("-----------------------")
				//log.Println("Module " + mod_name + " is supported: " + strconv.FormatBool(module_supported))
				//log.Println("Module " + mod_name + " is running: " + strconv.FormatBool(isModRunning(mod_num)))
				//log.Println("Module " + mod_name + " is enabled: " + strconv.FormatBool(modules_GL[mod_num].Enabled))

				if module_supported {
					if !isModRunning(mod_num) && modules_GL[mod_num].Enabled {
						//log.Println("Starting module: " + mod_name)

						modules_to_start[mod_num] = true
					} else if isModRunning(mod_num) && !modules_GL[mod_num].Enabled {
						//log.Println("Stopping module: " + mod_name)

						modules_to_stop[mod_num] = true
					}
				} else {
					if isModRunning(mod_num) {
						//log.Println("Stopping module: " + mod_name)

						modules_to_stop[mod_num] = true
					}
				}
			}

			// Start the modules
			for mod_num := 0; mod_num < Utils.MODS_ARRAY_SIZE; mod_num++ {
				if modules_to_start[mod_num] && modules_GL[mod_num].Enabled {
					modules_GL[mod_num].Stop = false
					var start_func = _MAP_MOD_NUM_START[mod_num]
					if start_func != nil {
						start_func(&modules_GL[mod_num])
					}
				}
			}

			// Stop the modules
			for mod_num := 0; mod_num < Utils.MODS_ARRAY_SIZE; mod_num++ {
				if modules_to_stop[mod_num] {
					modules_GL[mod_num].Stop = true
				}
			}

			//////////////////////////////////////////////////////////////////

			if Utils.WaitWithStop(module_stop, _TIME_SLEEP_S) {
				return
			}
		}
	}
}

func isModRunning(mod_num int) bool {
	return !modules_GL[mod_num].Stopped
}
