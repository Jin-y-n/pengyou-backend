package com.pengyou.service;


import com.pengyou.model.dto.postsection.PostSectionForAdd;
import com.pengyou.model.dto.postsection.PostSectionForDelete;

public interface PostSectionService {
    void add(PostSectionForAdd postSectionForAdd);
    void delete(PostSectionForDelete postSectionForDelete);
}
