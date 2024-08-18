package com.pengyou.controller;

/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/16/24
    @Description: 

*/

import com.pengyou.constant.RedisConstant;
import com.pengyou.exception.common.CaptchaErrorException;
import com.pengyou.model.Result;
import com.pengyou.model.dto.user.UserForAdd;
import com.pengyou.model.dto.user.UserForVerify;
import com.pengyou.service.UserService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.web.bind.annotation.*;

@Api
@RestController
@RequiredArgsConstructor
@RequestMapping("/user")
public class UserController {

    private final UserService service;

    private final RedisTemplate<String, String> template;

    @Api
    @PostMapping("/register")
    public Result register(
            @RequestBody UserForAdd user
            ) {

        String captcha = user.getCaptcha();

        if (captcha.equals(template.opsForValue().get(RedisConstant.USER_CAPTCHA+user.getEmail()))) {
            user.setPhone(null);
        } else if (captcha.equals(template.opsForValue().get(RedisConstant.USER_CAPTCHA+user.getPhone()))) {
            user.setEmail(null);
        } else {
            throw new CaptchaErrorException();
        }

        service.register(user);

        return Result.success("success -- register");
    }

    @Api
    @PostMapping("/verify")
    public Result verify(
            @RequestBody UserForVerify user
    ) {

        service.verify(user);

        return Result.success();
    }

}
