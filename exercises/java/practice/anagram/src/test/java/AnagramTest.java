import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;
import java.util.Arrays;
import java.util.List;

public class AnagramTest {
  @Test
  public void testNoMatches() {
    Anagram detector = new Anagram("diaper");
    assertTrue(detector.match(Arrays.asList("hello", "world", "zombies", "pants")).isEmpty());
  }

  @Test
  @Disabled("Remove to run test")
  public void testDetectMultipleAnagrams() {
    Anagram detector = new Anagram("master");
    List<String> expected = Arrays.asList("stream", "maters");
    List<String> actual = detector.match(Arrays.asList("stream", "pigeon", "maters"));
    assertAll(
        () -> assertEquals(expected.size(), actual.size()),
        () -> assertTrue(actual.containsAll(expected) && expected.containsAll(actual))
    );
  }

  @Test
  @Disabled("Remove to run test")
  public void testEliminateAnagramSubsets() {
    Anagram detector = new Anagram("good");
    assertTrue(detector.match(Arrays.asList("dog", "goody")).isEmpty());
  }

  @Test
  @Disabled("Remove to run test")
  public void testDetectLongerAnagram() {
    Anagram detector = new Anagram("listen");
    List<String> result = detector.match(Arrays.asList("enlists", "google", "inlets", "banana"));
    assertAll(
        () -> assertEquals(1, result.size()),
        () -> assertTrue(result.contains("inlets"))
    );
  }

  @Test
  @Disabled("Remove to run test")
  public void testDetectMultipleAnagramsForLongerWord() {
    Anagram detector = new Anagram("allergy");
    List<String> expected = Arrays.asList("gallery", "regally", "largely");
    List<String> actual = detector.match(Arrays.asList("gallery", "ballerina", "regally", "clergy", "largely", "leading"));
    assertAll(
        () -> assertEquals(expected.size(), actual.size()),
        () -> assertTrue(actual.containsAll(expected) && expected.containsAll(actual))
    );
  }

  @Test
  @Disabled("Remove to run test")
  public void testDetectsMultipleAnagramsWithDifferentCase() {
    Anagram detector = new Anagram("nose");
    List<String> expected = Arrays.asList("Eons", "ONES");
    List<String> actual = detector.match(Arrays.asList("Eons", "ONES"));
    assertAll(
        () -> assertEquals(expected.size(), actual.size()),
        () -> assertTrue(actual.containsAll(expected) && expected.containsAll(actual))
    );
  }

  @Test
  @Disabled("Remove to run test")
  public void testEliminateAnagramsWithSameChecksum() {
    Anagram detector = new Anagram("mass");
    assertTrue(detector.match(Arrays.asList("last")).isEmpty());
  }

  @Test
  @Disabled("Remove to run test")
  public void testCaseInsensitiveWhenBothAnagramAndSubjectStartWithUpperCaseLetter() {
    Anagram detector = new Anagram("Orchestra");
    List<String> result = detector.match(Arrays.asList("cashregister", "Carthorse", "radishes"));
    assertAll(
        () -> assertEquals(1, result.size()),
        () -> assertTrue(result.contains("Carthorse"))
    );
  }

  @Test
  @Disabled("Remove to run test")
  public void testCaseInsensitiveWhenSubjectStartsWithUpperCaseLetter() {
    Anagram detector = new Anagram("Orchestra");
    List<String> result = detector.match(Arrays.asList("cashregister", "carthorse", "radishes"));
    assertAll(
        () -> assertEquals(1, result.size()),
        () -> assertTrue(result.contains("carthorse"))
    );
  }

  @Test
  @Disabled("Remove to run test")
  public void testCaseInsensitiveWhenAnagramStartsWithUpperCaseLetter() {
    Anagram detector = new Anagram("orchestra");
    List<String> result = detector.match(Arrays.asList("cashregister", "Carthorse", "radishes"));
    assertAll(
        () -> assertEquals(1, result.size()),
        () -> assertTrue(result.contains("Carthorse"))
    );
  }

  @Test
  @Disabled("Remove to run test")
  public void testIdenticalWordRepeatedIsNotAnagram() {
    Anagram detector = new Anagram("go");
    assertTrue(detector.match(List.of("go Go GO")).isEmpty());
  }

  @Test
  @Disabled("Remove to run test")
  public void testAnagramMustUseAllLettersExactlyOnce() {
    Anagram detector = new Anagram("tapper");
    assertTrue(detector.match(List.of("patter")).isEmpty());
  }

  @Test
  @Disabled("Remove to run test")
  public void testWordsAreNotAnagramsOfThemselvesCaseInsensitive() {
    Anagram detector = new Anagram("BANANA");
    assertTrue(detector.match(Arrays.asList("BANANA", "Banana", "banana")).isEmpty());
  }

  @Test
  @Disabled("Remove to run test")
  public void testWordsOtherThanThemselvesCanBeAnagrams() {
    Anagram detector = new Anagram("LISTEN");
    List<String> result = detector.match(Arrays.asList("Listen", "Silent", "LISTEN"));
    assertAll(
        () -> assertEquals(1, result.size()),
        () -> assertTrue(result.contains("Silent"))
    );
  }
}