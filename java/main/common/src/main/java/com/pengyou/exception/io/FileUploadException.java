package com.pengyou.exception.io;

import com.pengyou.constant.FileConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class FileUploadException extends BaseException {
    public FileUploadException() {
        super(FileConstant.FILE_UPLOAD_FAILURE);
    }

    public FileUploadException(String msg) {
        super(msg);
    }
}
