package com.pengyou.model.entity;

import org.babyfish.jimmer.sql.*;

import org.jetbrains.annotations.Nullable;

import java.time.LocalDateTime;
import java.util.List;

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

    @ManyToMany
    @JoinTable(
            name = "user_tag_mapping",
            joinColumnName = "user_id",
            inverseJoinColumnName = "tag_id"
    )
    List<Tag> tags();

    @OneToOne(mappedBy = "user")
    @Nullable
    UserProfile profile();
}

