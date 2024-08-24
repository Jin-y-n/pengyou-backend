mkdir ./elasticsearch/node1/data -p

mkdir ./elasticsearch/node1/conf -p

chmod 777 ./elasticsearch -R

podman network create es-network

cat > "./elasticsearch/node1/conf/elasticsearch.yml" << EOF
cluster.name: "es-cluster"
network.host: 0.0.0.0
node.name: es-01
cluster.initial_master_nodes:
    - es-01
EOF

   cat > "./elasticsearch/node1/conf/log4j2.properties" << EOF
   # Default root logger level is set to WARN.
   # To enable additional logging, change the level below to DEBUG.
   #
   # For more information about configuring log4j2, see:
   # https://www.elastic.co/guide/en/elasticsearch/reference/current/logging.html
   #
   # Logging levels:
   #   OFF, FATAL, ERROR, WARN, INFO, DEBUG, TRACE, ALL
   #
   # Logging appenders:
   #   console - logs to the console
   #   file - logs to a file
   #   rolling_file - logs to a file and rolls over when it reaches a certain size
   #   daily_rolling_file - logs to a file and rolls over at a specific time each day
   #
   # Example configuration for logging to the console:
   #   path = stdout
   #   level = info
   #   append = true
   #
   # Example configuration for logging to a file:
   #   path = /var/log/elasticsearch/elasticsearch.log
   #   level = info
   #   append = true
   #
   # Example configuration for logging to a rolling file:
   #   path = /var/log/elasticsearch/elasticsearch.log
   #   level = info
   #   append = true
   #   max_size = 10MB
   #   max_files = 7
   #
   # Example configuration for logging to a daily rolling file:
   #   path = /var/log/elasticsearch/elasticsearch.log
   #   level = info
   #   append = true
   #   roll_at = 00:00
   #   roll_offset = -05:00
   #   keep_files = 7
   #
   # Example configuration for logging to both the console and a file:
   #   path = stdout
   #   level = info
   #   append = true
   #   path = /var/log/elasticsearch/elasticsearch.log
   #   level = info
   #   append = true
   #
   # Example configuration for logging to both the console and a rolling file:
   #   path = stdout
   #   level = info
   #   append = true
   #   path = /var/log/elasticsearch/elasticsearch.log
   #   level = info
   #   append = true
   #   max_size = 10MB
   #   max_files = 7
   #
   # Example configuration for logging to both the console and a daily rolling file:
   #   path = stdout
   #   level = info
   #   append = true
   #   path = /var/log/elasticsearch/elasticsearch.log
   #   level = info
   #   append = true
   #   roll_at = 00:00
   #   roll_offset = -05:00
   #   keep_files = 7
   #
   # Example configuration for logging to both a file and a rolling file:
   #   path = /var/log/elasticsearch/elasticsearch.log
   #   level = info
   #   append = true
   #   path = /var/log/elasticsearch/elasticsearch.log
   #   level = info
   #   append = true
   #   max_size = 10MB
   #   max_files = 7
   #
   # Example configuration for logging to both a file and a daily rolling file:
   #   path = /var/log/elasticsearch/elasticsearch.log
   #   level = info
   #   append = true
   #   path = /var/log/elasticsearch/elasticsearch.log
   #   level = info
   #   append = true
   #   roll_at = 00:00
   #   roll_offset = -05:00
   #   keep_files = 7
   #
   # Example configuration for logging to the console and a file:
   path = stdout
   level = info
   append = true
   path = /var/log/elasticsearch/elasticsearch.log
   level = info
   append = true
EOF


sudo podman run -d \
    -p 9200:9200 \
    -v ./elasticsearch/node1/data:/usr/share/elasticsearch/data \
    -v ./elasticsearch/node1/conf:/usr/share/elasticsearch/config \
    --name es-01 \
    --privileged=true \
    m.daocloud.io/docker.elastic.co/elasticsearch/elasticsearch:8.15.0

podman run -d \
    -p 9200:9200 \
    --name es-cluster-01 \
    --privileged=true \
    m.daocloud.io/docker.elastic.co/elasticsearch/elasticsearch:8.15.0

sleep 10

mkdir ./elasticsearch/node2/data -p
mkdir ./elasticsearch/node3/data -p

mkdir ./elasticsearch/node2/conf -p
mkdir ./elasticsearch/node3/conf -p

cat > "./elasticsearch/node2/conf/elasticsearch.yml" << EOF
cluster.name: "es-cluster"
network.host: 0.0.0.0
node.name: es-02
discovery.seed_hosts:
    - es-01
cluster.initial_master_nodes:
    - es-01
EOF

cat > "./elasticsearch/node3/conf/elasticsearch.yml" << EOF
cluster.name: "es-cluster"
network.host: 0.0.0.0
node.name: es-03
discovery.seed_hosts:
    - es-01
cluster.initial_master_nodes:
    - es-01
EOF
# sudo podman run -d \
#     -p 9201:9200 \
#     -v ./elasticsearch/node2/data:/usr/share/elasticsearch/data \
#     -v ./elasticsearch/node2/conf/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml \
#     --net es-network \
#     --name es-02 \
#     docker.elastic.co/elasticsearch/elasticsearch:8.15.0

# sudo podman run -d \
#     -p 9202:9200 \
#     -v ./elasticsearch/node3/data:/usr/share/elasticsearch/data \
#     -v ./elasticsearch/node3/conf/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml \
#     --net es-network \
#     --name es-03 \
#     docker.elastic.co/elasticsearch/elasticsearch:8.15.0




