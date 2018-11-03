echo "Clean..."
rm ./information.get_agent_balance
echo "Build..."
go build -o information.get_agent_balance main.go 
echo "Build Done..."
echo "Run..."
./information.get_agent_balance > information.get_agent_balance-rs.json
echo "Done."