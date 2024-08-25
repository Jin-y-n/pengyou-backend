package com.pengyou.controller;

import com.pengyou.model.Result;
import com.pengyou.model.dto.Friend.FriendForAccept;
import com.pengyou.model.dto.Friend.FriendForAdd;
import com.pengyou.model.dto.Friend.FriendForDelete;
import com.pengyou.model.dto.Friend.FriendForUpgrade;
import com.pengyou.model.dto.Socialaccount.AccountForAdd;
import com.pengyou.model.dto.Socialaccount.AccountForDelete;
import com.pengyou.service.FriendService;
import com.pengyou.service.PostService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.client.meta.Api;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.web.bind.annotation.*;

@Api
@Slf4j
@CrossOrigin
@RestController
@RequiredArgsConstructor
@RequestMapping("/user/friend")
public class FriendController {

    private final FriendService friendService;
    private final JSqlClient jSqlClient;

    @Api
    @PostMapping("/add")
    public Result add(
            @RequestBody FriendForAdd friend
    ) {
        friendService.add(friend);
        return Result.success();
    }

    @Api
    @PostMapping("/delete")
    public Result delete(
            @RequestBody FriendForDelete friend
    ) {
        friendService.delete(friend);
        return Result.success();
    }

    @Api
    @PostMapping("/accept")
    public Result accept(
            @RequestBody FriendForAccept friend
    ) {
        friendService.accept(friend);
        return Result.success();
    }

    @Api
    @PostMapping("/upgrade")
    public Result upgrade(
            @RequestBody FriendForUpgrade friend
    ) {
        friendService.upgrade(friend);
        return Result.success();
    }

    @Api
    @PostMapping("/social-account/add")
    public Result addSocialAccount(
            @RequestBody AccountForAdd account
    ) {
        friendService.addSocialAccount(account);
        return Result.success();
    }

    @Api
    @PostMapping("/social-account/delete")
    public Result deleteSocialAccount(
            @RequestBody AccountForDelete account
    ) {
        friendService.deleteSocialAccount(account);
        return Result.success();
    }
}
