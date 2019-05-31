// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: batch_cursor.gen.go.tmpl

package tsm1

import (
	"context"
	"sort"

	"github.com/influxdata/influxdb/query"
	"github.com/influxdata/influxdb/tsdb"
)

// buildFloatCursor creates a cursor for a float field.
func (e *Engine) buildFloatCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) floatCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newFloatCursor(opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}

// buildIntegerCursor creates a cursor for a integer field.
func (e *Engine) buildIntegerCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) integerCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newIntegerCursor(opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}

// buildUnsignedCursor creates a cursor for a unsigned field.
func (e *Engine) buildUnsignedCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) unsignedCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newUnsignedCursor(opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}

// buildStringCursor creates a cursor for a string field.
func (e *Engine) buildStringCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) stringCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newStringCursor(opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}

// buildBooleanCursor creates a cursor for a boolean field.
func (e *Engine) buildBooleanCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) booleanCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newBooleanCursor(opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}

// Cursors

type floatAscendingBatchCursor struct {
	cache struct {
		values Values
		pos    int
	}

	tsm struct {
		buf       []FloatValue
		values    []FloatValue
		pos       int
		keyCursor *KeyCursor
	}

	end int64
	t   []int64
	v   []float64
}

func newFloatAscendingBatchCursor() *floatAscendingBatchCursor {
	return &floatAscendingBatchCursor{
		t: make([]int64, tsdb.DefaultMaxPointsPerBlock),
		v: make([]float64, tsdb.DefaultMaxPointsPerBlock),
	}
}

func (c *floatAscendingBatchCursor) reset(seek, end int64, cacheValues Values, tsmKeyCursor *KeyCursor) {
	c.end = end
	c.cache.values = cacheValues
	c.cache.pos = sort.Search(len(c.cache.values), func(i int) bool {
		return c.cache.values[i].UnixNano() >= seek
	})

	c.tsm.keyCursor = tsmKeyCursor
	c.tsm.values, _ = c.tsm.keyCursor.ReadFloatBlock(&c.tsm.buf)
	c.tsm.pos = sort.Search(len(c.tsm.values), func(i int) bool {
		return c.tsm.values[i].UnixNano() >= seek
	})
}

func (c *floatAscendingBatchCursor) Err() error { return nil }

// close closes the cursor and any dependent cursors.
func (c *floatAscendingBatchCursor) Close() {
	c.tsm.keyCursor.Close()
	c.tsm.keyCursor = nil
	c.cache.values = nil
	c.tsm.values = nil
}

// Next returns the next key/value for the cursor.
func (c *floatAscendingBatchCursor) Next() ([]int64, []float64) {
	var ckey, tkey int64
	var cvalue, tvalue float64

	pos := 0
	for ; pos < cap(c.t); pos++ {
		tkey, tvalue = c.peekTSM()

		if c.cache.pos < len(c.cache.values) {
			ckey, cvalue = c.peekCache()

			var cache, tsm bool

			// Both cache and tsm files have the same key, cache takes precedence.
			if ckey == tkey {
				cache, tsm = true, true
				tkey = ckey
				tvalue = cvalue
			} else if ckey < tkey || tkey == tsdb.EOF {
				// Buffered cache key precedes that in TSM file.
				cache = true
				tkey = ckey
				tvalue = cvalue
			} else {
				// Buffered TSM key precedes that in cache.
				tsm = true
			}

			if cache {
				c.nextCache()
			}

			if tsm {
				c.nextTSM()
			}
		} else {
			if tkey == tsdb.EOF {
				break
			}
			c.nextTSM()
		}

		c.t[pos] = tkey
		c.v[pos] = tvalue
	}

	if pos > 0 && c.t[pos-1] > c.end {
		pos -= 2
		for pos >= 0 && c.t[pos] > c.end {
			pos--
		}
		pos++
	}

	return c.t[:pos], c.v[:pos]
}

// peekCache returns the current time/value from the cache.
func (c *floatAscendingBatchCursor) peekCache() (t int64, v float64) {
	item := c.cache.values[c.cache.pos]
	return item.UnixNano(), item.(FloatValue).value
}

// nextCache returns the next value from the cache.
func (c *floatAscendingBatchCursor) nextCache() {
	if c.cache.pos < len(c.cache.values) {
		c.cache.pos++
	}
}

// peekTSM returns the current time/value from tsm.
func (c *floatAscendingBatchCursor) peekTSM() (t int64, v float64) {
	if c.tsm.pos >= len(c.tsm.values) {
		return tsdb.EOF, 0
	}

	item := c.tsm.values[c.tsm.pos]
	return item.UnixNano(), item.value
}

// nextTSM returns the next value from the TSM files.
func (c *floatAscendingBatchCursor) nextTSM() {
	c.tsm.pos++
	if c.tsm.pos >= len(c.tsm.values) {
		c.tsm.keyCursor.Next()
		c.tsm.values, _ = c.tsm.keyCursor.ReadFloatBlock(&c.tsm.buf)
		c.tsm.pos = 0
	}
}

type floatDescendingBatchCursor struct {
	cache struct {
		values Values
		pos    int
	}

	tsm struct {
		buf       []FloatValue
		values    []FloatValue
		pos       int
		keyCursor *KeyCursor
	}

	end int64
	t   []int64
	v   []float64
}

func newFloatDescendingBatchCursor() *floatDescendingBatchCursor {
	return &floatDescendingBatchCursor{
		t: make([]int64, tsdb.DefaultMaxPointsPerBlock),
		v: make([]float64, tsdb.DefaultMaxPointsPerBlock),
	}
}

