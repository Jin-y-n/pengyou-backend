package com.pengyou.service;

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

import java.util.List;

public interface SearchService {
    List<SearchForUsersView> searchUsers(SearchForUsers searchForUsers);
    List<SearchForTagsView> searchTags(SearchForTags searchForTags);
    List<SearchForLabelsView> searchLabels(SearchForLabels searchForLabel);
    List<SearchForSectionsView> searchSections(SearchForSections searchForSections);
    List<SearchForPostsView> searchPosts(SearchForPosts searchForPosts);
}
