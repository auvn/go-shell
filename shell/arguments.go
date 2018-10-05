package shell

import "github.com/pkg/errors"

func StringArgs(args ...interface{}) ([]string, error) {
	stringArgs := make([]string, 0, len(args))

	for i := range args {
		switch arg := args[i].(type) {
		case string:
			stringArgs = append(stringArgs, arg)
		default:
			return nil, errors.Errorf("invalid argument type: %T", arg)
		}
	}
	return stringArgs, nil
}
