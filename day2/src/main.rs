use std::fs;

#[derive(PartialEq, Debug)]
enum Points {
    ROCK,
    PAPER,
    SCISSORS,
    UNKNOWN
}

struct Object {
    typ: Points,
    lose: Points,
    win: Points
}

impl Object {
    fn win(&self, other: &Object) -> i8 {
        if other.typ == self.typ {
            return 0;
        } else if other.typ == self.lose {
            return -1;
        } else {
            return 1;
        }
    }

    fn compute_score(&self, other: &Object) -> i8 {
        let score: i8 = match self.win(other) {
            1 => 6, // Victory
            0 => 3, // Draw
            -1 => 0, // Lose
            _ => -1 // Default case
        };
        score
    }

    fn which_type(&self, nb_points: i8) -> i32 {
        let mut result: i32 = 0;
        if nb_points == 0 {
            result = match self.win {
                Points::SCISSORS => 3,
                Points::PAPER => 2,
                Points::ROCK => 1,
                Points::UNKNOWN => 0
            };
        } else if nb_points == 3 {
            result = match self.typ {
                Points::SCISSORS => 3,
                Points::PAPER => 2,
                Points::ROCK => 1,
                Points::UNKNOWN => 0
            };
        } else {
            result = match self.lose {
                Points::SCISSORS => 3,
                Points::PAPER => 2,
                Points::ROCK => 1,
                Points::UNKNOWN => 0
            };
        }
        return result;
    }
}

fn main() {
    let content: String = fs::read_to_string("./data/day2.txt")
        .expect("Should have been able to read the file");
    let lines: Vec<&str> = content.lines().collect();
    let mut part1: i32 = 0;
    let mut part2: i32 = 0;

    for line in lines {
        let round: Vec<&str> = line.split(" ").collect();

        let p1 = match round[0] {
            "A" => Object{typ: Points::ROCK, lose: Points::PAPER, win: Points::SCISSORS},
            "B" => Object{typ: Points::PAPER, lose: Points::SCISSORS, win: Points::ROCK},
            "C" => Object{typ: Points::SCISSORS, lose: Points::ROCK, win: Points::PAPER},
            _ => Object{typ: Points::UNKNOWN, lose: Points::UNKNOWN, win: Points::UNKNOWN},
        };

        // Part 1
        let p2 = match round[1] {
            "X" => Object{typ: Points::ROCK, lose: Points::PAPER, win: Points::SCISSORS},
            "Y" => Object{typ: Points::PAPER, lose: Points::SCISSORS, win: Points::ROCK},
            "Z" => Object{typ: Points::SCISSORS, lose: Points::ROCK, win: Points::PAPER},
            _ => Object{typ: Points::UNKNOWN, lose: Points::UNKNOWN, win: Points::UNKNOWN},
        };
        let outcome: i32 = i32::from(p2.compute_score(&p1));
        match p2.typ {
            Points::ROCK => part1 += 1 + outcome,
            Points::PAPER => part1 += 2 + outcome,
            Points::SCISSORS => part1 += 3 + outcome,
            Points::UNKNOWN => part1 += 0,
        };
        // Part 2
        let p2: i8 = match round[1] {
                "X" => 0,
                "Y" => 3,
                "Z" => 6,
                _ => -1
            };
        let new_type: i32 = p1.which_type(p2);
        part2 += i32::from(p2) + new_type;        
    }

    println!("Part 1: {part1}");
    println!("Part 2: {part2}");
}