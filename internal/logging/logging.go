/*
   Copyright © 2021 Arsen Musayelyan

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

package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Hook(FatalHook{})

// Fatal hook to run in case of Fatal error
type FatalHook struct {
	WorkDir string
}

// Run function on trigger
func (hook FatalHook) Run(_ *zerolog.Event, level zerolog.Level, _ string) {
	// If log event is fatal
	if level == zerolog.FatalLevel {
		// Attempt removal of opensend directory
		_ = os.RemoveAll(hook.WorkDir)
	}
}
