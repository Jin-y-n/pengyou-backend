package com.pengyou.service.impl;

import com.pengyou.constant.AccountConstant;
import com.pengyou.constant.AuthorityConstant;
import com.pengyou.constant.RedisConstant;
import com.pengyou.exception.BaseException;
import com.pengyou.exception.InitFailedException;
import com.pengyou.exception.LoginFailedException;
import com.pengyou.exception.NoAuthorityException;
import com.pengyou.exception.common.InputInvalidException;
import com.pengyou.model.dto.admin.*;
import com.pengyou.model.entity.Admin;
import com.pengyou.model.entity.AdminTable;
import com.pengyou.service.AdminService;
import com.pengyou.util.security.AccountUtil;
import com.pengyou.util.verify.CaptchaGenerator;
import com.pengyou.util.verify.MailUtil;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.sql.JSqlClient;
import org.babyfish.jimmer.sql.ast.Predicate;
import org.babyfish.jimmer.sql.ast.mutation.SimpleSaveResult;
import org.babyfish.jimmer.sql.ast.tuple.Tuple3;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
@RequiredArgsConstructor
public class AdminImpl implements AdminService {
    private final JSqlClient sqlClient;
    private final RedisTemplate<String, String> template;
    private final MailUtil mailUtil;
    private final AdminTable adminTable = AdminTable.$;

    @Override
    public void register(AdminForRegister adminForRegister) {
        List<Tuple3<String, String, String>> execute = sqlClient
                .createQuery(adminTable)
                .where(
                        Predicate.or(
                                adminTable.email().eq(adminForRegister.getEmail()),
                                adminTable.phone().eq(adminForRegister.getPhone()),
                                adminTable.username().eq(adminForRegister.getUsername())
                        )
                ).select(
                        adminTable.email(),
                        adminTable.phone(),
                        adminTable.username()
                ).execute();

        if (!execute.isEmpty()) {
            execute.forEach(
                    tuple -> {
                        if (tuple.get_1().equals(adminForRegister.getEmail())) {
                            throw new InitFailedException(AccountConstant.ACCOUNT_EMAIL_EXISTS);
                        }

                        if (tuple.get_2().equals(adminForRegister.getPhone())) {
                            throw new InitFailedException(AccountConstant.ACCOUNT_PHONE_EXISTS);
                        }

                        if (tuple.get_3().equals(adminForRegister.getUsername())) {
                            throw new InitFailedException(AccountConstant.ACCOUNT_NAME_EXISTS);
                        }
                    }
            );
        }
        SimpleSaveResult<Admin> insert = sqlClient
                .insert(adminForRegister);

        sqlClient
                .createUpdate(adminTable)
                .set(adminTable.createdPerson(), insert.getModifiedEntity().id())
                .set(adminTable.modifiedPerson(), insert.getModifiedEntity().id())
                .where(adminTable.username().eq(insert.getModifiedEntity().username()))
                .execute();


    }

    @Override
    public void verify(AdminForVerify adminForVerify) {
        if (AccountUtil.checkEmail(adminForVerify.getEmail())) {
            String captcha = CaptchaGenerator.generateCaptcha(6);

            template.opsForValue().set(RedisConstant.ADMIN_CAPTCHA + adminForVerify.getEmail(), captcha);
            mailUtil.sendCaptcha(captcha, adminForVerify.getEmail());
            return;
        }

        if (AccountUtil.checkPhone(adminForVerify.getPhone())) {
            String captcha = CaptchaGenerator.generateCaptcha(6);

            template.opsForValue().set(RedisConstant.ADMIN_CAPTCHA + adminForVerify.getEmail(), captcha);
            // TODO: SEND SMS

            return;
        }

        throw new InputInvalidException();
    }

    @Override
    public void delete(AdminForDelete adminForDelete) {
        sqlClient
                .deleteByIds(Admin.class, adminForDelete.getIds());
    }

    @Override
    public void update(AdminForUpdate adminForUpdate) {
        if (adminForUpdate.getModifiedPerson() == adminForUpdate.getCreatedPerson()) {
            throw new NoAuthorityException(AuthorityConstant.NO_AUTHORIZATION);
        }

        sqlClient
                .update(adminForUpdate);
    }

    @Override
    public List<AdminForView> query(AdminForQuery adminForQuery) {

        List<AdminForView> execute = sqlClient
                .createQuery(adminTable)
                .where(adminForQuery)
                .select(
                        adminTable.fetch(AdminForView.class)
                )
                .execute();

        if (execute.isEmpty()) {
            throw new BaseException("Admin查询失败");
        }
        return execute;
    }

    @Override
    public AdminForLoginView login(AdminForLogin adminForLogin) {

        List<AdminForLoginView> execute = sqlClient
                .createQuery(adminTable)
                .where(adminForLogin)
                .select(
                        adminTable.fetch(AdminForLoginView.class)
                )
                .execute();
        Optional<AdminForLoginView> first = execute.stream().findFirst();

        if (first.isEmpty()) {
            throw new LoginFailedException("Admin登录失败");
        }

        return first.get();
    }

    @Override
    public void logout(AdminForLogout adminForLogout) {

    }
}
