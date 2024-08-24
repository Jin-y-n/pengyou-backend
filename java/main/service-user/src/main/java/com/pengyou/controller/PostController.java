package com.pengyou.controller;

import com.pengyou.model.Result;
import com.pengyou.model.dto.post.UserPostForAdd;
import com.pengyou.model.dto.post.UserPostForDelete;
import com.pengyou.model.dto.post.UserPostForQuery;
import com.pengyou.model.dto.post.UserPostForUpdate;
import com.pengyou.service.PostService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.*;

@Api
@Slf4j
@CrossOrigin
@RestController
@RequiredArgsConstructor
@RequestMapping("/user/post")
public class PostController {
    private final PostService postService;

    @Api
    @PostMapping("/add")
    public Result addPost(
            @RequestBody UserPostForAdd userPostForAdd
    ) {
        postService.addPost(userPostForAdd);
        return Result.success("Post添加成功");
    }

    @Api
    @PostMapping("/delete")
    public Result deletePost(
            @RequestBody UserPostForDelete userPostForDelete
    ) {
        postService.deletePost(userPostForDelete);
        return Result.success("Post删除成功");
    }

    @Api
    @PostMapping("/update")
    public Result updatePost(
            @RequestBody UserPostForUpdate userPostForUpdate
    ) {
        postService.updatePost(userPostForUpdate);
        return Result.success("Post更新成功");
    }

    @Api
    @PostMapping("/query")
    public Result queryPost(
            @RequestBody UserPostForQuery userPostForQuery
    ) {
        return Result.success("Post查询成功", postService.queryPost(userPostForQuery));
    }
}