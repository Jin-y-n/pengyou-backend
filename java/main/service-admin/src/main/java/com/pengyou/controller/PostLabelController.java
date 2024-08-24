package com.pengyou.controller;


import com.pengyou.constant.LabelConstant;
import com.pengyou.model.Result;
import com.pengyou.model.dto.postlabel.PostLabelForAdd;
import com.pengyou.model.dto.postlabel.PostLabelForDelete;
import com.pengyou.model.dto.postlabel.PostLabelForQuery;
import com.pengyou.service.PostLabelService;
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
@RequestMapping("/admin/post-label")
public class PostLabelController {
    private final PostLabelService postLabelService;

    @Api
    @PostMapping("/add")
    public Result add(
            @RequestBody PostLabelForAdd postLabelForAdd
    ) {
        postLabelService.add(postLabelForAdd);
        log.info("Admin: [" + UserContext.getUserId() + "] add PostLabel: [" + postLabelForAdd.getLabel() + "] at " + new Date());
        return Result.success(LabelConstant.LABEL_ADD_SUCCESS);
    }

    @Api
    @PostMapping("/delete")
    public Result delete(
            @RequestBody PostLabelForDelete postLabelForDelete
    ) {
        postLabelService.delete(postLabelForDelete);
        if (!postLabelForDelete.getIds().isEmpty()) {
            log.info("Admin: [" + UserContext.getUserId() + "] delete PostLabel: [" + postLabelForDelete.getIds() + "] at " + new Date());
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
        log.info("Admin: [" + UserContext.getUserId() + "] query PostLabel: [" + postLabelForQuery.getLabel() + "] at " + new Date());
        return Result.success(postLabelService.query(postLabelForQuery));
    }
}
