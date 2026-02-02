# Data structures in Go

Here are some simple implementation. For more: 

GoDS (Go Data Structures) - [https://github.com/emirpasic/gods/](https://github.com/emirpasic/gods/)

TheAlgorithms - [https://github.com/TheAlgorithms/Go/tree/master/structure](https://github.com/TheAlgorithms/Go/tree/master/structure)

## Available Structures

### MultiSet (multiset.go)

A sorted multiset (bag) data structure built on top of GoDS RedBlackTree. Supports:
- O(log n) insertion and deletion
- O(log n) min/max access
- O(1) count lookups and contains checks
- Allows duplicate elements while maintaining sorted order

### KSmallestTracker (ksmallesttracker.go)

Efficiently maintains the k smallest elements from a dynamic set with O(1) sum queries.
Built using two MultiSets to track k smallest and remaining elements separately. Supports:
- O(log n) insertion and deletion
- O(1) sum of k smallest elements
- Automatic rebalancing between sets
- Useful for sliding window problems requiring k smallest elements
