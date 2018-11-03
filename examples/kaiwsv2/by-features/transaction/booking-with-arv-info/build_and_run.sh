echo "Clean..."
rm ./transaction.booking_with_arv_info
echo "Build..."
go build -o transaction.booking_with_arv_info main.go 
echo "Build Done..."
echo "Run..."
./transaction.booking_with_arv_info > transaction.booking_with_arv_info-rs.json
echo "Done."