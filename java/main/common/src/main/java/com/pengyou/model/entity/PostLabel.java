package com.pengyou.model.entity;

import org.babyfish.jimmer.sql.*;


import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;

/**
 * Entity for table "post_label"
 */
@Entity
public interface PostLabel {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY
    )
    long id();

    @NotNull
    String label();

    @Nullable
    String description();

    @ManyToMany(mappedBy = "labels")
    List<Post> posts();
}

