package com.pengyou.exception.account;

import com.pengyou.constant.AccountConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 * 密码错误异常
 */
public class PasswordErrorException extends BaseException {

    public PasswordErrorException() {
        super(AccountConstant.ACCOUNT_PASSWORD_WRONG);
    }

    public PasswordErrorException(String msg) {
        super(msg);
    }

}
