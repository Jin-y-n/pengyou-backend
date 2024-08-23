package com.pengyou.controller;

import com.pengyou.model.Result;
import com.pengyou.model.dto.postlabel.LabelForQuery;
import com.pengyou.model.dto.postsection.SectionForQuery;
import com.pengyou.model.dto.tag.TagForQuery;
import com.pengyou.service.QueryService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.*;

@Api
@Slf4j
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
        log.info("Query PostLabel");
        return Result.success("PostLabel查询成功", queryService.queryLabel(labelForQuery));
    }

    @Api
    @PostMapping("/section")
    public Result querySection(
            @RequestBody SectionForQuery sectionForQuery
    ) {
        log.info("Query PostSection");
        return Result.success("PostSection查询成功", queryService.querySection(sectionForQuery));
    }

    @Api
    @PostMapping("/tag")
    public Result queryTag(
            @RequestBody TagForQuery tagForQuery
    ) {
        log.info("Query Tag");
        return Result.success("Tag查询成功", queryService.queryTag(tagForQuery));
    }

}
