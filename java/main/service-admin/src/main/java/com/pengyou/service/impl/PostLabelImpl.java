package com.pengyou.service.impl;


import com.pengyou.constant.LabelConstant;
import com.pengyou.constant.SectionConstant;
import com.pengyou.exception.BaseException;
import com.pengyou.exception.label.LabelExistsException;
import com.pengyou.model.dto.postlabel.PostLabelForAdd;
import com.pengyou.model.dto.postlabel.PostLabelForDelete;
import com.pengyou.model.entity.*;
import com.pengyou.service.PostLabelService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.sql.JSqlClient;
import org.babyfish.jimmer.sql.ast.Predicate;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class PostLabelImpl implements PostLabelService {

    private final JSqlClient sqlClient;
    private final PostLabelTable postLabelTable = PostLabelTable.$;

    @Override
    public void add(PostLabelForAdd postLabelForAdd) {
        List<String> execute = sqlClient
                .createQuery(postLabelTable)
                .where(postLabelTable.label().eq(postLabelForAdd.getLabel()))
                .select(postLabelTable.label())
                .execute();
        if (!execute.isEmpty()){
            throw new BaseException(LabelConstant.LABEL_EXISTS);
        }

        sqlClient
                .insert(postLabelForAdd);
    }

    @Override
    public void delete(PostLabelForDelete postLabelForDelete) {
        sqlClient
                .deleteByIds(PostLabel.class, postLabelForDelete.getIds());
    }


}
