
package main



func main() {
	done := make(chan bool)
	// doSort is a lambda function, so a closure which knows the channel done:
	doSort := func(s []int){
		sort(s)
		done <- true
	}
	i := pivot(s)
	go doSort(s[:i])
	go doSort(s[i:])
	<-done
	<-done
}