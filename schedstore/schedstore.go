// Copyright 2016 Netflix, Inc.
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

package schedstore

import (
	"errors"
	"time"

	"github.com/netflix/chaosmonkey/schedule"
)

// ErrAlreadyExists is returned when calling Publish if a schedule already
// exists
var ErrAlreadyExists = errors.New("schedule already exists")

// SchedStore stores schedule of terminations
type SchedStore interface {
	// Retrieve retrieves the schedule for the given date
	// The date must be in the local time zone
	Retrieve(date time.Time) (*schedule.Schedule, error)

	// Publish publishes the schedule for the given date
	// The date must be in the local time zone
	Publish(date time.Time, sched *schedule.Schedule) error
}