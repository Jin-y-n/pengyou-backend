package com.pengyou.exception.io;

import com.pengyou.constant.FileConstant;
import com.pengyou.exception.BaseException;

/*
 * Author: Napbad
 * Version: 1.0
 */
public class FileDownloadException extends BaseException {
    public FileDownloadException() {
        super(FileConstant.FILE_DOWNLOAD_FAILURE);
    }
    public FileDownloadException(String msg) {
        super(msg);
    }
}
