package com.pengyou.exception.common;

import com.pengyou.constant.CommonConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class InputInvalidException extends BaseException {

    public InputInvalidException() {
        super(CommonConstant.INPUT_INVALID);
    }

    public InputInvalidException(String msg) {
        super(msg);
    }
}
