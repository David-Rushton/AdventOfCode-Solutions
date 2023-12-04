import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class ScratchCard:
    name: str
    winning_numbers: list[int]
    numbers: list[int]
    winners: int
    copies: int
    points: int


def main(is_test_mode: bool) -> None:
    cards = list(get_input(is_test_mode))
    for index in range(len(cards)):
        card = cards[index]
        for winning_number in card.winning_numbers:
            if winning_number in card.numbers:
                card.winners += 1
                if card.points == 0:
                    card.points = 1
                else:
                    card.points = card.points * 2
        if card.points > 0:
            update_copies(cards, index, card.winners, card.copies)

    print(f'Points: {sum(card.points for card in cards)}')
    print(f'Cards: {sum(card.copies for card in cards)}')

def update_copies(cards: list[ScratchCard], start: int, number_of_cards: int, factor: int = 1):
    if start == len(cards):
        return

    for index in range(start + 1, min(start + 1 + number_of_cards, len(cards))):
        cards[index].copies += factor


def get_input(get_test: bool) -> Iterator[ScratchCard]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path

    for line in open(path, 'rt').read().splitlines():
        content = line.split(':')
        numbers = content[1].split('|')
        yield ScratchCard(
            content[0],
            [int(n) for n in numbers[0].split(' ') if n.isdigit()],
            [int(n) for n in numbers[1].split(' ') if n.isdigit()],
            0,
            1,
            0)


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
