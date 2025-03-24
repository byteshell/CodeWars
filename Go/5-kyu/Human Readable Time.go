package kata

import "fmt"

func HumanReadableTime(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	remainingSeconds := seconds % 60
	
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, remainingSeconds)
}
