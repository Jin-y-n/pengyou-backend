docker stop mysql1
docker stop redis1
docker stop redis2

docker rm mysql1
docker rm redis1
docker rm redis2

Remove-Item -Path ./pengyou -Recurse -Force