package com.pengyou.controller;

import com.pengyou.model.Result;
import com.pengyou.model.dto.post.SearchForPosts;
import com.pengyou.model.dto.postlabel.SearchForLabels;
import com.pengyou.model.dto.postsection.SearchForSections;
import com.pengyou.model.dto.tag.SearchForTags;
import com.pengyou.model.dto.user.SearchForUsers;
import com.pengyou.service.SearchService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.web.bind.annotation.*;

@Api
@Slf4j
@CrossOrigin
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
        log.info("Search Users");
        return Result.success("Users查询成功", searchService.searchUsers(searchForUsers));
    }

    @Api
    @PostMapping("/posts")
    public Result posts(
            @RequestBody SearchForPosts searchForPosts
    ) {
        log.info("Search Posts");
        return Result.success("Posts查询成功", searchService.searchPosts(searchForPosts));
    }

    @Api
    @PostMapping("/tags")
    public Result tags(
            @RequestBody SearchForTags searchFortags
    ) {
        log.info("Search Tags");
        return Result.success("Tags查询成功", searchService.searchTags(searchFortags));
    }

    @Api
    @PostMapping("/labels")
    public Result labels(
            @RequestBody SearchForLabels searchForlabels
    ) {
        log.info("Search Labels");
        return Result.success("Labels查询成功", searchService.searchLabels(searchForlabels));
    }

    @Api
    @PostMapping("/sections")
    public Result sections(
            @RequestBody SearchForSections searchForSections
    ) {
        log.info("Search Sections");
        return Result.success("Sections查询成功", searchService.searchSections(searchForSections));
    }
}
