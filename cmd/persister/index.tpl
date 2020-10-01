package main

var Index []byte = []byte{ {{range $i, $a := .}} {{$a}},{{end}} }
