package com.pengyou.exception.account;

import com.pengyou.constant.AccountConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class UserNotLoginException extends BaseException {

    public UserNotLoginException() {
        super(AccountConstant.ACCOUNT_NOT_LOGIN);
    }

    public UserNotLoginException(String msg) {
        super(msg);
    }

}
