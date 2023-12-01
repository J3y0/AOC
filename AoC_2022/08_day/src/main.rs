use std::fs;

enum Direction {
    UP,
    DOWN,
    RIGHT,
    LEFT
}
struct Place {
    grid: Vec<Vec<i8>>,
    nb_arbre_visible: i32,
    by_dir_score: [usize; 4],
}

impl Place {
    fn check_direction(&mut self, pos: (usize, usize), depth: usize, dir: Direction, temp_score: &mut [usize; 4]) -> bool { 
        match dir {
            Direction::UP => {
                if pos.0 - depth == 0 {
                    temp_score[0] = depth;
                    return self.grid[0][pos.1] < self.grid[pos.0][pos.1];
                } else {
                    if self.grid[pos.0 - depth][pos.1] < self.grid[pos.0][pos.1] {
                        return self.check_direction(pos, depth + 1, Direction::UP, temp_score);
                    } else {
                        temp_score[0] = depth;
                        return false;
                    }
                }
            },
            Direction::DOWN => {
                if pos.0 + depth == self.grid.len() - 1 {
                    temp_score[1] = depth;
                    return self.grid[self.grid.len() - 1][pos.1] < self.grid[pos.0][pos.1];
                } else {
                    if self.grid[pos.0 + depth][pos.1] < self.grid[pos.0][pos.1] {
                        return self.check_direction(pos, depth + 1, Direction::DOWN, temp_score);
                    } else {
                        temp_score[1] = depth;
                        return false;
                    }
                }
            },
            Direction::LEFT => {
                if pos.1 - depth == 0 {
                    temp_score[2] = depth;
                    return self.grid[pos.0][0] < self.grid[pos.0][pos.1];
                } else {
                    if self.grid[pos.0][pos.1 - depth] < self.grid[pos.0][pos.1] {
                        return self.check_direction(pos, depth + 1, Direction::LEFT, temp_score);
                    } else {
                        temp_score[2] = depth;
                        return false;
                    }
                }
            },
            Direction::RIGHT => {
                if pos.1 + depth == self.grid[0].len() - 1 {
                    temp_score[3] = depth;
                    return self.grid[pos.0][self.grid[0].len() - 1] < self.grid[pos.0][pos.1];
                } else {
                    if self.grid[pos.0][pos.1 + depth] < self.grid[pos.0][pos.1] {
                        return self.check_direction(pos, depth + 1, Direction::RIGHT, temp_score);
                    } else {
                        temp_score[3] = depth;
                        return false;
                    }
                }
            }
        }
    }

    fn visible(&mut self, pos: (usize, usize)) -> bool {
        if pos.0 == 0 || pos.0 == self.grid.len() - 1 || pos.1 == 0 || pos.1 == self.grid[0].len() - 1 {
            return true;
        }

        let mut temp_score: [usize; 4] = [1; 4];
        let down: bool = self.check_direction(pos, 1, Direction::DOWN, &mut temp_score);
        let up: bool = self.check_direction(pos, 1, Direction::UP, &mut temp_score);
        let left: bool = self.check_direction(pos, 1, Direction::LEFT, &mut temp_score);
        let right: bool = self.check_direction(pos, 1, Direction::RIGHT, &mut temp_score);
        if down || up || left || right {
                if temp_score.iter().fold(1, |res, x| x * res) > self.by_dir_score.iter().fold(1, |res, x| x * res) {
                    self.by_dir_score = temp_score;
                }
                return true;
        }
        return false;
    }
}

fn main() {
    let content: String = fs::read_to_string("./data/day8.txt")
        .expect("Should have been able to read file");
    let lines: Vec<&str> = content.lines().collect();

    let mut grid: Vec<Vec<i8>> = Vec::new();
    let part2: usize;

    for i in 0..lines.len() {
        let mut l: Vec<i8> = Vec::new();
        for j in 0..lines[i].len() {
            let value = lines[i].chars().nth(j).unwrap().to_digit(10).unwrap() as i8;
            l.push(value);
        }
        grid.push(l);
    }

    let mut place = Place{grid: grid, nb_arbre_visible: 0, by_dir_score: [1;4]};

    for i in 0..place.grid.len() {
        for j in 0..place.grid[0].len() {
            if place.visible((i, j)) {
                place.nb_arbre_visible += 1;
            }
        }
    }
    part2 = place.by_dir_score.iter().fold(1, |res, x| x * res);
    println!("Part 1: {}", place.nb_arbre_visible);
    println!("Part 2: {}", part2);
}