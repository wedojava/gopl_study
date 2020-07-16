go build -o clock.out gopl.io/ch8/exercise/ex8.01/ray-g/clock
go build -o clockwall.out gopl.io/ch8/exercise/ex8.01/ray-g/clockwall
TZ=US/Eastern    ./clock.out -port 8010 &
TZ=Asia/Tokyo    ./clock.out -port 8020 &
TZ=Europe/London ./clock.out -port 8030 &
./clockwall.out NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
killall clock.*
