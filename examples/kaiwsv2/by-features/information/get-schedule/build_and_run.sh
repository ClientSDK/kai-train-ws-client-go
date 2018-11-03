echo "Clean..."
rm ./information.get_schedule
echo "Build..."
go build -o information.get_schedule main.go 
echo "Build Done..."
echo "Run..."
./information.get_schedule > information.get_schedule-rs.json
echo "Done."