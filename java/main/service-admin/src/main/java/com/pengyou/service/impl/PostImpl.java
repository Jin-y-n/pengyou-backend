package com.pengyou.service.impl;


import com.pengyou.exception.BaseException;
import com.pengyou.model.dto.post.PostForDelete;
import com.pengyou.model.dto.post.PostForQuery;
import com.pengyou.model.dto.post.PostForUpdate;
import com.pengyou.model.dto.post.PostForView;
import com.pengyou.model.entity.AdminTable;
import com.pengyou.model.entity.Post;
import com.pengyou.model.entity.PostTable;
import com.pengyou.service.PostService;
import com.pengyou.util.RedisLock;
import com.pengyou.util.UserContext;
import io.reactivex.rxjava3.exceptions.QueueOverflowException;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.Page;
import org.babyfish.jimmer.sql.JSqlClient;
import org.babyfish.jimmer.sql.ast.mutation.SaveMode;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Service
@RequiredArgsConstructor
public class PostImpl implements PostService {
    private final JSqlClient sqlClient;
    private final PostTable postTable = PostTable.$;
    private final RedisLock redisLock;

    @Override
    public Page<PostForView> query(PostForQuery postForQuery) {
        Page<PostForView> page = sqlClient
                .createQuery(postTable)
                .where(postForQuery)
                .select(
                        postTable.fetch(PostForView.class)
                )
                .fetchPage(postForQuery.getPageIndex() == null ? 0 : postForQuery.getPageIndex()
                        , postForQuery.getPageSize() == null ? 10 : postForQuery.getPageSize());
        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Post查询失败");
        }
        return page;
    }

    @Override
    public void delete(PostForDelete postForDelete) {
        sqlClient
                .deleteByIds(Post.class, postForDelete.getIds());
    }

    @Transactional
    @Override
    public void update(PostForUpdate postForUpdate) {
        redisLock.lock();
        sqlClient
                .createUpdate(postTable)
                .set(postTable.modifiedAt(), LocalDateTime.now())
                .execute();

        List<Short> execute = sqlClient
                .createQuery(AdminTable.$)
                .where(AdminTable.$.id().eq(Long.valueOf(UserContext.getUserId())))
                .select(AdminTable.$.role())
                .execute();
        Short first = execute.stream().findFirst().get();
        if (first == 1) {
            sqlClient
                    .createUpdate(postTable)
                    .set(postTable.modifiedPerson(), Long.valueOf(UserContext.getUserId()))
                    .set(postTable.author().modifiedByAdmin(), (short) 1)
                    .where(postTable.id().eq(postForUpdate.getId()))
                    .execute();
        } else if (first == 2) {
            sqlClient
                    .createUpdate(postTable)
                    .set(postTable.modifiedPerson(), Long.valueOf(UserContext.getUserId()))
                    .set(postTable.author().modifiedByAdmin(), (short) 0)
                    .where(postTable.id().eq(postForUpdate.getId()))
                    .execute();
        }

        sqlClient
                .update(postForUpdate);
        redisLock.unlock();
    }
}
