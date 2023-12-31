// Code generated by "enumer -type=Verbosity -json -text -yaml -transform=snake"; DO NOT EDIT.

package log

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _VerbosityName = "silenterrorinfowarningtrace"

var _VerbosityIndex = [...]uint8{0, 6, 11, 15, 22, 27}

const _VerbosityLowerName = "silenterrorinfowarningtrace"

func (i Verbosity) String() string {
	if i < 0 || i >= Verbosity(len(_VerbosityIndex)-1) {
		return fmt.Sprintf("Verbosity(%d)", i)
	}
	return _VerbosityName[_VerbosityIndex[i]:_VerbosityIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _VerbosityNoOp() {
	var x [1]struct{}
	_ = x[Silent-(0)]
	_ = x[Error-(1)]
	_ = x[Info-(2)]
	_ = x[Warning-(3)]
	_ = x[Trace-(4)]
}

var _VerbosityValues = []Verbosity{Silent, Error, Info, Warning, Trace}

var _VerbosityNameToValueMap = map[string]Verbosity{
	_VerbosityName[0:6]:        Silent,
	_VerbosityLowerName[0:6]:   Silent,
	_VerbosityName[6:11]:       Error,
	_VerbosityLowerName[6:11]:  Error,
	_VerbosityName[11:15]:      Info,
	_VerbosityLowerName[11:15]: Info,
	_VerbosityName[15:22]:      Warning,
	_VerbosityLowerName[15:22]: Warning,
	_VerbosityName[22:27]:      Trace,
	_VerbosityLowerName[22:27]: Trace,
}

var _VerbosityNames = []string{
	_VerbosityName[0:6],
	_VerbosityName[6:11],
	_VerbosityName[11:15],
	_VerbosityName[15:22],
	_VerbosityName[22:27],
}

// VerbosityString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func VerbosityString(s string) (Verbosity, error) {
	if val, ok := _VerbosityNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _VerbosityNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Verbosity values", s)
}

// VerbosityValues returns all values of the enum
func VerbosityValues() []Verbosity {
	return _VerbosityValues
}

// VerbosityStrings returns a slice of all String values of the enum
func VerbosityStrings() []string {
	strs := make([]string, len(_VerbosityNames))
	copy(strs, _VerbosityNames)
	return strs
}

// IsAVerbosity returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Verbosity) IsAVerbosity() bool {
	for _, v := range _VerbosityValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Verbosity
func (i Verbosity) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Verbosity
func (i *Verbosity) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Verbosity should be a string, got %s", data)
	}

	var err error
	*i, err = VerbosityString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for Verbosity
func (i Verbosity) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Verbosity
func (i *Verbosity) UnmarshalText(text []byte) error {
	var err error
	*i, err = VerbosityString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for Verbosity
func (i Verbosity) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Verbosity
func (i *Verbosity) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = VerbosityString(s)
	return err
}
