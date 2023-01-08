rm -rf ../build/linux/*

cd ..
go build -o ./build/linux/ControllerGo

cd ./scripts
chmod +x ./run.sh
./run.sh