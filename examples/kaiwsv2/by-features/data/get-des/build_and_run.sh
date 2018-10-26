echo "Clean..."
rm ./data.get_des
echo "Build..."
go build -o data.get_des main.go 
echo "Build Done..."
echo "Run..."
./data.get_des > data.get_des.json
echo "Done."