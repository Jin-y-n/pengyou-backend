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
import org.babyfish.jimmer.Page;



public interface SearchService {
    Page<SearchForUsersView> searchUsers(SearchForUsers searchForUsers);
    Page<SearchForTagsView> searchTags(SearchForTags searchForTags);
    Page<SearchForLabelsView> searchLabels(SearchForLabels searchForLabel);
    Page<SearchForSectionsView> searchSections(SearchForSections searchForSections);
    Page<SearchForPostsView> searchPosts(SearchForPosts searchForPosts);
}
