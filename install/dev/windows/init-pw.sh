# Create the directory 'pengyou'
New-Item -ItemType Directory -Path .\pengyou

# Change directory to 'pengyou'
Set-Location .\pengyou

# Create nested directories using New-Item
New-Item -ItemType Directory -Path .\redis\db1\conf -Force
New-Item -ItemType Directory -Path .\redis\db1\data -Force
New-Item -ItemType Directory -Path .\redis\db2\conf -Force
New-Item -ItemType Directory -Path .\redis\db2\data -Force
New-Item -ItemType Directory -Path .\mysql\db1\conf -Force
New-Item -ItemType Directory -Path .\mysql\db1\data -Force

# Write content to 'redis.conf' files
@"
port 6379
requirepass [this should be your password]
bind 0.0.0.0
protected-mode no
daemonize no
"@ | Set-Content -Path .\redis\db1\conf\redis.conf

@"
port 6379
requirepass [this should be your password]
bind 0.0.0.0
protected-mode no
daemonize no
"@ | Set-Content -Path .\redis\db2\conf\redis.conf

# Start container for Redis instances
docker run -d --name redis1 -p 12345:6379 --restart always --privileged -v "$(Get-Location).Path\redis\db1\conf:/usr/local/etc/redis" -v "$(Get-Location).Path\redis\db1\data:/data" redis redis-server /usr/local/etc/redis/redis.conf

docker run -d --name redis2 -p 12346:6379 --restart always --privileged -v "$(Get-Location).Path\redis\db2\conf:/usr/local/etc/redis" -v "$(Get-Location).Path\redis\db2\data:/data" redis redis-server /usr/local/etc/redis/redis.conf

# Start container for MySQL instance
docker run --name mysql1 -d --restart always -e MYSQL_ROOT_PASSWORD=[this should be your password] -p 3306:3306 -v "$(Get-Location).Path\mysql\db1\data:/var/lib/mysql" --privileged mysql