func (c *floatDescendingBatchCursor) reset(seek, end int64, cacheValues Values, tsmKeyCursor *KeyCursor) {
	c.end = end
	c.cache.values = cacheValues
	if len(c.cache.values) > 0 {
		c.cache.pos = sort.Search(len(c.cache.values), func(i int) bool {
			return c.cache.values[i].UnixNano() >= seek
		})
		if c.cache.pos == len(c.cache.values) {
			c.cache.pos--
		} else if t, _ := c.peekCache(); t != seek {
			c.cache.pos--
		}
	} else {
		c.cache.pos = -1
	}

	c.tsm.keyCursor = tsmKeyCursor
	c.tsm.values, _ = c.tsm.keyCursor.ReadFloatBlock(&c.tsm.buf)
	c.tsm.pos = sort.Search(len(c.tsm.values), func(i int) bool {
		return c.tsm.values[i].UnixNano() >= seek
	})
	if len(c.tsm.values) > 0 {
		if c.tsm.pos == len(c.tsm.values) {
			c.tsm.pos--
		} else if t, _ := c.peekTSM(); t != seek {
			c.tsm.pos--
		}
	} else {
		c.tsm.pos = -1
	}
}

func (c *floatDescendingBatchCursor) Err() error { return nil }

// close closes the cursor and any dependent cursors.
func (c *floatDescendingBatchCursor) Close() {
	c.tsm.keyCursor.Close()
	c.tsm.keyCursor = nil
	c.cache.values = nil
	c.tsm.values = nil
}

// nextFloat returns the next key/value for the cursor.
func (c *floatDescendingBatchCursor) Next() ([]int64, []float64) {
	var ckey, tkey int64
	var cvalue, tvalue float64

	pos := 0
	for ; pos < cap(c.t); pos++ {
		tkey, tvalue = c.peekTSM()

		if c.cache.pos >= 0 {
			ckey, cvalue = c.peekCache()

			var cache, tsm bool

			// Both cache and tsm files have the same key, cache takes precedence.
			if ckey == tkey {
				cache, tsm = true, true
				tkey = ckey
				tvalue = cvalue
			} else if ckey > tkey || tkey == tsdb.EOF {
				// Buffered cache key succeeds that in TSM file.
				cache = true
				tkey = ckey
				tvalue = cvalue
			} else {
				// Buffered TSM key succeeds that in cache.
				tsm = true
			}

			if cache {
				c.nextCache()
			}

			if tsm {
				c.nextTSM()
			}
		} else {
			if tkey == tsdb.EOF {
				break
			}
			c.nextTSM()
		}

		c.t[pos] = tkey
		c.v[pos] = tvalue
	}

	// strip out remaining points
	if pos > 0 && c.t[pos-1] < c.end {
		pos -= 2
		for pos >= 0 && c.t[pos] < c.end {
			pos--
		}
		pos++
	}

	return c.t[:pos], c.v[:pos]
}

// peekCache returns the current time/value from the cache.
func (c *floatDescendingBatchCursor) peekCache() (t int64, v float64) {
	item := c.cache.values[c.cache.pos]
	return item.UnixNano(), item.(FloatValue).value
}

// nextCache returns the next value from the cache.
func (c *floatDescendingBatchCursor) nextCache() {
	if c.cache.pos >= 0 {
		c.cache.pos--
	}
}

// peekTSM returns the current time/value from tsm.
func (c *floatDescendingBatchCursor) peekTSM() (t int64, v float64) {
	if c.tsm.pos < 0 {
		return tsdb.EOF, 0
	}

	item := c.tsm.values[c.tsm.pos]
	return item.UnixNano(), item.value
}

// nextTSM returns the next value from the TSM files.
func (c *floatDescendingBatchCursor) nextTSM() {
	c.tsm.pos--
	if c.tsm.pos < 0 {
		c.tsm.keyCursor.Next()
		c.tsm.values, _ = c.tsm.keyCursor.ReadFloatBlock(&c.tsm.buf)
		c.tsm.pos = len(c.tsm.values) - 1
	}
}

type integerAscendingBatchCursor struct {
	cache struct {
		values Values
		pos    int
	}

	tsm struct {
		buf       []IntegerValue
		values    []IntegerValue
		pos       int
		keyCursor *KeyCursor
	}

	end int64
	t   []int64
	v   []int64
}

func newIntegerAscendingBatchCursor() *integerAscendingBatchCursor {
	return &integerAscendingBatchCursor{
		t: make([]int64, tsdb.DefaultMaxPointsPerBlock),
		v: make([]int64, tsdb.DefaultMaxPointsPerBlock),
	}
}

func (c *integerAscendingBatchCursor) reset(seek, end int64, cacheValues Values, tsmKeyCursor *KeyCursor) {
	c.end = end
	c.cache.values = cacheValues
	c.cache.pos = sort.Search(len(c.cache.values), func(i int) bool {
		return c.cache.values[i].UnixNano() >= seek
	})

	c.tsm.keyCursor = tsmKeyCursor
	c.tsm.values, _ = c.tsm.keyCursor.ReadIntegerBlock(&c.tsm.buf)
	c.tsm.pos = sort.Search(len(c.tsm.values), func(i int) bool {
		return c.tsm.values[i].UnixNano() >= seek
	})
}

func (c *integerAscendingBatchCursor) Err() error { return nil }

// close closes the cursor and any dependent cursors.
func (c *integerAscendingBatchCursor) Close() {
	c.tsm.keyCursor.Close()
	c.tsm.keyCursor = nil
	c.cache.values = nil
	c.tsm.values = nil
}

