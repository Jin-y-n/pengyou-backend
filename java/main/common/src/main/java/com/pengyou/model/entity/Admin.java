package com.pengyou.model.entity;

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

    long createdPerson();

    long modifiedPerson();

    @Nullable
    @LogicalDeleted("now")
    LocalDateTime deleteAt();

    short role(); // 1：超级管理员 2：普通管理员

    short modifiedByRoot(); // 1:被超级修改 0：被普通修改
}

