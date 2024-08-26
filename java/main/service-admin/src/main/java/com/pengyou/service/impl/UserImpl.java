package com.pengyou.service.impl;

import com.pengyou.exception.BaseException;
import com.pengyou.model.dto.user.AdminUserForQuery;
import com.pengyou.model.dto.user.AdminUserForQueryView;
import com.pengyou.model.entity.UserTable;
import com.pengyou.service.UserService;
import lombok.RequiredArgsConstructor;
import org.babyfish.jimmer.Page;
import org.babyfish.jimmer.sql.JSqlClient;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class UserImpl implements UserService {
    private final JSqlClient sqlClient;
    private final UserTable userTable = UserTable.$;

    @Override
    public Page<AdminUserForQueryView> query(AdminUserForQuery adminUserForQuery) {
        Page<AdminUserForQueryView> page = sqlClient
                .createQuery(userTable)
                .where(adminUserForQuery)
                .select(
                        userTable.fetch(AdminUserForQueryView.class)
                )
                .fetchPage(adminUserForQuery.getPageIndex() == null ? 0 : adminUserForQuery.getPageIndex(),
                        adminUserForQuery.getPageSize() == null? 10 : adminUserForQuery.getPageSize());

        if(page.getTotalRowCount() == 0){
            throw new BaseException("User未找到");
        }
        return page;
    }
}