// Next returns the next key/value for the cursor.
func (c *integerAscendingBatchCursor) Next() ([]int64, []int64) {
	var ckey, tkey int64
	var cvalue, tvalue int64

	pos := 0
	for ; pos < cap(c.t); pos++ {
		tkey, tvalue = c.peekTSM()

		if c.cache.pos < len(c.cache.values) {
			ckey, cvalue = c.peekCache()

			var cache, tsm bool

			// Both cache and tsm files have the same key, cache takes precedence.
			if ckey == tkey {
				cache, tsm = true, true
				tkey = ckey
				tvalue = cvalue
			} else if ckey < tkey || tkey == tsdb.EOF {
				// Buffered cache key precedes that in TSM file.
				cache = true
				tkey = ckey
				tvalue = cvalue
			} else {
				// Buffered TSM key precedes that in cache.
				tsm = true
			}

			if cache {
				c.nextCache()
			}

			if tsm {
				c.nextTSM()
			}
		} else {
			if tkey == tsdb.EOF {
				break
			}
			c.nextTSM()
		}

		c.t[pos] = tkey
		c.v[pos] = tvalue
	}

	if pos > 0 && c.t[pos-1] > c.end {
		pos -= 2
		for pos >= 0 && c.t[pos] > c.end {
			pos--
		}
		pos++
	}

	return c.t[:pos], c.v[:pos]
}

// peekCache returns the current time/value from the cache.
func (c *integerAscendingBatchCursor) peekCache() (t int64, v int64) {
	item := c.cache.values[c.cache.pos]
	return item.UnixNano(), item.(IntegerValue).value
}

// nextCache returns the next value from the cache.
func (c *integerAscendingBatchCursor) nextCache() {
	if c.cache.pos < len(c.cache.values) {
		c.cache.pos++
	}
}

// peekTSM returns the current time/value from tsm.
func (c *integerAscendingBatchCursor) peekTSM() (t int64, v int64) {
	if c.tsm.pos >= len(c.tsm.values) {
		return tsdb.EOF, 0
	}

	item := c.tsm.values[c.tsm.pos]
	return item.UnixNano(), item.value
}

// nextTSM returns the next value from the TSM files.
func (c *integerAscendingBatchCursor) nextTSM() {
	c.tsm.pos++
	if c.tsm.pos >= len(c.tsm.values) {
		c.tsm.keyCursor.Next()
		c.tsm.values, _ = c.tsm.keyCursor.ReadIntegerBlock(&c.tsm.buf)
		c.tsm.pos = 0
	}
}

type integerDescendingBatchCursor struct {
	cache struct {
		values Values
		pos    int
	}

	tsm struct {
		buf       []IntegerValue
		values    []IntegerValue
		pos       int
		keyCursor *KeyCursor
	}

	end int64
	t   []int64
	v   []int64
}

func newIntegerDescendingBatchCursor() *integerDescendingBatchCursor {
	return &integerDescendingBatchCursor{
		t: make([]int64, tsdb.DefaultMaxPointsPerBlock),
		v: make([]int64, tsdb.DefaultMaxPointsPerBlock),
	}
}

func (c *integerDescendingBatchCursor) reset(seek, end int64, cacheValues Values, tsmKeyCursor *KeyCursor) {
	c.end = end
	c.cache.values = cacheValues
	if len(c.cache.values) > 0 {
		c.cache.pos = sort.Search(len(c.cache.values), func(i int) bool {
			return c.cache.values[i].UnixNano() >= seek
		})
		if c.cache.pos == len(c.cache.values) {
			c.cache.pos--
		} else if t, _ := c.peekCache(); t != seek {
			c.cache.pos--
		}
	} else {
		c.cache.pos = -1
	}

	c.tsm.keyCursor = tsmKeyCursor
	c.tsm.values, _ = c.tsm.keyCursor.ReadIntegerBlock(&c.tsm.buf)
	c.tsm.pos = sort.Search(len(c.tsm.values), func(i int) bool {
		return c.tsm.values[i].UnixNano() >= seek
	})
	if len(c.tsm.values) > 0 {
		if c.tsm.pos == len(c.tsm.values) {
			c.tsm.pos--
		} else if t, _ := c.peekTSM(); t != seek {
			c.tsm.pos--
		}
	} else {
		c.tsm.pos = -1
	}
}

func (c *integerDescendingBatchCursor) Err() error { return nil }

// close closes the cursor and any dependent cursors.
func (c *integerDescendingBatchCursor) Close() {
	c.tsm.keyCursor.Close()
	c.tsm.keyCursor = nil
	c.cache.values = nil
	c.tsm.values = nil
}

// nextInteger returns the next key/value for the cursor.
func (c *integerDescendingBatchCursor) Next() ([]int64, []int64) {
	var ckey, tkey int64
	var cvalue, tvalue int64

	pos := 0
	for ; pos < cap(c.t); pos++ {
		tkey, tvalue = c.peekTSM()

		if c.cache.pos >= 0 {
			ckey, cvalue = c.peekCache()

			var cache, tsm bool

			// Both cache and tsm files have the same key, cache takes precedence.
			if ckey == tkey {
				cache, tsm = true, true
				tkey = ckey
				tvalue = cvalue
			} else if ckey > tkey || tkey == tsdb.EOF {
				// Buffered cache key succeeds that in TSM file.
				cache = true
				tkey = ckey
				tvalue = cvalue
			} else {
				// Buffered TSM key succeeds that in cache.
				tsm = true
			}

			if cache {
				c.nextCache()
			}

			if tsm {
				c.nextTSM()
			}
		} else {
			if tkey == tsdb.EOF {
				break
			}
			c.nextTSM()
		}

		c.t[pos] = tkey
		c.v[pos] = tvalue
	}

	// strip out remaining points
	if pos > 0 && c.t[pos-1] < c.end {
		pos -= 2
		for pos >= 0 && c.t[pos] < c.end {
			pos--
		}
		pos++
	}

	return c.t[:pos], c.v[:pos]
}

