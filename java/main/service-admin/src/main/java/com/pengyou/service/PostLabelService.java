package com.pengyou.service;


import com.pengyou.model.dto.postlabel.PostLabelForAdd;
import com.pengyou.model.dto.postlabel.PostLabelForDelete;
import com.pengyou.model.dto.postlabel.PostLabelForQuery;
import com.pengyou.model.dto.postlabel.PostLabelForQueryView;
import com.pengyou.model.dto.postsection.PostSectionForQuery;
import com.pengyou.model.dto.postsection.PostSectionForQueryView;
import org.babyfish.jimmer.Page;

public interface PostLabelService {
    void add(PostLabelForAdd postLabelForAdd);
    void delete(PostLabelForDelete postLabelForDelete);
    Page<PostLabelForQueryView> query(PostLabelForQuery postLabelForQuery);

}
