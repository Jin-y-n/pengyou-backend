package com.pengyou.service.impl;


import com.pengyou.model.dto.postsection.PostSectionForAdd;
import com.pengyou.model.dto.postsection.PostSectionForDelete;
import com.pengyou.model.entity.PostSection;
import com.pengyou.model.entity.PostSectionTable;
import com.pengyou.service.PostSectionService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class PostSectionImpl implements PostSectionService {
    private final JSqlClient sqlClient;
    private final PostSectionTable postSectionTable = PostSectionTable.$;


    @Override
    public void add(PostSectionForAdd postSectionForAdd) {
        sqlClient
                .insert(postSectionForAdd);
    }

    @Override
    public void delete(PostSectionForDelete postSectionForDelete) {
        sqlClient
                .deleteByIds(PostSection.class, postSectionForDelete.getIds());
    }
}
