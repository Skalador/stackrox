package flags

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

const (
	timeoutFlagName = "timeout"
)

// AddTimeoutWithDefault adds a timeout flag to the given command, with the given default.
func AddTimeoutWithDefault(c *cobra.Command, defaul time.Duration) {
	c.PersistentFlags().DurationP(timeoutFlagName, "t", defaul, "timeout for API requests")
}

// AddTimeout adds a timeout flag to the given command, with the global default value.
func AddTimeout(c *cobra.Command) {
	AddTimeoutWithDefault(c, 10*time.Second)
}

// Timeout returns the set timeout.
func Timeout(c *cobra.Command) time.Duration {
	duration, err := c.Flags().GetDuration(timeoutFlagName)
	if err != nil {
		// This is a programming error. You shouldn't use the timeout flag unless you've added it to your command!
		// This helps us fail explicitly instead of defaulting to a zero timeout and allowing people to think it worked.
		panic(fmt.Sprintf("command does not have a timeout flag: %v", err))
	}
	return duration
}
