package com.pengyou.config;



/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/15/24
    @Description: 

*/


import com.pengyou.util.RedisLock;
import lombok.RequiredArgsConstructor;
import org.redisson.api.RedissonClient;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
@RequiredArgsConstructor
public class RedisLockConfig {

    private final RedissonClient redisson;

    @Bean
    public RedisLock redisLock() {

        return new RedisLock(redisson, redisson.getLock(RedisLock.DEFAULT_LOCK_KEY));
    }
}
