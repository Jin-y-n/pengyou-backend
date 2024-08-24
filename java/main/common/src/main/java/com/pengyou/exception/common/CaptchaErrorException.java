package com.pengyou.exception.common;

import com.pengyou.constant.VerifyConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class CaptchaErrorException extends BaseException {
    public CaptchaErrorException()
    {
        super(VerifyConstant.CAPTCHA_ERROR);
    }

    public CaptchaErrorException(String message)
    {
        super(message);
    }
}
