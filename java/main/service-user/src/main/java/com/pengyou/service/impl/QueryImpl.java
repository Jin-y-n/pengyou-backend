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
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
@RequiredArgsConstructor
public class QueryImpl implements QueryService {
    private final JSqlClient sqlClient;
    @Override
    public TagForQueryView queryTag(TagForQuery tagForQuery) {
        List<TagForQueryView> execute = sqlClient
                .createQuery(TagTable.$)
                .where(tagForQuery)
                .select(
                        TagTable.$.fetch(TagForQueryView.class)
                )
                .execute();

        Optional<TagForQueryView> first = execute.stream().findFirst();

        if (first.isEmpty()) {
            throw new BaseException("Tag不存在");
        }
        return first.get();
    }

    @Override
    public LabelForQueryView queryLabel(LabelForQuery labelForQuery) {
        List<LabelForQueryView> execute = sqlClient
                .createQuery(PostLabelTable.$)
                .where(labelForQuery)
                .select(
                        PostLabelTable.$.fetch(LabelForQueryView.class)
                )
                .execute();

        Optional<LabelForQueryView> first = execute.stream().findFirst();

        if (first.isEmpty()) {
            throw new BaseException("Label不存在");
        }
        return first.get();
    }

    @Override
    public SectionForQueryView querySection(SectionForQuery sectionForQuery) {
        List<SectionForQueryView> execute = sqlClient
                .createQuery(PostSectionTable.$)
                .where(sectionForQuery)
                .select(
                        PostSectionTable.$.fetch(SectionForQueryView.class)
                )
                .execute();

        Optional<SectionForQueryView> first = execute.stream().findFirst();

        if (first.isEmpty()) {
            throw new BaseException("Section不存在");
        }
        return first.get();
    }
}
