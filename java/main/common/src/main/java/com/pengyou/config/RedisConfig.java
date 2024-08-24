package com.pengyou.config;



/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/15/24
    @Description: 

*/
import lombok.RequiredArgsConstructor;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.GenericJackson2JsonRedisSerializer;

@Configuration
@RequiredArgsConstructor
public class RedisConfig {

    @Bean
    public RedisTemplate<String, Object>  redisTemplate(RedisConnectionFactory redisConnectionFactory) {
        RedisTemplate<String, Object>  template = new RedisTemplate<>();
        template.setConnectionFactory(redisConnectionFactory);
        //对象的序列化，GenericJackson2JsonRedisSerializer实现了RedisSerializer接口
        GenericJackson2JsonRedisSerializer serializer = new GenericJackson2JsonRedisSerializer();
        template.setDefaultSerializer(serializer);
        return template;
    }


}
