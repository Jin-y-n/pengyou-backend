package com.pengyou.service;


import com.pengyou.model.dto.post.PostForDelete;
import com.pengyou.model.dto.post.PostForQuery;
import com.pengyou.model.dto.post.PostForUpdate;
import com.pengyou.model.dto.post.PostForView;
import com.pengyou.model.entity.Post;

import java.util.List;

public interface PostService {
    List<PostForView> query(PostForQuery postForQuery);
    void delete(PostForDelete postForDelete);
    void update(PostForUpdate postForUpdate);
}
