package main
import "errors"
const (
	planFree = "free"
	planPro = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == "pro" {
		rmessages := messages[:]
		return rmessages, nil
	} else if plan == "free" {
		rmessages := messages[:2]
		return rmessages, nil
	} else {
		errornew := errors.New("unsupported plan")
		return nil, errornew
	}
}