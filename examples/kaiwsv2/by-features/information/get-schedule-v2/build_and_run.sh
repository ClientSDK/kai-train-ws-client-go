echo "Clean..."
rm ./information.get_schedule_v2
echo "Build..."
go build -o information.get_schedule_v2 main.go 
echo "Build Done..."
echo "Run..."
./information.get_schedule_v2 > information.get_schedule_v2-rs.json
echo "Done."