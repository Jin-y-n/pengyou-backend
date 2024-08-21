package com.pengyou.controller;


import com.pengyou.constant.LabelConstant;
import com.pengyou.model.Result;
import com.pengyou.model.dto.postlabel.PostLabelForAdd;
import com.pengyou.model.dto.postlabel.PostLabelForDelete;
import com.pengyou.service.PostLabelService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@Api
@RestController
@RequiredArgsConstructor
@RequestMapping("/admin/post-label")
public class PostLabelController {
    private final PostLabelService postLabelService;

    @Api
    @PostMapping("/add")
    public Result add(
            @RequestBody PostLabelForAdd postLabelForAdd
    ){
        postLabelService.add(postLabelForAdd);
        return Result.success(LabelConstant.LABEL_ADD_SUCCESS);
    }
    @Api
    @PostMapping("/delete")
    public Result delete(
            @RequestBody PostLabelForDelete postLabelForDelete
    ){
        postLabelService.delete(postLabelForDelete);
        return Result.success(LabelConstant.LABEL_DELETE_SUCCESS);
    }
}
