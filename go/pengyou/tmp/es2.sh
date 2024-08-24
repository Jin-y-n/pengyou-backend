mkdir -p ./es/data/node1/data ./es/data/node1/config
mkdir -p ./es/data/node2/data ./es/data/node2/config
mkdir -p ./es/data/node3/data ./es/data/node3/config

# ./es/data/node1/config/elasticsearch.yml
cluster.name: my-elasticsearch-cluster
node.name: node1
path.data: ./es/data/node1/data
network.host: 0.0.0.0
discovery.seed_hosts: ["node1", "node2", "node3"]
cluster.initial_master_nodes: ["node1", "node2", "node3"]#

# ./es/data/node1/config/elasticsearch.yml
cluster.name: my-elasticsearch-cluster
node.name: node2
path.data: ./es/data/node2/data
network.host: 0.0.0.0
discovery.seed_hosts: ["node1", "node2", "node3"]
cluster.initial_master_nodes: ["node1", "node2", "node3"]#

# ./es/data/node1/config/elasticsearch.yml
cluster.name: my-elasticsearch-cluster
node.name: node3
path.data: ./es/data/node3/data
network.host: 0.0.0.0
discovery.seed_hosts: ["node1", "node2", "node3"]
cluster.initial_master_nodes: ["node1", "node2", "node3"]

podman run -d --name node1 -v ./es/data/node1/data:/usr/share/elasticsearch/data -v ./es/data/node1/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -p 19200:9200 -p 19300:9300 docker.elastic.co/elasticsearch/elasticsearch:8.15.0

docker run -d --name node2 -v ./es/data/node2/data:/usr/share/elasticsearch/data -v ./es/data/node2/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -p 19201:9200 -p 19301:9300 docker.elastic.co/elasticsearch/elasticsearch:8.15.0

docker run -d --name node3 -v ./es/data/node3/data:/usr/share/elasticsearch/data -v ./es/data/node3/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -p 19202:9200 -p 19302:9300 docker.elastic.co/elasticsearch/elasticsearch:8.15.0

docker run -d --name es -v ./elasticsearch/date:/usr/share/elasticsearch/data -p 9200:9200 -p 9300:9300 docker.elastic.co/elasticsearch/elasticsearch:8.15.0

podman run -d \
--name=es_node_1 \
--restart=always \
-p 19201:9200 \
-p 19301:9300 \
--privileged=true \
-v ./es/es_cluster/node_1/data:/usr/share/elasticsearch/data \
-v ./es/es_cluster/node_1/logs:/usr/share/elasticsearch/logs \
-v ./es/es_cluster/node_1/plugins:/usr/share/elasticsearch/plugins \
-e "cluster.name=my-cluster" \
-e "node.name=node-1" \
-e "node.master=true" \
-e "node.data=true" \
-e "network.host=0.0.0.0" \
-e "transport.tcp.port=9300" \
-e "http.port=9200" \
-e "cluster.initial_master_nodes=node-1" \
-e "discovery.seed_hosts=127.0.0.1" \
-e "gateway.auto_import_dangling_indices=true" \
-e "http.cors.enabled=true" \
-e "http.cors.allow-origin=*" \
-e "ES_JAVA_OPTS=-Xms512m -Xmx512m" \
-e "TAKE_FILE_OWNERSHIP=true" \
elasticsearch:8.15.0

podman run -d \
--name=es_node_2 \
--restart=always \
-p 19202:9200 \
-p 19302:9300 \
--privileged=true \
-v ./es/es_cluster/node_2/data:/usr/share/elasticsearch/data \
-v ./es/es_cluster/node_2/logs:/usr/share/elasticsearch/logs \
-v ./es/es_cluster/node_2/plugins:/usr/share/elasticsearch/plugins \
-e "cluster.name=my-cluster" \
-e "node.name=node-2" \
-e "node.master=true" \
-e "node.data=true" \
-e "network.host=0.0.0.0" \
-e "transport.tcp.port=9300" \
-e "http.port=9200" \
-e "discovery.seed_hosts=127.0.0.1" \
-e "gateway.auto_import_dangling_indices=true" \
-e "http.cors.enabled=true" \
-e "http.cors.allow-origin=*" \
-e "ES_JAVA_OPTS=-Xms512m -Xmx512m" \
-e "TAKE_FILE_OWNERSHIP=true" \
elasticsearch:8.15.0


