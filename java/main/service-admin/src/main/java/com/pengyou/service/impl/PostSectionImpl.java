package com.pengyou.service.impl;


import com.pengyou.constant.SectionConstant;
import com.pengyou.exception.BaseException;
import com.pengyou.model.dto.postsection.PostSectionForAdd;
import com.pengyou.model.dto.postsection.PostSectionForDelete;
import com.pengyou.model.dto.postsection.PostSectionForQuery;
import com.pengyou.model.dto.postsection.PostSectionForQueryView;
import com.pengyou.model.entity.PostSection;
import com.pengyou.model.entity.PostSectionTable;
import com.pengyou.service.PostSectionService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.Page;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class PostSectionImpl implements PostSectionService {
    private final JSqlClient sqlClient;
    private final PostSectionTable postSectionTable = PostSectionTable.$;


    @Override
    public void add(PostSectionForAdd postSectionForAdd) {
        List<String> execute = sqlClient
                .createQuery(postSectionTable)
                .where(postSectionTable.section().eq(postSectionForAdd.getSection()))
                .select(postSectionTable.section())
                .execute();
        if (!execute.isEmpty()) {
            throw new BaseException(SectionConstant.SECTION_EXISTS);
        }

        sqlClient
                .insert(postSectionForAdd);
    }

    @Override
    public void delete(PostSectionForDelete postSectionForDelete) {
        sqlClient
                .deleteByIds(PostSection.class, postSectionForDelete.getIds());
    }

    @Override
    public Page<PostSectionForQueryView> query(PostSectionForQuery postSectionForQuery) {
        Page<PostSectionForQueryView> execute = sqlClient
                .createQuery(postSectionTable)
                .where(postSectionForQuery)
                .select(
                        postSectionTable.fetch(PostSectionForQueryView.class)
                )
                .fetchPage(postSectionForQuery.getPageIndex(), postSectionForQuery.getPageSize());

        if (execute.getTotalRowCount() == 0) {
            throw new BaseException("PostSection查询失败");
        }
        return execute;
    }
}
