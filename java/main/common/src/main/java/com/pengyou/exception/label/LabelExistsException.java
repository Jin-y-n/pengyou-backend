package com.pengyou.exception.label;

import com.pengyou.constant.LabelConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class LabelExistsException extends BaseException {

    public LabelExistsException() {
        super(LabelConstant.LABEL_EXISTS);
    }

    public LabelExistsException(String msg) {
        super(msg);
    }
}
