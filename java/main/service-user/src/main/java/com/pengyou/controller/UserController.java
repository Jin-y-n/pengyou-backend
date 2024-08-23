package com.pengyou.controller;

/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/16/24
    @Description: 

*/

import com.pengyou.config.properties.JwtProperties;
import com.pengyou.constant.AccountConstant;
import com.pengyou.constant.JwtClaimsConstant;
import com.pengyou.constant.RedisConstant;
import com.pengyou.exception.common.CaptchaErrorException;
import com.pengyou.model.Result;
import com.pengyou.model.dto.admin.AdminForUpdate;
import com.pengyou.model.dto.user.*;
import com.pengyou.model.response.AdminLoginResult;
import com.pengyou.model.response.UserLoginResult;
import com.pengyou.service.UserService;
import com.pengyou.util.security.JwtUtil;
import com.pengyou.util.security.SHA256Encryption;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;

@Api
@RestController
@RequiredArgsConstructor
@RequestMapping("/user")
public class UserController {

    private final UserService userService;
    private final JwtProperties jwtProperties;
    private final RedisTemplate<String, String> template;

    @Api
    @PostMapping("/account/register")
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

        userService.register(user);

        return Result.success("success -- register");
    }

    @Api
    @PostMapping("/account/verify")
    public Result verify(
            @RequestBody UserForVerify user
    ) {

        userService.verify(user);

        return Result.success();
    }

    @Api
    @PostMapping("/login")
    public Result login(
            @RequestBody UserForLogin userForLogin
    ) {
        UserForLoginView user = this.userService.login(userForLogin);
        if (user != null) {
            HashMap<String, Object> map = new HashMap<>();
            map.put(JwtClaimsConstant.ID, user.getId());
            String jwt = JwtUtil.createJWT(this.jwtProperties.getSecretKey(), this.jwtProperties.getTtl(), map);

            return Result.success(new UserLoginResult(user.toEntity(), jwt));

        } else {
            return Result.error(AccountConstant.ACCOUNT_LOGIN_FAILURE);
        }
    }

    @Api
    @PostMapping("/update")
    public Result update(
            @RequestBody UserForUpdate user
    ) {

        userService.update(user);
        return Result.success("Profile update successful");
    }


}
