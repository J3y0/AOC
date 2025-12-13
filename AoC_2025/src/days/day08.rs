use crate::days::Solution;

pub struct Day08;

#[derive(Debug)]
pub struct Pos {
    x: usize,
    y: usize,
    z: usize,
}

impl Pos {
    fn distance(&self, p2: &Pos) -> usize {
        // No need to compute sqrt -> more computation for nothing -> enough to order items
        self.x.abs_diff(p2.x).pow(2) + self.y.abs_diff(p2.y).pow(2) + self.z.abs_diff(p2.z).pow(2)
    }
}

impl From<&str> for Pos {
    fn from(value: &str) -> Self {
        let coords: Vec<usize> = value
            .splitn(3, ',')
            .map(|n| n.trim().parse::<usize>().unwrap())
            .collect();
        let (x, y, z) = (coords[0], coords[1], coords[2]);
        Self { x, y, z }
    }
}

// Disjoint Set Union with rank optimisation
// reason with indexes of Pos
#[derive(Debug)]
struct DSU {
    parents: Vec<usize>,
    sizes: Vec<u32>,
}

impl DSU {
    fn new(n: usize) -> Self {
        Self {
            parents: (0..n).collect(),
            sizes: vec![1u32; n],
        }
    }

    fn find_root(&self, i: usize) -> usize {
        if self.parents[i] == i {
            i
        } else {
            self.find_root(self.parents[i])
        }
    }

    fn union(&mut self, i1: usize, i2: usize) {
        let r1 = self.find_root(i1);
        let r2 = self.find_root(i2);
        if r1 != r2 {
            if self.sizes[r1] < self.sizes[r2] {
                self.sizes.swap(r1, r2);
            }
            self.parents[r2] = r1;
            self.sizes[r1] += self.sizes[r2];
        }
    }

    fn is_one_set(&self) -> bool {
        let max_size = self.sizes.iter().max().unwrap_or(&0);
        *max_size == self.sizes.len() as u32
    }
}

pub struct Boxes {
    positions: Vec<Pos>,
    distances: Vec<(usize, usize, usize)>,
}

impl Solution for Day08 {
    type Input = Boxes;

    fn parse(data: &str) -> Self::Input {
        let positions: Vec<Pos> = data.lines().map(|l| Pos::from(l)).collect();

        // Precompute all distances:
        //   - distance first to sort the elements based on it
        //   - index of first pos
        //   - index of second pos
        let mut distances: Vec<(usize, usize, usize)> = Vec::new();

        for i in 0..positions.len() {
            for j in (i + 1)..positions.len() {
                let p1 = &positions[i];
                let p2 = &positions[j];
                let dist = p1.distance(p2);
                distances.push((dist, i, j));
            }
        }
        distances.sort_unstable();

        Boxes {
            positions,
            distances,
        }
    }

    fn part1(input: &Self::Input) -> usize {
        let mut sizes = run_dsu(&input, 1000);
        sizes.sort_unstable();

        sizes.iter().rev().take(3).product::<u32>() as usize
    }

    fn part2(input: &Self::Input) -> usize {
        let mut dsu = DSU::new(input.positions.len());
        let mut i = 0;
        let mut res = 0;

        loop {
            let (_, i1, i2) = input.distances[i];
            dsu.union(i1, i2);

            // Don't check for max if all positions are not added at least one
            if i > input.positions.len() && dsu.is_one_set() {
                res = input.positions[i1].x * input.positions[i2].x;
                break;
            }

            i += 1;
        }

        res
    }
}

// Inner part 1 for testing (number of steps differs)
fn run_dsu(boxes: &Boxes, steps: usize) -> Vec<u32> {
    let mut dsu = DSU::new(boxes.positions.len());
    for k in 0..steps {
        let (_, i1, i2) = boxes.distances[k];
        dsu.union(i1, i2);
    }

    dsu.sizes
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Boxes {
        Day08::parse(
            "162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689",
        )
    }

    #[test]
    fn part1_test() {
        let boxes = &example_data();
        let mut sizes = run_dsu(&boxes, 10);
        sizes.sort_unstable();

        let res = sizes.iter().rev().take(3).product::<u32>();
        assert_eq!(res, 40);
    }

    #[test]
    fn part2_test() {
        assert_eq!(Day08::part2(&example_data()), 25272);
    }
}
