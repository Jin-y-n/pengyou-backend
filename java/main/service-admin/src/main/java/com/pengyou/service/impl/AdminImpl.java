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
import com.pengyou.model.entity.User;
import com.pengyou.service.AdminService;
import com.pengyou.util.UserContext;
import com.pengyou.util.security.AccountUtil;
import com.pengyou.util.verify.CaptchaGenerator;
import com.pengyou.util.verify.MailUtil;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.Page;
import org.babyfish.jimmer.sql.JSqlClient;
import org.babyfish.jimmer.sql.ast.Predicate;
import org.babyfish.jimmer.sql.ast.mutation.SimpleSaveResult;
import org.babyfish.jimmer.sql.ast.query.ConfigurableRootQuery;
import org.babyfish.jimmer.sql.ast.tuple.Tuple2;
import org.babyfish.jimmer.sql.ast.tuple.Tuple3;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
@Slf4j
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
    public String verify(AdminForVerify adminForVerify) {
        if (AccountUtil.checkEmail(adminForVerify.getEmail())) {
            String captcha = CaptchaGenerator.generateCaptcha(6);

            template.opsForValue().set(RedisConstant.ADMIN_CAPTCHA + adminForVerify.getEmail(), captcha);
            mailUtil.sendCaptcha(captcha, adminForVerify.getEmail());
            return captcha;
        }

        if (AccountUtil.checkPhone(adminForVerify.getPhone())) {
            String captcha = CaptchaGenerator.generateCaptcha(6);

            template.opsForValue().set(RedisConstant.ADMIN_CAPTCHA + adminForVerify.getEmail(), captcha);
            // TODO: SEND SMS

            return captcha;
        }

        throw new InputInvalidException();
    }

    @Override
    public void delete(AdminForDelete adminForDelete) {
        List<Tuple2<Long, Short>> execute = sqlClient
                .createQuery(adminTable)
                .where(
                        Predicate.or(
                                adminTable.id().eq(Long.valueOf(UserContext.getUserId())),
                                adminTable.id().in(adminForDelete.getIds())
                        )
                )
                .select(adminTable.id(), adminTable.role())
                .execute();
        execute.forEach(
                tuple -> {
                    if (tuple.get_1().equals(Long.valueOf(UserContext.getUserId())) && tuple.get_2() == 2) {
                        throw new NoAuthorityException(AuthorityConstant.NO_AUTHORIZATION);
                    }

                    if (adminForDelete.getIds().contains(tuple.get_1()) && tuple.get_2() == 1) {
                        throw new NoAuthorityException(AuthorityConstant.NO_AUTHORIZATION);
                    }
                }
        );
        sqlClient
                .createUpdate(adminTable)
                .set(adminTable.modifiedPerson(), Long.valueOf(UserContext.getUserId()))
                .set(adminTable.modifiedByRoot(), (short) 1)
                .where(adminTable.id().in(adminForDelete.getIds()))
                .execute();
        sqlClient
                .deleteByIds(Admin.class, adminForDelete.getIds());
    }

    @Override
    public void update(AdminForUpdate adminForUpdate) {
        if (UserContext.getUserId().longValue() == (adminForUpdate.getId())) {
            sqlClient
                    .update(adminForUpdate);
            return;
        }

        List<Tuple2<Long, Short>> select = sqlClient
                .createQuery(adminTable)
                .where(
                        Predicate.or(
                                adminTable.id().eq(Long.valueOf(UserContext.getUserId())),
                                adminTable.id().eq(adminForUpdate.getId())
                        )
                )
                .select(adminTable.id(), adminTable.role())
                .execute();

        select.forEach(
                tuple -> {
                    if (tuple.get_1().equals(Long.valueOf(UserContext.getUserId())) && tuple.get_2() == 2) {
                        throw new BaseException(AccountConstant.INSUFFICIENT_AUTHORITY);
                    } else if (tuple.get_1().equals(adminForUpdate.getId()) && tuple.get_2() == 1) {
                        throw new BaseException(AccountConstant.INSUFFICIENT_AUTHORITY);
                    }

                }
        );
        sqlClient
                .createUpdate(adminTable)
                .set(adminTable.modifiedByRoot(), (short) 1)
                .where(adminTable.id().eq(adminForUpdate.getId()))
                .execute();
        sqlClient
                .update(adminForUpdate);

    }

    @Override
    public Page<AdminForView> query(AdminForQuery adminForQuery) {

        Page<AdminForView> page = sqlClient
                .createQuery(adminTable)
                .where(adminForQuery)
                .select(
                        adminTable.fetch(AdminForView.class)
                )
                .fetchPage(adminForQuery.getPageIndex(), adminForQuery.getPageSize());

        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Admin查询失败");
        }
        return page;
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
        first.get().setPassword(AccountConstant.ACCOUNT_PASSWORD_ENCODE);
        return first.get();
    }

    @Override
    public void logout(AdminForLogout adminForLogout) {
        if (!(UserContext.getUserId() == adminForLogout.getId())) {
            throw new BaseException("登出失败");
        }
    }
}
