// Code generated by "stringer -type PieMatchResult"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Win-0]
	_ = x[Loss-1]
	_ = x[Tie-2]
	_ = x[Bye-3]
}

const _PieMatchResult_name = "WinLossTieBye"

var _PieMatchResult_index = [...]uint8{0, 3, 7, 10, 13}

func (i PieMatchResult) String() string {
	if i < 0 || i >= PieMatchResult(len(_PieMatchResult_index)-1) {
		return "PieMatchResult(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PieMatchResult_name[_PieMatchResult_index[i]:_PieMatchResult_index[i+1]]
}