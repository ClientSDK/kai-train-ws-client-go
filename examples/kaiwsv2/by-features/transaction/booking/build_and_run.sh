echo "Clean..."
rm ./transaction.booking
echo "Build..."
go build -o transaction.booking main.go 
echo "Build Done..."
echo "Run..."
./transaction.booking > transaction.booking-rs.json
echo "Done."