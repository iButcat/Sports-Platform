# this script is deleting the old binary file
# and then build and execute it

clear
printf "Compiling and Running..."
FILE=main
if test -f "$FILE"; then
  rm -f "$FILE"
else
  continue
fi
go build main.go
clear
./main
