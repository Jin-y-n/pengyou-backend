package com.pengyou.util;

/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/15/24
    @Description: 

*/

import lombok.RequiredArgsConstructor;
import org.redisson.api.RLock;
import org.redisson.api.RedissonClient;

@RequiredArgsConstructor
public class RedisLock {

    private final RedissonClient client;

    public static final String DEFAULT_LOCK_KEY = "lock-default-lock-key";

    private final RLock DEFAULT_LOCK;

    public void lock() {
        DEFAULT_LOCK.lock();
    }

    public void unlock() {
        DEFAULT_LOCK.unlock();
    }

    public void lock(String key) {
        client.getLock(key).lock();
    }

    public void unlock(String key) {
        client.getLock(key).unlock();
    }

}
