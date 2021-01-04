// Code generated by "stringer -type=connState"; DO NOT EDIT.

package telegram

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[connCreated-0]
	_ = x[connConnecting-1]
	_ = x[connConnected-2]
	_ = x[connInitializing-3]
	_ = x[connIdle-4]
	_ = x[connActive-5]
	_ = x[connReconnecting-6]
	_ = x[connClosing-7]
	_ = x[connClosed-8]
}

const _connState_name = "connCreatedconnConnectingconnConnectedconnInitializingconnIdleconnActiveconnReconnectingconnClosingconnClosed"

var _connState_index = [...]uint8{0, 11, 25, 38, 54, 62, 72, 88, 99, 109}

func (i connState) String() string {
	if i >= connState(len(_connState_index)-1) {
		return "connState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _connState_name[_connState_index[i]:_connState_index[i+1]]
}