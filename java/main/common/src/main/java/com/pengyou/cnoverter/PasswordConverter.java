package com.pengyou.cnoverter;

import com.pengyou.constant.AccountConstant;
import org.babyfish.jimmer.jackson.Converter;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class PasswordConverter implements Converter<String, String> {
    @Override
    public String output(String value) {
        return AccountConstant.ACCOUNT_PASSWORD_ENCODE;
    }
}
