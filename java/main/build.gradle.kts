plugins {
    id("java")
    id("org.springframework.boot") version "3.2.5"
    id("io.spring.dependency-management") version "1.1.4"
}

group = "com.pengyou"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

val feignCoreVersion = "13.3"
val springVersion = "3.2.5"
val springCloudLoadBalancerVersion = "4.1.3"
val jimmerVersion = "0.8.147"
val druidVersion = "1.2.21"
val jwtVersion = "0.9.1"
val jaxbVersion = "4.0.2"
val jaxbRuntimeVersion = "2.3.1"
val lombokVersion = "1.18.32"
val aspectVersion = "1.9.21"
val junitPlatformLauncherVersion = "1.10.2"
val javaMailVersion = "1.6.2"
val commonsPoolVersion = "2.11.1"
val nacosVersion = "2023.0.1.0"
val aliOOSVersion = "3.17.4"
val jsrVersion = "2.17.2"
val prometheusVersion = "1.13.3"
val commonsPool2Version = "2.12.0"
val redissonVersion = "3.34.1"

allprojects {
    apply {
        plugin("java")
        plugin("org.springframework.boot")
        plugin("io.spring.dependency-management")
    }

    dependencies {

        implementation("org.springframework.boot:spring-boot-starter-web")

        // Data
        implementation("org.springframework.boot:spring-boot-starter-data-redis")
        implementation("org.springframework.boot:spring-boot-starter-data-redis-reactive")
        implementation("org.babyfish.jimmer:jimmer-spring-boot-starter:$jimmerVersion")
        implementation("org.babyfish.jimmer:jimmer-sql:$jimmerVersion")
        implementation("com.alibaba:druid:$druidVersion")
//        implementation("com.aliyun.oss:aliyun-sdk-oss:$aliOOSVersion")
        implementation("org.apache.commons:commons-pool2:$commonsPool2Version")

        runtimeOnly("com.mysql:mysql-connector-j")


        // Utils
        implementation("com.fasterxml.jackson.datatype:jackson-datatype-jsr310:$jsrVersion")
        implementation("io.jsonwebtoken:jjwt:$jwtVersion")
        implementation("org.springframework.boot:spring-boot-starter-mail")
//        implementation("io.github.openfeign:feign-core:$feignCoreVersion")
//        implementation("com.alibaba.cloud:spring-cloud-starter-alibaba-nacos-discovery:$nacosVersion")
//        implementation("org.springframework.cloud:spring-cloud-starter-loadbalancer:$springCloudLoadBalancerVersion")
        implementation("io.micrometer:micrometer-registry-prometheus:$prometheusVersion")
        implementation("org.redisson:redisson-spring-boot-starter:$redissonVersion")


        compileOnly("org.projectlombok:lombok")
        annotationProcessor("org.projectlombok:lombok")
        annotationProcessor("org.babyfish.jimmer:jimmer-apt:$jimmerVersion")


        // test
        testImplementation("org.springframework.boot:spring-boot-starter-test")
        testImplementation(platform("org.junit:junit-bom:5.10.0"))
        testImplementation("org.junit.jupiter:junit-jupiter")
        testRuntimeOnly("org.junit.platform:junit-platform-launcher")
    }

    tasks.withType<JavaCompile> {
        options.encoding = "UTF-8"
    }

}

tasks.test {
    useJUnitPlatform()
}