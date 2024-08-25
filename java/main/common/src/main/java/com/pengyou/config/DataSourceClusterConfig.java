package com.pengyou.config;


/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/25/24
    @Description: 

*/

import com.pengyou.config.properties.DBSingleConfigProperties;
import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

import java.util.List;

@Data
@Configuration
@ConfigurationProperties(prefix = "pengyou.cluster-db")
public class DataSourceClusterConfig {
    List<DBSingleConfigProperties> dbs;
}
