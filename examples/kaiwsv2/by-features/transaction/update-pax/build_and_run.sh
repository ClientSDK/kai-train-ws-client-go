echo "Clean..."
rm ./transaction.update_pax
echo "Build..."
go build -o transaction.update_pax main.go 
echo "Build Done..."
echo "Run..."
./transaction.update_pax > transaction.update_pax-rs.json
echo "Done."