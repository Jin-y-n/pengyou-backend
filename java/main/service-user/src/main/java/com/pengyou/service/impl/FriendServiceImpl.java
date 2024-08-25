package com.pengyou.service.impl;

import com.pengyou.model.dto.Friend.FriendForAccept;
import com.pengyou.model.dto.Friend.FriendForAdd;
import com.pengyou.model.dto.Friend.FriendForDelete;
import com.pengyou.model.dto.Friend.FriendForUpgrade;
import com.pengyou.model.dto.Socialaccount.AccountForAdd;
import com.pengyou.model.dto.Socialaccount.AccountForDelete;
import com.pengyou.model.entity.SocialAccount;
import com.pengyou.model.entity.UserFriend;
import com.pengyou.service.FriendService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class FriendServiceImpl implements FriendService {

    private final JSqlClient jSqlClient;

    @Override
    public void add(FriendForAdd friend) {
        jSqlClient.save(friend);
    }

    @Override
    public void delete(FriendForDelete friend){
        jSqlClient.deleteById(UserFriend.class, friend.getId());
    };

    @Override
    public void accept(FriendForAccept friend) {
        jSqlClient.save(friend);
    }

    @Override
    public void upgrade(FriendForUpgrade friend) {
        jSqlClient.save(friend);
    }

    @Override
    public void addSocialAccount(AccountForAdd account) {
        jSqlClient.save(account);
    }

    @Override
    public void deleteSocialAccount(AccountForDelete account) {
        jSqlClient.deleteById(SocialAccount.class,account.getId());
    }

}
