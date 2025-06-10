package greedy

import "sort"

// Leetcode #846
func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand)%groupSize != 0 {
        return false
    }

    // Frequency map to count occurrences of each number
    freq := make(map[int]int)
    for _, n := range hand {
        freq[n]++
    }

    // Extract keys and sort them
    keys := make([]int, 0, len(freq))
    for k := range freq {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    // Process until all elements are used
    for len(keys) > 0 {
        first := keys[0]

        // Attempt to form a consecutive group starting from 'first'
        for i := first; i < first+groupSize; i++ {
            count, ok := freq[i]
            if !ok {
                return false
            }
            if count == 1 {
                // Remove element from map and keys slice
                delete(freq, i)
                // Remove i from keys slice
                idx := sort.SearchInts(keys, i)
                keys = append(keys[:idx], keys[idx+1:]...)
            } else {
                freq[i]--
            }
        }
    }

    return true
}
