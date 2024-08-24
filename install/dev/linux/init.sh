mkdir pengyou


# shellcheck disable=SC2164
cd pengyou

mkdir redis/db1/conf -p
mkdir redis/db1/data -p
mkdir redis/db2/conf -p
mkdir redis/db2/data -p

mkdir mysql/db1/conf -p
mkdir mysql/db1/data -p

mkdir es/es01/data -p
mkdir es/es02/data -p
mkdir es/es03/data -p


chmod 777 -R ./
chown -R 1000:1000 ./

    cat > "redis/db1/conf/redis.conf" <<EOF
port 6379
requirepass 12345678
bind 0.0.0.0
protected-mode no
daemonize no
EOF

    cat > "redis/db2/conf/redis.conf" <<EOF
port 6379
requirepass 12345678
bind 0.0.0.0
protected-mode no
daemonize no
EOF


# shellcheck disable=SC2046
podman run -d --name redis1 -p 6379:6379 --restart always --privileged -v $(pwd)/redis/db1/conf:/usr/local/etc/redis -v $(pwd)/redis/db1/data:/data redis redis-server /usr/local/etc/redis/redis.conf

# shellcheck disable=SC2046
podman run -d --name redis2 -p 6380:6379 --restart always --privileged -v $(pwd)/redis/db2/conf:/usr/local/etc/redis -v $(pwd)/redis/db2/data:/data redis redis-server /usr/local/etc/redis/redis.conf

# shellcheck disable=SC2046
podman run --name mysql1 -d --restart always -e MYSQL_ROOT_PASSWORD=12345678 -p 3306:3306 -v $(pwd)/mysql/db1/data:/var/lib/mysql --privileged mysql

cd .. 

podman-compose up -d
