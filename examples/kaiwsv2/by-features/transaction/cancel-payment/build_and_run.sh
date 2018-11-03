echo "Clean..."
rm ./transaction.cancel_payment
echo "Build..."
go build -o transaction.cancel_payment main.go 
echo "Build Done..."
echo "Run..."
./transaction.cancel_payment > transaction.cancel_payment-rs.json
echo "Done."