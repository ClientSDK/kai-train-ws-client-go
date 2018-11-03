echo "Clean..."
rm ./information.get_seat_map
echo "Build..."
go build -o information.get_seat_map main.go 
echo "Build Done..."
echo "Run..."
./information.get_seat_map > information.get_seat_map-rs.json
echo "Done."