package main

import "fmt"

/* input:
VISIT ***acm.ashland.edu/
VISIT ***acm.ashland.edu/ecmicpc/
BACK
BACK
BACK
FORWARD
VISIT ***###.ibm.com/
BACK
BACK
FORWARD
FORWARD
FORWARD
QUIT
 */
func main() {
	backward := make([]string, 0)
	forward := make([]string, 0)
	forLabel:
	for true {
		var in string
		fmt.Scan(&in)
		var p string
		switch in {
		case "VISIT":
			fmt.Scan(&p)
			forward = forward[:0]
			backward = append(backward, p)
		case "BACK":
			if len(backward) <= 1 {
				p = "Ignored"
			} else {
				p = backward[len(backward)-2]
				forward = append(forward, backward[len(backward)-1])
				backward = backward[:len(backward)-1]
			}
		case "FORWARD":
			if len(forward) == 0 {
				p = "Ignored"
			} else {
				p = forward[len(forward) -1]
				forward = forward[:len(forward)-1]
				backward = append(backward, p)
			}
		case "QUIT":
			break forLabel
		default:
			p = "ERROR"
		}
		fmt.Println(p)
	}
}