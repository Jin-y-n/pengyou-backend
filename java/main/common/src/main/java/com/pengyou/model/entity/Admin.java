package com.pengyou.model.entity;

import com.pengyou.cnoverter.PasswordConverter;
import org.babyfish.jimmer.jackson.JsonConverter;
import org.babyfish.jimmer.sql.*;


import org.jetbrains.annotations.NotNull;
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

    @NotNull
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

    Short role(); // 1：超级管理员 2：普通管理员

    Short modifiedByRoot(); // 1:被超级修改 0：被普通修改
}

