package com.pengyou.service.impl;



/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/16/24
    @Description: 

*/

import com.pengyou.model.dto.user.UserForAdd;
import com.pengyou.service.UserService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

@Slf4j
@Service
public class UserServiceImpl implements UserService {

    @Override
    public void register(UserForAdd user) {

      log.info("user register: {}", user.getUsername());
    }
}
