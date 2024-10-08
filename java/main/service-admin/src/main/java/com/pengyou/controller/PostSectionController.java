package com.pengyou.controller;

import com.pengyou.constant.SectionConstant;
import com.pengyou.model.Result;
import com.pengyou.model.dto.postsection.PostSectionForAdd;
import com.pengyou.model.dto.postsection.PostSectionForDelete;
import com.pengyou.model.dto.postsection.PostSectionForQuery;
import com.pengyou.service.PostSectionService;
import com.pengyou.util.UserContext;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.*;

import java.util.Date;

@Api
@Slf4j
@CrossOrigin
@RestController
@RequiredArgsConstructor
@RequestMapping("/admin/post-section")
public class PostSectionController {
    private final PostSectionService postSectionService;

    @Api
    @PostMapping("/add")
    public Result add(
            @RequestBody PostSectionForAdd postSectionForAdd
    ) {
        postSectionService.add(postSectionForAdd);
        log.info("Admin: [" + UserContext.getUserId() + "] add PostSection: [" + postSectionForAdd.getSection() + "] at " + new Date());
        return Result.success(SectionConstant.SECTION_ADD_SUCCESS);
    }

    @Api
    @PostMapping("/delete")
    public Result delete(
            @RequestBody PostSectionForDelete postSectionForDelete
    ) {
        postSectionService.delete(postSectionForDelete);
        if (!postSectionForDelete.getIds().isEmpty()) {
            log.info("Admin: [" + UserContext.getUserId() + "] delete PostSection: [" + postSectionForDelete.getIds() + "] at " + new Date());
        } else {
            log.info("False delete");
        }
        return Result.success(SectionConstant.SECTION_DELETE_SUCCESS);
    }

    @Api
    @PostMapping("/query")
    public Result query(
            @RequestBody PostSectionForQuery postSectionForQuery
    ) {
        log.info("Admin: [" + UserContext.getUserId() + "] query PostSection: [" + postSectionForQuery.getSection() + "] at " + new Date());
        return Result.success(postSectionService.query(postSectionForQuery));
    }
}
