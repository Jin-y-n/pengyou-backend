package com.pengyou.model.entity;

import org.babyfish.jimmer.sql.*;


import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;

/**
 * Entity for table "post_section"
 */
@Entity
public interface PostSection {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY
    )
    long id();

    @NotNull
    String section();

    @Nullable
    String description();

    @ManyToMany(mappedBy = "sections")
    List<Post> posts();
}

