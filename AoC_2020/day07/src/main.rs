use regex::Regex;
use std::collections::HashMap;
use std::hash::{DefaultHasher, Hash, Hasher};
use std::{fs, time::Instant};

#[derive(Debug)]
pub struct Bag {
    color: String,
    amount: u32,
}

fn main() {
    let content =
        fs::read_to_string("./input/day07.txt").expect("Something went wrong reading the file");

    let bags = parse(&content);

    let time = Instant::now();
    let p1 = part1(&bags);
    let p2 = part2(&bags);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(bags: &HashMap<u64, Vec<Bag>>) -> usize {
    fn contain_shiny_gold(
        cur: &u64,
        bags: &HashMap<u64, Vec<Bag>>,
        cache: &mut HashMap<u64, bool>,
    ) -> bool {
        if let Some(&cached) = cache.get(cur) {
            return cached;
        }

        let children = bags.get(cur).unwrap();
        if children.len() == 0 {
            return false;
        }

        if children.iter().any(|b| b.color == "shiny gold") {
            cache.insert(*cur, true);
            return true;
        }

        let result = children
            .iter()
            .any(|b| contain_shiny_gold(&hash(&b.color), bags, cache));

        cache.insert(*cur, result);
        return result;
    }

    let mut cache: HashMap<u64, bool> = HashMap::new();

    bags.keys()
        .filter(|&b| contain_shiny_gold(b, bags, &mut cache))
        .count()
}

pub fn part2(bags: &HashMap<u64, Vec<Bag>>) -> u32 {
    fn count_children(
        cur: &u64,
        bags: &HashMap<u64, Vec<Bag>>,
        cache: &mut HashMap<u64, u32>,
    ) -> u32 {
        if let Some(cached) = cache.get(cur) {
            return *cached;
        }

        let children = bags.get(cur).unwrap();
        if children.len() == 0 {
            cache.insert(*cur, 0);
            return 0;
        }

        let count = children.iter().fold(0, |acc, bag| {
            acc + bag.amount + bag.amount * count_children(&hash(&bag.color), bags, cache)
        });

        cache.insert(*cur, count);
        count
    }

    let mut cache: HashMap<u64, u32> = HashMap::new();
    count_children(&hash("shiny gold"), bags, &mut cache)
}

fn parse(input: &str) -> HashMap<u64, Vec<Bag>> {
    let mut h = HashMap::new();
    for line in input.lines() {
        let re = Regex::new(r"^(\w*\s\w*) bags contain (.*).$").unwrap();

        if let Some(cap) = re.captures(line) {
            let (_, [color, content]) = cap.extract();

            if content.contains("no other") {
                h.insert(hash(color), Vec::new());
                continue;
            }

            let mut children = Vec::new();
            let splitted: Vec<_> = content
                .split([',', ' '])
                .filter(|s| !s.is_empty())
                .collect();

            splitted.chunks(4).for_each(|chunk| {
                children.push(Bag {
                    amount: chunk[0].parse().unwrap(),
                    color: chunk[1].to_string() + " " + chunk[2],
                })
            });

            h.insert(hash(color), children);
        }
    }

    h
}

fn hash(val: &str) -> u64 {
    let mut h = DefaultHasher::new();
    val.hash(&mut h);
    h.finish()
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> HashMap<u64, Vec<Bag>> {
        parse(
            "light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 4);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&example_data()), 32);
    }
}
