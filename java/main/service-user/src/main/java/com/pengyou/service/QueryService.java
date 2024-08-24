package com.pengyou.service;

import com.pengyou.model.dto.postlabel.LabelForQuery;
import com.pengyou.model.dto.postlabel.LabelForQueryView;
import com.pengyou.model.dto.postsection.SectionForQuery;
import com.pengyou.model.dto.postsection.SectionForQueryView;
import com.pengyou.model.dto.tag.TagForQuery;
import com.pengyou.model.dto.tag.TagForQueryView;
import org.babyfish.jimmer.Page;

public interface QueryService {
    Page<TagForQueryView> queryTag(TagForQuery tagForQuery);
    Page<LabelForQueryView> queryLabel(LabelForQuery labelForQuery);
    Page<SectionForQueryView> querySection(SectionForQuery sectionForQuery);
}
