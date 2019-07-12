pub fn is_armstrong_number(num: u32) -> bool {
    let i = num.to_string().len() as u32;
    let sum: u32 = num.to_string()
        .chars()
        .map(|n| {
            n.to_digit(10).unwrap().pow(i)
        })
        .sum();
            
    num == sum
}
