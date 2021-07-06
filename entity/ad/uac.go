package ad

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"text/tabwriter"

	"github.com/audibleblink/msldapuac"
)

type UserAccountControl int

func NewUAC(uacStr string) (uac UserAccountControl, err error) {
	uacInt, err := strconv.ParseInt(uacStr, 10, 64)
	if err != nil {
		return
	}

	return UserAccountControl(uacInt), nil
}

// GetFlagNames returns the string representation of all set UAC flags
func (uac UserAccountControl) GetFlagNames() ([]string, error) {
	return msldapuac.ParseUAC(int64(uac))
}

// IsFlagSet
func (uac UserAccountControl) IsFlagSet(flag int) (bool, error) {
	return msldapuac.IsSet(int64(uac), flag)
}

func GetUACFlagName(flag int) string {
	return msldapuac.PropertyMap[flag]
}

// UACPrint prints the available UAC options that are available for searching
func UACPrint(dest io.Writer) {
	w := new(tabwriter.Writer)
	w.Init(dest, 8, 8, 0, '\t', 0)
	defer w.Flush()

	template := "%s\t%d\n"
	var sorted []string
	for k, v := range msldapuac.PropertyMap {
		sorted = append(sorted, fmt.Sprintf(template, v, k))
	}

	sort.Strings(sorted)
	fmt.Fprint(w, "Property\tValue\n")
	fmt.Fprint(w, "---\t---\n")
	for _, line := range sorted {
		fmt.Fprint(w, line)
	}
}