// peekCache returns the current time/value from the cache.
func (c *integerDescendingBatchCursor) peekCache() (t int64, v int64) {
	item := c.cache.values[c.cache.pos]
	return item.UnixNano(), item.(IntegerValue).value
}

// nextCache returns the next value from the cache.
func (c *integerDescendingBatchCursor) nextCache() {
	if c.cache.pos >= 0 {
		c.cache.pos--
	}
}

// peekTSM returns the current time/value from tsm.
func (c *integerDescendingBatchCursor) peekTSM() (t int64, v int64) {
	if c.tsm.pos < 0 {
		return tsdb.EOF, 0
	}

	item := c.tsm.values[c.tsm.pos]
	return item.UnixNano(), item.value
}

// nextTSM returns the next value from the TSM files.
func (c *integerDescendingBatchCursor) nextTSM() {
	c.tsm.pos--
	if c.tsm.pos < 0 {
		c.tsm.keyCursor.Next()
		c.tsm.values, _ = c.tsm.keyCursor.ReadIntegerBlock(&c.tsm.buf)
		c.tsm.pos = len(c.tsm.values) - 1
	}
}

type unsignedAscendingBatchCursor struct {
	cache struct {
		values Values
		pos    int
	}

	tsm struct {
		buf       []UnsignedValue
		values    []UnsignedValue
		pos       int
		keyCursor *KeyCursor
	}

	end int64
	t   []int64
	v   []uint64
}

func newUnsignedAscendingBatchCursor() *unsignedAscendingBatchCursor {
	return &unsignedAscendingBatchCursor{
		t: make([]int64, tsdb.DefaultMaxPointsPerBlock),
		v: make([]uint64, tsdb.DefaultMaxPointsPerBlock),
	}
}

func (c *unsignedAscendingBatchCursor) reset(seek, end int64, cacheValues Values, tsmKeyCursor *KeyCursor) {
	c.end = end
	c.cache.values = cacheValues
	c.cache.pos = sort.Search(len(c.cache.values), func(i int) bool {
		return c.cache.values[i].UnixNano() >= seek
	})

	c.tsm.keyCursor = tsmKeyCursor
	c.tsm.values, _ = c.tsm.keyCursor.ReadUnsignedBlock(&c.tsm.buf)
	c.tsm.pos = sort.Search(len(c.tsm.values), func(i int) bool {
		return c.tsm.values[i].UnixNano() >= seek
	})
}

func (c *unsignedAscendingBatchCursor) Err() error { return nil }

// close closes the cursor and any dependent cursors.
func (c *unsignedAscendingBatchCursor) Close() {
	c.tsm.keyCursor.Close()
	c.tsm.keyCursor = nil
	c.cache.values = nil
	c.tsm.values = nil
}

// Next returns the next key/value for the cursor.
func (c *unsignedAscendingBatchCursor) Next() ([]int64, []uint64) {
	var ckey, tkey int64
	var cvalue, tvalue uint64

	pos := 0
	for ; pos < cap(c.t); pos++ {
		tkey, tvalue = c.peekTSM()

		if c.cache.pos < len(c.cache.values) {
			ckey, cvalue = c.peekCache()

			var cache, tsm bool

			// Both cache and tsm files have the same key, cache takes precedence.
			if ckey == tkey {
				cache, tsm = true, true
				tkey = ckey
				tvalue = cvalue
			} else if ckey < tkey || tkey == tsdb.EOF {
				// Buffered cache key precedes that in TSM file.
				cache = true
				tkey = ckey
				tvalue = cvalue
			} else {
				// Buffered TSM key precedes that in cache.
				tsm = true
			}

			if cache {
				c.nextCache()
			}

			if tsm {
				c.nextTSM()
			}
		} else {
			if tkey == tsdb.EOF {
				break
			}
			c.nextTSM()
		}

		c.t[pos] = tkey
		c.v[pos] = tvalue
	}

	if pos > 0 && c.t[pos-1] > c.end {
		pos -= 2
		for pos >= 0 && c.t[pos] > c.end {
			pos--
		}
		pos++
	}

	return c.t[:pos], c.v[:pos]
}

// peekCache returns the current time/value from the cache.
func (c *unsignedAscendingBatchCursor) peekCache() (t int64, v uint64) {
	item := c.cache.values[c.cache.pos]
	return item.UnixNano(), item.(UnsignedValue).value
}

// nextCache returns the next value from the cache.
func (c *unsignedAscendingBatchCursor) nextCache() {
	if c.cache.pos < len(c.cache.values) {
		c.cache.pos++
	}
}

// peekTSM returns the current time/value from tsm.
func (c *unsignedAscendingBatchCursor) peekTSM() (t int64, v uint64) {
	if c.tsm.pos >= len(c.tsm.values) {
		return tsdb.EOF, 0
	}

	item := c.tsm.values[c.tsm.pos]
	return item.UnixNano(), item.value
}

