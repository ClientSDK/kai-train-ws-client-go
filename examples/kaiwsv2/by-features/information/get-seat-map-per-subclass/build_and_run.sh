echo "Clean..."
rm ./information.get_seat_map_per_subclass
echo "Build..."
go build -o information.get_seat_map_per_subclass main.go 
echo "Build Done..."
echo "Run..."
./information.get_seat_map_per_subclass > information.get_seat_map_per_subclass-rs.json
echo "Done."