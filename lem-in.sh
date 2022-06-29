#!bin/sh

printf "Audit tests below:"
printf "\n"

go run . example01.txt
go run . example02.txt
go run . example03.txt
go run . example04.txt
go run . example05.txt
go run . example06.txt
go run . example07.txt
go run . badexample01.txt
go run . badexample02.txt