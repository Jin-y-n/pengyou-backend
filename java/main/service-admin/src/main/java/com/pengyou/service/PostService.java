package com.pengyou.service;


import com.pengyou.model.dto.post.PostForDelete;
import com.pengyou.model.dto.post.PostForQuery;
import com.pengyou.model.dto.post.PostForUpdate;
import com.pengyou.model.dto.post.PostForView;
import org.babyfish.jimmer.Page;


public interface PostService {
    Page<PostForView> query(PostForQuery postForQuery);
    void delete(PostForDelete postForDelete);
    void update(PostForUpdate postForUpdate);
}
