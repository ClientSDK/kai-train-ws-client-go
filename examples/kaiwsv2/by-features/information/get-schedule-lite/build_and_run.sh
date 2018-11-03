echo "Clean..."
rm ./information.get_schedule_lite
echo "Build..."
go build -o information.get_schedule_lite main.go 
echo "Build Done..."
echo "Run..."
./information.get_schedule_lite > information.get_schedule_lite-rs.json
echo "Done."