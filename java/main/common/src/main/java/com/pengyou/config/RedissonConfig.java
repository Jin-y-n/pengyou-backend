package com.pengyou.config;



/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/15/24
    @Description: 

*/

import com.pengyou.config.properties.RedisLockConfigProperties;
import lombok.RequiredArgsConstructor;
import org.redisson.Redisson;
import org.redisson.api.RedissonClient;
import org.redisson.config.Config;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
@RequiredArgsConstructor
public class RedissonConfig {

    private final RedisLockConfigProperties redisLockConfigProperties;

    @Bean
    public RedissonClient redissonClient() {
        Config config = new Config();

        config.useSingleServer()
                .setAddress(redisLockConfigProperties.getUrl())
                .setPassword(redisLockConfigProperties.getPassword())
                .setDatabase(redisLockConfigProperties.getDatabase())
                .setConnectTimeout(redisLockConfigProperties.getConnectTimeOut())
                .setTimeout(redisLockConfigProperties.getTimeOut());

        return Redisson.create(config);
    }

}
