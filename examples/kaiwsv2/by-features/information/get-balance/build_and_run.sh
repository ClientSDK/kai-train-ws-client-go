echo "Clean..."
rm ./information.get_balance
echo "Build..."
go build -o information.get_balance main.go 
echo "Build Done..."
echo "Run..."
./information.get_balance > information.get_balance-rs.json
echo "Done."