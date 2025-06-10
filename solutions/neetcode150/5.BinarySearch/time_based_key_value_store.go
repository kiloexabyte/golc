package main

// Leetcode #981
// TimeMap stores the key-value pairs with timestamps
type TimeMap struct {
	store map[string][]TimeValue
}

// TimeValue holds the value and timestamp as a pair
type TimeValue struct {
	value     string
	timestamp int
}

// Constructor initializes a new TimeMap
func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]TimeValue),
	}
}

// Set stores the value along with the timestamp for the given key
func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, exists := this.store[key]; !exists {
		this.store[key] = []TimeValue{}
	}
	this.store[key] = append(this.store[key], TimeValue{value: value, timestamp: timestamp})
}

// Get retrieves the value corresponding to the given key and timestamp (or closest earlier timestamp)
func (this *TimeMap) Get(key string, timestamp int) string {
	values, exists := this.store[key]
	if !exists {
		return ""
	}

	// Binary search for the closest timestamp
	l, r := 0, len(values)-1
	res := ""
	for l <= r {
		mid := (l + r) / 2
		if values[mid].timestamp <= timestamp {
			res = values[mid].value
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}