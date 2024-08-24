package com.pengyou.model.response;

import com.pengyou.model.entity.Admin;
import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class AdminLoginResult {
    private Admin admin;
    private String jwt;
}
