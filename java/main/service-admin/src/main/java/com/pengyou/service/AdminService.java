package com.pengyou.service;


import com.pengyou.model.dto.admin.*;
import org.babyfish.jimmer.Page;



public interface AdminService {
    void register(AdminForRegister adminForRegister);
    String verify(AdminForVerify adminForVerify);
    void delete(AdminForDelete adminForDelete);
    void update(AdminForUpdate adminForUpdate);
    Page<AdminForView> query(AdminForQuery adminForQuery);
    AdminForLoginView login(AdminForLogin adminForLogin);
    void logout(AdminForLogout adminForLogout);
}
