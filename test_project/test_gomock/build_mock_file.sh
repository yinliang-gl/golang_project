
#!  /bin/bash
cd "$(dirname "$0")"
export dir_name=$(pwd)

rm -rf mock/* && mockgen -source=./person/male.go -destination=./mock/male_mock.go -package=mock
