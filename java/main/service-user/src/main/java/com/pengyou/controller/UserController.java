package com.pengyou.controller;



/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/16/24
    @Description: 

*/

import com.pengyou.model.Result;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@Api
@RestController
@RequestMapping("/user")
public class UserController {

    @Api
    @GetMapping("/test")
    public Result test() {
        return Result.success("success");
    }

    @Api
    @GetMapping("/test2")
    public Result test2() {
        return Result.success("success -- 2");
    }

}
