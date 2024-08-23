package com.pengyou.service.impl;



/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/16/24
    @Description: 

*/

import com.pengyou.constant.RedisConstant;
import com.pengyou.exception.LoginFailedException;
import com.pengyou.exception.common.InputInvalidException;
import com.pengyou.model.dto.user.*;
import com.pengyou.model.entity.AdminTable;
import com.pengyou.model.entity.UserTable;
import com.pengyou.service.UserService;
import com.pengyou.util.security.AccountUtil;
import com.pengyou.util.verify.CaptchaGenerator;
import com.pengyou.util.verify.MailUtil;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Slf4j
@Service
@RequiredArgsConstructor
public class UserServiceImpl implements UserService {

    private final JSqlClient jSqlClient;
    private final RedisTemplate<String, String> template;
    private final MailUtil mailUtil;
    private final UserTable userTable = UserTable.$;

    @Override
    public void register(UserForAdd user) {
        jSqlClient.save(user);

        log.info("user register: {}", user.getUsername());
    }

    @Override
    public void verify(UserForVerify user) {
        if (AccountUtil.checkEmail(user.getEmail())) {
            String captcha = CaptchaGenerator.generateCaptcha(6);

            template.opsForValue().set(RedisConstant.USER_CAPTCHA+user.getEmail(), captcha);
            mailUtil.sendCaptcha(captcha, user.getEmail());
            return;
        }

        if (AccountUtil.checkPhone(user.getPhone())) {
            String captcha = CaptchaGenerator.generateCaptcha(6);

            template.opsForValue().set(RedisConstant.USER_CAPTCHA+user.getEmail(), captcha);
            // TODO: SEND SMS

            return;
        }

        throw new InputInvalidException();
    }

    @Override
    public UserForLoginView login(UserForLogin userForLogin) {

        List<UserForLoginView> execute = jSqlClient
                .createQuery(userTable)
                .where(userForLogin)
                .select(
                        userTable.fetch(UserForLoginView.class)
                )
                .execute();
        Optional<UserForLoginView> first = execute.stream().findFirst();

        if (first.isEmpty()) {
            throw new LoginFailedException("");
        }

        return first.get();
    }

    @Override
    public void update(UserForUpdate user) {
        jSqlClient
                .update(user);
    }
}
