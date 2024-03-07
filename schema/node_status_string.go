// Code generated by "enumer --values --type=NodeStatus --linecomment --output node_status_string.go --json --yaml --sql"; DO NOT EDIT.

package schema

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _NodeStatusName = "registeredonlineofflineexiting"

var _NodeStatusIndex = [...]uint8{0, 10, 16, 23, 30}

const _NodeStatusLowerName = "registeredonlineofflineexiting"

func (i NodeStatus) String() string {
	if i < 0 || i >= NodeStatus(len(_NodeStatusIndex)-1) {
		return fmt.Sprintf("NodeStatus(%d)", i)
	}
	return _NodeStatusName[_NodeStatusIndex[i]:_NodeStatusIndex[i+1]]
}

func (NodeStatus) Values() []string {
	return NodeStatusStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NodeStatusNoOp() {
	var x [1]struct{}
	_ = x[NodeStatusRegistered-(0)]
	_ = x[NodeStatusOnline-(1)]
	_ = x[NodeStatusOffline-(2)]
	_ = x[NodeStatusExited-(3)]
}

var _NodeStatusValues = []NodeStatus{NodeStatusRegistered, NodeStatusOnline, NodeStatusOffline, NodeStatusExited}

var _NodeStatusNameToValueMap = map[string]NodeStatus{
	_NodeStatusName[0:10]:       NodeStatusRegistered,
	_NodeStatusLowerName[0:10]:  NodeStatusRegistered,
	_NodeStatusName[10:16]:      NodeStatusOnline,
	_NodeStatusLowerName[10:16]: NodeStatusOnline,
	_NodeStatusName[16:23]:      NodeStatusOffline,
	_NodeStatusLowerName[16:23]: NodeStatusOffline,
	_NodeStatusName[23:30]:      NodeStatusExited,
	_NodeStatusLowerName[23:30]: NodeStatusExited,
}

var _NodeStatusNames = []string{
	_NodeStatusName[0:10],
	_NodeStatusName[10:16],
	_NodeStatusName[16:23],
	_NodeStatusName[23:30],
}

// NodeStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NodeStatusString(s string) (NodeStatus, error) {
	if val, ok := _NodeStatusNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NodeStatusNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to NodeStatus values", s)
}

// NodeStatusValues returns all values of the enum
func NodeStatusValues() []NodeStatus {
	return _NodeStatusValues
}

// NodeStatusStrings returns a slice of all String values of the enum
func NodeStatusStrings() []string {
	strs := make([]string, len(_NodeStatusNames))
	copy(strs, _NodeStatusNames)
	return strs
}

// IsANodeStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i NodeStatus) IsANodeStatus() bool {
	for _, v := range _NodeStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for NodeStatus
func (i NodeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for NodeStatus
func (i *NodeStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("NodeStatus should be a string, got %s", data)
	}

	var err error
	*i, err = NodeStatusString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for NodeStatus
func (i NodeStatus) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for NodeStatus
func (i *NodeStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = NodeStatusString(s)
	return err
}

func (i NodeStatus) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *NodeStatus) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of NodeStatus: %[1]T(%[1]v)", value)
	}

	val, err := NodeStatusString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
