mod day01;
mod day02;
mod day03;
mod day04;
mod day05;
mod day06;
mod day07;
mod day08;

pub use day01::Day01;
pub use day02::Day02;
pub use day03::Day03;
pub use day04::Day04;
pub use day05::Day05;
pub use day06::Day06;
pub use day07::Day07;
pub use day08::Day08;

pub trait Solution {
    type Input;

    fn parse(data: &str) -> Self::Input;
    fn part1(input: &Self::Input) -> usize;
    fn part2(input: &Self::Input) -> usize;
}
