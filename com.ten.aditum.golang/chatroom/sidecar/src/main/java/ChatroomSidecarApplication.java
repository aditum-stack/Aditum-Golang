import com.ctrip.framework.apollo.spring.annotation.EnableApolloConfig;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.sidecar.EnableSidecar;

/**
 * 监听30021端口
 */
@SpringBootApplication
@EnableSidecar
@EnableApolloConfig
public class ChatroomSidecarApplication {
    public static void main(String[] args) {
        SpringApplication.run(ChatroomSidecarApplication.class, args);
    }
}
