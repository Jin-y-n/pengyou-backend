package com.pengyou.model.entity;

import org.babyfish.jimmer.sql.Entity;
import org.babyfish.jimmer.sql.Id;
import org.babyfish.jimmer.sql.GeneratedValue;
import org.babyfish.jimmer.sql.Key;

import org.babyfish.jimmer.sql.GenerationType;
import org.jetbrains.annotations.Nullable;

import java.time.LocalDateTime;

/**
 * Entity for table "user"
 */
@Entity
public interface User {

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
    LocalDateTime loginTime();

    @Nullable
    LocalDateTime createdAt();

    @Nullable
    LocalDateTime modifiedAt();

    @Nullable
    LocalDateTime deleteAt();

    @Nullable
    Short status();

    @Nullable
    LocalDateTime heartBeatTime();

    @Nullable
    String clientIp();

    @Nullable
    Short isLogout();

    @Nullable
    LocalDateTime logOutTime();

    @Nullable
    String deviceInfo();

    @Nullable
    Long createdPerson();

    @Nullable
    Long modifiedPerson();

    Short modifiedByAdmin();
}

