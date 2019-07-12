pub fn reverse(input: &str) -> String {
    let out: String = input.chars()
         .rev()
         .collect();
    out
}