// nextTSM returns the next value from the TSM files.
func (c *unsignedAscendingBatchCursor) nextTSM() {
	c.tsm.pos++
	if c.tsm.pos >= len(c.tsm.values) {
		c.tsm.keyCursor.Next()
		c.tsm.values, _ = c.tsm.keyCursor.ReadUnsignedBlock(&c.tsm.buf)
		c.tsm.pos = 0
	}
}

type unsignedDescendingBatchCursor struct {
	cache struct {
		values Values
		pos    int
	}

	tsm struct {
		buf       []UnsignedValue
		values    []UnsignedValue
		pos       int
		keyCursor *KeyCursor
	}

	end int64
	t   []int64
	v   []uint64
}

func newUnsignedDescendingBatchCursor() *unsignedDescendingBatchCursor {
	return &unsignedDescendingBatchCursor{
		t: make([]int64, tsdb.DefaultMaxPointsPerBlock),
		v: make([]uint64, tsdb.DefaultMaxPointsPerBlock),
	}
}

func (c *unsignedDescendingBatchCursor) reset(seek, end int64, cacheValues Values, tsmKeyCursor *KeyCursor) {
	c.end = end
	c.cache.values = cacheValues
	if len(c.cache.values) > 0 {
		c.cache.pos = sort.Search(len(c.cache.values), func(i int) bool {
			return c.cache.values[i].UnixNano() >= seek
		})
		if c.cache.pos == len(c.cache.values) {
			c.cache.pos--
		} else if t, _ := c.peekCache(); t != seek {
			c.cache.pos--
		}
	} else {
		c.cache.pos = -1
	}

	c.tsm.keyCursor = tsmKeyCursor
	c.tsm.values, _ = c.tsm.keyCursor.ReadUnsignedBlock(&c.tsm.buf)
	c.tsm.pos = sort.Search(len(c.tsm.values), func(i int) bool {
		return c.tsm.values[i].UnixNano() >= seek
	})
	if len(c.tsm.values) > 0 {
		if c.tsm.pos == len(c.tsm.values) {
			c.tsm.pos--
		} else if t, _ := c.peekTSM(); t != seek {
			c.tsm.pos--
		}
	} else {
		c.tsm.pos = -1
	}
}

func (c *unsignedDescendingBatchCursor) Err() error { return nil }

// close closes the cursor and any dependent cursors.
func (c *unsignedDescendingBatchCursor) Close() {
	c.tsm.keyCursor.Close()
	c.tsm.keyCursor = nil
	c.cache.values = nil
	c.tsm.values = nil
}

// nextUnsigned returns the next key/value for the cursor.
func (c *unsignedDescendingBatchCursor) Next() ([]int64, []uint64) {
	var ckey, tkey int64
	var cvalue, tvalue uint64

	pos := 0
	for ; pos < cap(c.t); pos++ {
		tkey, tvalue = c.peekTSM()

		if c.cache.pos >= 0 {
			ckey, cvalue = c.peekCache()

			var cache, tsm bool

			// Both cache and tsm files have the same key, cache takes precedence.
			if ckey == tkey {
				cache, tsm = true, true
				tkey = ckey
				tvalue = cvalue
			} else if ckey > tkey || tkey == tsdb.EOF {
				// Buffered cache key succeeds that in TSM file.
				cache = true
				tkey = ckey
				tvalue = cvalue
			} else {
				// Buffered TSM key succeeds that in cache.
				tsm = true
			}

			if cache {
				c.nextCache()
			}

			if tsm {
				c.nextTSM()
			}
		} else {
			if tkey == tsdb.EOF {
				break
			}
			c.nextTSM()
		}

		c.t[pos] = tkey
		c.v[pos] = tvalue
	}

	// strip out remaining points
	if pos > 0 && c.t[pos-1] < c.end {
		pos -= 2
		for pos >= 0 && c.t[pos] < c.end {
			pos--
		}
		pos++
	}

	return c.t[:pos], c.v[:pos]
}

// peekCache returns the current time/value from the cache.
func (c *unsignedDescendingBatchCursor) peekCache() (t int64, v uint64) {
	item := c.cache.values[c.cache.pos]
	return item.UnixNano(), item.(UnsignedValue).value
}

// nextCache returns the next value from the cache.
func (c *unsignedDescendingBatchCursor) nextCache() {
	if c.cache.pos >= 0 {
		c.cache.pos--
	}
}

// peekTSM returns the current time/value from tsm.
func (c *unsignedDescendingBatchCursor) peekTSM() (t int64, v uint64) {
	if c.tsm.pos < 0 {
		return tsdb.EOF, 0
	}

	item := c.tsm.values[c.tsm.pos]
	return item.UnixNano(), item.value
}

// nextTSM returns the next value from the TSM files.
func (c *unsignedDescendingBatchCursor) nextTSM() {
	c.tsm.pos--
	if c.tsm.pos < 0 {
		c.tsm.keyCursor.Next()
		c.tsm.values, _ = c.tsm.keyCursor.ReadUnsignedBlock(&c.tsm.buf)
		c.tsm.pos = len(c.tsm.values) - 1
	}
}

type stringAscendingBatchCursor struct {
	cache struct {
		values Values
		pos    int
	}

	tsm struct {
		buf       []StringValue
		values    []StringValue
		pos       int
		keyCursor *KeyCursor
	}

	end int64
	t   []int64
	v   []string
}

func newStringAscendingBatchCursor() *stringAscendingBatchCursor {
	return &stringAscendingBatchCursor{
		t: make([]int64, tsdb.DefaultMaxPointsPerBlock),
		v: make([]string, tsdb.DefaultMaxPointsPerBlock),
	}
}

