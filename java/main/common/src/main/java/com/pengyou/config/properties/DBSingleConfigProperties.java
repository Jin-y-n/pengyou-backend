package com.pengyou.config.properties;



/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/15/24
    @Description: 

*/

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@Data
@Configuration
@ConfigurationProperties(prefix = "spring.datasource")
public class DBSingleConfigProperties {

    private String url;
    private String username;
    private String password;
    private String driverClassName;
}
