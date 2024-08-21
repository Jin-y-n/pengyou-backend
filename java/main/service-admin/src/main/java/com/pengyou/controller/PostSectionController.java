package com.pengyou.controller;

import com.pengyou.constant.SectionConstant;
import com.pengyou.model.Result;
import com.pengyou.model.dto.postsection.PostSectionForAdd;
import com.pengyou.model.dto.postsection.PostSectionForDelete;
import com.pengyou.service.PostSectionService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@Api
@RestController
@RequiredArgsConstructor
@RequestMapping("/admin/post-section")
public class PostSectionController {
    private final PostSectionService postSectionService;

    @Api
    @PostMapping("/add")
    public Result add(
            @RequestBody PostSectionForAdd postSectionForAdd
    ){
        postSectionService.add(postSectionForAdd);
        return Result.success(SectionConstant.SECTION_ADD_SUCCESS);
    }

    @Api
    @PostMapping("/delete")
    public Result delete(
            @RequestBody PostSectionForDelete postSectionForDelete
    ){
        postSectionService.delete(postSectionForDelete);
        return Result.success(SectionConstant.SECTION_DELETE_SUCCESS);
    }
}
