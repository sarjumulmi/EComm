package main

func main() {
	a := App{}
	a.Initialize("root", "Oracle123", "EComm")
	a.Run(":6000")
}
