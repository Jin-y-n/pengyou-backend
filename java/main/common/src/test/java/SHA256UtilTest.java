/*
    @Author: Napbad
    @Version: 0.1    
    @Date: 8/25/24
    @Description: 

*/


import com.pengyou.util.security.SHA256Encryption;
import org.junit.jupiter.api.Test;

public class SHA256UtilTest {


    @Test
    void test() {
        System.out.println(SHA256Encryption.getSHA256("123456"));
    }

}
