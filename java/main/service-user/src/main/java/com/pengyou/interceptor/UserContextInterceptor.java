package com.pengyou.interceptor;

import com.pengyou.config.properties.AuthExcludePathProperties;
import com.pengyou.config.properties.JwtProperties;
import com.pengyou.constant.JwtClaimsConstant;
import com.pengyou.model.entity.User;
import com.pengyou.util.UserContext;
import com.pengyou.util.security.JwtUtil;
import io.jsonwebtoken.Claims;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.extern.slf4j.Slf4j;
import org.jetbrains.annotations.NotNull;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.util.AntPathMatcher;
import org.springframework.util.StringUtils;
import org.springframework.web.servlet.HandlerInterceptor;

import java.io.IOException;


/**
 * jwt令牌校验的拦截器
 */
@Slf4j
@Component
public class UserContextInterceptor implements HandlerInterceptor {

    private final AuthExcludePathProperties authExcludePathProperties;
    private final JwtProperties jwtProperties;
    private final AntPathMatcher antPathMatcher = new AntPathMatcher();

    public UserContextInterceptor(AuthExcludePathProperties authExcludePathProperties, JwtProperties jwtProperties) {
        this.authExcludePathProperties = authExcludePathProperties;
        this.jwtProperties = jwtProperties;
    }

    /**
     * 校验jwt
     *
     * @param request  请求对象，用于从请求头中获取JWT令牌
     * @param response 响应对象，用于在令牌校验失败时设置HTTP状态码
     * @param handler  当前拦截的处理器对象，用于判断是否为Controller方法
     * @return 是否允许继续执行后续处理器（即Controller方法）
     */
    // TODO Authorization update
    @Override
    public boolean preHandle(
            @NotNull HttpServletRequest request,
            @NotNull HttpServletResponse response,
            @NotNull Object handler) {

//        String userId = request.getHeader("userId");
//        if (StringUtils.hasLength(userId)) {
//            return false;
//        }
//
//        if (isExcluded(request.getServletPath())) {
//            UserContext.setUserId(0);
//            return true;
//        }
//
//        UserContext.setUserId(Integer.valueOf(userId));
        return true;
//        UserContext.setUserId(1);
//        return true;

//        String token = request.getHeader(jwtProperties.getTokenName());
//
//        if (isExcluded(request.getServletPath())) {
//            UserContext.setUserId(0);
//
//            return true;
//        }
//
//        if (!StringUtils.hasLength(token)) {
//            try {
//                response.sendError(HttpServletResponse.SC_UNAUTHORIZED, "未登录");
//            } catch (IOException e) {
//                throw new RuntimeException(e);
//            }
//            return false;
//        }
//
//        try {
//            log.info("校验token:{}", token);
//            Claims claims = JwtUtil.parseJWT(jwtProperties.getSecretKey(), token);
//            String s = claims.get(JwtClaimsConstant.ID).toString();
//            long empId = Long.parseLong(s);
//            log.info("登录用户:{}", empId);
//            // 3、令牌校验通过，放行请求，继续执行后续处理器（Controller方法）
//
//           UserContext.setUserId(Integer.valueOf(s));
//           return true;
//        } catch (Exception ex) {
//            // 4、令牌校验不通过，响应401 Unauthorized状态码，并阻止执行后续处理器
////            exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
////
////            if (ex.getMessage().contains("JWT expired at")) {
////                return exchange.getResponse().setComplete();
////            }
////
////            return exchange.getResponse().setComplete();
//            log.error("校验token失败:{}", ex.getMessage());
//            try {
//                response.sendError(HttpServletResponse.SC_UNAUTHORIZED, "未登录");
//            } catch (IOException e) {
//                throw new RuntimeException(e);
//            }
//            return false;
//        }
    }

    @Override
    public void afterCompletion(@NotNull HttpServletRequest request,
                                @NotNull HttpServletResponse response,
                                @NotNull Object handler, Exception ex) throws Exception {
        UserContext.remove();
    }

    @NotNull
    private Boolean isExcluded(String path) {
        return authExcludePathProperties
                .getExcludePaths()
                .stream()
                .anyMatch(excludePath -> antPathMatcher.match(excludePath, path));
    }
}