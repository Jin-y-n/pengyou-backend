package com.pengyou.controller;

import com.pengyou.model.Result;
import com.pengyou.model.dto.post.PostForDelete;
import com.pengyou.model.dto.post.PostForQuery;
import com.pengyou.model.dto.post.PostForUpdate;
import com.pengyou.service.PostService;
import com.pengyou.util.UserContext;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.*;

import java.util.Date;

@Api
@CrossOrigin
@Slf4j
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
        log.info("Admin: [" + UserContext.getUserId() + "] query Post: [" + postForQuery + "] at " + new Date());
        return Result.success("Post查询成功", postService.query(postForQuery));
    }

    @Api
    @PostMapping("/delete")
    public Result delete(
            @RequestBody PostForDelete postForDelete
    ) {
        postService.delete(postForDelete);
        if (!postForDelete.getIds().isEmpty()) {
            log.info("Admin: [" + UserContext.getUserId() + "] delete Post: [" + postForDelete.getIds() + "] at " + new Date());
        } else {
            log.info("False delete");
        }
        return Result.success("Post删除成功");
    }

    @Api
    @PostMapping("/update")
    public Result update(
            @RequestBody PostForUpdate postForUpdate
    ) {
        postService.update(postForUpdate);
        log.info("Admin: [" + UserContext.getUserId() + "] update Post: [" + postForUpdate.getId() + "] at " + new Date());
        return Result.success("Post更新成功");
    }
}
