// this file was generated by gomacro command: import _b "os/signal"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"os/signal"
)

// reflection: allow interpreted code to import "os/signal"
func init() {
	Packages["os/signal"] = Package{
	Binds: map[string]Value{
		"Ignore":	ValueOf(signal.Ignore),
		"Notify":	ValueOf(signal.Notify),
		"Reset":	ValueOf(signal.Reset),
		"Stop":	ValueOf(signal.Stop),
	}, 
	}
}
