package lrfu

import (
	"errors"
	"math"
	"time"
)

/*------------------ DOMAIN TYPES ------------------*/

type cachePolicy func(entry *entry)

/*------------------ CACHE DEFINITIONS ------------------*/

// Cache min-heap tree cache utilizing a LRFU eviction policy as per
// the LRFU papers specification.
type Cache struct {
	// Backing data-structure for operations on cache entries.
	entries arrayHeap

	// Cache metadata. Note that this data is not based on cache policy
	// but instead during cache operations.
	maxSize  uint
	currSize uint
}

// NewCache returns an instance of cache.
func NewCache(maxSize uint) Cache {
	return Cache{
		entries:  newArrayHeap(maxSize),
		maxSize:  maxSize,
		currSize: 0,
	}
}

// Insert a new piece of data into the cache utilizing given key.
// Returns error if data does not fit into cache.
func (cache *Cache) Insert(key interface{}, data *[]byte) error {
	var _ entryKey = key
	var _ entryData = data

	return nil
}

// Fetch data from cache entry identified by key. Nil if no entry
// with given key is found.
func (cache *Cache) Fetch(key interface{}) *[]byte {
	var _ entryKey = key

	return nil

}

// Contains returns true if cache entry with given key exists in cache.
// False otherwise. This operation does not increment an entries reference.
func (cache *Cache) Contains(key interface{}) bool {
	var _ entryKey = key

	return false

}

// Eject entry with given key from cache. The data of the entry is returned
// if the entries data was dirty, nil otherwise.
func (cache *Cache) Eject() *[]byte {
	return nil
}

// Len returns the number of entries in the cache.
func (cache *Cache) Len() int {
	return len(cache.entries)
}

// Size returns the current number of bytes allocated by the cache.
func (cache *Cache) Size() uint {
	return cache.currSize
}

// Capacity returns the maximum number of allocatable bytes by the cache.
func (cache *Cache) Capacity() uint {
	return cache.maxSize
}

// Purge flushes the entire cache, returning any dirty entries.
func (cache *Cache) Purge() {

}

/*------------------ LRFU DEFINITIONS ------------------*/
// Return an instance of the LRFU policy configured with the given parameters.
// Lambda must be in the range of [0, 1] where LRFU will behave as LFU at 0,
// and LRU at 1. Lambda values between the extremes will yield an LRFU policy
// between the two. If lambda is out of range, an error is returned.
func lrfuPolicy(lambda float64) error {
	if lambda < 0 || lambda > 1 {
		return errors.New("lambda must be between 0 and 1 inclusive")
	}

	_ = func(refTime refStamp, crf crfValue) crfValue {
		currTime := time.Now().Unix()

		weightFunction := func(refSpan int64) float64 {
			//Wish there was a way to inline this
			if refSpan == 0 {
				return 1
			}

			return math.Pow(1>>2, lambda*float64(refSpan))
		}

		return crfValue(weightFunction(0) + weightFunction(currTime-int64(refTime))*float64(crf))
	}

	return nil
}
