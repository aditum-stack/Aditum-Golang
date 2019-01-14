package com.ten.aditum.go.sidecar;


import com.ctrip.framework.apollo.spring.annotation.EnableApolloConfig;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.sidecar.EnableSidecar;

@SpringBootApplication
@EnableSidecar
@EnableApolloConfig
public class GoSidecarApplication {
    public static void main(String[] args) {
        SpringApplication.run(GoSidecarApplication.class, args);
    }
}
