package com.pengyou.exception.account;

import com.pengyou.constant.AccountConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class AccountPermissionException extends BaseException {
    public AccountPermissionException()
    {
        super(AccountConstant.INSUFFICIENT_AUTHORITY);
    }

    public AccountPermissionException(String message)
    {
        super(message);
    }
}
