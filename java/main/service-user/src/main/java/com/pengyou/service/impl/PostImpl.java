package com.pengyou.service.impl;

import com.pengyou.exception.BaseException;
import com.pengyou.model.dto.post.*;
import com.pengyou.model.entity.PostTable;
import com.pengyou.service.PostService;
import com.pengyou.util.UserContext;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.Page;
import org.babyfish.jimmer.sql.JSqlClient;
import org.babyfish.jimmer.sql.ast.Predicate;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class PostImpl implements PostService {
    private final JSqlClient sqlClient;
    private final PostTable postTable = PostTable.$;

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
                .createDelete(postTable)
                .where(
                        Predicate.and(
                                PostTable.$.id().in(userPostForDelete.getIds()),
                                PostTable.$.author().id().eq(Long.valueOf(UserContext.getUserId()))
                        )
                )
                .execute();
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
