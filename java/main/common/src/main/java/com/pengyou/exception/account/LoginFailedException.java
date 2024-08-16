package com.pengyou.exception.account;

import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 * 登录失败
 */
public class LoginFailedException extends BaseException {
    public LoginFailedException(String msg){
        super(msg);
    }
}
