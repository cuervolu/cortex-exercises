use space_age::*;

fn assert_in_delta(expected: f64, actual: f64) {
    let diff: f64 = (expected - actual).abs();
    let delta: f64 = 0.01;
    if diff > delta {
        panic!("Your result of {} should be within {} of the expected result {}", actual, delta, expected);
    }
}

#[test]
fn earth_age() {
    let age = SpaceAge::new(1_000_000_000);
    assert_in_delta(31.69, age.on_earth());
}

#[test]
fn mercury_age() {
    let age = SpaceAge::new(2_134_835_688);
    assert_in_delta(280.88, age.on_mercury());
}

#[test]
fn venus_age() {
    let age = SpaceAge::new(189_839_836);
    assert_in_delta(9.78, age.on_venus());
}

#[test]
fn mars_age() {
    let age = SpaceAge::new(2_329_871_239);
    assert_in_delta(39.25, age.on_mars());
}

#[test]
fn jupiter_age() {
    let age = SpaceAge::new(901_876_382);
    assert_in_delta(2.41, age.on_jupiter());
}

#[test]
fn saturn_age() {
    let age = SpaceAge::new(3_000_000_000);
    assert_in_delta(3.23, age.on_saturn());
}

#[test]
fn uranus_age() {
    let age = SpaceAge::new(3_210_123_456);
    assert_in_delta(1.21, age.on_uranus());
}

#[test]
fn neptune_age() {
    let age = SpaceAge::new(8_210_123_456);
    assert_in_delta(1.58, age.on_neptune());
}
