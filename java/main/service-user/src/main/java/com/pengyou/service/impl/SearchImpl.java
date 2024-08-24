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
import org.babyfish.jimmer.Page;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;


@Service
@RequiredArgsConstructor
public class SearchImpl implements SearchService {
    private final JSqlClient sqlClient;

    @Override
    public Page<SearchForUsersView> searchUsers(SearchForUsers searchForUsers) {
        Page<SearchForUsersView> page = sqlClient
                .createQuery(UserTable.$)
                .where(searchForUsers)
                .select(
                        UserTable.$.fetch(SearchForUsersView.class)
                )
                .fetchPage(searchForUsers.getPageIndex(), searchForUsers.getPageSize());

        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Post不存在");
        }
        return page;
    }

    @Override
    public Page<SearchForTagsView> searchTags(SearchForTags searchForTags) {
        Page<SearchForTagsView> page = sqlClient
                .createQuery(TagTable.$)
                .where(searchForTags)
                .select(
                        TagTable.$.fetch(SearchForTagsView.class)
                )
                .fetchPage(searchForTags.getPageIndex(), searchForTags.getPageSize());

        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Post不存在");
        }
        return page;
    }

    @Override
    public Page<SearchForLabelsView> searchLabels(SearchForLabels searchForLabels) {
        Page<SearchForLabelsView> page = sqlClient
                .createQuery(PostLabelTable.$)
                .where(searchForLabels)
                .select(
                        PostLabelTable.$.fetch(SearchForLabelsView.class)
                )
                .fetchPage(searchForLabels.getPageIndex(), searchForLabels.getPageSize());

        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Post不存在");
        }
        return page;
    }

    @Override
    public Page<SearchForSectionsView> searchSections(SearchForSections searchForSections) {
        Page<SearchForSectionsView> page = sqlClient
                .createQuery(PostSectionTable.$)
                .where(searchForSections)
                .select(
                        PostSectionTable.$.fetch(SearchForSectionsView.class)
                )
                .fetchPage(searchForSections.getPageIndex(), searchForSections.getPageSize());

        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Post不存在");
        }
        return page;
    }

    @Override
    public Page<SearchForPostsView> searchPosts(SearchForPosts searchForPosts) {
        Page<SearchForPostsView> page = sqlClient
                .createQuery(PostTable.$)
                .where(searchForPosts)
                .select(
                        PostTable.$.fetch(SearchForPostsView.class)
                )
                .fetchPage(searchForPosts.getPageIndex(), searchForPosts.getPageSize());


        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Post不存在");
        }
        return page;
    }
}
