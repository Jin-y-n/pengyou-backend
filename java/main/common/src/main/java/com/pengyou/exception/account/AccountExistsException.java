package com.pengyou.exception.account;

import com.pengyou.constant.AccountConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class AccountExistsException extends BaseException {
    public AccountExistsException()
    {
        super(AccountConstant.ACCOUNT_EXISTS);
    }

    public AccountExistsException(String message)
    {
        super(message);
    }
}
