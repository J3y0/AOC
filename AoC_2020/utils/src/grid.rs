use crate::point::Point;
use std::ops::Index;

#[derive(Clone)]
pub struct Grid {
    pub grid: Vec<u8>,
    pub height: i32,
    pub width: i32,
}

impl Grid {
    pub fn parse(input: &str) -> Grid {
        let raw: Vec<_> = input.lines().map(str::as_bytes).collect();
        let width = raw[0].len() as i32;
        let height = raw.len() as i32;

        let mut grid = Vec::with_capacity((height * width) as usize);
        raw.iter().for_each(|&l| grid.extend_from_slice(l));

        Grid {
            grid,
            height,
            width,
        }
    }

    pub fn print(&self) {
        for i in 0..self.height {
            for j in 0..self.width {
                print!("{}", self[Point { x: i, y: j }] as char);
            }
            println!();
        }
    }

    pub fn outside(&self, p: &Point) -> bool {
        if p.x < 0 || p.y < 0 {
            return true;
        }

        if p.x >= self.height || p.y >= self.width {
            return true;
        }

        false
    }

    pub fn len(&self) -> usize {
        (self.height * self.width) as usize
    }
}

impl Index<Point> for Grid {
    type Output = u8;

    fn index(&self, index: Point) -> &Self::Output {
        self.grid.index((index.x * self.width + index.y) as usize)
    }
}
