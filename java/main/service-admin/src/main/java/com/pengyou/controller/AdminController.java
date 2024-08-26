package com.pengyou.controller;

import com.pengyou.config.properties.JwtProperties;
import com.pengyou.constant.AccountConstant;
import com.pengyou.constant.JwtClaimsConstant;
import com.pengyou.constant.RedisConstant;
import com.pengyou.constant.VerifyConstant;
import com.pengyou.exception.common.CaptchaErrorException;
import com.pengyou.model.Result;
import com.pengyou.model.dto.admin.*;
import com.pengyou.model.response.AdminLoginResult;
import com.pengyou.service.AdminService;
import com.pengyou.util.UserContext;
import com.pengyou.util.security.JwtUtil;
import com.pengyou.util.security.SHA256Encryption;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.client.meta.Api;
import org.babyfish.jimmer.sql.ast.tuple.Tuple2;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;
import java.util.Date;
import java.util.HashMap;

@Api
@Slf4j
@CrossOrigin
@RestController
@RequiredArgsConstructor
@RequestMapping("/admin/account")
public class AdminController {
    private final AdminService adminService;
    private final JwtProperties jwtProperties;
    private final RedisTemplate<String, String> template;

    @Api
    @PostMapping("/register")
    public Result register(
            @RequestBody AdminForRegister adminForRegister
    ) {

        String captcha = adminForRegister.getCaptcha();

        if (captcha.equals(template.opsForValue().get(RedisConstant.ADMIN_CAPTCHA + adminForRegister.getEmail()))) {
            adminForRegister.setPhone(null);
        } else if (captcha.equals(template.opsForValue().get(RedisConstant.ADMIN_CAPTCHA + adminForRegister.getPhone()))) {
            adminForRegister.setEmail(null);
        } else {
            throw new CaptchaErrorException();
        }
        // 密码加密
        adminForRegister.setPassword(SHA256Encryption.getSHA256(adminForRegister.getPassword()));

        if (adminForRegister.getRole() != 1) {
            adminForRegister.setRole((short) 2);
        }
        log.info("Admin: [" + adminForRegister.getUsername() + "], role：[" + adminForRegister.getRole() + "], createTime: [" + LocalDateTime.now() + "]");
        adminService.register(adminForRegister);
        return Result.success(AccountConstant.ACCOUNT_REGISTER_SUCCESS);
    }

    @Api
    @PostMapping("/verify")
    public Result verify(
            @RequestBody AdminForVerify adminForVerify
    ) {
        String verification = adminService.verify(adminForVerify);
        log.info("Verification code: [" + verification + "] has been sent at "+ new Date());
        return Result.success(VerifyConstant.VERIFY_CODE_SENT);
    }

    @Api
    @PostMapping("/login")
    public Result login(
            @RequestBody AdminForLogin adminForLogin
    ) {
        // 密码加密
        adminForLogin.setPassword(SHA256Encryption.getSHA256(adminForLogin.getPassword()));
        AdminForLoginView admin = this.adminService.login(adminForLogin);
        if (admin != null) {
            HashMap<String, Object> map = new HashMap<>();
            map.put(JwtClaimsConstant.ID, admin.getId());
            String jwt = JwtUtil.createJWT(this.jwtProperties.getSecretKey(), this.jwtProperties.getTtl(), map);
            log.info("Admin: [" + UserContext.getUserId() + "] login at " + new Date());
            return Result.success(new AdminLoginResult(admin.toEntity(), jwt));

        } else {
            log.info("Admin: [" + UserContext.getUserId() + "] login failure at " + new Date());
            return Result.error(AccountConstant.ACCOUNT_LOGIN_FAILURE);
        }
    }

    @Api
    @PostMapping("/logout")
    public Result logout(
            @RequestBody AdminForLogout adminForLogout
    ) {
        adminService.logout(adminForLogout);
        log.info("Admin: [" + UserContext.getUserId() + "] logout at " + new Date());
        return Result.success(AccountConstant.ACCOUNT_LOGOUT_SUCCESS);
    }

    @Api
    @PostMapping("/delete")
    public Result delete(
            @RequestBody AdminForDelete adminForDelete
    ) {
        adminService.delete(adminForDelete);
        if (!adminForDelete.getIds().isEmpty()) {
            log.info("Admin: [" + UserContext.getUserId() + "] delete [" + adminForDelete.getIds() + "] at " + new Date());
        } else {
            log.info("False delete");
        }
        return Result.success("Admin删除成功");
    }

    @Api
    @PostMapping("/update")
    public Result update(
            @RequestBody AdminForUpdate adminForUpdate
    ) {
        // 密码加密
        boolean passwordLoaded = adminForUpdate.isPasswordLoaded();

        if (passwordLoaded) {
            adminForUpdate.setPassword(SHA256Encryption.getSHA256(adminForUpdate.getPassword()));
        }

        adminService.update(adminForUpdate);
        log.info("Admin: [" + UserContext.getUserId() + "] update " + adminForUpdate.getId() + "] at " + new Date());
        return Result.success(AccountConstant.ACCOUNT_CHANGE_SUCCESS);
    }

    @Api
    @PostMapping("/query")
    public Result query(
            @RequestBody AdminForQuery adminForQuery
    ) {
        log.info("Admin: [" + UserContext.getUserId() + "] is queried [" + adminForQuery + "] at " + new Date());
        return Result.success("Admin查询成功", adminService.query(adminForQuery));
    }
}
