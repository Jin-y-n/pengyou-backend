mkdir pengyou

chmod 777 -R ./pengyou
cd pengyou

mkdir redis/db1/conf -p
mkdir redis/db1/data -p
mkdir redis/db2/conf -p
mkdir redis/db2/data -p

mkdir mysql/db1/conf -p
mkdir mysql/db1/data -p


    cat > "redis/db1/conf/redis.conf" <<EOF
port 6379
requirepass [this should be your password]
bind 0.0.0.0
protected-mode no
daemonize no
EOF

    cat > "redis/db2/conf/redis.conf" <<EOF
port 6379
requirepass [this should be your password]
bind 0.0.0.0
protected-mode no
daemonize no
EOF


podman run -d --name redis1 -p 12345:6379 --restart always --privileged -v $(pwd)/redis/db1/conf:/usr/local/etc/redis -v $(pwd)/redis/db1/data:/data redis redis-server /usr/local/etc/redis/redis.conf

podman run -d --name redis2 -p 12346:6379 --restart always --privileged -v $(pwd)/redis/db2/conf:/usr/local/etc/redis -v $(pwd)/redis/db2/data:/data redis redis-server /usr/local/etc/redis/redis.conf


podman run --name mysql1 -d --restart always -e MYSQL_ROOT_PASSWORD=[this should be your password] -p 3306:3306 -v $(pwd)/mysql/db1/data:/var/lib/mysql --privileged mysql

