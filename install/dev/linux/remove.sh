podman stop mysql1
podman stop redis1
podman stop redis2
podman stop es01
podman stop es02
podman stop es03

podman rm mysql1
podman rm redis1
podman rm redis2
podman rm es01
podman rm es02
podman rm es03

podman volume rm data01
podman volume rm data02
podman volume rm data03

podman volume rm linux_data01
podman volume rm linux_data02
podman volume rm linux_data03

podman network rm elastic
podman network rm linux_elastic

sudo rm -Rf pengyou

