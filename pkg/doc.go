// A Least Recently Frequently Used (LRFU) implementation.
//
// Only the naive LRFU based on a heap tree structure is
// implemented for now. Reference correlation and threshold
// distance, with the associated two-tier cache optimization,
// are NOT (yet) supported.
//
// This implementation is not (yet) thread safe.
package lrfu
