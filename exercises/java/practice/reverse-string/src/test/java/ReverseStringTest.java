import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class ReverseStringTest {

  @Test
  @DisplayName("Cadena vacía")
  public void testEmptyString() {
    assertEquals("", ReverseString.reverse(""));
  }

  @Test
  @DisplayName("Palabra")
  public void testWord() {
    assertEquals("tobor", ReverseString.reverse("robot"));
  }

  @Test
  @DisplayName("Frase con espacios")
  public void testCapitalisedWord() {
    assertEquals("gnidoc evol I", ReverseString.reverse("I love coding"));
  }

  @Test
  @DisplayName("Frase con puntuación")
  public void testSentenceWithPunctuation() {
    assertEquals("!dlroW ,olleH", ReverseString.reverse("Hello, World!"));
  }

  @Test
  @DisplayName("Palíndromo")
  public void testPalindrome() {
    assertEquals("racecar", ReverseString.reverse("racecar"));
  }
}