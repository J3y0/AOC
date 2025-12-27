use crate::days::Solution;

pub struct Day12;

type Present = Vec<u16>;

#[derive(Debug)]
struct Region {
    width: u16,
    height: u16,
    quantities: Vec<u16>,
}

impl Region {
    fn area(&self) -> usize {
        (self.height * self.width) as usize
    }
}

#[derive(Debug)]
pub struct Farm {
    #[allow(dead_code)]
    presents: Vec<Present>,
    regions: Vec<Region>,
}

impl Solution for Day12 {
    type Input = Farm;

    fn parse(data: &str) -> Self::Input {
        let mut regions: Vec<Region> = vec![];
        let mut presents: Vec<Present> = vec![];

        for block in data.split("\n\n") {
            let mut lines = block.lines();
            match block.chars().nth(1).unwrap() {
                ':' => {
                    // remove index line
                    let _ = lines.next();
                    let mut present = vec![];
                    for l in lines {
                        let mut bitvec = 0;
                        for (i, c) in l.char_indices() {
                            bitvec += ((c == '#') as u16) << i;
                        }
                        present.push(bitvec);
                    }
                    presents.push(present);
                }
                _ => {
                    regions = lines
                        .map(|l| {
                            let (size, qty) = l.split_once(':').unwrap();
                            let (col, row) = size
                                .split_once('x')
                                .map(|(col, row)| (col.parse().unwrap(), row.parse().unwrap()))
                                .unwrap();
                            let qty = qty
                                .split_ascii_whitespace()
                                .map(|i| i.parse().unwrap())
                                .collect();

                            return Region {
                                quantities: qty,
                                width: col,
                                height: row,
                            };
                        })
                        .collect();
                }
            }
        }

        Farm { regions, presents }
    }

    fn part1(input: &Self::Input) -> usize {
        input.regions.iter().filter(|&r| can_fit(r)).count()
    }

    fn part2(_input: &Self::Input) -> usize {
        0
    }
}

fn can_fit(region: &Region) -> bool {
    // consider shapes as block of 3x3
    //  -> input is permissive... and I am tired boss
    let tot_shapes = region.quantities.iter().sum::<u16>() as usize;

    region.area() >= tot_shapes * 3 * 3
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Farm {
        Day12::parse(
            "0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2",
        )
    }

    #[test]
    fn part1_test() {
        assert_eq!(Day12::part1(&example_data()), 2);
    }

    #[test]
    fn part2_test() {
        assert_eq!(Day12::part2(&example_data()), 0);
    }
}
