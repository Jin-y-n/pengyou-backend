package com.pengyou.service;



/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/16/24
    @Description: 

*/

import com.pengyou.model.dto.user.*;

public interface UserService {
    void register(UserForAdd user);

    void verify(UserForVerify user);

    void update(UserForUpdate user);

    UserForLoginView login(UserForLogin userForLogin);
}
