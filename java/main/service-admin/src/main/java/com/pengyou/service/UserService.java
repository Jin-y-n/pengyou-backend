package com.pengyou.service;


import com.pengyou.model.dto.user.AdminUserForQuery;
import com.pengyou.model.dto.user.AdminUserForQueryView;
import org.babyfish.jimmer.Page;

public interface UserService {
//    void update(AdminUserForUpdate adminUserForUpdate);
    Page<AdminUserForQueryView> query(AdminUserForQuery adminUserForQuery);
}
