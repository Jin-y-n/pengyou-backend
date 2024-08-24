package com.pengyou.exception.account;

import com.pengyou.constant.AccountConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class AccountNameErrorException extends BaseException {
    public AccountNameErrorException()
    {
        super(AccountConstant.ACCOUNT_NAME_ERROR);
    }

    public AccountNameErrorException(String msg)
    {
        super(msg);
    }
}
