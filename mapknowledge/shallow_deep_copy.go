package mapknowledge

//  map assignment is shallow copy
// what is shallow copy
// it is share the same reference for inner value
// if change the new one map like newMap["apple"] = 5  and the original one will change too

// what is deep copy
// the new one and the old one is total independence

// how to deep copy

// range the map, if the value is complexity value like map or slice , process recursively
// dont need to care about key, because the key of map is only accept the comparably type
