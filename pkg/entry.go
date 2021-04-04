package lrfu

import "time"

/*------------------ DOMAIN TYPES ------------------*/

type crfValue float64     // crfValue type represents captures a cache entries CRF value.
type refStamp int64       // refStamp type captures the time at which a reference occured.
type entryData *[]byte    // entryData type captures the data value associated with a cache entry.
type entryKey interface{} // entryKey type captures values used to identify one entry uniquely from another entry.

/*------------------ ENTRY DEFINITIONS ------------------*/

// Entry contains the cache entries data and key, as well LRFU policy required attributes such
// as time of last reference and associated CRF (combined recency frequency).
type entry struct {
	// Fundamental cache entry attributes. These attributes would
	// be necessary for a cache to be of any use, indifferent of policy.
	key   entryKey
	data  entryData
	dirty bool

	// LRFU policy based attributes. Each entry is required to have these values
	// for determining eviction candidacy by LRFU.
	lastReference refStamp
	crf           crfValue
}

// New cache entry returns a new cache entry initialized with default state.
func newEntry(key interface{}, data *[]byte) entry {

	return entry{
		key:           key,
		data:          data,
		dirty:         false,
		lastReference: refStamp(time.Now().Unix()),
		crf:           crfValue(0),
	}
}
