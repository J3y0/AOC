use std::collections::HashMap;

use crate::days::Solution;

pub type Point = (usize, usize);

pub struct Day09;

impl Solution for Day09 {
    type Input = Vec<Point>;

    fn parse(data: &str) -> Self::Input {
        data.lines()
            .map(|l| {
                let (c, r) = l.split_once(',').unwrap();
                let c = c.parse().unwrap();
                let r = r.parse().unwrap();

                (r, c)
            })
            .collect()
    }

    fn part1(input: &Self::Input) -> usize {
        let mut max_area = usize::MIN;
        for i in 0..input.len() {
            for j in (i + 1)..input.len() {
                max_area = max_area.max(area(&input[i], &input[j]));
            }
        }

        max_area
    }

    fn part2(input: &Self::Input) -> usize {
        let mut xs = vec![];
        let mut ys = vec![];
        for p in input {
            xs.push(p.0);
            ys.push(p.1);
        }
        xs.sort_unstable();
        xs.dedup();
        ys.sort_unstable();
        ys.dedup();

        let x_lookup: HashMap<usize, usize> =
            xs.iter().enumerate().map(|(i, &x)| (x, 2 * i)).collect();

        let y_lookup: HashMap<usize, usize> =
            ys.iter().enumerate().map(|(i, &y)| (y, 2 * i)).collect();

        let mut cgrid = compressed_grid(input, &x_lookup, &y_lookup);
        // Compute sum_cgrid: each index holds the sum of the region at its top left
        //
        //  ///// |
        //  ///// |
        //  ///// |
        // -------i - -
        //        |
        for i in 0..cgrid.len() {
            for j in 0..cgrid[i].len() {
                let top = if i > 0 { cgrid[i - 1][j] } else { 0 };
                let left = if j > 0 { cgrid[i][j - 1] } else { 0 };
                let topleft = if i > 0 && j > 0 {
                    cgrid[i - 1][j - 1]
                } else {
                    0
                };
                cgrid[i][j] = cgrid[i][j] + top + left - topleft;
            }
        }

        let mut max_valid = usize::MIN;
        for i in 0..input.len() {
            for j in (i + 1)..input.len() {
                let p1 = &input[i];
                let p2 = &input[j];
                let ((cx1, cy1), (cx2, cy2)) = compressed_coords(p1, p2, &x_lookup, &y_lookup);

                let rec_left = if cy1 > 0 { cgrid[cx2][cy1 - 1] } else { 0 };
                let rec_top = if cx1 > 0 { cgrid[cx1 - 1][cy2] } else { 0 };
                let rec_topleft = if cx1 > 0 && cy1 > 0 {
                    cgrid[cx1 - 1][cy1 - 1]
                } else {
                    0
                };
                let csum = cgrid[cx2][cy2] + rec_topleft - rec_left - rec_top;

                if area(&(cx1, cy1), &(cx2, cy2)) == csum {
                    max_valid = max_valid.max(area(p1, p2));
                }
            }
        }

        max_valid
    }
}

fn area(p1: &Point, p2: &Point) -> usize {
    let r = p1.0.max(p2.0) - p1.0.min(p2.0) + 1;
    let c = p1.1.max(p2.1) - p1.1.min(p2.1) + 1;

    r * c
}

fn compressed_coords(
    (x1, y1): &Point,
    (x2, y2): &Point,
    xs: &HashMap<usize, usize>,
    ys: &HashMap<usize, usize>,
) -> (Point, Point) {
    // compressed index of x1
    let cx1 = xs[x1];
    let cx2 = xs[x2];
    let cxmin = cx1.min(cx2);
    let cxmax = cx1.max(cx2);

    // compressed index of y1
    let cy1 = ys[y1];
    let cy2 = ys[y2];
    let cymin = cy1.min(cy2);
    let cymax = cy1.max(cy2);

    ((cxmin, cymin), (cxmax, cymax))
}

fn compressed_grid(
    points: &[Point],
    xs: &HashMap<usize, usize>,
    ys: &HashMap<usize, usize>,
) -> Vec<Vec<usize>> {
    // compressed grid - 1 is red or green tile, 0 is outside
    let mut cgrid = vec![vec![0; ys.len() * 2 - 1]; xs.len() * 2 - 1];

    // fill polygon sides
    let points_len = points.len();
    for i in 0..points_len {
        let p1 = &points[i];
        let p2 = &points[(i + 1) % points_len];
        let ((cx1, cy1), (cx2, cy2)) = compressed_coords(p1, p2, xs, ys);
        for row in cgrid.iter_mut().take(cx2 + 1).skip(cx1) {
            for elt in row.iter_mut().take(cy2 + 1).skip(cy1) {
                *elt = 1;
            }
        }
    }

    // fill within polygon
    for row in &mut cgrid {
        let mut start = 0;
        let mut out = true;
        for j in 0..row.len() {
            if row[j] == 0 {
                continue;
            }
            if out {
                start = j + 1;
                out = false;
            } else {
                out = true;
                for within in row.iter_mut().take(j).skip(start) {
                    *within = 1;
                }
            }
        }
    }

    cgrid
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<Point> {
        Day09::parse(
            "7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3",
        )
    }

    #[test]
    fn part1_test() {
        assert_eq!(Day09::part1(&example_data()), 50);
    }

    #[test]
    fn part2_test() {
        assert_eq!(Day09::part2(&example_data()), 24);
    }
}
