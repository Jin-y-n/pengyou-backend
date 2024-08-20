package com.pengyou.service;


import com.pengyou.model.dto.postlabel.PostLabelForAdd;
import com.pengyou.model.dto.postlabel.PostLabelForDelete;

public interface PostLabelService {
    void add(PostLabelForAdd postLabelForAdd);
    void delete(PostLabelForDelete postLabelForDelete);
}