func (c *stringAscendingBatchCursor) reset(seek, end int64, cacheValues Values, tsmKeyCursor *KeyCursor) {
	c.end = end
	c.cache.values = cacheValues
	c.cache.pos = sort.Search(len(c.cache.values), func(i int) bool {
		return c.cache.values[i].UnixNano() >= seek
	})

	c.tsm.keyCursor = tsmKeyCursor
	c.tsm.values, _ = c.tsm.keyCursor.ReadStringBlock(&c.tsm.buf)
	c.tsm.pos = sort.Search(len(c.tsm.values), func(i int) bool {
		return c.tsm.values[i].UnixNano() >= seek
	})
}

func (c *stringAscendingBatchCursor) Err() error { return nil }

// close closes the cursor and any dependent cursors.
func (c *stringAscendingBatchCursor) Close() {
	c.tsm.keyCursor.Close()
	c.tsm.keyCursor = nil
	c.cache.values = nil
	c.tsm.values = nil
}

// Next returns the next key/value for the cursor.
func (c *stringAscendingBatchCursor) Next() ([]int64, []string) {
	var ckey, tkey int64
	var cvalue, tvalue string

	pos := 0
	for ; pos < cap(c.t); pos++ {
		tkey, tvalue = c.peekTSM()

		if c.cache.pos < len(c.cache.values) {
			ckey, cvalue = c.peekCache()

			var cache, tsm bool

			// Both cache and tsm files have the same key, cache takes precedence.
			if ckey == tkey {
				cache, tsm = true, true
				tkey = ckey
				tvalue = cvalue
			} else if ckey < tkey || tkey == tsdb.EOF {
				// Buffered cache key precedes that in TSM file.
				cache = true
				tkey = ckey
				tvalue = cvalue
			} else {
				// Buffered TSM key precedes that in cache.
				tsm = true
			}

			if cache {
				c.nextCache()
			}

			if tsm {
				c.nextTSM()
			}
		} else {
			if tkey == tsdb.EOF {
				break
			}
			c.nextTSM()
		}

		c.t[pos] = tkey
		c.v[pos] = tvalue
	}

	if pos > 0 && c.t[pos-1] > c.end {
		pos -= 2
		for pos >= 0 && c.t[pos] > c.end {
			pos--
		}
		pos++
	}

	return c.t[:pos], c.v[:pos]
}

// peekCache returns the current time/value from the cache.
func (c *stringAscendingBatchCursor) peekCache() (t int64, v string) {
	item := c.cache.values[c.cache.pos]
	return item.UnixNano(), item.(StringValue).value
}

// nextCache returns the next value from the cache.
func (c *stringAscendingBatchCursor) nextCache() {
	if c.cache.pos < len(c.cache.values) {
		c.cache.pos++
	}
}

// peekTSM returns the current time/value from tsm.
func (c *stringAscendingBatchCursor) peekTSM() (t int64, v string) {
	if c.tsm.pos >= len(c.tsm.values) {
		return tsdb.EOF, ""
	}

	item := c.tsm.values[c.tsm.pos]
	return item.UnixNano(), item.value
}

// nextTSM returns the next value from the TSM files.
func (c *stringAscendingBatchCursor) nextTSM() {
	c.tsm.pos++
	if c.tsm.pos >= len(c.tsm.values) {
		c.tsm.keyCursor.Next()
		c.tsm.values, _ = c.tsm.keyCursor.ReadStringBlock(&c.tsm.buf)
		c.tsm.pos = 0
	}
}

type stringDescendingBatchCursor struct {
	cache struct {
		values Values
		pos    int
	}

	tsm struct {
		buf       []StringValue
		values    []StringValue
		pos       int
		keyCursor *KeyCursor
	}

	end int64
	t   []int64
	v   []string
}

func newStringDescendingBatchCursor() *stringDescendingBatchCursor {
	return &stringDescendingBatchCursor{
		t: make([]int64, tsdb.DefaultMaxPointsPerBlock),
		v: make([]string, tsdb.DefaultMaxPointsPerBlock),
	}
}

func (c *stringDescendingBatchCursor) reset(seek, end int64, cacheValues Values, tsmKeyCursor *KeyCursor) {
	c.end = end
	c.cache.values = cacheValues
	if len(c.cache.values) > 0 {
		c.cache.pos = sort.Search(len(c.cache.values), func(i int) bool {
			return c.cache.values[i].UnixNano() >= seek
		})
		if c.cache.pos == len(c.cache.values) {
			c.cache.pos--
		} else if t, _ := c.peekCache(); t != seek {
			c.cache.pos--
		}
	} else {
		c.cache.pos = -1
	}

	c.tsm.keyCursor = tsmKeyCursor
	c.tsm.values, _ = c.tsm.keyCursor.ReadStringBlock(&c.tsm.buf)
	c.tsm.pos = sort.Search(len(c.tsm.values), func(i int) bool {
		return c.tsm.values[i].UnixNano() >= seek
	})
	if len(c.tsm.values) > 0 {
		if c.tsm.pos == len(c.tsm.values) {
			c.tsm.pos--
		} else if t, _ := c.peekTSM(); t != seek {
			c.tsm.pos--
		}
	} else {
		c.tsm.pos = -1
	}
}

func (c *stringDescendingBatchCursor) Err() error { return nil }

// close closes the cursor and any dependent cursors.
func (c *stringDescendingBatchCursor) Close() {
	c.tsm.keyCursor.Close()
	c.tsm.keyCursor = nil
	c.cache.values = nil
	c.tsm.values = nil
}

