package com.pengyou.model.entity;

import org.babyfish.jimmer.sql.*;
import org.jetbrains.annotations.Nullable;

@Entity
public interface SocialAccount {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    long id();

    @IdView
    long userId();

    @Nullable
    String platform();

    @Nullable
    String link();

    @OneToOne
    @JoinColumn(name = "user_id")
    User user();
}
