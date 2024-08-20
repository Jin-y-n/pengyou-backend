package com.pengyou.service;


import com.pengyou.model.dto.admin.*;

import java.util.List;


public interface AdminService {
    void register(AdminForRegister adminForRegister);
    void verify(AdminForVerify adminForVerify);
    void delete(AdminForDelete adminForDelete);
    void update(AdminForUpdate adminForUpdate);
    List<AdminForView> query(AdminForQuery adminForQuery);
    AdminForLoginView login(AdminForLogin adminForLogin);
    void logout(AdminForLogout adminForLogout);

}
