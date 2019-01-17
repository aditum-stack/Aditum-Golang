import com.ctrip.framework.apollo.spring.annotation.EnableApolloConfig;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.sidecar.EnableSidecar;

/**
 * 监听30022端口
 */
@SpringBootApplication
@EnableSidecar
@EnableApolloConfig
public class EmailSidecarApplication {
    public static void main(String[] args) {
        SpringApplication.run(EmailSidecarApplication.class, args);
    }
}
