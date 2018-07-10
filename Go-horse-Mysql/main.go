// main.go

package main

func main() {
	a := App{}
	a.Initialize("root", "Debian23@", "gobase")

	a.Run(":8080")
}
