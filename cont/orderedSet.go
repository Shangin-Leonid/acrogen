package cont /* Containers */

// TODO docs
type OrderedSet[T comparable] struct { // TODO maybe not "comparable"
	set      Set[*zeroLevelSkipListNode] // For searching; stores pointers to zero-level list nodes
	skipList SkipList[T]                 // For ordering; stores original data
}

// TODO
NewSet(map[value]score)
Insert
Delete
Walk
WalkInOrder
GetScore