// nextString returns the next key/value for the cursor.
func (c *stringDescendingBatchCursor) Next() ([]int64, []string) {
	var ckey, tkey int64
	var cvalue, tvalue string

	pos := 0
	for ; pos < cap(c.t); pos++ {
		tkey, tvalue = c.peekTSM()

		if c.cache.pos >= 0 {
			ckey, cvalue = c.peekCache()

			var cache, tsm bool

			// Both cache and tsm files have the same key, cache takes precedence.
			if ckey == tkey {
				cache, tsm = true, true
				tkey = ckey
				tvalue = cvalue
			} else if ckey > tkey || tkey == tsdb.EOF {
				// Buffered cache key succeeds that in TSM file.
				cache = true
				tkey = ckey
				tvalue = cvalue
			} else {
				// Buffered TSM key succeeds that in cache.
				tsm = true
			}

			if cache {
				c.nextCache()
			}

			if tsm {
				c.nextTSM()
			}
		} else {
			if tkey == tsdb.EOF {
				break
			}
			c.nextTSM()
		}

		c.t[pos] = tkey
		c.v[pos] = tvalue
	}

	// strip out remaining points
	if pos > 0 && c.t[pos-1] < c.end {
		pos -= 2
		for pos >= 0 && c.t[pos] < c.end {
			pos--
		}
		pos++
	}

	return c.t[:pos], c.v[:pos]
}

// peekCache returns the current time/value from the cache.
func (c *stringDescendingBatchCursor) peekCache() (t int64, v string) {
	item := c.cache.values[c.cache.pos]
	return item.UnixNano(), item.(StringValue).value
}

// nextCache returns the next value from the cache.
func (c *stringDescendingBatchCursor) nextCache() {
	if c.cache.pos >= 0 {
		c.cache.pos--
	}
}

// peekTSM returns the current time/value from tsm.
func (c *stringDescendingBatchCursor) peekTSM() (t int64, v string) {
	if c.tsm.pos < 0 {
		return tsdb.EOF, ""
	}

	item := c.tsm.values[c.tsm.pos]
	return item.UnixNano(), item.value
}

// nextTSM returns the next value from the TSM files.
func (c *stringDescendingBatchCursor) nextTSM() {
	c.tsm.pos--
	if c.tsm.pos < 0 {
		c.tsm.keyCursor.Next()
		c.tsm.values, _ = c.tsm.keyCursor.ReadStringBlock(&c.tsm.buf)
		c.tsm.pos = len(c.tsm.values) - 1
	}
}

type booleanAscendingBatchCursor struct {
	cache struct {
		values Values
		pos    int
	}

	tsm struct {
		buf       []BooleanValue
		values    []BooleanValue
		pos       int
		keyCursor *KeyCursor
	}

	end int64
	t   []int64
	v   []bool
}

func newBooleanAscendingBatchCursor() *booleanAscendingBatchCursor {
	return &booleanAscendingBatchCursor{
		t: make([]int64, tsdb.DefaultMaxPointsPerBlock),
		v: make([]bool, tsdb.DefaultMaxPointsPerBlock),
	}
}

func (c *booleanAscendingBatchCursor) reset(seek, end int64, cacheValues Values, tsmKeyCursor *KeyCursor) {
	c.end = end
	c.cache.values = cacheValues
	c.cache.pos = sort.Search(len(c.cache.values), func(i int) bool {
		return c.cache.values[i].UnixNano() >= seek
	})

	c.tsm.keyCursor = tsmKeyCursor
	c.tsm.values, _ = c.tsm.keyCursor.ReadBooleanBlock(&c.tsm.buf)
	c.tsm.pos = sort.Search(len(c.tsm.values), func(i int) bool {
		return c.tsm.values[i].UnixNano() >= seek
	})
}

func (c *booleanAscendingBatchCursor) Err() error { return nil }

// close closes the cursor and any dependent cursors.
func (c *booleanAscendingBatchCursor) Close() {
	c.tsm.keyCursor.Close()
	c.tsm.keyCursor = nil
	c.cache.values = nil
	c.tsm.values = nil
}

// Next returns the next key/value for the cursor.
func (c *booleanAscendingBatchCursor) Next() ([]int64, []bool) {
	var ckey, tkey int64
	var cvalue, tvalue bool

	pos := 0
	for ; pos < cap(c.t); pos++ {
		tkey, tvalue = c.peekTSM()

		if c.cache.pos < len(c.cache.values) {
			ckey, cvalue = c.peekCache()

			var cache, tsm bool

			// Both cache and tsm files have the same key, cache takes precedence.
			if ckey == tkey {
				cache, tsm = true, true
				tkey = ckey
				tvalue = cvalue
			} else if ckey < tkey || tkey == tsdb.EOF {
				// Buffered cache key precedes that in TSM file.
				cache = true
				tkey = ckey
				tvalue = cvalue
			} else {
				// Buffered TSM key precedes that in cache.
				tsm = true
			}

			if cache {
				c.nextCache()
			}

			if tsm {
				c.nextTSM()
			}
		} else {
			if tkey == tsdb.EOF {
				break
			}
			c.nextTSM()
		}

		c.t[pos] = tkey
		c.v[pos] = tvalue
	}

	if pos > 0 && c.t[pos-1] > c.end {
		pos -= 2
		for pos >= 0 && c.t[pos] > c.end {
			pos--
		}
		pos++
	}

	return c.t[:pos], c.v[:pos]
}

// peekCache returns the current time/value from the cache.
func (c *booleanAscendingBatchCursor) peekCache() (t int64, v bool) {
	item := c.cache.values[c.cache.pos]
	return item.UnixNano(), item.(BooleanValue).value
}

