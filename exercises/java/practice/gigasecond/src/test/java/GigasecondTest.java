import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

import java.time.LocalDateTime;
import java.time.Month;

public class GigasecondTest {

  @Test
  @DisplayName("Fecha")
  public void testDate() {
    LocalDateTime birthDate = LocalDateTime.of(2011, Month.APRIL, 25, 0, 0);
    assertEquals(LocalDateTime.of(2043, Month.JANUARY, 1, 1, 46, 40), Gigasecond.after(birthDate));
  }

  @Test
  @DisplayName("Otra fecha")
  public void testAnotherDate() {
    LocalDateTime birthDate = LocalDateTime.of(1977, Month.JUNE, 13, 0, 0);
    assertEquals(LocalDateTime.of(2009, Month.FEBRUARY, 19, 1, 46, 40), Gigasecond.after(birthDate));
  }

  @Test
  @DisplayName("Tercera fecha")
  public void testThirdDate() {
    LocalDateTime birthDate = LocalDateTime.of(1959, Month.JULY, 19, 0, 0);
    assertEquals(LocalDateTime.of(1991, Month.MARCH, 27, 1, 46, 40), Gigasecond.after(birthDate));
  }

  @Test
  @DisplayName("Fecha y hora")
  public void testDatetime() {
    LocalDateTime birthDate = LocalDateTime.of(2015, Month.JANUARY, 24, 22, 0);
    assertEquals(LocalDateTime.of(2046, Month.OCTOBER, 2, 23, 46, 40), Gigasecond.after(birthDate));
  }

  @Test
  @DisplayName("Otra fecha y hora")
  public void testAnotherDatetime() {
    LocalDateTime birthDate = LocalDateTime.of(2015, Month.JANUARY, 24, 23, 59, 59);
    assertEquals(LocalDateTime.of(2046, Month.OCTOBER, 3, 1, 46, 39), Gigasecond.after(birthDate));
  }
}