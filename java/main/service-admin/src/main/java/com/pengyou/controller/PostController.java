package com.pengyou.controller;

import com.pengyou.model.Result;
import com.pengyou.model.dto.admin.AdminForLogin;
import com.pengyou.model.dto.post.PostForDelete;
import com.pengyou.model.dto.post.PostForQuery;
import com.pengyou.model.dto.post.PostForUpdate;
import com.pengyou.service.PostService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@Api
@RestController
@RequiredArgsConstructor
@RequestMapping("/admin/post")
public class PostController {
    private final PostService postService;

    @Api
    @PostMapping("/query")
    public Result query(
            @RequestBody PostForQuery postForQuery
    ) {
        return Result.success("Post查询成功", postService.query(postForQuery));
    }

    @Api
    @PostMapping("/delete")
    public Result delete(
            @RequestBody PostForDelete postForDelete
    ) {
        postService.delete(postForDelete);
        return Result.success("Post删除成功");
    }

    @Api
    @PostMapping("/update")
    public Result update(
            @RequestBody PostForUpdate postForUpdate
    ) {
        postService.update(postForUpdate);
        return Result.success("Post更新成功");
    }
}
