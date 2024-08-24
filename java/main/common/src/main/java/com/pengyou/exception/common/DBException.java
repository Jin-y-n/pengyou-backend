package com.pengyou.exception.common;

import com.pengyou.constant.CommonConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class DBException extends BaseException {
    public DBException() {
        super(CommonConstant.DB_EXCEPTION);
    }

    public DBException(String msg) {
        super(msg);
    }
}
