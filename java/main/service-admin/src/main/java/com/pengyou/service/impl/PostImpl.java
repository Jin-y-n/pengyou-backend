package com.pengyou.service.impl;


import com.pengyou.model.dto.post.PostForDelete;
import com.pengyou.model.dto.post.PostForQuery;
import com.pengyou.model.dto.post.PostForUpdate;
import com.pengyou.model.dto.post.PostForView;
import com.pengyou.model.entity.Post;
import com.pengyou.model.entity.PostTable;
import com.pengyou.service.PostService;
import lombok.RequiredArgsConstructor;
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
    public List<PostForView> query(PostForQuery postForQuery) {
        return sqlClient
                .createQuery(postTable)
                .where(postForQuery)
                .select(
                        postTable.fetch(PostForView.class)
                )
                .execute();
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
