package com.pengyou.controller;

/*
 * Author: Napbad
 * Version: 1.0
 */

import com.pengyou.constant.FileConstant;
import com.pengyou.exception.io.FileDownloadException;
import com.pengyou.exception.io.FileUploadException;
import com.pengyou.model.Result;
import jakarta.servlet.ServletOutputStream;
import jakarta.servlet.http.HttpServletResponse;
import lombok.extern.slf4j.Slf4j;
import org.babyfish.jimmer.client.meta.Api;
import org.springframework.core.io.FileSystemResource;
import org.springframework.core.io.Resource;
import org.springframework.scheduling.annotation.Async;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.io.*;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Objects;
import java.util.UUID;
import java.util.concurrent.CompletableFuture;

import static com.pengyou.constant.FileConstant.FILE_DIRECTORY;

@Api
@Slf4j
@RestController
@RequestMapping("/file")
public class FileController {

//    private final AliOssUtil aliOssUtil;
//
//    // 构造函数注入OutSideProperty，用于获取文件存储路径和服务器地址
//    public FileController(AliOssUtil aliOssUtil) {
//        this.aliOssUtil = aliOssUtil;
//    }

    /**
     * 异步上传file
     *
     * @param file 接收上传的文件
     * @return CompletableFuture<Result> 异步返回文件上传结果，其中包含文件的访问URL
     */
    @Api // 标识异步执行的API接口
    @Async // 异步执行
    @PostMapping("/upload")
    public CompletableFuture<Result> upload(@RequestParam("file") MultipartFile file) {
        // 获取文件原始名称
        String filename = file.getOriginalFilename();

        // 检查文件名格式是否合法
//        if (!FileUtil.checkFileNameFormat(Objects.requireNonNull(filename))) {
//            throw new FileUploadException(FileConstant.FILE_NOT_SUPPORTED);
//        }

        // 检查文件是否为空
        if (file.isEmpty()) {
            throw new FileUploadException(FileConstant.FILE_IS_NULL);
        }

        try {
            UUID uuid = UUID.randomUUID();

            if (filename == null) {
                throw new FileUploadException(FileConstant.FILE_NOT_FOUND);
            }

            int index = filename.lastIndexOf('.');

            if (index == -1) {
                throw new FileUploadException();
            }

            filename = uuid + filename.substring(index);


//            String path = aliOssUtil.upload(file.getBytes(), filename);

            String path = FILE_DIRECTORY + filename;

            InputStream inputStream = file.getInputStream();
            File file1 = new File(path);
            if (!file1.exists()) {
                boolean b = file1.createNewFile();
                if (!b) {
                    throw new FileUploadException();
                }
            }
            FileOutputStream outputStream = new FileOutputStream(file1);

            byte[] buf = new byte[4096];
            int read = inputStream.read(buf, 0, buf.length);
            while (read > buf.length) {
                read = inputStream.read(buf, 0, buf.length);
                outputStream.write(buf, 0, read);
            }

            // 记录文件上传日志
            log.info("文件上传成功: {}", path);
            // 返回文件的访问URL
            return CompletableFuture.completedFuture(Result.success(filename));
        } catch (IOException e) {
            e.printStackTrace();
            throw new FileUploadException();
        }
    }


    @Api
    @GetMapping("/files/{filename}")
    public void getFile(
            @PathVariable String filename,
            HttpServletResponse response) throws IOException {
        File file = new File(FileConstant.FILE_DIRECTORY + filename);

        if (!file.exists()) {
            throw new FileDownloadException(FileConstant.FILE_DOWNLOAD_FAILURE);
        }

        try {

            response.setCharacterEncoding("UTF-8");
            //Content-Disposition的作用：告知浏览器以何种方式显示响应返回的文件，用浏览器打开还是以附件的形式下载到本地保存
            //attachment表示以附件方式下载   inline表示在线打开   "Content-Disposition: inline; filename=文件名.mp3"
            // filename表示文件的默认名称，因为网络传输只支持URL编码的相关支付，因此需要将文件名URL编码后进行传输,前端收到后需要反编码才能获取到真正的名称
            response.addHeader("Content-Disposition", "attachment;filename=" + filename);
            // 告知浏览器文件的大小
            response.addHeader("Content-Length", "" + file.length());
            response.setContentType("application/octet-stream");

            FileInputStream inputStream = new FileInputStream(file);

            byte[] bytes = new byte[4096];
            ServletOutputStream outputStream = response.getOutputStream();
            while (inputStream.read(bytes) != -1) {
                outputStream.write(bytes);
            }
            outputStream.flush();
        } catch (Exception e) {
            e.printStackTrace();
            throw new FileDownloadException(FileConstant.FILE_DOWNLOAD_FAILURE);
        }
    }
}
