import java.util.Arrays;
import java.util.Collection;

import org.junit.Test;

import static org.junit.Assert.*;

import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;
import org.junit.runners.Parameterized.Parameters;

/**
 * @author Mariana
 */
@RunWith(Parameterized.class)
public class CustomMathTest {
    @Parameters
    public static Collection sumValues() {
        return Arrays.asList(new Object[][]{
                {10, 5, 2},
                {20, 5, 4},
                {24, 6, 4}});
    }

    int x, y, sumResult;

    public CustomMathTest(int x, int y, int sumResult) {
        this.x = x;
        this.y = y;
        this.sumResult = sumResult;
    }

    @Test
    public void testDivisionByZero() {
        int expResult = sumResult;
        assertTrue(y > 0);
        assertFalse(y <= 0);
        try {
            int result = CustomMath.division(x, y);
            assertEquals(expResult, result);
            if (y == 0) fail("Деление на ноли не создает исключителънои ситуации");
        } catch (IllegalArgumentException e) {
            if (y != 0) fail("Генерация исключения при ненулевом знаменателе");
        }
    }
}