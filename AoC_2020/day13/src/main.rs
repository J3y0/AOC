use std::{fs, time::Instant};

pub struct Input {
    time: u32,
    buses: Vec<Option<u32>>,
}

fn main() {
    let content =
        fs::read_to_string("./input/day13.txt").expect("Something went wrong reading the file");

    let input = parse(&content);

    let time = Instant::now();
    let p1 = part1(&input);
    let p2 = part2(&input);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(input: &Input) -> u32 {
    let (wait, id) = input
        .buses
        .iter()
        .filter(|bus| bus.is_some())
        .map(|bus| {
            let b = bus.unwrap();
            (b - input.time % b, b)
        })
        .min_by(|a, b| a.0.cmp(&b.0))
        .unwrap();

    wait * id
}

pub fn part2(input: &Input) -> u64 {
    // Use of chinese remainder theorem to find particular solution
    let n: i64 = input
        .buses
        .iter()
        .filter(|bus| bus.is_some())
        .map(|&b| b.unwrap() as i64)
        .product();

    let mut sol: i64 = 0;
    for (i, &b) in input.buses.iter().enumerate() {
        if let Some(bus) = b {
            let bus_signed = bus as i64;
            let nb = n / bus_signed;
            let ai = -(i as i64) % bus_signed;
            let inv_modulo = extended_euclidean_algorithm(nb, bus_signed).1;
            let coef = ai * nb * inv_modulo;

            sol = (sol + (coef % n + n) % n)%n;
        }
    }

    sol as u64
}

fn parse(input: &str) -> Input {
    let mut lines = input.lines();
    let time: u32 = lines.next().unwrap().parse().unwrap();

    let buses = lines
        .next()
        .unwrap()
        .split(',')
        .map(|b| b.parse().ok())
        .collect();

    Input { time, buses }
}

/// assert_eq!(extended_euclidean_algorithm(55, 79), (1, 23, -16));
//  assert_eq!(extended_euclidean_algorithm(33, 44), (11, -1, 1));
//  assert_eq!(extended_euclidean_algorithm(50, 70), (10, 3, -2));
fn extended_euclidean_algorithm(a: i64, b: i64) -> (i64, i64, i64) {
    fn update_step(a: &mut i64, old_a: &mut i64, quotient: i64) {
        let temp = *a;
        *a = *old_a - quotient * temp;
        *old_a = temp;
    }

    let (mut old_r, mut rem) = (a, b);
    let (mut old_s, mut coeff_s) = (1, 0);
    let (mut old_t, mut coeff_t) = (0, 1);

    while rem != 0 {
        let quotient = old_r / rem;

        update_step(&mut rem, &mut old_r, quotient);
        update_step(&mut coeff_s, &mut old_s, quotient);
        update_step(&mut coeff_t, &mut old_t, quotient);
    }

    (old_r, old_s, old_t)
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Input {
        parse(
            "939
7,13,x,x,59,x,31,19",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 295);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&parse("1234\n17,x,13,19")), 3417);
        assert_eq!(part2(&parse("1234\n67,7,59,61")), 754018);
        assert_eq!(part2(&parse("1234\n67,x,7,59,61")), 779210);
        assert_eq!(part2(&example_data()), 1068781);
    }
}
