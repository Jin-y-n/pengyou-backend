package com.pengyou.service;


import com.pengyou.model.dto.postsection.PostSectionForAdd;
import com.pengyou.model.dto.postsection.PostSectionForDelete;
import com.pengyou.model.dto.postsection.PostSectionForQuery;
import com.pengyou.model.dto.postsection.PostSectionForQueryView;
import org.babyfish.jimmer.Page;

public interface PostSectionService {
    void add(PostSectionForAdd postSectionForAdd);
    void delete(PostSectionForDelete postSectionForDelete);
    Page<PostSectionForQueryView> query(PostSectionForQuery postSectionForQuery);
}
