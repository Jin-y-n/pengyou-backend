package com.pengyou.service.impl;

import com.pengyou.exception.BaseException;
import com.pengyou.model.dto.postlabel.LabelForQuery;
import com.pengyou.model.dto.postlabel.LabelForQueryView;
import com.pengyou.model.dto.postsection.SectionForQuery;
import com.pengyou.model.dto.postsection.SectionForQueryView;
import com.pengyou.model.dto.tag.TagForQuery;
import com.pengyou.model.dto.tag.TagForQueryView;
import com.pengyou.model.entity.PostLabelTable;
import com.pengyou.model.entity.PostSectionTable;
import com.pengyou.model.entity.TagTable;
import com.pengyou.service.QueryService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.Page;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service
@RequiredArgsConstructor
public class QueryImpl implements QueryService {
    private final JSqlClient sqlClient;

    @Override
    public Page<TagForQueryView> queryTag(TagForQuery tagForQuery) {
        Page<TagForQueryView> page = sqlClient
                .createQuery(TagTable.$)
                .where(tagForQuery)
                .select(
                        TagTable.$.fetch(TagForQueryView.class)
                )
                .fetchPage(tagForQuery.getPageIndex(), tagForQuery.getPageSize());


        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Post不存在");
        }
        return page;
    }

    @Override
    public Page<LabelForQueryView> queryLabel(LabelForQuery labelForQuery) {
        Page<LabelForQueryView> page = sqlClient
                .createQuery(PostLabelTable.$)
                .where(labelForQuery)
                .select(
                        PostLabelTable.$.fetch(LabelForQueryView.class)
                )
                .fetchPage(labelForQuery.getPageIndex(), labelForQuery.getPageSize());

        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Post不存在");
        }
        return page;
    }

    @Override
    public Page<SectionForQueryView> querySection(SectionForQuery sectionForQuery) {
        Page<SectionForQueryView> page = sqlClient
                .createQuery(PostSectionTable.$)
                .where(sectionForQuery)
                .select(
                        PostSectionTable.$.fetch(SectionForQueryView.class)
                )
                .fetchPage(sectionForQuery.getPageIndex(), sectionForQuery.getPageSize());


        if (page.getTotalRowCount() == 0) {
            throw new BaseException("Post不存在");
        }
        return page;
    }
}
