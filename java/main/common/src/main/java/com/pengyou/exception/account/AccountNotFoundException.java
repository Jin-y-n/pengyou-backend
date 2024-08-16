package com.pengyou.exception.account;

import com.pengyou.constant.AccountConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 * 账号不存在异常
 */
public class AccountNotFoundException extends BaseException {

    public AccountNotFoundException() {
        super(AccountConstant.ACCOUNT_NOT_FOUND);
    }

    public AccountNotFoundException(String msg) {
        super(msg);
    }
}
