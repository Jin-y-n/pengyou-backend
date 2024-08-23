package com.pengyou.service.impl;


import com.pengyou.exception.BaseException;
import com.pengyou.model.dto.post.PostForDelete;
import com.pengyou.model.dto.post.PostForQuery;
import com.pengyou.model.dto.post.PostForUpdate;
import com.pengyou.model.dto.post.PostForView;
import com.pengyou.model.entity.Post;
import com.pengyou.model.entity.PostTable;
import com.pengyou.service.PostService;
import io.reactivex.rxjava3.exceptions.QueueOverflowException;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.Page;
import org.babyfish.jimmer.sql.JSqlClient;
import org.babyfish.jimmer.sql.ast.mutation.SaveMode;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class PostImpl implements PostService {
    private final JSqlClient sqlClient;
    private final PostTable postTable = PostTable.$;

    @Override
    public Page<PostForView> query(PostForQuery postForQuery) {
        Page<PostForView> page = sqlClient
                .createQuery(postTable)
                .where(postForQuery)
                .select(
                        postTable.fetch(PostForView.class)
                )
                .fetchPage(postForQuery.getPageIndex(), postForQuery.getPageSize());
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

    @Override
    public void update(PostForUpdate postForUpdate) {
        sqlClient
                .update(postForUpdate);
    }
}
