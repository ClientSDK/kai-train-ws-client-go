echo "Clean..."
rm ./data.get_pay_type
echo "Build..."
go build -o data.get_pay_type main.go 
echo "Build Done..."
echo "Run..."
./data.get_pay_type > data.get_pay_type-rs.json
echo "Done."