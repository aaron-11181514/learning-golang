package mapknowledge

// 在遍歷 map 時，為什麼每次的遍歷順序都可能不同？這樣設計的原因是什麼？請寫出一個保證遍歷順序的解決方案。

// unordered data struct designed for efficient lookups rather than maintaining the order of insertion

// 1. using a slice to storage the key or value that want to be sorted
// range this slice to have sorted key and get the value from map

// 2. using slice not using map