// nextCache returns the next value from the cache.
func (c *booleanAscendingBatchCursor) nextCache() {
	if c.cache.pos < len(c.cache.values) {
		c.cache.pos++
	}
}

// peekTSM returns the current time/value from tsm.
func (c *booleanAscendingBatchCursor) peekTSM() (t int64, v bool) {
	if c.tsm.pos >= len(c.tsm.values) {
		return tsdb.EOF, false
	}

	item := c.tsm.values[c.tsm.pos]
	return item.UnixNano(), item.value
}

// nextTSM returns the next value from the TSM files.
func (c *booleanAscendingBatchCursor) nextTSM() {
	c.tsm.pos++
	if c.tsm.pos >= len(c.tsm.values) {
		c.tsm.keyCursor.Next()
		c.tsm.values, _ = c.tsm.keyCursor.ReadBooleanBlock(&c.tsm.buf)
		c.tsm.pos = 0
	}
}

type booleanDescendingBatchCursor struct {
	cache struct {
		values Values
		pos    int
	}

	tsm struct {
		buf       []BooleanValue
		values    []BooleanValue
		pos       int
		keyCursor *KeyCursor
	}

	end int64
	t   []int64
	v   []bool
}

func newBooleanDescendingBatchCursor() *booleanDescendingBatchCursor {
	return &booleanDescendingBatchCursor{
		t: make([]int64, tsdb.DefaultMaxPointsPerBlock),
		v: make([]bool, tsdb.DefaultMaxPointsPerBlock),
	}
}

func (c *booleanDescendingBatchCursor) reset(seek, end int64, cacheValues Values, tsmKeyCursor *KeyCursor) {
	c.end = end
	c.cache.values = cacheValues
	if len(c.cache.values) > 0 {
		c.cache.pos = sort.Search(len(c.cache.values), func(i int) bool {
			return c.cache.values[i].UnixNano() >= seek
		})
		if c.cache.pos == len(c.cache.values) {
			c.cache.pos--
		} else if t, _ := c.peekCache(); t != seek {
			c.cache.pos--
		}
	} else {
		c.cache.pos = -1
	}

	c.tsm.keyCursor = tsmKeyCursor
	c.tsm.values, _ = c.tsm.keyCursor.ReadBooleanBlock(&c.tsm.buf)
	c.tsm.pos = sort.Search(len(c.tsm.values), func(i int) bool {
		return c.tsm.values[i].UnixNano() >= seek
	})
	if len(c.tsm.values) > 0 {
		if c.tsm.pos == len(c.tsm.values) {
			c.tsm.pos--
		} else if t, _ := c.peekTSM(); t != seek {
			c.tsm.pos--
		}
	} else {
		c.tsm.pos = -1
	}
}

func (c *booleanDescendingBatchCursor) Err() error { return nil }

// close closes the cursor and any dependent cursors.
func (c *booleanDescendingBatchCursor) Close() {
	c.tsm.keyCursor.Close()
	c.tsm.keyCursor = nil
	c.cache.values = nil
	c.tsm.values = nil
}

// nextBoolean returns the next key/value for the cursor.
func (c *booleanDescendingBatchCursor) Next() ([]int64, []bool) {
	var ckey, tkey int64
	var cvalue, tvalue bool

	pos := 0
	for ; pos < cap(c.t); pos++ {
		tkey, tvalue = c.peekTSM()

		if c.cache.pos >= 0 {
			ckey, cvalue = c.peekCache()

			var cache, tsm bool

			// Both cache and tsm files have the same key, cache takes precedence.
			if ckey == tkey {
				cache, tsm = true, true
				tkey = ckey
				tvalue = cvalue
			} else if ckey > tkey || tkey == tsdb.EOF {
				// Buffered cache key succeeds that in TSM file.
				cache = true
				tkey = ckey
				tvalue = cvalue
			} else {
				// Buffered TSM key succeeds that in cache.
				tsm = true
			}

			if cache {
				c.nextCache()
			}

			if tsm {
				c.nextTSM()
			}
		} else {
			if tkey == tsdb.EOF {
				break
			}
			c.nextTSM()
		}

		c.t[pos] = tkey
		c.v[pos] = tvalue
	}

	// strip out remaining points
	if pos > 0 && c.t[pos-1] < c.end {
		pos -= 2
		for pos >= 0 && c.t[pos] < c.end {
			pos--
		}
		pos++
	}

	return c.t[:pos], c.v[:pos]
}

// peekCache returns the current time/value from the cache.
func (c *booleanDescendingBatchCursor) peekCache() (t int64, v bool) {
	item := c.cache.values[c.cache.pos]
	return item.UnixNano(), item.(BooleanValue).value
}

// nextCache returns the next value from the cache.
func (c *booleanDescendingBatchCursor) nextCache() {
	if c.cache.pos >= 0 {
		c.cache.pos--
	}
}

// peekTSM returns the current time/value from tsm.
func (c *booleanDescendingBatchCursor) peekTSM() (t int64, v bool) {
	if c.tsm.pos < 0 {
		return tsdb.EOF, false
	}

	item := c.tsm.values[c.tsm.pos]
	return item.UnixNano(), item.value
}

// nextTSM returns the next value from the TSM files.
func (c *booleanDescendingBatchCursor) nextTSM() {
	c.tsm.pos--
	if c.tsm.pos < 0 {
		c.tsm.keyCursor.Next()
		c.tsm.values, _ = c.tsm.keyCursor.ReadBooleanBlock(&c.tsm.buf)
		c.tsm.pos = len(c.tsm.values) - 1
	}
}
