use std::{fs, time::Instant};

#[derive(Copy, Clone, Debug)]
struct Range {
    start: usize,
    end: usize,
}

impl From<&str> for Range {
    fn from(value: &str) -> Range {
        let values: Vec<usize> = value.split('-').map(|e| e.parse().unwrap()).collect();
        Range {
            start: values[0],
            end: values[1],
        }
    }
}

impl Range {
    fn new(start: usize, end: usize) -> Range {
        Range { start, end }
    }

    fn overlap(self, other: &Range) -> bool {
        (self.start <= other.start && other.start <= self.end)
            || (self.start <= other.end && other.end <= self.end)
    }

    fn merge(self, other: &Range) -> Range {
        Range {
            start: self.start.min(other.start),
            end: self.end.max(other.end),
        }
    }

    fn contains(self, value: usize) -> bool {
        self.start <= value && value <= self.end
    }
}

fn merge_ranges(ranges: &mut Vec<Range>) -> Vec<Range> {
    ranges.sort_by_key(|r| r.start);

    let mut cur = ranges[0];
    let mut merged = vec![];
    for r in ranges.iter().skip(1) {
        if cur.overlap(r) {
            cur = cur.merge(r);
        } else {
            merged.push(cur);
            cur = *r;
        }
    }
    merged.push(cur);

    merged
}

type Ticket = Vec<usize>;

#[derive(Debug)]
struct Input {
    ranges: Vec<Range>,
    ticket: Ticket,
    nearby: Vec<Ticket>,
}

fn main() {
    let content =
        fs::read_to_string("./input/day16.txt").expect("Something went wrong reading the file");

    let input = parse(&content);

    let time = Instant::now();
    let p1 = part1(&input);
    let p2 = part2(&input);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(input: &Input) -> usize {
    let mut cloned_ranges = input.ranges.clone();
    let merged = merge_ranges(&mut cloned_ranges);

    println!("{:?}", merged);
    let mut sum = 0;
    for t in &input.nearby {
        for &val in t {
            if !merged.iter().any(|rg| rg.contains(val)) {
                sum += val;
            }
        }
    }

    sum
}

#[allow(unused)]
pub fn part2(input: &Input) -> usize {
    0
}

fn parse(input: &str) -> Input {
    let [first, second, third] = input
        .splitn(3, "\n\n")
        .collect::<Vec<&str>>()
        .try_into()
        .unwrap();

    // Collect ranges from first part
    let mut ranges = Vec::new();
    for l in first.lines() {
        let rule = l.split(": ").skip(1).next().unwrap();
        ranges.extend(rule.split(" or ").map(Range::from));
    }

    // Collect ticket from second part
    let ticket = second.lines().skip(1).next().unwrap();
    let yours: Ticket = ticket.split(',').map(|e| e.parse().unwrap()).collect();

    // Collect nearby tickets from second part
    let nearby_tickets: Vec<Ticket> = third
        .lines()
        .skip(1)
        .map(|l| l.split(',').map(|e| e.parse().unwrap()).collect::<Ticket>())
        .collect();

    Input {
        ranges: ranges,
        ticket: yours,
        nearby: nearby_tickets,
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Input {
        parse(
            "class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 71);
    }
}
