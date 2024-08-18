echo -e "stoping redis...... "
echo -e "stoping redis-cluster1......"
docker stop redis-cluster9101
echo -e "stopped successfully\n stoping redis-cluster2......"
docker stop redis-cluster9102
echo -e "stopped successfully\n stoping redis-cluster3......"
docker stop redis-cluster9103
echo -e "stopped successfully\n stoping redis-cluster4......"
docker stop redis-cluster9104
echo -e "stopped successfully\n stoping redis-cluster5......"
docker stop redis-cluster9105
echo -e "stopped successfully\n stoping redis-cluster6......"
docker stop redis-cluster9106

echo -e "stopped successfully\n stoping redis-for-lock"
docker stop redis-for-lock
echo -e "stopped successfully\n"

echo -e "remove redis container......"
echo -e "removing redis-cluster1......"
docker rm redis-cluster9101
echo -e "removed successfully\n removing redis-cluster2......"
docker rm redis-cluster9102
echo -e "removed successfully\n removing redis-cluster3......"
docker rm redis-cluster9103
echo -e "removed successfully\n removing redis-cluster4......"
docker rm redis-cluster9104
echo -e "removed successfully\n removing redis-cluster5......"
docker rm redis-cluster9105
echo -e "removed successfully\n removing redis-cluster6......"
docker rm redis-cluster9106

echo -e "removed successfully\n removing redis-for-lock......"
docker rm redis-for-lock
echo -e "removed successfully"


echo -e "stoping mysql...... "
echo -e "stopped successfully\n stoping mysql-node2......"
docker stop mysql-node2
echo -e "stopped successfully\n stoping mysql-node3......"
docker stop mysql-node3
echo -e "stopped successfully\n stoping mysql-node4......"
docker stop mysql-node4
echo -e "stopped successfully\n stoping mysql-node5......"
docker stop mysql-node5
echo -e "stopped successfully\n stoping mysql-node1......"
docker stop mysql-node1
echo -e "stopped successfully\n"

echo -e "remove mysql container......"
echo -e "removing mysql-node1......"
docker rm mysql-node1
echo -e "removed successfully\n removing mysql-node2......"
docker rm mysql-node2
echo -e "removed successfully\n removing mysql-node3......"
docker rm mysql-node3
echo -e "removed successfully\n removing mysql-node4......"
docker rm mysql-node4
echo -e "removed successfully\n removing mysql-node5......"
docker rm mysql-node5
echo -e "removed successfully"


echo -e "stoping grafana container......"
docker stop grafana_pengyou

echo -e "removing grafana container......"
docker rm grafana_pengyou


echo -e "stoping prometheus container......"
docker stop prometheus_pengyou

echo -e "removing prometheus container......"
docker rm prometheus_pengyou


echo -e "removing data"
rm -Rf ./grafana
rm -Rf ./main
rm -Rf ./mysql
rm -Rf ./nacos
rm -Rf ./prometheus
rm -Rf ./redis
rm -Rf ./tmp

echo -e "data have removed successfully"

echo -e "removing networks......"
docker network rm redis-cluster-network
docker network rm mysql-cluster-network
echo -e "networks have removed successfully"

