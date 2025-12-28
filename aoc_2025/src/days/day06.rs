use crate::days::Solution;

pub struct Day06;

#[derive(Debug)]
enum Operator {
    Add,
    Mul,
    Undefined,
}

impl Operator {
    fn get_init_val(&self) -> usize {
        match self {
            Operator::Add => 0,
            Operator::Mul => 1,
            Operator::Undefined => 0,
        }
    }
}

impl From<&str> for Operator {
    fn from(value: &str) -> Self {
        match value {
            "+" => Self::Add,
            "*" => Self::Mul,
            _ => Self::Undefined,
        }
    }
}

#[derive(Debug)]
pub struct Calculation {
    symbols: Vec<Operator>,
    lines: Vec<String>,
}

impl Solution for Day06 {
    type Input = Calculation;

    fn parse(data: &str) -> Self::Input {
        let mut lines = data.lines().rev();
        let symbols = lines.next().unwrap();
        let symbols: Vec<Operator> = symbols
            .split_ascii_whitespace()
            .filter(|s| !s.is_empty())
            .map(Operator::from)
            .collect();

        Calculation {
            lines: lines.rev().map(String::from).collect(),
            symbols,
        }
    }

    fn part1(input: &Self::Input) -> usize {
        // Parse for part1
        let numbers: Vec<Vec<usize>> = input
            .lines
            .iter()
            .map(|s| {
                s.split_ascii_whitespace()
                    .filter(|s| !s.is_empty())
                    .map(|s| s.parse::<usize>().unwrap())
                    .collect()
            })
            .collect();

        assert_eq!(numbers[0].len(), input.symbols.len());

        let mut tot = 0;
        for c in 0..numbers[0].len() {
            let symbol = &input.symbols[c];
            let mut res = symbol.get_init_val();
            // Compute for the entire column
            for nb in &numbers {
                match symbol {
                    Operator::Add => res += nb[c],
                    Operator::Mul => res *= nb[c],
                    Operator::Undefined => {}
                }
            }

            tot += res;
        }

        tot
    }

    fn part2(input: &Self::Input) -> usize {
        // Find the line 'real' length as the end of line has been trimmed somehow
        // The largest one is the one with the full-length number, all others
        // should be pad with a whitespace for instance
        let mut max_len = 0;
        for l in &input.lines {
            max_len = max_len.max(l.len());
        }

        // Parse vertical numbers
        let mut numbers = Vec::with_capacity(max_len);
        for c in 0..max_len {
            let mut col_nb: usize = 0;
            for r in 0..input.lines.len() {
                // Pad if whitespace has been stripped
                let n = input.lines[r].as_bytes().get(c).unwrap_or(&b' ');
                if let b'0'..=b'9' = n {
                    col_nb = col_nb * 10 + (n - b'0') as usize;
                }
            }
            // numbers for one operation are separated by 0 (column with whitespaces only)
            numbers.push(col_nb);
        }

        let mut tot = 0;
        let mut op_idx = 0;
        let mut cur = input.symbols[op_idx].get_init_val();
        for n in numbers {
            if n == 0 {
                tot += cur;
                op_idx += 1;
                cur = input.symbols[op_idx].get_init_val();
                continue;
            }

            let symbol = &input.symbols[op_idx];
            match symbol {
                Operator::Add => cur += n,
                Operator::Mul => cur *= n,
                Operator::Undefined => {}
            }
        }

        tot += cur;

        tot
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Calculation {
        Day06::parse(
            "123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  ",
        )
    }

    #[test]
    fn part1_test() {
        assert_eq!(Day06::part1(&example_data()), 4277556);
    }

    #[test]
    fn part2_test() {
        assert_eq!(Day06::part2(&example_data()), 3263827);
    }
}
