package file

import "os"

const (
	javaCodeRetractSpace_0 = ""
	javaCodeRetractSpace_1 = "    "
	javaCodeRetractSpace_2 = "        "
)

var pathSeparator string

func init() {
	pathSeparator = string(os.PathSeparator)
}
