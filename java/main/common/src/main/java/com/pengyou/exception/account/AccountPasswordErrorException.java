package com.pengyou.exception.account;

import com.pengyou.constant.AccountConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class AccountPasswordErrorException extends BaseException {
    public AccountPasswordErrorException() {
        super(AccountConstant.ACCOUNT_PASSWORD_WRONG);
    }

    public AccountPasswordErrorException(String message) {
        super(message);
    }
}
