package com.pengyou.service.impl;

import com.pengyou.exception.BaseException;
import com.pengyou.model.dto.post.SearchForPosts;
import com.pengyou.model.dto.post.SearchForPostsView;
import com.pengyou.model.dto.postlabel.SearchForLabels;
import com.pengyou.model.dto.postlabel.SearchForLabelsView;
import com.pengyou.model.dto.postsection.SearchForSections;
import com.pengyou.model.dto.postsection.SearchForSectionsView;
import com.pengyou.model.dto.tag.SearchForTags;
import com.pengyou.model.dto.tag.SearchForTagsView;
import com.pengyou.model.dto.user.SearchForUsers;
import com.pengyou.model.dto.user.SearchForUsersView;
import com.pengyou.model.entity.*;
import com.pengyou.service.SearchService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class SearchImpl implements SearchService {
    private final JSqlClient sqlClient;

    @Override
    public List<SearchForUsersView> searchUsers(SearchForUsers searchForUsers) {
        List<SearchForUsersView> execute = sqlClient
                .createQuery(UserTable.$)
                .where(searchForUsers)
                .select(
                        UserTable.$.fetch(SearchForUsersView.class)
                )
                .execute();
        if (execute.isEmpty()){
            throw new BaseException("用户查询失败");
        }
        return execute;
    }

    @Override
    public List<SearchForTagsView> searchTags(SearchForTags searchForTags) {
        List<SearchForTagsView> execute = sqlClient
                .createQuery(TagTable.$)
                .where(searchForTags)
                .select(
                        TagTable.$.fetch(SearchForTagsView.class)
                )
                .execute();
        if (execute.isEmpty()){
            throw new BaseException("用户查询失败");
        }
        return execute;
    }

    @Override
    public List<SearchForLabelsView> searchLabels(SearchForLabels searchForLabels) {
        List<SearchForLabelsView> execute = sqlClient
                .createQuery(PostLabelTable.$)
                .where(searchForLabels)
                .select(
                        PostLabelTable.$.fetch(SearchForLabelsView.class)
                )
                .execute();
        if (execute.isEmpty()){
            throw new BaseException("用户查询失败");
        }
        return execute;
    }

    @Override
    public List<SearchForSectionsView> searchSections(SearchForSections searchForSections) {
        List<SearchForSectionsView> execute = sqlClient
                .createQuery(PostSectionTable.$)
                .where(searchForSections)
                .select(
                        PostSectionTable.$.fetch(SearchForSectionsView.class)
                )
                .execute();
        if (execute.isEmpty()){
            throw new BaseException("用户查询失败");
        }
        return execute;
    }

    @Override
    public List<SearchForPostsView> searchPosts(SearchForPosts searchForPosts) {
        List<SearchForPostsView> execute = sqlClient
                .createQuery(PostTable.$)
                .where(searchForPosts)
                .select(
                        PostTable.$.fetch(SearchForPostsView.class)
                )
                .execute();
        if (execute.isEmpty()){
            throw new BaseException("用户查询失败");
        }
        return execute;
    }
}
