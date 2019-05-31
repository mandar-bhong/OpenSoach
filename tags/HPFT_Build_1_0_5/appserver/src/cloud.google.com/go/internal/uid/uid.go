// Copyright 2017 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package uid supports generating unique IDs. Its chief purpose is to prevent
// multiple test executions from interfering with each other, and to facilitate
// cleanup of old entities that may remain if tests exit early.
package uid

import (
	"fmt"
	"regexp"
	"strconv"
	"sync/atomic"
	"time"
)

// A Space manages a set of unique IDs distinguished by a prefix.
type Space struct {
	Prefix string    // Prefix of UIDs. Read-only.
	Sep    rune      // Separates UID parts. Read-only.
	Time   time.Time // Timestamp for UIDs. Read-only.
	re     *regexp.Regexp
	count  int32 // atomic
	short  bool
}

// Options are optional values for a Space.
type Options struct {
	Sep  rune      // Separates parts of the UID. Defaults to '-'.
	Time time.Time // Timestamp for all UIDs made with this space. Defaults to current time.

	// Short, if true, makes the result of space.New shorter by 6 characters.
	// This can be useful for character restricted IDs. It will use a shorter
	// but less readable time representation, and will only use two characters
	// for the count suffix instead of four.
	//
	// e.x. normal: gotest-20181030-59751273685000-0001
	// e.x. short:  gotest-1540917351273685000-01
	Short bool
}

// NewSpace creates a new UID space. A UID Space is used to generate unique IDs.
func NewSpace(prefix string, opts *Options) *Space {
	var short bool
	sep := '-'
	tm := time.Now().UTC()
	if opts != nil {
		short = opts.Short
		if opts.Sep != 0 {
			sep = opts.Sep
		}
		if !opts.Time.IsZero() {
			tm = opts.Time
		}
	}
	var re string

	if short {
		re = fmt.Sprintf(`^%s%[2]c(\d+)%[2]c\d+$`, regexp.QuoteMeta(prefix), sep)
	} else {
		re = fmt.Sprintf(`^%s%[2]c(\d{4})(\d{2})(\d{2})%[2]c(\d+)%[2]c\d+$`,
			regexp.QuoteMeta(prefix), sep)
	}

	return &Space{
		Prefix: prefix,
		Sep:    sep,
		Time:   tm,
		re:     regexp.MustCompile(re),
		short:  short,
	}
}

// New generates a new unique ID. The ID consists of the Space's prefix, a
// timestamp, and a counter value. All unique IDs generated in the same test
// execution will have the same timestamp.
//
// Aside from the characters in the prefix, IDs contain only letters, numbers
// and sep.
func (s *Space) New() string {
	c := atomic.AddInt32(&s.count, 1)

	if s.short && c > 99 {
		// Short spaces only have space for 99 IDs. (two characters)
		panic("Short space called New more than 99 times. Ran out of IDs.")
	} else if c > 9999 {
		// Spaces only have space for 9999 IDs. (four characters)
		panic("New called more than 9999 times. Ran out of IDs.")
	}

	if s.short {
		return fmt.Sprintf("%s%c%d%c%02d", s.Prefix, s.Sep, s.Time.UnixNano(), s.Sep, c)
	}

	// Write the time as a date followed by nanoseconds from midnight of that date.
	// That makes it easier to see the approximate time of the ID when it is displayed.
	y, m, d := s.Time.Date()
	ns := s.Time.Sub(time.Date(y, m, d, 0, 0, 0, 0, time.UTC))
	// Zero-pad the counter for lexical sort order for IDs with the same timestamp.
	return fmt.Sprintf("%s%c%04d%02d%02d%c%d%c%04d",
		s.Prefix, s.Sep, y, m, d, s.Sep, ns, s.Sep, c)
}

// Timestamp extracts the timestamp of uid, which must have been generated by
// s. The second return value is true on success, false if there was a problem.
func (s *Space) Timestamp(uid string) (time.Time, bool) {
	subs := s.re.FindStringSubmatch(uid)
	if subs == nil {
		return time.Time{}, false
	}

	if s.short {
		ns, err := strconv.ParseInt(subs[1], 10, 64)
		if err != nil {
			return time.Time{}, false
		}
		return time.Unix(ns/1e9, ns%1e9), true
	}

	y, err1 := strconv.Atoi(subs[1])
	m, err2 := strconv.Atoi(subs[2])
	d, err3 := strconv.Atoi(subs[3])
	ns, err4 := strconv.Atoi(subs[4])
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return time.Time{}, false
	}
	return time.Date(y, time.Month(m), d, 0, 0, 0, ns, time.UTC), true
}

// Older reports whether uid was created by m and has a timestamp older than
// the current time by at least d.
func (s *Space) Older(uid string, d time.Duration) bool {
	ts, ok := s.Timestamp(uid)
	if !ok {
		return false
	}
	return time.Since(ts) > d
}
