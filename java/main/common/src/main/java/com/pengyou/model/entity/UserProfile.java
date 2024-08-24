package com.pengyou.model.entity;

import org.babyfish.jimmer.sql.*;


import org.jetbrains.annotations.Nullable;

import java.time.LocalDate;
import java.time.LocalDateTime;

/**
 * Entity for table "user_profile"
 */
@Entity
public interface UserProfile {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY
    )
    long id();

    @IdView
    long userId();

    @Nullable
    String displayName();

    @Nullable
    String avatarId();

    @Nullable
    String bio();

    @Nullable
    Short gender();

    @Nullable
    LocalDate birthday();

    @Nullable
    String location();

    @Nullable
    String occupation();

    @Nullable
    String education();

    @Nullable
    String school();

    @Nullable
    String major();

    @Nullable
    String company();

    @Nullable
    String position();

    @Nullable
    String website();

    @Nullable
    LocalDateTime createdAt();

    @Nullable
    LocalDateTime modifiedAt();

    @Nullable
    LocalDateTime deleteAt();

    @Nullable
    Long createdPerson();

    @Nullable
    Long modifiedPerson();

    Short modifiedByAdmin();

    @OneToOne
    @JoinColumn(name = "user_id")
    User user();
}

