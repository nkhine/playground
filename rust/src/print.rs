pub fn run() {
    // print to console
    println!("Hello, world!");
    // basic formatting
    println!("{} is from {}", "Norman", "Sofia");
    // positional arguments
    println!(
        "{0} is the capital of {1} which is in {2}",
        "Sofia", "Bulgaria", "Europe"
    );
    // Named arguments
    println!(
        "{name} likes to {activity}",
        name = "Norman",
        activity = "sail"
    );
    // placeholder traits
    println!("Binary: {:b} Hex: {:x} Octal: {:o}", 10, 10, 10);

    // placeholder for debug traits
    // returns a tuple
    println!("{:?}", (12, true, "hello"));

    // basic math
    println!("10 + 10 = {}", 10 + 10);
}
