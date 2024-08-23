package com.pengyou.service.impl;

import com.pengyou.exception.BaseException;
import com.pengyou.model.dto.post.*;
import com.pengyou.model.entity.PostTable;
import com.pengyou.service.PostService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

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
    public List<UserPostForQueryView> queryPost(UserPostForQuery userPostForQuery) {
        List<UserPostForQueryView> execute = sqlClient
                .createQuery(PostTable.$)
                .where(userPostForQuery)
                .select(
                        PostTable.$.fetch(UserPostForQueryView.class)
                )
                .execute();

        if (execute.isEmpty()) {
            throw new BaseException("Post不存在");
        }
        return execute;

    }
}
