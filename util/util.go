package util

import "fmt"

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("t add --task TASKNAME --start STARTTIME --end ENDTIME --comment COMMENT")
	fmt.Println("\tEither --start or --end is mandatory. Both can be specified at the same call")
	fmt.Println("\t--comment is optional")
	fmt.Println("t print")
	fmt.Println("\tPrints all log from current month.")
}
