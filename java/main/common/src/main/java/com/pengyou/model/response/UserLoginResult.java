package com.pengyou.model.response;

import com.pengyou.model.entity.Admin;
import com.pengyou.model.entity.User;
import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class UserLoginResult {
    private User user;
    private String jwt;
}