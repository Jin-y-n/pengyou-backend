package com.pengyou.service;

import com.pengyou.model.dto.post.*;
import com.pengyou.model.entity.Post;

import java.util.List;

public interface PostService {
    void addPost(UserPostForAdd userPostForAdd);
    void updatePost(UserPostForUpdate userPostForUpdate);
    void deletePost(UserPostForDelete userPostForDelete);
    List<UserPostForQueryView> queryPost(UserPostForQuery userPostForQuery);
}
