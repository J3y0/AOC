use std::fs;

fn main() {
    let content: String = fs::read_to_string("./data/day10.txt")
        .expect("Should have been able to read file");

    let lines: Vec<&str> = content.lines().collect();
    // let mut signal_strengths: [i32; 6] = [0; 6];
    let mut register: i32 = 1;
    let mut clock: usize = 0;
    
    for line in lines {
        if line.contains("noop") {
            clock += 1;
            // Part 1
            // if clock % 40 == 20 {
            //     let index: usize = clock/40;
            //     signal_strengths[index] = compute_power(register, index);
            // }

            // Part 2
            // The drawer is clock - 1 (starts at 0)
            if (clock as i32 - 1) % 40 >= register - 1 && (clock as i32 - 1) % 40<= register + 1 { 
                print!("#");
            } else {
                print!(".");
            }
            if clock % 40 == 0 {
                println!();
            }

        } else {
            let mut i: i32 = 0;
            while i < 2 {
                i += 1;
                clock += 1;
                // Part 1
                // if clock % 40 == 20 {
                //     let index: usize = clock/40;
                //     signal_strengths[index] = compute_power(register, index);
                // }

                // Part 2
                if (clock as i32 - 1) % 40 >= register - 1 && (clock as i32 - 1) % 40 <= register + 1 {
                    print!("#");
                } else {
                    print!(".");
                }
                if clock % 40 == 0 {
                    println!();
                }
            }
            let temp: Vec<&str> = line.split(" ").collect();
            let value: i32 = temp[1].parse().unwrap();
            register += value;
        }
    }

    // println!("Part 1: {}", signal_strengths.iter().sum::<i32>());
}

// Part 1
// fn compute_power(register_value: i32, nb_cycle: usize) -> i32 {
//     let nb_c:i32 = nb_cycle as i32;
//     return register_value * (nb_c*40 + 20);
// }