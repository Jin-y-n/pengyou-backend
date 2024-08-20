package com.pengyou.model.entity;

import com.esotericsoftware.kryo.util.Null;
import org.babyfish.jimmer.sql.*;


import org.jetbrains.annotations.Nullable;

import java.time.LocalDateTime;

/**
 * Entity for table "admin"
 */
@Entity
public interface Admin {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY
    )
    long id();

    @Key
    String username();

    String password();

    @Nullable
    String email();

    @Nullable
    String phone();

    @Nullable
    LocalDateTime createdTime();

    @Nullable
    LocalDateTime modifiedTime();

    @Nullable
    Long createdPerson();

    @Nullable
    Long modifiedPerson();

    @Nullable
    @LogicalDeleted("now")
    LocalDateTime deleteAt();

    @Nullable
    Short role();

    @Nullable
    Short modifiedByRoot();
}

