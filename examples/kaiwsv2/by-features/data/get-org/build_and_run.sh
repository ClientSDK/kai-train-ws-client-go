echo "Clean..."
rm ./data.get_org
echo "Build..."
go build -o data.get_org main.go 
echo "Build Done..."
echo "Run..."
./data.get_org > data.get_org-rs.json
echo "Done."