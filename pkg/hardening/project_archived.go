/*
 * umoci: Umoci Modifies Open Containers' Images
 * Copyright (C) 2016-2020 SUSE LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package hardening

// NOTE: This project has been moved to the OCI. This build warning exists to
//       tell any downstream users that they should update their import paths.
//       This is necessary because we needed to keep an old version of the
//       umoci source code in the old repo because the Go module name has
//       changed. However, as a result anyone using the old import path will
//       never see new umoci versions.

//#warning This module has been moved to <github.com/opencontainers/umoci>. \
//         Please update your import paths, as <github.com/openSUSE/umoci> \
//         will not receive any further updates.
import "C"
