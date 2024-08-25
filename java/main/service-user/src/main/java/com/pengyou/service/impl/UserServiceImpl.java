package com.pengyou.service.impl;



/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/16/24
    @Description: 

*/

import com.pengyou.constant.AccountConstant;
import com.pengyou.constant.RedisConstant;
import com.pengyou.exception.BaseException;
import com.pengyou.exception.InitFailedException;
import com.pengyou.exception.LoginFailedException;
import com.pengyou.exception.common.InputInvalidException;
import com.pengyou.model.dto.profile.UserForUpdate;
import com.pengyou.model.dto.user.*;
import com.pengyou.model.entity.Admin;
import com.pengyou.model.entity.AdminTable;
import com.pengyou.model.entity.User;
import com.pengyou.model.entity.UserTable;
import com.pengyou.service.UserService;
import com.pengyou.util.UserContext;
import com.pengyou.util.security.AccountUtil;
import com.pengyou.util.verify.CaptchaGenerator;
import com.pengyou.util.verify.MailUtil;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.sql.JSqlClient;
import org.babyfish.jimmer.sql.ast.Predicate;
import org.babyfish.jimmer.sql.ast.mutation.SimpleSaveResult;
import org.babyfish.jimmer.sql.ast.tuple.Tuple2;
import org.babyfish.jimmer.sql.ast.tuple.Tuple3;
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
        List<Tuple3<String, String, String>> execute = jSqlClient
                .createQuery(userTable)
                .where(
                        Predicate.or(
                                userTable.email().eq(user.getEmail()),
                                userTable.phone().eq(user.getPhone()),
                                userTable.username().eq(user.getUsername())
                        )
                ).select(
                        userTable.email(),
                        userTable.phone(),
                        userTable.username()
                ).execute();

        if (!execute.isEmpty()) {
            execute.forEach(
                    tuple -> {
                        if (tuple.get_1().equals(user.getEmail())) {
                            throw new InitFailedException(AccountConstant.ACCOUNT_EMAIL_EXISTS);
                        }

                        if (tuple.get_2().equals(user.getPhone())) {
                            throw new InitFailedException(AccountConstant.ACCOUNT_PHONE_EXISTS);
                        }

                        if (tuple.get_3().equals(user.getUsername())) {
                            throw new InitFailedException(AccountConstant.ACCOUNT_NAME_EXISTS);
                        }
                    }
            );
        }
        SimpleSaveResult<User> insert = jSqlClient
                .insert(user);

        jSqlClient
                .createUpdate(userTable)
                .set(userTable.createdPerson(), insert.getModifiedEntity().id())
                .set(userTable.modifiedPerson(), insert.getModifiedEntity().id())
                .where(userTable.username().eq(insert.getModifiedEntity().username()))
                .execute();
    }

    @Override
    public String verify(UserForVerify user) {
        if (AccountUtil.checkEmail(user.getEmail())) {
            String captcha = CaptchaGenerator.generateCaptcha(6);

            template.opsForValue().set(RedisConstant.USER_CAPTCHA + user.getEmail(), captcha);
            mailUtil.sendCaptcha(captcha, user.getEmail());
            return captcha;
        }

        if (AccountUtil.checkPhone(user.getPhone())) {
            String captcha = CaptchaGenerator.generateCaptcha(6);

            template.opsForValue().set(RedisConstant.USER_CAPTCHA + user.getEmail(), captcha);
            // TODO: SEND SMS

            return captcha;
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
//        if (UserContext.getUserId().longValue() == (user.getId())) {
//            jSqlClient
//                    .update(user);.
//            return;
//        }
//
//        List<Tuple2<Long, Short>> select = jSqlClient
//                .createQuery(userTable)
//                .where(
//                        Predicate.or(
//                                userTable.id().eq(Long.valueOf(UserContext.getUserId())),
//                                userTable.id().eq(user.getId())
//                        )
//                )
//                .select(userTable.id(), userTable.role())
//                .execute();
//
//        select.forEach(
//                tuple -> {
//                    if (tuple.get_1().equals(Long.valueOf(UserContext.getUserId())) && tuple.get_2() == 2) {
//                        throw new BaseException(AccountConstant.INSUFFICIENT_AUTHORITY);
//                    } else if (tuple.get_1().equals(user.getId()) && tuple.get_2() == 1) {
//                        throw new BaseException(AccountConstant.INSUFFICIENT_AUTHORITY);
//                    }
//
//                }
//        );
//        jSqlClient
//                .createUpdate(userTable)
//                .set(userTable.modifiedByRoot(), (short) 1)
//                .where(userTable.id().eq(user.getId()))
//                .execute();
        jSqlClient
                .update(user);
    }

    @Override
    public void updateSensitive(UserForUpdateSensitive user) {

        jSqlClient
                .update(user);
    }
}
