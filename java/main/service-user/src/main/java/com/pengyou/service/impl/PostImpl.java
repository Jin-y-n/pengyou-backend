package com.pengyou.service.impl;

import com.pengyou.exception.BaseException;
import com.pengyou.model.dto.post.*;
import com.pengyou.model.entity.PostTable;
import com.pengyou.service.PostService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.Page;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class PostImpl implements PostService {
    private final JSqlClient sqlClient;

    @Override
    public void addPost(UserPostForAdd userPostForAdd) {
        sqlClient
                .insert(userPostForAdd);
    }

    @Override
    public void updatePost(UserPostForUpdate userPostForUpdate) {
        sqlClient
                .update(userPostForUpdate);
    }

    @Override
    public void deletePost(UserPostForDelete userPostForDelete) {
        sqlClient
                .deleteByIds(UserPostForDelete.class, userPostForDelete.getIds());
    }

    @Override
    public Page<UserPostForQueryView> queryPost(UserPostForQuery userPostForQuery) {
        Page<UserPostForQueryView> page = sqlClient
                .createQuery(PostTable.$)
                .where(userPostForQuery)
                .select(
                        PostTable.$.fetch(UserPostForQueryView.class)
                )
                .fetchPage(userPostForQuery.getPageIndex(), userPostForQuery.getPageSize());

        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Post不存在");
        }
        return page;
    }

}
