echo "Clean..."
rm ./information.get_seat_null
echo "Build..."
go build -o information.get_seat_null main.go 
echo "Build Done..."
echo "Run..."
./information.get_seat_null > information.get_seat_null-rs.json
echo "Done."