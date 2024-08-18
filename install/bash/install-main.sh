#!/bin/bash


mkdir ./tmp
mkdir ./main
mkdir ./mysql
mkdir ./mysql/node1/data -p
mkdir ./mysql/node2/data -p
mkdir ./mysql/node3/data -p
mkdir ./mysql/node4/data -p
mkdir ./mysql/node5/data -p
mkdir ./redis
mkdir ./redis/for-lock/data -p
mkdir ./redis/for-lock/conf
mkdir ./nacos
mkdir ./grafana
mkdir ./prometheus

chmod 777 ./mysql -R
chmod 777 ./redis -R

############### redis cluster create ###############

echo "creating docker network......"
docker network create --subnet=172.191.111.0/24 redis-cluster-network 
docker network create --subnet=172.191.191.0/24 mysql-cluster-network


echo "creating redis configures......"
# create configuration
for port in $(seq 9101 9106); do
    node_path="./redis/node$port/conf"
    mkdir -p "$node_path"
    mkdir -p "$node_path/../data"
    conf_file_path="$node_path/redis.conf"

    cat > "$conf_file_path" <<EOF
port $port
requirepass [there should be your password]
bind 0.0.0.0
protected-mode no
daemonize no
appendonly yes
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip [this should be your ip address]
cluster-announce-port $port
cluster-announce-bus-port 1$port
EOF
done

cat > "./redis/for-lock/conf/redis.conf" << EOF
requirepass [there should be your password]
bind 0.0.0.0
EOF

# start container
for port in $(seq 9101 9106); do
    echo "runing redis cluster: redis-cluster$port"
    docker run -it -d \
               -p "$port:$port" -p "1$port:1$port" \
               --privileged \
               -v "$(pwd)/redis/node$port/conf:/usr/local/etc/redis" \
               -v "$(pwd)/redis/node$port/data:/data" \
               --restart always \
               --name "redis-cluster$port" \
               --net redis-cluster-network \
               redis redis-server /usr/local/etc/redis/redis.conf
done

# wait for begin
for port in $(seq 9101 9106); do
    while true; do
        container_status=$(docker inspect --format='{{.State.Status}}' "redis-cluster$port")
        if [ "$container_status" == "running" ]; then
            break
        fi
        sleep 0.5
    done
done

# get ip address
nodes_info=()
for port in $(seq 9101 9106); do
    inspect_result=$(docker inspect "redis-cluster$port")
    container_ip=$(echo "$inspect_result" | jq -r '.[0].NetworkSettings.Networks["redis-cluster-network"].IPAddress')
    nodes_info+=("$container_ip:$port")
done

# create cluster
cluster_nodes=$(IFS=" "; echo "${nodes_info[*]}")
create_cluster_cmd="redis-cli --cluster create -a [there should be your password] [this should be your ip address]:9101 [this should be your ip address]:9102 [this should be your ip address]:9103 [this should be your ip address]:9104 [this should be your ip address]:9105 [this should be your ip address]:9106 --cluster-replicas 1"

echo "$create_cluster_cmd"

# execute create cluster
docker exec -it redis-cluster9101 /bin/bash -c "$create_cluster_cmd"


############### grafana ###############
touch ./grafana/grafana.ini

############### prometheus

cat > "./prometheus/prometheus.yml" <<EOF
global:
  scrape_interval: 10s

scrape_configs:
- job_name: pengyou-user
  static_configs:
  - targets: ['service:10000']

- job_name: pengyou-admin
  static_configs:
  - targets: ['service:11000']

EOF

############### MySQL cluster, granfana, prometheus create ###############


docker run --name mysql-node1 -d \
           --restart always \
           -e MYSQL_ROOT_PASSWORD=[there should be your password] \
           -e CLUSTER_NAME=PXC \
           -e XTRABACKUP_PASSWORD=[there should be your password_XtraBackup] \
           -p 19601:3306 \
           -v $(pwd)/mysql/node1/data:/var/lib/mysql \
           --network mysql-cluster-network \
           --ip 172.191.191.2 \
           --privileged \
           pxc


docker run --name mysql-node2 -d \
           --restart always \
           -e MYSQL_ROOT_PASSWORD=[there should be your password] \
           -e CLUSTER_NAME=PXC \
           -e XTRABACKUP_PASSWORD=[there should be your password_XtraBackup] \
           -e CLUSTER_JOIN=mysql-node1 \
           -p 19602:3306 \
           -v $(pwd)/mysql/node2/data:/var/lib/mysql \
           --network mysql-cluster-network \
           --ip 172.191.191.3 \
           --privileged \
           pxc


docker run --name mysql-node3 -d \
           --restart always \
           -e MYSQL_ROOT_PASSWORD=[there should be your password] \
           -e CLUSTER_NAME=PXC \
           -e XTRABACKUP_PASSWORD=[there should be your password_XtraBackup] \
           -e CLUSTER_JOIN=mysql-node1 \
           -p 19603:3306 \
           -v $(pwd)/mysql/node3/data:/var/lib/mysql \
           --network mysql-cluster-network \
           --ip 172.191.191.4 \
           --privileged \
           pxc
           
docker run --name mysql-node4 -d \
           --restart always \
           -e MYSQL_ROOT_PASSWORD=[there should be your password] \
           -e CLUSTER_NAME=PXC \
           -e XTRABACKUP_PASSWORD=[there should be your password_XtraBackup] \
           -e CLUSTER_JOIN=mysql-node2 \
           -p 19604:3306 \
           -v $(pwd)/mysql/node4/data:/var/lib/mysql \
           --network mysql-cluster-network \
           --ip 172.191.191.5 \
           --privileged \
           pxc

docker run --name mysql-node5 -d \
           --restart always \
           -e MYSQL_ROOT_PASSWORD=[there should be your password] \
           -e CLUSTER_NAME=PXC \
           -e XTRABACKUP_PASSWORD=[there should be your password_XtraBackup] \
           -e CLUSTER_JOIN=mysql-node2 \
           -p 19605:3306 \
           -v $(pwd)/mysql/node5/data:/var/lib/mysql \
           --network mysql-cluster-network \
           --ip 172.191.191.6 \
           --privileged \
           pxc

echo "please wait a moment (about 10 seconds) ......"
sleep 10

file_path="./mysql/node1/data/grastate.dat"

if [ ! -f "$file_path" ]; then
  echo "Error: File '$file_path' does not exist."
  exit 1
fi

sed -i.bak 's/safe_to_bootstrap: 0/safe_to_bootstrap: 1/g' "$file_path"

if [ $? -eq 0 ]; then
  echo "The string 'safe_to_bootstrap: 0' has been successfully replaced with 'safe_to_bootstrap: 1'."
else
  echo "Failed to replace the string."
fi

echo "starting mysql cluster node1, please wait"
docker restart mysql-node1
sleep 30
echo "starting mysql cluster node2"
docker restart mysql-node2
echo "starting mysql cluster node3"
docker restart mysql-node3
echo "starting mysql cluster node4"
docker restart mysql-node4
echo "starting mysql cluster node5"
docker restart mysql-node5

docker-compose up -d
