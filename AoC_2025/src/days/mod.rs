pub mod day01;

pub trait Solution {
    type Input;

    fn parse(&self, data: &str) -> Self::Input;
    fn part1(&self, input: &Self::Input) -> usize;
    fn part2(&self, input: &Self::Input) -> usize;
}
