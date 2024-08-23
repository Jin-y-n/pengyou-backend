package com.pengyou.controller;


import com.pengyou.constant.LabelConstant;
import com.pengyou.model.Result;
import com.pengyou.model.dto.postlabel.PostLabelForAdd;
import com.pengyou.model.dto.postlabel.PostLabelForDelete;
import com.pengyou.model.dto.postlabel.PostLabelForQuery;
import com.pengyou.service.PostLabelService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Date;

@Api
@Slf4j
@RestController
@RequiredArgsConstructor
@RequestMapping("/admin/post-label")
public class PostLabelController {
    private final PostLabelService postLabelService;

    @Api
    @PostMapping("/add")
    public Result add(
            @RequestBody PostLabelForAdd postLabelForAdd
    ) {
        log.info("PostLabel:" + postLabelForAdd.getLabel() + "is got added at " + new Date());
        postLabelService.add(postLabelForAdd);
        return Result.success(LabelConstant.LABEL_ADD_SUCCESS);
    }

    @Api
    @PostMapping("/delete")
    public Result delete(
            @RequestBody PostLabelForDelete postLabelForDelete
    ) {
        postLabelService.delete(postLabelForDelete);
        if (!postLabelForDelete.getIds().isEmpty()) {
            log.info("Admin:" + postLabelForDelete.getIds().toString() + "delete at " + new Date());
        } else {
            log.info("False delete");
        }
        return Result.success(LabelConstant.LABEL_DELETE_SUCCESS);
    }

    @Api
    @PostMapping("/query")
    public Result query(
            @RequestBody PostLabelForQuery postLabelForQuery
    ) {
        log.info("PostLabel: is queried at " + new Date());
        return Result.success(postLabelService.query(postLabelForQuery));
    }
}
