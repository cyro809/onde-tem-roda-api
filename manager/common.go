package manager

import "fmt"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkAffectedRows(affec int64) {
	if affec > 0 {
		fmt.Println(fmt.Printf("Affected %d rows", affec))
	} else {
		fmt.Println("No rows were affected")
	}
}
