package com.pengyou.controller;

import com.pengyou.model.Result;
import com.pengyou.model.dto.postlabel.LabelForQuery;
import com.pengyou.model.dto.postsection.SectionForQuery;
import com.pengyou.model.dto.tag.TagForQuery;
import com.pengyou.service.QueryService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.*;

@Api
@RestController
@RequiredArgsConstructor
@RequestMapping("/user/query")
public class QueryController {
    private final QueryService queryService;

    @Api
    @PostMapping("/label")
    public Result queryLabel(
            @RequestBody LabelForQuery labelForQuery
    ) {
        return Result.success("Label查询成功",queryService.queryLabel(labelForQuery));
    }

    @Api
    @PostMapping("/section")
    public Result querySection(
            @RequestBody SectionForQuery sectionForQuery
    ) {
        return Result.success("Section查询成功",queryService.querySection(sectionForQuery));
    }

    @Api
    @PostMapping("/tag")
    public Result queryTag(
            @RequestBody TagForQuery tagForQuery
    ) {
        return Result.success("Tag查询成功",queryService.queryTag(tagForQuery));
    }

}
