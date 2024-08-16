package com.pengyou.controller;



/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/16/24
    @Description: 

*/

import com.pengyou.model.Result;
import com.pengyou.model.dto.user.UserForAdd;
import com.pengyou.service.UserService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.*;

@Api
@RestController
@RequiredArgsConstructor
@RequestMapping("/user")
public class UserController {

    private final UserService service;

    @Api
    @PostMapping("/register")
    public Result register(
            @RequestBody UserForAdd user
            ) {

        service.register(user);

        return Result.success("success -- register");
    }

}
