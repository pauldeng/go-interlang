#!bin/sh
CURPATH=$(pwd)
cd greetings
rm -Rf build
mkdir build
cd build
cmake -G 'Unix Makefiles' -DCMAKE_BUILD_TYPE=Debug ..
cmake --build . --config Release -- -j3

cd $CURPATH
go build -o static_c_lib