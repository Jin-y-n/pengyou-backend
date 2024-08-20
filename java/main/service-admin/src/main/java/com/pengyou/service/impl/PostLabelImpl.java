package com.pengyou.service.impl;


import com.pengyou.model.dto.postlabel.PostLabelForAdd;
import com.pengyou.model.dto.postlabel.PostLabelForDelete;
import com.pengyou.model.entity.Fetchers;
import com.pengyou.model.entity.PostLabel;
import com.pengyou.model.entity.PostLabelTable;
import com.pengyou.model.entity.PostLabelTableEx;
import com.pengyou.service.PostLabelService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class PostLabelImpl implements PostLabelService {

    private final JSqlClient sqlClient;
    private final PostLabelTable postLabelTable = PostLabelTable.$;

    @Override
    public void add(PostLabelForAdd postLabelForAdd) {
        sqlClient
                .insert(postLabelForAdd);
    }

    @Override
    public void delete(PostLabelForDelete postLabelForDelete) {
        sqlClient
                .deleteByIds(PostLabel.class, postLabelForDelete.getIds());
    }


}
