docker stop mysql1
docker stop redis1
docker stop redis2
docker stop es01
docker stop es02
docker stop es03

docker rm mysql1
docker rm redis1
docker rm redis2
docker rm es01
docker rm es02
docker rm es03

docker volume rm data01
docker volume rm data02
docker volume rm data03

docker volume rm linux_data01
docker volume rm linux_data02
docker volume rm linux_data03

docker network rm elastic
docker network rm linux_elastic



Remove-Item -Path ./pengyou -Recurse -Force