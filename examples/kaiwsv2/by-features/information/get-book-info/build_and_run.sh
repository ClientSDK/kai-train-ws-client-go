echo "Clean..."
rm ./information.get_book_info
echo "Build..."
go build -o information.get_book_info main.go 
echo "Build Done..."
echo "Run..."
./information.get_book_info > information.get_book_info-rs.json
echo "Done."