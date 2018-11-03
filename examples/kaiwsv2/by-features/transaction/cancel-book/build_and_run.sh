echo "Clean..."
rm ./transaction.cancel_book
echo "Build..."
go build -o transaction.cancel_book main.go 
echo "Build Done..."
echo "Run..."
./transaction.cancel_book > transaction.cancel_book-rs.json
echo "Done."