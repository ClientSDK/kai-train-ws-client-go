echo "Clean..."
rm ./transaction.manual_seat
echo "Build..."
go build -o transaction.manual_seat main.go 
echo "Build Done..."
echo "Run..."
./transaction.manual_seat > transaction.manual_seat-rs.json
echo "Done."