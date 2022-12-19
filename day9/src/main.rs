use std::fs;

enum Direction {
    UP,
    DOWN,
    RIGHT,
    LEFT
}

struct Knot {
    tail: (i32, i32),
    head: (i32, i32),
    visited: Vec<(i32, i32)>,
    nb_visited: i32,
}

impl Knot {
    /* (0,0) coordinates are located at the bottom left of the 'imaginary grid' */
    fn knot_move(&mut self, dir: &Direction, nb_steps: u8) {
        /* Move an entire knot */
        for _ in 0..nb_steps {
            match dir {
                Direction::UP => {
                    self.head.0 += 1;
                },
                Direction::DOWN => {
                    self.head.0 -= 1;
                },
                Direction::RIGHT => {
                    self.head.1 += 1;
                },
                Direction::LEFT => {
                    self.head.1 -= 1
                },
            }

            if !self.is_stuck() {
                self.tail_move();
                // Comment those lines for part 2 to gain in efficiency
                // if !self.visited.contains(&self.tail) {
                //     self.nb_visited += 1;
                // }
                // self.visited.push(self.tail);
            }
        }
    }

    fn tail_move(&mut self) {
        /* We know the tail has to move. So both coordinates of tail and head are not equal */
        if self.head.0 > self.tail.0 {
            if self.head.1 == self.tail.1 {
                self.tail.0 += 1;
            } else if self.head.1 > self.tail.1 {
                self.tail.0 += 1;
                self.tail.1 += 1;
            } else {
                self.tail.0 += 1;
                self.tail.1 -= 1;
            }
        } else if self.head.0 < self.tail.0 {
            if self.head.1 == self.tail.1 {
                self.tail.0 -= 1;
            } else if self.head.1 > self.tail.1 {
                self.tail.0 -= 1;
                self.tail.1 += 1;
            } else {
                self.tail.0 -= 1;
                self.tail.1 -= 1;
            }
        } else if self.head.1 > self.tail.1 {
            if self.head.0 == self.tail.0 {
                self.tail.1 += 1;
            // Cases already handled

            // } else if self.head.0 > self.tail.0 {
            //     self.tail.0 += 1;
            //     self.tail.1 += 1;
            // } else {
            //     self.tail.0 -= 1;
            //     self.tail.1 += 1;
            }
        } else {
            if self.head.0 == self.tail.0 {
                self.tail.1 -= 1;
            // Cases already handled

            // } else if self.head.0 > self.tail.0 {
            //     self.tail.0 += 1;
            //     self.tail.1 -= 1;
            // } else {
            //     self.tail.0 -= 1;
            //     self.tail.1 -= 1;
            }
        }
    }

    fn is_stuck(&self) -> bool {
        /* Check whether the head and the tail are stuck together or whether they are tore apart */
        if i32::abs(self.head.0 - self.tail.0) > 1
            || i32::abs(self.head.1 - self.tail.1) > 1 {
                return false;
        }
        return true;
    }
}
 
struct Rope {
    knots: Vec<Knot>,
    visited: Vec<(i32, i32)>,
    nb_visited: i32,
}

impl Rope {
    fn rope_move(&mut self, dir: Direction, nb_steps: u8) {
        for _ in 0..nb_steps {
            self.knots[0].knot_move(&dir, 1);
            for i in 1..self.knots.len() {
                self.knots[i].head = self.knots[i - 1].tail; 
                
                if !self.knots[i].is_stuck() && i != self.knots.len() - 1 {
                    self.knots[i].tail_move(); // Moving the tail of the current knot
                } else if i == self.knots.len() - 1 { // Adding a visited position (last knot doesn't have a tail)
                    if !self.visited.contains(&self.knots[i].head) {
                        self.nb_visited += 1;
                    }
                    self.visited.push(self.knots[i].head);
                }
            }
        }
    }
}

fn main() {
    let content: String = fs::read_to_string("./data/day9.txt")
        .expect("Should have been able to read file");

    let lines: Vec<&str> = content.lines().collect();

    // Part 1
    let mut knot = Knot{tail: (0, 0), head: (0, 0), visited: Vec::new(), nb_visited: 1};
    knot.visited.push((0, 0));

    // Part 2
    let mut rope = Rope{knots: Vec::new(), visited: Vec::new(), nb_visited: 1};
    rope.visited.push((0, 0));
    for _ in 0..10 {
        let mut k = Knot{tail: (0, 0), head: (0, 0), visited: Vec::new(), nb_visited: 1};
        k.visited.push((0, 0));
        rope.knots.push(k);
    }
    
    for line in lines {
        let temp: Vec<&str> = line.split(" ").collect();
        let steps: u8 = temp[1].parse().unwrap();
        let dir: Direction;
        match temp[0] {
            "U" => dir = Direction::UP,
            "D" => dir = Direction::DOWN,
            "R" => dir = Direction::RIGHT,
            _ => dir = Direction::LEFT,
        }

        knot.knot_move(&dir, steps);
        rope.rope_move(dir, steps);
    }

    println!("Part 1: {}", knot.nb_visited);
    println!("Part 1: {}", rope.nb_visited);
}