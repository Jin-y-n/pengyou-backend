package com.pengyou.service;

import com.pengyou.model.dto.postlabel.LabelForQuery;
import com.pengyou.model.dto.postlabel.LabelForQueryView;
import com.pengyou.model.dto.postsection.SectionForQuery;
import com.pengyou.model.dto.postsection.SectionForQueryView;
import com.pengyou.model.dto.tag.TagForQuery;
import com.pengyou.model.dto.tag.TagForQueryView;

public interface QueryService {
    TagForQueryView queryTag(TagForQuery tagForQuery);
    LabelForQueryView queryLabel(LabelForQuery labelForQuery);
    SectionForQueryView querySection(SectionForQuery sectionForQuery);
}
