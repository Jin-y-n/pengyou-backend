package com.pengyou.controller;

import com.pengyou.model.Result;
import com.pengyou.model.dto.post.SearchForPosts;
import com.pengyou.model.dto.postlabel.SearchForLabels;
import com.pengyou.model.dto.postsection.SearchForSections;
import com.pengyou.model.dto.tag.SearchForTags;
import com.pengyou.model.dto.user.SearchForUsers;
import com.pengyou.service.SearchService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@Api
@RestController
@RequiredArgsConstructor
@RequestMapping("/user/search")
public class SearchController {
    private final SearchService searchService;

    @Api
    @PostMapping("/users")
    public Result users(
            @RequestBody SearchForUsers searchForUsers
    ) {

        return Result.success("Users查询成功",searchService.searchUsers(searchForUsers));
    }

    @Api
    @PostMapping("/posts")
    public Result posts(
            @RequestBody SearchForPosts searchForPosts
    ) {
        return Result.success("Posts查询成功",searchService.searchPosts(searchForPosts));
    }

    @Api
    @PostMapping("/tags")
    public Result tags(
            @RequestBody SearchForTags searchFortags
    ) {
        return Result.success("Tags查询成功",searchService.searchTags(searchFortags));
    }

    @Api
    @PostMapping("/labels")
    public Result labels(
            @RequestBody SearchForLabels searchForlabels
    ) {
        return Result.success("Labels查询成功",searchService.searchLabels(searchForlabels));
    }

    @Api
    @PostMapping("/sections")
    public Result sections(
            @RequestBody SearchForSections searchForSections
    ) {
        return Result.success("Sections查询成功",searchService.searchSections(searchForSections));
    }
}
