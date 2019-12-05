package constants

/*
 Copyright (c) 2019 Dell Inc, or its subsidiaries.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

const (
	// PluginName is the name of the CSI plug-in.
	PluginName = "csi-isilon.dellemc.com"

	// DefaultAccessZone is "System"
	DefaultAccessZone = "System"
	// ModeNode is csi driver's "mode "deployment mode
	ModeNode = "node"
	// ModeController is csi driver's "controller "deployment mode
	ModeController = "controller"

	// DefaultVolumeSizeInBytes is default volume sgolang/protobuf/blob/master/ptypesize to create on an Isilon
	// cluster when no size is given, expressed in bytes
	DefaultVolumeSizeInBytes = 3 * BytesInGiB

	// BytesInGiB is the number of bytes in a gibibyte
	BytesInGiB = 1024 * 1024 * 1024
	// TRUE constant
	TRUE = "TRUE"
	// FALSE constant
	FALSE = "FALSE"

	// DefaultPortNumber is the port number in default to set the HTTPS port number of the Isilon OneFS API server
	DefaultPortNumber = "8080"

	// DefaultIsiPath is the default isiPath which will be used if there's
	// no proper isiPath value set in neither storageclass.yaml nor values.yaml
	DefaultIsiPath = "/ifs"
)
