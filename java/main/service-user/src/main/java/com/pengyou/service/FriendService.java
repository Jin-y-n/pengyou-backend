package com.pengyou.service;

import com.pengyou.model.dto.Friend.FriendForAccept;
import com.pengyou.model.dto.Friend.FriendForAdd;
import com.pengyou.model.dto.Friend.FriendForDelete;
import com.pengyou.model.dto.Friend.FriendForUpgrade;
import com.pengyou.model.dto.Socialaccount.AccountForAdd;
import com.pengyou.model.dto.Socialaccount.AccountForDelete;

public interface FriendService {
    void add(FriendForAdd friend);
    void delete(FriendForDelete friend);
    void accept(FriendForAccept friend);
    void upgrade(FriendForUpgrade friend);
    void addSocialAccount(AccountForAdd account);
    void deleteSocialAccount(AccountForDelete account);
}
