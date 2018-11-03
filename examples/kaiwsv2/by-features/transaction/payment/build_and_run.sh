echo "Clean..."
rm ./transaction.payment
echo "Build..."
go build -o transaction.payment main.go 
echo "Build Done..."
echo "Run..."
./transaction.payment > transaction.payment-rs.json
echo "Done."