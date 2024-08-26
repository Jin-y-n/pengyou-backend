package com.pengyou.model.entity;

import com.esotericsoftware.kryo.util.Null;
import org.babyfish.jimmer.sql.*;


import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.time.LocalDateTime;
import java.util.List;

/**
 * Entity for table "post"
 */
@Entity
public interface Post {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY
    )
    long id();

    @IdView
    Long authorId();

    @Nullable
    String title();

    @Nullable
    String content();

    @Nullable
    LocalDateTime createdAt();

    @Nullable
    LocalDateTime modifiedAt();

    @Nullable
    Short status();

    @Nullable
    Long createdPerson();

    @Nullable
    Long modifiedPerson();

    @Nullable
    @LogicalDeleted("now")
    LocalDateTime deleteAt();

    @ManyToMany
    @JoinTable(
            name = "post_label_mapping",
            joinColumnName = "post_id",
            inverseJoinColumnName = "label_id"
    )
    List<PostLabel> labels();

    @ManyToMany
    @JoinTable(
            name = "post_section_mapping",
            joinColumnName = "post_id",
            inverseJoinColumnName = "section_id"
    )
    List<PostSection> sections();

    @Nullable
    @ManyToOne(inputNotNull = true)
    @JoinColumn(name = "author")
    @OnDissociate(DissociateAction.SET_NULL)
    User author();
}

