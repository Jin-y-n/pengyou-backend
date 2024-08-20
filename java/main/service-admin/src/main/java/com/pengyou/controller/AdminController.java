package com.pengyou.controller;

import com.pengyou.config.properties.JwtProperties;
import com.pengyou.constant.AccountConstant;
import com.pengyou.constant.JwtClaimsConstant;
import com.pengyou.constant.RedisConstant;
import com.pengyou.exception.common.CaptchaErrorException;
import com.pengyou.model.Result;
import com.pengyou.model.dto.admin.*;
import com.pengyou.model.response.AdminLoginResult;
import com.pengyou.service.AdminService;
import com.pengyou.util.security.JwtUtil;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.client.meta.Api;
import org.babyfish.jimmer.sql.ast.tuple.Tuple2;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@Api
@RestController
@RequiredArgsConstructor
@RequestMapping("/admin")
public class AdminController {
    private final AdminService adminService;
    private final JwtProperties jwtProperties;
    private final RedisTemplate<String,String> template;

    @Api
    @PostMapping("/register")
    public Result register(
            @RequestBody AdminForRegister adminForRegister
    ) {

        String captcha = adminForRegister.getCaptcha();

        if (captcha.equals(template.opsForValue().get(RedisConstant.ADMIN_CAPTCHA+adminForRegister.getEmail()))) {
            adminForRegister.setPhone(null);
        } else if (captcha.equals(template.opsForValue().get(RedisConstant.ADMIN_CAPTCHA+adminForRegister.getPhone()))) {
            adminForRegister.setEmail(null);
        } else {
            throw new CaptchaErrorException();
        }

        adminService.register(adminForRegister);
        return Result.success(AccountConstant.ACCOUNT_REGISTER_SUCCESS);
    }

    @Api
    @PostMapping("/verify")
    public Result verify(
            @RequestBody AdminForVerify adminForVerify
    ) {
        adminService.verify(adminForVerify);
        return Result.success("Admin验证成功");
    }

    @Api
    @PostMapping("/login")
    public Result login(
            @RequestBody AdminForLogin adminForLogin
    ) {
        AdminForLoginView admin = this.adminService.login(adminForLogin);
        if (admin != null) {
            HashMap<String, Object> map = new HashMap<>();
            map.put(JwtClaimsConstant.ID, admin.getId());
            String jwt = JwtUtil.createJWT(this.jwtProperties.getSecretKey(), this.jwtProperties.getTtl(), map);

            return Result.success(new AdminLoginResult(admin.toEntity(), jwt));

        } else {
            return Result.error(AccountConstant.ACCOUNT_LOGIN_FAILURE);
        }
    }

    @Api
    @PostMapping("/logout")
    public Result logout(
            @RequestBody AdminForLogout adminForLogout
    ) {
        adminService.logout(adminForLogout);
        return Result.success("Admin登出成功");
    }

    @Api
    @PostMapping("/delete")
    public Result delete(
            @RequestBody AdminForDelete adminForDelete
    ) {
        adminService.delete(adminForDelete);
        return Result.success("Admin删除成功");
    }

    @Api
    @PostMapping("/update")
    public Result update(
            @RequestBody AdminForUpdate adminForUpdate
    ) {
        adminService.update(adminForUpdate);
        return Result.success("Admin更新成功");
    }

    @Api
    @PostMapping("/query")
    public Result query(
            @RequestBody AdminForQuery adminForQuery
    ) {
        return Result.success("Admin查询成功",adminService.query(adminForQuery));
    }
}
