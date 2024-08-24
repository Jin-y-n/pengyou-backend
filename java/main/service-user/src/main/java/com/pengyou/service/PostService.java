package com.pengyou.service;

import com.pengyou.model.dto.post.*;
import org.babyfish.jimmer.Page;


public interface PostService {
    void addPost(UserPostForAdd userPostForAdd);
    void updatePost(UserPostForUpdate userPostForUpdate);
    void deletePost(UserPostForDelete userPostForDelete);
    Page<UserPostForQueryView> queryPost(UserPostForQuery userPostForQuery);
}
