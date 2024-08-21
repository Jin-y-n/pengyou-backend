package com.pengyou.model.entity;

import org.babyfish.jimmer.sql.*;


import org.jetbrains.annotations.Nullable;

import java.util.List;

/**
 * Entity for table "tag"
 */
@Entity
public interface Tag {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY
    )
    long id();

    @Key
    String name();

    @Nullable
    String description();

    @ManyToMany(mappedBy = "tags")
    List<User> user();

}

