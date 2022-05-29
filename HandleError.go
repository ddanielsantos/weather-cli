package main

func HandleError(err error) {
	if err != nil {
		print("Error: " + err.Error())
	}
}