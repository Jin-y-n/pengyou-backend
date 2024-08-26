package com.pengyou.controller;


import com.pengyou.model.Result;
import com.pengyou.model.dto.user.AdminUserForQuery;
import com.pengyou.model.dto.user.UserForLogin;
import com.pengyou.service.UserService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.*;

@Api
@Slf4j
@CrossOrigin
@RestController
@RequiredArgsConstructor
@RequestMapping("/admin/user/")
public class UserController {
    private final UserService userService;

    @Api
    @PostMapping("/query")
    public Result query(
            @RequestBody AdminUserForQuery adminUserForQuery
            ) {
        return Result.success(userService.query(adminUserForQuery));
    }


}
